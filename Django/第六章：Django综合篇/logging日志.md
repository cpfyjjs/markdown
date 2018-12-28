##  logging日志

### 一、基本流程
![img](http://static.zybuluo.com/feixuelove1009/g09y63fuzu9omngnn6k2rego/image.png)
### 二、基本参数配置
#### 1.Logger记录器
```python
logger = logging.getLogger(__name__)
```
**常用方法一览**
* logger.setLevel()
* logger.addHandler()
* logger.removeHandler()
* logger.addFilter()
* logger.removeFilter()

**创建对应日志记录**
* logger.debug()
* logger.info()
* logger.warning()
* logger.error()
* logger.critical()

#### 2.Handlers处理器
**内置Handler处理器**

* StreamHandler  
* FileHandler
* BaseRotatingHandler
* RotatingFileHandler
* TimedRotatingFileHandler
* SocketHandler
* DatagramHandler
* SMTPHandler
* SysLogHandler
* NTEventLogHandler
* HTTPHandler
* WatchedFileHandler
* QueueHandler
* NullHandler

### 三、logging高级用法事例
#### 1、创建loggers、hangers、formatters来使用logging
例子一
```python
# _*_coding:utf-8_*_

import logging

# 创建一个logging记录器
logger = logging.getLogger('simple_logger')

# 设置等级
logger.setLevel(logging.DEBUG)

# 创建一个控制台处理器，并将日志级别设置为debug
ch = logging.StreamHandler()
ch.setLevel(logging.DEBUG)

# 创建formatter格式化器
formatter = logging.Formatter('%(asctime)s - %(name)s -%(levelname)s - %(message)s')

# 将formatter添加到ch处理器
ch.setFormatter(formatter)

# 将ch添加到logger
logger.addHandler(ch)

# 将logging输出到文件
fh = logging.FileHandler('D:/markdown/Django/第六章：Django综合篇/debug.log',encoding='utf-8')
fh.setLevel(logging.DEBUG)
fh.setFormatter(formatter)
logger.addHandler(fh)


logger.debug('debug message')
logger.info('info message')
logger.warn('warn message')
logger.error('error message')
logger.critical('critical message')
```
### 2.创建一个字典然后将它传递给dictConfig()方法
例子一
```python
import logging.config

LOGGING = {
    'version':1,
    'disable_existing_loggers':False,
    'handlers':{
        'console':{
            'level':'DEBUG',
            'class':'logging.StreamHandler',
        },
        'file':{
            'level':'DEBUG',
            'class':'logging.FileHandler',
            'filename':'D:/markdown/Django/第六章：Django综合篇/debug.log',
            'encoding':'utf-8',
        },
    },
    'loggers':{
        'django':{
            'handlers':['console','file'],
            'level':'DEBUG',      
        },
    },
}

# 读取配置信息
logging.config.dictConfig(LOGGING)

# 获取logger记录器
logger = logging.getLogger('django') 
# 使用日志功能
logger.debug('debug message')
logger.info('info message')
logger.warn('warn message')
logger.error('error message')
logger.critical('critical message')
```
**一个复杂的logging配置**
```python
LOGGING = {
    'version': 1,
    'disable_existing_loggers': False,
    'formatters': {
        'verbose': {
            'format': '%(levelname)s %(asctime)s %(module)s %(process)d %(thread)d %(message)s'
        },
        'simple': {
            'format': '%(levelname)s %(message)s'
        },
    },
    'filters': {
        'special': {
            '()': 'project.logging.SpecialFilter',
            'foo': 'bar',
        },
        'require_debug_true': {
            '()': 'django.utils.log.RequireDebugTrue',
        },
    },
    'handlers': {
        'console': {
            'level': 'INFO',
            'filters': ['require_debug_true'],
            'class': 'logging.StreamHandler',
            'formatter': 'simple'
        },
        'mail_admins': {
            'level': 'ERROR',
            'class': 'django.utils.log.AdminEmailHandler',
            'filters': ['special']
        }
    },
    'loggers': {
        'django': {
            'handlers': ['console'],
            'propagate': True,
        },
        'django.request': {
            'handlers': ['mail_admins'],
            'level': 'ERROR',
            'propagate': False,
        },
        'myproject.custom': {
            'handlers': ['console', 'mail_admins'],
            'level': 'INFO',
            'filters': ['special']
        }
    }
}
```
**formatter格式**

属性 | 格式 | 描述
---|----|---
asctime | %(asctime)s | 日志产生的时间，默认格式为2003-07-08 16:49:45,896
created | %(created)f | time.time()生成的日志创建时间戳
filename | %(filename)s | 生成日志的程序名
funcName | %(funcName)s | 调用日志的函数名
levelname | %(levelname)s | 日志级别 ('DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL')
levelno | %(levelno)s | 日志级别对应的数值
lineno | %(lineno)d | 日志所针对的代码行号（如果可用的话）
module | %(module)s | 生成日志的模块名
msecs | %(msecs)d | 日志生成时间的毫秒部分
message | %(message)s | 具体的日志信息
name | %(name)s | 日志调用者
pathname | %(pathname)s | 生成日志的文件的完整路径
process | %(process)d | 生成日志的进程ID（如果可用）
processName | %(processName)s | 进程名（如果可用）
thread | %(thread)d | 生成日志的线程ID（如果可用）
threadName | %(threadName)s | 线程名（如果可用）