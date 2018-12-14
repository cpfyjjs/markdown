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
    
    #允许的域名,即只允许爬去此域名下的网页，
    allowde_domians = ['chouti.com']
    
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

**自定义pipline**

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
        """
```







