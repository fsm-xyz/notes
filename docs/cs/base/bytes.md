# 字节

## 数据

数据存储在计算机中都是流的形式，即一串字节数据, 都是010101001的

## 字节序

字节的顺序，大于一个字节类型的数据在内存中的存放顺序。其实大部分人在实际的开 发中都很少会直接和字节序打交道。唯有在跨平台以及网络程序中字节序才是一个应该被考虑的问题。

## 网络序

先发送高字节数据再发送低字节数据（如：四个字节由高位向低位用0-31标示后，将首先发送0-7位第一字节，在发送8-15第二字节，以此下去）。IP协议中定义大端为网络字节序。

## 主机字节序

## 大端小端

1.概念
         大端：多字节的低位在内存的高位，高位在内存的低位
         小端：多字节的低位在内存的低位，高位在内存的高位

         比如：int a = 0x12345678
                           0x0001                  0x0002                   0x0003                  0x0004         内存由低变高
         大端             0x12                  0x34                   0x56                     0x78
         小端          0x78                  0x56                  0x34                  0x12                  

采用大端方式进行数据存放符合人类的正常思维,小端方式进行数据存放利于计算机处理
大端类似处理字符串，小端类似逆序放置字符串

2.现状

在网络上传输数据普遍采用的都是大端，而小端模式处理器的字节序到网络字节必须要进行转换
一般操作系统都是小端，而通讯协议是大端的。

- CPU的字节序
         Big Endian : PowerPC、IBM、Sun
         Little Endian : x86、DEC、Intel的Pentuim
         ARM既可以工作在大端模式，也可以工作在小端模式。
         MIPS等芯片要么采用全部大端的方式储存，要么提供选项支持大端——可以在大小端之间切换

- 常见文件的字节序
         Big Endian
                  Adobe PS
                  JPEG
                  MacPaint
         Little Endian
                  BMP
                  GIF
                  RTF
         Variable
                  DXF(AutoCAD)

- 编译器

在C语言中，默认是小端
Java和所有的网络通讯协议都是使用Big-Endian的编码

## 代码检测

方法1：

```sh
         #include <stdio.h>  
         int check()  
         {  
                  int i = 1;  
                  i = *(char*)&i;//取 i 的地址 强制类型转换后解引用  
                  return i;  
         }  
         int main(void)  
         {  
                  if(check()==1)  
                           printf("小端模式存储！\n");  
                  else  
                           printf("大端模式存储！\n");  
                  return 0;  
         }
```

方法2：

```sh
         #include <stdio.h>  
         int check()  
         {  
             union UN  
             {  
                 char c;  
                 int i;  
             }un;  
             un.i = 1;  
             return un.c;  
         }

         int main(void)  
         {  
             if(check()==1)  
                 printf("小端模式存储！\n");  
             else  
                 printf("大端模式存储！\n");  
             return 0;  
         }
```

方法3：
```sh
         #include <stdio.h>  
         int check()  
         {  
             union UN  
             {  
                 char a [4];  
                 int i ;  
             } un ;  
             un .i = 1;  
                 //02 是整数不够2位就补上0  x是以16进制输出  hhx 表示只输出两位  
             printf ("%02hhx %02hhx %02hhx %02hhx\n", un .a [0], un. a [1],un . a[2], un .a [3]);  
             return un . a[0];  
         }  

         int main(void)  
         {  
             if(check()==1)  
                 printf("小端模式存储！\n");  
             else  
                 printf("大端模式存储！\n");  
             return 0;  
         }
```

### 使用工具来查看顺序

od -b 单字节八进制显示
od -c ASCII码进行输出，其中包括转义字符
od -t 单字节十进制进行解释
hexdump

## 编码

ASCII 单字节编码
UTF16BE
UTF16LE
UTF-8 with BOM
UTF-8 withoutBOM
UTF-8 UTF（UCS Transfer Format) 变长编码 

### BOM

BOM是Byte Order Mark的缩写，它用来指明编码，如下所示：

BOM                         编码
FE FF                     UTF16BE
FF FE                     UTF16LE
EF BB BF                   UTF-8

### 字符集

UNICODE 多字节编码
为每一个「字符」分配一个唯一的 ID（学名为码位 / 码点 / Code Point）
「知」的码位是 30693，记作 U+77E5（30693 的十六进制为 0x77E5）。
UTF-8 顾名思义，是一套以 8 位为一个编码单位的可变长编码。会将一个码位编码为 1 到 4 个字节：U+ 0000 ~ U+ 007F: 0XXXXXXX
U+ 0080 ~ U+ 07FF: 110XXXXX 10XXXXXX
U+ 0800 ~ U+ FFFF: 1110XXXX 10XXXXXX 10XXXXXX
U+10000 ~ U+1FFFF: 11110XXX 10XXXXXX 10XXXXXX 10XXXXXX
