### 1.自定义管理器

在语句`Book.objects.all()`中`objects`是一个特俗的属性，通过它来查询数据库，它就是模型的一个`Manager`。

每个**Django**模型至少有一个`Manager`，你可以创建自定义`Manager`以定制数据库的访问。

#### 添加额外的Manager

添加额外的manager是为模块添加**表级功能**的首选方法。（至于**行级功能**，也就是只作用于模型实例对象的函数，则通过自定义模型方法实现）。

```python
# models.py

#_*_coding:utf-8_*_

from django.db import models

class BookManager(models.Manager):

    def title_count(self,keyworld):
        return self.filter(title_icoutains=keyworld).count()


class BookModel(models.Model):
    title = models.CharField(max_length=100)
    pub_date = models.DateField()
    press = models.ForeignKey(to = 'press',on_delete=models.CASCADE)
    description = models.CharField(255)
    authors = models.ManyToManyField(to='Author')

    # 添加额外的管理器
    objects = BookManager()

    def __str__(self):
        return self.title

# 自定义的查询方法
BookModel.objects.title_count('演绎')
# 默认的查询方法依然可用
BookModel.objects.filter(title__icontains="演绎").count()
```

#### 修改初始Manager Queryset

manager的基础Queryset返回系统中的所有对象.例如,`Book.objects.all()`返回book数据库中的所有书籍.你而已通过覆盖`Manager.get_queryset()`方法来重写manager的基础Queryset.`get_queryset()`应该按照你的需求返回一个Queryset.

```python
# models.py

#_*_coding:utf-8_*_

from django.db import models

class BookManager(models.Manager):

    def title_count(self,keyworld):
        return self.filter(title_icoutains=keyworld).count()

    def get_querset(self):
        return super(BookManager.self).get_queryset().filter(press="上海出版社")


class BookModel(models.Model):
    title = models.CharField(max_length=100)
    pub_date = models.DateField()
    press = models.ForeignKey(to = 'press',on_delete=models.CASCADE)
    description = models.CharField(255)
    authors = models.ManyToManyField(to='Author')
	
    objects = models.Manager()
    # 添加额外的管理器
    shanghai = BookManager()

    def __str__(self):
        return self.title

# 自定义的查询方法
BookModel.shanghai.all()
# 默认的查询方法依然可用
BookModel.objects.all()
```

自定义的Manager对象，请注意，Django遇到的第一个Manager(以它在模型中被定义的位置为准)会有一个特殊状态。 Django将会把第一个Manager 定义为默认Manager ，Django的许多部分(但是不包括admin应用)将会明确地为模型使用这个manager。 结论是，你应该小心地选择你的默认manager。因为覆盖`get_queryset()`了，你可能接受到一个无用的返回对像，你必须避免这种情况.

### 2.自定义模型的方法

模型方法只对特殊模型实例起作用

```python
# models.py

#_*_coding:utf-8_*_

from django.db import models

class Person(models.Model):
    first_name = models.CharField(max_length=30)
    last_name = models.CharField(max_length=30)
    birth_date = models.DateField()

    def body_status(self):
        # 返回年龄状态
        from datetime import datetime,timedelta
        now = datetime.now()
        if self.birth_date > (now - timedelta(days=60*365)):
            return "老年"
        elif self.birth_date > (now - timedelta(days=18 * 365)):
            return "成年"
        else:
            return "未成年"
    
    def _age_status(self):
        from datetime import datetime, timedelta
        now = datetime.now()
        if self.birth_date > (now - timedelta(days=60 * 365)):
            return "老年"
        elif self.birth_date > (now - timedelta(days=18 * 365)):
            return "成年"
        else:
            return "未成年"
        
    age_status = property(_age_status)      # 将方法包装成属性
    
p = Person.objects.get(first_name='lilei')
p.birth_date
p.body_status()
p.age_status
```

### 3.重写预定义方法

通过重写模型预定的`save(),delete(),clean()`,来完成一些特殊的操作。

比如重写`save()`方法，可以通过日志记录一些操作。