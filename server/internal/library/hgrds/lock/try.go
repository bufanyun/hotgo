package lock

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

type memEntry struct {
	mu   sync.Mutex
	held bool
}

var memEntries sync.Map

// Try 尝试获取命名锁。ok 为 false 表示没拿到。release 在拿到锁之后调用，用来释放。
func Try(ctx context.Context, name string) (release func(), ok bool) {
	switch currentDriver(ctx) {
	case "redis":
		return tryRedis(ctx, name)
	case "memory":
		return tryMemory(name)
	case "pgsql", "postgres", "postgresql":
		return tryPgsql(ctx, name)
	case "mysql":
		return tryMysql(ctx, name)
	default:
		return tryMemory(name)
	}
}

// currentDriver 决定用哪把锁。显式配置优先
func currentDriver(ctx context.Context) string {
	if d := strings.TrimSpace(g.Cfg().MustGet(ctx, "lock.driver").String()); d != "" {
		return strings.ToLower(d)
	}
	dbType := strings.TrimSpace(g.Cfg().MustGet(ctx, "database.default.type").String())
	if dbType == "" {
		link := g.Cfg().MustGet(ctx, "database.default.link").String()
		if i := strings.IndexByte(link, ':'); i > 0 {
			dbType = link[:i]
		}
	}
	return strings.ToLower(dbType)
}

// tryMemory 进程内互斥。第二个同名 Try 在释放前会失败。
func tryMemory(name string) (func(), bool) {
	actual, _ := memEntries.LoadOrStore(name, &memEntry{})
	entry := actual.(*memEntry)
	entry.mu.Lock()
	defer entry.mu.Unlock()
	if entry.held {
		return func() {}, false
	}
	entry.held = true
	return func() {
		entry.mu.Lock()
		entry.held = false
		entry.mu.Unlock()
	}, true
}

// tryRedis 使用 Redis 锁。未配置地址时直接失败，避免误连。
func tryRedis(ctx context.Context, name string) (func(), bool) {
	addr := strings.TrimSpace(g.Cfg().MustGet(ctx, "redis.default.address").String())
	if addr == "" {
		g.Log().Warning(ctx, "lock.driver=redis 但未配置 redis.default.address")
		return func() {}, false
	}
	l := Mutex(name)
	if err := l.TryLock(ctx); err != nil {
		return func() {}, false
	}
	return func() {
		_ = l.Unlock(context.Background())
	}, true
}

// tryPgsql 使用 pg_try_advisory_lock。锁绑定在这条连接上，用完或连接断开即释放。
func tryPgsql(ctx context.Context, name string) (func(), bool) {
	sqlDB, err := g.DB().GetCore().Master()
	if err != nil {
		g.Log().Warningf(ctx, "pgsql lock master: %+v", err)
		return func() {}, false
	}
	conn, err := sqlDB.Conn(ctx)
	if err != nil {
		g.Log().Warningf(ctx, "pgsql lock conn: %+v", err)
		return func() {}, false
	}
	k1, k2 := advisoryKeys(name)
	var locked bool
	if err = conn.QueryRowContext(ctx, "SELECT pg_try_advisory_lock($1::int4, $2::int4)", k1, k2).Scan(&locked); err != nil || !locked {
		_ = conn.Close()
		if err != nil {
			g.Log().Warningf(ctx, "pgsql lock: %+v", err)
		}
		return func() {}, false
	}
	var once sync.Once
	release := func() {
		once.Do(func() {
			_, _ = conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock($1::int4, $2::int4)", k1, k2)
			_ = conn.Close()
		})
	}
	return release, true
}

// advisoryKeys 把锁名哈希成两个 int32，供 pg_try_advisory_lock 使用。
func advisoryKeys(name string) (int32, int32) {
	h := fnv.New64a()
	_, _ = h.Write([]byte(name))
	sum := h.Sum64()
	return int32(sum), int32(sum >> 32)
}

// tryMysql 使用 GET_LOCK，会话级，超时 0 表示抢不到立刻返回。连接断开即释放。
func tryMysql(ctx context.Context, name string) (func(), bool) {
	sqlDB, err := g.DB().GetCore().Master()
	if err != nil {
		g.Log().Warningf(ctx, "mysql lock master: %+v", err)
		return func() {}, false
	}
	conn, err := sqlDB.Conn(ctx)
	if err != nil {
		g.Log().Warningf(ctx, "mysql lock conn: %+v", err)
		return func() {}, false
	}
	lockName := name
	if len(lockName) > 64 {
		k1, k2 := advisoryKeys(name)
		lockName = fmt.Sprintf("hg:%d:%d", k1, k2)
	}
	var got int
	if err = conn.QueryRowContext(ctx, "SELECT GET_LOCK(?, 0)", lockName).Scan(&got); err != nil || got != 1 {
		_ = conn.Close()
		if err != nil {
			g.Log().Warningf(ctx, "mysql lock: %+v", err)
		}
		return func() {}, false
	}
	var once sync.Once
	return func() {
		once.Do(func() {
			_, _ = conn.ExecContext(context.Background(), "SELECT RELEASE_LOCK(?)", lockName)
			_ = conn.Close()
		})
	}, true
}
