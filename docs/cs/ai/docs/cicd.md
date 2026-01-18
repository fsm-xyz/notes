# CICD

尽量跨云，不依赖某个厂商，可以随便复制

## 需求方

c++这边需要同时 apollo镜像和orin镜像的自动构建托管 Charlie Ren
流水线跑代码的测试用例，还有静态检查之类的
可以，不仅仅是提供环境，可以跑完用例和检查，根据结果决定是否允许代码合入

## 工具选型


+ TekTon
+ Argo CD
+ Gocd
+ flux
+ openkruise
+ ByteBase
+ devtron
+ flagger
+ harness

## IaS

devops平台目前主要的有GUI工具和IaS两种，IaS是声明式的，GUI类似命令式的

GUI在当前的情况下不适合(需要开发的工作多，周期长)。

IaS可以根据GitOps进行部署，只要搭建好整体流程就可以了

### 基本的方案

1. 使用Tekton进行CI构建，kusemize进行k8s资源文件yaml生成，提交到devops的git仓库
2. Argo CD进行集群部署

先Run起来后面再用更高级的别的工具