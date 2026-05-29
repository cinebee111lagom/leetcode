# Binary Search 二分查找 - LeetCode Deep Dive

## 目录
- [Problem 4: Median of Two Sorted Arrays](#problem-4-median-of-two-sorted-arrays-寻找两个正序数组的中位数)
- [Problem 34: Find First and Last Position](#problem-34-find-first-and-last-position-in-sorted-array-在排序数组中查找元素的第一个和最后一个位置)
- [Problem 35: Search Insert Position](#problem-35-search-insert-position-搜索插入位置)
- [Problem 74: Search 2D Matrix](#problem-74-search-a-2d-matrix-搜索二维矩阵)

---

## Problem 4: Median of Two Sorted Arrays 寻找两个正序数组的中位数

### 1. 题目核心与隐藏考点

**核心本质**: 归并两个有序数组的变形，通过二分查找在较短数组上找分割点，时间复杂度 O(log(m+n))。

**隐藏考点**:
- 为什么需要 O(log(m+n)) 而不能用简单的归并？
- 分割点的边界条件
- 奇偶长度不同处理

```
二分查找分割点图解:

nums1 = [1, 3], nums2 = [2]

寻找分割点 i, j:
- i 将 nums1 分为左半 [0..i-1] 和右半 [i..]
- j 将 nums2 分为左半 [0..j-1] 和右半 [j..]

条件:
- i + j = (m + n + 1) / 2 (左半元素个数)
- nums1[i-1] <= nums2[j] (左半最大值 <= 右半最小值)
- nums2[j-1] <= nums1[i] (对称条件)

中位数:
- 如果总长度是奇数: max(left1, left2)
- 如果总长度是偶数: (max(left1, left2) + min(right1, right2)) / 2

分割示例:
  nums1 = [1, 3], nums2 = [2]
  m=2, n=1, totalLen=3, halfLen=2

  i=1 时: j = 2 - 1 = 1
  nums1[0]=1, nums2[0]=2
  1 <= 2 ✓
  中位数 = max(1, 2) = 2
```

---

### 2. 思路演进

#### 解法一：归并 O(m+n) - 不是最优
```go
merged := make([]int, m+n)
i, j, k := 0, 0, 0
for i < m && j < n {
    if nums1[i] < nums2[j] {
        merged[k++] = nums1[i++]
    } else {
        merged[k++] = nums2[j++]
    }
}
// 找中位数
```

#### 解法二：二分查找 O(log(min(m,n))) - 最优

```
核心思想：
1. 在较短的数组上二分
2. 找到分割点使得:
   - 左半元素个数 = (m+n+1)/2
   - 左半最大值 <= 右半最小值
3. 根据奇偶计算中位数

决策过程:
  i=1, j=2-1=1, 检查 nums1[0] <= nums2[1]? 
  i=2, j=2-2=0, 检查 nums2[0] <= nums1[2]?
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 归并 | O(m+n) | O(m+n) | 简单 | 非最优 |
| 二分查找 | O(log(min(m,n))) | O(1) | 最优 | 边界复杂 |

---

### 4. 工业级 Go 源码与详细注释

```go
package binary_search

// FindMedianSortedArrays 寻找两个正序数组的中位数
//
// 核心思想：二分查找
//
// 为什么用二分？
// - 需要 O(log(m+n)) 时间复杂度
// - 普通的归并需要 O(m+n)，不够快
// - 二分可以在一个数组上找分割点，将复杂度降到 O(log(min(m,n)))
//
// 算法步骤：
// 1. 确保 nums1 是较短的数组（便于二分）
// 2. 在 nums1 上二分查找分割点 i
// 3. 计算 nums2 上的分割点 j = (m+n+1)/2 - i
// 4. 检查是否满足条件：nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
// 5. 根据奇偶计算中位数
//
// 分割点的含义：
// - 左半包含 m+n+1)/2 个元素
// - 如果 totalLen 是奇数，中位数就是左半最大值
// - 如果 totalLen 是偶数，中位数是左半最大值和右半最小值的平均值
//
// 时间复杂度：O(log(min(m,n)))
// 空间复杂度：O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // 确保 nums1 是较短的数组，便于二分
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)
    totalLen := m + n
    halfLen := (totalLen + 1) / 2 // 左半元素个数

    // 在 nums1 上二分
    left, right := 0, m

    for left <= right {
        // nums1 的分割点
        i := (left + right) / 2
        // nums2 的分割点
        j := halfLen - i

        // 计算左半最大值和右半最小值
        // 考虑四种边界情况

        // nums1 左半最大值（i==0 时为负无穷）
        var maxLeft1 int
        if i == 0 {
            maxLeft1 = nums2[j-1]
        } else if j == 0 {
            maxLeft1 = nums1[i-1]
        } else if nums1[i-1] > nums2[j-1] {
            maxLeft1 = nums1[i-1]
        } else {
            maxLeft1 = nums2[j-1]
        }

        // nums1 右半最小值（i==m 时为正无穷）
        var minRight1 int
        if i == m {
            minRight1 = nums2[j]
        } else if j == n {
            minRight1 = nums1[i]
        } else if nums1[i] < nums2[j] {
            minRight1 = nums1[i]
        } else {
            minRight1 = nums2[j]
        }

        // 检查是否找到正确的分割点
        if maxLeft1 <= minRight1 {
            // 找到分割点
            if totalLen%2 == 1 {
                return float64(maxLeft1)
            }
            // 计算右半最小值
            var maxLeft2 int
            if j == 0 {
                maxLeft2 = nums1[i-1]
            } else if i == 0 {
                maxLeft2 = nums2[j-1]
            } else if nums1[i-1] > nums2[j-1] {
                maxLeft2 = nums1[i-1]
            } else {
                maxLeft2 = nums2[j-1]
            }
            return float64(maxLeft1+minRight1) / 2.0
        }

        // 调整二分边界
        if nums1[i-1] < nums2[j-1] {
            left = i + 1
        } else {
            right = i - 1
        }
    }

    return 0
}
```

---

## Problem 34: Find First and Last Position in Sorted Array 在排序数组中查找元素的第一个和最后一个位置

### 1. 题目核心与隐藏考点

**核心本质**: 二分查找的变形，分别找左边界和右边界。

**隐藏考点**:
- 如何找左边界？找第一个 >= target 的位置
- 如何找右边界？找第一个 > target 的位置
- 边界条件处理

```
二分查找左右边界图解:

nums = [5, 7, 7, 8, 8, 10], target = 8

找左边界 (第一个 >= 8):
  left=0, right=5, mid=2
  nums[2]=7 < 8, left=mid+1=3
  left=3, right=5, mid=4
  nums[4]=8 >= 8, right=mid=4
  left=3, right=4, mid=3
  nums[3]=8 >= 8, right=mid=3
  left=3, right=3, mid=3
  nums[3]=8 == target, 返回 left=3

找右边界 (第一个 > 8):
  left=0, right=5, mid=2
  nums[2]=7 <= 8, left=mid+1=3
  left=3, right=5, mid=4
  nums[4]=8 <= 8, left=mid+1=5
  left=5, right=5, mid=5
  nums[5]=10 > 8, right=mid-1=4
  left=5 > right=4, 结束
  右边界 = left-1 = 4

结果: [3, 4]
```

---

### 2. 工业级 Go 源码与详细注释

```go
// SearchRange 在排序数组中查找元素的第一个和最后一个位置
//
// 核心思想：二分查找找左右边界
//
// 为什么需要两次二分？
// - 普通二分找到目标就停止，无法找到边界
// - 找左边界时，需要继续向左搜索
// - 找右边界时，需要继续向右搜索
//
// 左边界查找：
// - 当 nums[mid] < target 时，left = mid + 1
// - 当 nums[mid] >= target 时，right = mid
// - 结束条件：left == right，此时 left 就是第一个 >= target 的位置
//
// 右边界查找：
// - 当 nums[mid] <= target 时，left = mid + 1
// - 当 nums[mid] > target 时，right = mid
// - 结束条件：left == right，此时 left 就是第一个 > target 的位置
// - 右边界 = left - 1
//
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func SearchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    // 找左边界
    left := findLeftBound(nums, target)
    if left == -1 {
        return []int{-1, -1}
    }

    // 找右边界
    right := findRightBound(nums, target)
    return []int{left, right}
}

// findLeftBound 找第一个 >= target 的位置
func findLeftBound(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right-left)/2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    if nums[left] == target {
        return left
    }
    return -1
}

// findRightBound 找第一个 > target 的位置
func findRightBound(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left < right {
        // 注意这里 +1 是为了避免死循环
        mid := left + (right-left)/2 + 1
        if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid
        }
    }

    return left
}
```

---

## Problem 35: Search Insert Position 搜索插入位置

### 1. 题目核心与隐藏考点

**核心本质**: 二分查找找插入位置，等价于找第一个 >= target 的位置。

**隐藏考点**:
- 为什么返回值就是插入位置？
- 边界条件：目标插入到开头或结尾

```
二分查找图解:

nums = [1, 3, 5, 6], target = 2

Step 1: left=0, right=3, mid=1
        nums[1]=3 > 2, right=mid-1=1

Step 2: left=0, right=1, mid=0
        nums[0]=1 < 2, left=mid+1=1

Step 3: left=1, right=1
        返回 left=1，即插入位置

验证: [1, 2, 3, 5, 6] 中 2 在索引 1
```

---

### 2. 工业级 Go 源码与详细注释

```go
// SearchInsert 搜索插入位置
//
// 核心思想：二分查找
//
// 查找第一个 >= target 的位置，这个位置就是插入位置
// - 如果 target 存在，返回其索引
// - 如果 target 不存在，返回第一个大于 target 的位置
//
// 为什么这样正确？
// - 二分查找找到的位置满足：所有 < target 的元素都在左边
// - 这正是插入位置的定义
//
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func SearchInsert(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    // left 就是插入位置
    // 因为当循环结束时，left > right
    // 且所有 < target 的元素都在 left 左侧
    return left
}
```

---

## Problem 74: Search a 2D Matrix 搜索二维矩阵

### 1. 题目核心与隐藏考点

**核心本质**: 两次二分：先按行二分找到目标行，再在行内二分找目标。

**隐藏考点**:
- 每行第一个元素大于上一行最后一个元素
- 如何将二维转化为一维？

```
二维矩阵搜索图解:

matrix = [
  [1, 3, 5, 7],
  [10, 11, 16, 20],
  [23, 30, 34, 60]
], target = 3

Step 1: 按行二分找目标行
        rows = 3, left=0, right=2, mid=1
        matrix[1][0]=10 > 3? 不是
        matrix[1][0]=10 > 3? 否，搜索上半部分
        right=mid-1=0

Step 2: left=0, right=0, mid=0
        matrix[0][0]=1 < 3? 是，left=mid+1=1
        循环结束，row=0

Step 3: 在第 0 行二分查找 3
        left=0, right=3, mid=1
        matrix[0][1]=3 == target
        返回 true

结果: true
```

---

### 2. 工业级 Go 源码与详细注释

```go
// SearchMatrix 搜索二维矩阵
//
// 核心思想：两次二分
//
// 为什么可以两次二分？
// 1. 矩阵的特性：每行第一个元素大于上一行最后一个元素
// 2. 因此整个矩阵可以看作一个有序数组
// 3. 可以先按行二分找到目标可能所在的行
// 4. 再在那一行内二分找目标
//
// 为什么按行二分是正确的？
// - 如果 target 小于某行第一个元素，那么 target 只可能在上方的行
// - 如果 target 大于某行最后一个元素，那么 target 只可能在下方的行
// - 这个二分可以将搜索范围缩小到一行
//
// 时间复杂度：O(log(m) + log(n)) = O(log(mn))
// 空间复杂度：O(1)
func SearchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    m, n := len(matrix), len(matrix[0])

    // 第一次二分：按行找目标行
    top, bottom := 0, m-1
    for top < bottom {
        mid := (top + bottom) / 2
        if matrix[mid][n-1] < target {
            // 目标在 mid 行下方
            top = mid + 1
        } else if matrix[mid][0] > target {
            // 目标在 mid 行上方
            bottom = mid - 1
        } else {
            // mid 行可能包含目标
            break
        }
    }

    // 确定目标行
    row := (top + bottom) / 2

    // 第二次二分：在目标行内找目标
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [33. Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | 二分查找变形 | Medium |
| [81. Search in Rotated Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/) | 二分 + 去重 | Medium |

---

## 二分查找总结

### 常见模式

| 模式 | 题目 | 关键点 |
| :--- | :--- | :--- |
| 标准二分 | 35 | 找目标位置 |
| 左右边界 | 34 | 两次二分 |
| 二维展开 | 4, 74 | 两次二分 |

### 二分模板

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1 // 或 left（插入位置）
}
```

---

*本文件由 Claude Code 自动生成*