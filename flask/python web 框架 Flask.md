# python web 框架 Flask

Flask 是使用python 编写的一款轻量级web框架。其WSGI 工具箱采用的werkzeug，模板引擎是Jinjia2。同时，具有极好的可扩展性。http://flask.pocoo.org/extensions/

### flask 简单使用

Flask 的安装` pip install flask`    

#### 实例一  

```python
from flask import Flask

app = Flask(__name__)	# 创建一个falsk 实例

@app.route("/index")	# 通过装饰器，将url 与视图函数相绑定
def index():
    return "hello world"

if __name__ == "__main__":
    app.run()	# 启动flask,监听浏览器发过来的请求
```



## 一、配置文件

`app = Flask(__name__,tempalte_floder="templates",static_url_path="/static/")`

模板路径：` template_folder='templates'`

静态文件路径：`static_url_path='/static/'`

静态文件引入别名：`static_path='/zhanggen'`

设置为调试环境：`app.debug=True` （代码修改自动更新）

设置json编码格式 如果为False 就不使用ascii编码：`app.config['JSON_AS_ASCII']=False `

设置响应头信息`Content-Type   app.config['JSONIFY_MIMETYPE'] ="application/json;charset=utf-8"  `（注意 ;charset=utf-8）



## 二、路由系统

### 1.动态路由

#### 1.1接受字符串参数

```python
from flask import Flask

app = Flask(__name__)

@app.route('/<name>')
def index(name):
	return "hello %s"%name
	
if __name__ == "__main__":
    app.run
```



#### 1.2接受整型参数

```python
from flask import Flask

app = Flask(__name__)

@app.route("/students/<int:pid>")
def student(pid):
    return "hello %d"%pid

if __name__ == "__main__":
    app.run()
```



#### 1.3接受浮点型

```python
from flask import Flask

app = Flask(__name__)

@app.route("/post/<float:salary>")
def sale(salary):
    return "...."

if __name__ == "__main__":
    app.run
```



#### 1.4 接收URL链接类型参数

```python
from flask import Flask
app=Flask(__name__)
@app.route('/<path:url>/')  #设置url传参数：http://127.0.0.1:5000/http://www.baiu.com/
def url_flask(url):  #视图必须有对应接收参数
    print(url)
    return 'Hello World'  #response

if __name__ == '__main__':
    app.run()
```



### 2.制定允许的请求方法

```python
from flask import Flask
app=Flask(__name__)
@app.route('/inex',methods=['GET',"POST"]) #只允许get、post请求
def method_flask(url):
    print(url)
    return 'Hello World'  

if __name__ == '__main__':
    app.run()
```



### 3.通过别名反向生成url

```python
from flask import Flask,url_for

app = Flask(__name__)

@app.route("/students/<int:pid>",endpoint = "name1")
def student(pid):
    print(url_for("name1"，pid = pid))
    return "..."

if __name__ == "__main__":
    app.run()
```



### 4.通过app.add_url_rule()调用路由

```python
from flask import Flask,url_for

app = Flask(__name__)

def student(pid):
    print(url_for("name1"，pid = pid))
    return "..."

app.add_url_rule(rule='/index/',endpoint='name1',view_func=first_flask,methods=['GET'])
#app.add_url_rule(rule=访问的url,endpoint=路由别名,view_func=视图名称,methods=[允许访问的方法])


if __name__ == "__main__":
    app.run()
```



### 5.扩展路由功能：正则匹配url

```python
from flask import Flask,views,url_for 
from werkzeug.routing import BaseConverter

app = Flask(__name__)

class RegexConverter(BaseConverter):
    """
    自定义URL匹配正则表达式
    """
    
    def __init__(self,map,regex):		# map参数不需要传入，自动传入
        super(RegexConverter,self).__init__(map)
        self.regex = regex
        
    def to_python(self,value):
        """
        路由匹配时，匹配成功后传递给视图参数的值
        """
        return int(value)
    
    def to_url(self,value):
        """
        使用url_for 反向生成url时，传递的参数经过该方法处理，返回的值用于生成url中的参数
        """
		val = super(RegexConverter, self).to_url(value)
        return val
    
app.url_map.converters['regex'] = RegexConverter

@app.route("/index/<regex('\d+'):nid>")
def index(nid):
    print(url_for("index",nid ='8888'))
    return "..."

if __name__ == "__main__":
    app.run
```



## 三、视图

### 1.给视图函数加装饰器

如果要给视图函数加装饰器，一定要加载在路由装饰器的下面。

```python
from flask import Flask
from functiontools import wapper

def auth(func):
    @wapper(func)
    def inner(*args,**kwargs):
        print("权限认证")
        result = func(*args,**kwargs)
        print("。。。")
        return result
    
app = Flask(__name__)

@app.route("/index")	# 通过装饰器，将url 与视图函数相绑定
@auth
def index():
    return "hello world"

if __name__ == "__main__":
    app.run()	# 启动flask,监听浏览器发过来的请求
```



### 2.request 和 resopnse

`request.method`:获取请求的方法

`request.json`

`request.args`:获取get的请求参数

`request.from`:获取post的请求参数

`request.form.get("name")`:获取单个值

`request.from.getlist("name_list")`:获取参数列表。（多个值）

`request.values.get("age")`:获取GET和POST请求携带的所有参数（GET/POST通用）

