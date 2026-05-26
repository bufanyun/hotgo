# HotGo 项目开发规范

**生成或修改任何代码前，必须完整阅读本文档。**  
本规范是项目唯一事实来源，所有新增代码必须严格对齐现有模块风格，不得自创新的代码模式。

---

## 技术栈

| 层 | 技术                                                             |
|----|----------------------------------------------------------------|
| 后端语言 | Go，module 名 `hotgo`                                            |
| 后端框架 | GoFrame v2（gdb、ghttp、gcron、gerror、gctx、gerror、gjson、gvar） |
| 前端框架 | Vue 3 + Vite + TypeScript                                      |
| UI 组件库 | **Naive UI**（优先使用，未经允许不得引入其他 UI 库）                             |
| 状态管理 | Pinia                                                          |
| HTTP 客户端 | 项目封装 `@/utils/http/axios`（`http.request`），禁止直接使用 axios         |

---

## 后端目录结构

```
server/
├── addons/                       # 插件模块（每个插件独立微架构，见"插件开发"章节）
│   ├── modules/                  # 隐式注册文件（每个插件一个 .go，仅含 import）
│   └── <addonName>/              # 插件目录，与主模块结构镜像
├── api/admin/<module>/           # 主模块 HTTP 契约：Req/Res + g.Meta 路由声明
├── internal/
│   ├── consts/                   # 全局常量与字典选项（枚举在此定义并注册）
│   ├── global/                   # 全局公用变量（谨慎添加，尽量少用）
│   ├── controller/admin/sys/     # 薄控制器，只做参数透传，不含业务逻辑
│   ├── crons/                    # 主模块定时任务注册与处理
│   ├── queues/                   # 主模块消息队列消费者注册与处理
│   ├── logic/sys/                # 业务实现（sSysXxx），init() 注册到 service
│   ├── library/                  # 内部基础库（addons、dict、hgorm、cache、queue、cron 等）
│   ├── service/                  # 接口声明（ISysXxx）+ 注册函数
│   ├── dao/                      # 对外 DAO；internal/dao/internal/ 为 CLI 生成，禁止手改
│   ├── model/
│   │   ├── entity/               # 表结构体（CLI 生成，禁止手改）
│   │   ├── do/                   # ORM 写入结构（CLI 生成，禁止手改）
│   │   └── input/sysin/          # 业务入参、出参、过滤结构
│   └── router/
│       ├── admin.go              # 主路由，末尾调用 genrouter.Register()
│       └── genrouter/            # 每个生成模块一个文件，init() 追加到 LoginRequiredRouter
├── resource/
│   ├── generate/default/curd/    # 主模块 CRUD 代码生成模板（参考，禁止直接修改）
│   └── generate/default/addon/   # 插件脚手架代码生成模板（参考，禁止直接修改）
└── utility/                      # 通用工具：convert、excel、validate 等
```

---

## 后端分层约定

### 分层职责（禁止跨层调用）

```
api（HTTP契约） → controller（参数透传） → service（接口） → logic（业务实现） → dao（数据访问）
```

- **controller 层**：不写任何业务逻辑，只转调 service
- **logic 层**：不出现 HTTP 相关代码，所有校验通过 `Filter()` 或 `g.Validator()` 完成
- **dao 层**：`internal/dao/internal/` 由 CLI 自动生成，禁止手改

---

### 1. API 层 `api/admin/<module>/<module>.go`

- 每个接口一对 `XxxReq` / `XxxRes`，`g.Meta` 声明路由
- GET 接口用 `method:"get"`，写操作用 `method:"post"`
- 请求体嵌入 `sysin.XxxInp`，响应体嵌入对应 Model 或空 `struct{}`

```go
type ListReq struct {
    g.Meta `path:"/xxxModule/list" method:"get" tags:"XXX模块" summary:"获取XXX列表"`
    sysin.XxxListInp
}
type ListRes struct {
    form.PageRes
    List []*sysin.XxxListModel `json:"list" dc:"数据列表"`
}
```

### 2. Controller 层 `internal/controller/admin/sys/<module>.go`

- 变量：`var XxxModule = cXxxModule{}`，类型：`type cXxxModule struct{}`
- 只做参数转发和结果组装，禁止直接操作 dao 或写业务逻辑

```go
func (c *cXxxModule) List(ctx context.Context, req *xxxmodule.ListReq) (res *xxxmodule.ListRes, err error) {
    list, totalCount, err := service.SysXxxModule().List(ctx, &req.XxxListInp)
    if err != nil {
        return
    }
    if list == nil {
        list = []*sysin.XxxListModel{}
    }
    res = new(xxxmodule.ListRes)
    res.List = list
    res.PageRes.Pack(req, totalCount)
    return
}
```

### 3. Service 层 `internal/service/sys.go`

- 在已有文件中**追加**接口声明和注册函数，禁止新建文件
- 声明 `ISysXxxModule` 接口，提供 `SysXxxModule()` 获取函数和 `RegisterSysXxxModule()` 注册函数

### 4. Logic 层 `internal/logic/sys/<module>.go`

- 结构体：`type sSysXxxModule struct{}`
- `init()` 里调用 `service.RegisterSysXxxModule(NewSysXxxModule())`
- `Model()` 方法：`return handler.Model(dao.XxxTable.Ctx(ctx), option...)`
- 列名引用必须用 `dao.XxxTable.Columns().FieldName`，禁止硬编码字符串字段名
- `Edit` 方法以 `in.Id > 0` 区分更新/新增，更新用 `XxxUpdateFields`，新增用 `XxxInsertFields`
- 字段校验在 `sysin` 的 `Filter(ctx)` 方法完成，logic 层不重复校验

### 5. Model 层 `internal/model/input/sysin/<module>.go`

| 类型 | 用途 |
|------|------|
| `XxxUpdateFields` | 更新时字段白名单 |
| `XxxInsertFields` | 新增时字段白名单 |
| `XxxEditInp` | 嵌入 `entity.XxxTable`，含 `Filter()` 校验方法 |
| `XxxListInp` | 列表查询入参，嵌入 `form.PageInp` |
| `XxxListModel` | 列表返回字段 |
| `XxxViewInp` / `XxxViewModel` | 详情入参/出参 |
| `XxxDeleteInp` | 删除入参（通常含 `Ids []int64`） |
| `XxxStatusInp`、`XxxMaxSortInp` 等 | 其他操作 |

### 6. 路由注册 `internal/router/genrouter/<module>.go`

```go
package genrouter

import "hotgo/internal/controller/admin/sys"

func init() {
    LoginRequiredRouter = append(LoginRequiredRouter, sys.XxxModule) // XXX模块
}
```

新建此文件即可，`genrouter/init.go` 通过 `init()` 链自动加载，无需修改 `admin.go`。

---

## 字典与枚举规范

**字典由后端统一维护**，前端通过固定接口拉取，禁止在前端硬编码枚举值。

### 后端：定义与注册（`internal/consts/` 下对应文件）

```go
// 1. 定义常量
const (
    XxxStatusEnabled = 1 // 启用
    XxxStatusDisable = 2 // 禁用
)

// 2. 定义选项（选择合适的样式函数）
var XxxStatusOptions = []*model.Option{
    dict.GenSuccessOption(XxxStatusEnabled, "启用"),
    dict.GenWarningOption(XxxStatusDisable, "禁用"),
}

// 3. 在 init() 中注册（同文件）
func init() {
    dict.RegisterEnums("XxxStatusOptions", "XXX状态选项", XxxStatusOptions)
}
```

**样式函数选择规则：**

| 函数 | 适用场景 |
|------|---------|
| `dict.GenSuccessOption` | 正常、启用、成功 |
| `dict.GenWarningOption` | 待处理、禁用、警告 |
| `dict.GenErrorOption` | 失败、封禁、危险 |
| `dict.GenInfoOption` | 信息类、中性状态 |
| `dict.GenPrimaryOption` | 主要操作、强调 |
| `dict.GenDefaultOption` | 默认、无特殊含义 |
| `dict.GenHashOption` | 不确定数量的动态枚举 |

### 后端：Logic 层中使用字典

```go
// 获取标签文本
label := dict.GetOptionLabel(consts.XxxStatusOptions, in.Status)

// 校验是否为合法值
if !dict.HasOptionKey(consts.XxxStatusOptions, in.Status) {
    return gerror.New("状态值不合法")
}
```

### 前端：加载与使用字典

**加载**（在 `model.ts` 的 `loadOptions()` 中）：

```ts
export function loadOptions() {
    dict.loadOptions(['XxxStatusOptions', 'sys_normal_disable']);
}
```

**在 `model.ts` 的 columns 中渲染字典标签：**

```ts
import { useDictStore } from '@/store/modules/dict';
const dict = useDictStore();

// columns 中
{
    title: '状态',
    key: 'status',
    render(row) {
        return dict.getLabel('XxxStatusOptions', row.status);
    }
}
```

**在搜索表单 schemas 中使用字典下拉：**

```ts
{
    field: 'status',
    component: 'NSelect',
    label: '状态',
    componentProps: {
        options: computed(() => dict.getOptions('XxxStatusOptions')),
        placeholder: '请选择状态',
    },
}
```

---

## 插件（Addon）开发规范

> 定位：独立、临时性、工具类型的功能推荐插件化开发，例如小游戏、广告管理、文章管理、小程序、微商城等。插件之间完全隔离，方便多项目复用。

### 插件目录结构

每个插件拥有完整的独立微架构，与主模块结构镜像，但 **包路径前缀为 `hotgo/addons/<addonName>/`**：

```
server/addons/<addonName>/
├── main.go           # 插件入口：定义 module 结构体，实现 Module 接口，init() 注册
├── global/
│   ├── global.go     # 插件级全局变量
│   └── init.go       # Init(ctx, skeleton)，GetSkeleton() 方法
├── api/
│   ├── admin/        # 后台接口契约（同主模块 api 写法）
│   ├── api/          # 前台 API 接口契约
│   └── home/         # 首页接口契约（如有）
├── controller/
│   ├── admin/sys/    # 后台薄控制器
│   ├── api/          # 前台控制器
│   └── home/         # 首页控制器（如有）
├── logic/
│   ├── logic.go      # 匿名导入各 logic 子包触发 init()
│   └── sys/          # 业务实现（同主模块 logic 写法）
├── service/          # 插件内接口声明与注册（独立，不与主模块 service 混用）
├── model/
│   └── input/sysin/  # 插件内入参/出参（调用主模块服务时继承主模块 input 结构）
├── router/
│   ├── admin.go      # 注册后台路由，使用 addons.RouterPrefix() 获取路由前缀
│   ├── api.go        # 注册前台路由
│   ├── home.go       # 注册首页路由（如有）
│   ├── websocket.go  # 注册 WebSocket 路由（如有）
│   └── genrouter/
│       ├── init.go   # 插件版 genrouter：Register() 使用插件前缀
│       └── *.go      # 每个生成模块一个文件，追加到插件 LoginRequiredRouter
├── consts/           # 插件内常量与字典（同主模块 consts 写法）
├── crons/            # 插件定时任务（可选）
├── queues/           # 插件消息队列消费者（可选）
└── resource/         # 插件静态资源与模板（可选）
```

**对应前端目录：**

```
web/src/
├── api/addons/<addonName>/       # 插件前端 API 封装
└── views/addons/<addonName>/     # 插件前端页面
```

### 插件入口 `main.go`

```go
package <addonName>

import (
    "context"
    "github.com/gogf/gf/v2/net/ghttp"
    _ "hotgo/addons/<addonName>/crons"
    "hotgo/addons/<addonName>/global"
    _ "hotgo/addons/<addonName>/logic"
    _ "hotgo/addons/<addonName>/queues"
    "hotgo/addons/<addonName>/router"
    "hotgo/internal/library/addons"
    "hotgo/internal/service"
    "sync"
)

type module struct {
    skeleton *addons.Skeleton
    ctx      context.Context
    sync.Mutex
}

func init() { newModule() }

func newModule() {
    m := &module{
        skeleton: &addons.Skeleton{
            Label:       "插件显示名",
            Name:        "<addonName>",  // 与目录名一致
            Group:       1,
            Brief:       "简介",
            Description: "详细描述",
            Author:      "作者",
            Version:     "v1.0.0",
        },
        ctx: gctx.New(),
    }
    addons.RegisterModule(m)
}

func (m *module) Start(option *addons.Option) (err error) {
    global.Init(m.ctx, m.skeleton)
    option.Server.Group("/", func(group *ghttp.RouterGroup) {
        group.Middleware(service.Middleware().Addon)
        router.Admin(m.ctx, group)
        router.Api(m.ctx, group)
    })
    return
}
// 其余方法：Stop、Ctx、GetSkeleton、Install、Upgrade、UnInstall（默认留空）
```

### 插件路由前缀规则

```
后台：/admin/<addonName>/接口路径    → 对应 controller/admin/sys/
前台：/api/<addonName>/接口路径      → 对应 controller/api/
首页：/home/<addonName>/接口路径     → 对应 controller/home/
WebSocket：/socket/<addonName>/接口路径
```

路由前缀通过 `addons.RouterPrefix(ctx, consts.AppAdmin, global.GetSkeleton().Name)` 自动生成，**禁止在路由文件里硬编码路径前缀**。

### 插件 genrouter 与主模块的区别

插件 `genrouter/init.go` 的 `Register()` 函数需要传入 `ctx` 和 `group`，并使用插件路由前缀：

```go
// 插件 genrouter/init.go
func Register(ctx context.Context, group *ghttp.RouterGroup) {
    prefix := addons.RouterPrefix(ctx, consts.AppAdmin, global.GetSkeleton().Name)
    group.Group(prefix, func(group *ghttp.RouterGroup) {
        group.Middleware(service.Middleware().AdminAuth)
        if len(LoginRequiredRouter) > 0 {
            group.Bind(LoginRequiredRouter...)
        }
    })
}

// 插件 genrouter/<module>.go
func init() {
    LoginRequiredRouter = append(LoginRequiredRouter, sys.XxxModule)
}
```

参考实现：`server/addons/hgexample/router/genrouter/`

### 插件调用主模块服务

**在插件 input 层继承主模块 input 结构，解耦参数依赖**，避免 import cycle：

```go
// 插件 model/input/sysin/config.go
package sysin

import "hotgo/internal/model/input/sysin"

type UpdateConfigInp struct {
    sysin.UpdateAddonsConfigInp
}
```

```go
// 插件 logic 中调用主模块服务
import isc "hotgo/internal/service"

func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in sysin.UpdateConfigInp) error {
    in.UpdateAddonsConfigInp.AddonName = global.GetSkeleton().Name
    return isc.SysAddonsConfig().UpdateConfigByGroup(ctx, in.UpdateAddonsConfigInp)
}
```

### 插件获取自身信息

```go
// 插件内部
global.GetSkeleton()

// 任意位置判断当前请求是否为插件请求
contexts.IsAddonRequest(ctx)
contexts.GetAddonName(ctx)
```

### 插件隐式注册

在 `server/addons/modules/` 下新建 `<addonName>.go`，内容仅含一行匿名导入：

```go
package modules

import _ "hotgo/addons/<addonName>"
```

### 新增插件检查清单

**后端：**
- [ ] `addons/<addonName>/main.go` — 插件入口，实现 Module 接口
- [ ] `addons/<addonName>/global/` — global.go + init.go
- [ ] `addons/modules/<addonName>.go` — 隐式注册
- [ ] `addons/<addonName>/service/` — 插件内接口声明（独立包）
- [ ] `addons/<addonName>/logic/logic.go` — 匿名导入触发 init()
- [ ] `addons/<addonName>/logic/sys/<module>.go` — 业务实现
- [ ] `addons/<addonName>/model/input/sysin/<module>.go` — 入参/出参
- [ ] `addons/<addonName>/api/admin/<module>/<module>.go` — API 契约
- [ ] `addons/<addonName>/controller/admin/sys/<module>.go` — 薄控制器
- [ ] `addons/<addonName>/router/admin.go` — 路由注册
- [ ] `addons/<addonName>/router/genrouter/init.go` — 插件 genrouter
- [ ] `addons/<addonName>/router/genrouter/<module>.go` — 模块路由追加

**前端：**
- [ ] `web/src/api/addons/<addonName>/index.ts` — API 封装
- [ ] `web/src/views/addons/<addonName>/model.ts` — State、schemas、columns
- [ ] `web/src/views/addons/<addonName>/index.vue` — 列表页
- [ ] `web/src/views/addons/<addonName>/edit.vue` — 编辑页

参考实现：`server/addons/hgexample/`（完整插件示例）、`server/resource/generate/default/addon/`（生成模板）

---

## 标准 CRUD 接口清单

| 接口 | method | 说明 |
|------|--------|------|
| `/xxx/list` | GET | 分页列表 |
| `/xxx/view` | GET | 详情 |
| `/xxx/edit` | POST | 新增/编辑（id>0 为编辑） |
| `/xxx/delete` | POST | 删除（支持批量，传 ids） |
| `/xxx/status` | POST | 启用/禁用状态切换 |
| `/xxx/switch` | POST | 开关字段切换 |
| `/xxx/maxSort` | GET | 获取最大排序值 |
| `/xxx/export` | GET | 导出列表 |

按实际需求裁剪，不强制全部实现。

---

## 前端目录结构

```
web/src/
├── api/<module>/index.ts         # HTTP 请求封装
├── views/<module>/
│   ├── index.vue                 # 列表页
│   ├── edit.vue                  # 新增/编辑弹窗
│   └── model.ts                  # State 类、rules、schemas、columns、loadOptions
├── components/                   # 通用组件（BasicTable、BasicForm、BasicModal 等）
├── store/modules/dict.ts         # 字典 store（loadOptions、getLabel、getOptions）
└── utils/
    ├── http/axios.ts             # HTTP 封装（http.request / jumpExport）
    └── dateUtil.ts               # 日期工具（defRangeShortcuts 等）
```

---

## 前端编码约定

### API 文件 `api/<module>/index.ts`

- 使用 `http.request`，禁止直接使用 axios
- 函数名与接口语义一致：`List`、`Edit`、`Delete`、`View`、`Status`、`Switch`、`Export`、`MaxSort`
- 导出接口用 `jumpExport('/xxx/export', params)`

```ts
import { http, jumpExport } from '@/utils/http/axios';

export function List(params) {
    return http.request({ url: '/xxxModule/list', method: 'get', params });
}
export function Edit(params) {
    return http.request({ url: '/xxxModule/edit', method: 'POST', params });
}
export function Export(params) {
    jumpExport('/xxxModule/export', params);
}
```

### model.ts 结构

- `class State`：属性与后端 entity 字段对应，提供合理默认值
- `newState()`：工厂函数，支持深拷贝（`cloneDeep`）
- `rules`：Naive UI 表单校验规则
- `schemas`：搜索栏 `FormSchema[]`（`ref` 包装）
- `columns`：表格列定义，含权限守卫的操作用 `hasPermission(['/xxx/edit'])`
- `loadOptions()`：调用 `dict.loadOptions([...])`，预加载字典数据

### 组件使用规范

**必须优先使用 Naive UI 原生组件**，能用 Naive UI 实现的禁止自行封装替代组件：

| 场景 | 必须使用的组件 |
|------|--------------|
| 表格 | 项目封装 `BasicTable`（基于 Naive UI） |
| 表单 | 项目封装 `BasicForm`（基于 Naive UI） |
| 弹窗 | `n-modal` 或项目封装 `BasicModal` |
| 按钮 | `n-button`，类型用 `primary`/`default`/`error` 等语义值 |
| 选择器 | `n-select`（字典下拉）、`n-tree-select`（树形） |
| 开关 | `n-switch`，配合 Switch 接口 |
| 输入框 | `n-input` |
| 日期选择 | `n-date-picker`，范围选择配合 `defRangeShortcuts()` |
| 图片渲染 | `renderImage(url)` |
| 文件渲染 | `renderFile(url)` |
| 成员信息 | `renderPopoverMemberSumma(summa)` |

**禁止行为：**
- 禁止引入 Element Plus、Ant Design Vue 等其他 UI 库
- 禁止使用内联 style（使用 Naive UI 的 props 或 class）
- 禁止在模板里硬编码字典值文本，必须通过 `dict.getLabel()` 渲染

### 权限控制

```ts
import { usePermission } from '@/hooks/web/usePermission';
const { hasPermission } = usePermission();

// 模板中
v-if="hasPermission(['/xxxModule/edit'])"

// columns 中开关操作
disabled: !hasPermission(['/xxxModule/switch'])
```

权限标识与后端路由路径保持一致，**路由由后端菜单数据驱动，禁止在源码里硬编码路由路径**。

### Vue 3 代码规范

- 使用 Composition API（`<script setup>`）
- 响应式数据用 `ref` / `reactive`，计算属性用 `computed`
- 字典选项传给组件时用 `computed(() => dict.getOptions('XxxOptions'))`，保持响应式

---

## 新增 CRUD 模块检查清单

**后端（必须全部完成，顺序建议从下往上）：**

- [ ] `internal/consts/<module>.go` — 常量定义、选项变量、`init()` 注册字典
- [ ] 使用 goframe CLI 生成：`entity`、`do`、`dao/internal`、`dao`（禁止手写）
- [ ] `internal/model/input/sysin/<module>.go` — 入参/出参/过滤模型
- [ ] `internal/service/sys.go` — 追加接口声明与注册函数
- [ ] `internal/logic/sys/<module>.go` — 业务实现
- [ ] `internal/controller/admin/sys/<module>.go` — 薄控制器
- [ ] `api/admin/<module>/<module>.go` — API 契约
- [ ] `internal/router/genrouter/<module>.go` — 路由注册（`LoginRequiredRouter` 追加）

**前端（必须全部完成）：**

- [ ] `web/src/api/<module>/index.ts` — HTTP 封装
- [ ] `web/src/views/<module>/model.ts` — State、schemas、columns、rules、loadOptions
- [ ] `web/src/views/<module>/index.vue` — 列表页
- [ ] `web/src/views/<module>/edit.vue` — 编辑页

---

## 参考实现

**所有新模块必须严格对齐以下文件的风格，生成前请逐一阅读：**

**主模块 CRUD 参考：**

| 文件 | 说明 |
|------|------|
| `server/internal/consts/status.go` | 字典常量定义与注册标准写法 |
| `server/api/admin/curddemo/curddemo.go` | API 契约标准写法 |
| `server/internal/controller/admin/sys/curd_demo.go` | 薄控制器标准写法 |
| `server/internal/logic/sys/curd_demo.go` | Logic 标准写法（含关联查询、导出、字段过滤） |
| `server/internal/model/input/sysin/curd_demo.go` | 入参/出参模型标准写法 |
| `server/internal/router/genrouter/curd_demo.go` | 主模块路由注册标准写法 |
| `web/src/api/curdDemo/index.ts` | 前端 API 封装标准写法 |
| `web/src/views/curdDemo/model.ts` | 前端 model 标准写法（含字典加载） |
| `web/src/views/curdDemo/index.vue` | 列表页标准写法 |
| `web/src/views/curdDemo/edit.vue` | 编辑页标准写法 |

**插件参考（以 hgexample 为标准）：**

| 文件 | 说明 |
|------|------|
| `server/addons/hgexample/main.go` | 插件入口标准写法 |
| `server/addons/hgexample/global/init.go` | 插件 global 标准写法 |
| `server/addons/hgexample/router/admin.go` | 插件路由注册标准写法 |
| `server/addons/hgexample/router/genrouter/init.go` | 插件 genrouter 标准写法 |
| `server/addons/hgexample/router/genrouter/tenant_order.go` | 插件模块路由追加标准写法 |
| `server/addons/hgexample/logic/sys/tenant_order.go` | 插件业务逻辑标准写法 |
| `server/addons/hgexample/logic/sys/config.go` | 插件调用主模块服务标准写法 |

代码生成模板：
- 主模块 CRUD：`server/resource/generate/default/curd/`
- 插件脚手架：`server/resource/generate/default/addon/`

以上模板可参考逻辑，**禁止直接修改模板文件**。

---

## 错误处理规范

### 基本原则

- 业务错误直接用 `gerror.New("描述清晰的中文信息")` 返回，GoFrame 框架会统一处理响应码
- 错误信息要面向用户，清晰说明**发生了什么**，而不是堆栈路径

### 错误创建方式

```go
import "github.com/gogf/gf/v2/errors/gerror"

// 普通业务错误（最常用）
return gerror.New("用户名已存在")

// 包装底层错误，追加上下文（调用第三方/dao 层出错时使用）
if err := dao.User.Ctx(ctx).Data(data).Insert(); err != nil {
    return gerror.Wrap(err, "创建用户失败")
}

// 带格式化参数
return gerror.Newf("订单 %d 状态不合法，当前状态：%d", in.Id, order.Status)
```

### 错误传递原则

- **logic 层**：遇到底层错误用 `gerror.Wrap` 追加上下文后返回；纯业务校验失败用 `gerror.New` 直接返回
- **controller 层**：直接 `return` 透传 logic 的 error，**不再包装**，不打印日志（框架统一处理）
- **禁止**在 controller 层 `fmt.Println(err)` 或 `g.Log().Error` 打印已经会被框架记录的错误

---

## 缓存使用规范

### 统一使用 `cache.Instance()`

**禁止直接调用 `g.Redis().Do()`**，项目统一通过 `cache.Instance()` 操作缓存，底层适配器（Memory / Redis / File）由配置决定，业务层无需关心。

```go
import (
    "hotgo/internal/library/cache"
    "github.com/gogf/gf/v2/os/gcache"
    "time"
)

// 设置缓存（带过期时间）
err := cache.Instance().Set(ctx, "user:info:1", userInfo, 30*time.Minute)

// 获取缓存
val, err := cache.Instance().Get(ctx, "user:info:1")
if err != nil || val.IsNil() {
    // 缓存未命中，查数据库
}

// 删除缓存
_, err = cache.Instance().Remove(ctx, "user:info:1")

// GetOrSet：缓存不存在时自动执行函数并写入（推荐用于只读类缓存）
val, err := cache.Instance().GetOrSet(ctx, "user:info:1", func(ctx context.Context) (interface{}, error) {
    return dao.AdminMember.Ctx(ctx).Where(...).One()
}, 30*time.Minute)
```

### 缓存 Key 命名规范

- 格式：`模块:实体:唯一标识`，例如：`user:info:1`、`config:site:basic`、`dict:options:XxxStatus`
- 同一模块的 Key 统一在 `internal/consts/cache.go`（或对应模块的 consts 文件）中定义为常量，禁止散落在 logic 里硬编码
- 插件缓存 Key 加插件名前缀：`hgexample:order:1`

### 缓存使用原则

- 更新或删除数据后，**主动清除**相关缓存，不依赖过期自动失效
- 缓存时间：频繁更新的数据（配置、状态）用较短时间（5～30 分钟）；基本不变的数据（字典、枚举）可用更长时间
- 缓存穿透场景：查询结果为空时也写入短时缓存（如 1 分钟）防止击穿

---

## 日志规范

### 日志级别选择

| 级别 | 方法 | 适用场景 |
|------|------|---------|
| Info | `g.Log().Info` / `g.Log().Infof` | 正常业务流程的关键节点（登录成功、订单创建、支付回调等） |
| Warning | `g.Log().Warning` / `g.Log().Warningf` | **遇到错误但不影响主流程**（重试、降级、外部服务异常、非致命校验失败） |
| Panic | `g.Log().Panic` / `g.Log().Panicf` | 极其严重、必须立即终止的情况（配置缺失、核心依赖不可用） |

### 使用示例

```go
// Info：记录正常流程关键事件
g.Log().Infof(ctx, "用户登录成功 uid:%d ip:%s", member.Id, ip)

// Info：记录外部调用结果
g.Log().Info(ctx, "支付回调处理完成", g.Map{"orderId": orderId, "status": status})

// Warning：外部接口失败但已降级处理
g.Log().Warningf(ctx, "短信发送失败，已跳过 mobile:%s err:%+v", mobile, err)

// Warning：业务异常但不中断流程
g.Log().Warning(ctx, "用户余额不足，跳过自动扣款", g.Map{"uid": uid, "balance": balance})

// Panic：启动时核心配置缺失
g.Log().Panicf(ctx, "缓存未初始化，无法启动: %+v", err)
```

### 日志上下文原则

- 日志**必须传入 `ctx`**（第一个参数），便于链路追踪
- 关键日志需携带足够的业务字段（uid、orderId、requestId 等），用 `g.Map{}` 附加结构化信息
- **禁止**用 `fmt.Println` 替代日志，禁止在 logic/controller 层裸打印调试信息

---

## 数据库建表约定

建表时遵循以下字段命名约定，代码生成器和 ORM hook 会自动识别并处理：

| 字段名 | 类型 | 说明 |
|--------|------|------|
| `id` | bigint(20) NOT NULL AUTO_INCREMENT | 主键，必须 |
| `created_at` | datetime | 创建时间，**自动写入**，禁止手动赋值 |
| `updated_at` | datetime | 更新时间，**自动维护**，禁止手动赋值 |
| `deleted_at` | datetime | 软删除时间，存在时查询自动加 `IS NULL`，删除只更新此字段 |
| `created_by` | bigint(20) | 创建者 ID，**自动写入** |
| `updated_by` | bigint(20) | 更新者 ID，**自动维护** |
| `deleted_by` | bigint(20) | 删除者 ID，**自动写入** |
| `status` | tinyint | 状态字段，默认使用系统字典 `sys_normal_disable` |
| `sort` | int | 排序字段，存在时编辑表单自动获取最大值+1填充 |
| `pid` | bigint(20) | 树表上级ID（树表必须） |
| `level` | int | 树等级（树表必须，自动维护） |
| `tree` | varchar(512) | 关系树（树表必须，自动维护） |
| `tenant_id` | bigint(20) | 租户ID（多租户场景按需加） |

**特殊字段查询方式（代码生成器默认生成）：**
- `string` 类型字段 → `LIKE` 模糊查询
- `int/uint/int64` → `=` 精确查询
- `created_by`/`updated_by`/`deleted_by` → 关键词查询（ID/用户名/手机号匹配），调用 `service.AdminMember().GetIdsByKeyword()`

**数据类型 → Go 类型映射（影响生成代码的字段类型）：**
- `tinyint/int/smallint` → `int`；加 `unsigned` → `uint`
- `bigint` → `int64`；加 `unsigned` → `uint64`
- `varchar/text/char` → `string`
- `datetime/timestamp` → `*gtime.Time`
- `json/jsonb` → `*gjson.Json`
- `decimal/float/double` → `float64`

---

## 数据权限（权限过滤）

业务查询中如需过滤数据权限（仅自己/所属部门/下级等），在 ORM 链式调用中加入 handler：

```go
import "hotgo/internal/library/hgorm/handler"

// 按登录用户的数据权限范围过滤（表需含 created_by 或 member_id 字段）
dao.XxxTable.Ctx(ctx).Handler(handler.FilterAuth).Scan(&res)

// 表中无标准字段时，指定自定义字段
dao.XxxTable.Ctx(ctx).Handler(handler.FilterAuthWithField("custom_field")).Scan(&res)

// 多租户数据权限过滤（表需含 tenant_id 字段）
// 在 Model() 方法中通过 handler.Option 开启
func (s *sSysXxx) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
    if len(option) == 0 {
        option = append(option, &handler.Option{FilterTenant: true})
    }
    return handler.Model(dao.XxxTable.Ctx(ctx), option...)
}
```

多租户增改时自动维护租户关系，在 ORM 操作后加 `.Hook(hook.SaveTenant)`：

```go
import "hotgo/internal/library/hgorm/hook"

// 新增时自动写入 tenant_id/merchant_id/user_id
dao.XxxTable.Ctx(ctx).Fields(sysin.XxxInsertFields{}).Hook(hook.SaveTenant).Data(in).Insert()

// 更新时自动维护
s.Model(ctx).Fields(sysin.XxxUpdateFields{}).WherePri(in.Id).Data(in).Hook(hook.SaveTenant).Update()
```

---

## 定时任务（Crons）

定时任务通过实现接口 + `init()` 注册，在后台「系统设置 → 定时任务」中配置执行策略：

**主模块**：新建 `server/internal/crons/<name>.go`  
**插件**：新建 `server/addons/<addonName>/crons/<name>.go`

```go
package crons

import (
    "context"
    "hotgo/internal/library/cron"
)

func init() {
    cron.Register(MyTask)
}

var MyTask = &cMyTask{name: "myTask"} // name 与后台配置的任务名称一致

type cMyTask struct{ name string }

func (c *cMyTask) GetName() string { return c.name }

func (c *cMyTask) Execute(ctx context.Context, parser *cron.Parser) (err error) {
    parser.Logger.Infof(ctx, "任务执行：%v", ...)
    return
}
```

---

## 消息队列（Queues）

队列采用接口 + `init()` 注册，topic 统一定义在 `internal/consts/queue.go`（或对应模块 consts 文件）。

**主模块**：新建 `server/internal/queues/<name>.go`  
**插件**：新建 `server/addons/<addonName>/queues/<name>.go`

```go
package queues

import (
    "context"
    "encoding/json"
    "hotgo/internal/consts"
    "hotgo/internal/library/queue"
)

func init() {
    queue.RegisterConsumer(MyConsumer)
}

var MyConsumer = &qMyConsumer{}

type qMyConsumer struct{}

func (q *qMyConsumer) GetTopic() string { return consts.QueueMyTopic }

func (q *qMyConsumer) Handle(ctx context.Context, mqMsg queue.MqMsg) (err error) {
    var data MyData
    if err = json.Unmarshal(mqMsg.Body, &data); err != nil {
        return
    }
    // 处理消息...
    return
}
```

**发送消息：**

```go
// 即时消息
queue.Push(consts.QueueMyTopic, data)

// 延迟消息（仅 redis/rocketmq 驱动支持）
queue.SendDelayMsg(consts.QueueMyTopic, data, 10) // redis: 延迟秒数；rocketmq: 延迟级别
```

---

## 工具库使用规范

**后端（`server/utility/`）优先使用已有工具，禁止重复造轮子：**

| 包 | 用途 |
|----|------|
| `utility/convert` | 数据类型转换 |
| `utility/encrypt` | 加密/解密 |
| `utility/excel` | 电子表格导出/导入（导出接口用此包） |
| `utility/validate` | 数据验证工具 |
| `utility/tree` | 树形结构处理 |
| `utility/simple` | 简捷函数集合 |
| `utility/format` | 数据格式化 |

**前端（`web/src/utils/`）优先使用已有工具：**

| 工具 | 用途 |
|------|------|
| `@/utils/http/axios`（`http.request`） | HTTP 请求，**唯一入口** |
| `@/utils/dateUtil`（`defRangeShortcuts`） | 日期范围快捷选项 |
| `@/utils`（`renderImage`、`renderFile`、`renderPopoverMemberSumma`） | 表格渲染工具函数 |
| `@/store/modules/dict`（`useDictStore`） | 字典数据加载与渲染 |
| `@/hooks/web/usePermission`（`hasPermission`） | 权限判断 |

---

## 树表（Tree CRUD）特殊约定

树表在普通 CRUD 基础上有以下差异：

- 表中必须含 `pid`、`level`、`tree` 字段（参见建表约定）
- `pid`/`level`/`tree` 字段**由系统自动维护**，禁止在编辑表单中手动设置这三个字段
- 前端列表页使用树形表格（`BasicTable` 的 `childrenKey` 或 `treeData` 模式），而非普通分页列表
- 参考实现：`server/internal/logic/sys/tree_demo.go`、`web/src/views/develop/curd/treeDemo/`

---

## 响应格式约定

所有接口统一响应格式（由框架中间件自动处理，**无需在 controller 手动包装**）：

```json
{ "code": 0, "message": "操作成功", "timestamp": 1234567890, "traceID": "xxx", "data": {} }
```

- 成功状态码：`0`；失败状态码：`-1`（GoFrame 内置，无需手动设置）
- `gerror.Wrap(err, "用户友好的提示")` → 客户端只看到最外层提示，开发者日志可见完整堆栈
- 禁止在 controller 层手动 `Write`/`WriteJSON` 等直接写响应（破坏统一格式）

---

## 常见错误禁止清单

**后端：**
- 禁止在 controller 层直接调用 dao
- 禁止硬编码字段名字符串（用 `dao.Table.Columns().Field`）
- 禁止手改 `entity/`、`do/`、`dao/internal/` 目录下的文件
- 禁止在 logic 层重复做 `Filter()` 已完成的校验
- 禁止在 `internal/service/sys.go` 之外新建主模块 service 文件（同文件追加）
- 禁止插件 import 另一个插件的包（插件间完全隔离，通过主模块 service 通信）
- 禁止在插件路由文件里硬编码路由前缀（用 `addons.RouterPrefix()` 生成）
- 禁止在 `internal/library/` 下修改基础库（如需扩展，提 issue 或在 utility/ 添加）

**前端：**
- 禁止直接使用 axios，必须用 `http.request`
- 禁止硬编码字典值文本，必须用 `dict.getLabel()` 渲染
- 禁止引入 Naive UI 以外的 UI 框架
- 禁止内联 style
- 禁止在源码里硬编码菜单路由路径
