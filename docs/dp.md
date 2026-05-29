# Dynamic Programming 动态规划 - LeetCode Deep Dive

## 目录
- [Problem 53: Maximum Subarray](#problem-53-maximum-subarray-最大子数组和)
- [Problem 55: Jump Game](#problem-55-jump-game-跳跃游戏)
- [Problem 62: Unique Paths](#problem-62-unique-paths-不同路径)
- [Problem 70: Climbing Stairs](#problem-70-climbing-stairs-爬楼梯)
- [Problem 72: Edit Distance](#problem-72-edit-distance-编辑距离)

---

## Problem 53: Maximum Subarray 最大子数组和

### 1. 题目核心与隐藏考点

**核心本质**: Kadane 算法，利用「前缀和」的概念，如果前一个位置的最大子序和为负，则丢弃，从当前位置重新开始。

**隐藏考点**:
- 全负数组的情况
- 为什么前缀为负要丢弃？
- 动态规划状态的定义

```
Kadane 算法图解:

nums: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Step 1: current_sum = -2, max_sum = -2
        [-2] vs [] → -2 最大

Step 2: current_sum + nums[1] = -2 + 1 = -1 < 0? 是的
        丢弃前一个，重置 current_sum = 1
        max_sum = 1

Step 3: current_sum = 1 + (-3) = -2 < 0，重置
        current_sum = 4
        max_sum = 4

... 继续

最终: 4 + (-1) + 2 + 1 = 6
max_sum = 6
```

---

### 2. 思路演进

#### 解法一：暴力枚举 O(n²) - 超时
```go
maxSum := math.MinInt
for i := 0; i < n; i++ {
    sum := 0
    for j := i; j < n; j++ {
        sum += nums[j]
        maxSum = max(maxSum, sum)
    }
}
```

#### 解法二：Kadane 算法 O(n) - 最优

```
动态规划状态:
  dp[i] = 以 nums[i] 结尾的最大子序和

状态转移:
  dp[i] = max(dp[i-1] + nums[i], nums[i])
        = max(前一个子序和 + 当前值, 从当前值重新开始)

解释:
  如果前一个子序和为负，加上它只会让和更小
  所以选择从当前值重新开始
```

---

### 3. 多解法对比表格

| 解法 | 时间 | 空间 | 优点 | 缺点 |
| :--- | :--- | :--- | :--- | :--- |
| 暴力枚举 | O(n²) | O(1) | 简单 | 超时 |
| 分治 | O(n log n) | O(log n) | 通用 | 复杂 |
| Kadane | O(n) | O(1) | 最优 | - |

---

### 4. 工业级 Go 源码与详细注释

```go
package dp

// MaxSubArray 最大子序和（Kadane 算法）
//
// 核心思想：动态规划 + 空间优化
//
// 动态规划状态定义：
// - dp[i] = 以 nums[i] 结尾的最大子序和
// - 状态转移：dp[i] = max(dp[i-1] + nums[i], nums[i])
//
// 为什么可以空间优化？
// - dp[i] 只依赖 dp[i-1]
// - 所以只需要一个变量记录前一个状态
// - 空间从 O(n) 优化到 O(1)
//
// 正确性证明（数学归纳法）：
// 1. base: dp[0] = nums[0]，以第一个元素结尾的最大子序和就是它本身
// 2. inductive: 假设 dp[i-1] 已知，dp[i] 有两种选择：
//    - 继承：以 nums[i] 结尾的子序包含 nums[i-1] 的最优子序
//    - 重置：从 nums[i] 开始一个新的子序
//    - 取最大值即为 dp[i]
// 3. 遍历过程中记录 maxSum，即为全局最大子序和
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func MaxSubArray(nums []int) int {
    maxSum := nums[0]     // 全局最大和
    currentSum := nums[0] // 以当前元素结尾的最大子序和

    // 从第二个元素开始
    for i := 1; i < len(nums); i++ {
        // 两种选择：继承前一个子序，或者从当前重新开始
        // 如果 currentSum < 0，加上它只会让和更小
        currentSum = max(currentSum+nums[i], nums[i])
        maxSum = max(maxSum, currentSum)
    }

    return maxSum
}
```

---

## Problem 55: Jump Game 跳跃游戏

### 1. 题目核心与隐藏考点

**核心本质**: 贪心，维护「最远可达位置」，如果当前位置超出最远可达，则失败。

**隐藏考点**:
- 局部最优是否能推出全局最优？
- 为什么每步都要更新最远距离？

```
贪心算法图解:

nums: [2, 3, 1, 1, 4]

Step 1: 位置 0，num=2，最远可达 = 0+2 = 2
        当前位置 0 <= 2，继续

Step 2: 位置 1，最远可达 = max(2, 1+3) = 4
        当前位置 1 <= 4，继续

Step 3: 位置 2，最远可达 = max(4, 2+1) = 4
        当前位置 2 <= 4，继续

... 继续检查，位置 4 <= 4，成功!
```

---

### 2. 工业级 Go 源码与详细注释

```go
// CanJump 判断是否能到达最后一个位置
//
// 核心思想：贪心
//
// 贪心策略：
// - 维护一个变量 maxReach，表示最远能到达的位置
// - 遍历数组，检查每个位置是否可达
// - 更新 maxReach = max(maxReach, i + nums[i])
// - 如果当前索引 i > maxReach，说明前面没有位置能跳到这里，返回 false
//
// 为什么这样正确？
// - 我们维护的是「从起点能达到的最远位置」
// - 每一步都考虑所有可能的选择（跳到 i+nums[i] 范围内的任何位置）
// - 如果当前位置超出最远可达，说明无论如何都无法到达这里
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func CanJump(nums []int) bool {
    maxReach := nums[0] // 初始最远可达

    for i := 1; i < len(nums); i++ {
        // 如果当前位置超出了最远可达范围，说明不可达
        if i > maxReach {
            return false
        }

        // 更新最远可达位置
        maxReach = max(maxReach, i+nums[i])
    }

    return true
}
```

---

## Problem 62: Unique Paths 不同路径

### 1. 题目核心与隐藏考点

**核心本质**: 网格路径计数，唯一路径数 = 上边路径数 + 左边路径数。

**隐藏考点**:
- 为什么是最优子结构？
- 状态转移方程如何推导？

```
动态规划图解:

m=3, n=7

  0 1 2 3 4 5 6
0 1 1 1 1 1 1 1
1 1 2 3 4 5 6 7
2 1 3 6 10 15 21 28

状态转移:
dp[i][j] = dp[i-1][j] + dp[i][j-1]

解释：
- 到达 (i,j) 的路径，要么从上方来，要么从左方来
- 上方路径数 = dp[i-1][j]
- 左方路径数 = dp[i][j-1]
- 总和就是唯一的路径数
```

---

### 2. 工业级 Go 源码与详细注释

```go
// UniquePaths 不同路径数
//
// 核心思想：动态规划
//
// 状态定义：
// - dp[i][j] = 到达格子 (i,j) 的唯一路径数
//
// 状态转移方程：
// - dp[i][j] = dp[i-1][j] + dp[i][j-1]
// - 到达 (i,j) 的路径，要么从上方来，要么从左方来
//
// 初始化：
// - dp[0][j] = 1：最上面一行只能从左边来
// - dp[i][0] = 1：最左边一列只能从上边来
//
// 空间优化：
// - 当前行只依赖上一行，可以优化到 O(n) 空间
// - 但 O(m*n) 的二维数组也很简单且足够
//
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func UniquePaths(m int, n int) int {
    // 创建 dp 表
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1 // 每行第一个格子只能从上边来
    }
    // 第一行只能从左边来
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }

    // 填表
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}
```

---

## Problem 70: Climbing Stairs 爬楼梯

### 1. 题目核心与隐藏考点

**核心本质**: 斐波那契数列，第 n 项等于第 n-1 项 + 第 n-2 项。

**隐藏考点**:
- 为什么是斐波那契？
- 如何从爬楼梯问题推导出斐波那契？

```
爬楼梯图解:

n=1: 1种 (1)
n=2: 2种 (1+1, 2)
n=3: 3种 (1+1+1, 1+2, 2+1)
n=4: 5种 (1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2)

递推关系:
f(1) = 1
f(2) = 2
f(n) = f(n-1) + f(n-2)

这正是斐波那契数列!
```

---

### 2. 工业级 Go 源码与详细注释

```go
// ClimbStairs 爬楼梯
//
// 核心思想：斐波那契数列
//
// 动态规划状态定义：
// - dp[i] = 爬到第 i 阶的不同方法数
//
// 状态转移方程：
// - dp[i] = dp[i-1] + dp[i-2]
// - 最后一步可以是爬 1 阶或爬 2 阶
// - 所以方法数 = 爬到第 i-1 阶的方法 + 爬到第 i-2 阶的方法
//
// 空间优化：
// - 只需要保存前两个状态
// - 空间从 O(n) 优化到 O(1)
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func ClimbStairs(n int) int {
    if n <= 2 {
        return n
    }

    // a = f(1), b = f(2)
    a, b := 1, 2

    for i := 3; i <= n; i++ {
        // 计算 f(i) = f(i-1) + f(i-2)
        // 然后更新 a, b
        a, b = b, a+b
    }

    return b
}
```

---

## Problem 72: Edit Distance 编辑距离

### 1. 题目核心与隐藏考点

**核心本质**: 二维 DP，将 word1 转换为 word2 的最少操作数。

**隐藏考点**:
- 为什么需要二维 DP？
- 状态转移方程如何推导？

```
DP 表格:

     "" h  o  r  s  e
""    0  1  2  3  4  5
r     1  1  2  2  3  4
o     2  2  1  2  3  4
s     3  3  2  2  2  3

状态转移:
dp[i][j] = min(dp[i-1][j] + 1,      # 删除
               dp[i][j-1] + 1,      # 插入
               dp[i-1][j-1] + cost) # 替换

cost = 0 if word1[i-1] == word2[j-1] else 1
```

---

### 2. 工业级 Go 源码与详细注释

```go
// MinDistance 编辑距离
//
// 核心思想：二维动态规划
//
// 状态定义：
// - dp[i][j] = word1[0..i-1] 转换为 word2[0..j-1] 的最少操作数
//
// 状态转移方程：
// - 如果 word1[i-1] == word2[j-1]，不需要操作
//   dp[i][j] = dp[i-1][j-1]
// - 否则，可以：
//   1. 删除 word1[i-1]：dp[i-1][j] + 1
//   2. 插入 word2[j-1]：dp[i][j-1] + 1
//   3. 替换 word1[i-1] 为 word2[j-1]：dp[i-1][j-1] + 1
//   取最小值
//
// 初始化：
// - dp[0][j] = j：空字符串转换为 word2[0..j-1] 需要 j 次插入
// - dp[i][0] = i：word1[0..i-1] 转换为空字符串需要 i 次删除
//
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func MinDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)

    // 创建 dp 表
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // 初始化：第一行和第一列
    for i := 0; i <= m; i++ {
        dp[i][0] = i // 删除 i 个字符
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j // 插入 j 个字符
    }

    // 填表
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                // 字符相同，不需要操作
                dp[i][j] = dp[i-1][j-1]
            } else {
                // 取三种操作的最小值
                dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
            }
        }
    }

    return dp[m][n]
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [198. House Robber](https://leetcode.com/problems/house-robber/) | 1D DP | Medium |
| [64. Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) | 2D DP | Medium |
| [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | 1D DP + 二分 | Medium |

---

## 动态规划总结

### DP 问题分类

| 类型 | 题目 | 特点 |
| :--- | :--- | :--- |
| 一维 DP | 53, 70 | 线性递推 |
| 二维 DP | 62, 72 | 网格/矩阵 |
| 贪心 + DP | 55 | 贪心选择 |

### DP 解题步骤

1. **定义状态**：dp[i] 或 dp[i][j] 表示什么
2. **写出转移方程**：如何从子问题推导
3. **确定初始值**：base case
4. **确定遍历顺序**：从前往后还是从后往前
5. **空间优化（如需要）**：观察依赖关系

---

*本文件由 Claude Code 自动生成*