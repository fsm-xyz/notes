# 构建golang应用镜像

## 方案大致如下

### 基础镜像直接添加

直接添加编译好的二进制到其他基础镜像
缺点: 需要手动编译二进制，添加到基础镜像，过程不连续

### 基于golang的镜像上面构建

添加项目到镜像，编译
优点: 可以debug, 方便后续调试
缺点: 开发环境可以, 生产环境的话镜像体积太大

### 分段构建

  在golang基础镜像上构建，然后放到其他镜像

GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o app app.go && tar c app | docker import - app:latest
