目录

> 1. 常见排序算法
> 2. 入门题目
>    1. 排序链表
>    2. 合并区间
> 3. 进阶题目：
>    1. 数组中第k个最大元素
>       2. 寻找两个正序数组的中位数
> 4. 小结



## 1. 常见排序算法

排序算法，常见的有插入排序【O(n^2)】，归并排序【O(nlogn)】，堆排序【O(nlogn)】以及快速排序【O(nlogn) -> O(n^2)】。

不熟悉排序算法的同学可以在 Visualgo 算法数据结构[`https://visualgo.net/en/sorting`]上学习，里面有视频和算法定义详解，文章里就不予赘述。

> 不熟悉力扣刷题顺序的可以看这篇文章：<a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484332&idx=1&sn=4fb1eb8580d93dc4a0dff69b713cce1f&chksm=ecac57a5dbdbdeb3dc0816468c3986f66cc483a6adcd7f80e6b208ff945fef5e8b538322aed8&token=609371871&lang=zh_CN#rd">学习路线导图</a>



## 2. 入门题目

#### 1）排序链表

力扣题目148：

![image-20230306091126288](img\lc-148.png)

这是 leetcode 的 148 题，题意很简单：对无序链表排序，进阶要求是：在 `O(nlogn)` 时间复杂度和常数级空间复杂度下，对链表进行排序。

其中，排序算法中**时间复杂度为 O(nlongn) 的算法有归并排序、堆排序和快排**，但快排的最差时间复杂度为 O(n^2) 在本题中不适用，堆排序和归并排序相比，归并排序是更适合链表的排序算法。



所以，我们用归并排序来实现这道题目，它的思想为**分治递归**，步骤为：

1. 将链表从中间节点一分为二，可以采用快慢指针的方式来实现；
2. 对两个子链表分别排序；
3. 将排好序的链表进行合并，即可得到完整的排序链表。

Go 代码如下：

``` go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 快慢指针，找到链表中间节点
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    next := slow.Next
    slow.Next = nil
    return mergeTwoList(sortList(head), sortList(next))
}

// mergeTwoList 合并两个子链表
func mergeTwoList(h1, h2 *ListNode) *ListNode {
    pre := new(ListNode)
    hair := pre
    for h1 != nil && h2 != nil {
        if h1.Val < h2.Val {
            pre.Next = h1
            h1 = h1.Next
        } else {
            pre.Next = h2
            h2 = h2.Next
        }
        pre = pre.Next
    }
    if h1 != nil {
        pre.Next = h1
    }
    if h2 != nil {
        pre.Next = h2
    }
    return hair.Next
}
```

代码用递归的方式写非常简单，效率也不错：

![image-20230306092258310](img\lc-148-submit.png)



#### 2）合并区间

力扣题目56：

![image-20230307085227415](img\lc-56.png)

大致题意为：对数组区间进行合并，合并之后不能有重叠区间。假设有区间 [1,3] 和 [2,6]，就要将其合并为 [1,6]。

这道题的解决思想是：将所有子数组的第 2 个元素与后面数组的第 1 个元素相比，如果有交叉，则找一个更大的元素作为新数组的第 2 个元素。即 `[1, 3] 和 [2, 6] `比较时，先对比 3 大于 2，可见两个数组是有交叉的，且 6 大于 3， 所以新数组的第 2 个元素为 6 => `[1, 6]`。

那我们如何保证第 1 个元素一定唯一呢？比如这 2 个数组：`[3, 5], [1, 8]`，我们在合并的时候就必须得判断第 1 个元素的大小，所以，为了简便，我们可以先对所有数组作一个排序，保证第 1 个元素更小的数组排在前面。在 Go 语言里面，可用简单的函数实现数组排序：

``` go
sort.Slice(arr, func(i, j int) bool {
	return arr[i][0] < arr[j][0] // 表示对子数组中下标为0的元素，进行升序排序
})
```

接下来，我们就可以写出完整代码：

``` go
import "sort"

func merge(intervals [][]int) [][]int {
    // 对数组的第1个元素（下标为0）进行升序排序
    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })
    var res [][]int
    for _, interval := range intervals{
        var size = len(res)
        // 如果当前的第1个元素小于前一个数组的第2个元素
        if size > 0 && interval[0] <= res[size-1][1] {
            res[size-1][1] = max(res[size-1][1], interval[1])
        }else{
            res = append(res, interval)
        }
    }
    return res
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
```

效率超过 94% 的 golang 提交：

![image-20230307090828110](img\lc-56-submit.png)



## 3. 进阶题目

### 1）求数组中的第 K 个最大元素

力扣题目215：

![image-20230307091243716](img\lc-215.png)

数组中的第 K 个最大元素，这道题是互联网大厂的常考面试题，我在字节、百度面试的时候都被考过，所以必须得熟练掌握。

题目意思很简单，就是寻找数组排序后的第 k 个最大的元素，比如 [1,2,3,4,5]，k=2，那第 2 个最大的元素就是 4。如果有题目中有相同元素，我们也不用理会，比如[5,5,4,3,1]，k=2，那第 2 个最大的元素就是 5，而非 4。

**进阶要求是**：设计并实现时间复杂度为 `O(n)` 的算法解决此问题。

这道题是排序算法中的经典题目，其中，如果我们不要求时间复杂度，可以直接用各类排序算法先对数组进行排序，再获取第 k 大的元素。但是，无论是面试，还是在日常使用中，我们都需要追求算法的时间复杂度最优解，而高效的排序算法有堆排序（nlogn）、归并排序（nlogn）和快速排序（nlogn -> n^2）。

那时间复杂度为 O(n) 该如何满足呢？

**快速排序**

答案在《算法导论》第三版的 9.2 有提到，`randomized_select` 算法是期望为线性时间的选择算法，它以快速排序为模型，思想仍是递归分治。但与快排不同的是，`randomized_select` 算法只处理划分的一边，我们用代码描述一下：

``` go
func findKthLargest(nums []int, k int) int {
    rand.Seed(time.Now().UnixNano())
    return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
    q := partition(a, l, r)
    if q == index {
        return a[q]
    } else if q < index {
        // 根据partition元素，只处理划分的一边
        return quickSelect(a, q + 1, r, index)
    }
    return quickSelect(a, l, q - 1, index)
}

func partition(a []int, l, r int) int {
    // 选取随机数作为边界元素
    i := rand.Intn(r-l+1) + l
    a[i], a[r] = a[r], a[i]
    
    x := a[r]
    i := l - 1
    for j := l; j < r; j++ {
        if a[j] <= x {
            i++
            a[i], a[j] = a[j], a[i]
        }
    }
    a[i+1], a[r] = a[r], a[i+1]
    return i + 1
}
```

>不熟悉快排代码的可以先看这篇文章：<a href="https://mp.weixin.qq.com/s?__biz=MzI5Nzk2MDgwNg==&mid=2247484207&idx=1&sn=77098e438cb574d8707fe44f9eb51298&chksm=ecac5726dbdbde30c17a8ae8abe7450dd7bdccd4d732c395c48bb923aaf7538bc93ddee525e5&scene=21#wechat_redirect">最大数（No179）</a>

提交结果：

![image-20230307092939507](img\lc-215-submit.png)

我们发现，快排的优化算法 randomized_select 时间复杂度超过了 99% 的用户。



**堆排序**

接下来我们再用堆排序的方式来实现，堆排序是利用**堆**这种数据结构所设计的一种排序算法。堆是一个近似完全二叉树的结构，并同时满足堆积的形式：即子节点的键值总是小于（或大于）其父节点的键值：

* 小顶堆：每个节点的值都小于或等于其子节点的值；
* 大顶堆：每个节点的值都大于或等于其子节点的值。

这道题目中是求最大的第 k 个数，所以我们构建大顶堆，每次找到最大数后，去除这个最大数（实现方式：将堆顶和堆尾元素互换，再将堆的大小减去一），然后求第 k-1 个数，代码如下：

``` go
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)+1-k; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, headSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < headSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < headSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, headSize)
	}
}
```

可以看出，堆排序的时间复杂度为 O(nlogn)，比快排的 `randomized_select` 更耗费时间一些：

![image-20230312102810219](img\lc-215-heap.png)



### 2）寻找两个正序数组的中位数

力扣题目4：

![image-20230312102920688](img\lc-4.png)

题目大意是找出两个有序数组 nums1 和 nums2 的中位数，并且要求时间复杂度为 `O（log(m+n)）`，可以假设两个数组不会同时为空。

这题，首先能想到的最简单方法就是将两个数组合并，然后取出中位数。但是合并数组的操作是 m+n 的复杂度了，不符合题意要求。看到 log 的时间复杂度，我们联想到**二分搜索**。

1. 两个数组的总数为奇数时，中间数为中位数；总数为偶数时，中间两个数的均值为中位数；
2. 已知数组的长度，要找两个数组合并以后的中位数，只需要二分搜索找到数组中间位置的 1 或 2 个元素即可；
3. 由于中间数的位置确定，所以两个数组中只需要**二分寻找一个数组**即可，为了复杂度最小，我们搜索长度较小的数组；
4. 最关键的：二分切割时应该怎么判断是否结束，我们假设有数组 A 和数组 B（长度分别为 m 和 n），先对数组 A 进行二分切割，那切割点 i，和 j【由于 i + j = (m+n)/2，所以 j = (m+n)/2 - i】需要满足，**A[i]、B[j] 左边的元素都要小于右边元素，即可保证 i 和 j 是中位数元素之一**。

图形表示如下：

![image-20230309111804177](img\image-20230309111804177.png)

1次二分拆分：

![image-20230309111918676](img\image-20230309111918676.png)

2次二分拆分：

![image-20230309112551595](img\image-20230309112551595.png)

再次拆分以后，left = 3， right = 3，数组 A 二分切割完毕。这时，数组 A 和 B 的切割下标分别为 i = left = 3, j = mid - left = 2。

![image-20230309114019032](img\image-20230309114019032.png)

而中位数的结果，只会在这 `i, j, i-1, j-1` 几个数里面产生：

1. 当数组总数为奇数时，中位数是左边分区（即 i-1 和 j-1）里面较大的那个；
2. 当数组总数为偶数时，中位数是左边分区最大数和右边分区最小数的均值。

除此之外，我们得考虑边界情况，比如 i, j 为 0 时，下标 i-1 或 j-1 不存在；当 i, j 为 m, n 时，下标 i 或 j 不存在。由于题目中说明了 m 和 n 不会同时为 0，所有 i 和 j 至少有一个存在。

Go 代码如下：

```go
func findMedianSortedArrays(nums1, nums2 []int) float64 {
   // 设置两个值，m+n
   m, n := len(nums1), len(nums2)
   if m > n {
      return findMedianSortedArrays(nums2, nums1)
   }
   left, right, mid := 0, m, (m+n+1)/2
   for left < right {
      i := (left + right + 1) / 2
      j := mid - i
      // 若数组A下标的左边分区数 > 数组B下标及右边分区时
      if nums1[i-1] > nums2[j] {
         right = i - 1
      } else {
         left = i
      }
   }
   i, j := left, mid-left
   var leftMax float64
   if i > 0 && j > 0 {
      leftMax = float64(max(nums1[i-1], nums2[j-1]))
   } else if i > 0 {
      leftMax = float64(nums1[i-1])
   } else {
      leftMax = float64(nums2[j-1])
   }
   if (m+n)%2 == 1 {
      return leftMax
   }
   var rightMin float64
   if i < m && j < n {
      rightMin = float64(min(nums1[i], nums2[j]))
   } else if i < m {
      rightMin = float64(nums1[i])
   } else {
      rightMin = float64(nums2[j])
   }
   return (rightMin+leftMax)/2
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a<b {
        return a
    }
    return b
}
```

提交结果：

![image-20230312103144870](img\lc-4-submit.png)



## 4. 小结

力扣排序类的几个经典题目都在上面了，它们在面试或笔试中遇到的概率非常之高。本人就曾在字节和百度的技术面试中遇到过 `数组中第k个最大元素` 和 `合并区间` 两道题，而 `链表排序` 在华为的机试题当中出现的频率也是很高的，除了最后一道题因为难度稍大考的有点少以外，其余几道题基本上 10 次面试中至少会出现 1 次。

而我们在练习算法题之时，也最好是按照它们的分类去针对性学习，以做到事半功倍之效。后续还会继续归纳下一分类的算法题目，敬请期待~

![公众号二维码](D:\envionment\coder\img\公众号二维码.jpg)