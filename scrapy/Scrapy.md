# Scrapy

scrapy 是一个为了爬取网站数据，提取结构性数据而编写的应用框架。其可以应用在数据挖掘，信息处理和储存历史数据等一系列的程序中。

整体架构如下

![](D:\markdown\scrapy\scrapy.jpg)

Scrapy主要包括了以下组件：

- **引擎(Scrapy)**
  *用来处理整个系统的数据流处理, 触发事务(框架核心)*
- **调度器(Scheduler)**
  *用来接受引擎发过来的请求, 压入队列中, 并在引擎再次请求的时候返回. 可以想像成一个URL（抓取网页的网址或者说是链接）的优先队列, 由它来决定下一个要抓取的网址是什么, 同时去除重复的网址*
- **下载器(Downloader)**
  *用于下载网页内容, 并将网页内容返回给蜘蛛(Scrapy下载器是建立在twisted这个高效的异步模型上的)*
- **爬虫(Spiders)**
  *爬虫是主要干活的, 用于从特定的网页中提取自己需要的信息, 即所谓的实体(Item)。用户也可以从中提取出链接,让Scrapy继续抓取下一个页面*
- **项目管道(Pipeline)**
  *负责处理爬虫从网页中抽取的实体，主要的功能是持久化实体、验证实体的有效性、清除不需要的信息。当页面被爬虫解析后，将被发送到项目管道，并经过几个特定的次序处理数据。*
- **下载器中间件(Downloader Middlewares)**
  *位于Scrapy引擎和下载器之间的框架，主要是处理Scrapy引擎与下载器之间的请求及响应。*
- **爬虫中间件(Spider Middlewares)**
  *介于Scrapy引擎和爬虫之间的框架，主要工作是处理蜘蛛的响应输入和请求输出。*
- **调度中间件(Scheduler Middewares)**
  *介于Scrapy引擎和调度之间的中间件，从Scrapy引擎发送到调度的请求和响应。*

Scrapy运行流程大概如下：

1. 引擎从调度器中取出一个链接(URL)用于接下来的抓取
2. 引擎把URL封装成一个请求(Request)传给下载器
3. 下载器把资源下载下来，并封装成应答包(Response)
4. 爬虫解析Response
5. 解析出实体（Item）,则交给实体管道进行进一步的处理
6. 解析出的是链接（URL）,则把URL交给调度器等待抓取

**安装**

`pip install scrapy`

### 1.基本命令

`scrapy startproject projectname`创建一个爬虫项目

`scrapy genspider [-t template] <name> <domain>`创建一个爬虫应用

如：

​	`scrapy gensipder -t basic chouti chou.com`

​	`scrapy genspider -t xmlfedd autohome autohome.com.cn`

PS:

​	查看所有命令：`scrapy genspider -l`

`scrapy list`:查看爬虫应用列表

`scrapy crawl spidername`:运行单独爬虫应用

### 2.项目结构及爬虫应用简介

```
project_name``/
   ``scrapy.cfg
   ``project_name``/
       ``__init__.py
       ``items.py
       ``pipelines.py
       ``settings.py
       ``spiders``/
           ``__init__.py
           ``爬虫``1.py
           ``爬虫``2.py
           ``爬虫``3.py
```

文件说明：

* **scrapy.cfg ** 项目的主配置信息。（真正爬虫相关的配置信息在setting.py中）
* **item.py** 设置数据存储模板，用于结构化数据。
* **piplines.py** 数据处理行为。如：一般结构化的数据持久化
* **setting.py**配置文件，如爬虫深度、并发数、延迟下载等
* **spiders** 爬虫目录。

*注意：一般创建爬虫文件时，以网站域名命名*

### 3.简单实例

```python
import scrapy
from scrapy.select import HtmlXpathSelector
from scrapy.http.request import Request


class DigSpider(scrapy.Spider):
    #爬虫应用名称，通过此名称启动爬虫命令
    name = "dig"
    
    #
    
```

