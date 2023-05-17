#!/usr/bin/env python3
# -*- coding: utf-8 -*-

<<<<<<< HEAD
def prod(L):
    def fn(x,y):
        return x*y
    return reduce(fn, L)
print('3 * 5 * 7 * 9 =', prod([3, 5, 7, 9]))
if prod([3, 5, 7, 9]) == 945:
    print('测试成功!')
else:
    print('测试失败!')
=======
from enum import Enum
@unique
class Weekday(Enum):
    Sun = 0
    Mon = 1
    Tue = 2
    Wed = 3
    Thu = 4
    Fri = 5
    Sat = 6

day1 = Weekday.Mon
print(day1)

Weekday.Mon
>>>>>>> refs/remotes/origin/main
