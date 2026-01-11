# server

## web服务器模型

* 单进程I/O模型

    服务端开启一个进程，一个进程仅能处理一个请求，并且对请求顺序处理

* 多进程I/O模型

    服务端并行开启多个进程，同样的一个进程只能处理一个请求，这样服务端就可以同时处理多个请求

* 多线程I/O模型

    服务端开启一个进程，但是呢，同时开启多个线程，一个线程响应一个请求，同样可以达到同时处理多个请求，线程间并发执行

* 复用多线程I/O模型

    服务端并行开启多个进程，同时每个进程开启多个线程，这样服务端可以同时处理进程数M*每个进程的线程数N个请求

* 多路复用

## CA证书

### 名词解释

* DV Domain validtion
* OV Orgnaztion validtion
* EV Extend validtion

### 安装Lets'Script

官网有文档利用certbot

### 常见问题

1. python需要先安装ssl等才能用pip，不然得重新编译python
2. The requested nginx plugin does not appear to be installed
系统是python3，而certbot-2指定的是2，不一致
3. ImportError: 'pyOpenSSL' module missing required functionality.
rpm -Uvh http://cbs.centos.org/kojifiles/packages/pyOpenSSL/0.15.1/1.el7/noarch/pyOpenSSL-0.15.1-1.el7.noarch.rpm 
4. AttributeError: 'module' object has no attribute 'SSL_ST_INIT' 解决办法

```sh
pip uninstall pyopenssl
pip install pyopenssl
```
