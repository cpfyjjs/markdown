# Python笔记

## Python中的数据结构

### 数据结构

数据结构是将数据组织在一起的数据结构。数据结构是用来存储一系列相关的东西。在Python中有四种基本的内建的数据结构，分别是List、Tuple、Dictionary、以及Set。大部分情况下，这四种数据结构基本上能满足大多数操作。其中List、Dictionary的功能最为强大。当然也有一些其他的数据结构可供选择，例如Collection、Array、Heapq、Bisect、Werakref、Copy以及Pprint。

***

#### 列表 list

列表是可变的，它支持序列的标准操作索引、分片、连接和乘法，同时也支持一些特有的方法：元素赋值、元素删除、分片赋值、以及列表方法。

```python
a = [] #声明一个列表
a.append(1)
print (a)   #1

a+= [1,2,3,4]
print (a)   #[1, 1, 2, 3, 4]

b = a   #a,b引用的是同一个变量所以b改变，a也改变
b.append(3)
print (a)   #[1, 1, 2, 3, 4, 3]

del a[2:4]
print (a)   #[1, 1, 4, 3]

print (a[1]) #1
```

现在通过冒泡排序算法，集中展示一下列表的特性,函数没有返回值，但他同样改变带排序的内容。

```python
def bubblesorted(A = []):
    n = len(A)
    sorted = False
    while not sorted:
        sorted = True
        for i in range(1,n):
            if A[i-1] > A[i]:
                A[i-1],A[i] = A[i],A[i-1]
                sorted = False
        n-= 1

a = [2,3,5,1,6,8,4,9,7]

bubblesorted(a)
print (a)
#[1, 2, 3, 4, 5, 6, 7, 8, 9]
```

当然我们可以通过浅拷贝，技术来实现列表复制。

```python
def bubblesorted(A = []):
    B = A[:]	#完成数据的浅拷贝
    n = len(B)
    sorted = False
    while not sorted:
        sorted = True
        for i in range(1,n):
            if B[i-1] > B[i]:
                B[i-1],B[i] = B[i],B[i-1]
                sorted = False
        n-= 1
    return B

a = [2,3,5,1,6,8,4,9,7]

b = bubblesorted(a)
print (a)
[2, 3, 5, 1, 6, 8, 4, 9, 7]
print (b)
#[1, 2, 3, 4, 5, 6, 7, 8, 9]
```

当然如果，是嵌套的列表，A[ : ] 并不能完整的拷贝复制数据，效果是不是看起很感人。使用这种方法复制和使用copy.copy函数的结果一样，只能复制浅层数据，如果想复制深层数据，就需要用到copy.decopy函数

```python
a = [[1,2,3],2,3,4,[4,5,6,7]]
b = a[:]
b.append(3)
print (b)   #[[1, 2, 3], 2, 3, 4, [4, 5, 6, 7], 3]
print (a)   #[[1, 2, 3], 2, 3, 4, [4, 5, 6, 7]]

b[0][0] = 10000000
print (b)   #[[10000000, 2, 3], 2, 3, 4, [4, 5, 6, 7], 3]
print (a)   #[[10000000, 2, 3], 2, 3, 4, [4, 5, 6, 7]]

c = copy.copy(a)
c[0][0] = [0]
print (c)   #[[[0], 2, 3], 2, 3, 4, [4, 5, 6, 7]]
print (a)   #[[[0], 2, 3], 2, 3, 4, [4, 5, 6, 7]]

d = copy.deepcopy(a)
d[0][0] = 1111
print(d)    #[[1111, 2, 3], 2, 3, 4, [4, 5, 6, 7]]
print(a)    #[[[0], 2, 3], 2, 3, 4, [4, 5, 6, 7]]
```

所以当使用列表的数据结构过程中需要注意到列表的易变性，这方面字典和列表极为相像。

***

#### 字典 Dictionary

字典是Python中的唯一一个内建的映射类型。字典中值并没有特殊的顺序，但都存在一个特定的键（Key）。键可以是数字、字符串、甚至元组以及任何可哈希对象。值（value）可以是任何对象。一个键只能对应一个值，但一个值可以对应多个键。

字典的创建方式

```python
d = {'Alex':123,'LiLei':234,'HanMeiMei':456}
print (d['Alex'])   #123

items = [('name','Alex'),('Age',18)]
d = dict(items)
print (d)   #{'name': 'Alex', 'Age': 18}

items =  [('name',('Alex','HanLei')),('Age',(18,24))]
d = dict(items)
print (d)   #{'name': ('Alex', 'HanLei'), 'Age': (18, 24)}

d = dict(name = 'Hanlei',age = 18)
print (d) #{'name': 'Hanlei', 'age': 18}
```

字典的基本行为在很多方面和序列极为相似（sequence）：  

-[ ] len(d)	返回序列长度
-[ ] d[k]         返回关联在键key上的值
-[ ] d[k] = v   将值v 关联在key上，如果k存在将覆盖，如果不存在将创建
-[ ] del d[k]   删除键为k的项
-[ ] k in d       检查d中是否有键k的项

#### 默认字典 Defaultdict

这个类型除了在处理不存在的键的操作之外与普通的字典完全相同。当查找一个不存在的键操作发生时，它的default_factory会被调用，提供一个默认的值，并且将这对键值存储下来。其他的参数同普通的字典方法dict()一致，一个defaultdict的实例同内建dict一样拥有同样地操作。

```python
from collections import defaultdict

s = "the quick brown fox jumps over the lazy dog"

words = s.split()
location = defaultdict(list)
for m, n in enumerate(words):
    location[n].append(m)

print location

# defaultdict(<type 'list'>, {'brown': [2], 'lazy': [7], 'over': [5],'fox': [3],'dog': [8], 'quick': [1], 'the': [0, 6], 'jumps': [4]})
```

#### 优先级对列 Heapq

heapq模块使用一个用堆实现的优先级队列。堆是一种简单的有序列表，并且置入了堆的相关规则。

```python
import heapq
 
heap = []
 
for value in [20, 10, 30, 50, 40]:
    heapq.heappush(heap, value)
 
while heap:
    print heapq.heappop(heap) 	#10,20,30,40,50
```

### 自定义数据结构

#### 优先级队列

```python
class MaxPQ:
    def __init__(self,items=[]):
        self.items = items
        
    def isempty(self):
        return len(self.items) == 0

    def __len__(self):
        return len(self.items)
    #得到最大元素
    def getmax(self):
        return self.items[0]
	#将存储的元素有序化
    def initialize(self):
        size = len(self.items) -1
        for i in range(size,-1,-1):
            self.sink(i)
      
b = [1,2,43,5,6,7,8,13,45,687,34,56,8]
a = MaxPQ(b)
a.initialize()
print(a.delmax())   #687
print(a.delmax())   #56
print(a.delmax())   #45
```

插入一个新的元素

```python
def swim(self,h):
    if h > 0:            
        p = (h - 1)//2
        while self.items[h] > self.items[p]:
            self.items[h],self.items[p] = self.items[p],self.items[h]
            h = p
            p = (h - 1)//2
            if h == 0: break

def insert(self,v):
    self.items.append(v)
    h = len(self.items)-1
    self.swim(h)
```

删除最大元素。

```python
def sink(self,p):
    while 2*(p+1) <= len(self.items)-1:
        #找出左右孩子中较大的元素
        v = (2*p+1)if self.items[2*p +1] > self.items [2*(p+1)] else 2*(p+1)
        if self.items[v] > self.items[p]:
            #将父子节点元素内容互换
            self.items[v],self.items[p] = self.items[p],self.items[v]
            p = v
        else:break

    if (2*p+1) == len(self.items)-1:
        v = 2 * p + 1
        if self.items[v] > self.items[p]:
            self.items[v],self.items[p] = self.items[p],self.items[v]

def delmax(self):
    temp = self.items[0]
    self.items[0] = self.items[-1]
    del self.items[-1]
    self.sink(0)
    return temp
```

优先级对列在数据结构中极为重要，他的全部操作的复杂度都会在log(n)的范围之内。同时在线性扫描算法，最短路径算法中都扮演着极为重要的角色。

#### 有向图 digraph

图的介绍

图作为数据结构的重要组成部分，扮演着极为重要的角色。图本身作为非线性的数据结构，往往用来描述定义一组对象的二元关系，比如城市交通，以及人与人之间的关系。实现图的方式有两种，邻接矩阵和邻接表。邻接矩阵描述的是节点与节点之间的关系而邻接表描述的是节点与边之间的关系。若边（u，v)所对应顶点u 与 v 的次序无所谓，则称为无向边，例如同学之间的关系。反之，称为有向边，例如单行道。

有向图实现邻接矩阵实现方式。为了节省空间使用字典的方式来表示节点与节点之间的关系。

```python
class DiGraph:
    #有向图初始化
    def __init__(self,data = None):
        self.graph = {} #图属性
        self.node = {}  #节点属性
        self.adj = {}   #邻接矩阵属性
        self.age = self.adj
```

增加节点节点

本身需要是可哈希的对象，例如整型、字符串、元组等。列表不能作为字典的键。节点和属性对应着字典中的键和值。

```python
def add_node(self,n,attr_dict = None ,**attr):
    #设置节点属性
    if attr_dict is None:
        attr_dict = attr
    else:
        try:
            #更新节点属性
            attr_dict.update(attr)
        except AttributeError:
            raise TypeError("The attr_dict argument must be a dictionary.")
    
    if n not in self.node:
        self.node[n] = attr_dict
    else:
        self.node[n].update(attr_dict)
```

增加一系列的节点

```python
def add_nodes_from(self,nodes,**attr):
    for n in nodes:
        try:
            #判断是否是新加入的节点
            newnodes = n not in self.nodes:
        except TypeError:
            nn,ndict = n
            if nn not in self.node:
                newdict = attr.copy()
                newdict.update(ndict)
                self.node[nn] = newdict
            else:
                olddict = slef.node[nn]
                odldict.update(attr)
                olddict.update(ndict)
                
        if newnode:
            self.node[n] = attr.copy()
        else:
            self.node[n].update(attr)
```

移除节点

在移除节点时，我们同时需要移除对应的边

```python
def remove_node(self,n):
        try:
            nbrs = self.adj[n]
            #删除节点
            del self.node[n]
        except KeyError:    #节点不存在与图中
            raise TypeError ("不存在节点%n."%(n,))
        
        #删除边（n,u)    
        for u in nbrs:
            del self.adj[n][u]
        del self.adj[n]
        #删除边（u,n) 
        for u in self.adj:
            if u in self.adj[u]:
                del self.adj[u][n]
```

移除一系列的节点

```python
def remove_nodes_from(self,nbunch):
    for n in nbunch:
        remove_node(n)
```

添加边，如果边上节点不存在，同时添加节点

```python
def add_edge(self,u,v,attr_dict = None,**attr):
    if attr_dict ids None:
        attr_dict = attr
    else:
        try :
            attr_dict.update(attr)
        except AttributeError:
            raise AttributeError("attr_dict 必须是一个字典"）
            
    #添加节点
    if u not in self.node:
        self.node[u] = {}
        self.adj[u] = {}
        
    if v not in self.node:
        self.node[v] = {}
        self.adj[v] = {}
        
    datadict = self.adj[u].get(v,{})
    datadict.update(attr_dict)
```

添加一系列的边

```python
def add_edges_from(self,ebunch,attr_dict = None,**attr):
    #设置边属性
    if attr_dcit is None:
        attr_dict = attr
    else:
        try:
            attr_dict.update(attr)
        except AttributeError:
            raise TypeError("attr_dict 必须是一个字典")
    
    #判断输入元组，或者列表的长度，采取不同的策略        
    for e in ebunch:
        ne = len(e)
        if ne == 3:
            u,v ,dd = e
            assert hasattr(dd,"update")
        elif ne ==2:
            u,v = e 
            dd = {}
        else:
            raise TypeError("Edge tuple %s must be a 2-tuple or 3-tuple."%(e,))
        
        #判断节点是否存在   
        if u not in self.node:
            self.node[n] = {}
            self.adj[n] = {}
        
        if v  not in self.node:
            self.node[v] = {}
            self.adj[v] = {}
        
        #更新边的属性   
        datadict = self.adj[u].get(v,{})
        datadict = update(attr_dict)
        datadict.update(dd)
        self.pred[v][u] = datadict
```

删除边

```python
def remove_age(self,u,v):
    try:
        del self.adj[u][v]
    except KeyError:
        raise keyError("The edge %s-%s not in graph."%(u,v))
```

删除一系列的边

```python
def remove_edges_from(self,ebunch):
    for e in enbunch:
        (u,v) = e[0:2]
        if v in self.adj[u]:
            del self.adj[u][v]
```

以上就是有向图的一些基本操作，如果对有向图很有兴趣，可以查看networkx的源码，里面有详细的定义以及实现。









