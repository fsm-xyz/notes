# file mode

ls -al
-rw-r--r--@

第一个表示文件类型, User, Group, Other, @表示extended attributes.

## 文件权限

chmod下面的0755, 可以通过设置第一位的值，设置可执行文件的权限
4: SUID permission      -r-sr-xr-x
2: SGID permission      -rwxr-sr-x
1: Sticky bit           -rwxrwxr-t

chmod会把777当做八进制处理,同0777的

在程序里面以0开头表示这是一个八进制，0755映射chmod的

## 文件权限掩码

默认的掩码0022，
文件的权限 = 文件权限 - 文件权限掩码

```go
func main() {
    oldMask := syscall.Umask(0)
    os.MkdirAll("abc/a", a)
    syscall.Umask(oldMask)
}
```

### 参考资料

[1](https://www.thegeekdiary.com/what-is-suid-sgid-and-sticky-bit/)

[2](https://digitalfortress.tech/php/difference-file-mode-0777-vs-777/)

[3](https://unix.stackexchange.com/questions/103413/is-there-any-difference-between-mode-value-0777-and-777)

## FAQ

如果一个目录的缺少写权限，则不能删除下面的文件
如果缺少执行权限，则不能读取下面的文件
