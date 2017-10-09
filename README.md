# Golang Web Scaffold
开箱即用的golang echo框架脚手架

**献给所有冰岩人，愿能为你的伟大创造添砖加瓦**



## 准备工作
* 开发、部署环境的配置都使用Docker：[Install Docker](https://docs.docker.com/engine/installation/)
* 使用docker-compose，现在桌面操作系统上docker-compose已经和docker打包安装，无须独立安装，不过请确保已经安装，尤其是使用linux操作系统的：[Install Docker Compose](https://docs.docker.com/compose/install/)
* 使用git做版本控制：[Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* golang环境已经集成在docker容器中，包含以下库
  * `github.com/labstack/echo` 
  * `github.com/codegangsta/gin` 代码热更新工具
  * `github.com/jmoiron/sqlx` 对原生库sql的扩展，提供数据库字段映射
* 下载好预先制作好的docker镜像
  * `docker pull redis`
  * `docker pull postgres:9.4`
  * `docker pull zhihaojun/golang-web-scaffold`
  * `docker pull zhihaojun/phppgadmin`
  * **如果pull的速度很慢，请尝试使用国内的docker镜像加速器来加速**
  * [镜像加速](https://www.docker-cn.com/registry-mirror)




## 快速开始

1. clone这个仓库到本地 `git clone git@github.com:ZhihaoJun/GolangWebScaffold.git`
2. 在项目根目录运行 `docker-compose up -d`
3. 访问 [http://localhost:10060](http://localhost:10060/)
4. 访问 [http://localhost:10061](http://localhost:10061) 进入数据库管理
   * 用户名 `postgres`
   * 密码为空



## 开发

* 脚手架使用了gin来热更新加载代码，尝试修改[]()的代码，刷新网页可以看到即时更新的效果。
* 阅读docker-compose.yml可以获得更多关于配置的信息
* 已经内置了postgresql 9.4版本和redis最新版本
  * postgresql的访问域名是`postgres:5432`
  * redis的访问域名是 `redis:6379`
* **postgresql的版本是9.4，意味着不能使用on conflict子句，而选择这个版本是因为阿里云目前rds主要为9.4，他们适配的10.0版本还在计划中（捂脸**
* 版本管理，使用自己的仓库使用代码，务必重新 `git init`
* golang库依赖都是在容器内部，为了让ide能够正确提示和补全代码，在本机上也需要安装相应的golang依赖库，推荐使用 [govendor](https://github.com/kardianos/govendor) 工具完成依赖库管理。



## 部署

* [Docker Registry](https://docs.docker.com/registry/)



## 最佳实践

