# 参数

```sh
$0 脚本文件名称
$1 第一个参数
$# 参数个数
$$ 脚本当前PID
$! 最后一个进程ID
$- 同set,显示当前选项
$? 显示退出状态
$@ 以一个单字符串显示所有向脚本传递的参数
$* 以一个多字符串显示所有向脚本传递的参数

shell 变量包含空格
p='abc 123'
"${p}" abc 123
'${p}' ${p}
${p} abc
echo -e 换行  ./需要加-e, sh run.sh 不需要
echo -e hello,world\n
```

## 调试shell脚本过程

[参考](https://www.ibm.com/developerworks/cn/linux/l-cn-shell-debug/index.html)
