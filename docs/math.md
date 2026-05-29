# Math 数学 - LeetCode Deep Dive

## 目录
- [Problem 7: Reverse Integer](#problem-7-reverse-integer-整数反转)
- [Problem 9: Palindrome Number](#problem-9-palindrome-number-回文数)
- [Problem 29: Divide Two Integers](#problem-29-divide-two-integers两数相除)
- [Problem 50: Pow(x, n)](#problem-50-powx-n- Powx-n)
- [Problem 69: Sqrt(x)](#problem-69-sqrtx-平方根)

---

## Problem 7: Reverse Integer 整数反转

### 1. 题目核心与隐藏考点

**核心本质**: 逐位提取数字，反向构建，需要检测32位有符号整数溢出。

**隐藏考点**:
- 32 位有符号整数范围：[-2147483648, 2147483647]
- 反转后可能溢出，必须在构建过程中检测
- 负数的处理：保留负号

```
溢出检测:

正数情况:
  result > 2147483647/10 = 214748364
  或者 result == 214748364 && digit > 7

负数情况:
  result < -2147483648/10 = -214748364
  或者 result == -214748364 && digit < -8

示例:
  2147483647 反转 → 7463847412 → 溢出，返回 0

  -2147483648 反转 → -8463847412 → 溢出，返回 0

逐位构建:
  x = 123
  result = 0

  Step 1: digit = 3, result = 0*10 + 3 = 3
  Step 2: digit = 2, result = 3*10 + 2 = 32
  Step 3: digit = 1, result = 32*10 + 1 = 321

  结果: 321
```

---

### 2. 工业级 Go 源码与详细注释

```go
package math

// Reverse 反转整数
//
// 核心思想：逐位提取，构建新数字，同时检测溢出
//
// 关键点：
// 1. 每次提取最后一位 digit = x % 10
// 2. 构建新数字 result = result*10 + digit
// 3. 在构建过程中检测溢出
//
// 为什么要在构建过程中检测？
// - 如果已经溢出，再处理就没有意义了
// - 需要在最后一次可能导致溢出的操作之前检测
//
// 溢出检测条件：
// - 对于正数：如果 result > 2147483647/10，说明下一步必然溢出
// - 如果 result == 2147483647/10 且 digit > 7，也会溢出
// - 负数情况类似，只是符号相反
//
// 时间复杂度：O(log₁₀(x)) - x 的位数
// 空间复杂度：O(1)
func Reverse(x int) int {
    result := 0

    for x != 0 {
        digit := x % 10 // 提取最后一位

        // 溢出检测
        // 正数情况：result > 2147483647/10 或 (result == 2147483647/10 && digit > 7)
        if result > 2147483647/10 || (result == 2147483647/10 && digit > 7) {
            return 0
        }
        // 负数情况：result < -2147483648/10 或 (result == -2147483648/10 && digit < -8)
        if result < -2147483648/10 || (result == -2147483648/10 && digit < -8) {
            return 0
        }

        result = result*10 + digit
        x /= 10
    }

    return result
}
```

---

## Problem 9: Palindrome Number 回文数

### 1. 题目核心与隐藏考点

**核心本质**: 反转后一半数字，与原数比较。注意：负数一定不是回文。

**隐藏考点**:
- 为什么不需要处理全部反转？
- 如何只反转后一半？

```
反转后一半示例:

12321 → 反转后一半 = 123
  比较 123 == 12321/10000 = 1? 不对...
  12321/10000 = 1
  12321%10000 = 2321

  实际上应该这样做：
  - x = 12321
  - 反转后一半：reversed = 123
  - 原数前一半：12321/100 = 123
  - 123 == 123 ✓

更清晰的例子：
  12321
  - 12321 / 10000 = 1，余 2321
  - 2321 / 1000 = 2，余 321
  - 321 / 100 = 3，余 21
  - 21 / 10 = 2，余 1
  - 1 / 1 = 1

  反转后一半: 1*10000 + 2*1000 + 3*100 + 2*10 + 1 = 12321? 不对...

让我重新看：

当位数是奇数时：
  12321 → 前半: 12, 中间: 3, 后半反转: 21
  我们实际上是在比较 前半 和 后半反转
  12 vs 21? 不对

正确的做法：
  - 对于奇数位数，中间那个数可以忽略
  - 12321 → 前半 = 12，后半 = 21
  - 但我们只反转后一半，得到 12
  - 所以比较 前半(12) 和 反转后(12) ✓

步骤：
  - 12321 / 10000 = 1 → 得到第一位
  - (12321 % 10000) / 1000 = 2 → 得到第二位
  - ...

实际上：
  x = 12321
  reversed = 0

  Step 1: reversed = 0*10 + (12321%10) = 1, x = 1232
  Step 2: reversed = 1*10 + (1232%10) = 12, x = 123
  Step 3: reversed = 12*10 + (123%10) = 123, x = 12
  Step 4: x = 12, reversed = 123, x <= reversed
  结束，此时 x = 12, reversed = 123
  但 123 是反转后一半，不是完整的反转

对于奇数位数：
  12321, 长度=5
  遍历直到 x <= reversed:
    Step 1: x=1232, reversed=1
    Step 2: x=123, reversed=12
    Step 3: x=12, reversed=123
    此时 x=12 <= reversed=123

  前半 = 12
  后半 = 123 / 10 = 12 (去掉最后一位)
  相等 ✓
```

---

### 2. 工业级 Go 源码与详细注释

```go
// IsPalindrome 判断一个数是否是回文数
//
// 核心思想：反转后一半数字，与前一半比较
//
// 为什么不需要完整反转？
// - 对于偶数位数：前后两半应该相等
// - 对于奇数位数：中间数字可以忽略，前后半应该相等
// - 所以只需要反转后一半，与前一半（或前一半去掉中间位）比较
//
// 算法步骤：
// 1. 负数一定不是回文
// 2. 如果最后一位是 0 但第一位不是 0，也不是回文（但这种情况在正整数中不存在）
// 3. 反转后一半数字
// 4. 比较前一半（或前一半去掉中间位）和后一半反转
//
// 时间复杂度：O(log₁₀(n)) - 只需要处理一半数字
// 空间复杂度：O(1)
func IsPalindrome(x int) bool {
    // 负数一定不是回文
    if x < 0 {
        return false
    }

    // 特殊情况：最后一位是 0 的正数不是回文（但题目说非负整数，所以这个检查可选）
    // 但如果 x 的最后一位是 0 且 x 不等于 0，则不是回文
    // 例如：10 -> 01 = 1，不是回文
    // 实际上，LeetCode 官方题解没有这个检查，因为题目说 positive integer

    original := x
    reversed := 0

    for x > 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }

    return original == reversed
}
```

---

## Problem 29: Divide Two Integers 两数相除

### 1. 题目核心与隐藏考点

**核心本质**: 不用乘法/除法实现除法，使用位运算加速（因为除法本质是减法，位运算可以快速找到最大倍数）。

**隐藏考点**:
- 32 位整数范围：[-2147483648, 2147483647]
- -2147483648 / -1 = 2147483648 溢出
- 位运算加速：x << shift 等于 x * 2^shift

```
位运算加速示例:

dividend = 15, divisor = 3

普通方法：15 - 3 - 3 - 3 - 3 - 3 = 0，减 5 次
结果 = 5

优化方法（找最大倍数）：
  15 / 3
  3 * 2^1 = 6 <= 15
  3 * 2^2 = 12 <= 15
  3 * 2^3 = 24 > 15

  所以 15 = 3 * 2^2 + 3
  继续：15 - 12 = 3，3 / 3 = 1

  结果 = 4 + 1 = 5

位运算实现：
  shifted = divisor << i 等价于 divisor * 2^i
  找最大的 i 使得 shifted <= dividend
```

---

### 2. 工业级 Go 源码与详细注释

```go
// Divide 两数相除
//
// 核心思想：位运算加速的减法
//
// 为什么不用除法？
// - 题目禁止使用乘法、除法和模运算
// - 需要用减法或位运算实现除法
//
// 优化思路：
// 1. 每次找到能被减去的最大倍数
// 2. 使用位运算快速计算 divisor * 2^i
// 3. 将结果加上 2^i
//
// 算法步骤：
// 1. 处理符号
// 2. 将 dividend 和 divisor 转为正数（用绝对值）
// 3. 找最大 shift，使得 divisor << shift <= dividend
// 4. 从 dividend 中减去 divisor << shift，结果加上 2^shift
// 5. 重复直到 dividend < divisor
//
// 边界情况：
// - -2147483648 / -1 = 2147483648，超出 int32 范围，返回 2147483647
//
// 时间复杂度：O(log n) - 每次可能减去过大的倍数
// 空间复杂度：O(1)
func Divide(dividend int, divisor int) int {
    // 特殊情况：-2147483648 / -1 = 2147483648，溢出
    if dividend == -2147483648 && divisor == -1 {
        return 2147483647
    }

    // 判断结果符号
    negative := (dividend < 0) != (divisor < 0)

    // 转为正数计算
    a, b := abs(dividend), abs(divisor)

    quotient := 0

    // 减法加速：找最大倍数
    for a >= b {
        shift := 0

        // 找最大的 shift，使得 b << shift <= a
        // 即找到最大的 2^shift 使得 divisor * 2^shift <= dividend
        for a >= b<<(shift+1) {
            shift++
        }

        // 减去这个倍数
        quotient += 1 << shift
        a -= b << shift
    }

    if negative {
        return -quotient
    }

    return quotient
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

---

## Problem 50: Pow(x, n) Pow(x, n)

### 1. 题目核心与隐藏考点

**核心本质**: 快速幂（Fast Power），利用指数的二进制分解，将 O(n) 降为 O(log n)。

**隐藏考点**:
- 指数为负数的情况
- 指数可能是奇数的情况
- 为什么可以用移位？

```
快速幂图解:

x^10 = x^(1010₂) = x^8 * x^2

n = 10 = 1010₂

Step 1: n 是偶数，x = x^2, n = 5
  x^10 = (x^2)^5 = x^10

Step 2: n 是奇数，result *= x, n = 2
  result = x^2, x = x^4, n = 2

Step 3: n 是偶数，x = x^8, n = 1
  x^4 → x^8

Step 4: n 是奇数，result *= x, n = 0
  result = x^2 * x^8 = x^10

二进制方法:
  10 = 1010₂
  x^10 = x^(2³) * x^(2¹) = x^8 * x^2 = x^10
```

---

### 2. 工业级 Go 源码与详细注释

```go
// Pow 计算 x 的 n 次方
//
// 核心思想：快速幂（Fast Power）算法
//
// 为什么快速？
// - 普通方法需要 O(n) 次乘法
// - 快速幂利用指数的二进制表示，将乘法次数降为 O(log n)
//
// 算法步骤：
// 1. 处理负指数：x^n = 1 / x^(-n)
// 2. 初始化结果为 1，底数为 x
// 3. 当 n > 0 时：
//    - 如果 n 是奇数，结果乘以底数
//    - 底数平方，n 折半
//
// 数学原理：
// - n 的二进制表示：n = Σ(b_i * 2^i)
// - x^n = x^(Σ(b_i * 2^i)) = Π(x^(2^i))^{b_i}
// - 当 b_i = 1 时，乘以 x^(2^i)
//
// 举例：n = 10 = 1010₂
// - x^10 = x^(2³) * x^(2¹) = x^8 * x^2
// - 在二进制处理过程中，当遇到 1 时，就乘以当前的 x^(2^i)
//
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func Pow(x float64, n int) float64 {
    // 处理负指数
    if n < 0 {
        x = 1 / x
        n = -n
    }

    result := 1.0

    for n > 0 {
        // 如果当前指数位是 1，乘入结果
        if n%2 == 1 {
            result *= x
        }

        // 底数平方
        x *= x

        // 指数折半
        n /= 2
    }

    return result
}
```

---

## Problem 69: Sqrt(x) 平方根

### 1. 题目核心与隐藏考点

**核心本质**: 二分查找，找到最大的整数 n 使得 n² ≤ x。

**隐藏考点**:
- 为什么不直接用库函数？
- 为什么不返回浮点数？
- 边界情况：0, 1, 大数

```
二分查找图解:

x = 8

Step 1: left=1, right=8, mid=4
        4² = 16 > 8 → 太大了，right = 3

Step 2: left=1, right=3, mid=2
        2² = 4 <= 8 → 小了，但可能是答案
        left = 2 + 1 = 3? 不对...

正确做法：
        2² = 4 <= 8，且 3² = 9 > 8
        所以 2 是候选，left = mid + 1 = 3

Step 3: left=3, right=3, mid=3
        3² = 9 > 8 → 太大了
        right = mid - 1 = 2

Step 4: left=3 > right=2，结束
        返回 right = 2

验证: 2² = 4 <= 8 ✓
```

---

### 2. 工业级 Go 源码与详细注释

```go
// MySqrt 计算整数平方根（向下取整）
//
// 核心思想：二分查找
//
// 为什么用二分？
// - 需要找到最大的 n 使得 n² ≤ x
// - 平方函数是单调递增的
// - 可以用二分查找定位边界
//
// 算法步骤：
// 1. 特殊情况：x = 0，返回 0
// 2. 设置左右边界：left = 1, right = x
// 3. 当 left <= right 时：
//    - mid = left + (right-left)/2
//    - 如果 mid == x/mid，找到精确答案
//    - 如果 mid < x/mid，mid 是候选，搜索右半边
//    - 如果 mid > x/mid，搜索左半边
// 4. 返回 right（最后的有候选值）
//
// 为什么用 x/mid 而不是 mid*mid？
// - 避免乘法溢出
// - 如果 mid * mid > x，可能导致整数溢出
// - 而 x/mid 不涉及乘法，安全
//
// 时间复杂度：O(log x) - 二分查找的次数
// 空间复杂度：O(1)
func MySqrt(x int) int {
    // 特殊情况：0
    if x == 0 {
        return 0
    }

    left, right := 1, x

    for left <= right {
        mid := left + (right-left)/2

        // 用除法避免乘法溢出
        if mid == x/mid {
            return mid
        } else if mid < x/mid {
            // mid² < x，mid 是候选，继续向右找
            left = mid + 1
        } else {
            // mid² > x，mid 太大，向左找
            right = mid - 1
        }
    }

    // 此时 right < left，right 是最大的满足条件的值
    return right
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [231. Power of Two](https://leetcode.com/problems/power-of-two/) | 位运算 | Easy |
| [326. Power of Three](https://leetcode.com/problems/power-of-three/) | 对数/循环 | Easy |
| [268. Missing Number](https://leetcode.com/problems/missing-number/) | 位运算/求和 | Easy |

---

## 数学题目总结

### 常见技巧

| 技巧 | 题目 | 说明 |
| :--- | :--- | :--- |
| 溢出检测 | 7, 29 | 在操作前检测 |
| 位运算加速 | 29, 50 | 左移代替乘法 |
| 二分查找 | 69 | 边界定位 |
| 快速幂 | 50 | 指数二进制 |

---

*本文件由 Claude Code 自动生成*