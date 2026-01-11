# rm

linux上删除文件都是rm，会直接删除，导致发生不可预期的后果，这里使用假删除，替换原来的rm

## 流程

```sh
# 1. 创建回收的文件目录
mkdir ~/.trash

# 2. 增加删除脚本.rm.sh

TRASH_DIR="$HOME/.trash"

for i in $*; do
    STAMP=`date "+%m-%d"`
    fileName=`basename $i`
    mv $i $TRASH_DIR/$fileName.$STAMP
done

# 3. 在.profile里面增加alisa, 替换原来的rm

alias rm="$HOME/.rm.sh"

# 4. 设置定时删除垃圾文件，crontab -e, 添加如下规则(每天的0点0分删除)

0 0 * * * rm -rf ~/.trash/*
```
