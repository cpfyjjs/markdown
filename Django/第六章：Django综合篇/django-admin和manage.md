## django-admin和manage.py
### 一、Django内置命令选项
#### 1.check
检查整个Django项目是否存在常见问题。默认情况下，所有应用都被选中。也可以指定app的名字检查指定应用。

`django-admin check app01 app02 app03`
#### 2.deffsetings
查看当前设置文件与Django的默认设置之间的差异

`django-admin deffsetings`
#### 3.flush
只删除具体数据，不删除数据表

`django-admin flush`

#### 4.makemigrations
根据检测道德模型创建新的迁移。迁移的作用，更多的是将数据库的操作，以文件的形式记录下来，以便以后检查、调用、重做等等。尤其是对于Git版本管理，它无法获知数据库是如何变化的，只能通过迁移文件中的记录来追溯和保存

`python mamange.py makemigrations [app01 app02 app03]`
#### 5.migrate
是数据库状态与当前模型集和迁移集同步。

`python manage.py migrate [app01 app02 app03]`

#### 6.runserver
#### 7.shell 
启动带有Django环境的Python交互式解释器，

`django-admin shell -i ipthon`

`djang0-admin shell -i bpython`

#### 8.startapp

### 二、app提供的命令
django-admin command [option]
python manage.py command [option]
* changepassword
* createsuperuser
* clearsessions
  

#### traceback
当引发CommandError时，显示完整的错误栈信息。默认情况下，django-admin将显示一个简单的错误消息。

用法示例：

django-admin migrate --traceback
