# Backtracking 回溯 - LeetCode Deep Dive

## 目录
- [Problem 17: Letter Combinations](#problem-17-letter-combinations-of-phone-number-电话号码的字母组合)
- [Problem 22: Generate Parentheses](#problem-22-generate-parentheses-括号生成)
- [Problem 79: Word Search](#problem-79-word-search-单词搜索)

---

## Problem 17: Letter Combinations of Phone Number 电话号码的字母组合

### 1. 题目核心与隐藏考点

**核心本质**: 组合问题，每个数字对应多个字母，生成所有可能的组合。

**隐藏考点**:
- 回溯的模板
- 如何处理空输入
- 树的深度优先遍历

```
回溯树图解:

digits = "23"

数字到字母映射:
  2 → abc
  3 → def

        ""
       /|\
      a b c
     /| |...
    ad ae af bd be bf cd ce cf

遍历过程:
  Step 1: 选择 'a' → ad
  Step 2: 选择 'd' → ad
  Step 3: 选择 'e' → ad e? 不对...

实际上:
  Path "ad":
    选择 '2' 对应的 'a'
    选择 '3' 对应的 'd'
    得到 "ad"

  然后回溯，换 'd' 为 'e' → "ae"
  然后回溯，换 'a' 为 'b' → "bd"...

最终结果: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
```

---

### 2. 工业级 Go 源码与详细注释

```go
package backtracking

// 电话号码映射表
var phone = map[byte][]string{
    '2': {"a", "b", "c"},
    '3': {"d", "e", "f"},
    '4': {"g", "h", "i"},
    '5': {"j", "k", "l"},
    '6': {"m", "n", "o"},
    '7': {"p", "q", "r", "s"},
    '8': {"t", "u", "v"},
    '9': {"w", "x", "y", "z"},
}

// LetterCombinations 生成所有可能的字母组合
//
// 核心思想：回溯（Backtracking）
//
// 什么是回溯？
// - 一种深度优先搜索（DFS）策略
// - 在搜索过程中「尝试」一个选择，如果不成功就「撤销」这个选择
// - 类似于走迷宫，如果一条路走不通，就退回上一步尝试另一条路
//
// 算法步骤：
// 1. 从空字符串开始
// 2. 对于当前位置的每个可能选择：
//    - 做出选择
//    - 递归处理下一个位置
//    - 撤销选择（回溯）
// 3. 当路径长度等于输入长度时，得到一个完整解
//
// 时间复杂度：O(4^n * n) - n 是数字个数，4 是最大选项数
// 空间复杂度：O(n) - 递归栈深度
func LetterCombinations(digits string) []string {
    // 空输入处理
    if len(digits) == 0 {
        return []string{}
    }

    var result []string

    // 回溯函数
    // current: 当前构建的字符串
    // index: 下一个要处理的数字索引
    var backtrack func(current string, index int)
    backtrack = func(current string, index int) {
        // 如果当前字符串长度等于数字个数，说明找到了完整解
        if len(current) == len(digits) {
            result = append(result, current)
            return
        }

        // 获取当前数字对应的字母列表
        letters := phone[digits[index]]

        // 遍历所有可能的字母选择
        for _, letter := range letters {
            // 做选择：添加当前字母
            backtrack(current+letter, index+1)
            // 撤销选择（隐式，因为函数返回后 current 恢复原值）
        }
    }

    backtrack("", 0)
    return result
}
```

---

## Problem 22: Generate Parentheses 括号生成

### 1. 题目核心与隐藏考点

**核心本质**: 括号生成，经典回溯 + 剪枝，约束条件是「左括号数 >= 右括号数」。

**隐藏考点**:
- 剪枝条件：左括号数 <= 右括号数（合法括号串的前缀）
- 递归树的深度
- 为什么是卡特兰数？

```
回溯树图解:

n = 3

                 ""
            /
          "("          ← 左括号数 < n
         /
       "(("           ← 左括号数 < n
      /
    "((("             ← 左括号数 < n
      \
    "(()"             ← 右括号数 < 左括号数，可以加右括号
       \
     "(())"           ← 继续...

递归选择:
  - 如果左括号数 < n，可以加左括号
  - 如果右括号数 < 左括号数，可以加右括号

剪枝条件:
  - 右括号数不能超过左括号数（否则不合法）
```

---

### 2. 工业级 Go 源码与详细注释

```go
// GenerateParenthesis 生成所有有效的括号组合
//
// 核心思想：回溯 + 剪枝
//
// 约束条件：
// 1. 左括号数不能超过 n
// 2. 右括号数不能超过左括号数（否则非法）
//
// 为什么需要这些约束？
// - 有效的括号串中，任意前缀都满足「左括号数 >= 右括号数」
// - 这是因为每个右括号必须匹配一个之前的左括号
//
// 卡特兰数：
// - n 对括号的有效组合数是第 n 个卡特兰数 C(2n, n) / (n+1)
// - 当 n=3 时，有 5 种组合：((())), (()()), (())(), ()(()), ()()()
//
// 时间复杂度：O(4^n / sqrt(n)) - 卡特兰数级别
// 空间复杂度：O(n) - 递归栈深度
func GenerateParenthesis(n int) []string {
    var result []string

    // 回溯函数
    // current: 当前构建的字符串
    // open: 已使用的左括号数
    // close: 已使用的右括号数
    var backtrack func(current string, open int, close int)
    backtrack = func(current string, open int, close int) {
        // 如果当前字符串长度等于 2n，说明找到了完整解
        if len(current) == n*2 {
            result = append(result, current)
            return
        }

        // 选择 1：添加左括号（如果还有剩余）
        if open < n {
            backtrack(current+"(", open+1, close)
        }

        // 选择 2：添加右括号（如果右括号数 < 左括号数）
        if close < open {
            backtrack(current+")", open, close+1)
        }
    }

    backtrack("", 0, 0)
    return result
}
```

---

## Problem 79: Word Search 单词搜索

### 1. 题目核心与隐藏考点

**核心本质**: 二维网格的 DFS 回溯，需要标记已访问的单元格。

**隐藏考点**:
- 网格边界处理
- 为什么需要回溯（标记已访问）？
- 时间复杂度分析

```
DFS 回溯图解:

board:
  A B C E
  S F C S
  A D E E

word = "ABCCED"

搜索路径:
  A(0,0) → B(0,1) → C(0,2) → C(1,2) → E(1,3) → D(2,3) ✓

关键点:
  - 每走一步，标记当前单元格为已访问 '#'
  - 四个方向尝试：上、下、左、右
  - 如果走到死路，回溯（撤销标记），尝试其他方向

标记机制:
  - 访问 A 后，board[0][0] = '#'
  - 防止重复访问同一单元格
  - 回溯时恢复为原值
```

---

### 2. 工业级 Go 源码与详细注释

```go
// Exist 判断单词是否存在于网格中
//
// 核心思想：DFS 回溯
//
// 算法步骤：
// 1. 遍历网格中的每个单元格作为起点
// 2. 从起点开始深度优先搜索
// 3. 如果当前字符匹配 word 的当前索引：
//    - 标记当前单元格为已访问（临时替换为 '#'）
//    - 递归检查四个方向
//    - 如果找到完整匹配，返回 true
//    - 如果没找到，撤销标记（回溯）
// 4. 如果所有起点都尝试过仍未找到，返回 false
//
// 为什么需要回溯？
// - 一个单元格只能使用一次
// - 需要防止搜索过程中回到已访问的单元格
// - 但不同路径可能经过同一个单元格，所以搜索完一条路径后要恢复
//
// 时间复杂度：O(m * n * 4^L) - m*n 是网格大小，L 是单词长度
// 空间复杂度：O(L) - 递归栈深度
func Exist(board [][]byte, word string) bool {
    // 边界检查
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }

    m, n := len(board), len(board[0])

    // DFS 辅助函数
    var dfs func(i, j, index int) bool
    dfs = func(i, j, index int) bool {
        // 如果 word 的所有字符都已匹配，返回 true
        if index == len(word) {
            return true
        }

        // 边界检查和字符匹配检查
        // 注意：index 不会越界，因为上面已经检查过了
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[index] {
            return false
        }

        // 标记当前单元格为已访问
        board[i][j] = '#'

        // 递归检查四个方向
        // 只要有一个方向返回 true，就找到了解
        found := dfs(i+1, j, index+1) ||
            dfs(i-1, j, index+1) ||
            dfs(i, j+1, index+1) ||
            dfs(i, j-1, index+1)

        // 撤销标记（回溯）
        board[i][j] = word[index]

        return found
    }

    // 尝试每个单元格作为起点
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }

    return false
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [39. Combination Sum](https://leetcode.com/problems/combination-sum/) | 回溯 + 剪枝 | Medium |
| [46. Permutations](https://leetcode.com/problems/permutations/) | 回溯 | Medium |
| [51. N-Queens](https://leetcode.com/problems/n-queens/) | 回溯 | Hard |

---

## 回溯题目总结

### 回溯算法模板

```
func backtrack(params) {
    // 1. 终止条件
    if is终点 {
        添加结果
        return
    }

    // 2. 选择列表
    for 选择 in 选择列表:
        // 做选择
        做选择

        // 递归
        backtrack(新的params)

        // 撤销选择（回溯）
        撤销选择
}
```

### 常见模式

| 模式 | 题目 | 特点 |
| :--- | :--- | :--- |
| 组合 | 17, 22 | 选择子集 |
| 排列 | - | 全排列 |
| 搜索 | 79 | 二维网格 |

---

*本文件由 Claude Code 自动生成*