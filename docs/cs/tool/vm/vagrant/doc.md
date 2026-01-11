# vagrant

vagrant是一款可以快速搭建虚拟机, 自动调用VirtualBox、VMware等虚拟软件

## boxs

[boxs](https://app.vagrantup.com/boxes/search)
[details](https://www.vagrantup.com/intro/getting-started/boxes.html)

## 命令

```sh
vagrant box list
vagrant box add {title} {url}
vagrant init {title}
vagrant up
```

安装centsos8

```sh
vagrant init generic/centos8
vagrant up
vagrant ssh
```

```sh
vagrant init  # 初始化
vagrant up  # 启动虚拟机
vagrant halt  # 关闭虚拟机
vagrant suspend # 挂起
vagrant reload  # 重启虚拟机
vagrant ssh  # SSH 至虚拟机
vagrant status  # 查看虚拟机运行状态
vagrant destroy  # 销毁当前虚拟机
```
