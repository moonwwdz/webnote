## 这是一个记录笔记的应用

平时不爱记笔记，所以学东西慢，也没效率，虽然有blog,但每次写太长的东西，肚子里根本没东西，写不出来。     
而且太长的文章，确实是懒的写  
做这个应用就是单纯的想有个记录平时学习笔记的地方，当然也是为了学习golang

* [ ] 动态的，因为会有一个搜索，可以快速的搜索自己的笔记
* [ ] 多人的，虽然一定不会有别人用，但练手么
* [x] 尽可能的在单页面上完成
* [ ] ~~后期可能会加入react (脸好疼）~~

## 引用库
1.  [blackfriday](https://github.com/russross/blackfriday) *__MarkDown转换__*




## 安装使用
1. 自己编译
    
 ```bash
 #下载 
 $ git clone https://github.com/moonwwdz/webnote.git
 $ cd webnote
 #安装引用库
 $ go get github.com/astaxie/beego
 $ go get github.com/beego/bee
 $ go get github.com/russross/blackfriday
 #运行
 bee run

 ```
 *有其它问题。。。概不负责，自己搞定*

