# coder

> Learning route and article of a backstage development programmer/gopher



GitHub地址：<a href="https://github.com/yangfx15/coder">点击进入GitHub</a>

> 防止上面的链接无法跳转：https://github.com/yangfx15/coder



说明：学习路线中**未打上链接的知识点都还未完善**，⭐️和评论越多 ，更新越快！



### 学习路线图：

![image-20230814171118235](img/image-20230814171118235.png)

 

### 文章目录

> 1. 算法与数据结构
>    1. 高频算法题Top系列
>    2. 海量数据问题
>    3. 智力测试题
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
>       1. 入门
>       2. Web应用
>       3. 并发安全
>       4. go-swagger框架
>    2. Python 语言基础与高级
>       1. 入门
> 9. 架构设计
>    1. 架构优化
>    2. 常见系统设计
>       1. 短链接生成系统
>       2. 微博架构
>       3. 网约车系统
>       4. 网盘系统
> 10. 分布式与高可用
>        1. CAP/BASE 理论
>        2. 分布式事务
>        3. 系统高可用设计
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
>     2. 码间逸事
>     3. 个人面经



# 1. 算法与数据结构

## 1.1 高频算法题 Top 系列

### 1）<a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484365&idx=1&sn=d0a1195fa364ec9780132900fd21c90d&chksm=ecac57c4dbdbded26bbcb03a222f74cf60f99a668d6deaae11731bd075fed799cc3f5d65296d&token=55044055&lang=zh_CN#rd">排序类题目（涵盖入门和进阶题目，含代码解析）</a>

#### 题目列表

- 排序链表（No148）
- 合并区间（No56）

- 数组中第K个最大元素（No215）
- 寻找两个正序数组的中位数（No4）

#### 附加

* [高效排序算法之快排](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484200&idx=1&sn=5ea830216cf99ae87ff82f55978ef7a1&chksm=ecac5721dbdbde37d01b069b1dd61cf417aed5ca4e8caac5ac572ac5d375c43de96f42f70868&scene=21#wechat_redirect)
* [最大数（No179）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484207&idx=1&sn=77098e438cb574d8707fe44f9eb51298&chksm=ecac5726dbdbde30c17a8ae8abe7450dd7bdccd4d732c395c48bb923aaf7538bc93ddee525e5&scene=21#wechat_redirect)



### 2）链表操作类

#### 入门题目

- 反转链表（No206）
- 删除链表的倒数第K个节点（No19）
- 相交链表（No160）
- 重排链表（No143）
- 删除排序链表中的重复元素II（No82）

#### 进阶题目

- K个一组反转链表（No25）
- 反转链表II（No92）
- 奇偶链表（No328）
- 回文链表（234）
- 合并K个排序链表（No23）



### 3）数据结构之堆栈、队列、Map类

#### 入门题目

- 最小栈（No155）
- 栈实现队列（No232）
- 循环队列（No622）
- 数组中第K个最大元素（No215-堆实现）

#### 进阶题目

- LRU缓存（No146）
- 最长连续序列（No128）
- 数据流的中位数（No295）
- LFU缓存（No460）



### 4）二分法类

#### 入门题目

- 排序数组中查找元素的第1个和最后1个位置（No34）
- 寻找峰值（No162）
- 搜索二维矩阵（No74）
- 搜索旋转排序数组（No33）
- x的平方根（No69）

#### 进阶题目

- 搜索二维矩阵 II（No240）
- 按权重随机选择（No528）



### 5）双指针类

#### 入门题目

- 最长回文子串（No5）
- 无重复字符的最长子串（No3）
- 最大连续1的个数 III（No1004）
- 删除链表的倒数第N个节点（No19）

#### 进阶题目

- 四数之和（No18）
- 最小覆盖子串（No76）



### 6）BFS（广度优先搜索）

#### 入门题目

- 二叉树的锯齿形（No103）
- 二叉树的序列化（No297）
- 岛屿数量（No200）

#### 进阶题目

- 课程表（No207-拓扑排序）
- 克隆图（No133）
- 打开转盘锁（No752）
- 网格中的最短路径（No1293）



### 7）DFS（深度优先搜索）

#### 入门题目

- 二叉树的直径（No543）
- 二叉树中的最大路径和（No124）
- 二叉树的最近公共祖先（No236）
- 验证二叉搜索树（No98）
- 单词拆分（No139）

#### 进阶题目

- 矩阵中的最长递增（No329）
- 二叉搜索树中第K小的元素（No230）
- 字符串解码（No394）
- 复原IP地址（No93）
- 解数独（No37）
- 编辑距离（No72）
- 交错字符串（No97）
- N皇后（No51）

#### 分治/回溯题目

- 从前序和中序获取树（No105）
- 子集（No78）
- 括号生产（No22）
- 全排列（No46）
- 组合总和（No39）



### 8）单调栈/队列类

#### 入门题目

- 每日温度（No739）

#### 进阶题目

- 滑动窗口最大值（No239）



### 9）前缀和

#### 入门题目

- 最大子序和（No53）
- 连续的子数组和（No523）

#### 进阶题目

- 滑动窗口最大值（No239）



### 10）动态规划

#### 入门题目

- 不同路径（No62）
- 最小路径和（No64）
- 最长上升子序列（No300）
- 买卖股票的最佳时机（No121）
- 最大正方形（No221）
- 打家劫舍II（No213）
- 最长有效括号（No32）

#### 进阶题目

- 跳跃游戏II（No45）
- 字符串匹配（No44）
- 零钱兑换 II（No518）
- 解码方法（No91）
- 最长公共子序列（No1143）
- 编辑距离（No72）
- 正则表达式匹配（No10）



## 1.2 智力题和海量数据问题

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484469&idx=1&sn=ef1f8e05dd4656a3f0379d953050fe71&chksm=ecac503cdbdbd92abd193c55778e7aa6216d7c66666a12ed2b457000a520b45ceed45297a919&token=55044055&lang=zh_CN#rd">经典海量数据问题</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484385&idx=1&sn=f2b84bb6906bfc11b3f3a496019430b0&chksm=ecac57e8dbdbdefe637eb44a28ddf54760e6ee70577a85ea3cc5211f95fadf24e16a9af2707e&token=55044055&lang=zh_CN#rd">整理一些不考智力的智力题</a>



# 2. 计算机网络

## 2.1 TCP/UDP 协议

- [TCP 三次握手与四次挥手](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484223&idx=1&sn=cc70375c292d8e3ba966771965d75efc&chksm=ecac5736dbdbde204e0a4f243ff7b575602738c502e1b90cb5a21e848a3d7bf65cc5b4bd29bd&scene=21#wechat_redirect)



## 2.2 HTTP/HTTPs 协议

- [告诉老默，我想用 HTTPs 了](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484330&idx=1&sn=95b0ed6637ad7dcea0983fa1b7abe662&chksm=ecac57a3dbdbdeb53c6b264a9e91abd4e90b4cad33b09b8cedc9b4ce771ad458c3e8066f2d68&scene=21#wechat_redirect)



## 2.3 TCP/IP 网络编程

## 2.4 Cookie 与 Session

* [状态管理小能手：Cookie 和 Session](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485011&idx=1&sn=298e9e6ccf4a0403675efca14b76c748&chksm=ecac525adbdbdb4cb9ada7f57d998fe8ef2747924e154fa7fb7860aaa971eded44c213e4b3ff#rd)



# 3. 操作系统与计算机组原

### 3.1 内存管理

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484428&idx=1&sn=732ab17abcb07262d48b5bec64bdb8d6&chksm=ecac5005dbdbd9132424f59f0cf6188ce5412c75083fc781b2bc3c68592759c1a1b0542635cf&token=1255539741&lang=zh_CN#rd">一文带你认识内存管理</a>

### 3.2 进程调度

### 3.3 网络 IO 模型

### 3.4 进程间通信



# 4. 数据库

### 4.1 MySQL

- [MySQL硬核知识点总结](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484042&idx=1&sn=1620b3df43419745708f6f4c60a9ad9a&chksm=ecac5683dbdbdf95197214ec82fe119ce0bc64c95b6806a8124a381f2c01131d41d6678b87e6&scene=21#wechat_redirect)
- [MySQL高可用：分库分表你学废了吗？](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484972&idx=1&sn=421dc791c11563d9681782963a6b65d9&chksm=ecac5225dbdbdb33852d455876fded0df4f08f7e211675b8ea14d3bf7bd2512c5cfcee6951c4#rd)
- [解锁MySQL的黑科技：事务与隔离](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484917&idx=1&sn=b37e3edd5752e68a646edacec85f7395&chksm=ecac51fcdbdbd8ea1939e6fb7aa493a9304856cd6eafb6d417bb586b5cb073ae63721b55be93#rd)
- [一张图看懂 SQL 执行过程](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484875&idx=1&sn=cc741c241cb4b0ee54713b9a2469e54e&chksm=ecac51c2dbdbd8d48bdb3d0c922f98b160f8ab54537454adbdfdca925b28a7019e7d74c8353f#rd)



### 4.2 Redis

- [Redis数据结构的底层原理](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484077&idx=1&sn=05b873433b81351db83db4d90a569192&chksm=ecac56a4dbdbdfb2562a08f47e9c2b240008903f6cd8b91df6cde69bdb2814908d771be846cb&scene=21#wechat_redirect)
- [Redis IO多路复用](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484093&idx=1&sn=e173c294e8362e43fb36abf27d20cc30&chksm=ecac56b4dbdbdfa271cedbf3cd9e7c87ad7c7318d493f560b4fd814f277a3d50415ee8fd56f4&scene=21#wechat_redirect)
- <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484373&idx=1&sn=ca61fbfae1958908688d8d404ef9acaf&chksm=ecac57dcdbdbdeca1b0458e2bf91aae680b85867a5ef4c359fea22d96519675a61338e99c093&token=55044055&lang=zh_CN#rd">Redis持久化都说不明白？那今天先到这吧~</a>
- [救命！只有我还不明白Redis主从复制的原理吗？](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485225&idx=1&sn=761feb02c973cc7b2ecdcc9b9124b070&chksm=ecac5320dbdbda36d147b027db413b324b5e0d31760e71b5b52fbec9af51bb3261475a27b973#rd)
- [深入浅出Redis高可用：哨兵机制](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485363&idx=1&sn=c3568d661ab3136bf38accb0b3720192&chksm=ecac53badbdbdaac73ae2e87ae827b94564e9cc06b17a3ba82148a32913bc9690da99b81ed2b#rd)
- [Redis高可用：武林秘籍存在集群里，那稳了~](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485487&idx=1&sn=b005f993944ed72783675d110de410f1&chksm=ecac5c26dbdbd53001969519a6c90b156600ac6fe43ada939f5ac5335f644981cffcc726f212#rd)
- [Redis如何为 List/Set/Hash 的元素设置单独的过期时间](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485532&idx=1&sn=52dc4322513628e4e6f4d289eb6247f5&chksm=ecac5c55dbdbd543a1cc0ca0b147db63cc1a7ebe46bbefeec6182b0a0bc92c838c6f6c7108c4#rd)



### 4.3 ElasticSearch

- [Es 数据库从入门到实践](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484152&idx=1&sn=99d17f281baaa62f0b50d6ffb9f07354&chksm=ecac56f1dbdbdfe739e5977ec228eaee0755312611568408df49cdf55f2e6d0f0ca94aaa2d41&scene=21#wechat_redirect)



# 5. 消息中间件

* [应对流量高峰的利器——消息中间件](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485140&idx=1&sn=62aa2a762363cc8c704fc427ef50a2b6&chksm=ecac52dddbdbdbcb3f2aa1e0178e17acfa0357b3f945e15caf92f62a3cfbce4c28b3a143250c#rd)



### 5.1 Kafka

* [走近Kafka：大数据领域的不败王者](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485668&idx=1&sn=df37577dcd24d4a7cc32408af7b18be7&chksm=ecac5ceddbdbd5fb3fca2ddf12f7e3e8d81a0f1fa33bc2cadd5284e2c3b9c925eeb3b864e9fe&token=786211625&lang=zh_CN#rd)
* [深入浅出Kafka：高可用、顺序消费及幂等性](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485692&idx=1&sn=ac6e0961e888cd69697512030fcabeda&chksm=ecac5cf5dbdbd5e3967ffd986214d9363ddadeb202c86e9fa6507c3f643b8917881c8aceb6ce#rd)



### 5.2 RabbitMQ

* [深入浅出RabbitMQ：顺序消费、死信队列和延时队列](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485213&idx=1&sn=62e175c9fe41b5db13081dc8c975b298&chksm=ecac5314dbdbda02183d57501c4e15e9e432689c0ef94cfd17f0a062c544ce65e0e781091c83#rd)



# 6. Linux 入门与进阶

### 6.1 常用 Linux 命令大全

- [常用 Linux 命令大全](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484024&idx=1&sn=67d881a51cb315aa86697d5f4e49da03&chksm=ecac5671dbdbdf67a5b969fa621485f3967f613b42a275b8313d371c4b58d701d8b8291a00f4&scene=21#wechat_redirect)



### 6.2 Linux 系统高级操作



# 7. 重构与设计模式

### 7.1 重构

* [外甥女问我什么是代码洁癖，我是这么回答的...](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484696&idx=1&sn=d49f5d1fdddf54733aab6a4b98eca1c2&chksm=ecac5111dbdbd807045cfca4b6f521b75443cdcde5c311acb4ea9e82cf5e171b032a38ad9bac#rd)
* [让代码在键盘上跳“华尔兹”的10大原则](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485995&idx=1&sn=8df64b40589df43a75343cafe7c85ae0&chksm=ecac5e22dbdbd7349159a867c9d5d7f6a181ac91c0354ce21e98d62bfec7b0078a2611ae0b7e#rd)



### 7.2 设计模式

- [单例模式](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484007&idx=1&sn=3a4f8d40811db9c4f202ae52f810729b&chksm=ecac566edbdbdf78fc03d35c52c6d5a71c4696607b22781fbdf89e2f02809da6c3c04473b710&scene=21#wechat_redirect)
- 工厂模式
- 建造者模式
- 适配器设计模式
- 装饰模式
- 策略模式
- 观察者模式



# 8. 编程语言

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
- <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484344&idx=1&sn=66545b924dffa4ccb6d0cd334f56ef93&chksm=ecac57b1dbdbdea76ea384e2cf36bca800ce8961a4a9002d9d98f2e2ea722355905a43f98f82&token=55044055&lang=zh_CN#rd">map都不知道，怎么写代码？</a>
- [深入浅出内存管理：空间分配及逃逸分析](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485750&idx=1&sn=e98b1546e9a822ec87da48f5c7f17dca&chksm=ecac5d3fdbdbd42967da00fd73d758ab5b2c7a2fe5dcb02042c2843ed1a20e2d7a086acb3eb5#rd)



#### 4）周边常用工具

- [Go-swagger 应用](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483965&idx=1&sn=5d07d2df5a19db748e4646803e60cdd1&chksm=ecac5634dbdbdf22dcc63709cf627e03aa49ee21a8c990fe111fe5f4080e2d9553f2a2f09e41&scene=21#wechat_redirect)
- [Swagger 自动生成 API 文档（开发者实用工具）](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484288&idx=1&sn=974a56eaa286c7c83c25b9c077e8388b&chksm=ecac5789dbdbde9fbf82682500f3c753d0538ed69e1d7bdc9d87a31eed8dca3f91334e2a56ab&scene=21#wechat_redirect)

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484452&idx=1&sn=14f0267e9bf15ae94fc53cb90eda7c3c&chksm=ecac502ddbdbd93b37ab5cf11cd38660c760a844e4455990a53500cd4986befd031a4a68f093&token=55044055&lang=zh_CN#rd">gRPC响应ChatGPT流式问答</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484618&idx=1&sn=e6c36c497df7acac9239747e6fe36142&chksm=ecac50c3dbdbd9d57d3662b840a9f653f2bca931909ab01a751aa374a31e25e3f0e7d905afd4&token=1097967098&lang=zh_CN#rd">Go错误码设计与管理实践</a>
* [听说你会内存分析？来，pprof一下](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247486018&idx=1&sn=68e7a7176c5a3558705fd14d6d5391d1&chksm=ecac5e4bdbdbd75d8f039e9737609e7781cfb76a5e48c1fd79a655e1409f997052a1ae855d99#rd)



### 8.2 Python

#### 1）入门

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484497&idx=1&sn=4ef211b62caa60d97dd1fd26c2898f8d&chksm=ecac5058dbdbd94e68dce8034f01765370ba8fbc30436b998b12e786bde294923612532e8ac4&token=55044055&lang=zh_CN#rd">Python入门篇（上）</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484501&idx=1&sn=f913e2a8811f35f3d694cfe492327845&chksm=ecac505cdbdbd94a6f04d347be615f9ebe2c14ce934da15c22a3bb38906da38474e147a71422&token=55044055&lang=zh_CN#rd">Python入门篇（中）</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484539&idx=1&sn=7b66b06d4011597f448248fe41c41b72&chksm=ecac5072dbdbd96407419718b0408653a9fa54825741cbdf60257e7d815d57c40a3975c52b5a&token=55044055&lang=zh_CN#rd">Python入门篇（下）</a>



# 9. 架构设计

### 9.1 常见系统设计

- [听说你学过架构设计？来，弄个短链系统](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484230&idx=1&sn=89cd8cc063700baa28d2190287742a38&chksm=ecac574fdbdbde591a43b97147cd4e201eff5b0f8a16e607ea0e82885db0226481ad10e60911&scene=21#wechat_redirect)
- [听说你学过架构设计？来，弄个微博系统](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484310&idx=1&sn=33a4df6dea986ca17da75cf8513b1dbb&chksm=ecac579fdbdbde89f92927f16505eef7d31c826b6d911f1257c3bdfa7e16f133b666bc8492d0&scene=21#wechat_redirect)
- <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484653&idx=1&sn=44bef7d1e9530018da6be609fc4e57d2&chksm=ecac50e4dbdbd9f23ac9d1ce81628662e482ff57c95c51b8a6fb9858f8128793066128acbca7&token=1097967098&lang=zh_CN#rd">听说你学过架构设计？来，弄个打车系统</a>
- [听说你会架构设计？来，弄一个评论系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484740&idx=1&sn=a92c71bf4b70d611ca880ee38d4393bf&chksm=ecac514ddbdbd85b6c8ca03e900d2f7850fb2795aa7d1e8b1d73a4bb9a1c4e527eb309b480cf#rd)
- [听说你会架构设计？来，弄一个网盘系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484952&idx=1&sn=baa9da16bb52cbf6ee17d9868cacc5c6&chksm=ecac5211dbdbdb07ea31acb4fbf77cfab87c2d45bd89637d86b23499b0e2f6bc101686883748#rd)
- [听说你会架构设计？来，弄一个公交&地铁乘车系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485064&idx=1&sn=c640638fd65e559c3a747bbdfcf189d4&chksm=ecac5281dbdbdb9770e2f048a23e71064176a30b044f60add9b273683279fc541e2031de3a6e#rd)
- [听说你会架构设计？来，弄一个微信群聊系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485201&idx=1&sn=5b14c748222a42cb474c01d07bfc8601&chksm=ecac5318dbdbda0ee59db60d7502afcafea5e120039b091fffd06e2eeacd9c456f0ca6b9cf7c#rd)
- [听说你会架构设计？来，弄一个红包系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485554&idx=1&sn=55446985bb6b185babbe444fc73423cf&chksm=ecac5c7bdbdbd56dd5ea8ecdb8cbd472cb2f73fb949ae196ff819df7858459fdb9925669bbe1#rd)
- [听说你会架构设计？来，弄一个交友系统](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485865&idx=1&sn=74e7e9136ebd4c0285d23c3f0b5e1fdc&chksm=ecac5da0dbdbd4b60a07194733acd33c3b2c42f878eb75a7dbbd990f6000e6cf0a2588c219c0#rd)



### 9.2 架构优化



# 10. 分布式与高可用

### 10.1 分布式

* [深入浅出：分布式、CAP 和 BASE 理论](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484896&idx=1&sn=60dd09486fc9ecc652af917d8a311419&chksm=ecac51e9dbdbd8ffc10b79699ea7e4a8fb00aabc743b15cc5c3311970a9e3046592cbb879364#rd)
* [说出来你可能不信，分布式锁竟然这么简单...](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485092&idx=1&sn=f948b050ae2656e00e841c14c75c9aef&chksm=ecac52addbdbdbbb0b7c622972cdac50b25e4be177993df8282ba9f6f22fab11ea37ecd9cf11#rd)
* [算法江湖：揭秘分布式框架下的四大高手](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485935&idx=1&sn=d9120e18bc2ed5f8cb9eabffd0f3eca9&chksm=ecac5de6dbdbd4f0b83ac8bf3558cb9c8bd8bc9572123298e6b71c7fffc02c5185e4fceb84ce#rd)
* [数据齐舞：深入浅出分布式事务的八奇技](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485762&idx=1&sn=7770c1de5e22b62a5cee1658432e8ff5&chksm=ecac5d4bdbdbd45dd071c9f08a47cb93649773af9170ab41dd719be4c4708c3e67cc82132fff#rd)



### 10.2 高可用、高并发

* [面试官：若我问到高可用，阁下又该如何应对呢？](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484680&idx=1&sn=42d76236dfaf5220fa09c1f36eb1d421&chksm=ecac5101dbdbd81706efd91490aa69ec56d9f92687f664cf48f52f64709ce1b5f10eaa0e3b2a#rd)

* [面试薪水被压？那是你还不懂多线程和高并发](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485109&idx=1&sn=c73198ee1095ad9270f9d27c9c1e87fd&chksm=ecac52bcdbdbdbaa52eca5855617ea88993e2136ded91aff191297f002d0e3a19b645c98d66e#rd)



# 11. 微服务

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484630&idx=1&sn=efb9ba13b0db966c5d48ab49a7a57ec0&chksm=ecac50dfdbdbd9c907885a765166cba3ccd4b3a11727e5dea3d0c077dbcdab1fc2ab64646b4b&token=1097967098&lang=zh_CN#rd">都2023了，还有人不了解微服务吗</a>



# 12. Docker 与 K8s

### 12.1 docker 容器化实践

### 12.2 k8s 入门与实践

### 12.3 docker-compose 管理工具



# 13. 数据序列化协议

### 13.1 Json

### 13.2 ProtoBuffer



# 14. 常用开发工具

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484586&idx=1&sn=8eac490d874e3fee2cf6c2ce541ed661&chksm=ecac50a3dbdbd9b51320e3c6f1884e6a01abbd9badd0e119f777d2c205af2ff6aee3b6bc4189&token=1097967098&lang=zh_CN#rd">代码管理工具的扛把子-Git</a>

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484611&idx=1&sn=a0050c2b7b27099028b16fba7749005e&chksm=ecac50cadbdbd9dc83e995da039c5ea524d665316a2ca6d774cc8a5d5abf5154e0af97440a41&token=1097967098&lang=zh_CN#rd">如何DIY你自己的Typora文档</a>
* [JetBrains 破解教程最新版本，一键激活【亲测有效】](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484726&idx=1&sn=e2dcc9100a129a695d6c9b6e03f4a70a&chksm=ecac513fdbdbd8296a31ddae121a9ea0821b3d36c13089294d1f14a91a6d0bc0501ed156b431#rd)
* [免费资源整理](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484931&idx=1&sn=b3528f69d246e93e5cabab9a4cb95bfe&chksm=ecac520adbdbdb1cf8a8ce5555801cdeb230eb6cb78c3ed5a3c3114a228e50ffa1dc34210978#rd)
* [谁懂啊，语雀故障的那7个小时我是怎么过来的](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485183&idx=1&sn=ac23358c549d926df8e6d046c4f35e80&chksm=ecac52f6dbdbdbe0d96e795f0247ebda84c74ae4b125620672e63dfb149b0a1a3438868f18b1#rd)



# 15. 饮码江湖

### 15.1 README

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484522&idx=1&sn=d716a78a0526dd69291832be705d20b6&chksm=ecac5063dbdbd9759252e8b31fec30d9b7ebf4d7b2c1049de34882d62925a2d20c27e2de9690&token=55044055&lang=zh_CN#rd">26岁程序员体检，脑袋中有一个鸽子蛋般的肿瘤...</a>
* [打工人回村的双节怎么过的](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485040&idx=1&sn=ecce5aa477b8878ed059f8104e154dda&chksm=ecac5279dbdbdb6f01a5a2f192a042516ebd50f3bd66f5ace41e29e1917db883f55f0498c9e9#rd)



### 15.2 码间逸事

- [七年荣耀，一如既往？](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483881&idx=1&sn=f6f682a35910f8e753ec49b5ae9e38c7&chksm=ecac55e0dbdbdcf6b3b90f45aa1123eced26c1a6a45724e2b73b03dd639dd4ce2a6579c05d15&scene=21#wechat_redirect)
- [要脸，还是要方便？](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483917&idx=1&sn=5703164683f0a3dda4f7d560e83bc1f2&chksm=ecac5604dbdbdf128d8d288c0673ae0d5fd8d7c6974fdedc37b2d5001debafa0dae29a2337db&scene=21#wechat_redirect)
- [PDD 事件看内卷](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247483986&idx=1&sn=9c28064604aa497271c6fa4eff70e9a6&chksm=ecac565bdbdbdf4da52b6c3027ce2d5f7bdfc1a1812a4d1c6bd3eb823fa86a2e902072241d5d&scene=21#wechat_redirect)

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484410&idx=1&sn=e96622284a64204e2bf2d79bdadb4370&chksm=ecac57f3dbdbdee5fff3575a44c801f5841ba700b2ff1c04b0e55a4762405a6ad49984a4a3a0&token=55044055&lang=zh_CN#rd">OpenAI 发布 GPT-4，有哪些技术上的优化或突破？</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484476&idx=1&sn=d566a34e0764bff9580079261e5903dd&chksm=ecac5035dbdbd923ba373c595b9b01c0fb387386b3300104c8f5ab32fdc659b0291369ef07d3&token=55044055&lang=zh_CN#rd">从数据标注看机器学习</a>
* [震惊！翻墙3年挣百万，非法收入被没收](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485066&idx=1&sn=fa0a23433c8cb101b9a86c4747ac30d8&chksm=ecac5283dbdbdb9564a8f9e857021cbdd8bc7110b687d8fe43bc308d303a438a279e56e7f425#rd)
* [谁懂啊，语雀故障的那7个小时我是怎么过来的](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485183&idx=1&sn=ac23358c549d926df8e6d046c4f35e80&chksm=ecac52f6dbdbdbe0d96e795f0247ebda84c74ae4b125620672e63dfb149b0a1a3438868f18b1#rd)
* [【Error】阿里全系产品崩了](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485382&idx=1&sn=b914f23bf0d68cbf0e80683ed774da04&chksm=ecac53cfdbdbdad98b93b260e02ea2e5a0577c70da3e24a81bd7faa323db956d936eb75e5c7a#rd)
* [一句“裁员广进”，揭开了多少互联网大厂的遮羞布](https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247485972&idx=1&sn=3b36922da5891b1651fc8043cdcb7494&chksm=ecac5e1ddbdbd70b0d722ec27ba3a16dfbd46b89b8de1abf76f10f0fac0efe35fe3126426436#rd)



### 15.3 面经

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484469&idx=1&sn=ef1f8e05dd4656a3f0379d953050fe71&chksm=ecac503cdbdbd92abd193c55778e7aa6216d7c66666a12ed2b457000a520b45ceed45297a919&token=55044055&lang=zh_CN#rd">经典海量数据问题</a>

* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484461&idx=1&sn=a278bc04ae023fba780ef792fead2362&chksm=ecac5024dbdbd9325eb5535fecf717085625e17f8245d0bcd4aec7d2d942cd813ec6fdeb3ae3&token=55044055&lang=zh_CN#rd">大厂面试都会问什么</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484428&idx=1&sn=732ab17abcb07262d48b5bec64bdb8d6&chksm=ecac5005dbdbd9132424f59f0cf6188ce5412c75083fc781b2bc3c68592759c1a1b0542635cf&token=1255539741&lang=zh_CN#rd">一文带你认识内存管理</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484385&idx=1&sn=f2b84bb6906bfc11b3f3a496019430b0&chksm=ecac57e8dbdbdefe637eb44a28ddf54760e6ee70577a85ea3cc5211f95fadf24e16a9af2707e&token=55044055&lang=zh_CN#rd">整理一些不考智力的智力题</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484373&idx=1&sn=ca61fbfae1958908688d8d404ef9acaf&chksm=ecac57dcdbdbdeca1b0458e2bf91aae680b85867a5ef4c359fea22d96519675a61338e99c093&token=55044055&lang=zh_CN#rd">Redis持久化都说不明白？那今天先到这吧~</a>
* <a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484344&idx=1&sn=66545b924dffa4ccb6d0cd334f56ef93&chksm=ecac57b1dbdbdea76ea384e2cf36bca800ce8961a4a9002d9d98f2e2ea722355905a43f98f82&token=55044055&lang=zh_CN#rd">map都不知道，怎么写代码？</a>
* [告诉老默，我想用 HTTPs 了](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484330&idx=1&sn=95b0ed6637ad7dcea0983fa1b7abe662&chksm=ecac57a3dbdbdeb53c6b264a9e91abd4e90b4cad33b09b8cedc9b4ce771ad458c3e8066f2d68&scene=21#wechat_redirect)

* [TCP 三次握手与四次挥手](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484223&idx=1&sn=cc70375c292d8e3ba966771965d75efc&chksm=ecac5736dbdbde204e0a4f243ff7b575602738c502e1b90cb5a21e848a3d7bf65cc5b4bd29bd&scene=21#wechat_redirect)

* [GO 垃圾回收](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484217&idx=1&sn=d464e95dcec6252b5e00a5f634a46892&chksm=ecac5730dbdbde26005ba864ab324a3a2c9b90f42c5ce037eb4f5eee36a4e55a4385421d9dd3&scene=21#wechat_redirect)

* [GPM 调度模型](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484182&idx=1&sn=6d3f54eea5622a2d7f6323cbb553fdd8&chksm=ecac571fdbdbde09cc8beb982e5df0caafdf5c87587cd3fbd69ca86c33724e9368ab957beac3&scene=21#wechat_redirect)
* [Redis IO多路复用](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484093&idx=1&sn=e173c294e8362e43fb36abf27d20cc30&chksm=ecac56b4dbdbdfa271cedbf3cd9e7c87ad7c7318d493f560b4fd814f277a3d50415ee8fd56f4&scene=21#wechat_redirect)

- [Redis数据结构的底层原理](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484077&idx=1&sn=05b873433b81351db83db4d90a569192&chksm=ecac56a4dbdbdfb2562a08f47e9c2b240008903f6cd8b91df6cde69bdb2814908d771be846cb&scene=21#wechat_redirect)
- [MySQL硬核知识点总结](http://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484042&idx=1&sn=1620b3df43419745708f6f4c60a9ad9a&chksm=ecac5683dbdbdf95197214ec82fe119ce0bc64c95b6806a8124a381f2c01131d41d6678b87e6&scene=21#wechat_redirect)

