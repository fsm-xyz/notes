# kubernetes

## 定义

容器编排和部署的一套解决方案

Container Manager
Container Moniter
Container (Scheduler, Load Balance, Scalability)

## 核心

>* POD      一组容器
>* RC       Replication Controller Monitor & Control
>* Service  外部访问,LB

## command

```sh
kubectl cluster-info
kubectl get cs
kubectl get nodes
kubectl get pods
kubectl get rc
kubectl get replicasets
kubectl get svc
kubectl get namespace
kubectl describe pod nginx
kubectl get pods --all-namespaces

kubectl create -f filename
kubectl replace -f filename
kubectl patch pod nginx
kubectl edit pod nginx
kubectl delete -f filename
kubectl apply -f filename
kubectl logs nginx
kubectl rolling-update
kubectl scale
kubectl autoscale
kubectl attach
kubectl exec
kubectl run

kubectl get pods -n kube-system
watch kubectl get pods --all-namespaces
```

## option

-o wide

## 例子

```sh
kubectl run nginx3 --image=nginx --port=80 --replicas=10 --labels="run=lb"

kubectl expose deployment nginx3  --type="LoadBalancer" --name="my-nginx"

kubectl run nginx --image nginx --port 8080
kubectl expose rc nginx --type=LoadBalancer
kubectl scala rc nginx --replicas=3

watch -n 1 `kubectl get services`
watch -n 1 `kubectl get pod`
```
