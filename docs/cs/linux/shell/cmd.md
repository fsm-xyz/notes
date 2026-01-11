# 命令

## 工具

```sh
htop
top
vmstat
glances
netstat
iproute2
dig
nslookup
drill
curl
iputils
rsync
tcpdump
brctl
chpasswd
rz
nc
tcpdump port 80 host 127.0.0.1
mtr  10.111.47.27
pidstat -u -l 1
```

## GUI

* finalshell
* xShell
* SecureCRT
* Putty

## 复制

```bash
# 递归拷贝
cp -Rf *
# 递归查找
grep -R '' *

gzip

bgrep
bcat
```

## 切割

### 方案1

```sh
x="a,b,ci"
OLD_IFS="$IFS"
IFS=","
arr=(${x})
IFS="$OLD_IFS"
for s in ${arr[@]}
do
    echo "$s"
done

echo "${arr[0]}"
```

### 方案2

```sh
#!/bin/bash
string="hello,shell,haha"  
array=(${string//,/ })  
for var in ${array[@]}
do
   echo $var
done
```

#### 缺陷

1. 会把空格也当成分隔符
