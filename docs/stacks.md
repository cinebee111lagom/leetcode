# Stacks 栈 - LeetCode Deep Dive

## 目录
- [Problem 20: Valid Parentheses](#problem-20-valid-parentheses-有效的括号)
- [Problem 32: Longest Valid Parentheses](#problem-32-longest-valid-parentheses-最长有效括号)
- [Problem 71: Simplify Path](#problem-71-simplify-path-简化路径)
- [Problem 84: Largest Rectangle in Histogram](#problem-84-largest-rectangle-in-histogram-柱状图中最大的矩形)

---

## Problem 20: Valid Parentheses 有效的括号

### 1. 题目核心与隐藏考点

**核心本质**: 栈的后进先出（LIFO）特性，正好匹配括号的匹配特性——最近打开的括号应该最先关闭。

**隐藏考点**:
- 右括号可能没有对应的左括号
- 左括号可能多余
- 括号可能交错，如 `([)]`

```
匹配过程:
输入: "([])"

Step 1: 遇到 '('，push
        stack: [(]

Step 2: 遇到 '['，push
        stack: [(, []

Step 3: 遇到 ']'，栈顶是 '[', 匹配，pop
        stack: [(]

Step 4: 遇到 ')'，栈顶是 '('，匹配，pop
        stack: []

结果: 栈空，合法!
```

---

### 2. 思路演进

#### 解法：栈匹配

```
不合法情况:
  "([)]" → stack: [(, [] → 遇到 )，栈顶是 ]，不匹配!

  "(" → stack: [(] → 栈不空，非法!

核心：每遇到右括号，检查栈顶是否是匹配的左括号
```

---

### 3. 工业级 Go 源码与详细注释

```go
package stacks

// IsValid 检查括号字符串是否合法
//
// 核心思想：栈
// 1. 遇到左括号 '(', '[', '{'，入栈
// 2. 遇到右括号 ')', ']', '}'，检查栈顶是否是对应的左括号
//    - 如果是，pop
//    - 如果不是，返回 false
// 3. 遍历结束后，栈应该为空
//
// 为什么栈适合这个问题？
// - 括号必须「最近打开的先关闭」——这是 LIFO 的定义
// - 栈顶元素就是「最近打开」的那个括号
//
// 时间复杂度：O(n)
// 空间复杂度：O(n) - 最坏情况全是左括号
func IsValid(s string) bool {
    stack := make([]byte, 0)

    // 配对映射表
    mapping := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for i := 0; i < len(s); i++ {
        // 检查是否是右括号
        if closing, ok := mapping[s[i]]; ok {
            // 是右括号，检查栈顶
            // 注意：stack 为空时 len(stack) == 0
            if len(stack) == 0 || stack[len(stack)-1] != closing {
                return false
            }
            // 匹配，pop
            stack = stack[:len(stack)-1]
        } else {
            // 是左括号，入栈
            stack = append(stack, s[i])
        }
    }

    // 栈空表示完全匹配
    return len(stack) == 0
}
```

---

## Problem 32: Longest Valid Parentheses 最长有效括号

### 1. 题目核心与隐藏考点

**核心本质**: 动态规划或两次遍历，用栈记录非法位置，从而计算最大连续有效子串长度。

**隐藏考点**:
- "(()" 这种以左括号结尾的情况
- ")()())" 这种中间有多个有效段的情况

```
两次遍历法核心思路:
从左到右遍历:
  "(()"
  - 遇到 '('，left++
  - 遇到 ')'，right++
  - left == right，记录最大长度
  - right > left，重置计数

从右到左遍历（处理以 '(' 结尾的情况）:
  "(()"
  - 遇到 ')'，right++
  - 遇到 '('，left++
  - left == right，记录最大长度
  - left > right，重置计数
```

---

### 2. 工业级 Go 源码与详细注释

```go
// LongestValidParentheses 最长有效括号
//
// 方法一：两次遍历（推荐）
//
// 核心思想：
// 1. 从左到右遍历，统计左右括号数量
// 2. 当 left == right 时，说明找到了一个有效子串
// 3. 当 right > left 时，说明之前的选择无效，重置
// 4. 但这种方法无法处理 "(()" 这种情况（以左括号结尾）
// 5. 从右到左再遍历一次，可以处理这种情况
//
// 为什么需要两次遍历？
// - "(()" 从左到右：left=2, right=1，永远不会 left==right
// - 从右到左遍历时：left=2, right=1，此时 left > right
//   但重新计算后，当 left==right 时可以找到最长有效子串
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func LongestValidParentheses(s string) int {
    maxLen := 0
    left, right := 0, 0

    // 从左到右
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }

        if left == right {
            maxLen = max(maxLen, left*2)
        } else if right > left {
            // 无效，重置
            left, right = 0, 0
        }
    }

    // 从右到左
    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }

        if left == right {
            maxLen = max(maxLen, left*2)
        } else if left > right {
            // 无效，重置
            left, right = 0, 0
        }
    }

    return maxLen
}
```

---

## Problem 71: Simplify Path 简化路径

### 1. 题目核心与隐藏考点

**核心本质**: 栈处理路径规范化，遇到 ".." 弹栈，遇到 "." 或空跳过。

```
示例: "/home/foo/../bar/"

解析过程:
  /home/foo/../bar/
  ├── /        → 根目录，入栈
  ├── home     → 普通目录，入栈 [/, home]
  ├── foo      → 普通目录，入栈 [/, home, foo]
  ├── ..       → 返回上级，弹栈 [/, home]
  └── bar      → 普通目录，入栈 [/, home, bar]

结果: /home/bar
```

---

### 2. 工业级 Go 源码与详细注释

```go
// SimplifyPath 简化路径
//
// 核心思想：栈
// 1. 按 '/' 分隔路径
// 2. 遇到空字符串或 "."，跳过
// 3. 遇到 ".." 且栈不为空，pop（返回上级目录）
// 4. 遇到其他目录名，入栈
// 5. 最后拼接结果
//
// 为什么要用栈？
// - 路径的层级关系是 LIFO（后进先出）
// - 遇到 ".." 时，应该返回最近的那个目录
//
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func SimplifyPath(path string) string {
    stack := make([]string, 0)

    // 用 '/' 分隔路径
    parts := split(path, '/')

    for _, part := range parts {
        switch part {
        case "", ".":
            // 空字符串或当前目录，跳过
            continue
        case "..":
            // 返回上级目录，如果栈不为空则 pop
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        default:
            // 普通目录名，入栈
            stack = append(stack, part)
        }
    }

    // 拼接结果
    result := ""
    for _, dir := range stack {
        result += "/" + dir
    }

    // 特殊情况：空路径返回根目录
    if result == "" {
        result = "/"
    }

    return result
}

// 自定义 split 函数，避免 strings.Split 带来的空字符串
func split(s string, delimiter byte) []string {
    var parts []string
    start := 0
    for i := 0; i <= len(s); i++ {
        if i == len(s) || s[i] == delimiter {
            if start < i {
                parts = append(parts, s[start:i])
            }
            start = i + 1
        }
    }
    return parts
}
```

---

## Problem 84: Largest Rectangle in Histogram 柱状图中最大的矩形

### 1. 题目核心与隐藏考点

**核心本质**: 单调递增栈，找到每个柱子作为最矮柱子时的最大矩形边界。

**隐藏考点**:
- 单调栈中元素代表什么？
- 为什么要在数组末尾加一个高度为 0 的柱子？

```
单调栈图解:

heights: [2, 1, 5, 6, 2, 3]

Step 1: i=0, h=2, stack=[0]
Step 2: i=1, h=1
        1 < 2, pop! 计算以 height[0]=2 为高的矩形
        width = i = 1, area = 2*1 = 2
        1 < 1, pop! 计算以 height[1]=1 为高的矩形
        width = i = 1, area = 1*1 = 1

        stack=[]

Step 3: i=2, h=5, stack=[2]
Step 4: i=3, h=6, stack=[2,3]
Step 5: i=4, h=2
        2 < 6, pop! width = 4, height=6, area=24
        ... 类似处理

最终最大矩形: 5*3=15? 不对，应该是 5*3=15 或 4*5=20??

正确: 2*5=10??

等等让我重新看... 实际上 [2,1,5,6,2,3] 最大矩形是 5*2=10? 不对...

最大矩形是高度 5，宽度 3 (位置 2,3,4)，面积 = 15
或者高度 2，宽度 5 (位置 0-4)，面积 = 10

不对... 让我重新计算
输入: [2, 1, 5, 6, 2, 3]

可视化:
  █
  █ █
  █ █
█ █ █
█ █ █

最大矩形: 高度 3（第三根），宽度 3（位置 2,3,4），面积 = 9
不对...

让我看答案... 实际上答案是 10 (高度 2，宽度 5)

等等，题目说 sample 输入 [2, 1, 5, 6, 2, 3] 输出是 10
高度 2，宽度 5 = 10

正确答案是 10
```

---

### 2. 工业级 Go 源码与详细注释

```go
// LargestRectangleArea 柱状图中最大的矩形
//
// 核心思想：单调递增栈
//
// 单调栈的核心：
// - 维护一个递增的栈，存储柱子的索引
// - 当遇到比栈顶更低的柱子时，说明栈顶柱子的「最大宽度」已经确定
// - 可以弹出栈顶，计算以其为高的最大矩形
//
// 算法步骤：
// 1. 遍历每个柱子 i
// 2. 当 heights[i] < heights[stack[-1]] 时，说明栈顶柱子的右边边界确定了
// 3. 弹出栈顶，计算以其为高的矩形
// 4. 最后，栈中可能还有元素，需要处理
//
// 为什么在末尾加 0？
// - 保证最后所有柱子都能被处理
// - 当 i = n 时，所有剩余柱子的右边边界都确定了
//
// 时间复杂度：O(n) - 每个柱子最多入栈出栈各一次
// 空间复杂度：O(n)
func LargestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }

    stack := make([]int, 0) // 存储索引
    maxArea := 0

    for i := 0; i <= n; i++ {
        h := 0
        if i < n {
            h = heights[i]
        }

        // 当遇到比栈顶更低的柱子，或者遍历结束时
        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]

            // 计算宽度
            width := i
            if len(stack) > 0 {
                // 栈不为空时，左边界是栈顶元素的下一个位置
                width = i - stack[len(stack)-1] - 1
            }

            maxArea = max(maxArea, height*width)
        }

        stack = append(stack, i)
    }

    return maxArea
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [155. Min Stack](https://leetcode.com/problems/min-stack/) | 辅助栈 | Easy |
| [42. Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | 单调栈/双指针 | Hard |
| [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/) | 单调栈 | Easy |

---

## 栈题目总结

### 常见模式

| 模式 | 题目 | 关键点 |
| :--- | :--- | :--- |
| 括号匹配 | 20 | 配对检查 |
| 单调栈 | 84 | 递增/递减栈 |
| 路径简化 | 71 | 分割+弹栈 |

---

*本文件由 Claude Code 自动生成*