package hook

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"hotgo/internal/library/contexts"
	"hotgo/utility/simple"
	"sync"
	"time"
)

type visitor struct {
	lastSeen *gtime.Time
}

var (
	visitors = make(map[int64]*visitor)
	mtx      sync.Mutex
)

// Run a background goroutine to remove old entries from the visitors map.
func init() {
	go cleanupVisitors()
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	gtimer.AddSingleton(gctx.New(), time.Minute, func(ctx context.Context) {
		mtx.Lock()
		for memberId, v := range visitors {
			if gtime.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, memberId)
			}
		}
		mtx.Unlock()
	})
}

func allow(memberId int64) bool {
	mtx.Lock()
	defer mtx.Unlock()
	v, exists := visitors[memberId]
	if !exists {
		visitors[memberId] = &visitor{gtime.Now()}
		return true
	}

	if gtime.Now().Sub(v.lastSeen) > time.Second*3 {
		v.lastSeen = gtime.Now()
		return true
	}

	return false
}

// LastActive 更新用户最后活跃
func (s *sHook) LastActive(r *ghttp.Request) {
	if r.IsFileRequest() {
		return
	}

	var (
		ctx      = r.Context()
		memberId = contexts.GetUserId(ctx)
	)

	if memberId == 0 {
		return
	}

	if allow(memberId) {
		simple.SafeGo(ctx, func(ctx context.Context) {
			_, err := g.Model("admin_member").Ctx(ctx).
				Where("id", memberId).
				WhereLT("last_active_at", gtime.Now()).
				Data(g.Map{"last_active_at": gtime.Now()}).
				Update()
			if err != nil {
				g.Log().Warningf(ctx, "hook LastActive err:%+v, memberId:%v", err, memberId)
			}
		})
	}
}
