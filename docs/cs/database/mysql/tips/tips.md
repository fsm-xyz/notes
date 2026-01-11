# 技巧

## 

```sh
# 查看当前连接的信息
\s

# 查看字符编码
show variables like '%char%';
```

### 无法输入特殊字符

1. 本地新建一个sql文件
2. 输入sql语句
3. 保存
4. 登录mysql
5. source sql文件