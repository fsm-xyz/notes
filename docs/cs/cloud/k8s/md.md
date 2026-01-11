### kubctl

```sh
kubectl -n monitoring get po -o wide

kubectl get crd

kubectl -n monitoring get prometheus

kubectl -n monitoring get all

kubectl  patch svc  grafana -n monitoring -p '{"spec":{"type":"NodePort","ports":[{"name":"http","port":3000,"protocol":"TCP","targetPort":"http","nodePort":30030}]}}'

```

### prometheus

```sh

# 查询prometheus的匹配label

kubectl get prometheus  -o yaml | grep -A4 serviceMonitorSelector

/targets
/service-discovery
/metrics

```

prometheus-prometheus.yaml

#### 自定义存储

```yaml
spec:
  storage:
    volumeClaimTemplate:
      spec:
        storageClassName: nfs-client
        resources:
          requests:
            storage: 40Gi
```

### grafana

https://grafana.com/dashboards/315
https://grafana.com/dashboards/8919