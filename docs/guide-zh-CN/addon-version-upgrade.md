## RageFrame 2.x 插件升级到 3.x

目录

- 路由替换
- 配置移除
- UI 替换
- Icons 替换
- 表单组件 替换
- 安装

> 注意：默认当已安装 PhpStorm 等可以批量替换的编译器/软件

### 路由替换

1. 查找 `插件/common/config` 目录

2. 替换 `route` 为 `name`

### 配置移除

1. 查找 `插件/AddonConfig.php` 文件

2. 找到 `appsConfig` 变量的 `merapi` 删除掉，移除掉以后的配置

```
public $appsConfig = [
   'backend' => 'common/config/backend.php',
   'frontend' => 'common/config/frontend.php',
   'merchant' => 'common/config/merchant.php',
   'html5' => 'common/config/html5.php',
   'api' => 'common/config/api.php',
   'oauth2' => 'common/config/oauth2.php',
];
```

### UI 替换

1. 查找 `插件` 目录

2. 替换 `<div class='col-sm-1 text-right'>{label}</div><div class='col-sm-11'>{input}{hint}{error}</div>` 为 `<div class='row'><div class='col-sm-1 text-right'>{label}</div><div class='col-sm-11'>{input}\n{hint}\n{error}</div></div>`

3. 替换 `<div class='col-sm-2 text-right'>{label}</div><div class='col-sm-10'>{input}{hint}{error}</div>` 为 `<div class='row'><div class='col-sm-2 text-right'>{label}</div><div class='col-sm-10'>{input}\n{hint}\n{error}</div></div>`

4. 替换 `<div class='col-sm-2 text-right'>{label}</div><div class='col-sm-10'>{input}\n{hint}\n{error}</div>` 为 `<div class='row'><div class='col-sm-2 text-right'>{label}</div><div class='col-sm-10'>{input}\n{hint}\n{error}</div></div>`

5. 替换 `<div class='col-sm-3 text-right'>{label}</div><div class='col-sm-9'>{input}{hint}{error}</div>` 为 `<div class='row'><div class='col-sm-3 text-right'>{label}</div><div class='col-sm-9'>{input}\n{hint}\n{error}</div></div>`

6. 替换 `modal` 

> 批量替换估计难查找到，可以看见了手动替换掉，关键词搜索 `基本信息`

```
<div class="modal-header">
    <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">×</span><span class="sr-only">关闭</span></button>
    <h4 class="modal-title">基本信息</h4>
</div>
```

为

```
<div class="modal-header">
    <h4 class="modal-title">基本信息</h4>
    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>
</div>
```

7. 替换布局

如果单独使用了 `col-*` 这种布局出现页面失调的，需要在 `col-*` 外层加上 `<div class='row'></div>`。

例如:

```
<div class="row">
    <div class="col-6">
        
    </div>
    <div class="col-6">
        
    </div>
</div>
```

### Icons 替换

Icons 找不到请使用最新的 Icons 库 https://fontawesome.com/v5/search?s=solid

### 表单组件替换

个别表单组件报错请参考最新的使用文档 [表单控件](sys-widget.md)

### 安装

1. 放入 3.x 的 `addons` 目录，在后台->应用管理查找安装即可

2. 个别页面出现大小表格缩短显示问题可以修改替换

 - `col-xs-12` 为 `col-12`
 - `col-lg-12` 为 `col-12`
 - `col-sm-12` 为 `col-12`

> 个别报错找不到的组件需要自己替换修复