package member

import (
	"context"

	v1 "hotgo/api/api/member/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetIdByCode(ctx context.Context, req *v1.GetIdByCodeReq) (res *v1.GetIdByCodeRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World api member!")
	return
}
