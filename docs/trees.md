# Trees 树 - LeetCode Deep Dive

## 目录
- [Problem 94: Binary Tree Inorder Traversal](#problem-94-binary-tree-inorder-traversal-二叉树的中序遍历)
- [Problem 98: Validate Binary Search Tree](#problem-98-validate-binary-search-tree-验证二叉搜索树)
- [Problem 100: Same Tree](#problem-100-same-tree-相同的树)

---

## Problem 94: Binary Tree Inorder Traversal 二叉树的中序遍历

### 1. 题目核心与隐藏考点

**核心本质**: 「左-根-右」的递归遍历顺序，递归的本质是深度优先搜索（DFS）。

**隐藏考点**:
- 递归 vs 迭代（栈模拟递归）
- Morris 遍历 O(1) 空间
- 递归的隐式栈 vs 显式栈

```
中序遍历图解:

       1
      / \
     2   3
    / \
   4   5

遍历顺序: 左 → 根 → 右
结果: 4 → 2 → 5 → 1 → 3

递归调用栈:
  inorder(1)
    → inorder(2)
      → inorder(4)
        → 4 无左子树，输出 4
        → 4 无右子树，返回
      → 输出 2
      → inorder(5)
        → 5 无左子树，输出 5
        → 5 无右子树，返回
      → 返回
    → 输出 1
    → inorder(3)
      → 3 无左子树，输出 3
      → 3 无右子树，返回
```

---

### 2. 思路演进

#### 解法一：递归（最简洁）
```go
func inorderTraversal(root *TreeNode) []int {
    var result []int
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        result = append(result, node.Val)
        inorder(node.Right)
    }
    inorder(root)
    return result
}
```

#### 解法二：迭代（显式栈）
```go
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    stack := []*TreeNode{}
    curr := root

    for curr != nil || len(stack) > 0 {
        // 一直向左走，压入栈
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        // 弹出栈顶，访问
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, curr.Val)

        // 转向右子树
        curr = curr.Right
    }

    return result
}
```

---

### 3. 工业级 Go 源码与详细注释

```go
package trees

// TreeNode 二叉树节点定义
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// InorderTraversal 中序遍历（递归版本）
//
// 核心思想：深度优先搜索 (DFS)
//
// 遍历顺序：左 → 根 → 右
//
// 为什么是「左根右」而不是其他顺序？
// - 对于 BST，中序遍历可以得到有序序列
// - 左子树所有节点小于根，根小于右子树
// - 中序遍历正好按升序输出
//
// 递归的本质：
// - 函数调用栈隐式保存了「下一步去哪里」的信息
// - 每个节点被访问三次：第一次进入（处理左子树）、
//   第二次从左子树返回（处理根）、第三次从右子树返回
//
// 时间复杂度：O(n) - 每个节点访问一次
// 空间复杂度：O(h) - h 为树的高度，递归栈的深度
func InorderTraversal(root *TreeNode) []int {
    var result []int

    // 递归函数
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        // 先遍历左子树
        inorder(node.Left)

        // 处理当前节点
        result = append(result, node.Val)

        // 再遍历右子树
        inorder(node.Right)
    }

    inorder(root)
    return result
}
```

---

## Problem 98: Validate Binary Search Tree 验证二叉搜索树

### 1. 题目核心与隐藏考点

**核心本质**: BST 的性质：左子树所有节点 < 根 < 右子树所有节点，且递归成立。

**隐藏考点**:
- 不能只比较左右子节点与根，要确保整个左子树都小于根
- 需要传递有效的上界和下界
- 递归时的上下界传递

```
验证过程图解:

       2
      / \
     1   3
    / \
   4   5   ← 这是 BST 吗？不是！

验证节点 4:
- 当前节点值 = 4
- 左边界 = 1, 右边界 = 2
- 4 >= 右边界(2)，违反 BST 性质!

正确 BST:
       2
      / \
     1   3      ← 1 < 2 < 3，有效

验证过程:
- 根 2：范围 (-∞, +∞)
- 左 1：范围 (-∞, 2)，1 在范围内 ✓
- 右 3：范围 (2, +∞)，3 在范围内 ✓
```

---

### 2. 工业级 Go 源码与详细注释

```go
// IsValidBST 验证二叉搜索树
//
// 核心思想：递归 + 上下界传递
//
// BST 的性质：
// - 左子树中所有节点的值都小于根节点
// - 右子树中所有节点的值都大于根节点
// - 且这个性质对所有子树都成立
//
// 递归思路：
// - 对于每个节点，传递一个有效范围 [min, max]
// - 如果节点值不在这个范围内，不是 BST
// - 递归时，左子节点的范围变为 [min, node.Val]
// - 递归时，右子节点的范围变为 [node.Val, max]
//
// 为什么需要上下界？
// - 只比较左右子节点与根不够
// - 考虑这样的情况：根=2，左子节点=1，右子节点=3，看起来没问题
// - 但左子节点的右子树可能是 1.5，仍然合法
// - 关键是：左子树的最大值必须小于根，右子树的最小值必须大于根
//
// 时间复杂度：O(n) - 每个节点访问一次
// 空间复杂度：O(h) - h 为树的高度
func IsValidBST(root *TreeNode) bool {
    // 递归函数，min 和 max 表示当前节点的有效范围
    var validate func(node, min, max *TreeNode) bool
    validate = func(node, min, max *TreeNode) bool {
        // 空树是有效的 BST
        if node == nil {
            return true
        }

        // 检查节点值是否在有效范围内
        // 条件：min != nil && node.Val <= min.Val
        // 解释：如果有下界限制且节点值小于等于下界，则无效
        if min != nil && node.Val <= min.Val {
            return false
        }
        if max != nil && node.Val >= max.Val {
            return false
        }

        // 递归检查左子树和右子树
        // 左子树的节点值必须在 (min, node.Val) 范围内
        // 右子树的节点值必须在 (node.Val, max) 范围内
        return validate(node.Left, min, node) && validate(node.Right, node, max)
    }

    return validate(root, nil, nil)
}
```

---

## Problem 100: Same Tree 相同的树

### 1. 题目核心与隐藏考点

**核心本质**: 递归比较两棵树的对应节点。

**隐藏考点**:
- 为什么是「与」而不是「或」？
- 空树如何比较？

```
比较过程图解:

p:       1       q:       1
       / \             / \
      2   3           2   3

递归比较:
- p.Val == q.Val? 1 == 1 ✓
- 比较左子树: p.Left(2) vs q.Left(2)
  - 2 == 2 ✓
  - 比较 2 的左子树: nil vs nil ✓
  - 比较 2 的右子树: nil vs nil ✓
- 比较右子树: p.Right(3) vs q.Right(3)
  - 3 == 3 ✓
  - ...

结果: 两棵树相同 ✓
```

---

### 2. 工业级 Go 源码与详细注释

```go
// IsSameTree 判断两棵二叉树是否相同
//
// 核心思想：深度优先搜索 (DFS) - 递归比较
//
// 比较逻辑：
// 1. 两棵树都为空 → 相同
// 2. 一棵树为空，另一棵非空 → 不同
// 3. 两棵树都非空但值不同 → 不同
// 4. 两棵树都非空且值相同 → 递归比较左右子树
//
// 为什么用 && 连接？
// - 只有当左右子树都相同时，两棵树才相同
// - 如果左子树相同但右子树不同，结果应该是不同
//
// 时间复杂度：O(n) - 每个节点最多访问一次
// 空间复杂度：O(h) - h 为树的高度，递归栈深度
func IsSameTree(p *TreeNode, q *TreeNode) bool {
    // 情况1：两棵树都为空
    if p == nil && q == nil {
        return true
    }

    // 情况2：只有一棵树为空
    if p == nil || q == nil {
        return false
    }

    // 情况3：两棵树都非空，但值不同
    if p.Val != q.Val {
        return false
    }

    // 情况4：值相同，递归比较左右子树
    // 只有左右子树都相同，才认为树相同
    return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
```

---

## 举一反三

| 相似题目 | 核心思想 | 难度 |
| :--- | :--- | :--- |
| [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | 递归比较 | Easy |
| [104. Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | 递归/迭代 | Easy |
| [110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) | 递归 | Easy |

---

## 树题目总结

### 常见遍历模式

| 遍历 | 顺序 | BST 性质 |
| :--- | :--- | :--- |
| 前序 | 根-左-右 | - |
| 中序 | 左-根-右 | 有序 |
| 后序 | 左-右-根 | - |
| 层序 | 按层 | - |

### 递归解题模板

```
func solution(root *TreeNode) {
    if root == nil {
        return // base case
    }

    // 处理当前节点
    // process(root.Val)

    // 递归处理左右子树
    solution(root.Left)
    solution(root.Right)
}
```

---

*本文件由 Claude Code 自动生成*