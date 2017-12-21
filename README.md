# Golang Web Scaffold
开箱即用的golang echo框架脚手架

**献给所有冰岩人，愿能为你的伟大创造添砖加瓦**

## 准备工作
* 使用git做版本控制：[Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* 开发、部署环境的配置都使用Docker：[Install Docker](https://docs.docker.com/engine/installation/)
* 使用docker-compose，现在桌面操作系统上docker-compose已经和docker打包安装，无须独立安装，不过请确保已经安装，尤其是使用linux操作系统的：[Install Docker Compose](https://docs.docker.com/compose/install/)
* 下载好预先制作好的docker镜像
  * `docker pull redis`
  * `docker pull postgres:9.4`
  * `docker pull zhihaojun/golang-web-scaffold`
  * `docker pull zhihaojun/phppgadmin`
  * **如果pull的速度很慢，请尝试使用国内的docker镜像加速器来加速**
  * [镜像加速](https://www.docker-cn.com/registry-mirror)
* golang环境已经集成在docker容器中，包含以下库
  * `github.com/labstack/echo` 
  * `github.com/utahta/echo-sessions`
  * `github.com/boj/redistore`
  * `github.com/garyburd/redigo`
  * `github.com/codegangsta/gin` 代码热更新工具


## 快速开始
1. clone这个仓库到本地 `git clone git@github.com:ZhihaoJun/GolangWebScaffold.git`
2. 在项目根目录运行 `docker-compose up -d`
3. 访问 [http://localhost:10060/?name=Bingyan](http://localhost:10060/?name=Bingyan) 看到 `Hello, Bingyan!`
4. 访问 [http://localhost:10060/count](http://localhost:10060/count) 看到 `Page has been visited 1 times` 刷新可以看到次数在增加
4. 访问 [http://localhost:10061](http://localhost:10061) 进入数据库管理
   * 用户名 `postgres`
   * 密码为空


## 开发
* 脚手架使用了gin来热更新加载代码，尝试修改代码，刷新网页可以看到即时更新的效果。
* 阅读docker-compose.yml可以获得更多关于配置的信息
* 已经内置了postgresql 9.4版本和redis最新版本
  * postgresql的访问地址是`postgres:5432`
  * redis的访问地址是 `redis:6379`
* **postgresql的版本是9.4，意味着不能使用on conflict子句，而选择这个版本是因为阿里云目前rds主要为9.4，他们适配的10.0版本还在计划中（捂脸**
* 版本管理，务必删除`.git`文件夹，并重新设置自己的仓库地址
* golang库依赖都是在容器内部，为了让ide能够正确提示和补全代码，在本机上也需要安装相应的golang依赖库，推荐使用 [govendor](https://github.com/kardianos/govendor) 工具完成依赖库管理。


## 部署
* **不应该直接使用golang-web-scaffold容器放到线上部署**
* 开发完成的代码应该构建出一个新的镜像，push到类似[Docker Registry](https://docs.docker.com/registry/)的私有镜像管理的地方，再进行部署
* 根目录下的`Dockerfile`文件是一个实例的构建文件，可以直接使用
* `Dockerfile` 默认为你打开了1323端口


## 最佳实践
* 如果需求替换已经在`golang-web-scaffold`镜像中的库，请参考使用 [govendor](https://github.com/kardianos/govendor) 或者 [glide](https://github.com/Masterminds/glide) 工具
* 最好是自己修改Dockerfile的构建命令，脱离`golang-web-scaffold`镜像，直接从`golang:1.9`构建，并且自己安装所需库的依赖
