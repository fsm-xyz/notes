# Jaeger

## 组件

[组件列表](https://www.jaegertracing.io/download/)

## 后端存储

[jaeger](https://www.jaegertracing.io/docs/1.19/faq/)

## 安装

[docs](https://www.jaegertracing.io/docs/1.19/)
[operator](https://www.jaegertracing.io/docs/1.19/operator/)

### emptyDir

不使用外部的PVC资源

```sh
cat <<EOF | kubectl apply -f -
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: quickstart
spec:
  version: 7.9.1
  nodeSets:
  - name: default
    count: 3
    config:
      node.master: true
      node.data: true
      node.ingest: true
      node.store.allow_mmap: false
    podTemplate:
      spec:
        volumes:
        - name: elasticsearch-data
          emptyDir: {}
EOF
```

### 验证

```sh
PASSWORD=$(kubectl get secret quickstart-es-elastic-user -o go-template='{{.data.elastic | base64decode}}')

# inside
curl -u "elastic:M70U7ADV3f14ueL6eRa7Z31e" -k "https://quickstart-es-http:9200"

# local
kubectl port-forward service/quickstart-es-http 9200

```


kubectl -n observability create secret generic jaeger-secret --from-literal=ES_PASSWORD=M70U7ADV3f14ueL6eRa7Z31e --from-literal=ES_USERNAME=elastic
### jaeger

```sh
apiVersion: jaegertracing.io/v1
kind: Jaeger
metadata:
  name: simple-prod
spec:
  strategy: production
  storage:
    type: elasticsearch
    options:
      es:
        server-urls: https://quickstart-es-http.db.svc:9200
        index-prefix: my-prefix
        tls:
          ca: /es/certificates/ca.crt
    secretName: jaeger-secret
  volumeMounts:
    - name: certificates
      mountPath: /es/certificates/
      readOnly: true
  volumes:
    - name: certificates
      secret:
        secretName: db/quickstart-es-http-certs-public
```
