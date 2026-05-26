## SaaS多租户


目录

- 介绍
- 职责划分
- 如何在HotGo开发多租户业务
- 多租户数据库设计
- 多租户功能演示

### 介绍

SaaS系统多租户多应用设计，已成为互联网企业的重要发展建设方向。核心在于多租户SAAS系统独立前台、共享后台、共享数据库的平台应用架构。

> 目前HotGo的部分基础功能并不完全支持多租户设计，但这并不妨碍开发者基于HotGo构建自己的多租户产品


### 职责划分

在SaaS系统多租户中，不同身份用户的职责功能划分假设可以是这样：

> 这只是一个粗略的概念，实际开发时应根据业务来进行调整，每个身份也都不是必须的。

| 身份        | 标识       | 职责和功能划分                                               |
|-----------------|----------|-------------------------------------------------------|
| 公司| company  | 管理整个平台，包括商户和用户账户、系统设置以及其他全局性业务流程。                     |
| 租户          | tenant   | 多租户系统中顶层实体客户、组织或实体。有自己的多个商户、用户、产品、订单等。拥有独立的数据隔离和安全边界。 |
| 商户          | merchant | 受租户的监管和管理，可独立经营的实体。提供产品或服务，管理自己的业务，包括库存管理、订单处理、结算等。          |
| 用户          | user     | 真正购买产品或享受服务的人，与商户互动，管理个人信息等个性化功能。                     |


### 如何在HotGo开发多租户业务
#### 一、应用功能

根据角色来划分用户的后台功能，在创建用户时为其绑定角色，然后为不同角色分配不同的功能菜单

请参考： [权限控制](sys-auth.md)


#### 二、 数据隔离

根据部门来划定用户的数据权限范围，在创建用户时为其绑定部门

- 在用户登录成功后，server端可通过上下文来获取用户部门类型来确定用户身份
- 文件路径：server/internal/library/contexts/context.go

```go
package contexts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/consts"
	"hotgo/internal/model"
)

// GetDeptType 获取用户部门类型
func GetDeptType(ctx context.Context) string {
	user := GetUser(ctx)
	if user == nil {
		return ""
	}
	return user.DeptType
}

// IsCompanyDept 是否为公司部门
func IsCompanyDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeCompany
}

// IsTenantDept 是否为租户部门
func IsTenantDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeTenant
}

// IsMerchantDept 是否为商户部门
func IsMerchantDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeMerchant
}

// IsUserDept 是否为普通用户部门
func IsUserDept(ctx context.Context) bool {
	return GetDeptType(ctx) == consts.DeptTypeUser
}
```

- 在用户登录成功后，web端可通`useUserStore`来获取用户部门类型来确定用户身份
- 文件路径：web/src/store/modules/user.ts

```vue
<script lang="ts" setup>
import { useUserStore } from '@/store/modules/user';

const userStore = useUserStore();

console.log('用户部门类型:' + userStore.info?.deptType);
console.log('是否为公司:' + userStore.isCompanyDept);
console.log('是否为租户:' + userStore.isTenantDept);
console.log('是否为商户:' + userStore.isMerchantDept);
console.log('是否为用户:' + userStore.isUserDept);
</script>
```

### 多租户数据库设计

HotGo定位是中小型应用开发，推荐采用一套数据库不同Schema。就是在多租户业务表中加入用户标识字段，来区分不同用户的数据，如：`tenant_id`

- 参考文章：https://blog.csdn.net/haponchang/article/details/104246317


### 多租户功能演示

请登录后台【插件应用】-【功能案例】-【多租户功能演示】查看

#### 自动维护租户关系

- 只需在表设计时包含以下字段，即可使用handler和hook实现租户权限过滤和租户关系维护

| 字段名称 | 数据类型    | 字段注释 | 必选                                                  |
|------|----------|------|-----------------------------------------------------|
| tenant_id   | bigint(20)  | 租户ID | 否                                                   |
| merchant_id   | bigint(20)   | 商户ID | 否 |
| user_id   | bigint(20) | 用户ID | 否   |

下面是多租户功能演示例子代码中的使用片段

- 封装查询Model

```go
// Model 多租户功能演示ORM模型
func (s *sSysTenantOrder) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	if len(option) == 0 {
		// 过滤多租户数据权限
		option = append(option, &handler.Option{
			FilterTenant: true,
			//FilterAuth:   true, // 如果还需要维护created_by、member_id等部门数据权限范围可放开注释
		})
	}
	return handler.Model(dao.AddonHgexampleTenantOrder.Ctx(ctx), option...)
}
```

- 增改数据自动维护租户关系
```go
// Edit 修改/新增多租户功能演示
func (s *sSysTenantOrder) Edit(ctx context.Context, in *sysin.TenantOrderEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.TenantOrderUpdateFields{}).
				WherePri(in.Id).Data(in).
				Hook(hook.SaveTenant). // 自动维护租户关系更新
				Update(); err != nil {
			}
			return
		}

		// 新增
		if _, err = dao.AddonHgexampleTenantOrder.Ctx(ctx).
			Fields(sysin.TenantOrderInsertFields{}).
			Hook(hook.SaveTenant). // 自动维护租户关系更新
			Data(in).
			Insert(); err != nil {
		}
		return
	})
}
```

相关代码文件：/server/addons/hgexample/logic/sys/tenant_order.go

