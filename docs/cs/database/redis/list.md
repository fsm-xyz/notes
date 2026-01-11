# list

## 场景

+ 队列
+ 评论

## 注意点

### blpop阻塞读取，连接断开，应用程序不能继续读取

redis在blpop命令处理过程时，首先会去查找key对应的list，如果存在，则pop出数据响应给客户端。否则将对应的key push到blocking_keys数据结构当中，对应的value是被阻塞的client。当下次push命令发出时，服务器检查blocking_keys当中是否存在对应的key，如果存在，则将key添加到ready_keys链表当中，同时将value插入链表当中并响应客户端。

服务端在每次的事件循环当中处理完客户端请求之后，会遍历ready_keys链表，并从blocking_keys链表当中找到对应的client，进行响应，整个过程并不会阻塞事件循环的执行。所以， 总的来说，redis server是通过ready_keys和blocking_keys两个链表和事件循环来处理阻塞事件的。