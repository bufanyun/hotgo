// Package views
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package views

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hggen/internal/cmd/gendao"
	"hotgo/internal/library/hggen/internal/utility/utils"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/convert"
	"hotgo/utility/file"
	"runtime"
	"strings"
)

var Curd = gCurd{}

type gCurd struct{}

type CurdStep struct {
	HasMaxSort       bool `json:"hasMaxSort"`
	HasAdd           bool `json:"hasAdd"`
	HasBatchDel      bool `json:"hasBatchDel"`
	HasExport        bool `json:"hasExport"`
	HasNotFilterAuth bool `json:"hasNotFilterAuth"`
	HasEdit          bool `json:"hasEdit"`
	HasDel           bool `json:"hasDel"`
	HasView          bool `json:"hasView"`
	HasStatus        bool `json:"hasStatus"`
	HasSwitch        bool `json:"hasSwitch"`
	HasCheck         bool `json:"hasCheck"`
	HasMenu          bool `json:"hasMenu"`
}

type CurdOptionsJoin struct {
	Uuid        string                           `json:"uuid"`
	LinkTable   string                           `json:"linkTable"`
	Alias       string                           `json:"alias"`
	LinkMode    int                              `json:"linkMode"`
	Field       string                           `json:"field"`
	MasterField string                           `json:"masterField"`
	DaoName     string                           `json:"daoName"`
	Columns     []*sysin.GenCodesColumnListModel `json:"columns"`
}

type CurdOptionsMenu struct {
	Icon string `json:"icon"`
	Pid  int    `json:"pid"`
	Sort int    `json:"sort"`
}

type CurdOptions struct {
	AutoOps       []string           `json:"autoOps"`
	ColumnOps     []string           `json:"columnOps"`
	HeadOps       []string           `json:"headOps"`
	Join          []*CurdOptionsJoin `json:"join"`
	Menu          *CurdOptionsMenu   `json:"menu"`
	TemplateGroup string             `json:"templateGroup"`
	ApiPrefix     string             `json:"apiPrefix"`
	Step          *CurdStep          // 转换后的流程控制条件
	dictMap       g.Map              // 字典选项 -> 字段映射关系
}

type CurdPreviewInput struct {
	In           *sysin.GenCodesPreviewInp        // 提交参数
	DaoConfig    gendao.CGenDaoInput              // 生成dao配置
	Config       *model.GenerateConfig            // 生成配置
	view         *gview.View                      // 视图模板
	content      *sysin.GenCodesPreviewModel      // 页面代码
	masterFields []*sysin.GenCodesColumnListModel // 主表字段属性
	pk           *sysin.GenCodesColumnListModel   // 主键属性
	options      *CurdOptions                     // 生成选项
}

type CurdBuildEvent map[string]func(ctx context.Context) (err error)

type CurdBuildInput struct {
	PreviewIn   *CurdPreviewInput // 预览参数
	BeforeEvent CurdBuildEvent    // 前置事件
	AfterEvent  CurdBuildEvent    // 后置事件
}

func (l *gCurd) initInput(ctx context.Context, in *CurdPreviewInput) (err error) {
	in.content = new(sysin.GenCodesPreviewModel)
	in.content.Views = make(map[string]*sysin.GenFile)

	// 加载主表配置
	if err = in.In.MasterColumns.Scan(&in.masterFields); err != nil {
		return
	}

	if len(in.masterFields) == 0 {
		if in.masterFields, err = DoTableColumns(ctx, &sysin.GenCodesColumnListInp{Name: in.In.DbName, Table: in.In.TableName}, in.DaoConfig); err != nil {
			return
		}
	}

	// 主键属性
	in.pk = l.getPkField(in)
	if in.pk == nil {
		return gerror.New("initInput no primary key is set in the table!")
	}

	// 加载选项
	if err = in.In.Options.Scan(&in.options); err != nil {
		return
	}

	initStep(in)
	in.options.dictMap = make(g.Map)

	if len(in.Config.Application.Crud.Templates)-1 < in.In.GenTemplate {
		return gerror.New("没有找到生成模板的配置，请检查！")
	}

	// api前缀
	apiPrefix := gstr.LcFirst(in.In.VarName)
	if in.Config.Application.Crud.Templates[in.In.GenTemplate].IsAddon {
		apiPrefix = in.In.AddonName + "/" + apiPrefix
	}
	in.options.ApiPrefix = apiPrefix

	if err = checkCurdPath(in.Config.Application.Crud.Templates[in.In.GenTemplate], in.In.AddonName); err != nil {
		return
	}
	in.options.TemplateGroup = in.Config.Application.Crud.Templates[in.In.GenTemplate].MasterPackage
	return
}

func initStep(in *CurdPreviewInput) {
	in.options.Step = new(CurdStep)
	in.options.Step.HasMaxSort = HasMaxSort(in.masterFields)
	in.options.Step.HasAdd = gstr.InArray(in.options.HeadOps, "add")
	in.options.Step.HasBatchDel = gstr.InArray(in.options.HeadOps, "batchDel")
	in.options.Step.HasExport = gstr.InArray(in.options.HeadOps, "export")
	in.options.Step.HasNotFilterAuth = gstr.InArray(in.options.ColumnOps, "notFilterAuth")
	in.options.Step.HasEdit = gstr.InArray(in.options.ColumnOps, "edit")
	in.options.Step.HasDel = gstr.InArray(in.options.ColumnOps, "del")
	in.options.Step.HasView = gstr.InArray(in.options.ColumnOps, "view")
	in.options.Step.HasStatus = HasStatus(in.options.ColumnOps, in.masterFields)
	in.options.Step.HasSwitch = HasSwitch(in.options.ColumnOps, in.masterFields)
	in.options.Step.HasCheck = gstr.InArray(in.options.ColumnOps, "check")
	in.options.Step.HasMenu = gstr.InArray(in.options.AutoOps, "genMenuPermissions")
}

func (l *gCurd) loadView(ctx context.Context, in *CurdPreviewInput) (err error) {
	temp := in.Config.Application.Crud.Templates[in.In.GenTemplate]
	view := gview.New()
	err = view.SetConfigWithMap(g.Map{
		"Paths":      temp.TemplatePath,
		"Delimiters": in.Config.Delimiters,
	})
	if err != nil {
		return
	}

	now := gtime.Now()
	view.BindFuncMap(g.Map{
		"NowYear": now.Year,        // 当前年
		"ToLower": strings.ToLower, // 全部小写
		"LcFirst": gstr.LcFirst,    // 首字母小写
		"UcFirst": gstr.UcFirst,    // 首字母大写
	})

	dictOptions, err := l.generateWebModelDictOptions(ctx, in)
	if err != nil {
		return
	}

	modName, err := GetModName(ctx)
	if err != nil {
		return
	}
	importApi := gstr.Replace(temp.ApiPath, "./", modName+"/") + "/" + strings.ToLower(in.In.VarName)
	importInput := gstr.Replace(temp.InputPath, "./", modName+"/")
	importController := gstr.Replace(temp.ControllerPath, "./", modName+"/")
	importService := "hotgo/internal/service"
	if temp.IsAddon {
		importService = "hotgo/addons/" + in.In.AddonName + "/service"
	}

	importWebApi := "@/api/" + gstr.LcFirst(in.In.VarName)
	if temp.IsAddon {
		importWebApi = "@/api/addons/" + in.In.AddonName + "/" + gstr.LcFirst(in.In.VarName)
	}

	componentPrefix := gstr.LcFirst(in.In.VarName)
	if temp.IsAddon {
		componentPrefix = "addons/" + in.In.AddonName + "/" + componentPrefix
	}

	nowTime := now.Format("Y-m-d H:i:s")
	view.Assigns(gview.Params{
		"templateGroup":    in.options.TemplateGroup,                                    // 生成模板分组名称
		"servFunName":      l.parseServFunName(in.options.TemplateGroup, in.In.VarName), // 业务服务名称
		"nowTime":          nowTime,                                                     // 当前时间
		"version":          runtime.Version(),                                           // GO 版本
		"hgVersion":        consts.VersionApp,                                           // HG 版本
		"varName":          in.In.VarName,                                               // 实体名称
		"tableComment":     in.In.TableComment,                                          // 对外名称
		"daoName":          in.In.DaoName,                                               // ORM模型
		"masterFields":     in.masterFields,                                             // 主表字段
		"pk":               in.pk,                                                       // 主键属性
		"options":          in.options,                                                  // 提交选项
		"dictOptions":      dictOptions,                                                 // web字典选项
		"importApi":        importApi,                                                   // 导入goApi包
		"importInput":      importInput,                                                 // 导入input包
		"importController": importController,                                            // 导入控制器包
		"importService":    importService,                                               // 导入业务服务
		"importWebApi":     importWebApi,                                                // 导入webApi
		"apiPrefix":        in.options.ApiPrefix,                                        // api前缀
		"componentPrefix":  componentPrefix,                                             // vue子组件前缀
	})

	in.view = view
	return
}

func (l *gCurd) DoBuild(ctx context.Context, in *CurdBuildInput) (err error) {
	preview, err := l.DoPreview(ctx, in.PreviewIn)
	if err != nil {
		return
	}

	// 前置操作
	if len(in.BeforeEvent) > 0 {
		for name, f := range in.BeforeEvent {
			if gstr.InArray(in.PreviewIn.options.AutoOps, name) {
				if err = f(ctx); err != nil {
					return gerror.Newf("in doBuild operation beforeEvent to '%s' failed:%v", name, err)
				}
			}
		}
	}

	// 处理sql文件
	handleSqlFile := func(vi *sysin.GenFile) (err error) {
		// 无需生成
		if vi.Meth != consts.GenCodesBuildMethCreate && vi.Meth != consts.GenCodesBuildMethCover {
			return
		}

		if err = gfile.PutContents(vi.Path, strings.TrimSpace(vi.Content)); err != nil {
			return gerror.Newf("writing content to '%s' failed: %v", vi.Path, err)
		}

		// 导入失败，将sql文件删除
		if err = ImportSql(ctx, vi.Path); err != nil {
			_ = gfile.Remove(vi.Path)
		}
		return
	}

	// 将sql文件提取出来优先处理
	// sql执行过程出错是高概率事件，后期在执行前要进行预效验，尽量减少在执行过程中出错的可能性
	sqlGenFile, ok := preview.Views["source.sql"]
	if ok {
		delete(preview.Views, "source.sql")
		if err = handleSqlFile(sqlGenFile); err != nil {
			return
		}
	}

	for _, vi := range preview.Views {
		// 无需生成
		if vi.Meth != consts.GenCodesBuildMethCreate && vi.Meth != consts.GenCodesBuildMethCover {
			continue
		}

		if err = gfile.PutContents(vi.Path, strings.TrimSpace(vi.Content)); err != nil {
			return gerror.Newf("writing content to '%s' failed: %v", vi.Path, err)
		}

		if gstr.Str(vi.Path, `.`) == ".go" {
			utils.GoFmt(vi.Path)
		}
	}

	// 后置操作
	if len(in.AfterEvent) > 0 {
		for name, f := range in.AfterEvent {
			if gstr.InArray(in.PreviewIn.options.AutoOps, name) {
				if err = f(ctx); err != nil {
					return gerror.Newf("in doBuild operation afterEvent to '%s' failed:%v", name, err)
				}
			}
		}
	}
	return
}

func (l *gCurd) DoPreview(ctx context.Context, in *CurdPreviewInput) (res *sysin.GenCodesPreviewModel, err error) {
	// 初始化
	if err = l.initInput(ctx, in); err != nil {
		return nil, err
	}

	// 加载模板
	if err = l.loadView(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateApiContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateInputContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateControllerContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateLogicContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateRouterContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebApiContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebModelContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebIndexContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebEditContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateWebViewContent(ctx, in); err != nil {
		return nil, err
	}

	if err = l.generateSqlContent(ctx, in); err != nil {
		return nil, err
	}

	in.content.Config = in.Config
	res = in.content
	return
}

func (l *gCurd) generateApiContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "api.go"
		tplData = g.Map{}
		genFile = new(sysin.GenFile)
	)
	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].ApiPath, strings.ToLower(in.In.VarName), strings.ToLower(in.In.VarName)+".go")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}

	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateInputContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "input.go"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.inputTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}
	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].InputPath, convert.CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateControllerContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "controller.go"
		tplData = g.Map{}
		genFile = new(sysin.GenFile)
	)

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}
	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].ControllerPath, convert.CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateLogicContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "logic.go"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.logicTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}
	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].LogicPath, convert.CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateRouterContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "router.go"
		tplData = g.Map{}
		genFile = new(sysin.GenFile)
	)
	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].RouterPath, convert.CamelCaseToUnderline(in.In.VarName)+".go")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebApiContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "web.api.ts"
		tplData = g.Map{}
		genFile = new(sysin.GenFile)
	)
	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebApiPath, gstr.LcFirst(in.In.VarName), "index.ts")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebModelContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "web.model.ts"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.webModelTplData(ctx, in)
	if err != nil {
		return
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "model.ts")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}
	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebIndexContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "web.index.vue"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.webIndexTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "index.vue")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}
	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebEditContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "web.edit.vue"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.webEditTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "edit.vue")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true
	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	if !in.options.Step.HasEdit {
		genFile.Meth = consts.GenCodesBuildIgnore
		genFile.Required = false
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateWebViewContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "web.view.vue"
		genFile = new(sysin.GenFile)
	)

	tplData, err := l.webViewTplData(ctx, in)
	if err != nil {
		return err
	}

	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].WebViewsPath, gstr.LcFirst(in.In.VarName), "view.vue")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if genFile.Meth == consts.GenCodesBuildMethSkip && gstr.InArray(in.options.AutoOps, "forcedCover") {
		genFile.Meth = consts.GenCodesBuildMethCover
	}

	if !in.options.Step.HasView {
		genFile.Meth = consts.GenCodesBuildIgnore
		genFile.Required = false
	}

	in.content.Views[name] = genFile
	return
}

func (l *gCurd) generateSqlContent(ctx context.Context, in *CurdPreviewInput) (err error) {
	var (
		name    = "source.sql"
		config  = g.DB("default").GetConfig()
		tplData = g.Map{
			"dbName":        config.Name,
			"menuTable":     config.Prefix + "admin_menu",
			"mainComponent": "LAYOUT",
		}
		genFile = new(sysin.GenFile)
	)

	tplData["dirPid"], tplData["dirLevel"], tplData["dirTree"], err = hgorm.AutoUpdateTree(ctx, &dao.AdminMenu, 0, int64(in.options.Menu.Pid))
	if err != nil {
		return err
	}

	tplData["listLevel"] = tplData["dirLevel"].(int) + 1
	tplData["btnLevel"] = tplData["dirLevel"].(int) + 2
	tplData["sortLevel"] = tplData["dirLevel"].(int) + 3

	if in.options.Menu.Pid > 0 {
		tplData["mainComponent"] = "ParentLayout"
	}

	genFile.Path = file.MergeAbs(in.Config.Application.Crud.Templates[in.In.GenTemplate].SqlPath, convert.CamelCaseToUnderline(in.In.VarName)+"_menu.sql")
	genFile.Meth = consts.GenCodesBuildMethCreate
	if gfile.Exists(genFile.Path) {
		genFile.Meth = consts.GenCodesBuildMethSkip
	}
	genFile.Required = true

	if !in.options.Step.HasMenu {
		genFile.Meth = consts.GenCodesBuildIgnore
		genFile.Required = false
	}

	tplData["generatePath"] = genFile.Path
	genFile.Content, err = in.view.Parse(ctx, name+".template", tplData)
	if err != nil {
		return err
	}

	in.content.Views[name] = genFile
	return
}
