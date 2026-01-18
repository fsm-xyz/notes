# Go Programming Language

前言

    一·语言特性

        1.特性
            垃圾回收
            包系统
            一等函数    len make new copy cap delete
            词法作用域
            系统调用接口
            UTF-8字符串编码
        2.不支持
            隐式数值类型强制转换
            析构和构造函数
            运算重载符
            形参默认值
            继承
            泛型
            异常
            宏
            函数注记
            线程局部存储

TUTORIAL

    一·go 工具链
        go run            complie + execute
        go build        compile + save
        go install        compile + save(bin)
        go get            git clone + go install

        gofmt
        goimports

    二·关键字
        functions        func
        variables        var
        constants        const
        types            type
    三·表达式
        无表达式    j = i++
        没后缀        --i
    四·琐碎的知识点
        os.Args[]
        range   _
        strings.Join(os.Args[1:], " ")
        strings.Split()

        ioutil.ReadFile()
        ioutil.ReadAll()
        io.Copy()
        strconv.Atoi()

        变量声明, 实际中使用的情况
        s := ""                    短变量声明, 在方法内使用，不在包级别使用
        var s string             默认zero value
        var s = ""                多个变量
        var s string = ""        类型不同

程序结构

    一·声明
        var const type func
        作用域
            局部级别    package-level
            包级别        local
    二·变量
        变量声明
            var name type = expression
        短变量声明 适用于local variables
            name := expression
        指针
            指向局部变量的指针，返回，仍然有效
        通过内建new创建变量
            new(T)        unamed variable of type T, initiakizes zero value of T, return  its address, which is a value of type *T
            p := new(int)    0
            *p = 2
        生命周期
            包级别的是在整个执行过程
            local是动态的
            变量是否一直有效,取决于他是否可达(reachable)

            head-stack, memory-allocation
        赋值
            =
            元祖赋值
                x, y = y, x

                v, ok = m[key]                    //map lookup
                v, ok = x.(T)                    //type assertion
                v, ok = <-ch                    //channel receive
                _, err = io.Copry(dst, src)     //blank identifier 
            赋值性
                同类型
                nil <=> interface, reference

        Type Declarations
            type name underlying-type
            类性转换
                改变变类型
                T(x)    //converts x To type T, both have the same underlying type, or unnamed poiter types that point to variables of  the same underlying type 
                数值, 字符串, some slice types各自互转
                类型不同不能直接比较, 但可以和underlying type比较
            type's method
                给自定义的类型加上方法
                    func (c Celsius) String() string{ return fmt.Sprintf("%g℃", c)}

        包和文件
            是否被包外访问, 大小写
            包的初始化
                package-level 变量开始
                引用的包优先
                func init() 默认调用, 不可调用引用
                在main之前, 所有的包必须初始化

        作用域, 生命周期
            The scope of a declaration is a region of the program text; it is a compile-time property.
            The lifetime of a variable is the range of time during execution when the variable can be referred to by other parts of the program; it is a run-time property.

            块作用域
                for if switch select  {}

基础数据类型

    Go four types:
        basic types: numbers, strings and boolean
        aggregate types: arrays, structs
        reference types: pointers, slices, maps, functions and channels
        interface types

    Integer
        int8, int16, int32, int64
        uint8, uint16, uint32, uint64

        run <==> int32    byte <==> int8    unsigned integer type uintptr

        %b     %[1] %#[1]
    Folat
        float32, float64
    Complex
        complex64, complex128
    Booleans
        true, false

    Strings
        1.不可变的字节序列, immutable
        2.字符字面量 
        3.utf8.RuneCountInString(s)   utf8.DecodeRuneInString
        4.重要的4个库   strings, bytes, strconv, unicode

            \xhh         对应小于256的码点
            \uhhhh         对应16位的码点
            \Uhhhhhhhh    对应32为的码点

            0x        \xhh    16进制
            0        \ooo    八进制

        字符串和数值互转
            integer to string
                x := 123
                y := fmt.Sprintf("%d", x)
                fmt.Println(y, strconv.Itoa(x))                  //"123 123"
                fmt.Println(strconv.FormatInt(int64(x), 2))        //"1111011"
            string to integer
                x, err := strconv.Atoi("123")
                y, err := strconv.ParseInt("123", 10, 64)

        常量
            basic type: boolean, string, number
            compile time
            constant generator    iota

            enums    枚举

            Print         Println         Printf        //输出到标准输出
            Fprint      Fprintln         Fprintf        //输出到指定位置
            Sprint         Sprintln         Sprintf        //输出为字符串

Composite Types

    aggregate types: arrays, structs                                   //fixed size
    referrence types: maps, slices, pointers, channels, functions     //maps slices dynamic size

    数组
        数组类型
        [3]int, [...]int{1, 2, 3}

        类型转换
            []byte("X")  <=> "x" convert to []byte
        Go里面的数组是按值传递的, 其他语言是按指针或者引用传递的
            func zero(arr [3]int)
            arr 是值拷贝而不是指向原始的数组
            要指向原始的, 通过指针来实现

    Slices
        three components：
            poiter, length, capacity
            poiter: a poiter to an element of an array
            length:    the number of slice elements
            capacity: the start of slice and the end of underlying array

        不能直接用==去比较slice是否相同的元素

        bytes.Equal用于比较[]byte, 其他类型的需要我们自定义比较

        append(dst, src) 当len + 1 >  cap时会动态改变slice的大小, 重新分配空间,改变原始的指向

    Maps
        底层是hash table, k/v
        k必须是可比较的
        maps是不能比较的
        每次迭代结果是无序的
        1.    ages := make(map[string]int)
        2.    ages := map[string]int {"alice": 31, "charlie": 33,}
        3.    make(map[string]map[string]bool)    嵌套的map

        ok模式判断是否存在该键值
        if age, ok != ages["ages"]; !ok {}

    Structs 
        按值传递
        指针
            type Point struct{X, Y int}

            p = &Point{1, 2}  <=> pp := new(Point)  *pp = Point{1, 2}
            q = Point{1, 3}

            p.X  q.X     (*p).X
        可以比较
        可以做map的key
        struct embedding and anonymous fields

    Text and HTML Templates
        html: 自动转义元字符, template.HTML会不自动转义

函数

    函数签名:参数, 返回值
    按值传递
    不可比较
    有自己的类型

    递归:
        别的语言采用固定大小的栈来递归, Go用动态大小的栈实现递归调用

        判断结束条件, 处理要写在方法开头, 然后再递归调用

    多返回值
        Gc 会回收未使用的内存, 但是不会回收系统资源或者网络连接, 需要明确的关闭
        resp.Body.Close()
        _ 可以忽略使用

        带名字的返回值可以直接return, 不用跟其他的
            bare return <=> return each of the named result variables

    错误

    函数值(类似js中的函数)
        有自己的类型
        可以赋值给其他变量
        可以当作参数
        var f = func{}

    匿名函数
        闭包, 函数字面量, 做参数的函数,函数变量
        捕获迭代变量
            在循环结束前, 回调函数会延迟执行, 所以变量的值已经是最后的值了

    变长参数
        val ...int
    defer
        在return之前,执行该函数, 多个defer按照栈的顺序执行

    panic
        compile time: 类型错误
        run time: 数组越界, nil pointer dereference

        会执行以前定义的defer内容

        用于严重的错误, 其他的可预知的错误使用以前的错误机制(expected), 避免程序员的崩溃


    recover

        在defer中调用recover, deferred函数内部得到panic的信息

Methods

    OOP
        an object is simpley a value or variable that has methods, and a method is a function associated with a particular type.

    Two principles
        1.encapsulation        封装
        2.composition        组合

    声明
        1.方法接收者
        2.选择器
        可以给任何类型定义方法, 除了pointer, interface
    带指针接收器的方法
        避免值拷贝, 浪费资源
        在声明方法时，类型名本身不能是指针类型, 避免歧义  
        
            type P *int
            func (P) f() {} //编译报错

    nil is a valid receiver value

    组装类型通过嵌入结构体

    方法值和表达式
        方法值: p.Distance
        表达式: Ponit.Distance    T.f    *T.f

    例子: 位向量类型

    封装
        struct

接口

    接口作为合约
        满足接口中的合约, 即为该接口类型

    Interface Types
        接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

        接口嵌入

    接口满足
        一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口
        接口指定
            var w io.Writer
            w = os.Stdout           // OK: *os.File has Write method
            w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
            w = time.Second         // compile error: time.Duration lacks Write method

            var rwc io.ReadWriteCloser
            rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
            rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

            w = rwc                 // OK: io.ReadWriteCloser has Write method
            rwc = w                 // compile error: io.Writer lacks Close method

            右边比左边的大可以赋值, 反之不行
    接口值
        包含具体的动态类型和那个类型的动态值
        类型描述符

        var w io.Writer     //声明接口
        w = os.Stdout        //赋值

        类型断言
    接口的2種用途
        1.用于隐藏具体类型的细节
        2.用于可识别联合, 用空接口判断类型 interface{}

Goroutinues And Channels

    goroutine
    Channels
        ch = make(chan int)    create
        ch <- x         send
        x = <-ch         receive
        close(ch)        close

        make(chan int)    unbuffered
        make(chan int, 3) buffered

Concurrency with Shared Variables

    Race Condition
        数据竞争会在两个以上的goroutine并发访问相同的变量且至少其中一个为写操作时发生。
        根据上述定义，有三种方式可以避免数据竞争
            1.是不要去写变量
            2.避免从多个goroutine访问变量
            3.允许很多goroutine去访问变量, 但是在同一个时刻最多只有一个goroutine在访问(“互斥”_)
    Mutex
        sync.Mutex
        sync.RWMutex
        RLock只能在临界区共享变量没有任何写入操作时可用。
        RWMutex只有当获得锁的大部分goroutine都是读操作，而锁在竞争条件下，也就是说，goroutine们必须等待才能获取到锁的时候，RWMutex才是最能带来好处的。RWMutex需要更复杂的内部记录，所以会让它比一般的无竞争锁的mutex慢一些。

    Memory Synchronization
        所有并发的问题都可以用一致的、简单的既定的模式来规避。所以可能的话，将变量限定在goroutine内部；如果是多个goroutine都需要访问的变量，使用互斥条件来访问。

    Lazy Initialization
        sync.Once
    Race Detector
        -race
    解决方案:        
        1.基于互斥量的版本
        2.基于单独的monitor(channel) goroutine

    Goroutines vs Threads
        1.动态栈
            2M     Thread fixed size stack
            2K     Goroutine dynamic size stack
        2.Scheduling
            OS kernel, hardware timer, context switch
            Runtime, channel, mutex (m:n)
        3.GOMAXPROCS
            使用的最大的OS thread个数, 默认CPU核心数
        4.Have no Indentify
            thread-local storage

Packages and the Go tool

    Imports Paths

    The Package Declaration

    Import Declarations
        alternative name
    Blank Imports

    Packages and Naming
        包名一般单数,

    Go Tool
        Workapace Organization
            $GOPATH
        Downloading Packages
            真实地址
        Building Packages

        Documenting Packages
            go doc
            godoc    以web显示文档
        Internal Package
            限制能被否被外面的包访问
        Querying Packages
            go list

Testing

    Test
    Benchmark
    Example

    产品代码
    包内测试
    外部测试包

    测试覆盖
        go test -coverprofile=c.out
        go tool cover -html=c.out

Reflection

    reflect.Type: 一个接口, reflect.TypeOf 接受任意的 interface{} 类型, 并以reflect.Type形式返回其动态类型
    reflect.Value: reflect.ValueOf 接受任意的 interface{} 类型, 并返回一个装载着其动态值的 reflect.Value.

    dynamic type and value

    reflect.Value 和 interface{} 都能装载任意的值. 所不同的是, 一个空的接口隐藏了值内部的表示方式和所有方法, 因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值(就像上面那样), 内部值我们没法访问. 相比之下, 一个 Value 则有很多方法来检查其内容, 无论它的具体类型是什么

Low-Level Programming

    unsafe
    cgo
