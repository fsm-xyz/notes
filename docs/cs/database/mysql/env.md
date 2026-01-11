# MySQL

## 设置用户和权限

```bash
// 选择数据库
use mysql;
// 查看用户
select host, user, password from user;
// 创建用户
insert into user(host,user,password) values("localhost","remote",password("qwe1234"));
// 授权
grant all privileges on `test`.* to 'remote'@'%' identified by 'qwe1234';
grant all privileges on `test`.* to 'remote'@'localhost' identified by 'qwe1234';
// 取消授权
revoke all privileges on `test`.* from 'remote'@'%' identified by 'qwe1234';
revoke all privileges on `test`.* from 'remote'@'locahost' identified by 'qwe1234';
// 刷新权限
flush privileges;
// 查看用户权限
show grants for 'remote'@'%';
show grants for 'remote'@'localhost';
```

## mysql5.6

```bash
select host, user, password, ssl_cipher, x509_issuer, x509_subject from user;
insert into user(host,user,password, ssl_cipher, x509_issuer, x509_subject) values("localhost","remote",password("qwe1234"), '', '', '');
```

## 修改密码

```bash
select user();
flush privileges;
set password=password('123456');
```

## 配置

```bash
// 查看MySQL配置文件
mysql find my.cnf
// 查看编码设置
show variables like "char%"
// 显示数据目录
show global variables like "%datadir%";
```

### 解决乱码

```sh
# 在my.cnf中设置编码格式
[client]
   default-character-set=utf8
[mysqld]
    character-set-server=utf8
    collation-server=utf8_general_ci
```

#### mysqld的默认编码是Latin1, 不支持中文

#### client默认跟随系统

## 时区

mysql默认的时区不是东八区, 需要手动设置, 执行完重新登录

```bash
select now();                   // 查看mysql系统时间。和当前时间做对比
set global time_zone = '+8:00'; // 设置时区更改为东八区
flush privileges;               // 刷新权限
```

## 导出数据

```sh
mysqldump -u用戶名 -p密码 -d 数据库名 表名 脚本名;

1、导出数据库为dbname的表结构（其中用戶名为root,密码为dbpasswd,生成的脚本名为db.sql）
mysqldump -uroot -pdbpasswd -d dbname >db.sql;

2、导出数据库为dbname某张表(test)结构
mysqldump -uroot -pdbpasswd -d dbname test>db.sql;

3、导出数据库为dbname所有表结构及表数据（不加-d）
mysqldump -uroot -pdbpasswd  dbname >db.sql;

4、导出数据库为dbname某张表(test)结构及表数据（不加-d）
mysqldump -uroot -pdbpasswd dbname test>db.sql;

```

## 额外


