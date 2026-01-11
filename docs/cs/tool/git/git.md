# 基本的git操作

## 安装及配置

```bash
sudo apt-get install git
git config --list
git config --global user.name "rrf"
git config --global user.email "xxx@qq.com"
git config --gobal core.autocrlf input
git config --global color.ui true
git init
ssh-keygen -t rsa -C "youremail@example.com"
```

## 添加提交

```bash
git add <file>
git rm <file>
git commit -m "desc"
```

## 状态

```bash
git status
git diff <file>
git diff HEAD -- readme.txt
```

## 回退

```bash
git log --pretty=oneline
git reflog
git reset --hard HEAD(commit_id)
git checkout -- file
git reset HEAD file
```

## 远程

```bash
git remote add origin address
git push -u origin master
git clone address
git remote
git remote -v
git checkout -b dev origin/dev
git branch --set-upstream dev origin/dev
```

## 分支

```bash
git branch
git branch -a
git branch master
git checkout -b dev
git checkout -b dev origin/dev
git stash
git stash list
git stash apply
git stash drop
git branch -d dev
git branch -D
```

## 标签

```bash
git tag
git tag <name>
git tag <name> <id>
git show <tagname>
git tag -a v0.1 -m "version 0.1 released" 3628164
git tag -s v0.2 -m "signed version 0.2 released" fec145a
git tag -d v0.1
git push origin :refs/tags/v0.9
git push origin <tagname>
git push origin --tags
git tag --sort version:refname
```

## ignore

```bash
git add -f App.class
```

## 别名

```bash
git config --global alias.st status
git config --global alias.co checkout
git config --global alias.ci commit
git config --global alias.br branch
git check-ignore -v App.class
git config --global alias.unstage 'reset HEAD'
git config --global alias.last 'log -1'
```

## 搭建git服务器

```bash
sudo apt-get install git
sudo adduser git
sudo git init --bare sample.git
sudo chown -R git:git sample.git
cat /etc/passwd | grep git 
git:x:1001:1001:,,,:/home/git:/bin/bash
git:x:1001:1001:,,,:/home/git:/usr/bin/git-shell
```

## 忽略某些已关联的文件

```bash
git update-index --assume-unchanged FILENAME
git update-index --no-assume-unchanged FILENAME
```

## 移除无效的远程分支信息

```bash
git prune
git push origin --delete <branchName>
git fetch -prune [origin]


git remote show origin
git remote prune origin
```

## 按日期查看tag

```bash
git for-each-ref --sort=taggerdate --format '%(refname) %(taggerdate)' refs/tags
git for-each-ref --sort=taggerdate --format '%(refname) %' refs/tags
git log --tags --simplify-by-decoration --pretty="format:%ci %d"
```

## 协议

git 支持多种协议访问, git, ssh, https

## 查看大文件

```bash
git rev-list --all --objects | \
grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -n 3 | awk -F ' '  '{print $1}')"
```

## 删除仓库种的文件


[参考](https://www.jianshu.com/p/d333ab0e6818)
```bash
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *.pcd' --prune-empty --tag-name-filter cat -- --all  


git filter-branch --tree-filter 'rm -rf path/folder' HEAD
git filter-branch --tree-filter 'rm -f path/file' HEAD
git for-each-ref --format='delete %(refname)' refs/original
git for-each-ref --format='delete %(refname)' refs/original | git update-ref --stdin
git reflog expire --expire=now --all
git gc --prune=now

git push origin master --force
git push origin --force --tags
```

### 增加源仓库

git remote add upstream  git@github.com:goproxyio/goproxy.git
git remote -v

git fetch upstream
git merge upstream/master

https://help.github.com/en/articles/configuring-a-remote-for-a-fork
https://help.github.com/en/articles/syncing-a-fork

### git 历史记录

```sh

# 从master创建develop
git checkout -b develop master

# 拉取远程分支
git fetch origin

git commit --verbose

git push origin --delete <branch-name>

git branch --delete <branch-name>

git commit --fixup master

# 提交本次修改的内容, 并附加到上一次的commit
git commit --amend

# 撤销某次的提交
git revert [Commit-ID]

# 强制回退
git reset --hard HEAD@{2}

# 使用某个commit
git cherry-pick [Commit-ID]

git reflog
```

## 远程仓库

```sh
git ls-remote -q https://github.com/x/sub/1
```

## git merge

```sh
# 直接合并，保留原有分支的commit, 默认使用ff
git merge dev

# 直接合并, 不使用ff, 并使用自定义说明本次merge
git merge --no-ff -m "merge with no-ff" dev

# 将要合并分支的多个commit合并到一次，--squash 会暂停commit提交，需要重新提交
git merge --squash f1
git commit -m 'desc'
```

## git rebase

rebase自己的分支, 不要操作共享分支

```sh
git rebase [basebranch] [topicbranch]

git reabse master

git rebase origin/master

# 交互式
git rebase -i master

#  自动保留
git rebase -i --autosquash master
```

## git pull

```sh
git pull

git pull origin == git fetch + git merge

# 拉取的时候使用rebase
git pull --rebase
```

## 支持url替换

```sh
git config --global url."git@xxx.com:".insteadOf "https://xxx.com"
```

## 特定主机的端口

```sh
# ~/.ssh/config
Host xxx.com
Port 80
```