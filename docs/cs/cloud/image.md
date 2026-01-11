# image

减小容器镜像体积

alpine, distroless, scratch

## 基础镜像

### distroless

基于debian的，体积最小

gcr.io/distroless/static
gcr.io/distroless/base

gcr.io/distroless/base:debug

static, base

### alpine

基于busybox, 使用musl替换glibc

## 多级构建

## 资源复用

### 挂载重复使用的库

docker in docker