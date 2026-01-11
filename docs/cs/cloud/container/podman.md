

## 用户隔离


user1：
podman run --rm -it node
podman run --rm -it node

root：
podman run --rm -it node

各自不客见，属于运行者自己的权限范围内，不像docker允许在docker用户权限下



```sh
# 系统进程信息
UID          PID    PPID  CMD

fsm          619       1  /usr/lib/systemd/systemd --user
fsm          773     619   \_ /usr/bin/plasmashell --no-respawn
fsm         1036     773   |   \_ /usr/bin/konsole
fsm         1449    1036   |   |   \_ /bin/zsh
fsm       130398    1449   |   |   |   \_ podman

fsm       130713    1036   |   |   \_ /bin/zsh
fsm       130757  130713   |   |   |   \_ podman

fsm       131147    1036   |   |   \_ /bin/zsh
root      131194  131147   |   |   |   \_ su
root      131196  131194   |   |   |       \_ bash
root      131219  131196   |   |   |           \_ podman

fsm       130604     619   \_ /usr/bin/conmon 
fsm       130606  130604   |   \_ node
fsm       130779     619   \_ /usr/bin/conmon
fsm       130781  130779   |   \_ node
root      131591     619   \_ /usr/bin/conmon
root      131593  131591   |   \_ node
```
  