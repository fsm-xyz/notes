# Fedora安装

## 移除旧版本

```sh
dnf remove docker \
    docker-client \
    docker-client-latest \
    docker-common \
    docker-latest \
    docker-latest-logrotate \
    docker-logrotate \
    docker-selinux \
    docker-engine-selinux \
    docker-engine
```

## 安装使用的仓库

```sh
dnf -y install dnf-plugins-core
dnf config-manager \
    --add-repo \
    https://download.docker.com/linux/fedora/docker-ce.repo
```

## 安装

### 直接

```sh
dnf install docker-ce docker-ce-cli containerd.io
```

### 查询特定版本并安装

```sh
dnf list docker-ce  --showduplicates | sort -r
dnf -y install docker-ce-<VERSION_STRING> docker-ce-cli-<VERSION_STRING> containerd.io
```

## 启动

```sh
systemctl start docker
```

## 验证

```sh
docker run hello-world
```

## 自启动

```sh
systemctl enable docker
```
