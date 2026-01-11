# postgres简单入门

## 安装(源码或者二进制)

    官方都有对应的文档指导

## 用法

    默认会创建一个unix postgres用户, postgres超级数据库用户和数据库

```sh
su postgres
psql
create role/user name
drop role/user name
```

shell command

```sh
createuser name
dropuser name
createdb dbname
dropdb dbname
```

## 角色授权

```sh
SELECT rolname FROM pg_roles;
SELECT datname FROM pg_database;
CREATE DATABASE name OWNER renruifeng;
GRANT ALL PRIVILEGES ON database test TO renruifeng;
alter role renruifeng with LOGIN;
```

## 连接

1
```sh
psql
-U role
-d Database
```

## 远程访问

listen_addresses = '*'

## 自动服务启动

```sh
    systemctl enable postgresql-10
    systemctl start postgresql-10
    systemctl stop postgresql-10
    systemctl restart postgresql-10
```
