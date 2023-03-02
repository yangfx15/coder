# coder

Learning route of some programmer



GitHub地址：https://github.com/yangfx15/coder

说明：学习路线中**未打上链接的知识点都还未完善**，⭐️和评论越多 ，更新越快哦~~



学习路线图：

![img](https://mmbiz.qpic.cn/mmbiz_png/pQBexpeYRE6wldmHiaibhKj7l9s4bFreicRlP8NvRAMvpSjGgomILfWYNgseiaTdxLwia6K8gAHyxKJ1RtnFuFzUchA/640?wx_fmt=png)

 

### 文章目录

> 1. 算法与数据结构
>    1. 高频算法题Top系列
>    2. 海量数据问题
> 2. 计算机网络
>    1. TCP/UDP 协议
>    2. HTTP 与 HTTPs
>    3. TCP/IP网络编程
>    4. Cookie 与 Session
> 3. 操作系统与组原
>    1. 虚拟内存与物理内存
>    2. 进程调度
>    3. 网络 IO 模型
>    4. 进程间通讯
> 4. 数据库知识
>    1. MySQL
>    2. Redis
>    3. ElasticSearch
> 5. 消息中间件
>    1. kafka
>    2. RabbitMQ
> 6. Linux 入门与进阶
>    1. 常用 Linux 命令大全
>    2. Linux 系统高级操作
> 7. 重构与设计模式
>    1. 重构
>    2. 设计模式
> 8. 编程语言
>    1. Go 语言基础与高级
>       1. 三个入门程序
>       2. Go 数据结构
>       3. 并发安全
>       4. go-swagger框架
>    2. Java 语言基础与高级
> 9. 架构设计
>    1. 架构优化
>    2. 常见系统设计
>       1. 短链接生成系统
>       2. 微博架构
>       3. 网约车系统
>       4. 网盘系统
> 10. 分布式与高可用
>     1. CAP/BASE 理论
>     2. 分布式事务
>     3. 系统高可用设计
> 11. 微服务
>     1. 单体架构的演进
>     2. 服务发现
>     3. 服务间通信
> 12. Docker 与 K8s
>     1. docker容器化实践
>     2. k8s入门与实践
>     3. docker-compose管理工具
> 13. 数据序列化协议
>     1. Json
>     2. ProtoBuffer
> 14. 常用开发工具
>     1. Git
>     2. Typroa
> 15. 饮码江湖
>     1. README
>     2. 码间新闻



## 1. 算法与数据结构

### 1.1 高频算法题 Top 系列

#### 1）排序类题目

**入门题目：**

- 排序链表（No148）
- 合并区间（No56）
- [高效排序算法之快排](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484200&idx=1&sn=5ea830216cf99ae87ff82f55978ef7a1&chksm=ecac5721dbdbde37d01b069b1dd61cf417aed5ca4e8caac5ac572ac5d375c43de96f42f70868&scene=21#wechat_redirect)

**进阶题目：**

- [最大数（No179）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484207&idx=1&sn=77098e438cb574d8707fe44f9eb51298&chksm=ecac5726dbdbde30c17a8ae8abe7450dd7bdccd4d732c395c48bb923aaf7538bc93ddee525e5&scene=21#wechat_redirect)
- 数组中第K个最大元素（No215）
- 寻找两个正序数组的中位数（No4）



#### 2）链表操作类

**入门题目：**

- 反转链表（No206）
- 删除链表的倒数第K个节点（No19）
- 相交链表（No160）
- 重排链表（No143）
- 删除排序链表中的重复元素II（No82）

**进阶题目：**

- K个一组反转链表（No25）
- 反转链表II（No92）
- 奇偶链表（No328）
- 回文链表（234）
- 合并K个排序链表（No23）



#### 3）数据结构之堆栈、队列、Map类

**入门题目：**

- 最小栈（No155）
- 栈实现队列（No232）
- 循环队列（No622）
- 数组中第K个最大元素（No215-堆实现）

**进阶题目：**

- LRU缓存（No146）
- 最长连续序列（No128）
- 数据流的中位数（No295）
- LFU缓存（No460）



#### 4）二分法类

**入门题目：**

- 排序数组中查找元素的第1个和最后1个位置（No34）
- 寻找峰值（No162）
- 搜索二维矩阵（No74）
- 搜索旋转排序数组（No33）
- x的平方根（No69）

**进阶题目：**

- 搜索二维矩阵 II（No240）
- 按权重随机选择（No528）



#### 5）双指针类

**入门题目：**

- 最长回文子串（No5）
- 无重复字符的最长子串（No3）
- 最大连续1的个数 III（No1004）
- 删除链表的倒数第N个节点（No19）

**进阶题目：**

- 四数之和（No18）
- 最小覆盖子串（No76）



#### 6）BFS（广度优先搜索）

**入门题目：**

- 二叉树的锯齿形（No103）
- 二叉树的序列化（No297）
- 岛屿数量（No200）

**进阶题目：**

- 课程表（No207-拓扑排序）
- 克隆图（No133）
- 打开转盘锁（No752）
- 网格中的最短路径（No1293）



#### 7）DFS（深度优先搜索）

**入门题目：**

- 二叉树的直径（No543）
- 二叉树中的最大路径和（No124）
- 二叉树的最近公共祖先（No236）
- 验证二叉搜索树（No98）
- 单词拆分（No139）

**进阶题目：**

- 矩阵中的最长递增（No329）
- 二叉搜索树中第K小的元素（No230）
- 字符串解码（No394）
- 复原IP地址（No93）
- 解数独（No37）
- 编辑距离（No72）
- 交错字符串（No97）
- N皇后（No51）

**分治/回溯题目：**

- 从前序和中序获取树（No105）
- 子集（No78）
- 括号生产（No22）
- 全排列（No46）
- 组合总和（No39）



#### 8）单调栈/队列类

**入门题目：**

- 每日温度（No739）

**进阶题目：**

- 滑动窗口最大值（No239）



#### 9）前缀和

**入门题目：**

- 最大子序和（No53）
- 连续的子数组和（No523）

**进阶题目：**

- 滑动窗口最大值（No239）



#### 10）动态规划

**入门题目：**

- 不同路径（No62）
- 最小路径和（No64）
- 最长上升子序列（No300）
- 买卖股票的最佳时机（No121）
- 最大正方形（No221）
- 打家劫舍II（No213）
- 最长有效括号（No32）

**进阶题目：**

- 跳跃游戏II（No45）
- 字符串匹配（No44）
- 零钱兑换 II（No518）
- 解码方法（No91）
- 最长公共子序列（No1143）
- 编辑距离（No72）
- 正则表达式匹配（No10）



### 1.2 海量数据问题

#### 1）重复数搜索问题

#### 2）Top 频率问题



## 2. 计算机网络

### 2.1 TCP/UDP 协议

- [TCP 三次握手与四次挥手](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484223&idx=1&sn=cc70375c292d8e3ba966771965d75efc&chksm=ecac5736dbdbde204e0a4f243ff7b575602738c502e1b90cb5a21e848a3d7bf65cc5b4bd29bd&scene=21#wechat_redirect)



### 2.2 HTTP/HTTPs 协议

- [告诉老默，我想用 HTTPs 了](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484330&idx=1&sn=95b0ed6637ad7dcea0983fa1b7abe662&chksm=ecac57a3dbdbdeb53c6b264a9e91abd4e90b4cad33b09b8cedc9b4ce771ad458c3e8066f2d68&scene=21#wechat_redirect)



### 2.3 TCP/IP 网络编程

### 2.4 Cookie 与 Session



## 3. 操作系统与计算机组原

### 3.1 内存管理

### 3.2 进程调度

### 3.3 网络 IO 模型

### 3.4 进程间通信



## 4. 数据库

#### 4.1 MySQL

- [MySQL硬核知识点总结](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484042&idx=1&sn=1620b3df43419745708f6f4c60a9ad9a&chksm=ecac5683dbdbdf95197214ec82fe119ce0bc64c95b6806a8124a381f2c01131d41d6678b87e6&scene=21#wechat_redirect)

  

### 4.2 Redis

- [Redis数据结构的底层原理](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484077&idx=1&sn=05b873433b81351db83db4d90a569192&chksm=ecac56a4dbdbdfb2562a08f47e9c2b240008903f6cd8b91df6cde69bdb2814908d771be846cb&scene=21#wechat_redirect)
- [Redis IO多路复用](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484093&idx=1&sn=e173c294e8362e43fb36abf27d20cc30&chksm=ecac56b4dbdbdfa271cedbf3cd9e7c87ad7c7318d493f560b4fd814f277a3d50415ee8fd56f4&scene=21#wechat_redirect)



### 4.3 ElasticSearch

- [Es 数据库从入门到实践](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484152&idx=1&sn=99d17f281baaa62f0b50d6ffb9f07354&chksm=ecac56f1dbdbdfe739e5977ec228eaee0755312611568408df49cdf55f2e6d0f0ca94aaa2d41&scene=21#wechat_redirect)



## 5. 消息中间件

### 5.1 Kafka

### 5.2 RabbitMQ



## 6. Linux 入门与进阶

### 6.1 常用 Linux 命令大全

- [常用 Linux 命令大全](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484024&idx=1&sn=67d881a51cb315aa86697d5f4e49da03&chksm=ecac5671dbdbdf67a5b969fa621485f3967f613b42a275b8313d371c4b58d701d8b8291a00f4&scene=21#wechat_redirect)



### 6.2 Linux 系统高级操作



## 7. 重构与设计模式

### 7.1 重构



### 7.2 设计模式

- [单例模式](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484007&idx=1&sn=3a4f8d40811db9c4f202ae52f810729b&chksm=ecac566edbdbdf78fc03d35c52c6d5a71c4696607b22781fbdf89e2f02809da6c3c04473b710&scene=21#wechat_redirect)
- 工厂模式
- 建造者模式
- 适配器设计模式
- 装饰模式
- 策略模式
- 观察者模式



## 8. 编程语言

### 8.1 Golang

#### 1）入门

- [人生苦短，Let's Go](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483729&idx=1&sn=60fe53ac1ec196faa1637d5b539628cf&chksm=ecac5558dbdbdc4e3ddc1cb982d0026dc4b8dd631e563d49342e8a9f2f3db7edac32a6c0aa48&scene=21#wechat_redirect)
- [Go 安装与配置运行](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483828&idx=1&sn=15b08998d4febc600b9b87b852ac7491&chksm=ecac55bddbdbdcab6fae68becd1a4cf1260f6cee33e449a1a98a24accf1fbd71cc5714a8cdf7&scene=21#wechat_redirect)
- [程序建筑的原材料：基本数据类型](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483891&idx=1&sn=6ad3b114e23727e1b31b0a9a218e4357&chksm=ecac55fadbdbdcec62ba538ae6f30793ddec1ede4111f166eba307133a29eba7ad61da31e4d3&scene=21#wechat_redirect)
- [复合数据类型](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483937&idx=1&sn=456a1e7fa269a557e465fa2fef8f6d2a&chksm=ecac5628dbdbdf3e2008d02bbb4faba9785b972ee52d3450c9c29c1a2c1499e2a77360108813&scene=21#wechat_redirect)
- [函数](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483949&idx=1&sn=c31f4b7cd5ed6fe72c686d6e31c41550&chksm=ecac5624dbdbdf32b8fd08d49eb19744d3caf48e89be712646ab9ea3dbe7962c460f51cd6af3&scene=21#wechat_redirect)
- [接口-朦胧之美](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483953&idx=1&sn=5da81ccd439d0398bb78f1500ba11f31&chksm=ecac5638dbdbdf2e48126e94ce446ab0b488863855a0424706f274a4fa8ce9b82a98e15ff3e8&scene=21#wechat_redirect)



#### 2）Go 编写 Web 应用

- [快速搭建一个简洁服务器，Go 称霸微服务架构的秘诀](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483866&idx=1&sn=34fd6ecc9723ada234ab8e3743a248e3&chksm=ecac55d3dbdbdcc50e531ecd396c728d4420825b838153e1931f992633f0bfee258c8f46bc6a&scene=21#wechat_redirect)
- [编写 Web 应用（上篇）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483921&idx=1&sn=0b4c72b11bd9fa42f06ed64a169e0d98&chksm=ecac5618dbdbdf0e8f9068c134d85cc5839a491910662a3ad24d280b53eb789b16dd11f0889b&scene=21#wechat_redirect)
- [编写 Web 应用（下篇）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483941&idx=1&sn=857d5373b2152ce977f9e7338eb6f063&chksm=ecac562cdbdbdf3a71c2bd6a0faeb9ec5158f8cfef2401d34df5ce18868495edbd2840646867&scene=21#wechat_redirect)



#### 3）并发安全

- [CSP 并发实现](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483991&idx=1&sn=5df1ccb1e1e28be56a38dd37e59b2199&chksm=ecac565edbdbdf48a98d7ca44c9dda1fd549d805d29dc01f6794661cee515aa749856182b0ef&scene=21#wechat_redirect)
- [共享内存并发实现](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484002&idx=1&sn=1e9352dea27bd4c531d607dae664615d&chksm=ecac566bdbdbdf7d440dac7b1c460430e596665340f312d8bf9c6c1efc35be067262bf3fd206&scene=21#wechat_redirect)
- [GO 垃圾回收](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484217&idx=1&sn=d464e95dcec6252b5e00a5f634a46892&chksm=ecac5730dbdbde26005ba864ab324a3a2c9b90f42c5ce037eb4f5eee36a4e55a4385421d9dd3&scene=21#wechat_redirect)
- [并发才是时代的主题，goroutine 帮你快速实现](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483844&idx=1&sn=09b1841623eaedd17babb6b408eb72d7&chksm=ecac55cddbdbdcdb818cf493920199d1a0ccd7fa25bfd010a55a078abc0772022a4af4eb29d0&scene=21#wechat_redirect)
- [GPM 调度模型](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484182&idx=1&sn=6d3f54eea5622a2d7f6323cbb553fdd8&chksm=ecac571fdbdbde09cc8beb982e5df0caafdf5c87587cd3fbd69ca86c33724e9368ab957beac3&scene=21#wechat_redirect)



#### 4）Go-swagger 框架

- [Go-swagger 应用](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483965&idx=1&sn=5d07d2df5a19db748e4646803e60cdd1&chksm=ecac5634dbdbdf22dcc63709cf627e03aa49ee21a8c990fe111fe5f4080e2d9553f2a2f09e41&scene=21#wechat_redirect)
- [Swagger 自动生成 API 文档（开发者实用工具）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484288&idx=1&sn=974a56eaa286c7c83c25b9c077e8388b&chksm=ecac5789dbdbde9fbf82682500f3c753d0538ed69e1d7bdc9d87a31eed8dca3f91334e2a56ab&scene=21#wechat_redirect)



### 8.2 Java



## 9. 架构设计

### 9.1 架构优化

### 9.2 常见系统设计

- [听说你学过架构设计？来，弄个短链系统](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484230&idx=1&sn=89cd8cc063700baa28d2190287742a38&chksm=ecac574fdbdbde591a43b97147cd4e201eff5b0f8a16e607ea0e82885db0226481ad10e60911&scene=21#wechat_redirect)
- [听说你学过架构设计？来，弄个微博系统](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484310&idx=1&sn=33a4df6dea986ca17da75cf8513b1dbb&chksm=ecac579fdbdbde89f92927f16505eef7d31c826b6d911f1257c3bdfa7e16f133b666bc8492d0&scene=21#wechat_redirect)



## 10. 分布式与高可用

### 10.1 CAP/BASE 理论

### 10.2 分布式事务

### 10.3 系统高可用设计



## 11. 微服务

### 11.1 单体架构演进

### 11.2 服务发现

### 11.3 服务间通信



## 12. Docker 与 K8s

### 12.1 docker 容器化实践

### 12.2 k8s 入门与实践

### 12.3 docker-compose 管理工具



## 13. 数据序列号协议



## 14. 常用开发工具

### 14.1 Git

### 14.2 Typora



## 15. 饮码江湖

### 15.1 README



### 15.2 码间新闻

- [七年荣耀，一如既往？](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483881&idx=1&sn=f6f682a35910f8e753ec49b5ae9e38c7&chksm=ecac55e0dbdbdcf6b3b90f45aa1123eced26c1a6a45724e2b73b03dd639dd4ce2a6579c05d15&scene=21#wechat_redirect)
- [要脸，还是要方便？](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483917&idx=1&sn=5703164683f0a3dda4f7d560e83bc1f2&chksm=ecac5604dbdbdf128d8d288c0673ae0d5fd8d7c6974fdedc37b2d5001debafa0dae29a2337db&scene=21#wechat_redirect)
- [PDD 事件看内卷](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483986&idx=1&sn=9c28064604aa497271c6fa4eff70e9a6&chksm=ecac565bdbdbdf4da52b6c3027ce2d5f7bdfc1a1812a4d1c6bd3eb823fa86a2e902072241d5d&scene=21#wechat_redirect)

