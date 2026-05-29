# Linked Lists 链表 - LeetCode Deep Dive

## 目录
- [Problem 2: Add Two Numbers](#problem-2-add-two-numbers-两数相加)
- [Problem 19: Remove Nth Node](#problem-19-remove-nth-node-from-end-of-list)
- [Problem 21: Merge Two Sorted Lists](#problem-21-merge-two-sorted-lists)
- [Problem 24: Swap Pairs](#problem-24-swap-nodes-in-pairs)
- [Problem 83: Delete Duplicates](#problem-83-remove-duplicates-from-sorted-list)
- [Problem 92: Reverse Between](#problem-92-reverse-linked-list-ii)

---

## Problem 2: Add Two Numbers 两数相加

### 1. 题目核心与隐藏考点

**核心本质**: 逐位相加，使用进位carry模拟竖式加法，链表逆序存储正好符合加法顺序。

**隐藏考点**:
- 链表长度可能不同
- 最终可能产生额外的进位（如 999 + 1 = 1000）
- 需要dummy节点简化边界处理

```
链表表示:
  2 → 4 → 3  代表数字 342 (逆序存储!)
  5 → 6 → 4  代表数字 465

计算:
    342
  + 465
  ─────
    807

逐位相加过程:
  2 + 5 = 7,  进位 0
  4 + 6 = 10, 进位 1
  3 + 4 + 1 = 8, 进位 0
```

---

### 2. 思路演进

#### 解法：模拟加法（最优解）

```
l1: 2 → 4 → 3
l2: 5 → 6 → 4

dummy → 7 → 0 → 8
        ↑
      result

每步:
  sum = l1.val + l2.val + carry
  new_val = sum % 10
  carry = sum / 10
```

---

### 3. 工业级 Go 源码与详细注释

```go
package linked_lists

// ListNode 链表节点定义
type ListNode struct {
    Val  int
    Next *ListNode
}

// AddTwoNumbers 模拟竖式加法
//
// 核心思想：
// 1. 链表是逆序存储的，所以从头部开始正好是从低位到高位
// 2. 逐位相加，用 carry 记录进位
// 3. 使用 dummy 节点简化头节点处理
//
// 为什么需要 dummy 节点？
// - 结果链表的头可能变化（如最终进位产生新节点）
// - dummy.Next 永远指向结果链表的头
//
// 时间复杂度：O(max(m, n)) - m, n 为两个链表长度
// 空间复杂度：O(max(m, n)) - 忽略输出空间
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{} // 哑节点，简化边界处理
    current := dummy
    carry := 0

    // 遍历直到两个链表都为空且没有进位
    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry

        // 加上 l1 的值
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        // 加上 l2 的值
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        // 计算当前位的值和新的进位
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }

    return dummy.Next
}
```

---

## Problem 19: Remove Nth Node from End of List 删除链表的倒数第 N 个节点

### 1. 题目核心与隐藏考点

**核心本质**: 快慢指针，让快指针先走 n 步，然后一起移动，快指针到达末尾时，慢指针正好在倒数第 n 个节点。

**隐藏考点**:
- 删除头节点的特殊处理
- 需要 dummy 节点处理删除第一个节点的情况

```
示例: 删除倒数第 2 个节点

初始: 1 → 2 → 3 → 4 → 5, n = 2

Step 1: fast 先走 2 步
        dummy → 1 → 2 → 3 → 4 → 5
        fast
        dummy → 1 → 2 → 3 → 4 → 5
                    ↑     ↑
                   slow  fast

Step 2: 一起移动直到 fast 到达末尾
        dummy → 1 → 2 → 3 → 4 → 5
                            ↑     ↑
                          slow   fast (nil)

Step 3: slow.Next = slow.Next.Next 删除节点
        dummy → 1 → 2 → 3 → 5
```

---

### 2. 工业级 Go 源码与详细注释

```go
// RemoveNthFromEnd 删除倒数第 n 个节点
//
// 核心思想：快慢指针
// 1. 让 fast 先走 n+1 步
// 2. 然后 slow 和 fast 一起走
// 3. 当 fast 到达末尾时，slow 的下一个节点就是要删除的节点
//
// 为什么需要 n+1 而不是 n？
// - 要删除倒数第 n 个节点，需要找到它的前一个节点
// - fast 先走 n+1 步后，slow 和 fast 之间隔着 n 个节点
// - 这样当 fast 到达末尾时，slow 正好在目标节点的前一个位置
//
// 为什么需要 dummy？
// - 如果要删除的是第一个节点，没有前一个节点
// - dummy 让这个情况变得简单
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast := dummy
    slow := dummy

    // fast 先走 n+1 步
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // 一起移动
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    // 删除节点
    slow.Next = slow.Next.Next

    return dummy.Next
}
```

---

## Problem 21: Merge Two Sorted Lists 合并两个有序链表

### 1. 题目核心与隐藏考点

**核心本质**: 归并排序的merge步骤，逐个比较两个链表的节点，按顺序串接。

```
示例:
  l1: 1 → 2 → 4
  l2: 1 → 3 → 4

归并过程:
  比较 1 vs 1 → 选择 1
  比较 2 vs 1 → 选择 1
  比较 2 vs 3 → 选择 2
  比较 4 vs 3 → 选择 3
  选 4

结果: 1 → 1 → 2 → 3 → 4 → 4
```

---

### 2. 工业级 Go 源码与详细注释

```go
// MergeTwoLists 合并两个有序链表
//
// 核心思想：归并排序的 merge 步骤
// 1. 比较两个链表的头节点
// 2. 选取较小的节点接到结果链表
// 3. 移动选取节点的指针
// 4. 重复直到某个链表为空
// 5. 接入剩余的链表
//
// 时间复杂度：O(m + n)
// 空间复杂度：O(1) - 不计输出空间
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy

    // 同时遍历两个链表
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }

    // 接入剩余部分
    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }

    return dummy.Next
}
```

---

## Problem 24: Swap Nodes in Pairs 两两交换链表节点

### 1. 题目核心与隐藏考点

**核心本质**: 节点反转，但只反转相邻的两个节点，需要保存前后节点。

```
示例: 1 → 2 → 3 → 4

交换过程:
  Step 1: 交换 1 和 2
          dummy → 2 → 1 → 3 → 4

  Step 2: 交换 3 和 4
          dummy → 2 → 1 → 4 → 3

图解:
  before:  [prev] → [node1] → [node2] → ...
                  ↓
           swap: [prev] → [node2] → [node1] → ...
```

---

### 2. 工业级 Go 源码与详细注释

```go
// SwapPairs 两两交换相邻节点
//
// 核心思想：
// 1. 保存需要交换的两个节点：first = current.Next, second = first.Next
// 2. 反转这两个节点的指向：first.Next = second.Next, second.Next = first
// 3. 将 current 连接到 second：current.Next = second
// 4. 移动 current 到 first（下一个要处理的节点）
//
// 关键点：
// - 必须先保存 second.Next，因为会被 first 覆盖
// - 每次处理一对节点，需要记录前一个节点
func SwapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    current := dummy

    // 必须有两个节点可以交换
    for current.Next != nil && current.Next.Next != nil {
        first := current.Next
        second := current.Next.Next

        // 反转
        first.Next = second.Next
        second.Next = first

        // 连接
        current.Next = second

        // 移动到下一对
        current = first
    }

    return dummy.Next
}
```

---

## Problem 83: Remove Duplicates from Sorted List 删除排序链表的重复项

### 1. 题目核心与隐藏考点

**核心本质**: 单指针遍历，遇到重复就跳过。

```
示例: 1 → 1 → 2 → 3 → 3

处理过程:
  cur = 1 → 1 → 2 → 3 → 3
        ↑
  发现重复，cur.Next = cur.Next.Next

  cur = 1 → 2 → 3 → 3
          ↑
  不重复，移动cur
```

---

### 2. 工业级 Go 源码与详细注释

```go
// DeleteDuplicates 删除排序链表中的重复项
//
// 核心思想：
// - 链表已排序，重复元素一定相邻
// - 用 current 遍历，如果 current.Val == current.Next.Val
// - 说明有重复，跳过 current.Next
// - 否则，current 前进
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func DeleteDuplicates(head *ListNode) *ListNode {
    current := head

    for current != nil && current.Next != nil {
        if current.Val == current.Next.Val {
            // 发现重复，跳过
            current.Next = current.Next.Next
        } else {
            // 不重复，前进
            current = current.Next
        }
    }

    return head
}
```

---

## Problem 92: Reverse Linked List II 翻转链表 II

### 1. 题目核心与隐藏考点

**核心本质**: 部分翻转，需要定位翻转区间的前一个节点和后一个节点。

```
示例: 翻转 2-4 位
输入: 1 → 2 → 3 → 4 → 5, left=2, right=4
输出: 1 → 4 → 3 → 2 → 5

图解:
  定位: prev=1, start=2, end=4, next=5

  翻转 2→3→4:
  1 → [2 ← 3 ← 4] → 5
       ↑
     reverse

  结果: 1 → 4 → 3 → 2 → 5
```

---

### 2. 工业级 Go 源码与详细注释

```go
// ReverseBetween 翻转链表的指定区间
//
// 核心思想：
// 1. 找到翻转区间的前一个节点 prev
// 2. 记录翻转区间的起始节点 start
// 3. 翻转从 start 开始的 right-left 个节点
// 4. 将 prev.Next 指向翻转后的头，start.Next 指向原 next
//
// 具体步骤：
// - 先让 prev 走到 left-1 的位置
// - start = prev.Next，即翻转区间的第一个节点
// - 用标准的链表翻转方法翻转 right-left 次
// - 将 prev.Next 连接到翻转后的链表头
// - 将翻转链表的尾连接到原 next
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy

    // 1. 找到 left-1 位置的节点
    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }

    // 2. 开始翻转
    current := prev.Next
    for i := 0; i < right-left; i++ {
        next := current.Next
        current.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
    }

    return dummy.Next
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | 完全翻转 | Easy |
| [86. Partition List](https://leetcode.com/problems/partition-list/) | 双链表 | Medium |
| [328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/) | 分离再连接 | Medium |

---

## 链表题目总结

### 常见模式

| 模式 | 题目 | 关键点 |
| :--- | :--- | :--- |
| 快慢指针 | 19 | 先走 n 步 |
| 归并 | 21 | 比较 + 连接 |
| 两两交换 | 24 | 保存 next |
| 部分翻转 | 92 | 定位边界 |
| dummy 节点 | 全部 | 简化边界处理 |

---

*本文件由 Claude Code 自动生成*