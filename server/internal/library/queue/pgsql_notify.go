package queue

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lib/pq"
)

// pgsqlQueueChannel PostgreSQL NOTIFY 频道。消息体是主题名。
const pgsqlQueueChannel = "hotgo_queue"

var pgsqlListenOnce sync.Once

// notifyPgsql 写入成功后通知其他进程。本进程另有 publishQueue，不依赖这次通知。
func notifyPgsql(ctx context.Context, db gdb.DB, topic string) {
	if _, err := db.Exec(ctx, "SELECT pg_notify(?, ?)", pgsqlQueueChannel, topic); err != nil {
		g.Log().Warningf(ctx, "pgsql queue notify: %+v", err)
	}
}

// startPgsqlListen 进程内只开一条 LISTEN 连接。收到通知后唤醒对应主题，不再靠轮询发现新消息。
func startPgsqlListen(db gdb.DB) {
	pgsqlListenOnce.Do(func() {
		ctx := context.Background()
		dsn, err := postgresDSN(db)
		if err != nil {
			g.Log().Warningf(ctx, "pgsql queue listen skipped: %+v", err)
			return
		}
		listener := pq.NewListener(dsn, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
			if err != nil {
				g.Log().Warningf(context.Background(), "pgsql queue listen: %+v", err)
			}
		})
		if err = listener.Listen(pgsqlQueueChannel); err != nil {
			g.Log().Warningf(ctx, "pgsql queue listen: %+v", err)
			_ = listener.Close()
			return
		}
		go func() {
			ping := time.NewTicker(time.Minute)
			defer ping.Stop()
			for {
				select {
				case n, ok := <-listener.Notify:
					if !ok {
						return
					}
					if n != nil && n.Extra != "" {
						publishQueue(n.Extra)
					}
				case <-ping.C:
					if err := listener.Ping(); err != nil {
						g.Log().Warningf(context.Background(), "pgsql queue listen ping: %+v", err)
					}
				}
			}
		}()
	})
}

// postgresDSN 用当前库配置拼出 lib/pq 连接串，供 LISTEN 单独占用一条连接。
func postgresDSN(db gdb.DB) (string, error) {
	cfg := db.GetConfig()
	host, port, user, pass, name := cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name
	if host == "" && cfg.Link != "" {
		var err error
		user, pass, host, port, name, err = parsePgsqlLink(cfg.Link)
		if err != nil {
			return "", err
		}
	}
	if host == "" || name == "" {
		return "", fmt.Errorf("pgsql queue listen: database address is empty")
	}
	if port == "" {
		port = "5432"
	}
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, pass),
		Host:     net.JoinHostPort(host, port),
		Path:     name,
		RawQuery: "sslmode=disable",
	}
	return u.String(), nil
}

// parsePgsqlLink 解析 pgsql:user:pass@tcp(host:port)/db。
func parsePgsqlLink(link string) (user, pass, host, port, name string, err error) {
	rest := strings.TrimPrefix(link, "pgsql:")
	at := strings.LastIndex(rest, "@tcp(")
	if at < 0 {
		return "", "", "", "", "", fmt.Errorf("pgsql queue listen: bad link")
	}
	cred := rest[:at]
	tail := rest[at+len("@tcp("):]
	user, pass, _ = strings.Cut(cred, ":")
	hostport, db, ok := strings.Cut(tail, ")/")
	if !ok {
		return "", "", "", "", "", fmt.Errorf("pgsql queue listen: bad link")
	}
	host, port, _ = strings.Cut(hostport, ":")
	name = db
	if i := strings.IndexByte(name, '?'); i >= 0 {
		name = name[:i]
	}
	return user, pass, host, port, name, nil
}
