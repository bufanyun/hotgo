package migrations

import (
	"context"
	"io"
	"os"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/container/gvar"
)
func DoSqlContent(ctx context.Context, sqlPath string) (string, error) {
	open, err := os.Open(sqlPath)
	if err != nil {
		return "fail", err
	}
	defer open.Close()
	sqlContent, err := io.ReadAll(open)
	if err != nil {
		return "fail", err
	}

	// 首先尝试整个执行
	_, err = g.DB().Exec(ctx, string(sqlContent))
	if err != nil {
		g.Log().Error(ctx, "整个执行SQL失败，尝试分句执行:", err)

		// 整个执行失败，尝试按分号分割执行
		sqls := strings.Split(string(sqlContent), ";")
		for _, sql := range sqls {
			sql = strings.TrimSpace(sql)
			if sql != "" {
				_, err := g.DB().Exec(ctx, sql)
				if err != nil {
					g.Log().Error(ctx, "执行SQL失败:", err, sql)
					return "fail", err
				}
			}
		}
	}
	return "success", nil
}

func GetDbLink(ctx context.Context) *gvar.Var {
	link := g.Cfg().MustGet(ctx, "database.default")
	//读写分离
	if !link.IsSlice() {
		return g.Cfg().MustGet(ctx, "database.default.link")
	}
	for _, v := range link.Array() {
		// 只获取主库
		val := v.(map[string]interface{})
		if val["role"] == "master" {
			return gvar.New(val["link"])
		}
	}
	return gvar.New("database.default.0.link")
}

func GetDbType(ctx context.Context) string {
	var (
		link   = GetDbLink(ctx)
	)
	config := strings.SplitN(link.String(), ":", 2)
	return config[0]
}