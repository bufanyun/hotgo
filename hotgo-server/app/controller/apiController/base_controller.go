//
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package apiController

import (
	"context"
	"github.com/bufanyun/hotgo/app/com"
	"github.com/bufanyun/hotgo/app/consts"
	"github.com/bufanyun/hotgo/app/form/apiForm"
	"github.com/bufanyun/hotgo/app/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xuri/excelize/v2"
	"time"
)

// 基础
var Base = base{}

type base struct{}

//
//  @Title  获取lang信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *base) Lang(ctx context.Context, req *apiForm.BaseLangReq) (res *apiForm.BaseLangRes, err error) {

	return
}

//
//  @Title  获取IP归属地信息
//  @Description
//  @Author  Ms <133814250@qq.com>
//  @Param   ctx
//  @Param   req
//  @Return  res
//  @Return  err
//
func (controller *base) IpLocation(ctx context.Context, req *apiForm.IpLocationReq) (res *apiForm.IpLocationRes, err error) {

	panic("测试panic...")
	data := com.Ip.GetLocation(ctx, req.Ip)
	res = &apiForm.IpLocationRes{data}

	return
}

func (controller *base) Excel(ctx context.Context, req *apiForm.ExportReq) (res *apiForm.ExportRes, err error) {
	w := com.Context.Get(ctx).Request.Response

	// 文件名
	fileName := "demo.xlsx"
	// 创建excel文件 （第三方excel包）
	file := excelize.NewFile()
	// 填充数据
	index := file.NewSheet("Sheet1")
	err = file.SetCellValue("Sheet1", "A1", "Hello world.")
	if err != nil {
		g.Log().Print(ctx, "SetCellValue:", err)
		return nil, err
	}
	err = file.SetCellValue("Sheet1", "B1", 100)
	if err != nil {
		g.Log().Print(ctx, "SetCellValue2:", err)
		return nil, err
	}
	file.SetActiveSheet(index)
	// 设置header头
	w.Header().Add("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	// 写入字节数据
	err = file.Write(w.Writer)
	if err != nil {
		g.Log().Print(ctx, "Write:", err)
		return nil, err
	}

	// TODO  加入到上下文
	com.Context.SetResponse(ctx, &model.Response{
		Code:      consts.CodeOK,
		Message:   "",
		Timestamp: time.Now().Unix(),
		ReqId:     com.Context.Get(ctx).ReqId,
	})
	//com.Context.Get(ctx).Request.Exit()
	return
}
