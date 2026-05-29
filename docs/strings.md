# Strings 字符串 - LeetCode Deep Dive

## 目录
- [Problem 3: Longest Substring Without Repeating](#problem-3-longest-substring-without-repeating-characters-无重复字符的最长子串)
- [Problem 5: Longest Palindromic Substring](#problem-5-longest-palindromic-substring-最长回文子串)
- [Problem 14: Longest Common Prefix](#problem-14-longest-common-prefix-最长公共前缀)
- [Problem 28: Implement strStr](#problem-28-implement-strstr-implement-strstr)
- [Problem 43: Multiply Strings](#problem-43-multiply-strings-字符串相乘)
- [Problem 49: Group Anagrams](#problem-49-group-anagrams-字母异位词分组)

---

## Problem 3: Longest Substring Without Repeating Characters 无重复字符的最长子串

### 1. 题目核心与隐藏考点

**核心本质**: 滑动窗口，用哈希表记录字符最后出现的位置，右指针扩展，左指针收缩。

**隐藏考点**:
- 滑动窗口的收缩条件
- 字符重复时索引更新的策略
- 窗口内子串的表示

```
滑动窗口图解:

s = "abcabcbb"

Step 1: i=0, j=0 → 'a', map={}, window=[a], max=1
Step 2: j=1 → 'b', map={}, window=[ab], max=2
Step 3: j=2 → 'c', map={}, window=[abc], max=3
Step 4: j=3 → 'a', map={a:0}, a 在窗口内 (0>=0)
        left = 0+1 = 1, window=[bc], 但实际上...

正确做法:
  当遇到重复字符 'a' 时，left 应该更新到重复字符的下一个位置
  但如果重复字符不在当前窗口内（如之前的 'a' 在索引 0，而 current left 已经是 1），
  不应该回退

更新策略:
  当 s[j] 在 map 中存在且 map[s[j]] >= left 时
  更新 left = map[s[j]] + 1

Step 5: j=4 → 'b'
        map[b]=1 >= left=1? 是的
        left = 1+1 = 2
        window=[c]? 不对...

让我重新看:
s = "abcabcbb"
left=0, right=0

j=3, char='a', map[a]=0, left=0
map[a] >= left? 0 >= 0? 是的
left = 0+1 = 1

j=4, char='b', map[b]=1, left=1
map[b] >= left? 1 >= 1? 是的
left = 1+1 = 2

j=5, char='c', map[c]=2, left=2
map[c] >= left? 2 >= 2? 是的
left = 2+1 = 3

j=6, char='b', map[b]=1, left=3
map[b] >= left? 1 >= 3? 不是
不更新 left

j=7, char='b', map[b]=4? 不对...

等等，让我重新看
当 j=4 时，我们已经更新 left=2
此时 map[b]=1 < left=2，所以 1 >= 2 是 false
left 不变

但问题是 map[b] 应该更新为 4（最新的位置）

让我看代码逻辑：
```go
if idx, ok := charIndex[s[i]]; ok && idx >= start {
    start = idx + 1
}
charIndex[s[i]] = i
```

所以：
- charIndex 记录的是每个字符最后出现的位置
- 当发现重复时，如果重复字符在当前窗口内（idx >= start），才更新 start
- 然后更新 charIndex[s[i]] = i

对于 "abcabcbb":
- j=0, 'a': charIndex[a]=0, start=0
- j=1, 'b': charIndex[b]=1
- j=2, 'c': charIndex[c]=2
- j=3, 'a': charIndex[a]=0 >= start=0? 是, start=1; charIndex[a]=3
- j=4, 'b': charIndex[b]=1 >= start=1? 是, start=2; charIndex[b]=4
- j=5, 'c': charIndex[c]=2 >= start=2? 是, start=3; charIndex[c]=5
- j=6, 'b': charIndex[b]=4 >= start=3? 是, start=5; charIndex[b]=6
- j=7, 'b': charIndex[b]=6 >= start=5? 是, start=7; charIndex[b]=7

最后 start=7，但窗口应该是最大的无重复子串

等等，让我重新看：
- 最长无重复子串应该是 "abc"，长度 3
- 或者 "bca"，长度 3
- 或者 "cab"，长度 3
- 或者 "abc"

实际上 "abc" 长度是 3

让我重新分析：
s = "abcabcbb"
无重复字符的最长子串是 "abc"，长度 3

正确！让我验证代码逻辑是否正确
```

---

### 2. 工业级 Go 源码与详细注释

```go
package strings

// LengthOfLongestSubstring 无重复字符的最长子串长度
//
// 核心思想：滑动窗口 + 哈希表
//
// 滑动窗口原理：
// - 维护一个窗口 [left, right)
// - 右指针 right 不断扩展，尝试加入新字符
// - 如果新字符与窗口内某字符重复：
//   - 将左指针 left 移到重复字符的下一个位置
//   - 这样可以确保窗口内没有重复字符
//
// 哈希表作用：
// - charIndex 记录每个字符最后出现的位置
// - 当发现重复时，可以 O(1) 时间定位到重复字符
//
// 更新策略：
// - 只有当重复字符在当前窗口内时（charIndex[char] >= left）才更新 left
// - 因为重复字符可能在窗口之外（如之前的字符已被排除）
// - 然后更新 charIndex[char] = right
//
// 时间复杂度：O(n) - 每个字符最多被访问两次（左右各一次）
// 空间复杂度：O(min(m, n)) - m 是字符集大小，n 是字符串长度
func LengthOfLongestSubstring(s string) int {
    // 空字符串处理
    if len(s) == 0 {
        return 0
    }

    // 哈希表：存储每个字符最后出现的位置
    charIndex := make(map[byte]int)

    maxLen := 0     // 全局最大长度
    start := 0      // 窗口左边界

    for i := 0; i < len(s); i++ {
        // 检查当前字符是否在窗口内重复
        // charIndex[s[i]] >= start 说明重复字符在当前窗口内
        if idx, exists := charIndex[s[i]]; exists && idx >= start {
            // 更新左边界到重复字符的下一个位置
            start = idx + 1
        }

        // 更新字符最后出现位置
        charIndex[s[i]] = i

        // 更新最大长度
        // i - start + 1 是当前窗口的长度
        if i-start+1 > maxLen {
            maxLen = i - start + 1
        }
    }

    return maxLen
}
```

---

## Problem 14: Longest Common Prefix 最长公共前缀

### 1. 题目核心与隐藏考点

**核心本质**: 垂直扫描，逐列比较，最多比较到最短字符串。

**隐藏考点**:
- 字符数组 vs 字符串遍历
- 空数组处理
- 何时终止比较

```
垂直扫描图解:

strs = ["flower", "flow", "flight"]

     f l o w e r
     f l o w     ← 'flow' 是 'flower' 的前缀
     f l i g h t

第一列: all 'f' ✓
第二列: all 'l' ✓
第三列: 'o' vs 'o' vs 'i' → 'i' 不同，停止
结果: "fl"

关键点：
- 每次比较第 i 列的所有字符
- 遇到不匹配就停止
- 最坏情况需要比较 min(len(strs)) 列
```

---

### 2. 工业级 Go 源码与详细注释

```go
// LongestCommonPrefix 最长公共前缀
//
// 核心思想：垂直扫描（Horizontal Scan）的变种
//
// 方法：逐字符比较
// 1. 以第一个字符串为基准
// 2. 遍历第一个字符串的每个字符
// 3. 对于每个字符，遍历所有字符串的相同位置
// 4. 如果有不匹配，提前返回
//
// 为什么从第一个字符串开始？
// - 可以处理空数组情况
// - 第一个字符串就是最短的公共前缀的上限
//
// 优化考虑：
// - 如果某个字符串为空，直接返回 ""
// - 每次比较时，检查是否已达到某个字符串的末尾
//
// 时间复杂度：O(S) - S 是所有字符串的字符总数
// 空间复杂度：O(1)
func LongestCommonPrefix(strs []string) string {
    // 空数组检查
    if len(strs) == 0 {
        return ""
    }

    // 以第一个字符串为基准
    prefix := strs[0]

    // 遍历第一个字符串的每个字符
    for i := 0; i < len(prefix); i++ {
        // 遍历其他所有字符串
        for j := 1; j < len(strs); j++ {
            // 如果当前字符越界，或者不匹配，返回已找到的前缀
            // 注意条件：len(strs[j]) < i 或 prefix[i] != strs[j][i]
            if len(strs[j]) <= i || prefix[i] != strs[j][i] {
                return prefix[:i]
            }
        }
    }

    return prefix
}
```

---

## Problem 28: Implement strStr() 实现 strStr()

### 1. 题目核心与隐藏考点

**核心本质**: KMP 算法，避免朴素匹配的重复比较，将时间复杂度从 O(mn) 降到 O(m+n)。

**隐藏考点**:
- 为什么 KMP 比朴素匹配快？
- 部分匹配表（Prefix Table）的含义
- 回退位置的确定

```
KMP 算法图解:

haystack = "hello", needle = "ll"

朴素匹配（会超时）:
  hello 中从每个位置开始比较 ll
  从位置 2 开始找到匹配

KMP 优化:
  - 在比较过程中，当发生不匹配时
  - 不是从头开始，而是利用已经匹配的信息
  - 利用部分匹配表找到回退位置

部分匹配表计算:
  needle = "ll"

  前缀: "", "l"
  后缀: "", "l"

  对于 "ll":
    - 前缀: "", "l", "ll"
    - 后缀: "", "l", "ll"
    - 最长公共前缀后缀: "l", 长度 1

  但对于 "ll"，部分匹配值应该是 [0, 1]
  因为第一个 'l' 的前后缀为空，长度 0
  第二个 'l' 的前缀 "l" 和后缀 "l" 匹配，长度 1

LPS (Longest proper Prefix which is also Suffix):
  [0, 0, 1, 2, ...]? 不对...

LPS for "ll" should be [0, 1] where:
  - lps[0] = 0 (no prefix-suffix for single char)
  - lps[1] = 1 (for "ll", "l" is both prefix and suffix)

在 KMP 匹配中:
  当遇到不匹配时，看 needle 中不匹配位置的前一个字符的 lps 值
  回退到 lps[prev] 位置继续比较
```

---

### 2. 工业级 Go 源码与详细注释

```go
// StrStr 实现 strStr (KMP 算法)
//
// 核心思想：KMP（Knuth-Morris-Pratt）算法
//
// 为什么需要 KMP？
// - 朴素匹配在 haystack 中移动时，可能会「回退」
// - 例如：haystack = "aaaaaab", needle = "aab"
// - 朴素匹配会多次比较已知的匹配字符
// - KMP 通过预处理 needle，构建部分匹配表（LPS），避免回退
//
// LPS 表的含义：
// - lps[i] 表示 needle[0..i] 这个子串中，
//   最长的相等前后缀的长度
// - 例如：needle = "aabab"
//   - lps[0] = 0 ('a')
//   - lps[1] = 0 ('aa' 没有相等前后缀)
//   - lps[2] = 1 ('aab' 的前缀 'a' 和后缀 'a' 匹配)
//   - ...
//
// KMP 算法流程：
// 1. 预处理 needle，构建 lps 表
// 2. 用两个指针 i（haystack）和 j（needle）遍历
// 3. 当匹配时，两个指针都前进
// 4. 当不匹配时，根据 lps[j-1] 回退 j，而不是从头开始
//
// 时间复杂度：O(m + n) - m = len(haystack), n = len(needle)
// 空间复杂度：O(n) - lps 表的大小
func StrStr(haystack string, needle string) int {
    // 空 needle 返回 0（题目规定）
    if needle == "" {
        return 0
    }

    // needle 比 haystack 长，不可能匹配
    if len(needle) > len(haystack) {
        return -1
    }

    // 构建 LPS 表
    lps := computeLPS(needle)

    i, j := 0, 0 // i 是 haystack 索引，j 是 needle 索引

    for i < len(haystack) {
        // 匹配，移动指针
        if haystack[i] == needle[j] {
            i++
            j++

            // needle 完全匹配
            if j == len(needle) {
                return i - j
            }
        } else {
            // 不匹配，根据 LPS 回退
            if j != 0 {
                // j 回退到 lps[j-1]
                j = lps[j-1]
            } else {
                // j 已经回退到 0，i 前进
                i++
            }
        }
    }

    return -1
}

// computeLPS 构建 LPS (Longest Proper Prefix which is also Suffix) 表
//
// LPS[i] = needle[0..i] 子串的最长相等前后缀长度
//
// 算法：
// - 用两个指针 len 和 i
// - len 记录当前匹配的前后缀长度
// - i 遍历字符串
// - 当字符匹配时，len++，lps[i] = len
// - 当字符不匹配时，根据 lps[len-1] 回退 len
//
// 举例：needle = "aabab"
//
//  i=0: 'a', len=0, lps[0]=0
//  i=1: 'a', match, len=1, lps[1]=1
//  i=2: 'b', not match, len=lps[0]=0
//       'b', match, len=1, lps[2]=1
//  i=3: 'a', match, len=2, lps[3]=2
//  i=4: 'b', match, len=3, lps[4]=3
//
// 结果：lps = [0, 1, 1, 2, 3]
func computeLPS(pattern string) []int {
    lps := make([]int, len(pattern))
    length := 0 // 当前匹配的前后缀长度

    i := 1 // 从第二个字符开始
    for i < len(pattern) {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                // 不匹配，回退到上一个可能的匹配位置
                length = lps[length-1]
            } else {
                // length == 0，无法回退
                lps[i] = 0
                i++
            }
        }
    }

    return lps
}
```

---

## Problem 5: Longest Palindromic Substring 最长回文子串

### 1. 题目核心与隐藏考点

**核心本质**: 中心扩展法，从每个可能的中心向外扩展，记录最长回文。

**隐藏考点**:
- 奇数长度 vs 偶数长度回文的不同处理
- 为什么中心扩展比动态规划更好？

```
中心扩展图解:

s = "babad"

情况1: 奇数长度回文 "bab"
       中心在索引 1

情况2: 偶数长度回文 "aba"
       中心在索引 1 和 2 之间

扩展过程:
  i=1, "bab" → 长度 3
  i=2, "aba" → 长度 3

最终: "bab" 或 "aba" (都是最长回文)
```

---

### 2. 工业级 Go 源码与详细注释

```go
package strings

// LongestPalindrome 最长回文子串
//
// 核心思想：中心扩展法
//
// 为什么不用 DP？
// - DP 需要 O(n²) 空间
// - 中心扩展只需要 O(1) 空间
// - 中心扩展更容易理解和实现
//
// 算法步骤：
// 1. 遍历每个可能的中心位置
// 2. 对于每个中心，分别处理奇数和偶数长度的情况
// 3. 向外扩展直到不匹配
// 4. 记录最长的回文
//
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func LongestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }

    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        // 奇数长度：以 i 为中心
        len1 := expandAroundCenter(s, i, i)
        // 偶数长度：以 i 和 i+1 为中心
        len2 := expandAroundCenter(s, i, i+1)

        // 取较长的回文
        maxLen := len1
        if len2 > maxLen {
            maxLen = len2
        }

        // 如果找到更长的回文，更新起始和结束位置
        if maxLen > end-start+1 {
            start = i - (maxLen-1)/2
            end = i + maxLen/2
        }
    }

    return s[start : end+1]
}

// expandAroundCenter 从中心向外扩展，返回最大回文长度
func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}
```

---

## Problem 43: Multiply Strings 字符串相乘

### 1. 题目核心与隐藏考点

**核心本质**: 模拟竖式乘法，从低位到高位逐位相乘，结果逆序存储。

**隐藏考点**:
- 大数乘法，不能直接用数字
- 结果数组逆序存储的原因
- 进位的处理

```
竖式乘法图解:

num1 = "123", num2 = "456"

    1 2 3
  × 4 5 6
  ───────
    7 3 8   (123 × 6)
   6 1 5    (123 × 5, 左移一位)
  4 9 2     (123 × 4, 左移两位)
  ───────
  5 6 0 8 8

逆序存储: [8, 8, 0, 5, 6] → "56088"

每个位置的结果:
- pos 0: 3×6 = 18, 进位 1, 结果 8
- pos 1: 2×6 + 3×5 + 1 = 20, 进位 2, 结果 0
- pos 2: 1×6 + 2×5 + 3×4 + 2 = 6+10+12+2 = 30, 进位 3, 结果 0
- pos 3: 1×5 + 2×4 + 3 = 5+8+3 = 16, 进位 1, 结果 6
- pos 4: 1×4 + 1 = 5, 进位 0, 结果 5

结果逆序: 5, 6, 0, 8, 8 → "56088"
```

---

### 2. 工业级 Go 源码与详细注释

```go
// Multiply 字符串相乘
//
// 核心思想：模拟竖式乘法
//
// 为什么不能用数字直接相乘？
// - 输入可能是很大的数（如 100 位）
// - 超过 int64 范围，无法用普通数字存储
// - 需要用字符串模拟乘法
//
// 算法步骤：
// 1. 从低位到高位逐位相乘
// 2. 结果存储在数组中，索引代表位置（逆序）
// 3. 每个位置存储当前位的累加值（可能 > 10）
// 4. 最后统一处理进位
// 5. 转为字符串（注意去除前导零）
//
// 关键点：
// - 结果数组长度 = len(num1) + len(num2)
// - num1[i] × num2[j] 的结果影响 pos = i + j 和 i + j + 1
// - 最终需要逆序输出
//
// 时间复杂度：O(m × n)
// 空间复杂度：O(m + n)
func Multiply(num1 string, num2 string) string {
    // 处理特殊情况
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    m, n := len(num1), len(num2)
    // 结果数组，长度为 m+n（最大可能长度）
    result := make([]int, m+n)

    // 从低位到高位逐位相乘
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            // 相乘
            mul := (num1[i] - '0') * (num2[j] - '0')
            // 加上之前的结果（可能有进位）
            p1, p2 := i+j, i+j+1
            sum := mul + result[p2]

            result[p2] = sum % 10
            result[p1] += sum / 10
        }
    }

    // 转为字符串，跳过前导零
    resultStr := ""
    for _, v := range result {
        if !(len(resultStr) == 0 && v == 0) {
            resultStr += string(rune(v + '0'))
        }
    }

    if resultStr == "" {
        return "0"
    }
    return resultStr
}
```

---

## Problem 49: Group Anagrams 字母异位词分组

### 1. 题目核心与隐藏考点

**核心本质**: 哈希表分组，key 是排序后的字符串。

**隐藏考点**:
- 如何判断字母异位词？
- 排序 vs 计数哪种方式更好？

```
哈希表分组图解:

strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

排序后的 key:
  "eat" → "aet"
  "tea" → "aet"
  "tan" → "ant"
  "ate" → "aet"
  "nat" → "ant"
  "bat" → "abt"

哈希表:
  "aet" → ["eat", "tea", "ate"]
  "ant" → ["tan", "nat"]
  "abt" → ["bat"]

结果:
  [["eat","tea","ate"], ["tan","nat"], ["bat"]]
```

---

### 2. 工业级 Go 源码与详细注释

```go
package strings

import (
    "sort"
)

// GroupAnagrams 字母异位词分组
//
// 核心思想：哈希表
//
// 关键 insight：
// - 字母异位词排序后相同
// - 因此可以用排序后的字符串作为 key
// - 将所有字母异位词归到同一组
//
// 算法步骤：
// 1. 创建哈希表，key 是排序后的字符串，value 是字符串数组
// 2. 遍历每个字符串
// 3. 将字符串排序，得到 key
// 4. 将原字符串加入 key 对应的组
// 5. 将哈希表的值转为二维数组返回
//
// 时间复杂度：O(n × k log k) - n 是字符串个数，k 是平均长度
// 空间复杂度：O(n × k)
func GroupAnagrams(strs []string) [][]string {
    // 哈希表：key = 排序后的字符串，value = 同一组的所有字符串
    groups := make(map[string][]string)

    for _, s := range strs {
        // 将字符串转为字符数组
        chars := []byte(s)
        // 排序
        sort.Slice(chars, func(i, j int) bool {
            return chars[i] < chars[j]
        })
        // 排序后的字符串作为 key
        key := string(chars)

        // 加入对应组
        groups[key] = append(groups[key], s)
    }

    // 将哈希表的值转为二维数组
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | 滑动窗口 | Hard |
| [30. Substring with Concatenation](https://leetcode.com/problems/substring-with-concatenation-of-all-words/) | 滑动窗口 | Hard |

---

## 字符串题目总结

### 常见模式

| 模式 | 题目 | 关键点 |
| :--- | :--- | :--- |
| 滑动窗口 | 3 | 收缩+扩展 |
| 字符串匹配 | 28 | KMP |
| 前缀处理 | 14 | 逐字符比较 |
| 中心扩展 | 5 | 回文检测 |
| 大数乘法 | 43 | 竖式模拟 |
| 哈希分组 | 49 | 排序 key |

---

*本文件由 Claude Code 自动生成*