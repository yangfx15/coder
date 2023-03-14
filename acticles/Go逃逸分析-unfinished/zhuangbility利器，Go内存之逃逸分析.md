#### 1）堆与栈内存

虚拟内存里有两块比较重要的地址空间，分别为堆和栈空间。我们在函数中申请对象时，如果分配在栈上，函数执行结束后编译器会自动回收；如果分配在堆中，则在函数结束后的某个时间点由编译器和垃圾收集器进行回收。

在栈上分配和回收内存的开销很低，只需要 2 个指令：PUSH 和 POP。PUSH 将数据压入栈中，POP 释放空间，消耗的仅是将数据拷贝到内存的时间。

而在堆上分配内存时，不仅分配的时候慢，而且垃圾回收的时候也比较费劲，比如说 Go 在 1.8 以后就用到了三色标记法+混合写屏障的技术来做垃圾回收。总体来看，堆内存分配比栈内存分配导致的开销要大很多。



#### 2）逃逸分析

在 C 语言中，可以使用 malloc 和 free 手动在堆上分配和回收内存。Go 语言中，堆内存是编译器和垃圾收集器自动管理的，无需用户程序来指定。那么，Go 编译器如何知道某个变量需要分配在堆，还是栈上呢？

编译器决定内存分配位置的方式，就称之为逃逸分析。逃逸分析由编译器完成，作用于编译阶段。



#### 3）四种逃逸场景

**3.1 指针逃逸**

指针逃逸很容易理解，我们在函数中创建一个对象时，对象的生命周期随着函数结束而结束，这时候对象的内存就分配在栈上。



而如果返回了一个对象的指针，这种情况下，函数虽然退出了，但指针还在，对象的内存不能随着函数结束而回收，因此只能分配在堆上。

```go
package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}
// 要想不发生逃逸，返回 User 对象即可。
func GetUserInfo() *User {
	return &User{
		ID: 666666,
		Name: "sim lou",
		Avatar: "https://www.baidu.com/avatar/666666",
	}
}

func main()  {
	u := GetUserInfo()
	println(u.Name)
}
```

上面例子中，如果返回的是 User 对象，而非对象指针 *User，那么它就是一个局部变量，会分配在栈上；反之，指针作为引用，在 main 函数中还会继续使用，因此内存只能分配到堆上。



我们可以用编译器命令 go build -gcflags -m main.go 来查看变量逃逸的情况：

![img](https://cdn.nlark.com/yuque/0/2022/png/12709510/1668130326542-c5bd5539-d861-42c2-a01d-9e5ce45ea5e8.png)

&User{...} escapes to heap 即表示对象逃逸到堆上了。



**3.2 interface{} 动态类型逃逸**

在 Go 语言中，空接口即 interface{} 可以表示任意的类型，如果函数参数为 interface{}，编译期间很难确定其参数的具体类型，也会发生逃逸。



比如 Println 函数，入参是一个 interface{} 空类型：

```go
func Println(a ...interface{}) (n int, err error)
```



这时，返回的是一个 User 对象，也会发生对象逃逸，但逃逸节点是 fmt.Println 函数使用时

```go
func GetUserInfo() User {
	return User{
		ID: 666666,
		Name: "sim lou",
		Avatar: "https://www.baidu.com/avatar/666666",
	}
}

func main()  {
	u := GetUserInfo()
	fmt.Println(u.Name) // 对象发生逃逸
}
```



**3.3 栈空间不足**

操作系统对内核线程使用的栈空间是有大小限制的，64 位 Linux 系统上通常是 8 MB。可以使用 ulimit -a 命令查看机器上栈允许占用的内存的大小。

```go
root@cvm_172_16_10_34:~ # ulimit -a
-s: stack size (kbytes)             8192
-u: processes                       655360
-n: file descriptors                655360
```



因为栈空间通常比较小，因此递归函数实现不当时，容易导致栈溢出。

对于 Go 语言来说，运行时(runtime) 尝试在 goroutine 需要的时候动态地分配栈空间，goroutine 的初始栈大小为 2 KB。当 goroutine 被调度时，会绑定内核线程执行，栈空间大小也不会超过操作系统的限制。

对 Go 编译器而言，超过一定大小的局部变量将逃逸到堆上，不同的 Go 版本的大小限制可能不一样。我们来做一个实验（注意，分配 int[] 时，int 占 8 字节，所以 8192 个 int 就是 64 KB）：

```go
package main

import "math/rand"

func generate8191() {
	nums := make([]int, 8192) // <= 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generate8192() {
	nums := make([]int, 8193) // > 64KB
	for i := 0; i < 8193; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

func main() {
	generate8191()
	generate8192()
	generate(1)
}
```



编译结果如下：

![img](https://cdn.nlark.com/yuque/0/2022/png/12709510/1668131153785-d6a3f64c-e61c-4a73-96b1-8f0ad9422f1b.png)



可以发现，make([]int, 8192) 没有发生逃逸，make([]int, 8193) 和 make([]int, n) 逃逸到堆上。

也就是说，当切片占用内存超过一定大小，或无法确定当前切片长度时，对象占用内存将在堆上分配。



**3.4 闭包**

一个函数和对其周围状态（lexical environment，词法环境）的引用捆绑在一起（或者说函数被引用包围），这样的组合就是闭包（closure）。也就是说，闭包让你可以在一个内层函数中访问到其外层函数的作用域。

 

— [闭包](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Closures)

Go 语言中，当使用闭包函数时，也会发生内存逃逸。看一则示例代码：

```go
package main

func Increase() func() int {
    n := 0
    return func() int {
        n++
        return n
    }
}
func main() {
    in := Increase()
    fmt.Println(in()) // 1
    fmt.Println(in()) // 2
}
```

Increase() 返回值是一个闭包函数，该闭包函数访问了外部变量 n，那变量 n 将会一直存在，直到 in 被销毁。很显然，变量 n 占用的内存不能随着函数 Increase() 的退出而回收，因此将会逃逸到堆上。



#### 4）利用逃逸分析提升性能

**4.1 传值VS传指针**

传值会拷贝整个对象，而传指针只会拷贝指针地址，指向的对象是同一个。传指针可以减少值的拷贝，但是会导致内存分配逃逸到堆中，增加垃圾回收(GC)的负担。在对象频繁创建和删除的场景下，传递指针导致的 GC 开销可能会严重影响性能。

一般情况下，对于需要修改原对象值，或占用内存比较大的结构体，选择传指针。对于只读的占用内存较小的结构体，直接传值能够获得更好的性能。