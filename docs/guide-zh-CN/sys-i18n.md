## 国际化


目录

- 介绍
- 配置文件
- 服务端使用
- Web端使用
- 更多文档

### 介绍

从`v2.18.6`版本开始，HotGo 提供了完善的国际化（i18n）支持，内置简体中文、繁体中文和英文三种语言包，默认语言为简体中文。目前已对后台管理首页进行了多语言适配，开发者可基于该机制扩展实现业务模块的多语言功能。

### 配置文件

国际化配置位于 `server/manifest/config/config.yaml` 文件中：

```yaml
# hotgo系统配置
system:
  # ...
  # 国际化
  i18n:
    switch: true                                      # 国际化功能开关，可选值：false|true，默认：true
    defaultLanguage: "zh-CN"                          # 默认语言，可选值：zh-CN|zh-TW|en，默认：zh-CN
```

### 服务端使用

多语言配置文件存放目录：`server/manifest/i18n`，您可以在该目录下添加自定义的语言映射文件。

**使用示例：**

```go
// 设置当前上下文的语言为英文
ctx := gctx.New()
gi18n.WithLanguage(ctx, "en")

// 基础翻译
gi18n.T(ctx, "你好，美丽世界")
// 输出：Hello, Beautiful World

// 格式化翻译（支持参数替换）
gi18n.Tf(ctx, "剩余%v余额", 100)
// 输出：Remaining 100 Balance
```


### Web端使用

多语言配置文件存放目录：`web/src/locale`，您可以在该目录下添加自定义的语言映射文件。

> 提示：启用国际化功能后，后台管理界面右上角将显示语言切换选项。

**使用示例（假设当前语言为英文）：**

```vue
<!-- 基础翻译 -->
<div>{{ t('你好，美丽世界') }}</div>
<!-- 输出：Hello, Beautiful World -->

<!-- 格式化翻译（支持参数替换） -->
<div>{{ t('剩余{num}余额', {num:100}) }}</div>
<!-- 输出：Remaining 100 Balance -->
```


### 更多文档

- [gi18n](https://goframe.org/docs/core/gi18n)
- [vie-i18n](https://vue-i18n.intlify.dev)