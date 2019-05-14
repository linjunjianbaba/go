git相关

​     1、在本地创建一个版本库（即文件夹），通过git init把它变成Git仓库；

​     2、把项目复制到这个文件夹里面，再通过git add .把项目添加到仓库；

​     3、再通过git commit -m "注释内容"把项目提交到仓库；

​     4、在Github上设置好SSH密钥后，新建一个远程仓库，通过git remote add origin https://github.com/guyibang/TEST2.git将本地仓库和远程仓库进行关联；

​     5、最后通过git push -u origin master把本地仓库的项目推送到远程仓库（也就是Github）上；（若新建远程仓库的时候自动创建了README文件会报错，解决办法看上面）。

## Git 常用命令及使用

![img](https://timg01.bdimg.com/timg?pacompress&imgtype=1&sec=1439619614&autorotate=1&di=7fb97c104ea32ba9b27641fc1d6e6228&quality=90&size=b310_10055&cut_x=72&cut_y=0&cut_w=310&cut_h=310&src=http%3A%2F%2Ftimg01.bdimg.com%2Ftimg%3Fpacompress%26imgtype%3D1%26sec%3D1439619614%26autorotate%3D1%26di%3D3398dba125f5bc97513d59207a507397%26quality%3D90%26size%3Db870_10000%26src%3Dhttp%253A%252F%252Fpic.rmb.bdstatic.com%252F1525324741638a8a2937a4fe4b7d97ee5c232cca78.jpeg)

Git 常用命令使用

1）、本地库初始化 git init

2）、设置签名

作用：区分不同开发人员的身份。

说明：这里设置的签名和登录远程库（代码托管中心）的账户没有关系。

a)、项目级别签名:

git config user.name [AAA]

git config user.email [邮箱地址]

签名信息位置：**cat .git/config**

b)、系统级别签名:

git config **--global**user.name [AAA]

git config **--global** user.email [邮箱地址]

签名信息位置：cd ~ 、**cat .gitconfig**

3）、基本操作

a)、查看状态： **git status**(查看工作区、暂存区的状态)

b)、添加操作: **git add** **文件名**(将工作区新建/修改的内容添加到暂存区)

c)、提交操作： **git commit -m “commit message”** **文件名**(将暂存区的内容提交到本地库)

4）、查看历史记录

a)、git log

b)、git log --pretty=oneline

c)、git log --oneline

d)、git reflog (HEAD@{移动到当前版本需要多少步})

5）、前进和后退

a)、基于索引值的操作（推荐做法）

**git reset --hard** **哈希索引值**

示例：找回删除状态已经提交本地库的文件操作。

b)、使用^符号 （只能后退，一个^表示后退一步）

**git reset --hard HEAD^**

c)、使用~符号 （只能后退，n表示后退n步）

**git reset --hard HEAD~2**

6）、比较文件差异

a)、git diff [文件名] (将工作区中的文件和暂存区的进行比较)

b）、git diff [本地库历史版本] [文件名] (将工作区中的文件和本地库历史记录比较，不带文件名的话，会比较多个文件)

7）、分支管理

在版本控制过程中，使用多条线同时推进多个任务。



![img](https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=485545563,2438022825&fm=173&app=49&f=JPEG?w=640&h=336&s=119AAB7F1759546C52F175C20100E031)





分支的优势？

a)、同时并行推进多个功能开发，提高开发效率。

b)、各个分支在开发过程中，如果某个分支开发失败，不会对其他分支有影响，失败的分支可以删除，然后重新开始即可。

分支常用命令：

a)、git branch -v （查看本地库中的所有分支）

b)、git branch dev (创建一个新的分支)

c)、git checkout dev （切换分支）

d)、分支合并

i)、切换到接收修改的分支

**git checkout master**

ii)、执行merge命令

**git merge dev**

（注：切换分支后，在dev分支中做出的修改需要合并到被合并的分支master上)

**8****）、冲突解决**

当一个分支的内容和另一个分支的内容不同时，此时任一分支合并另一分支过程中就会出现冲突。



![img](https://ss0.baidu.com/6ONWsjip0QIZ8tyhnq/it/u=3840930747,1750568043&fm=173&app=49&f=JPEG?w=640&h=278&s=19C8AB5D165948680EBD7D6003007070)



冲突的解决办法：

a)、编辑文件，删除特殊符号。

b）、将文件修改完毕后，保存退出。

c)、git add [文件名]。

d)、git commit –m “日志信息”。

注意：此时commit时不能带文件名。