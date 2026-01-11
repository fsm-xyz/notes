# ssh

## 启动sshd

```sh
/usr/sbin/sshd
```

## 密码登录

PermitRootLogin yes         // 允许root用户登录
PasswordAuthentication yes  // 允许密码登录

## ssh登录

```bash
 ssh user@host
 ssh -p 2222 user@host // 指定端口号
```

远程登陆不要密码

```bash
ssh-keygen -t rsa
// 手动copy
cat ~/.ssh/id_rsa.pub >> authorized_keys // 添加公钥到服务器的~/.ssh目录下authorized_keys

# 简单的方式
ssh-copy-id -i ~/.ssh/mykey user@host
```

### FAQ

1. authorized_keys不生效的解决方法
2. 设置.ssh的目录700
3. 设置.authorized_keys文件的权限600

## 常见问题

### 启动sshd报错

+ sshd re-exec requires execution with an absolute path
使用sshd的绝对路径启动服务

+ sshd: no hostkeys available -- exiting.

```sh
ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key
ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key
```

## 安全

出于安全性考虑，只允许通过ssh key登陆并禁用了root登陆。

```bash
sudo su
vi /etc/ssh/sshd_config //编辑文件
# Authentication:
LoginGraceTime 120
PermitRootLogin yes //默认为no，需要开启root用户访问改为yes
StrictModes yes

# Change to no to disable tunnelled clear text passwords
PasswordAuthentication yes //默认为no，改为yes开启密码登陆

/etc/init.d/ssh restart
```

2.端口
设置特定端口开放

## 代理

-L local_socket:remote_socket

-N  Do not execute a remote command.  This is useful for just forwarding ports.
