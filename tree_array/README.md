树状数组(Binary Indexed Tree)
===

**树状数组**是一个查询和修改复杂度都为log(n)的数据结构。主要用于查询任意两位之间的所有元素之和，但是每次只能修改一个元素的值；经过简单修改可以在log(n)的复杂度下进行范围修改，但是这时只能查询其中一个元素的值(如果加入多个辅助数组则可以实现区间修改与区间查询).

主要用于区间求和的情况。

## lowbit(x)
计算低位bit位

## bitindex(x)
计算下一个需要改变的数组下标

## Add
给某位加上某个值

## GetSum
获得前N个(包含)值的和

# 用例
* 总分100(N)，统计某个分数区间的人数(0-59)。