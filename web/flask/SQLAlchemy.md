# SQLAlchemy

### 前言：

Django的ORM虽然强大，但是毕竟局限于Django，而SQLAlchemy是python中的ORM框架;

**SQLAlchemy的作用是：类/对象--->SQL语句--->通过pymysql/MySQLdb模块--->提交到数据库执行；**

组成部分：

* **Engine**:框架的引擎
* **Connection pooling**:数据库连接池
* **Dialect**:连接数据库的DB API种类
* **Schema/Types**：架构和类型
* **SQL Exprression Language**：SQL表达式语言

安装

`pip install sqlalchemy`

## 一、基本使用

### 1.原生SQL

```python
import time
import threding
import sqlalchemy
from sqlalchemy import create_engine
from sqlalchemy.engine.base import Engine



```



