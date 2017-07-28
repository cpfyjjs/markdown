## C#中的算法

#### 结合Grasshopper

#### 最短路路径

```c#
using Rhino;
using Rhino.Geometry;
using Rhino.DocObjects;
using Rhino.Collections;

using GH_IO;
using GH_IO.Serialization;
using Grasshopper;
using Grasshopper.Kernel;
using Grasshopper.Kernel.Data;
using Grasshopper.Kernel.Types;

using System;
using System.IO;
using System.Xml;
using System.Xml.Linq;
using System.Linq;
using System.Data;
using System.Drawing;
using System.Reflection;
using System.Collections;
using System.Windows.Forms;
using System.Collections.Generic;
using System.Runtime.InteropServices;

public class Script_Instance : GH_ScriptInstance
{
  private void RunScript(List<Curve> curves, List<Point3d> points, List<int> index01, List<int> index02, Point3d start, Point3d end, ref object passed_points, ref object weight, ref object passed_curves)
  {
    ArrayList nodeList = new ArrayList();
    //声明一个 Node 节点的动态数组。
    ArrayList edgeList = new ArrayList();
    //声明一个有向边 Edge 的动态数组。
    //为edgeList添加Edge实例成员。
    for (int i = 0; i < curves.Count;i++)
    {
      Edge edge = new Edge();
      edge.StartNodeID = points[index01[i]];
      edge.EndNodeID = points[index02[i]];
      edge.Weight = curves[i].GetLength();
      edge.Passed = curves[i];
      edgeList.Add(edge);
    }
    /*
    在这里用index 的原因在于使用curve.PointAtStart会生成新的Point3d，与生成Node节点的Point3d是不同的。
    各个类之间的联系是通过Point3d建立，如果不同，联系将建立不起来。
    */
    for (int i = 0; i < curves.Count;i++)
    {
      Edge edge = new Edge();
      edge.StartNodeID = points[index02[i]];
      edge.EndNodeID = points[index01[i]];
      edge.Weight = curves[i].GetLength();
      edge.Passed = curves[i];
      edgeList.Add(edge);
    }

    for (int i = 0; i < points.Count; i++)
    {
      Node node = new Node(points[i]);
      for (int k = 0; k < edgeList.Count; k++)
      {
        Edge edge = (Edge) edgeList[k];
        if (points[i] == edge.StartNodeID)
        {
          node.EdgeList.Add(edge);
          //节点和边建立联系。
        }
      }
      nodeList.Add(node);
      //将Node实例加入到nodeList动态数组里
    }

    RoutePlanner planner = new RoutePlanner();
    RoutePlanResult result = planner.Plan(nodeList, start, end);
    planner = null;

    weight = result.Value;
    passed_points = result.ResultNode;
    //等到最短路径的point3d数组。

    Curve [] new_curves = new Curve[result.PassedCurve.Length];
    new_curves = result.PassedCurve;
    passed_curves = Curve.JoinCurves(new_curves);
    //得到最短路径
  }

  // <Custom additional code> 
  //首先，我们可以将“有向边”抽象为Edge类；
  public class Edge
  {
    public Point3d StartNodeID;
    public Point3d EndNodeID;
    public double Weight;
    public Curve Passed;
  }

  //节点则抽象成Node类，一个节点上挂着以此节点作为起点边的“出边”表。
  public class Node
  {
    private Point3d iD;
    private ArrayList edgeList;

    //构造函数
    public Node(Point3d id)
    {
      this.iD = id;
      this.edgeList = new ArrayList();
    }
    #region property
    //只能的ID，不能设置
    public Point3d ID
    {
      get { return this.iD; }
    }
    public ArrayList EdgeList
    {
      get
      {
        return this.edgeList;
      }
    }

    #endregion
  }

  //PassedPath 用于缓存计算过程中到达某个节点的权值的最小路径
  public class PassedPath
  {
    private Point3d curNodeID;          //当前节点
    private bool beProcessed;           //是否被处理过
    private double weight;              //累积的权值
    private ArrayList passedIDList;     //缓存路径节点Point3d
    public ArrayList passedCurveList;   //缓存处理过路径Curve实例
    //声明构造函数，
    public PassedPath(Point3d ID)
    {
      this.curNodeID = ID;
      this.weight = double.MaxValue;    //将当前结点的权重设置为最大值
      this.passedIDList = new ArrayList();
      this.passedCurveList = new ArrayList();
      this.beProcessed = false;
    }
    #region property
    public bool BeProcessed
    {
      get{ return this.beProcessed;}
      set{ this.beProcessed = value;}
    }
    public Point3d CurNodeID
    {
      get{ return this.curNodeID;}
    }
    public double Weight
    {
      get{ return this.weight;}
      set{ this.weight = value;
      }
    }
    public ArrayList PassedIDList
    {
      get{ return this.passedIDList;}
    }
    #endregion
  }

  //另外，还需要一个表PlanCourse来记录规划的中间结果
  //即管理了每个节点的PassedPath。
  //PlanCourse缓存从节点到任意点的最小路径值=》路径表
  public class PlanCourse
  {
    private Hashtable htPassedPath;     //声明一个私有成员函数，结构字典
    public PlanCourse(ArrayList nodeList, Point3d originID)
    {
      this.htPassedPath = new Hashtable();
      //创建一个字典，键为Point3d,值为此节点的PassedPath
      Node originNode = null;
      foreach(Node node in nodeList)
      {
        if (node.ID == originID)
          originNode = node;
        else
        {
          PassedPath pPath = new PassedPath(node.ID);
          this.htPassedPath.Add(node.ID, pPath);
        }
      }
      if(originNode == null)
        throw new Exception("The origin node is not exist!");

      this.InitializeWeight(originNode);
    }
    private void InitializeWeight(Node originNode)      //  初始化权重
    {
      if((originNode.EdgeList == null) || originNode.EdgeList.Count == 0)
      {
        return;
      }

      foreach(Edge edge in originNode.EdgeList)
        //originNode上挂着以此节点作为起点边的“出边”
      {
        PassedPath pPath = this[edge.EndNodeID];
        if(pPath == null)
        {
          continue;
        }
        //初始化出边的end.
        pPath.PassedIDList.Add(originNode.ID);
        pPath.Weight = edge.Weight;
        pPath.passedCurveList.Add(edge.Passed);
      }
    }
    public PassedPath this[Point3d nodeID]
    {
      get
      {
        return (PassedPath) this.htPassedPath[nodeID];
      }
    }
  }

    /*
    在所有的基础构建好后，规划路径的算法就很容易实施，该算法的步骤如下：
    1.用一张表（PlanCourse)记录源点到任意其他节点的最小权值，初始化这张表时，
    如果源点能直通某节点，则权值设为对应的边的权，否则设为double.MaxValue。

    2.选取没有被处理并且当前累计权值最小的节点TargetNode，
    用其边的可达性来更新到达到达其它节点的路径和权值（如果其他节点
    经过此节点后权值变小更新，否则不更新），然后标记TargetNode为已处理。

    3.重复2，直到所有的可达节点都被处理一遍。
    */

    //RoutePlanner 提供图算法常用的路径规划功能。
  public class RoutePlanner
  {
    public RoutePlanner() { }
    //获取最小路径
    public RoutePlanResult Plan(ArrayList nodeList, Point3d originID, Point3d destID)
    {
      PlanCourse planCourse = new PlanCourse(nodeList, originID);
      //从PlaneCourse取出一个当前累计权值最小，并没有处理过的节点.
      Node curNode = this.GetMinWeightRudeNode(planCourse, nodeList, originID);
      //计算过程
      while (curNode != null)
      {
        PassedPath curPath = planCourse[curNode.ID];
        foreach (Edge edge in curNode.EdgeList)
        {
          if(edge.EndNodeID != originID)
          {
            //这个地方就是为什么一定要使用输入的Point3d,而不是新生成的
            //毕竟我们是用输入的Point3d来初始化planCourse（路径表）的。
            PassedPath targetPath = planCourse[edge.EndNodeID];
            double tempWeight = curPath.Weight + edge.Weight;

            //其它节点经此节点后权值变小则更新
            if (tempWeight < targetPath.Weight)
            {
              targetPath.Weight = tempWeight;
              targetPath.PassedIDList.Clear();
              targetPath.passedCurveList.Clear();


              for (int i = 0; i < curPath.PassedIDList.Count; i++)
              {
                targetPath.PassedIDList.Add(curPath.PassedIDList[i]);
                targetPath.passedCurveList.Add(curPath.passedCurveList[i]);
              }
              targetPath.PassedIDList.Add(curNode.ID);  //添加点
              targetPath.passedCurveList.Add(edge.Passed);  //添加边
            }
          }
        }
        //标记为已处理
        planCourse[curNode.ID].BeProcessed = true;
        //获取下一个未处理的累计权值最小的节点
        curNode = this.GetMinWeightRudeNode(planCourse, nodeList, originID);
      }
      //表示规划结束
      return this.GetResult(planCourse, destID);

    }

    //从PlanCourse表中取出节点的PassedPath，这个PassedPath即是规划结果，建立字典的目的。
    private RoutePlanResult GetResult(PlanCourse planCourse, Point3d destID)
    {
      PassedPath pPath = planCourse[destID];

      if(pPath.Weight == int.MaxValue)
      {
        RoutePlanResult result1 = new RoutePlanResult(null, null, int.MaxValue);
        return result1;
      }

      Point3d[] passedNodeIDs = new Point3d[pPath.PassedIDList.Count + 1];
      for (int i = 0; i < pPath.PassedIDList.Count;i++)
      {
        passedNodeIDs[i] = (Point3d) pPath.PassedIDList[i];
      }
      //将尾节点加入到数组里。
      passedNodeIDs[pPath.PassedIDList.Count] = destID;

      Curve[] passedCurves = new Curve[pPath.passedCurveList.Count];
      for (int i = 0; i < pPath.passedCurveList.Count;i++)
      {
        passedCurves[i] = (Curve) pPath.passedCurveList[i];
      }

      RoutePlanResult result = new RoutePlanResult(passedNodeIDs, passedCurves, pPath.Weight);
      return result;
    }

    //从PlaneCourse取出一个当前累计权值最小，并没有处理过的节点
    private Node GetMinWeightRudeNode(PlanCourse planCourse, ArrayList nodeList, Point3d originID)
    {
      double weight = double.MaxValue;
      Node destNode = null;
      foreach(Node node in nodeList)
      {
        if (node.ID != originID)
        {
          PassedPath pPath = planCourse[node.ID];
          //没有处理过的节点
          if(pPath.BeProcessed)
          {
            continue;
          }
          //当前累计权值最小
          if(pPath.Weight < weight)
          {
            weight = pPath.Weight;
            destNode = node;
          }
        }
      }
      return destNode;
    }
  }

  //最短路径的结果定义
  public class RoutePlanResult
  {
    public RoutePlanResult(Point3d [] passedNode, Curve [] passedCurve, double value)
    {
      this.m_resultNode = passedNode;
      this.m_value = value;
      this.m_passedCurve = passedCurve;


    }
    public Point3d[] ResultNode
    {
      get { return m_resultNode; }
    }

    public double Value
    {
      get { return m_value; }
    }
    public Curve [] PassedCurve
    {
      get { return m_passedCurve; }
    }

    private Point3d[] m_resultNode;
    private double m_value;
    private Curve [] m_passedCurve;


  }
  // </Custom additional code> 
}
```

