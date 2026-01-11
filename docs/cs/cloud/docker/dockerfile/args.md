# dockerfile 格式

FROM NGINX                                          基础镜像依赖
COPY index.html /usr/share/nginx/html/index.html    复制文件到容器
CMD                                                 执行的命令
VOLUME                                              把文件挂载到容器
EXPOSE                                              暴露端口
WORKDIR                                             工作目录
RUN                                                 构建执行的命令
ENV                                                 设置时区
MAINTAINER                                          维护信息

## 生成imgage

```shell
docker build -t nginx15
```

## CMD ENTRYPOINT

The CMD instruction has three forms:

### exec

CMD ["executable","param1","param2"] (exec form, this is the preferred form)

CMD ["param1","param2"] (as default parameters to ENTRYPOINT)

ENTRYPOINT ["executable","param1","param2"] (exec form, this is the preferred form)

不会解析环境变量

### shell

CMD command param1 param2 (shell form)

shell会解析环境变量

底层调用 /bin/sh -c [command]

## Note

shell会解析环境变量

CMD会被docker run后面的参数替换
