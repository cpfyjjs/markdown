# Scrapy

Scrapy 是一个为了爬取网站数据，提取结构性数据而编写的应用框架。其可以应用在数据挖掘，信息处理和储存历史数据等一系列的程序中。

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

官方文档数据流：

1. 引擎打开一个网站(open a domain)，找到处理该网站的Spider并向该spider请求第一个要爬取的URL(s)。
2. 引擎从Spider中获取到第一个要爬取的URL并在调度器(Scheduler)以Request调度。
3. 引擎向调度器请求下一个要爬取的URL。
4. 调度器返回下一个要爬取的URL给引擎，引擎将URL通过下载中间件(请求(request)方向)转发给下载器(Downloader)。
5. 一旦页面下载完毕，下载器生成一个该页面的Response，并将其通过下载中间件(返回(response)方向)发送给引擎。
6. 引擎从下载器中接收到Response并通过Spider中间件(输入方向)发送给Spider处理。
7. Spider处理Response并返回爬取到的Item及(跟进的)新的Request给引擎。
8. 引擎将(Spider返回的)爬取到的Item给Item Pipeline，将(Spider返回的)Request给调度器。
9. (从第二步)重复直到调度器中没有更多地request，引擎关闭该网站。

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
project_name/
   scrapy.cfg
   project_name/
       __init__.py
       items.py
       pipelines.py
       settings.py
       spiders/
           __init__.py
           爬虫1.py
           爬虫2.py
           爬虫3.py
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
    
    #允许的域名,即只允许爬去此域名下的网页，
    allowed_domians = ['chouti.com']
    
    #起始URL
    start_urls=['http://chouti.com']
    
    has_request_set ={}
    
    #解析响应
    def parse(self,response):  
        page_list=response.xpath("//div[@id="dig_lcpage"]//a[re:test(@href, "/all/hot/recent/\d+")]/@href").extrat()
        
        for page in page_list:
            page_url = "http://dig.chouti.com%s"%page
			key = self.md5(page_url)
            if key in self.has_request_set:
                pass
            else:
                self.has_request_set[key] = page_url
                obj = Request(url=page_url, method='GET', callback=self.parse)
                yield obj
                
    @staticmethod
    def md5(val):
        import hashlib
        has = hashlib.md5()
        ha.update(val,encoding="utf-8")
        key = ha.hexdigest()
        return key
         
```

启动项目

`scrapy crawl dig`

### 4.xpath选择器

```python
#!/usr/bin/env python
# -*- coding:utf-8 -*-
from scrapy.selector import Selector, HtmlXPathSelector
from scrapy.http import HtmlResponse
html = """<!DOCTYPE html>
<html>
    <head lang="en">
        <meta charset="UTF-8">
        <title></title>
    </head>
    <body>
        <ul>
            <li class="item-"><a id='i1' href="link.html">first item</a></li>
            <li class="item-0"><a id='i2' href="llink.html">first item</a></li>
            <li class="item-1"><a href="llink2.html">second item<span>vv</span></a></li>
        </ul>
        <div><a href="llink2.html">second item</a></div>
    </body>
</html>
"""
response = HtmlResponse(url='http://example.com', body=html,encoding='utf-8')
hxs = HtmlXPathSelector(response)
# 选取在子孙中的a标签
hxs = Selector(response=response).xpath('//a')
# 选取在子孙中的a标签的第二个
hxs = Selector(response=response).xpath('//a[2]')
# 选取在子孙中带有id属性的a标签
hxs = Selector(response=response).xpath('//a[@id]')
# 选取在子孙中id属性等于i1的a标签
hxs = Selector(response=response).xpath('//a[@id="i1"]')
# 选取在子孙中id属性等于i1、href属性等于link.html的a标签
hxs = Selector(response=response).xpath('//a[@href="link.html"][@id="i1"]')
# 选取在子孙中href属性包含link的a标签
hxs = Selector(response=response).xpath('//a[contains(@href, "link")]')
# 选取在子孙中href属性以link开头的a标签
hxs = Selector(response=response).xpath('//a[starts-with(@href, "link")]')
# 利用正则表达式
hxs = Selector(response=response).xpath('//a[re:test(@id, "i\d+")]')
# 提取文字
hxs = Selector(response=response).xpath('//a[re:test(@id, "i\d+")]/text()').extract()
# print(hxs)
hxs = Selector(response=response).xpath('//a[re:test(@id, "i\d+")]/@href').extract()
# 提取href属性
hxs = Selector(response=response).xpath('/html/body/ul/li/a/@href').extract()
# print(hxs)
hxs = Selector(response=response).xpath('//body/ul/li/a/@href').extract_first()
```

#### 登陆抽屉网,并点赞(cookies)

```python
#_*_ conding:utf-8 *_*
import scrapy
from scrapy.http.request import Request
from scrapy.http.cookies import CookieJar
from scrapy import FormRequest

class ChouTiSpider(scrapy.Spider):
    # 爬虫名称
    name = 'chouti'
    #允许域名
    allow_domians = ['chouti.com']
    
    cookie_dict={}
    has_request_set ={}
    
    def start_request(self,response):
        """起始方法"""
        url = 'http://dig/chouti.com'
        # return [Request(url = url,callback = self.login)]
        yield Request(url = url ,callback = self.login)
        
    def login(self,response):
        """登陆"""
        cookie_jar = CookieJar()
        # 提取cookies
        cookie_jar.extract_cookies(response,response.request)
        # 将提取的cookies 放入cookie_dict中
        for k,v in cookie_jar._cookies.items():
            for i,j in v.items():
                for m,n in j.items():
                    self.cookie_dict[m]=n.value
        req = Request(
            url = 'http://dig/chouti.com/login',
            method = "POST",
            headers = {'Content-Type':'appliction/x-www-form-urlencoded; charset=UTF-8'},
			body='phone=8615131255089&password=pppppppp&oneMonth=1',
            cookies = self.cookie_dict,
            callback = self.check_login)
        yield req
        
    def check_login(self,response):
        """返回首页"""
        req = Request(
        url = 'http://dig/chouti.com/',
        method = "GET",
        cookies = self.cookie_dict,
        dont_filter = True,
        callback = self.show)
        yield req
        
     def show(self, response):
        # print(response)
        hxs = HtmlXPathSelector(response)
        news_list = hxs.select('//div[@id="content-list"]/div[@class="item"]')
        for new in news_list:
            # temp = new.xpath('div/div[@class="part2"]/@share-linkid').extract()
            link_id = new.xpath('*/div[@class="part2"]/@share-linkid').extract_first()
            yield Request(
                url='http://dig.chouti.com/link/vote?linksId=%s' %(link_id,),
                method='POST',
                cookies=self.cookie_dict,
                callback=self.do_favor
            )

        page_list = hxs.select('//div[@id="dig_lcpage"]//a[re:test(@href, "/all/hot/recent/\d+")]/@href').extract()
        
        for page in page_list:
            page_url = 'http://dig.chouti.com%s' % page
            import hashlib
            hash = hashlib.md5()
            hash.update(bytes(page_url,encoding='utf-8'))
            key = hash.hexdigest()
            if key in self.has_request_set:
                pass
            else:
                self.has_request_set[key] = page_url
                yield Request(
                    url=page_url,
                    method='GET',
                    callback=self.show
                )

    def do_favor(self, response):
        print(response.text)
        
```

#### 处理cookies

```python
import scrapy
from scrapy.http.response.html import HtmlResponse
from scrapy.http import Request
from scrapy.http.cookies import CookieJar


class ChoutiSpider(scrapy.Spider):
    name = "chouti"
    allowed_domains = ["chouti.com"]
    start_urls = (
        'http://www.chouti.com/',
    )

    def start_requests(self):
        url = 'http://dig.chouti.com/'
        yield Request(url=url, callback=self.login, meta={'cookiejar': True})

    def login(self, response):
        print(response.headers.getlist('Set-Cookie'))
        req = Request(
            url='http://dig.chouti.com/login',
            method='POST',
            headers={'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'},
            body='phone=8613121758648&password=woshiniba&oneMonth=1',
            callback=self.check_login,
            meta={'cookiejar': True}
        )
        yield req

    def check_login(self, response):
        print(response.text)
```

### 5.格式化数据

```python
import scrapy

# 定义结构化数据
class XiaoHuaItem(scrapy.Item):
    name = scrapy.Field()
    school = scrapy.Field()
    url = scrapy.Field()
    
```

```python
import json
import os
import requests


class JsonPipeline(object):
    def __init__(self):
        self.file = open('xiaohua.txt', 'w')

    def process_item(self, item, spider):
        v = json.dumps(dict(item), ensure_ascii=False)
        self.file.write(v)
        self.file.write('\n')
        self.file.flush()
        return item


class FilePipeline(object):
    def __init__(self):
        if not os.path.exists('imgs'):
            os.makedirs('imgs')

    def process_item(self, item, spider):
        response = requests.get(item['url'], stream=True)
        file_name = '%s_%s.jpg' % (item['name'], item['school'])
        with open(os.path.join('imgs', file_name), mode='wb') as f:
            f.write(response.content)
        return item
```

```python
# setting.py
ITEM_PIPELINES = {
   'spider1.pipelines.JsonPipeline': 100,
   'spider1.pipelines.FilePipeline': 300,
}
# 每行后面的整型值，确定了他们运行的顺序，item按数字从低到高的顺序，通过pipeline，通常将这些数字定义在0-1000范围内。
```

#### pipeline

```python
#_*_ coding:UTF-8 _*_

class CustomPipeline(object):
    """自定义Pipeline"""
    
    def __init__(self,v):
        self.value = v
        
    @classmethod
    def from_crawler(cls,crawler):
        """
        初始化时候，用于创建Pipeline对象
        """
        val = crawler.settings.getint("MMM")
        retrun cls(val)
        
    def open_spider(self,spider):
        """
        爬虫开始执行的时候调用
        """
        # 可以用于链接数据库等初始化操作
        print("处理数据之前")
        
    def process_item(self，item,spider):
        """
        操作并进行持久化
        """
        # 表示会被后续的pipeline继续处理
        return item
    
    	# 表示将item丢弃，不会被后续的pipeline处理
        # raise DropItem()
        
    def close_spider(self,spider):
        """
        爬虫关闭时，被调用
        """
        print("处理数据之后")
        
```

### 6.中间件

#### 爬虫中间件

爬虫中间件是介入到Scrapy的spider处理机制的钩子框架，您可以添加代码来处理发送给 [Spiders](https://scrapy-chs.readthedocs.io/zh_CN/latest/topics/spiders.html#topics-spiders)的response及spider产生的item和request。

```python
class SpiderMiddleware(object):

    def process_spider_input(self,response, spider):
        """
        下载完成，执行，然后交给parse处理 
        """
        pass

    def process_spider_output(self,response, result, spider):
        """
        spider处理完成，返回时调用
        :return: 必须返回包含 Request 或 Item 对象的可迭代对象(iterable)
        """
        return result

    def process_spider_exception(self,response, exception, spider):
        """
        异常调用
        :return: None,继续交给后续中间件处理异常；含 Response 或 Item 的可迭代对象(iterable)，交给调度器或pipeline
        """
        return None


    def process_start_requests(self,start_requests, spider):
        """
        爬虫启动时调用
        :return: 包含 Request 对象的可迭代对象
        """
        return start_requests
```

#### 下载中间件

```python
# 定义下载中间件
class DownMiddleware(object):
    
    def process_request(self,request,spider):
        """
        下载器下载前执行
        retrun :
        	None:继续后续中间件去下载
        	Response对象：停止后续的process_request的执行，开始执行process_response
        	Request对象：停止下载中间件的执行，将request对象返回给调度器
        	raise IgnoreRequest:停止后去process_request的执行，开始执行processs_exception
        """
        pass
    
    def process_response(self,request,response,spider):
        """
        下载器下载完成后执行
        return：
        	Response对象：转交给后续的中间件执行process_response
        	Request对象:停止后续下载中间件的执行，将request对象返回给调度器
        	raise IgnoreRequest异常，停止process_request的执行，开始执行process_exception
        """
        return response
    
    def process_exception(self,request,response,spider):
        """
        当下载处理器(download handler)或 process_request() (下载中间件)抛出异常时执行。
        :return: 
            None：继续交给后续中间件处理异常；
            Response对象：停止后续process_exception方法
            Request对象：停止中间件，request将会被重新调用下载
        """
        return None
       
```

### 7.自定制命令

* 在spiders同级目录创建任意目录，如commands
* 在其中创建`rawlall.py`文件（此处文件名就是自定义的命令）

```python
from scrapy.commands import ScrapyCommand
from scrapy.untils.project import get_project_settings

class Command(ScrapyCommand):
    
    require_project = True
    
    def syntax(self):
        return '[options]'
    
    def short_desc(self):
        return 'Runs all of the spiders'
    
    def run(self,args,opts):
        spider_list = self.crawler_process.spiders.list()
        for name in spider_list:
            self.crawler_process.crawl(name,**opts.__dict__)
        self.crawler_process.start()
              
```

- 在settings.py 中添加配置 COMMANDS_MODULE = '项目名称.目录名称'
- 在项目目录执行命令：`scrapy crawlall `

**在编译器中运行scrapy**

```python
import sys
from scrapy.cmdline import execute

if __name__ == '__main__':
    execute(["scrapy","github","--nolog"])
```



### 8.自定义扩展

自定义扩展时，利用信号在制定位置注册制定操作

```python
#_*_ coding:UTF-8 _*_

from scrapy import signals


class MyExtension(object):
    
    def __init__(self,value):
        self.value = value
        
    @classmethod
    def from_crawler(cls, crawler):
        val = crawler.settings.getint('MMMM')
        ext = cls(val)
        # 信号注册
        crawler.signals.connect(ext.spider_opened, signal=signals.spider_opened)
        crawler.signals.connect(ext.spider_closed, signal=signals.spider_closed)
        return ext

    def spider_opened(self, spider):
        print('open')

    def spider_closed(self, spider):
        print('close')
```

### 9.避免重复访问

Scrapy默认使用`scrapy.dupefilter.RFDupefliter`进行去重，相关配置

```python
DUPEFILTER_CLASS = 'scrapy.dupefilter.RFPDupeFilter'
DUPEFILTER_DEBUG = False
JOBDIR = "保存范文记录的日志路径，如：/root/"  # 最终路径为 /root/requests.seen
```

**自定义去重**

```python
class RepeatUrl:
    def __init__(self):
        self.visited_url = set()

    @classmethod
    def from_settings(cls, settings):
        """
        初始化时，调用
        :param settings: 
        :return: 
        """
        return cls()

    def request_seen(self, request):
        """
        检测当前请求是否已经被访问过
        :param request: 
        :return: True表示已经访问过；False表示未访问过
        """
        if request.url in self.visited_url:
            return True
        self.visited_url.add(request.url)
        return False

    def open(self):
        """
        开始爬去请求时，调用
        :return: 
        """
        print('open replication')

    def close(self, reason):
        """
        结束爬虫爬取时，调用
        :param reason: 
        :return: 
        """
        print('close replication')

    def log(self, request, spider):
        """
        记录日志
        :param request: 
        :param spider: 
        :return: 
        """
        print('repeat', request.url)
```



### 10.其他配置

```python
# -*- coding: utf-8 -*-

# Scrapy settings for step8_king project
#
# For simplicity, this file contains only settings considered important or
# commonly used. You can find more settings consulting the documentation:
#
#     http://doc.scrapy.org/en/latest/topics/settings.html
#     http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html
#     http://scrapy.readthedocs.org/en/latest/topics/spider-middleware.html

# 1. 爬虫名称
BOT_NAME = 'step8_king'

# 2. 爬虫应用路径
SPIDER_MODULES = ['step8_king.spiders']
NEWSPIDER_MODULE = 'step8_king.spiders'

# Crawl responsibly by identifying yourself (and your website) on the user-agent
# 3. 客户端 user-agent请求头
# USER_AGENT = 'step8_king (+http://www.yourdomain.com)'

# Obey robots.txt rules
# 4. 禁止爬虫配置
# ROBOTSTXT_OBEY = False

# Configure maximum concurrent requests performed by Scrapy (default: 16)
# 5. 并发请求数
# CONCURRENT_REQUESTS = 4

# Configure a delay for requests for the same website (default: 0)
# See http://scrapy.readthedocs.org/en/latest/topics/settings.html#download-delay
# See also autothrottle settings and docs
# 6. 延迟下载秒数
# DOWNLOAD_DELAY = 2


# The download delay setting will honor only one of:
# 7. 单域名访问并发数，并且延迟下次秒数也应用在每个域名
# CONCURRENT_REQUESTS_PER_DOMAIN = 2
# 单IP访问并发数，如果有值则忽略：CONCURRENT_REQUESTS_PER_DOMAIN，并且延迟下次秒数也应用在每个IP
# CONCURRENT_REQUESTS_PER_IP = 3

# Disable cookies (enabled by default)
# 8. 是否支持cookie，cookiejar进行操作cookie
# COOKIES_ENABLED = True
# COOKIES_DEBUG = True

# Disable Telnet Console (enabled by default)
# 9. Telnet用于查看当前爬虫的信息，操作爬虫等...
#    使用telnet ip port ，然后通过命令操作
# TELNETCONSOLE_ENABLED = True
# TELNETCONSOLE_HOST = '127.0.0.1'
# TELNETCONSOLE_PORT = [6023,]


# 10. 默认请求头
# Override the default request headers:
# DEFAULT_REQUEST_HEADERS = {
#     'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
#     'Accept-Language': 'en',
# }


# Configure item pipelines
# See http://scrapy.readthedocs.org/en/latest/topics/item-pipeline.html
# 11. 定义pipeline处理请求
# ITEM_PIPELINES = {
#    'step8_king.pipelines.JsonPipeline': 700,
#    'step8_king.pipelines.FilePipeline': 500,
# }



# 12. 自定义扩展，基于信号进行调用
# Enable or disable extensions
# See http://scrapy.readthedocs.org/en/latest/topics/extensions.html
# EXTENSIONS = {
#     # 'step8_king.extensions.MyExtension': 500,
# }


# 13. 爬虫允许的最大深度，可以通过meta查看当前深度；0表示无深度
# DEPTH_LIMIT = 3

# 14. 爬取时，0表示深度优先Lifo(默认)；1表示广度优先FiFo

# 后进先出，深度优先
# DEPTH_PRIORITY = 0
# SCHEDULER_DISK_QUEUE = 'scrapy.squeue.PickleLifoDiskQueue'
# SCHEDULER_MEMORY_QUEUE = 'scrapy.squeue.LifoMemoryQueue'
# 先进先出，广度优先

# DEPTH_PRIORITY = 1
# SCHEDULER_DISK_QUEUE = 'scrapy.squeue.PickleFifoDiskQueue'
# SCHEDULER_MEMORY_QUEUE = 'scrapy.squeue.FifoMemoryQueue'

# 15. 调度器队列
# SCHEDULER = 'scrapy.core.scheduler.Scheduler'
# from scrapy.core.scheduler import Scheduler


# 16. 访问URL去重
# DUPEFILTER_CLASS = 'step8_king.duplication.RepeatUrl'


# Enable and configure the AutoThrottle extension (disabled by default)
# See http://doc.scrapy.org/en/latest/topics/autothrottle.html

"""
17. 自动限速算法
    from scrapy.contrib.throttle import AutoThrottle
    自动限速设置
    1. 获取最小延迟 DOWNLOAD_DELAY
    2. 获取最大延迟 AUTOTHROTTLE_MAX_DELAY
    3. 设置初始下载延迟 AUTOTHROTTLE_START_DELAY
    4. 当请求下载完成后，获取其"连接"时间 latency，即：请求连接到接受到响应头之间的时间
    5. 用于计算的... AUTOTHROTTLE_TARGET_CONCURRENCY
    target_delay = latency / self.target_concurrency
    new_delay = (slot.delay + target_delay) / 2.0 # 表示上一次的延迟时间
    new_delay = max(target_delay, new_delay)
    new_delay = min(max(self.mindelay, new_delay), self.maxdelay)
    slot.delay = new_delay
"""

# 开始自动限速
# AUTOTHROTTLE_ENABLED = True
# The initial download delay
# 初始下载延迟
# AUTOTHROTTLE_START_DELAY = 5
# The maximum download delay to be set in case of high latencies
# 最大下载延迟
# AUTOTHROTTLE_MAX_DELAY = 10
# The average number of requests Scrapy should be sending in parallel to each remote server
# 平均每秒并发数
# AUTOTHROTTLE_TARGET_CONCURRENCY = 1.0

# Enable showing throttling stats for every response received:
# 是否显示
# AUTOTHROTTLE_DEBUG = True

# Enable and configure HTTP caching (disabled by default)
# See http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html#httpcache-middleware-settings


"""
18. 启用缓存
    目的用于将已经发送的请求或相应缓存下来，以便以后使用
    
    from scrapy.downloadermiddlewares.httpcache import HttpCacheMiddleware
    from scrapy.extensions.httpcache import DummyPolicy
    from scrapy.extensions.httpcache import FilesystemCacheStorage
"""
# 是否启用缓存策略
# HTTPCACHE_ENABLED = True

# 缓存策略：所有请求均缓存，下次在请求直接访问原来的缓存即可
# HTTPCACHE_POLICY = "scrapy.extensions.httpcache.DummyPolicy"
# 缓存策略：根据Http响应头：Cache-Control、Last-Modified 等进行缓存的策略
# HTTPCACHE_POLICY = "scrapy.extensions.httpcache.RFC2616Policy"

# 缓存超时时间
# HTTPCACHE_EXPIRATION_SECS = 0

# 缓存保存路径
# HTTPCACHE_DIR = 'httpcache'

# 缓存忽略的Http状态码
# HTTPCACHE_IGNORE_HTTP_CODES = []

# 缓存存储的插件
# HTTPCACHE_STORAGE = 'scrapy.extensions.httpcache.FilesystemCacheStorage'


"""
19. 代理，需要在环境变量中设置
    from scrapy.contrib.downloadermiddleware.httpproxy import HttpProxyMiddleware
    
    方式一：使用默认
        os.environ
        {
            http_proxy:http://root:woshiniba@192.168.11.11:9999/
            https_proxy:http://192.168.11.11:9999/
        }
    方式二：使用自定义下载中间件
    
    def to_bytes(text, encoding=None, errors='strict'):
        if isinstance(text, bytes):
            return text
        if not isinstance(text, six.string_types):
            raise TypeError('to_bytes must receive a unicode, str or bytes '
                            'object, got %s' % type(text).__name__)
        if encoding is None:
            encoding = 'utf-8'
        return text.encode(encoding, errors)
        
    class ProxyMiddleware(object):
        def process_request(self, request, spider):
            PROXIES = [
                {'ip_port': '111.11.228.75:80', 'user_pass': ''},
                {'ip_port': '120.198.243.22:80', 'user_pass': ''},
                {'ip_port': '111.8.60.9:8123', 'user_pass': ''},
                {'ip_port': '101.71.27.120:80', 'user_pass': ''},
                {'ip_port': '122.96.59.104:80', 'user_pass': ''},
                {'ip_port': '122.224.249.122:8088', 'user_pass': ''},
            ]
            proxy = random.choice(PROXIES)
            if proxy['user_pass'] is not None:
                request.meta['proxy'] = to_bytes（"http://%s" % proxy['ip_port']）
                encoded_user_pass = base64.encodestring(to_bytes(proxy['user_pass']))
                request.headers['Proxy-Authorization'] = to_bytes('Basic ' + encoded_user_pass)
                print "**************ProxyMiddleware have pass************" + proxy['ip_port']
            else:
                print "**************ProxyMiddleware no pass************" + proxy['ip_port']
                request.meta['proxy'] = to_bytes("http://%s" % proxy['ip_port'])
    
    DOWNLOADER_MIDDLEWARES = {
       'step8_king.middlewares.ProxyMiddleware': 500,
    }
    
"""

"""
20. Https访问
    Https访问时有两种情况：
    1. 要爬取网站使用的可信任证书(默认支持)
        DOWNLOADER_HTTPCLIENTFACTORY = "scrapy.core.downloader.webclient.ScrapyHTTPClientFactory"
        DOWNLOADER_CLIENTCONTEXTFACTORY = "scrapy.core.downloader.contextfactory.ScrapyClientContextFactory"
        
    2. 要爬取网站使用的自定义证书
        DOWNLOADER_HTTPCLIENTFACTORY = "scrapy.core.downloader.webclient.ScrapyHTTPClientFactory"
        DOWNLOADER_CLIENTCONTEXTFACTORY = "step8_king.https.MySSLFactory"
        
        # https.py
        from scrapy.core.downloader.contextfactory import ScrapyClientContextFactory
        from twisted.internet.ssl import (optionsForClientTLS, CertificateOptions, PrivateCertificate)
        
        class MySSLFactory(ScrapyClientContextFactory):
            def getCertificateOptions(self):
                from OpenSSL import crypto
                v1 = crypto.load_privatekey(crypto.FILETYPE_PEM, open('/Users/wupeiqi/client.key.unsecure', mode='r').read())
                v2 = crypto.load_certificate(crypto.FILETYPE_PEM, open('/Users/wupeiqi/client.pem', mode='r').read())
                return CertificateOptions(
                    privateKey=v1,  # pKey对象
                    certificate=v2,  # X509对象
                    verify=False,
                    method=getattr(self, 'method', getattr(self, '_ssl_method', None))
                )
    其他：
        相关类
            scrapy.core.downloader.handlers.http.HttpDownloadHandler
            scrapy.core.downloader.webclient.ScrapyHTTPClientFactory
            scrapy.core.downloader.contextfactory.ScrapyClientContextFactory
        相关配置
            DOWNLOADER_HTTPCLIENTFACTORY
            DOWNLOADER_CLIENTCONTEXTFACTORY

"""



"""
21. 爬虫中间件
    class SpiderMiddleware(object):

        def process_spider_input(self,response, spider):
            '''
            下载完成，执行，然后交给parse处理
            :param response: 
            :param spider: 
            :return: 
            '''
            pass
    
        def process_spider_output(self,response, result, spider):
            '''
            spider处理完成，返回时调用
            :param response:
            :param result:
            :param spider:
            :return: 必须返回包含 Request 或 Item 对象的可迭代对象(iterable)
            '''
            return result
    
        def process_spider_exception(self,response, exception, spider):
            '''
            异常调用
            :param response:
            :param exception:
            :param spider:
            :return: None,继续交给后续中间件处理异常；含 Response 或 Item 的可迭代对象(iterable)，交给调度器或pipeline
            '''
            return None
    
    
        def process_start_requests(self,start_requests, spider):
            '''
            爬虫启动时调用
            :param start_requests:
            :param spider:
            :return: 包含 Request 对象的可迭代对象
            '''
            return start_requests
    
    内置爬虫中间件：
        'scrapy.contrib.spidermiddleware.httperror.HttpErrorMiddleware': 50,
        'scrapy.contrib.spidermiddleware.offsite.OffsiteMiddleware': 500,
        'scrapy.contrib.spidermiddleware.referer.RefererMiddleware': 700,
        'scrapy.contrib.spidermiddleware.urllength.UrlLengthMiddleware': 800,
        'scrapy.contrib.spidermiddleware.depth.DepthMiddleware': 900,

"""
# from scrapy.contrib.spidermiddleware.referer import RefererMiddleware
# Enable or disable spider middlewares
# See http://scrapy.readthedocs.org/en/latest/topics/spider-middleware.html
SPIDER_MIDDLEWARES = {
   # 'step8_king.middlewares.SpiderMiddleware': 543,
}


"""
22. 下载中间件
    class DownMiddleware1(object):
        def process_request(self, request, spider):
            '''
            请求需要被下载时，经过所有下载器中间件的process_request调用
            :param request:
            :param spider:
            :return:
                None,继续后续中间件去下载；
                Response对象，停止process_request的执行，开始执行process_response
                Request对象，停止中间件的执行，将Request重新调度器
                raise IgnoreRequest异常，停止process_request的执行，开始执行process_exception
            '''
            pass
    
    
    
        def process_response(self, request, response, spider):
            '''
            spider处理完成，返回时调用
            :param response:
            :param result:
            :param spider:
            :return:
                Response 对象：转交给其他中间件process_response
                Request 对象：停止中间件，request会被重新调度下载
                raise IgnoreRequest 异常：调用Request.errback
            '''
            print('response1')
            return response
    
        def process_exception(self, request, exception, spider):
            '''
            当下载处理器(download handler)或 process_request() (下载中间件)抛出异常
            :param response:
            :param exception:
            :param spider:
            :return:
                None：继续交给后续中间件处理异常；
                Response对象：停止后续process_exception方法
                Request对象：停止中间件，request将会被重新调用下载
            '''
            return None

    
    默认下载中间件
    {
        'scrapy.contrib.downloadermiddleware.robotstxt.RobotsTxtMiddleware': 100,
        'scrapy.contrib.downloadermiddleware.httpauth.HttpAuthMiddleware': 300,
        'scrapy.contrib.downloadermiddleware.downloadtimeout.DownloadTimeoutMiddleware': 350,
        'scrapy.contrib.downloadermiddleware.useragent.UserAgentMiddleware': 400,
        'scrapy.contrib.downloadermiddleware.retry.RetryMiddleware': 500,
        'scrapy.contrib.downloadermiddleware.defaultheaders.DefaultHeadersMiddleware': 550,
        'scrapy.contrib.downloadermiddleware.redirect.MetaRefreshMiddleware': 580,
        'scrapy.contrib.downloadermiddleware.httpcompression.HttpCompressionMiddleware': 590,
        'scrapy.contrib.downloadermiddleware.redirect.RedirectMiddleware': 600,
        'scrapy.contrib.downloadermiddleware.cookies.CookiesMiddleware': 700,
        'scrapy.contrib.downloadermiddleware.httpproxy.HttpProxyMiddleware': 750,
        'scrapy.contrib.downloadermiddleware.chunked.ChunkedTransferMiddleware': 830,
        'scrapy.contrib.downloadermiddleware.stats.DownloaderStats': 850,
        'scrapy.contrib.downloadermiddleware.httpcache.HttpCacheMiddleware': 900,
    }

"""
# from scrapy.contrib.downloadermiddleware.httpauth import HttpAuthMiddleware
# Enable or disable downloader middlewares
# See http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html
# DOWNLOADER_MIDDLEWARES = {
#    'step8_king.middlewares.DownMiddleware1': 100,
#    'step8_king.middlewares.DownMiddleware2': 500,
# }
```





