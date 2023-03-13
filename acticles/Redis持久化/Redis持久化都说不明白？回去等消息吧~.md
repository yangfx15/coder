目录

> 1. 引言
> 2. RDB
> 3. AOF
> 4. 两种持久化机制对比
> 5. 小结



## 1. 引言

==Q：什么是 Redis？==

A：Redis，全称是远程数据服务（Remote Dictionary Server），它作为一款由 C 语言编写的内存高速缓存数据库，在互联网产品中应用甚广。

无论是国内还是国外，从五百强公司到小型初创公司都在使用 Redis，很多云服务提供商还以 Redis 为基础构建了相应的缓存服务、消息队列服务以及内存存储服务，当你使用这些服务时，实际上就是在使用 Redis。

**作为一名开发者，即便没在工作中用到它，也一定会在面试中被问到！**

![image-20230313224738375](img\2.jpg)



==Q：为何 Redis 这么重要，它有哪些常见的应用场景？==

A：究其原因，在于 Redis 是纯内存操作，且执行效率高，本身也支持丰富的数据结构。而它广泛的应用场景，包括但不限于**缓存、事件发布或订阅、分布式锁**等等。

> 不熟悉 Redis 数据结构的，可以看这篇文章【Redis底层数据结构的实现】。



==Q：Redis 所有的操作都是内存操作吗？==

A：并不是，redis 的单线程是指在**接收客户端 IO 请求进行读写时是单线程操作**。但 redis  本身是存在多线程使用场景的，比如：异步删除，持久化以及集群同步。

而 Redis 作为面试中的常考知识点，常见的八股文想必各位已经烂熟于心。但是，在互联网行情日益内卷的今天，相关的知识点，我们是否都能答出面试官想要的深度和广度呢？





比如今天我们要复习的这个知识点，Redis 的持久化机制。

==Q：Redis 常见的持久化机制有哪些？==

A：RDB 和 AOF。

![image-20230313224738375](img\image-20230313224738375.png)



## 2. RDB

#### 2.1 简介

RDB（Redis Database Backup file），即快照模式，是 Redis 默认的数据持久化方式。

RDB 实际上是 Redis 内部的一个**定时器事件**，每隔一段时间就去检查当前数据发生改变的**次数**和**时间频率**，看它们是否满足配置文件中规定的触发条件。

当条件满足时，Redis 通过操作系统调用 fork() 函数创建一个子进程，这个子进程**和父进程享有同一个地址空间、文件系统和信号量**等系统数据。之后，Redis 通过子进程遍历整个内存空间，将数据集拷贝到一个临时文件，当拷贝完成后通知父进程将新的 RDB 文件替换掉原有的文件，以完成数据持久化的操作。

同时，在持久化过程中，父进程仍然可以对外提供服务，父子进程通过操作系统的多进程 **COW（copy and write）机制**实现了数据段分离，从而保证父子进程之间互不影响。



#### 2.2 优缺点总结

在 RDB 持久化的时候，Redis fock 子进程会把 Redis 的所有数据都保存到新建的 dump.rdb 文件中，这是一个既消耗资源又浪费时间的操作。因此，Redis 服务器不能过于频繁地创建 rdb 文件，否则会严重影响服务器的性能。

除此之外，RDB 持久化的最大不足之处在于：**最后一次持久化的过程中可能会出现大量的数据丢失**。我们想象一个场景，在 RDB 进行持久化的过程中，Redis 服务器突然宕机了，这时子进程可能已经生成了 rdb 文件，但是父进程还没来得及用它覆盖掉旧的 rdb 文件，缓冲区中的文件还未保存，就会导致大量数据的丢失。

而 RDB 数据持久化的优势在于它还原速度很快，所以比较适用于大规模的数据恢复，如果对数据的完整性不是特别敏感（允许持久化过程中数据丢失的情况），那么 RDB 持久化方式就非常适合。



## 3. AOF

#### 3.1 简介

AOF，append only log file，又被称为追加模式，或日志模式。

AOF 会记录服务器执行的所有写操作命令，**并且只记录对内存有过修改的命令**，Redis 会将这些命令定时写到 appendonly.aof 文件中。

我们在服务器启动时，可以重新执行 AOF 文件来还原数据集，这个过程被称为 **命令重演**。



#### 3.2 三种持久化机制

Redis 在收到客户端修改命令后，会先进行相应的校验，如果命令无误，就立即将该命令存进缓冲区，再以一定的速率将缓冲区数据追加到 .aof 文件中。

这样，就算遇到了突发的宕机情况，也只需将存储到 aof 文件中的命令，进行一次“命令重演”就可以恢复到宕机前的状态。

在上述执行过程中，有一个很重要的环节就是**命令的写入，这是一个磁盘 IO 操作**：Redis 为了提升写入效率，它不会将内容直接写入到磁盘中，而是先将其放到一个内存缓存区（buffer），等到缓存区被填满或者满足 AOF 的持久化策略时才真正将缓存区中的内容写入到磁盘里（fsync 操作）。

AOF 的持久化策略（即 fsync 操作的频率）分为三种：

- always：服务器每写入一个命令，就调用一次 fsync 函数，将缓冲区里面的命令写入到硬盘。这种模式下，服务器出现故障，也不会丢失任何已经成功执行的命令数据，但是其执行速度非常慢；
- everysec（默认）：服务器每一秒调用一次 fsync 函数，将缓冲区里面的命令写入到硬盘。这种模式下，服务器出现故障，最多只丢失一秒钟内的执行的命令数据，通常都使用它作为 AOF 配置策略；
- no：服务器不主动调用 fsync 函数，由操作系统决定何时将缓冲区里面的命令写入到硬盘。这种模式下，服务器遭遇意外停机时，丢失命令的数量是不确定的，所以这种策略，不确定性较大，也不常用。

Redis 在遇到宕机前，如果缓存内的数据未能写入到磁盘中，那么数据仍然会有丢失的风险。而丢失命令的数量，取决于命令被写入磁盘的时间：**越早地把命令写入到磁盘中，发生意外时丢失的数据就会越少。**

由于是 fsync 是磁盘 IO 操作，所以它很慢！如果 Redis 执行一条指令就要 fsync 一次（always），将会严重地影响到 Redis 的性能。

在生产环境的服务器中，Redis 通常是默认每隔 1s 左右执行一次 fsync 操作（ everysec），这样既保持了高性能，也让数据尽可能的少丢失。

而最后一种策略（no），让操作系统来决定何时将数据同步到磁盘，这种策略存在许多不确定性，所以不建议使用。

> 注意：sync 和 fsync 函数是操作系统提供的两个函数，用来防止“延迟写”造成的缓存和文件的数据不一致问题。
>
> 其中，sync 将修改过的数据放入缓存写队列中就返回，不等待 IO 操作结束。
>
> 而 fsync 会等待 IO 操作结束再返回，它会确保修改过的块立即写到磁盘上，来保证文件数据和缓存一致。
>
> 即，Linux 系统的 fsync() 函数可以将指定文件的内容从内核缓存同步刷新到硬盘中，sync() 是异步操作。



#### 3.3 重写机制

在 AOF 的持久化策略下，Redis 在长期运行的过程中，aof 文件会越变越长。而如果机器宕机重启，命令重演整个 aof 文件就会非常耗时，这将导致 Redis 长时间无法对外提供服务。

因此，为了让 aof 文件的大小控制在合理的范围内，Redis 提供了 AOF 重写机制，即对 aof 文件进行 “瘦身” ：Redis 服务器可以创建一个新的 AOF 文件来替代现有的 AOF 文件，**新旧两个文件所保存的数据库状态是相同的，不同的是新文件没有任务冗余命令**，所以文件大小会比旧文件小很多。

Redis 提供了两种 AOF 文件的重写方式：手动执行 BGREWRITEAOF 命令，或者配置策略实现自动重写。AOF 文件重写时和 RDB 持久化过程类似，都是 fork 一个子进程来操作原 AOF 文件。

![image-20230313222524259](img\image-20230313222524259.png)

如图：AOF 的重写过程中父进程继续处理新的请求，如果有新的命令添加，会追加到 AOF 重写缓存区，然后直接追加到新的 AOF 文件中。



## 4. AOF 和 RDB 对比

| RDB 持久化机制                                               | AOF 持久化机制                             |
| ------------------------------------------------------------ | ------------------------------------------ |
| 全量备份，一次保存整个数据库                                 | 增量备份，一次只保存一个修改数据库的命令   |
| 每次执行持久化操作的间隔时间较长                             | 保存的间隔默认为一秒钟（everysec）         |
| 数据保存为二进制格式，其还原速度快                           | 使用文本格式还原数据，所以数据还原速度一般 |
| 执行 SAVE 命令时会阻塞服务器，但手动或者自动触发的 BGSAVE 不会阻塞服务器 | AOF 持久化无论何时都不会阻塞服务器         |



在 Redis4.0 以前，我们只能选择 RDB 或 AOF 作为持久化机制；而 Redis4.0 以后，我们可以将 Redis 持久化配置成混合机制，即 RDB+AOF 都作为持久化方式，详情可见 redis.conf 文件：

``` yaml
# Compress string objects using LZF when dump .rdb databases?
# For default that's set to 'yes' as it's almost always a win.
# If you want to save some CPU in the saving child set it to 'no' but
# the dataset will likely be bigger if you have compressible values or keys.
rdbcompression yes

# RDB files created with checksum disabled have a checksum of zero that will
# tell the loading code to skip the check.
rdbchecksum yes

# The filename where to dump the DB
dbfilename dump.rdb

# The name of the append only file (default: "appendonly.aof")
appendfilename "appendonly.aof"

# appendfsync always
appendfsync everysec
# appendfsync no

auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

# When a child rewrites the AOF file, if the following option is enabled
# the file will be fsync-ed every 32 MB of data generated. This is useful
# in order to commit the file to the disk more incrementally and avoid
# big latency spikes.
aof-rewrite-incremental-fsync yes

# When redis saves RDB file, if the following option is enabled
# the file will be fsync-ed every 32 MB of data generated. This is useful
# in order to commit the file to the disk more incrementally and avoid
# big latency spikes.
rdb-save-incremental-fsync yes
```

> 如果进行数据恢复时，既有 RDB 文件，又有 AOF 文件，我们应该先通过 AOF 文件来恢复数据，因为这能最大程度地保证数据的安全性。



## 5. 小结

2023 了，互联网像前几年那样疯狂扩张、抢占应用风口的时代逐渐过去！如果说整个计算机行业的浪潮正在短暂或长期退去，那么能够决定程序员是否会在沙滩上“旱泳”的是什么呢？边缘业务，抑或大龄危机？

我认为都不是，其实涨潮和退潮都是有趋势的，有些人早早捕捉到了危机，及早穿好泳衣上岸；有些人不足现状，向着大海更深处游去。他们都是互联网浪潮中的智者，智者不会担心浪潮退去，因为他们已经提前做好了准备。而机会，也总是青睐这些早有准备的人！

![表情包：打工人必备内卷表情包- 知乎](https://pic2.zhimg.com/80/v2-11ba1c8bf2e25ad491220e7ca9d5ecb9_1440w.webp)

年轻人们，让我们携手并进，一起在互联网的浪潮下狗刨吧~



**参考资料：**

C语言中文网：http://c.biancheng.net/redis/

Redis数据刷盘：https://cloud.tencent.com/developer/news/607786

