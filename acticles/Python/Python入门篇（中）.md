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

和 for 或者其他代码块结构一样，函数的定义后跟冒号，以及缩进单位相同的代码块，用来表示函数的内容。



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

### 2.1 可迭代对象

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

#### 1）列表生成式（range）

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



#### 2）生成器（generator）

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

普通函数会从定义开始直到 `return` 或者执行完所有代码后结束。

但在 Python 中，如果一个函数定义中包含 `yield` 关键字，那么这个函数就是一个 `generator`，调用这种函数就会返回一个 `generator`：

``` python
def add(num):
    n = 1
    sum = 0
    while n<=num:
        sum += n
        yield sum
        n += 1
	return 'success'
```

注意，当我们调用这个 `generator` 函数时，如：add(100)，并非立即得出 1+2+...+100 的结果。而是依次返回数字：1、3、6、10...5050：

``` python
f = add(100)
next(f) # 1
next(f) # 3
for sum in f:
    print(sum) #1、3、6、10...5050
```

同时，如果用 for 循环获取 `generator` 的值时，无法获取 return 的返回，此时我们可以捕获`StopIteration`错误，返回值就包含在 value 里面：

``` python
f = add(100)
while True:
    try:
        sum = next(f)
        print('sum:', sum)
    except StopIteration as e:
        print('Generator return value:', e.value)
        break
```

执行结果：

> sum: 1
> sum: 3
> ...
> sum:5050
> Generator return value: success



### 2.3 迭代器

在 Python 中，可以被 `next()` 函数调用并不断返回下一个值的对象被称为迭代器：`Iterator`。

和可迭代对象（Iterable）不同，**列表、元组、字典、集合、字符串**等对象虽然是可迭代的，但是它们没法直接使用 `next()` 获取下一个值，因此它们不是迭代器。示例：

``` python
from collections.abc import Iterable, Iterator
isinstance([1,2,3], Iterable) # True，列表是一个可迭代对象
isinstance([1,2,3], Iterator) # False，列表不是一个迭代器
```

这时可能有人就蒙逼了，迭代器和可迭代对象有什么本质区别，Python 这两个迭代对象如何区分呢？

我们只需要关注：迭代器对象表示的是一个数据流，可以源源不断地取数据，这个数据流的长度我们在遍历前是不确定的。而可迭代对象可以通过 `len()` 来获取对象的长度，它们的长度在遍历之前就是确定的。



把可迭代对象变成迭代器，可使用 `iter()` 方法：

``` python
isinstance(iter([1,2,3]), Iterator) # True
```



## 3. 面向对象

面向对象是一种设计思想，大部分高级编程语言都采用了这种思想来管理代码结构，Python 也不例外。

对象是一种抽象概念，为了方便描述，面向对象中衍生了两个最重要的实体概念：类（Class）和实例（Instance）。

其中，**类**泛指具有相同特征的集合。比如：动物园中的所有动物，可以抽象为动物类。**实例**是根据类创建出的一个个具体对象，每个对象都有着相同的属性和行为，即这个类别共同拥有的属性和行为。

比如，动物拥有自己的类别，名字，叫声等属性；拥有自己的跑动，进食方式等。



### 3.1 类初始化 \__init__

Python 中用 class 来声明类，比如，定义一个动物类：

``` python
class Animal(object):
    def __init__(self, kind, name, sound):
        self.name = kind
        self.kind = name
        self.sound = sound
```

在 Python 中定义类时，必须指定继承一个父类，如果没有合适的父类，就用 `object` 祖先类。



由于类是所有实例的模板，所以我们在创建实例的时候必须要传入类的基本属性，把它们放入 `__init__` 方法里。并且，把传入的属性绑定到 `self` ，即当前实例上。

比如定义一只小狗旺财：

``` python
wc = Animal('旺财', '狗', '汪汪')
print(wc.name) # 旺财
```



### 3.2 基础属性

#### 1）实例的额外属性

和静态语言不同，Python 中同一个对象的实例可能有不同的属性，如：

``` python
kitty = Animal('kitty', '猫', '喵喵')
kitty.age = 1
print(kitty.age) # 1
print(wc.age) # 报错，旺财不存在该属性
```



#### 2）私有属性

Python 中可以在属性名称之前加两个下划线，声明为私有属性，私有属性不能在类外部访问，如：

``` python
class Student(object):
    def __init__(self, name, score):
        self.__name = name
        self.__score = score
    def print_score(self):
        print('%s: %s' % (self.__name, self.__score))
```

当我们创建实例时，只能通过类里面的方法打印学生的成绩：

``` python
s = Student('zhangsan', 95)
s.print_score() # zhangsan: 95
s.__score # 'Student' object has no attribute '__score'
```



#### 3）获取所有属性&方法

Python 中，我们可以用 `dir()` 来获取类所有的属性和方法，如：

``` python
dir(Animal)
['__class__', '__delattr__', '__dict__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__gt__', '__hash__', '__init__', '__init_subclass__', '__le__', '__lt__', '__module__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__']
```

可以看到有很多我们没主动定义的 `__xxx__`方法，这些方法在 Python 中是有特殊用途的。比如 `__str__` 返回一个实例的描述，我们可以调用 `str()` 方法，Python 会自动调用 `__str__` 方法：

``` python
print(str(Animal))
"<class '__main__.Animal'>"
print(Animal) # 相当于调用 Animal.__str__()
<class '__main__.Animal'>
```



**合理打印对象数据**

通过以上可知，一个类通过继承祖先 `object` 类，可以获取很多初始化方法，比如上述打印类的 `__str__` 方法。所以，我们还可以重写这些方法：

``` python
class Animal(object):
    def __init__(self, kind, name, sound):
        self.name = kind
        self.kind = name
        self.sound = sound
    def __str__(self):
        return 'Animal object(kind:%s) name:%s, sound:%s' % (self.kind, self.name, self.sound)
```

此时，直接打印对象数据时，就变成了：

``` python
>>> k = Animal('kitty', 'cat', '喵喵')
>>> print(k)
Animal object(kind:kitty) name:cat, sound:喵喵
```



#### 4）测试&新增属性

Python 提供了三个方法，可以操控一个类或对象的属性：

``` python
>>> k = Animal('kitty', '猫', '喵喵')
>>> hasattr(k, 'name') # 对象是否有name属性
True
>>> getattr(k, 'name') # 获取属性name
'kitty'
>>> hasattr(k, 'age')
False
>>> setattr(k, 'age', 1) # 设置属性age默认值为1
>>> getattr(k, 'age')
1
```

其中 `hasattr` 判断对象或类是否存在某个属性，`getattr` 获取对象或类的属性值，当获取不存在的属性时 Python 会报错；而 `setattr` 则为对象或类添加属性，并为其设置默认值。



### 3.3 高级属性

#### 1）\__slots__限制类属性

Python 作为一门动态语言，可以在定义好 class 之后给实例绑定任意的属性和方法，如：

``` python
k = Animal('kitty', '猫', '喵喵')
k.age = 1 # 指定猫咪的年龄为1岁
```

但是，假设我定义了一个类，任何人都可以添加类属性，这样势必会带来管理困难的问题。那我们如何让类保持初始的属性字段呢？

Python 提供了 `__slots__` 变量，来限制 class 类的实例属性：

``` python
class Animal(object):
    __slots__ = ('kind', 'name', 'sound')
    def __init__(self, kind, name, sound):
        self.name = kind
        self.kind = name
        self.sound = sound
```

这时，如果我们还为实例添加额外的属性，就会报错：

``` python
k = Animal('kitty', '猫', '喵喵')
k.age = 1 # AttributeError: 'Animal' object has no attribute 'age'
```



#### 2）装饰器@property

为了方便调用，Python 提供了装饰器 `@property`，可以将类的方法变成属性来调用：

``` python
class Animal(object):
    @property
    def kind(self):
        #为了和方法名做区分，在属性前面加一个下划线
        return self._kind
    @kind.setter
    def kind(self, value):
        if not isinstance(value, str):
            raise ValueError('kind must be a str!')
        self._kind = value
```

调用 @setter 和 @property 装饰器的方法：

``` python
k = Animal()
k.kind = 'cat' # OK，相当于调用 k.set_kind('cat')
k.kind # cat，相当于调用 k.get_kind()
```

> 特别注意：为了防止属性的方法名和实例变量重名，我们在设置变量时，可以在属性前面加一个下划线



### 3.4 多重继承

不同于 Java 等编程语言只能单继承，Python 一个类可以有多个父类，即多重继承。

多重继承在自然界中很常见，归因于不同的分类方式。比如动物根据类别可以分为哺乳类和鸟类，也可以根据运动方式分为天上飞的，水里游的，陆上走的。

我们可以声明一个跑动类动物：

``` python
class Runnable(Animal):
    pass
```

一个哺乳类动物：

``` python
class Mammal(Animal):
    pass
```

一个跑动的哺乳类动物：

``` python
class Cat(Runnable, Mammal):
    pass
```

继承了多个父类的子类，可以同时获得多个父类的功能。



