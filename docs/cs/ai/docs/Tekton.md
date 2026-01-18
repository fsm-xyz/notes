# Tekton

## 安装

```sh
# 添加资源文件到集群
kubectl apply --filename \
https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

# 

```

## 资源类型

+ Step
+ Task
+ Pipeline
+ TaskRun
+ PipelineRun
+ EventListener

## 安全

后续实现

+ cosign

## Token

有两种方式，当前直接使用k8s的Secret,还有HashiCrop Vault（使用sidecar或者init容器获取）

### k8s Secret

+ 当作环境变量注入
+ 当作文件注入

```sh

kubectl create secret generic my-thirdparty-token \
  --namespace=your-tekton-namespace \
  --from-literal=token=YOUR_ACTUAL_TOKEN_HERE


# 环境变量

apiVersion: tekton.dev/v1
kind: Task
metadata:
  name: call-external-api
spec:
  steps:
    - name: invoke-api
      image: appropriate/curl-image # 例如curl的镜像
      env:
        - name: API_TOKEN
          valueFrom:
            secretKeyRef:
              name: my-thirdparty-token # 你创建的Secret名称
              key: token # Secret中存储Token的键名
      command: ['/bin/sh']
      args: ['-c', 'curl -H "Authorization: Bearer $API_TOKEN" https://api.example.com/resource']

# 文件注入

apiVersion: tekton.dev/v1
kind: Task
metadata:
  name: call-external-api-with-file
spec:
  steps:
    - name: invoke-api
      image: appropriate/curl-image
      volumeMounts:
        - name: token-volume
          mountPath: /etc/secrets
          readOnly: true
      command: ['/bin/sh']
      args: ['-c', 'curl -H "Authorization: Bearer $(cat /etc/secrets/token)" https://api.example.com/resource']
  volumes:
    - name: token-volume
      secret:
        secretName: my-thirdparty-token # 你创建的Secret名称
        items:
          - key: token # Secret中存储Token的键名
            path: token # 挂载后的文件名
```

#### 安全

使用单独的ServiceAccount进行权限控制，最小权限原则

#### todo

+ 安全
+ 集成外部Secret管理系统
+ 定期轮换Token
+ 加密静态数据