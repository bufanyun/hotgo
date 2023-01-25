// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package views

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

const (
	IndexApiImport   = "  import {%v } from '@/api/%s';"
	IndexIconsImport = "  import {%v } from '@vicons/antd';"
)

func (l *gCurd) webIndexTplData(ctx context.Context, in *CurdPreviewInput) (g.Map, error) {
	var (
		data        = make(g.Map)
		apiImport   = []string{" List"}
		iconsImport []string
	)

	// 添加
	if in.options.Step.HasAdd {
		iconsImport = append(iconsImport, " PlusOutlined")
	}

	// 编辑
	if in.options.Step.HasEdit {
	}

	// 导出
	if in.options.Step.HasExport {
		iconsImport = append(iconsImport, " ExportOutlined")
		apiImport = append(apiImport, " Export")
	}

	// 删除
	if in.options.Step.HasDel || in.options.Step.HasBatchDel {
		iconsImport = append(iconsImport, " DeleteOutlined")
		apiImport = append(apiImport, " Delete")
	}

	// 导出
	if in.options.Step.HasStatus {
		apiImport = append(apiImport, " Status")
	}

	data["apiImport"] = fmt.Sprintf(IndexApiImport, gstr.Implode(",", apiImport), gstr.LcFirst(in.In.VarName))
	if len(iconsImport) > 0 {
		data["iconsImport"] = fmt.Sprintf(IndexIconsImport, gstr.Implode(",", iconsImport))
	}

	return data, nil
}
