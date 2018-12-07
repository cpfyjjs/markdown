# python web 框架 Flask

Flask 是使用python 编写的一款轻量级web框架。其WSGI 工具箱采用的werkzeug，模板引擎是Jinjia2。同时，具有极好的可扩展性。http://flask.pocoo.org/extensions/

### flask 简单使用

Flask 的安装` pip install flask`    

#### 实例一  

```python
from flask import Flask

app = Flask(__name__)	# 创建一个falsk 实例

@app.router("/index")	# 通过装饰器，将url 与视图函数相绑定
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



