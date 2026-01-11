# ntp

ntpd -n -q -p ntp.aliyun.com

## 极路由

```sh
/etc/config/system

# 修改list serve为阿里
config system
        option hostname 'Hiwifi'
        option timezone 'CST-8'
        option zonename 'Asia/Shanghai'

config timeserver 'ntp'
        list server 'ntp.aliyun.com'
        list server 'ntp.api.bz'
        list server 'time-a.nist.gov'
        option enabled '1'
        option enable_server '0'

config default 'dhcpc'
        option hostname 'Hiwifi'

config macoui 'maclist'
        list oui 'd4:ee:07'

```

### 定时同步

```sh

/etc/rc.common
/etc/crontabs/root

30 6 * * * ntpd -n -q -p ntp.aliyun.com
```
