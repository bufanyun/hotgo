package global

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

func TestLoadHTTPRoutesConcurrentInit(t *testing.T) {
	const (
		rounds     = 20
		goroutines = 50
		routeCount = 20
	)

	paths := make([]string, routeCount)
	for i := range paths {
		paths[i] = fmt.Sprintf("/route-cache/res%02d", i)
	}

	// 唯一 server 名称避免与其他测试注册的路由互相干扰
	srv := ghttp.GetServer(fmt.Sprintf("route-cache-race-test-%d", time.Now().UnixNano()))
	for _, p := range paths {
		srv.BindHandler("GET:"+p, func(r *ghttp.Request) {})
	}

	req := &ghttp.Request{
		Request: httptest.NewRequest(http.MethodGet, paths[0], nil),
		Server:  srv,
	}

	reset := func() {
		httpRoutes = nil
		httpRoutesOnce = sync.Once{}
	}
	reset()
	t.Cleanup(reset)

	keys := make([]string, routeCount)
	for i, p := range paths {
		keys[i] = GenRouteKey(http.MethodGet, p)
	}

	for round := range rounds {
		reset()
		var (
			wg    sync.WaitGroup
			start = make(chan struct{})
		)
		wg.Add(goroutines)
		for range goroutines {
			go func() {
				defer wg.Done()
				<-start
				got := LoadHTTPRoutes(req)
				// 拿到 map 后立刻遍历,旧实现会在填充过程中被读到空值
				for _, key := range keys {
					if got[key] == nil {
						t.Errorf("第 %d 轮缺少路由 %q", round, key)
					}
				}
			}()
		}
		close(start)
		wg.Wait()
	}

	if route := GetRequestRoute(req); route == nil {
		t.Fatal("GetRequestRoute 未返回已缓存的路由")
	}
}
