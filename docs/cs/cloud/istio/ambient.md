

## 证书

```sh


istioctl pc secret ds/ztunnel -n istio-system -o json | jq -r '.dynamicActiveSecrets[0].secret.tlsCertificate.certificateChain.inlineBytes' | base64 --decode | openssl x509 -noout -text -in /dev/stdin


istioctl pc secret ztunnel-tcn4m -n istio-system


WORKER1=$(kubectl -n istio-system get pods --field-selector spec.nodeName==ambient-worker -lapp=ztunnel -o custom-columns=:.metadata.name --no-headers)
WORKER2=$(kubectl -n istio-system get pods --field-selector spec.nodeName==ambient-worker2 -lapp=ztunnel -o custom-columns=:.metadata.name --no-headers)

kubectl -n istio-system port-forward $WORKER1 15000:15000&
kubectl -n istio-system port-forward $WORKER2 15001:15000&

curl -XPOST "localhost:15000/logging?level=debug"
curl -XPOST "localhost:15001/logging?level=debug"

kubectl -n istio-system logs -lapp=ztunnel -f
# Or,
kubectl -n istio-system logs $WORKER1 -f
kubectl -n istio-system logs $WORKER2 -f

curl "localhost:15000/config_dump"

```

## 批量

```sh

kubectl exec -it deploy/sleep -- sh -c 'for i in $(seq 1 100); do curl -s http://productpage:9080 | grep reviews; done'

```

## tcpdump

```sh
tcpdump -nAi  lo port 8080 or port 8443
```

## kube debug

```sh

kubectl debug -it -n istio-system ztunnel-9c8hc --image=nicolaka/netshoot

termshark
```
