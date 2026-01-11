# github

## 无权限

使用git协议下载github上面的库提示如下:
`Permission denied (publickey).
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.`

没有配置当前的机器公钥到github

## github仓库与源库代码保持一致

与源库建立关系

git remote add upstream git@github.com
git remote -v
git remote update upstream
<!-- 直接合并源库的变更 -->
git pull upstream master
<!-- 基于rebase -->
git rebase upstream/master

### rebase发生冲突

解决冲突文件
git add 冲突文件, 无需commit, git rebase --continue自动应用补丁

取消则执行git rebase --abort

## git reset, git revert, git rebase, git cherry-pick
