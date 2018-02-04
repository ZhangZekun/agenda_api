# Agenda

## 如何使用我们的镜像

### 1. 下载镜像

~~~go
docker pull 13829062426/agenda_api
~~~

### 2. 运行一个服务器Container

~~~go
sudo docker run --name agenda_server -d -p 8080:8080 13829062426/agenda_api service
~~~

这个服务器容器运行在后台，对外提供服务，可以通过localhost:8080进行访问。

### 3. 运行一个客户端Container

~~~go
sudo docker run --name agenda_cli_new -it --net=host 13829062426/agenda_api
~~~

因为客户端代码中访问api直接使用了localhost，所以这里的网络模式应使用--net=host.



运行该命令后，会直接进入container，我们可以运行如下命令来进行注册用户：

~~~go
agenda register -u kun4 -p kun4 -e kun4 -t kun4
~~~

但要注意的是，当进行这些命令操作时，最好进入到/go/src/github.com/ZhangZeMian/agenda_api/cli/agenda目录下，因为cookie.txt文件在这个目录下。

其他命令操作也是同样。



## 如何使用go get从github获取我们项目代码

### 1. 前言

为了让大家方便地获取我们项目以及项目的依赖包，我们在项目的首页加了一个main.go文件，这样方便大家直接用go get github.com/ZhangZekun/agenda_api获取项目代码和依赖包。

### 2. 使用go get获取项目包

~~~go
 go get -u -v github.com/ZhangZekun/agenda_api
~~~

因为依赖包里面有golang.org相关包，所以会遇到被墙了的问题。这里需要设置代理，主要分为两步：下载lantern以及设置代理。

### 3. 下载Lantern

下载安装方法见：

[lantern安装下载](https://www.jianshu.com/p/d62cd99961c3)

### 4. 设置代理

设置代理方法见：

[设置终端lantern代理](https://dade.io/archives/116/)

### 5. 设置代理后使用go get获取项目包

~~~go
 go get -u -v github.com/ZhangZekun/agenda_api
~~~

### 6. go install 客户端和服务端

进入cli/agenda/cmd，执行go install

进入service，执行go install

然后就可以直接执行service，运行一个server后台。

执行agenda login -u XXX -p XXX进行登录。

### 7. 构建docker image

在项目总目录下执行：

~~~go
docker build -t agenda_api .
~~~

这里同样会有一个要翻墙的问题，因为build过程中需要引入依赖包，其中就包括golang.org/x系列包。

docker是一个CS结构，即docker命令会被发送到本地运行的一个服务器daemon执行。所以需要为daemon服务器设置代理，以实现安装golang.org/x包。

设置代理方法如下：

[设置daemon代理](https://docs.docker.com/config/daemon/systemd/#httphttps-proxy)

## 运行结果展示

### 1. 拉取docker image

![](https://raw.githubusercontent.com/ZhangZekun/images/master/agenda_api/1.PNG)

### 2. 运行一个服务器containe

![](https://raw.githubusercontent.com/ZhangZekun/images/master/agenda_api/3.PNG)

### 3. 运行一个客户端contrainer

![](https://raw.githubusercontent.com/ZhangZekun/images/master/agenda_api/2.PNG)

### 4. travis测试结果

![](https://raw.githubusercontent.com/ZhangZekun/images/master/agenda_api/4.PNG)