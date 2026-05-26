## 生成模板开发

### 自定义生成模板

- HotGo允许你新建新的模板分组来满足你的需求，模板可根据现有模板基础拷贝一份出来做改造，默认模板目录：[server/resource/generate/default](../../server/resource/generate/default)

- 系统内置了两组CURD生成模板，请参考[生成模板配置](sys-code.md#生成模板配置)。default：是默认的生成到主模块下；addon：是默认生成到指定的插件下



### 内置gf-cli

- 为了确保生成代码的依赖稳定性，在面对`gf`版本更新可能导致向下不兼容情况时，HotGo将`gf-cli`工具内置到系统中并进行在线执行调整，从而提供更可靠和一致的生成代码功能。

- 后续我们也将开放在线运行`gf gen ...`功能。在做插件开发时也会支持到在线生成插件下的service接口，这将会使得插件开发更加方便



### 指定gf-cli版本

- HotGo多数情况下会和最新版本的gf-cli保持同步，如果更新不及时或你不想使用最新版本的gf-cli来生成代码，可以找到自己想要的版本进行替换即可。

- 下面大致做一些替换步骤说明：

1. 打开`github.com/gogf/gf` 找到你想要使用的版本`clone`下来
2. 将`clone`代码中`gf/cmd/gf/internal/`目录覆盖到`server/internal/library/hggen/internal`
3. 将覆盖过来的目录文件中引入包名`github.com/gogf/gf/cmd/gf/v2/`批量改为`hotgo/internal/library/hggen/`
4. 运行`go mod tidy`
5. 运行`go run main.go`，如果没有报错，那么恭喜你已经完成了。如果有报错一般都是版本差异带来的影响，需要根据情况自行调整



### 指定数据库驱动

> HotGo默认使用mysql驱动，如果你想用其他数据库驱动打开下方文件中注释即可，目前已支持mysql、pgsql

- 修改文件路径：[server/internal/library/hggen/internal/cmd/cmd_gen_dao.go](../../server/internal/library/hggen/internal/cmd/cmd_gen_dao.go)

```go
package cmd

import (
	//_ "github.com/gogf/gf/contrib/drivers/clickhouse/v2"
	//_ "github.com/gogf/gf/contrib/drivers/mssql/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	//_ "github.com/gogf/gf/contrib/drivers/oracle/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	//_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"hotgo/internal/library/hggen/internal/cmd/gendao"
)

type (
	cGenDao = gendao.CGenDao
)

```

修改完成后运行`go mod tidy`
