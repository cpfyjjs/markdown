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

