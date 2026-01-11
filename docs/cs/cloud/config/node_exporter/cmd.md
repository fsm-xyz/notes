# NodeExporter

[NodeExporter](https://github.com/prometheus/node_exporter)

## 运行

```sh

podman run -d   --net="host" --name="node_exporter"   --pid="host"   -v "/:/host:ro,rslave"   quay.io/prometheus/node-exporter:latest   --path.rootfs=/host
```
