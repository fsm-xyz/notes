# 解压缩工具

## xz

### 压缩

```bash
tar -cvf xxx.tar xxx
xz -z xxx.tar
```

-k 保留xxx.tar

## 解压

```bash
xz -d xxx.tar.xz
tar -xvf xxx.tar
```

## gzip

```bash
# 压缩
tar -czvf xxx.tar.gz xxx
# 解压
tar -xzvf xxx.tar.gz
```

## zip

```bash
# 压缩
zip xxx.zip xxx
# 解压
upzip xxx.zip
```

## 下载文件

> wget
> scp
> nc
> sz,rz
