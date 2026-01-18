# c++

## 静态检测

这个过程只需要研发提供那些目录需要静态检测

### 工具

+ Clang-Tidy
+ Cppcheck

### 集成代码质量平台

+ SonarQube
+ Codecov

### 质量门禁和报告


### 优化点

增量分析（只检查修改的文件）

## 自动化测试

单元测试、集成测试、系统测试，Mock 对象测试，性能测试

业务方提供具体的实现，比如在Makefile里面提供各种功能

Google Test (gtest)

Catch2

Boost.Test

CppUnit

## 构建产物推送到机器人

+ 基于 SCP/SFTP + SSH 的传统推送方案
+ 基于 Ansible/SaltStack 的配置管理工具方案
+ 基于容器（Docker）和仓库的现代方案
+ 基于 CI/CD 平台 Agent 的主动拉取方案
+ 通过消息队列或 Webhook 的异步触发方案

其实就是2种，主动推送升级和主动拉取升级

## C++工程主要在公司离线环境处理，产物管理和测试可以放公有云

## 硬件测试

+ QEMU模拟
+ 公司内部的测试设备
