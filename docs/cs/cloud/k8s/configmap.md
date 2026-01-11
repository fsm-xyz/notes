# ConfigMap

## 创建

```sh
# 直接基于文件 xxx.yml为ConfigMap资源定义
kubectl apply -f xxx.yml

# 从文件中读取，key为可选项
kubectl create configmap --from-file=[key]=source

# 基于目录，把目录下的文件名当作key，内容为value
kubectl create configmap --from-file=dir

# 基于字面量
kubectl create configmap --from-literal=key=value
```
## 读取

挂载为环境变量
挂载为容器内文件

## faq

使用path/to解决挂载时候覆盖原来的目录
