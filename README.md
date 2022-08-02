# gostu
go study recomand

* 生成ssh秘钥
```
ssh-keygen -t rsa -C "xxx"
```
* 存放秘钥目录
```
/c/Users/chihl/.ssh/id_rsa.pub
```
* git提交中文乱码问题
```
* 设置git 的界面编码：
git config --global gui.encoding utf-8

* 设置 commit log 提交时使用 utf-8 编码：
git config --global i18n.commitencoding utf-8

* 在 $ git log 时将 utf-8 编码转换成 gbk 编码：
git config --global i18n.logoutputencoding gbk

* git log 正常显示中文：
export LESSCHARSET=utf-8
```