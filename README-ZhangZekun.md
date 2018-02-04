# 张泽坤个人报告

张泽坤 15331410

## fork 项目位置

https://github.com/ZhangZekun/agenda_api

## 个人贡献简要说明

1. 所有 数据库相关代码的实现（entity, dao, service）
2. 所有 Meeting相关的api的实现 
3. 所有 客户端代码的实现（Commands） 
4. 完成、测试客户端和服务端的对接
5. 部署travis，解决travis遇到的import路径错误问题
6. 部署docker，构建docker image
7. 写README

## 完成项目所查阅收藏的文档、博客

![5](https://raw.githubusercontent.com/ZhangZekun/images/master/agenda_api/5.PNG)

## 项目小结

这次项目让我学到了很多东西。

客户端和服务端的编写难度不算特别大，但练手还是很不错的。

travis部署以及docker部署这两个工作，让我学到最多。

因为travis部署，需要去了解travis的工作机制，当遇到引用包路径错误时，我尝试着从travis的工作流程去发现问题，并通过google找到类似问题的解决方法，这类解决问题的思想对以后的学习工作会有帮助。

而docker部署，因为在国内，需要考虑防火墙的问题，所以需要设置代理、下载vpn等，而在实际操作过程中会出现浏览器自动设置代理，而终端以及docker server需要手动设置代理，这让我对代理这一机制有了更深的了解。

而容器技术有点类似虚拟机而又不尽相同，其中涉及的文件系统、端口映射、网络模式设置等，对我理解虚拟机的工作原理有很大帮助。

学到了不少东西，感谢老师和TA~