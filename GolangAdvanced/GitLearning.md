# Git Learning

## 命令掌握

1. git branch & git checkout

这个是切出自己的分支，企业开发不能直接在master/main上写代码

- git branch feature : 创建一个叫做feature的分支
- git checkout feature : 切换到feature分支

一般用这个：git checkout -b feature : 创建并且自动切换到feature分支

2. git status & git diff

这个是对当前分支的文件和代码进行检查：status告诉修改了哪些文件，diff告诉修改了哪几行代码；

3. git stash

这个情况是：在开发功能A的时候，突然有个紧急的需求B需要开发；此时先 git stash 暂存当前的修改，然后去开发需求B；提交完了之后，再来 git stash pop, 恢复之前开发需求 A 的状态；

4. git merge & git rebase

- git merge : 把 B 分支的代码合并到 A 分支，会产生一个新的 "Merge commit" 节点；
- git rebase : 这个的意思是，将自己的提交记录，剪切下来，粘贴到最新的公共代码后面；

## 关于 Git Flow

Git Flow 不是一个命令，而是一种约定的分支管理规范；

标准的 Git Flow 模型包括以下分支：

- Master(Main): 生产环境分支
- Develop(Dev): 开发主干分支
- Feature: 功能分支，用来从Dev中切出来，然后开发完了返回给Dev
- Release: 预发布分支，从Dev中切出来，用来测试和修复bug
