# ss

## 搭建shadowsocks服务器

       1. [官网](https://shadowsocks.org/en/index.html)        
       2. [server环境](https://shadowsocks.org/en/download/servers.html)
       3. 配置文件
        cat /etc/shadowsocks/config.json
        Shadowsocks accepts JSON format configs like this:
        {
            "server":"my_server_ip",
            "server_port":8388,
            "local_port":1080,
            "password":"barfoo!",
            "timeout":600,
            "method":"chacha20-ietf-poly1305"
        }

       4. 启动
        ssserver -c /etc/shadowsocks/config.json              // 先这样启动，能看到报错信息
        ssserver -c /etc/shadowsocks/config.json -d start

       5. 优化
        1.网站有自己的优化方案
        2.开启BBR加速  kernel 4.9以上支持bbr，需要手动开启 ([BBR介绍](https://vircloud.net/linux/start-bbr.html))
       6. 常见问题
           - undefined symbol EVP_CIPHER_CTX_cleanup
        ```bash
            vim /usr/local/lib/xxx/dist-packages/shadowsocks/crypto/openssl.py
            :%s/cleanup/reset/
            :x
        ```
     xxx为自己的py版本，在openssl1.1.0版本中，废弃了EVP_CIPHER_CTX_cleanup函数
           * socket.error: [Errno 99] Cannot assign requested
            修改配置文件的server为0.0.0.0

## client

命令行设置翻墙
pip install shadowsocks

```json /etc/shadowsocks.json
{
    "server":"91xgp.com",
    "server_port":"8000",
    "local_address":"127.0.0.1",
    "local_port":1087,
    "password":"rain12345",
    "timeout":600,
    "method":"aes-256-cfb"
}
```

```shell
yum install privoxy
/etc/privoxy/config
forward-socks5   /               127.0.0.1:1087 .
listen-address  localhost:8118
privoxy /etc/privoxy/config
```

```sh
export http_proxy=127.0.0.1:8118
export https_proxy=127.0.0.1:8118

export http_proxy=192.168.19.49:1087
export https_proxy=192.168.19.49:1087

sslocal -c /etc/shadowsocks.json -d start

export https_proxy=http://localhost:8118
```

FAQ

局域网连接. 设置http的ip为0.0.0.0

## Chrome没法使用代理

Windows宽带连接使用了中文, 使用英文名字即可
