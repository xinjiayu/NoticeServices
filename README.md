# NoticeServices 通用的通知服务


将通知类的信息推送到Android、IOS、web、短信、邮件、企业微信。方便引入到自己的应用程序中，并可以单独部署。

支持功能：

- 即时推送
- 预约推送
- 定期推送

## 逻辑架构图

![design01](document/design01.jpg)


## 技术栈

基础框架：[GoFrame](https://github.com/gogf/gf) 【 [中文文档](https://goframe.org/index) 】

数据库：SQLite   【 [中文文档](https://doc.yonyoucloud.com/doc/wiki/project/sqlite/sqlite-intro.html) 】      *GO驱动使用  github.com/mattn/go-sqlite3* 【 [接口文档](https://godoc.org/github.com/mattn/go-sqlite3) 】

目录结构说明：

参考基础框架中的说明 【[项目结构](https://itician.org/pages/viewpage.action?pageId=3670259#id-%E6%96%B0%E5%BB%BA%E9%A1%B9%E7%9B%AE-%E9%A1%B9%E7%9B%AE%E7%BB%93%E6%9E%84) 】



## 关于build.sh编译脚本

在使用build.sh脚本进行程序编译的时候，提示
```
fatal: No names found, cannot describe anything.
./build.sh linux|windows|mac

```
是因为源码没有进行git版本的标签设置。

支持将git的tag编译到程序中。需要创建git的tag。只有创建了tag，编译的程序才会显示版本号。

```
git tag v0.0.1

git push origin v0.0.1
```



## 感谢 JetBrains

<a href="https://www.jetbrains.com/?from=Mybatis-PageHelper" target="_blank">
<img src="https://user-images.githubusercontent.com/1787798/69898077-4f4e3d00-138f-11ea-81f9-96fb7c49da89.png" alt="JetBrains" height="200"/></a>