# shell

## 查看帮助

* man
* info
* help
* command -h or --help

## 命令行格式

```bash
# command, option, arguments
ls -l /
```

-c 把option不当作args[], $i的参数

## 分号(;)

大多数编程语言用";"标识这是一个语句，sh文件中可以忽略";"
终端命令行里面没法省略
编辑shell脚本的时候可以用"\"来实现多行书写, 下面的多行会被压缩成一行

```bash
echo 'Hello,World' > a.txt; \
while read line; \
do echo $line; \
done < a.txt
```

等同于

```bash
 echo 'Hello,World' > a.txt; while read line; do echo $line; done < a.txt
```

## 赋值

shell 变量赋值不能空格再赋值

```bash
 a=1 // right
 a= 1 // wrong
 a = 1 // wrong
```

## 特殊字符

反引号\`cmd\`, 单引号'', 双引号""
    1. \`date\` 执行date命令  
    2. 单引号默认去掉特殊字符含义
    3. 双引号类似，忽略大多数，($,\,`)字符例外

## 文本处理

> awk, sed, find, grep, sort, uniq
> join 合并2个文件行
> paset 合并2个文件列
paste -d "\t" eng.txt chi.txt
> 根据第一列相同显示
awk 'NR==FNR{a[$1]=$0;next}NR>FNR{if($1 in a)print $0}' 1.log 2.log
awk '{FNR==NR{a[FNR]=$1}NR>FNR{print $1"\t"a[FNR]}}' a.txt b.txt > test.txt

## curl

### query string

如果有&符号需要\转移一下, 或者加引号括起来
    ?a=a&b=c
    ?a=a\&b=c
    curl 'google.com?a=a&b=c'

```shell
# curl模拟post
curl -i -X POST -H "'Content-type':'application/json'" -d '{"user":"atime","company":"btime"}' http://localhost:8080/add
```

## 结束进程

```shell
killall
kill -9
```

## 修改文件权限

chmod -R 0755 /var/www/website
find /var/www/website -type f -exec chmod 0644 {} \;

find -type d|xargs chmod 745
find -type f|xargs chmod 644

## 进程

```sh

# ps
# UNIX 风格，选项可以组合在一起，并且选项前必须有“-”连字符
# BSD 风格，选项可以组合在一起，但是选项前不能有“-”连字符
# GNU 风格的长选项，选项前有两个“-”连字符

# 查看某个进程的线程情况
ps -T -p pid
ps -T pid
# 查看进程
ps -ef
# 结果显示全
ps -efww

# 显示进程树
ps -ef ef

pstree -l -a -A 20708

```
## 端口查看

ss -tpan
ss -s

```shell
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
```

## io

### 特殊标识

* `0` 标准输入
* `1` 标准输出
* `2` 标准错误输出
* `/dev/null` 类似于回收桶, 传给他的数据都将丢弃

```text
>  输出重定向
>> 输出重定向，追加的方式
<  输入重定向
<< 标准输入来自命令行的一对分割号的中间内容
|  管道，虚拟的文件，进程通信
```

```shell
# n> n和>必须挨着,才有特殊含义
cat 1.txt > 2.txt       # 标准输出重定向 >
cat 1.txt 1> 2.txt      # 标准输出重定向 1>
cat 1.txt 2> 2.txt      # 标准错误重定向 2>
cat 1.txt &> 2.txt      # 把1和2都放到file文件

# n>&m &m表示文件描述符
cat 1.txt >2.txt 2>&1   # 标准重定向, 然后把2放到1里面
cat 1.txt 2>r.txt 1>&2  # 标准错误重定向, 把1放到2里面
```

### EOF(end of file)

linux上ctrl+d就代表EOF

#### 分界符

以下的EOF可以替换成ABC

```sh
cat <<EOF        # 开始
xxxx             # 输入内容
EOF              # 结束
```

```sh
cat <<EOF >1.txt
>1
>2
>3
>EOF
```

```sh
cat > 1.txt <<EOF
>1
>2
>3
>EOF
```

```sh
# 自定义的
cat <<ABC        #开始
xxxx
ABC              #结束
```

```sh
cat <<EOF | tee 1.txt
>123
>EOF
```

#### `<<EOF`和`<<-EOF`

如果重定向的操作符是<<-，那么分界符（EOF）所在行的开头部分的制表符（Tab）都将被去除
