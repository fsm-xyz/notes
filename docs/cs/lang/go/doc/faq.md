# FAQ

>- Go中遇到“invalid character 'ï' looking for beginning of value”

[文本编码问题](https://stackoverflow.com/questions/31398044/got-error-invalid-character-ï-looking-for-beginning-of-value-from-json-unmar)

## cobra

现代化的命令行参数

### 格式讲解

cmd subcmd [flags]

cmd subcmd --help

## 作用域

:= 会生成个新的变量

```go
         a := 0
         if a := func(){}(); a > 0 {
    a是一个局部变量， 块作用域
         }
         a是一个全局变量

         a := 0
    if a = func(){}(); a > 0 {
    a 是全局变量
         }

         a := 0
         if a, b := func(); a > 0 {
             := 会生成个新的变量
    a 是局部
         }
```

v是一个副本即局部新变量，要想改变xxx里面的值需要可以使用下标a[i] 或者a []*这种slice

```go
for k, v := range map {
         k, v是一个局部变量，副本
}

for k, v := range a[] {
}
```

## Slice

基于同一个底层数组，减少某个值，引发底层数据发生错乱，如下：
基于同一个底层数组，增加某个值，会改变产生新的数组

```go
package main

import (
    "fmt"
)

func main() {
    a := []int{0,1,2,3,4,5,6,7,8,9}
    b := a

    b = append(b[:4], b[6:]...)
    fmt.Println(a,b)
}
```

### ERR

>- cannot assign to struct field xxx in map
原因是 map 元素是无法取址的
[stackoverflow](https://stackoverflow.com/questions/32751537/why-do-i-get-a-cannot-assign-error-when-setting-value-to-a-struct-as-a-value-i)

## 协程调度

## GC

### 回收算法

- 引用计数法
- 标记清除, 升级版三色标记法
- 分代收集

### 术语

- 生命周期
- 作用域
- 常量编译时期给定地址
- 堆栈
- 逃逸分析

### channel

遍历一个关闭的chan

```go
c := make(chan int)
c <- 6 // 发送数据, 向关闭的chan发送数据会panic
x :=  <- c // 从关闭的chan取出数据会得到0值
if x, ok := <- c; ok {} // 避免从关闭的chan取到0值
for range c {} // 关闭后退出for, 避免0值

// select 会随机取一个就绪的分支执行
// 如果c关闭则获取到的是0值(同简单取值一样)
for {
    select {
    case <- c:
    default :
    }
}

package main

import (
    "fmt"
    "time"
)

func main() {
    // c1()
    // c2()
    // c3()
    // c4()
    c5()
    time.Sleep(1 * time.Second)
}

// 向关闭channel发送数据
func c1() {
    c := make(chan int)
    go func() {
        for {
            fmt.Println(<-c)
        }
    }()
    c <- 6
    close(c)
    c <- 6
}

// 从关闭的channel获取数据
func c2() {
    c := make(chan int)
    go func() {
        for {
            fmt.Println(<-c)
        }
    }()
    c <- 6
    close(c)
}

// ok pattern模式
func c3() {
    c := make(chan int)
    go func() {
        for {
            if x, ok := <-c; ok {
                fmt.Println(x, ok)
            }
        }
    }()
    c <- 6
    close(c)
}

// for range
func c4() {
    c := make(chan int)
    go func() {
        for x := range c {
            fmt.Println(x)
        }
    }()
    c <- 6
    close(c)
}

func c5() {
    c := make(chan int)
    go func() {
        for {
            select {
            case x := <-c:
                fmt.Println(x)
            }
        }
    }()
    c <- 6
    c <- 6
    c <- 6
    close(c)
}

// select 就绪分支随机执行
func c6() {
    c := make(chan int)
    go func() {
        for {
            select {
            case x := <-c:
                fmt.Println(x)
            default:
                fmt.Println("default")
            }
        }
    }()
    c <- 6
    c <- 6
    c <- 6
    close(c)
}

```

使用一个chan 阻塞多个goroutine

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    stop := make(chan bool)

    go func() {
        time.Sleep(time.Second)
        close(stop)
        <-stop
        fmt.Println("Hello, playground1")
    }()

    go func() {
        <-stop
        fmt.Println("Hello, playground2")
    }()

    go func() {
        <-stop
        fmt.Println("Hello, playground3")
    }()
    time.Sleep(2 * time.Second)
    <-stop
    fmt.Println("Hello, playground")
}
```

### map

[并发读写](https://studygolang.com/articles/9537)

### goroutine

显式传值

```go
package main

import (
    "fmt"
    "time"
)

func main() {
         var m = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
         for k, v := range m {
    go func(k, v int) {
             fmt.Println(k, v)
    }(k, v)
         }
         for k, v := range m {
    go func() {
             fmt.Println(k, v)
    }()
         }
         time.Sleep(10 * time.Second)
}
```

格式化
报错: cannot use a (type []string) as type []interface {} in argument to fmt.Printf
fmt.Sprintf("%s%s", []string{"q","1"}...)
fmt.Sprintf("%s%s", []interface{}{"q","1"}...)

## 高级功能

- go:linkname sleep time.Sleep  必须引入unsafe包

## database

gorm 自动转换表名复数形式
sqlx 对于NULL自动换成对应类型的NULLTYPE类型, 需要处理转换为对应类型的nil
sqlx里面的unsafe:false可以保证字段的数量保持一致, true则不管

## 创建文件

os.ModePerm = 0777
unix unmask = 0022

所以都是filemode - unmask

创建文件使用filepath.Abs获取绝对地址

```go
# 创建~/1/2/3
os.MkDirAll("~/1/2/3", os.ModePerm)

# 得到的是当前目录下面的 "\~/1/2/3"
# 所以使用filepath.Abs("")获取绝对地址之后创建

## IO

ioutil.NopCloser 可以实现复用底层数据，不关闭

## http

go keepalive开启的话，会报EOF

刚好tcp的keepalive过期，导致EOF
