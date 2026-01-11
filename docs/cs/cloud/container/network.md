# network

## slirp4netns


默认rootless的容器网络模式

普通用户当设置了port mapping后，用bridge，其他为slirp4netns

rootful用户设置了--privileged， 用bidge， 其他slirp4netns


https://docs.podman.io/en/latest/markdown/podman-run.1.html

### FAQ

自动加入容器网络的dns服务，一些在docker下运行正常，但是这里会报DNS错误`dial tcp: lookup goproxy.cn on 10.0.2.3:53: no such host`

cat /etc/hosts
cat /etc/resolv.conf

```sh
# Run a command in a modified user namespace
podman unshare --rootless-netns