###  select 机制

Go 语言里的 select 思想来源于网络 IO 模型中的 select，本质上也是 IO 多路复用。只是这里的 IO 是基于 channel 而不是基于网络套接字。

#### 1）Go select 特性

1. 每个 case 都必须是 channel；
2. 所有被发送的 channel 表达式都会被求值；
3. 如果有多个 case 可以执行，select 会随机公平地选出一个执行；
4. 如果没有 default 语句，select 将阻塞，直到某个 channel 可以运行；
5. 当 case 中的通道关闭时，case 每次都会读到 chan 对应数据类型的零值，导致死循环；
6. 空 select 语句会死锁。

#### 2）特性验证

1. **select 已关闭通道和空通道场景，case closed/nil channel**

```go
func main() {
	c1, c2, c3 := make(chan bool), make(chan bool), make(chan bool)
	go func() {
		for {
			select {

			// 保证c1一定不会关闭，否则会死循环此case
			case <-c1:
				fmt.Println("case c1")

			// c2可以防止出现c1死循环的情况
            // 如果c2关闭了(ok==false)，将其置为nil，下次就会忽略此case
			case _, ok := <-c2:
				if !ok {
					fmt.Println("c2 is closed")
					c2 = nil
				}

			// 如果c3已关闭，则panic(不能往已关闭的channel里写数据)
			// 如果c3为nil，则ignore该case
			case c3 <- true:
				fmt.Println("case c3")

            case v <- c4:
                fmt.Println(v)
			}       
            
		}
	}()
	time.Sleep(10 * time.Second)
}
```

当 channel 关闭以后，case <- chan 会接收该通信对应数据类型的零值，所以会死循环。



1. **带 default 语句实现非阻塞读写**

```go
select {
    case <- c1:
    fmt.Println(":case c3")
    // 当c1没有消息时，不会一直阻塞，而是进入default
    default:
    fmt.Println(":select default")
}
```

常见误区：一些同学使用 select 时，习惯先将 default 写好，然后外层加上 for 循环【会出现死锁】。

Go 语言的 select 和 switch 在这里不太一样，switch 中一般会带有 default 判断分支，但 select 使用时，外层的 for 循环和 default 基本不会同时出现，否则会发生死锁。



1. **select 实现定时任务**

```go
func main() {
    done := make(chan bool)
    var selectTest = func() {
       for {
          select {
             case <-time.After(1 * time.Second):
             fmt.Println("second exec")
             case <-done:
             fmt.Println("over")
         }
      }
   }
    go selectTest()

    time.Sleep(3 * time.Second)
    done <- true
    time.Sleep(500 * time.Microsecond)
}
```

运行结果：

second exec

second exec

over

注意，如果定时器的另外 case 分支是上面已关闭 channel 场景，可能会出现异常：

```go
func main() {
	done := make(chan bool)
	t := time.Now()
	var selectTest = func() {
		for {
			select {
			case <-time.After(100 * time.Microsecond):
				fmt.Println(time.Since(t), " time.After exec, return!")
				return
			case <-done:
				fmt.Println("over")
			}
		}
	}
	close(done)
	go selectTest()

	time.Sleep(2 * time.Second)
}
```

不妨猜一下这段代码会输出什么，答案是：

...

over

over

over

601.3938ms  time.After exec, return!

done 已经被关闭，所以会死循环此 case 分支，定时器 case 100 ms 执行一次，在第 6 次执行时进入此 case 程序退出。

为什么会这样呢？请看下文。



1. **多个 case 满足读写条件**

如果多个 case 满足读写，select 会随机选择一个语句执行

```go
func main() {
    done := make(chan int, 1024)
    tick := time.NewTicker(time.Second)
    var selectTest = func() {
       for i := 0; i < 10; i++ {
          select {
             case done <- i:
             fmt.Println(i, ", done exec")
             case <-tick.C:
             fmt.Println(i, ", time.After exec")
         }
          time.Sleep(500 * time.Millisecond)
      }
   }
    go selectTest()

    time.Sleep(5 * time.Second)
}
```

执行结果：

0 , done exec

1 , done exec

2 , time.After exec

3 , done exec

4 , time.After exec

5 , done exec

6 , done exec

7 , done exec

8 , time.After exec

9 , done exec

可以发现，原本写入 done 通道的 2、4 和 8 不见了，说明在循环的过程中，tick.C 是执行了的。所以，当多个 case 同时满足时，select 会随机选择一个执行，也就解释了上面 **3.select 实现定时任务** 的疑问了。