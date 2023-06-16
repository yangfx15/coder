目录

> 1. 函数
> 2. 迭代与生成器
> 3. 面向对象
> 4. 错误处理
> 5. 文件操作
> 6. 总结



## 1. 函数

### 1.1 函数定义与调用

和别的语言一样，函数是一个完整的代码块，用于实现特定的任务，Python 定义函数用 `def` 关键字。

下面是一个打印问候语的简单函数：

``` python
def hello():
    print("hello")
    
hello() #函数调用
```

和 for 或者其他代码块结构一样，函数的定义后跟冒汗，以及缩进单位相同的代码块，用来表示函数的内容。



### 1.2 函数传参

函数是一系列相同动作的组合，所以我们可以向函数传入参数，让它显示不同的结果：

``` python
def hello(name):
    print('hello, '+ name)

hello('zhangsan')
```

接收参数打印结果为：

> hello, zhangsan



#### 1）默认参数

Python 传参时，如果没有特殊说明，这些参数都是必传参数。这时如果你不想传参，那函数就会报错，如何避免报错呢？最简单的方法就是给参数设置一个默认参数：

``` python
def hello(name="world"):
    print('hello, '+ name)

hello()
```

如果没有传参，结果就会打印为：

> hello, world

注意：**默认参数可以有多个，但默认参数只能在必选参数的后面**，否则 Python 的解释器会报错。

如果一个函数里有多个默认参数，可以在调用时加上参数名：

``` python
def person(name, gender, age=6, city='Beijing'):
    print('name:', name)
    print('gender:', gender)
    print('age:', age)
    print('city:', city)
>>> person('yfx', 'M', 18)
>>> person('yfx', 'F', city='shenzhen')
```

如上所示，如果不加参数名，函数会按传参顺序使用这些参数。



#### 2）接收任意数量的参数

当函数的传参个数不确定时，Python 可以用一个 `*` 参数接收任意数量的位置参数，如：

``` python
def avg(first, *rest):
    return (first + sum(rest)) / (1 + len(rest))

avg(1, 2) # 1.5
avg(1, 2, 3, 4) # 2.5
```

此例中，rest 是由除了第一个参数外的所有参数组成的 tuple。

除了接收任意数量的 tuple，我们还可以添加任意数量的 dict，用 `**` 表示：

``` python
def person(name, age, **attrs):
    print(name, age, attrs)

person('zhangsan', 18, {'gender':'Male', 'job':'student'})
```



### 1.3 返回值

函数更多时候是对一段代码的封装，即输入一些参数，经过函数处理后返回一个或一组输出值。Python 中，`return` 语句可以返回函数调用的返回值。

``` python
def add(a, b):
    return a+b
add(1,2) #输出求和结果3
```



#### 1）返回元组

除了返回单个值以为，Python 还可以返回多个值，通过元组的方式呈现：

``` python
def point():
    x = 1
    y = 2
    return x, y
point() # (1,2)
```

和返回单个值一样，返回多个值时也是用 return 语句。



#### 2）返回字典

``` python
def build_persion(f_name, l_name, age=0):
    person = {'first': f_name, 'last': l_name}
    if age:
        person['age'] = age
    return person
```



## 2. 迭代与生成器

### 2.1 迭代

Python 中的字符串、列表、元组和字典在获取单个元素时都需要用到迭代：

**字符串迭代：**

``` python
for c in 'abc':
    print(c)
# 打印结果：
a
b
c
```

**字典迭代：**

``` python
person = {'name':'zhangsan', 'age':18, 'gender':'Male'}
for key in person:
    print(key)
打印结果：
name
age
gender
```

如果要对字典的 value 和 key-value 同时迭代，可以用 `values()` 和 `items()`：

``` python
person = {'name':'zhangsan', 'age':18, 'gender':'Male'}
for v in person.values():
    print(v)
#打印结果：
zhangsan
18
Male

#通过key-value迭代
for k,v in person.items():
    print(k,v)
#打印结果：
name zhangsan
age 18
gender Male
```



同时，如果想对列表进行**下标迭代**，可以用 Python 内置的 `enumrate` 函数处理：

``` python
for i,v in enumerate([4, 5, 6]):
    print(i,v)
#打印结果：
0 4
1 5
2 6
```



**判断是否可迭代**

Python 通过 `Iterable` 判断某个数据结构是否可迭代：

``` python
from collections.abc import Iterable
isinstance([1,2,3], Iterable)
# 结果：True
```



### 2.2 生成器

#### 1）列表生成式

`range(n)` 的用法我们已经在 Python入门篇（上）中说过了，这里我们复习一下：

``` python
arr = list(range(1,6))
print(arr)
# [1,2,3,4,5]
```

而**列表生成式**是在一行代码中生成列表的用法，比如我们想生成列表 arr 的平方：

``` python
arrs = [x*x for x in arr]
print(arrs)
# [1,4,9,16,25]
```

还可以通过两层循环遍历两个列表：

``` python
comp = [m + n for m in 'abc' for n in '12']
print(comp)
# ['a1', 'a2', 'b1', 'b2', 'c1', 'c2']
```



#### 2）生成器

生成器 `generator` 是 Python 的一个独有特性，是一种随取随用的数据结构。例如：

``` python
g = (x*x for x in range(10))
g
#<generator object <genexpr> at 0x1022ef630>
```

生成器的创建与列表生成式十分类似，不同的是列表生成式用 `[]` 生成，结果是一个列表；而后者用 `()` 生成， 结果是一个 `generator`。

用 `next()` 函数可以获取 `generator` 的下一个结果：

``` python
next(g) # 0
next(g) # 1
...
next(g) # 81
```

或者使用 `for` 循环：

``` python
for v in g:
    print(v)
```



#### 3）函数变为 `generator`

