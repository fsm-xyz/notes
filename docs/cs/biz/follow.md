# 关注

## 概念

```text
a和b之间的关注关系:
    a单向关注b
    b单向关注a
    a和b互相关注
    a和b毫无关联
```

a和b的拉黑关系:

有关注, 就有了粉丝

我关注了别人, 我的粉丝有哪些, 和我互粉的有哪些人

## 不分表分库

```text
一张表, 1条记录
当uid, to_uid等价于to_uid, uid, 增加status

CRATE TABLE follow (
    uid
    to_uid
    status
)

uid < to_uid
select *from follow where uid = low and touid = high

select *from follow where (uid = low and status = 1) or  (touid = high and status =2)

0 无关系
1 单向关注 a->b
2 单向关注 a<-b
3 互相关注 a<->b

一张表, 2条记录
当uid, to_uid不等价于to_uid, uid

CRATE TABLE follow (
    uid
    to_uid
)

我的关注    select * from follow where uid = 1
我的粉丝    select * from follow where touid = 1
相互关注    (select *from follow where uid = 1 and touid = 2) and (select* from follow where uid = 2 and touid = 1)

2张表

follow表
fan表

不分表分库的话，这个基本是没意义的, 单表就能解决关注和粉丝
```

## 分表分库

follow表, fan表

uid 来做分表的字段

查询自己的关注, 粉丝
