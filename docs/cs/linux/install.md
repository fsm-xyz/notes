# Linux软件的安装方式

* 源码包安装
* 二进制安装
  * ReadHat系列rpm,yum安装(新增一种yum的变体dnf，类似yum)
  * Debian系列dpkg,apt-get安装(ppa,copr)
  * 通用的二进制包安装

## 区别

* 安装的位置不同
* 启动方式不同

## 源码安装

* 下载对应的源码包
* 解压缩
* 进入解压缩的目录
* 执行./configure --prefix=/usr/local/filename
* make 失败可以用make clean清除
* make install

## rpm安装

```sh
1.相关知识
  包依赖：树形依赖，环形依赖，库依赖
  查询模块依赖的网站：www.rpmfind.com
2.rpm介绍及使用
  a.rpm相关知识
    包的格式： 包名-发行版本-修正版-Linux平台-硬件平台
    安装的信息： /var/lib/rpm 保存包的安装信息
  b.rpm安装：
    rpm -ivh 包全名
      选项：
        -i： install
        -v： verbose 显示详细信息
        -h： #显示进度
        -t:  测试安装
        -p:  百分比进度
        -f:  忽略冲突
        --nodeps: 不检测依赖性（不建议加）
  c.rpm更新
    rpm -Uvh 包名
      选项
        -U upgrade
        如果存在则升级，不存在则安装
  d.rpm卸载
    rpm -e 包名
      选项
        -e erase
  e.rpm查询
    rpm -q 包名  查询指定包
    rpm -qa     查询所有已安装的软件包
      选项
        -q query
        -a all
        -i information
        -p package 查询未安装的包的信息
        -l location 查询软件包的安装位置
        -f 查询文件属于那个rpm包
        -R 查询依赖那些包
  f.rpm校验
    rpm -V 包名 查询包是否被修改过
      显示信息：
        S： 大小改变
        M： 文件权限改变
        5： 文件MD5校验改变
        D： 设备的主从代码
        L:  文件路径
        U:  文件所有者改变
        G： 文件的属组
        T： 文件的时间
      文件类型：
        c config file
        d documentation
        g ghost file
        L license file
        r readme file
  g.rpm包中文件的提取
    rpm2cpio 包全名 | cpio -idv 文件的路径
      rpm2cpio 将rpm包转换为cpio格式
      cpio  用于创建档案文件和从档案文件中提取文件
      选项
        -i  copy-in模式 还原
        -d  还原时自动新建目录
        -v  显示还原过程
```

## yum安装

```sh
  基于rpm，将所有软件包放到官方服务器，自动解决依赖性问题
  yum源：
    网络源： /etc/yum.repos.d/
    本地源:  挂载光盘
  查询
    yum list            显示所有能安装的包
    yum search  关键字   查询和关键字相关的包
    yum provides realplay 查询包含特定文件的包
    yum list installed  显示已安装的
    yum list updates    显示可更新的
    yum list extras     显示已安装的但不在资源库中的包
  列举资源信息
    yum info 包名
    yum info updates
    yum info installed
    yum info extras

  安装
    yum install 包名     安装
      -y  自动yes
      --downloadonly
  更新
    yum update
    yum upgrade     （大规模的版本升级,与yum update不同的是,连旧的淘汰的包也升级）
    yum update 包名
    yum check-update
  清除
    yum clean 参数
    yum clean packages
    yum clearn headers
    yum clean oldheaders
    yum clean all
  卸载
    yum remove 包名
      -y 自动yes
      服务器最小化安装，尽量不卸载
  只下载
    yum install --downloadonly --downloaddir=<directory> <package>
  yum组操作
    yum grouplist                     列出可用的软件组列表
    yum grouplistinstall 软件组名       安装指定组
    yum groupupdate      软件组名       更新
    yum groupremove      软件组名       卸载指定组

## rpm转deb包：
    在ubuntu上使用sudo apt-get install alien安装alien，用alien -d *.rpm可以转换为对应的deb
```
