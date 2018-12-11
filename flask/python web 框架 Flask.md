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

#### a.请求相关信息

`request.method`:获取请求的方法

`request.json`

`request.args`:获取get的请求参数

`request.from`:获取post的请求参数

`request.form.get("name")`:获取单个值

`request.from.getlist("name_list")`:获取参数列表。（多个值）

`request.values.get("age")`:获取GET和POST请求携带的所有参数（GET/POST通用）

`request.cookies.get("name")`:获取cookies信息

`request.headers.get("Host")`:获取请求头相关的信息

`request.path`:获取用户访问的url地址，例如（/，/login,/index)

`request.full_path`:获取用户访问的完整路径+参数，例如（/login/?age=18&name=boby)

`request.host`:获取主机地址

`request.files`:获取用户上传的文件

```python
obj = request.files['the_file_name']

obj.save('/var/www/uploads/' + secure_filename(f.filename))  # 直接保存
```

#### b.响应相关信息

```python
return "字符串"
return render_template("xx.html",**kwargs)
return redirect("/index")
```

相应json数据

```python
# 方式一
return jsonify(user_list)
# 配置方式
app.config['JSON_AS_ASCII']=False  #指定json编码格式 如果为False 就不使用ascii编码，
app.config['JSONIFY_MIMETYPE'] ="application/json;charset=utf-8" #指定浏览器渲染的文件类型，和解码格式；
```

```python
# 方式二
return Response(data,mimetype="application/json;charset=utf-8",)
```

```python
# 设置响应头
from flask import Flask,make_response,render_template

response = make_resopnse(render_template("index.html"))
response.set_cookie("key","value")
response.delete_cookie("key")
response.headers["X-something"] = "A value"

return reponse
```

#### C.Flask CBV 视图

```python
# CBV视图
from functools import wraps
from flask import Flask,url_for,views

app = Flask()

def auth(func):
    @wraps(func)
    def inner(*args,**kwargs):
        # 权限认证
        result = func(*args,**kwargs)
        return result
    return inner

# CBV shitu
class IndexView(views.MethodView):
    method = ["GET"]
    decorators = [auth,]
    
    def get(self):
        return "..."
    
    def post(self):
        return "???"

app.add_url_rule("/index",view_func=IndexView.as_view(name="name1"))

if __name__ == "__main__":
    app.run()

```



## 四、模板语言

Flask 使用的是Jinja2模板语言，所以其语法和Django无差别(Django的模板语言参考Jinja2)

### 1.引入静态文件

方式一、别名引入

```html
<linlk rel="stylesheet" href="/zhanggen/commons.css"
```

方式二、url_for()方式引入

```html
<link rel="stylesheet" href="{{ url_for('stactic',filename='commons.css')}}"
```

### 2.模板语言引用上下文对象

变量

```html
<span>{{name}}</span>
```

循环、索引取值

```html
<ul>
    {% for user in user_list %}
    <li>{{user.name}}</li>
    {% endfor}
</ul>
{{user_list.0}}
{{user_list[0]}}

```

Flask的Jinjia2可以通过Context把视图中的函数传递到模板语言中执行，simple_tag 、 simple_filter

simle_tag(只能传两个参数，支持for ,if)

```python
from flask import Flask

app = Flask(__name__)

@app.template_global()
def foo(arg):
    return "<input type='text'>"
```

```html
<h1>
    {{foo(1)|safe}}
</h1>
```

simple_filter(对参数无限制，不支持for、if)

```python
from flask import Flask

app = Flask(__name__)
@app.template_filter()
def foo(arg1,arg2,arg3):
    return arg1+arg2+arg3

```

```html
<h1>
   {{"boy"|foo("s","b")}} 
</h1>
```

### 3.wtforms(flask表单验证插件)

#### 3.0简介

Wtforms 是一个支持多个web框架的form组件，主要对用户请求数据进行表单验证

#### 3.1安装

`pip install wtforms`

#### 3.2简单使用

wtforms和Django自带额form验证插件功能相同，使用起来大同小异哦。

用户登陆验证

```python
#_*_ coding:utf-8 _*_

from flask import Flask,render_template,request,redirect
from wtforms import Form
from wtforms import core
from wtforms import htmls5
from wtforms import simple
from wtforms import validators
from wtforms import widgets

app = Flask(__name__,template_folder="templates")
app.debug = True

# 定义用于登陆验证的form类
class LoginForm(Form):
    
    name = simple.StringField(
    lable = '用户名',
        validators=[
            validators.DataRequire(message="用户名不能为空"),
            validators.Length(min=6,max=18,message="用户名长度必须大于%(min)d且小于%(max)d")
        ],
    widget=widgets.TextIput()		#前端页面显示的插件TextArea
	render_kw={'class': 'form-control'}      #设置form标签的class信息
    )
    pwd = simple.PasswordField(
        label='密码',
        validators=[
            validators.DataRequired(message='密码不能为空.'),
            validators.Length(min=8, message='用户名长度必须大于%(min)d'),
            validators.Regexp(regex="^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&])[A-Za-z\d$@$!%*?&]{8,}",message='密码至少8个字符，至少1个大写字母，1个小写字母，1个数字和1个特殊字符')],
        widget=widgets.PasswordInput(),
        render_kw={'class': 'form-control'}
    )
    
@app.route("/login",methods=["GET","POST"])
def login():
    if request.method = "GET":
        form = LoginForm()	# 实例化LoginForm
		return render_template("login.html",form = form)
    else:
        from = LoginForm(request.form)
        if form.validate():		#判断是否验证成功
            print("用户提交的数据的格式验证成功")
            # 进一步取数据库数据，查看用户名密码是否正确
            
        else:
            print("用户提交的数据的格式验证失败")
            print(form.errors)
        return render_template("login.html",form = form)
    
if __name__ == "__main__":
    app.run()

```

```html	
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>登录</h1>
<form method="post" novalidate>
    <!--<input type="text" name="name">-->
    <p>{{form.name.label}} {{form.name}} {{form.name.errors[0] }}</p>

    <!--<input type="password" name="pwd">-->
    <p>{{form.pwd.label}} {{form.pwd}} {{form.pwd.errors[0] }}</p>
    <input type="submit" value="提交">
</form>
</body>
</html>
```

用户注册页面

```python
#用户注册
from flask import Flask, render_template, request, redirect
from wtforms import Form
from wtforms.fields import core
from wtforms.fields import html5
from wtforms.fields import simple
from wtforms import validators
from wtforms import widgets

app = Flask(__name__, template_folder='templates')
app.debug = True



class RegisterForm(Form):
    name = simple.StringField(
        label='用户名',
        validators=[
            validators.DataRequired()
        ],
        widget=widgets.TextInput(),
        render_kw={'class': 'form-control'},
        default='张根'                                             #设置input标签中默认值
    )

    pwd = simple.PasswordField(
        label='密码',
        validators=[
            validators.DataRequired(message='密码不能为空.')
        ],
        widget=widgets.PasswordInput(),
        render_kw={'class': 'form-control'}
    )

    pwd_confirm = simple.PasswordField(                                #第二次输入密码
        label='重复密码',
        validators=[
            validators.DataRequired(message='重复密码不能为空.'),
            validators.EqualTo('pwd', message="两次密码输入不一致")  #验证2次输入的密码是否一致？
        ],
        widget=widgets.PasswordInput(),
        render_kw={'class': 'form-control'}
    )

    email = html5.EmailField(
        label='邮箱',
        validators=[
            validators.DataRequired(message='邮箱不能为空.'),
            validators.Email(message='邮箱格式错误')
        ],
        widget=widgets.TextInput(input_type='email'),    #生成email input标签
        render_kw={'class': 'form-control'}
    )

    gender = core.RadioField(
        label='性别',
        choices=(                                        #choice radio选项
            (1, '男'),
            (2, '女'),
        ),
        coerce=int                                       #讲用户提交过来的 '4' 强制转成 int 4
    )
    city = core.SelectField(
        label='城市',
        choices=(
            ('bj', '北京'),
            ('sh', '上海'),
        )
    )

    hobby = core.SelectMultipleField(                      #select 下拉框多选框
        label='爱好',
        choices=(
            (1, '篮球'),
            (2, '足球'),
        ),
        coerce=int
    )

    favor = core.SelectMultipleField(
        label='喜好',
        choices=(
            (1, '篮球'),
            (2, '足球'),
        ),
        widget=widgets.ListWidget(prefix_label=False),        #生成Checkbox 多选框
        option_widget=widgets.CheckboxInput(),
        coerce=int,
        default=[1, 2]
    )

    def __init__(self, *args, **kwargs):                        #重写form验证类的__init__方法可以实时同步数据中数据
        super(RegisterForm, self).__init__(*args, **kwargs)
        self.favor.choices = ((1, '篮球'), (2, '足球'), (3, '羽毛球'))


    def validate_pwd_confirm(self, field):                       #wtforms验证 钩子函数
        """
        自定义pwd_confirm字段规则，例：与pwd字段是否一致
        :param field:
        :return:
        """
        # 最开始初始化时，self.data中已经有所有的值

        if field.data != self.data['pwd']:
            # raise validators.ValidationError("密码不一致") # 继续后续验证
            raise validators.StopValidation("密码不一致")  # 不再继续后续验证


@app.route('/register/', methods=['GET', 'POST'])
def register():
    if request.method == 'GET':
        form = RegisterForm(data={'gender': 1})  #默认值
        return render_template('register.html', form=form)
    else:
        form = RegisterForm(formdata=request.form)
        if form.validate():
            print('用户提交数据通过格式验证，提交的值为：', form.data)
        else:
            print(form.errors)
        return render_template('register.html', form=form)



if __name__ == '__main__':
    app.run()
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h1>用户注册</h1>
<form method="post" novalidate style="padding:0  50px">
    {% for item in form %}
    <p>{{item.label}}: {{item}} {{item.errors[0] }}</p>
    {% endfor %}
    <input type="submit" value="提交">
</form>
</body>
</html>
```

## 五、session功能

### 1.Flask自带的session功能

```python

from  flask import Flask,session,request
import json

app = Flask(__name__,template_floder="templates",static_path="/static/",static_url_path="/static/")

app.debug = True
app.secret_key = "sdfadfasdf"		#设置session加密
app.config["JSON_AS_ASCII"] = False		#指定json编码格式，如果为False，就不能使用ASCII编码
app.config["JSON_MIMETYPE"] = "appliction/json;charset=utf-8"	#指定浏览器渲染的文件格式及编码格式

@app.route("/login",method=["GET","POST"])
def login():
    msg = ""
    if request.method == "POST":
        name = request.values.get("user")
        password = request.values.get("password")
        
        if name and password:
            session["user"]=name
            return redirect("/index")
        else:
            msg = "用户名或者密码错误"
            
    return render_template("login.html",msg = msg)

@app.route("/index")
def index():
    if session.get("user"):
        return render_template("index.html")
    else:
        return redirect("/login")
    
if __name__ == "__main__":
    app.run()
```

### 2.第三方session组件

安装 pip install flask_session

配置文件，将session 存在redis中

```python
from flask import session,Flask,request
from flask import make_resopnse,render_templte,redirect,jsonify,Response
from flask.ext.session import Session	# 引入第三方session
import json
from redis import Redis

app = Flask(__name__,template_floder="templates",static_path="/static/",static_url_path="/static/")

app.debug = True
app.secret_key = "dsfaskljlk;"	# 设置session加密
app.config["JSON_AS_ASCII"] = False	# 指定json编码格式，如果为Flase,就不能使用ascii编码
app.config['JSONIFY_MIMETYPE'] ="application/json;charset=utf-8" 

app.config["SESSION_TYPE"]='redis'

from redis import Redis        #引入连接 redis模块
app.config['SESSION_REDIS']=Redis(host='192.168.0.94',port=6379) #连接redis
Session(app)

@app.route("/login",method=["GET","POST"])
def login():
    msg = ""
    if request.method == "POST":
        name = request.values.get("user")
        password = request.values.get("password")
        
        if name and password:
            session["user"]=name
            return redirect("/index")
        else:
            msg = "用户名或者密码错误"
            
    return render_template("login.html",msg = msg)

@app.route("/index")
def index():
    if session.get("user"):
        return render_template("index.html")
    else:
        return redirect("/login")
    
if __name__ == "__main__":
    app.run()
```

### 3.自定义session组件

```python
#!/usr/bin/env python
# -*- coding:utf-8 -*-
import uuid
import json
from flask.sessions import SessionInterface
from flask.sessions import SessionMixin
from itsdangerous import Signer, BadSignature, want_bytes


class MySession(dict, SessionMixin):
    def __init__(self, initial=None, sid=None):
        self.sid = sid
        self.initial = initial
        super(MySession, self).__init__(initial or ())

    def __setitem__(self, key, value):
        super(MySession, self).__setitem__(key, value)

    def __getitem__(self, item):
        return super(MySession, self).__getitem__(item)

    def __delitem__(self, key):
        super(MySession, self).__delitem__(key)


class MySessionInterface(SessionInterface):
    session_class = MySession
    container = {}

    def __init__(self):
        import redis
        self.redis = redis.Redis()

    def _generate_sid(self):
        return str(uuid.uuid4())

    def _get_signer(self, app):
        if not app.secret_key:
            return None
        return Signer(app.secret_key, salt='flask-session',
                      key_derivation='hmac')

    def open_session(self, app, request):
        """
        程序刚启动时执行，需要返回一个session对象
        """
        sid = request.cookies.get(app.session_cookie_name)
        if not sid:
            sid = self._generate_sid()
            return self.session_class(sid=sid)

        signer = self._get_signer(app)
        try:
            sid_as_bytes = signer.unsign(sid)
            sid = sid_as_bytes.decode()
        except BadSignature:
            sid = self._generate_sid()
            return self.session_class(sid=sid)

        # session保存在redis中
        # val = self.redis.get(sid)
        # session保存在内存中
        val = self.container.get(sid)

        if val is not None:
            try:
                data = json.loads(val)
                return self.session_class(data, sid=sid)
            except:
                return self.session_class(sid=sid)
        return self.session_class(sid=sid)

    def save_session(self, app, session, response):
        """
        程序结束前执行，可以保存session中所有的值
        如：
            保存到resit
            写入到用户cookie
        """
        domain = self.get_cookie_domain(app)
        path = self.get_cookie_path(app)
        httponly = self.get_cookie_httponly(app)
        secure = self.get_cookie_secure(app)
        expires = self.get_expiration_time(app, session)

        val = json.dumps(dict(session))

        # session保存在redis中
        # self.redis.setex(name=session.sid, value=val, time=app.permanent_session_lifetime)
        # session保存在内存中
        self.container.setdefault(session.sid, val)

        session_id = self._get_signer(app).sign(want_bytes(session.sid))

        response.set_cookie(app.session_cookie_name, session_id,
                            expires=expires, httponly=httponly,
                            domain=domain, path=path, secure=secure)
```

```python
from flask import Flask
from flask import session
from my_session import MySessionInterface

app = Flask(__name__)

app.secret_key = 'A0Zr98j/3yX R~XHH!jmN]LWX/,?RT'
app.session_interface = MySessionInterface()


@app.route('/login/', methods=['GET', "POST"])
def login():
    print(session)
    session['user1'] = 'alex'
    session['user2'] = 'alex'
    del session['user2']

    return "内容"


if __name__ == '__main__':
    app.run()
```

## 六、蓝图

使用Flask自带的Blueprint模块，可以帮助我们进行目录结构的划分

```python
# 关于用户账户的蓝图

```







