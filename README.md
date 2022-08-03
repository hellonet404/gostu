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
显示 status 编码
git config --global core.quotepath false

图形界面编码
git config --global gui.encoding utf-8

提交信息编码
git config --global i18n.commit.encoding utf-8

git log 默认使用 less 分页，所以需要 bash 对 less 命令进行 utf-8 编码
export LESSCHARSET=utf-8 (windows下为：set LESSCHARSET=utf-8)
```