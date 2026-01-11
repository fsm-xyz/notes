# 优化镜像大小

想着如何减少镜像体积

## 思路

+ 使用更小的Base镜像
+ 合并RUN指令，减少Layer层数
+ 尽可能的清除中间过程的临时文件
+ 最小化可执行程序
+ multistage分段构建
+ .dockerignore
+ 移除不需要的文件

## 基础镜像

+ scratch 空的镜像，这是制作最小镜像的最佳选择
+ busybox 具有单一可执行文件的精简Unix工具集
+ alpine  包含apk和busybox

## 查询镜像

docker images
docker image inspect my_image:my_tag
docker image history my_image:my_tag

[Dive](https://github.com/wagoodman/dive)

## 参考

[slimming](https://towardsdatascience.com/slimming-down-your-docker-images-275f0ca9337e)
[瘦身](https://andyyoung01.github.io/2016/08/26/%E6%9E%84%E5%BB%BA%E5%B0%8F%E5%AE%B9%E9%87%8FDocker%E9%95%9C%E5%83%8F%E7%9A%84%E6%8A%80%E5%B7%A7/)
