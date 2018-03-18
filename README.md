# IMG HOSTING

IMG HOSTING 是个简单而实用的图库系统，Golang + Vue，前后分离。

[预览地址: http://119.23.232.219:8001/](http://119.23.232.219:8001/)

![img](https://github.com/pwcong/SnapShot/blob/master/img-hosting/01.png?raw=true)

## 安装

0. 搭建 Golang 开发环境，搭建 NodeJS 开发环境，准备好 MySQL 和 Redis 数据库
1. `项目根目录` 中执行 `go get ./` 安装所有 Golang 依赖包，然后执行 `go build` 构建二进制执行文件
2. `项目根目录/view` 中执行 `npm install` 安装所有 NodeJS 依赖包，然后执行 `npm run build` 生成前端文件
3. 修改 `项目根目录/config/default.toml` 配置文件
4. 执行第一步生成的二进制文件启动系统

## 技术栈

**后端(Golang)：**
* Echo
* Gorm
* jwt-go
* go-redis

**前端(Vue)：**
* Webpack
* Vue + Vuex + Vue-Router
* element-ui
* moment
* chance