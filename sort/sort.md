## 选择排序

### 算法特性

- 时间复杂度为O(n²)： 非自适应排序: 遇到有序数据时插入操作不会提前终止
  次，求和得到
- 空间复杂度O(1): 原地排序
- 非稳定排序: [4,4,3,1,5,2]

## 冒泡排序

### 算法特性

- 时间复杂度为O(n²)： 使用自适应排序，引入flag后，最佳时间复杂度为O(n)
- 空间复杂度为O(1): 原地排序
- 稳定排序

## 插入排序

### 算法特性

- 时间复杂度为O(n²)： 自适应排序，最差情况为O(n²)。当输入数组完全有序时，插入排序达到最佳时间复杂度O(n)
- 空间复杂度O(1)：原地排序
- 稳定排序

### 优势

虽然冒泡排序、选择排序和插入排序的时间复杂度都为 O(n²)，但在实际情况中，插入排序的使用频率显著高于冒泡排序和选择排序，主要有以下原因。

- 冒泡排序基于元素交换实现，需要借助一个临时变量，共涉及 3 个单元操作；插入排序基于元素赋值实现，仅需 1 个单元操作。因此，冒泡排序的计算开销通常比插入排序更高。
- 选择排序在任何情况下的时间复杂度都为O(n²) 。如果给定一组部分有序的数据，插入排序通常比选择排序效率更高。
- 选择排序不稳定，无法应用于多级排序。

## 快速排序

### 算法特性

- 时间复杂度O(nlogn)非自适应排序：在平均情况下，哨兵划分的递归层数为logn，每层中的总循环为n，总体使用O(nlogn)。最差情况下，每轮哨兵划分操作都将长度为n的数组划分为长度为0和n-1的两个子数组，此时递归层数达到n，每层中的循环数为n，总体使用O(n²)时间
- 空间复杂度为O(n): 原地排序，在输入数组完全倒序的情况下，达到最差递归深度n,使用O(n)栈帧空间
- 非稳定排序： 在哨兵划分最后一步，基准数可能会被交换至相等元素右侧
