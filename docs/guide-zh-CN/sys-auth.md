## 权限控制

- 控制说明
- 如何页面块/按钮自由控制显示权限
- 不进行权限认证的路由
- 新增应用的 RBAC 授权

> RageFrame已经内置的 RBAC 权限管理，并对其进行了二次开发，无限父子级权限分组、可自由分配子级权限

### 控制说明

可以把权限理解为一个大的库，往里面丢各种名称，可以是路由、自定义名称等

系统自定义了几个默认的名称

- menuCate:1 => 平台管理导航菜单
- menuCate:2 => 微信公众号导航菜单
- menuCate:3 => 系统管理导航菜单
- menuCate:4 => 应用中心导航菜单

> 为什么要命名为cate:1这种名称呢，因为是后面带的是菜单分类数据库自带的id，如果自己有改动比如删除后添加注意修改权限名称  

> 注意！！！  
> 注意！！！  
> 注意！！！  
>
> 权限的路由一定要包含菜单路由才会显示

**举个例子**

已授权权限

```
    [
        '/test/index',
        '/test/system',
        
        ...
    ],
```

你的菜单

```
    [
        'testList', // 这里是菜单的顶级别名
        'child' => [
            '/test/index', // 子菜单
            '/test/system', // 子菜单
            '/test/test', // 子菜单
            ...
         ]
         
         ...
    ],
```

这里看到你授权菜单的子菜单但是没有授权 testList 别名，那么你整个菜单都会被隐藏，如果想显示你必须把 testList 也加入权限

### 如何页面块/按钮自由控制显示权限

按钮快捷方式：使用 [Html](helper-html.md) 辅助类  
自定义验证如下：使用 [Auth](helper-auth.md) 辅助类

### 不进行权限认证的路由

> 修改地址 `backend/config/params` 找到 `noAuthRoute`

```
/**
 * 不需要验证的路由全称
 *
 * 注意: 前面以绝对路径/为开头
 */
'noAuthRoute' => [
    '/main/index',// 系统主页
    '/main/system',// 系统首页
    '/menu-provinces/index',// 微信个性化菜单省市区
    '/wechat/common/select-news',// 微信自动回复获取图文
    '/wechat/common/select-attachment',// 微信自动回复获取图片/视频/
    '/wechat/analysis/image',// 微信显示素材图片
],
```
    
### 新增应用的 RBAC 授权

获取分配角色列表

```
/**
 * @param string $app_id 应用id
 * @param bool $sourceAuthChild 权限来源(false:所有权限，true：当前角色)
 * @return array
 */
Yii::$app->services->rbacAuthRole->getDropDown($app_id, $sourceAuthChild)
```

分配为用户角色

```
/**
 * @param array $role_ids 角色id
 * @param int $user_id 用户id
 * @param string $app_id 应用id
 * @throws UnprocessableEntityHttpException
 */
Yii::$app->services->rbacAuthAssignment->assign($role_ids, $user_id, $app_id);
```

权限管理

> 下面为新增应用权限(比如增加前台会员的权限)案例

```
<?php

namespace addons\Merchants\backend\controllers;

use common\enums\AppEnum;
use common\models\rbac\AuthItem;
use common\traits\AuthItemTrait;

/**
 * Class AuthItemController
 * @package backend\modules\common\controllers
 * @author jianyan74 <751393839@qq.com>
 */
class AuthItemController extends BaseController
{
    use AuthItemTrait;

    /**
     * @var AuthItem
     */
    public $modelClass = AuthItem::class;

    /**
     * 默认应用
     *
     * @var string
     */
    public $appId = AppEnum::MERCHANT;

    /**
     * 渲染视图前缀(默认)
     *
     * @var string
     */
    public $viewPrefix = '@backend/modules/base/views/auth-item/';
}
```

角色管理

> 下面为新增应用权限(比如增加前台会员的角色)案例

```
<?php

namespace addons\Merchants\backend\controllers;

use Yii;
use common\traits\AuthRoleTrait;
use common\models\rbac\AuthRole;
use common\enums\AppEnum;

/**
 * Class RoleController
 * @package addons\Merchants\backend\controllers
 * @author jianyan74 <751393839@qq.com>
 */
class AuthRoleController extends BaseController
{
    use AuthRoleTrait;

    /**
     * @var AuthRole
     */
    public $modelClass = AuthRole::class;

    /**
     * 默认应用
     *
     * @var string
     */
    public $appId = AppEnum::MERCHANT;

    /**
     * 权限来源
     *
     * false:所有权限，true：当前角色
     *
     * @var bool
     */
    public $sourceAuthChild = false;

    /**
     * 渲染视图前缀(默认)
     *
     * @var string
     */
    public $viewPrefix = '@backend/modules/base/views/auth-role/';

    /**
     * @throws \yii\base\InvalidConfigException
     */
    public function init()
    {
        parent::init();

        $this->merchant_id = Yii::$app->request->get('merchant_id');
        $this->merchant_id && Yii::$app->services->merchant->setId($this->merchant_id);
    }
}
```
