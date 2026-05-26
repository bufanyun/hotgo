## 生成常见问题

### 生成完后页面提示：服务器错误,请稍候重试!

- 热编译环境下，web端往往会快于服务端重启并加载完成，此时接口访问会出现`服务器错误,请稍候重试!`或`404`。这是服务端正在重启导致的，一般稍等几秒就好，如果不行就手动重启下服务端。


### fetching tables failed: SHOW TABLES: Error 1045 (28000): Access denied for user * (using password: YES)

- 请去确认`server/manifest/config/config.yaml`和`server/hack/config.yaml`下的数据库配置一致并且权限正确
- 参考：[生成配置](code-config.md)


### 为什么后台找不到开发工具菜单

- 请去确认`server/manifest/config/config.yaml`中的`system.mode`不为`product`。product模式下后台【开发工具】菜单自动隐藏


