# ip link

## TUN/TAP

```sh
ip addr add 192.168.1.111/24 dev tun0
ip link set tun0 up
```

## veth

```sh
ip link add veth0 type veth peer name veth0
ip addr add 192.168.1.111/24 dev veth0
ip addr del 192.168.3.111/24 dev veth0
ip link set veth0 up
ip link set dev veth1 down
```

## bridge

```sh
ip link add name br0 type bridge
ip link set br0 up
```

## 连接

```sh
ip link set dev veth0 master br0
ip link set br0 up
```

## route

```sh
route -n
route
```

## brctl

```sh
brctl show
brctl addbr br0
brctl delbr br0
brctl addif br0 eth0
brctl delif br0 eth0
ifconfig eth0 0.0.0.0
```

## nmcli

```sh
sudo nmcli connection show
sudo nmcli connection delete xxx
```

[nmcli](https://ywnz.com/linuxjc/4595.html)
