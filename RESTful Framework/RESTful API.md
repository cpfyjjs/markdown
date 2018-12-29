## RESTful API
### 一、RESTful API 设计
#### 1.协议
https 协议
#### 2.域名
`https://api.example.com`: 会存在跨域问题
`https://example.com/api`
#### 3.版本
`https://example.com/api/v2/`
#### 4.路径
路径表示API的具体网址
在RESTful架构中，每种网址代表一种资源(resource)，所以网址中不能有动词，只能是名词，可以有复数形式。
#### 5.HTTP动词
* **GET** ：从服务器获取资源（一个或者多个）
* **POST** ：在服务器新建一个资源。
* **PUT** ：在服务器更新资源（客户端提供改变后的完整数据）
* **PATCH** ：在服务器更新资源（局部属性）

#### 6.过滤信息
```html
?limit=10
?offset=10
?page=2&per_page=100
?student_id=1
```
#### 7.状态码
* 200 OK
* 201 CREATE
* 202 Accepted
* 204 NO CONTENT
* 400 INVALL REQUEST
* 404 NO FOUND

#### 8.错误处理
`{status:200,msg:'',error:''`

#### 9.返回结果
```python
GET /collection：返回资源对象的列表（数组）
GET /collection/resource：返回单个资源对象
POST /collection：返回新生成的资源对象
PUT /collection/resource：返回完整的资源对象
PATCH /collection/resource：返回完整的资源对象
DELETE /collection/resource：返回一个空文档
```
#### 10.Hypermedia API  超媒体API
```python
{"link": {
  "rel":   "collection https://www.example.com/zoos",  #表示这个API与当前网址的关系（collection关系，并给出该collection的网址）
  "href":  "https://api.example.com/zoos",  #API路径
  "title": "List of zoos",  #API的标题
  "type":  "application/vnd.yourformat+json"  #返回类型
}}
```
