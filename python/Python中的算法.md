## Python中的算法

### 结合Grasshopper

#### 模拟processing的自治智能体

这个仿鱼群运动是模拟processing的自治智能体。听起来是不是很高大上，其实很简单。他只需要尊重几条重要的特性。

- 1.自治智能体对环境的感知能力是有限的。

- 2.自治智能体需要处理来自外部的信息，并由此计算具体的行为。
- 3.自治智能体没有领导者。我们并不需要关心第三特性。

上面三句话引自《代码本色》的第六章自治智能体。对于自治智能体，细胞自动机，神经网络感兴趣的同学可以找找看。

现在我们看看如何实现，所需的知识不多，大概高中物理的力学基础部分即可。

S为距离，V是速度，t是时间，a为加速度，m为质量，F是力。

位移公式：S = So + V * t   

速度公式：V = Vo +a * t

加速度公式：F = a * m

代码部分运动到了基础的面向对象的部分，不过不懂代码的也不用担心。面想对象可以理解为对事物的简化，它更关注于对象之间的行为，而不关心它是怎么做到的。就像你可以开车，刹车，但你不用懂发动机如何工作。

在这里我抽象了三个基本的类。头羊，追随的的羊，还有捕猎的狼。

```python
import rhinoscriptsyntax as rs #调入rhino基本的库函数，函数是想象过程的，所以很简单。
import random as r #调入随机函数
import math 
from  scriptcontext import sticky as st

class Mover(object):
    def __init__(self):
        self.pos=[0,0,0]  #定义基本类的位置属性。
        self.vec=[0,0,0]  #速度
        self.acc=[0,0,0]  #加速度

    #定义一个函数更新物体的运动状态
    def update(self):
        self.vec=rs.VectorAdd(self.vec,self.acc) #速度公式：V = Vo +a * t  时间为1.
        self.pos=rs.VectorAdd(self.pos,self.vec) #位移公式：S = So + V * t 
        
    #定义一个恒定力，让物体运动有一个方向。
    def force(self,force,mass):
        scale=1/mass #加速度公式：F = a * m
        self.acc=rs.VectorScale(force,scale)
        
    #定义一个变化力，让物体运动方向产生一定的随机波动
    def wind(self):
        a=r.randint(-5,5) #这个函数是产生-5到5之间的随机数。
        b=r.randint(-5,5)
        c=r.randint(-5,5)
        vec=rs.VectorCreate([0,0,0],[a/4,b/4,c/4])
        self.vec=rs.VectorAdd(self.vec,vec)

    #将物体的运动速度限制在一定的范围之内
    def limited(self,max):  
        if rs.VectorLength(self.vec)>max:
            self.vec=rs.VectorUnitize(self.vec) #单位化向量。
            self.vec=rs.VectorScale(self.vec,max)
            self.acc=[0,0,0]
            
class OtMover(Mover):
    """
    定义了一个追随排斥的类，继承于Mover类。继承的优势在于
    父类的方法函数不用再写一遍可以直接调用。
    """
    def attraction(self,sca,otpos):
        #吸引力，这里你可以理解为追随头羊
        distance=rs.Distance(self.pos,otpos)
        vector=(rs.VectorCreate(otpos,self.pos))*sca
        self.vec=rs.VectorAdd(self.vec,vector)

    def reject(self,otpos):
        #排斥力当两个物体相距小于50时，排除力产生，理解为躲避狼
        distance=rs.Distance(self.pos,otpos)
        if distance<50: #这里进行判断在50这个范围内是否有危险
            vector=rs.VectorCreate(otpos,self.pos)
            vector=(rs.VectorUnitize(vector))*(distance-50) 
            #如果有危险，距离危险越近，逃离的速度越快。
            self.vec=rs.VectorAdd(self.vec,vector) #位移公式：S = So + V * t 
    def checkEdges(self,x,y,z): 
    #检查边界例如：当物体超越顶面时，从地面出现。
        if self.pos[0] < 0:
            self.pos[0] = x
        if self.pos[0] > x:
            self.pos[0] = 0
        if self.pos[1] < 0:
            self.pos[1] = y
        if self.pos[1] > y:
            self.pos[1] = 0
        if self.pos[2] < 0:
            self.pos[2] = z
        if self.pos[2] > z:
            self.pos[2] = 0

    def display(self): #函数可以调用函数，这样更易简化代码
        self.wind()
        self.force(force,10)
        self.limited(10)
        self.checkEdges(x,y,z)
        self.update()

class MiMover(Mover):
    #定义一个运动的类，继承于Mover类
    
    #在一定范围内运动，如果到达边界，运动方向进行方向，弹回来了。
    def checkEdges(self,x,y,z):
        if self.pos[0]<0 or self.pos[0]>x:
            self.vec[0]*=-1
        if self.pos[1]<0 or self.pos[1]>y:
            self.vec[1]*=-1
        if self.pos[2]<0 or self.pos[2]>z:
            self.vec[2]*=-1
            
    def display(self):
        self.wind()
        self.force(force,10)
        self.limited(8)
        self.checkEdges(x,y,z)
        self.update()
c=[]
d=[]
op=[]
ov=[]

if "move" not in st:
    st["move"]=MiMover()  #初始化一个头羊。
    r.seed(seed)
if Toogle:
    st["move"].display() #更新斗羊位置
    point=st["move"].pos   #得到头羊的位置
    vector=st["move"].vec #得到头羊的运动方向。
else:
    del st["move"]

for i in range(num): #初始化一群大灰狼
    names='move%d'%i 
    if names not in st:
        st[names]=OtMover()
    if Toogle:
        st[names].display() #更新狼的位置
        a=st[names].pos #得到狼的位置
        b=st[names].vec #得到狼的方向
        op.append(a) #把每一头狼的位置放在一起
        ov.append(b) #把每一头狼的方向放在一起


for item in range(number):
    name='otmove%d'%item
    if name not in st: #初始化基本物体。。。
        st[name]=OtMover()
    if Toogle:
        st[name].attraction(sca,st["move"].pos)#追随头羊
        for i in range(num):
            names="move%d"%i
            st[name].reject(st[names].pos) #逃避每一头狼。
        st[name].display()
        a=st[name].pos  
        b=st[name].vec
        c.append(a) #把每一羊的位置放在一起
        d.append(b) #把每一羊的方向放在一起
    else:
        del st[name]

```

#### 利用磁场线模拟人流疏散

这一次介绍的是利用python调用grasshopper的vector面板下的field。实现一种模拟人流的。只是伪模拟，之间相互排斥没写，但这次重点不是这，而是如何调用grasshopper写的的类，即磁场线。最后当然你也可以直接使用grasshopper和循环实现。好了还是利用高中的物理知识。就一条   位移公式：S = So + V * t

```python
from Grasshopper.Kernel.Types import GH_LineCharge  
from Rhino.Geometry import Point3d
from Rhino.Geometry import Vector3d 
from  scriptcontext import sticky as st     

#这些是导入所需的基本库函数。
class Mover(object):
    def __init__(self,pos):
        self.pos=pos  #定义基本类的位置属性。

L=[]    #定义一个列表用来存放点

for i in range(len(pos)):     #写一个循环，len函数是求出列表的长度
    names='move%d'%i   #字符串方法

    if names not in st:   #判断mover是否存在
        st[names]=Mover(pos[i])
        gH_LineCharge =  GH_LineCharge()  #实例化一个类，这个类是墙体，有排斥力
        gH_LineCharge.Charge = 1  # charge是排斥力的大小
        gH_LineCharge.Limits = box  #范围

        gH_LineCharge2 =  GH_LineCharge()  #实例化一个类，这个类是走道，有吸引力
        gH_LineCharge2.Charge = -num #因为是负数，具有吸引力
        gH_LineCharge2.Limits = box  #范围

    if Toogle:  #是否开始
        a=Point3d(0,0,0)   #实例化一个类

        for line in lines:   #写一个循环求出墙体的排斥的和
            gH_LineCharge.Segment = line
            b=gH_LineCharge.Force(st[names].pos)
            a=Point3d.Add(a,b)
        for line in line_field: #写一个循环求出走道的吸引力的和
            gH_LineCharge2.Segment = line
            b=gH_LineCharge2.Force(st[names].pos)
            a=Point3d.Add(a,b)

        a=Vector3d(a.X,a.Y,0)
        a.Unitize()  #Vector3d转换为单位向量
        st[names].pos=Point3d.Add(st[names].pos,a) # 位移公式：S = So + V * t 时间为零
        L.append(st[names].pos) #将点放在同一个列表中
    else:
        del st[names] #删除实例化的类，不可忘记。
```

重点不是模拟人流，而是如何调用grasshopper的类，实现一些rhinocommon中没有的东西。

#### 凸包算法

示例代码中sorted方法可以用单次循环，n的复杂度结束，主要作用是找到凸包的一个极点。退化问题没有考虑。算法复杂度最大为n平方的量级。

```python
import Rhino.Geometry as rg

pts = sorted(pts,key = lambda pt :(pt.Y,pt.X))
line = rg.Line(pts[0],pts[1])
ptsbool = [False for i in range(len(pts))]

def leftbool (line,point):
    vect01 = line.Direction
    endpt = line.PointAt(0)
    vect02 = endpt - point
    
    crossvect = rg.Vector3d.CrossProduct(vect02,vect01)
    if crossvect.Z > 0:
        return True
    else:
        return False
        
lines =[]
k =0
bool = True
while bool== True:
    startpt = pts[k]
    line = rg.Line(pts[k],pts[k-1])
    for i in range(len(pts)):
        if not(leftbool(line,pts[i])):
            if rg.Line(startpt,pts[i]).Length>0:
                line = rg.Line(startpt,pts[i])
                k = i
    lines.append(line)
    if k == 0:
        bool = False
a =lines
```



