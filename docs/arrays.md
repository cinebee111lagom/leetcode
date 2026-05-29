# Arrays 数组 - LeetCode Deep Dive

## 目录
- [Problem 1: Two Sum](#problem-1-two-sum-两数之和)
- [Problem 11: Container With Most Water](#problem-11-container-with-most-water-容器装水)
- [Problem 15: 3Sum](#problem-15-3sum-三数之和)
- [Problem 26: Remove Duplicates](#problem-26-remove-duplicates-from-sorted-array)
- [Problem 27: Remove Element](#problem-27-remove-element-移除元素)
- [Problem 33: Search in Rotated Array](#problem-33-search-in-rotated-sorted-array)
- [Problem 35: Search Insert Position](#problem-35-search-insert-position)
- [Problem 75: Sort Colors](#problem-75-sort-colors-颜色分类)
- [Problem 80: Remove Duplicates II](#problem-80-remove-duplicates-from-sorted-array-ii)

---

## Problem 1: Two Sum 两数之和

### 1. 题目核心与隐藏考点

**核心本质**: 通过哈希表将「查找 complement」从 O(n) 降为 O(1)，实现空间换时间。

**隐藏考点**:
- 数组中可能存在多组解，需要返回的是第一组有效解
- 不能重复使用同一元素
- 暴力解 O(n²) 会超时，必须优化

```go
// 核心思想: target - nums[i] = complement
// 用 map 存储已遍历的元素，查找复杂度 O(1)
```

---

### 2. 思路演进 (Brute Force → Optimal)

#### 解法一：暴力枚举 O(n²)
```go
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
        if nums[i] + nums[j] == target {
            return []int{i, j}
        }
    }
}
// 瓶颈: 双层循环，每个元素都要和之后的所有元素比较
// 时间: O(n²), 空间: O(1)
```

#### 解法二：哈希表 O(n)

```
初始: nums = [2, 7, 11, 15], target = 9

Step 1: i=0, num=2, complement=7, map={}
        → 7 not found, map={2:0}

Step 2: i=1, num=7, complement=2, map={2:0}
        → 2 found at index 0
        → 返回 [0, 1]

动画示意:
index:  0   1   2   3
nums:  [2] [7] [11] [15]
        ↑
      complement=7, check map→found!
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 暴力枚举 | O(n²) | O(1) | 简单直接 | 数据量大时超时 |
| 哈希表 | O(n) | O(n) | 高效 | 需要额外空间 |

---

### 4. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 1: Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers
// such that they add up to target. You may assume that each input would have exactly one solution.
package arrays

// TwoSum 使用哈希表实现 O(n) 时间复杂度
//
// 为什么用哈希表？
// - 暴力解需要 O(n²) 因为要查找「target - num」需要遍历整个数组
// - 哈希表的查找复杂度是 O(1)，可以将「寻找 complement」的时间从 O(n) 降为 O(1)
// - 所以总时间复杂度从 O(n²) 降为 O(n)
//
// 关键 insight: 对于每个元素 nums[i]，我们只需要查找 target - nums[i] 是否在之前出现过
// 这就转化成了「在集合中查找元素」的问题
func TwoSum(nums []int, target int) []int {
    // 哈希表: key=数值, value=该数值对应的索引
    // 使用 map 而非数组，因为数组值范围未知，且 map 更节省空间
    m := make(map[int]int)

    for i, num := range nums {
        // 计算当前元素需要的 complement
        complement := target - num

        // 在 map 中查找 complement
        // 如果找到，说明之前遍历过某个元素与当前元素之和为 target
        // 注意：这里查找的是 complement，不是 num 本身
        if j, exists := m[complement]; exists {
            // 找到了！返回之前元素的索引和当前索引
            // 注意顺序：j 是之前元素的索引，i 是当前元素的索引
            return []int{j, i}
        }

        // 没找到，将当前元素及其索引存入 map
        // 为什么要在这里存入？因为后续元素可能需要当前元素作为 complement
        m[num] = i
    }

    // 按题目要求，正常情况下一定会找到解
    // 但为了编译通过，返回 nil
    return nil
}
```

---

### 5. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [15. 3Sum](https://leetcode.com/problems/3sum/) | 哈希表 + 双指针 | Medium |
| [170. Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | 双指针 | Easy |

---

## Problem 11: Container With Most Water 容器装水

### 1. 题目核心与隐藏考点

**核心本质**: 双指针从两端向中间收敛，每次移动较短的板，因为短板决定了容器高度。

**隐藏考点**:
- 移动指针的决策：为什么移动短板而不是长板？
- 收敛性证明：为什么双指针一定能找到最优解？
- 面积计算: min(height[i], height[j]) * (j - i)

```
初始状态:
    i=0, j=9 (高度最低的两个点)

为什么移动短板？
┌─────────────────────────────────────────────────────┐
│ 移动长板不可能增加面积，因为：                        │
│ 新面积 = min(新长板, 另一端) * 新宽度                  │
│       ≤ min(旧长板, 另一端) * 新宽度（更窄）          │
│                                                      │
│ 但移动短板可能找到更大的面积！                        │
└─────────────────────────────────────────────────────┘
```

---

### 2. 思路演进 (Brute Force → Optimal)

#### 解法一：暴力枚举 O(n²)
```go
maxArea := 0
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
        height := min(height[i], height[j])
        width := j - i
        maxArea = max(maxArea, height*width)
    }
}
// 遍历所有可能的容器，总共 C(n,2) 种组合
// 瓶颈: O(n²) 时间复杂度
```

#### 解法二：双指针 O(n)

```
高度数组: [1,8,6,2,5,4,8,3,7]

初始: i=0, j=8
┌──────────────────────────────────────────────────────────┐
│  1   8   6   2   5   4   8   3   7                       │
│  ↑                           ↑                           │
│  i                           j                           │
│  area = min(1,7) * 8 = 8                                  │
│                                                          │
│  短板是 height[i]=1，移动 i++                            │
└──────────────────────────────────────────────────────────┘

Step 1: i=1, j=8
┌──────────────────────────────────────────────────────────┐
│  1   8   6   2   5   4   8   3   7                       │
│      ↑                           ↑                       │
│      i                           j                       │
│  area = min(8,7) * 7 = 56 (最大值!)                      │
│                                                          │
│  长板是 height[j]=7，移动 j--                            │
└──────────────────────────────────────────────────────────┘

继续收敛...
```

**为什么正确？** (反证法)
1. 假设最优解的左右边界是 (a, b)，初始指针在 (i, j)
2. 如果 i < a 且 j > b，那么 min(height[i], height[j]) ≤ min(height[a], height[b])
3. 因为 width(j-i) > width(b-a)，所以只有移动短边才可能突破当前最优

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 暴力枚举 | O(n²) | O(1) | 简单 | 数据量大时超时 |
| 双指针 | O(n) | O(1) | 最优 | 需要数学证明 |

---

### 4. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 11: Container With Most Water
// Given n non-negative integers representing bar heights, find two lines that together
// with x-axis form a container that holds the most water.
package arrays

// MaxArea 双指针解法 O(n)
//
// 核心思想：
// 1. 左右指针从数组两端开始，此时宽度最大
// 2. 容器水量由短板决定：min(height[i], height[j])
// 3. 移动短板的指针，可能找到更高的短板，从而增加水量
// 4. 移动长板的指针不可能增加水量（宽度减少，高度最多不变）
//
// 为什么移动短边？
// - 假设 height[i] < height[j]，面积为 min(height[i], height[j]) * (j-i)
// - 如果移动 j（长板），新面积为 min(height[i], height[j-1]) * (j-1-i)
//   因为 min(height[i], height[j-1]) ≤ height[i]，且宽度减少，所以面积不会增加
// - 如果移动 i（短板），可能找到更高的板，从而增加水量
//
// 收敛性：
// - 每一步都移动一个指针，最坏情况 O(n) 步收敛
// - 正确性基于：任何最优解都会被检查到（可数学证明）
func MaxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0

    for left < right {
        // 计算当前容器面积
        // width = right - left，height 由短板决定
        width := right - left
        h := min(height[left], height[right])
        maxArea = max(maxArea, width*h)

        // 移动短板指针
        // 为什么移动短板？因为只有移动短板才有可能找到更大的面积
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}
```

---

### 5. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | 双指针 + 单调栈 | Hard |
| [Container With Most Water II](https://leetcode.com/problems/container-with-most-water-ii/) | 贪心 | Medium |

---

## Problem 15: 3Sum 三数之和

### 1. 题目核心与隐藏考点

**核心本质**: 排序 + 双指针 + 跳过重复解，将三维问题降为二维。

**隐藏考点**:
- 去重策略：如何避免重复的三元组？
- 排序的必要性：为双指针创造条件
- 边界条件：数组长度 < 3 的情况

```
排序后: [-4, -1, -1, 0, 1, 2]

去重关键:
  [-1, -1, 2] vs [-1, 0, 1]  ← -1 相同但下一个数不同，所以是不同的三元组
  [-1, 0, 1]  ← 这是正确的三元组
```

---

### 2. 思路演进

#### 解法一：暴力枚举 O(n³) - 不可行
```go
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
        for k := j+1; k < n; k++ {
            if nums[i] + nums[j] + nums[k] == 0 {
                // 添加到结果
            }
        }
    }
}
// O(n³) 时间复杂度，LeetCode 会超时
```

#### 解法二：排序 + 双指针 O(n²)

```
数组: [-1, 0, 1, 2, -1, -4] 排序后 → [-4, -1, -1, 0, 1, 2]

i=0, num=-4:
    left=1, right=5
    -4 + (-1) + 2 = -3 ≠ 0

i=1, num=-1:
    left=2, right=5
    -1 + (-1) + 2 = 0 → 找到一个解!
    添加 [-1, -1, 2]
    跳过重复: left++ → 3, right-- → 4
    继续: left=3, right=4
    -1 + 0 + 1 = 0 → 找到一个解!
    添加 [-1, 0, 1]

i=2, num=-1 (与 i=1 相同，跳过!)
i=3, num=0 (小于等于0，继续)
i=4, num=1 (大于0，停止循环)
```

**去重逻辑图解**:
```
     i=1        i=2
   [-1, -1]   [-1, -1]   ← 这两个指向相同的值，但产生不同组合
       ↓           ↓
    left=2      跳过
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 暴力枚举 | O(n³) | O(1) | 简单 | 超时 |
| 哈希表 | O(n²) | O(n) | 无需排序 | 去重复杂 |
| 排序+双指针 | O(n²) | O(1) | 去重简单 | 需排序 |

---

### 4. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 15: 3Sum
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
package arrays

import "sort"

// ThreeSum 排序 + 双指针解法
//
// 核心思想：
// 1. 排序：方便双指针移动，且容易去重
// 2. 固定一个数 nums[i]，用双指针找另外两个数
// 3. 双指针从 nums[i+1] 和 nums[n-1] 开始，根据和的大小移动
//
// 为什么排序后可以用双指针？
// - 排序后，如果 nums[i] + nums[left] + nums[right] > 0，说明和太大
// - 因为 nums[right] 是最大值，所以需要减小 right
// - 反之，如果 < 0，需要增大 left
//
// 去重策略：
// - 外层循环：if i > 0 && nums[i] == nums[i-1] { continue }
//   跳过相同的第一个数，避免重复三元组
// - 内层循环：同样跳过相同的 left 和 right
func ThreeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    n := len(nums)

    // i < n-2: 至少需要三个数
    // nums[i] <= 0: 如果 nums[i] > 0，则后面不可能有和为 0 的三元组（已排序）
    for i := 0; i < n-2 && nums[i] <= 0; i++ {
        // 去重：跳过相同的 nums[i]
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, n-1
        target := -nums[i] // 我们需要找到 nums[left] + nums[right] == target

        for left < right {
            sum := nums[left] + nums[right]

            if sum == target {
                // 找到一个三元组
                result = append(result, []int{nums[i], nums[left], nums[right]})

                // 跳过重复的 left
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // 跳过重复的 right
                for left < right && nums[right] == nums[right-1] {
                    right--
                }

                // 继续寻找下一个可能的组合
                left++
                right--

            } else if sum < target {
                // 和太小，需要增大 left（nums[left] 更大）
                left++
            } else {
                // 和太大，需要减小 right（nums[right] 更小）
                right--
            }
        }
    }

    return result
}
```

---

### 5. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [18. 4Sum](https://leetcode.com/problems/4sum/) | 排序 + 双指针 + 外层循环 | Medium |
| [16. 3Sum Closest](https://leetcode.com/problems/3sum-closest/) | 排序 + 双指针 | Medium |

---

## Problem 26: Remove Duplicates from Sorted Array 移除排序数组中的重复项

### 1. 题目核心与隐藏考点

**核心本质**: 快慢指针，快指针遍历数组，慢指针指向下一个不重复元素应放置的位置。

**隐藏考点**:
- 原地修改：不能使用额外数组
- 返回值意义：慢指针之前的元素都是不重复的
- 边界情况：空数组、单元素数组

```
原始数组: [1, 1, 2, 2, 3]
          ↑
         slow (第一个元素)

Step 1: fast=1, nums[1]=1 → 重复，fast++
Step 2: fast=2, nums[2]=2 ≠ nums[slow]=1 → 不重复
        nums[slow+1]=nums[2] → [1, 2, ...]
        slow++

结果: slow=2, 数组变为 [1, 2, 2, 2, 3]
      返回 2（不重复元素个数）
```

---

### 2. 思路演进

#### 解法一：双指针（最优解）
```
快指针 j 遍历数组，慢指针 i 指向下一个不重复元素应存放的位置

初始: i=1, j=1 (第二个元素开始检查)

情况1: nums[j] != nums[i-1] (发现了不重复元素)
      nums[i] = nums[j]
      i++

情况2: nums[j] == nums[i-1] (重复，跳过)
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 双指针 | O(n) | O(1) | 最优，原地 | 需要理解指针含义 |
| 额外数组 | O(n) | O(n) | 简单 | 非原地 |

---

### 4. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 26: Remove Duplicates from Sorted Array
// Given a sorted array nums, remove duplicates in-place such that each element appears only once.
package arrays

// RemoveDuplicates 快慢指针原地去重
//
// 核心思想：
// 1. 快指针 j 遍历整个数组，探测不重复元素
// 2. 慢指针 i 指向下一个不重复元素应存放的位置
// 3. 当发现 nums[j] != nums[i-1] 时，将 nums[j] 复制到 nums[i]
//
// 为什么这个算法是正确的？
// - 数组已排序，相等元素一定相邻
// - i-1 指向已确认的最后一个不重复元素
// - nums[j] 与 nums[i-1] 不同，说明 nums[j] 是一个新的不重复元素
//
// 为什么不用额外空间？
// - 我们在原数组上「覆盖」而非「交换」
// - 最终，i 就是不重复元素的个数，nums[0:i] 就是去重后的数组
func RemoveDuplicates(nums []int) int {
    // 边界情况：空数组或只有一个元素
    if len(nums) == 0 {
        return 0
    }

    unique := 1 // 第一个元素一定是唯一的（已排序）

    // 从第二个元素开始检查
    for i := 1; i < len(nums); i++ {
        // nums[i] 与前一个不重复元素比较
        // 如果不同，说明找到了新的不重复元素
        if nums[i] != nums[unique-1] {
            // 将新元素放到正确位置
            nums[unique] = nums[i]
            unique++
        }
        // 如果相同，什么都不做，继续检查下一个元素
    }

    return unique
}
```

---

### 5. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [80. Remove Duplicates II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) | 快慢指针 + 计数器 | Medium |
| [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/) | 快慢指针 | Easy |

---

## Problem 33: Search in Rotated Sorted Array 搜索旋转排序数组

### 1. 题目核心与隐藏考点

**核心本质**: 二分查找的变形，需要判断目标落在哪个有序区间。

**隐藏考点**:
- 旋转点导致数组分为两个有序区间
- 每次迭代需要判断哪半边是有序的
- 边界条件容易出错

```
旋转数组示例:
原始: [0,1,2,3,4,5,6,7]
旋转4: [4,5,6,7,0,1,2,3]

二分查找时需要判断：
- nums[left..mid] 是否有序
- target 落在哪个有序区间
```

---

### 2. 思路演进

#### 二分查找变形

```
搜索 target=0，数组=[4,5,6,7,0,1,2,3]

Step 1: left=0, right=7, mid=3
        nums[mid]=7 > nums[left]=4
        → 左半边 [4,5,6,7] 有序
        → target 不在左半边 (0 < 4)，搜索右半边
        left=4

Step 2: left=4, right=7, mid=5
        nums[mid]=1 < nums[left]=0? 不，nums[left]=0? 不是
        nums[mid]=1 > nums[left]=0
        → 左半边 [0,1] 有序
        → target 在左半边 (0 >= 0 且 0 < 1)
        right=mid-1=4

Step 3: left=4, right=4, mid=4
        nums[mid]=0 == target → 找到!

二分查找决策树:
┌─────────────────────────────────────┐
│  mid 在哪边有序？                    │
├─────────────────┬───────────────────┤
│  nums[left] ≤ mid                   │
│  左半边有序?                          │
├────────┬────────┬────────┬───────────┤
│  是    │   否   │  是    │   否      │
│target  │target  │target  │target     │
│在左区间│在右区间│在左区间│在右区间    │
└────────┴────────┴────────┴───────────┘
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 遍历 | O(n) | O(1) | 简单 | 无优化 |
| 二分查找 | O(log n) | O(1) | 最优 | 边界复杂 |

---

### 4. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 33: Search in Rotated Sorted Array
// Given the array after rotation and an integer target, return the index of target.
package arrays

// SearchRotated 二分查找变形
//
// 核心思想：
// 1. 旋转数组由两个有序数组组成
// 2. 在二分查找的每一步，判断 target 落在哪个有序区间
// 3. 通过比较 nums[left] 和 nums[mid] 判断左半边是否有序
//
// 判断逻辑：
// - 如果 nums[left] <= nums[mid]，说明左半边 [left..mid] 有序
//   - 如果 target 在 [nums[left], nums[mid]) 范围内，收缩 right
//   - 否则，扩展 left
// - 否则，右半边 [mid..right] 有序
//   - 如果 target 在 (nums[mid], nums[right]] 范围内，扩展 left
//   - 否则，收缩 right
//
// 为什么这样正确？
// - 旋转点左侧的所有元素都大于右侧的所有元素
// - 通过 nums[left] 和 nums[mid] 的关系，可以确定哪个半边是有序的
// - 确定有序半边后，可以精确判断 target 是否在该区间内
func SearchRotated(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        // 判断左半边是否有序
        if nums[left] <= nums[mid] {
            // target 在左半边的范围内
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // 右半边有序
            // target 在右半边的范围内
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}
```

---

### 5. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [81. Search in Rotated Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/) | 二分 + 去重 | Medium |
| [153. Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | 二分 | Medium |

---

## Problem 75: Sort Colors 颜色分类

### 1. 题目核心与隐藏考点

**核心本质**: 三路划分（Dutch National Flag Algorithm），将数组分为红色(0)、白色(1)、蓝色(2)。

**隐藏考点**:
- 一次遍历完成排序
- 不能使用库函数
- 需要 O(1) 空间

```
三路划分示意图:
[2, 0, 2, 1, 1, 0]

初始: low=0, mid=0, high=5

Step: mid=0, nums[mid]=2 (蓝色)
      swap(nums[mid], nums[high]) → [0, 0, 1, 1, 2, 2]
      high--

... 继续直到 mid > high
```

---

### 2. 工业级 Go 源码与详细注释

```go
// Package arrays - LeetCode Problem 75: Sort Colors (Dutch National Flag)
// Given an array with 0s, 1s and 2s, sort them in-place.
package arrays

// SortColors 三路划分算法（荷兰国旗问题）
//
// 核心思想：
// 维护三个指针：low（红色边界），mid（当前位置），high（蓝色边界）
// - [0..low-1]: 红色 (0)
// - [low..mid-1]: 白色 (1)
// - [mid..high]: 未处理
// - [high+1..n-1]: 蓝色 (2)
//
// 操作逻辑：
// - nums[mid] == 0: 交换到左边，low++, mid++
// - nums[mid] == 1: 继续检查，mid++
// - nums[mid] == 2: 交换到右边，high--
//
// 为什么正确？
// - 每次交换后，要么 mid 前进，要么 high 后退
// - 交换到 low 的元素一定是红色（0），不需要再检查
// - 交换到 high 的元素可能是任意值，需要再检查
func SortColors(nums []int) {
    low, mid, high := 0, 0, len(nums)-1

    for mid <= high {
        switch nums[mid] {
        case 0:
            // 红色，交换到左边
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        case 1:
            // 白色，继续
            mid++
        case 2:
            // 蓝色，交换到右边
            nums[mid], nums[high] = nums[high], nums[mid]
            high--
            // 注意：mid 不前进，因为交换过来的元素还没检查过
        }
    }
}
```

---

### 3. 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [Sort List](https://leetcode.com/problems/sort-list/) | 三路划分 | Medium |
| [75. Sort Colors](https://leetcode.com/problems/sort-colors/) | 同本题 | Medium |

---

## 更多题目

（后续会继续补充 Problem 27, 35, 80 等题目）

---

## 总结

### 数组类题目常见模式

| 模式 | 题目 | 关键点 |
| :--- | :--- | :--- |
| 双指针 | 11, 26, 27 | 从两端或两端向中间 |
| 哈希表 | 1, 15 | 空间换时间 |
| 二分查找 | 33, 35 | 有序数组 |
| 排序+双指针 | 15 | 去重 |
| 三路划分 | 75 | 荷兰国旗 |

---

*本文件由 Claude Code 自动生成，用于 LeetCode 1-100 深度讲解*