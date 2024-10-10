package base

import (
	"container/list"
	"fmt"
)

// 后面直接使用官方库list当作队列或栈
// 层序遍历

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   v,
	}
}

func levelOrder(root *TreeNode) []any {
	// 层序遍历利用队列FIFO的特性实现
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片用于保存遍历序列
	nums := make([]any, 0)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return nums
}

var nums = make([]any, 0)

// 层序遍历的传统实现使用队列。在递归实现层序遍历时，我们可能需要将当前层的节点存储在一个列表中，然后递归处理下一层
func levleOrderByRecursion(root *TreeNode) [][]any {
	result := [][]any{}
	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) == level {
			result = append(result, []any{})
		}
		result[level] = append(result[level], root.Val)
		traverse(root.Left, level+1)
		traverse(root.Right, level+1)
	}

	return result
}

func preOrder(root *TreeNode) []any {
	// 前序遍历 根节点、左子树、右子树
	// 递归
	if root == nil {
		return nil
	}
	nums = append(nums, root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
	return nil
}

func inOrder(root *TreeNode) []any {
	// 中序遍历 左子树、根节点、右子树
	// 递归
	if root == nil {
		return nil
	}
	preOrder(root.Left)
	nums = append(nums, root.Val)
	preOrder(root.Right)
	return nil
}

func postOrder(root *TreeNode) []any {
	// 后序遍历 左子树、右子树、根节点
	// 递归
	if root == nil {
		return nil
	}
	preOrder(root.Left)
	preOrder(root.Right)
	nums = append(nums, root.Val)
	return nil
}

// 深度优先遍历的非递归方式实现
// 绝大多数可以用递归解决的问题，其实都可以用栈来解决。因为递归和栈都有回溯的特性

func preOrderTraversalWithStack(root *TreeNode) {
	// 前序遍历 根节点、左子树、右子树
	// 非递归
	// 每次出栈都是因为没有左右孩子节点或者都已经访问过了
	if root == nil {
		return
	}
	node := root
	stack := list.New()
	for stack.Len() > 0 && node != nil {
		for node != nil {
			fmt.Println(node.Val)
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() > 0 {
			node = stack.Remove(stack.Back()).(*TreeNode)
			node = node.Right
		}
	}
}

func preOrderTraversalWithStack1(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		fmt.Println(node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
}

func inOrderTraversalWithStack(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	node := root
	for stack.Len() > 0 && node != nil {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		if stack.Len() > 0 {
			node = stack.Remove(stack.Back()).(*TreeNode)
			fmt.Println(node.Val)

			node = node.Right
		}
	}
}

func postOrderTraversalWithStack(root *TreeNode) {
	if root == nil {
		return
	}

	stack := list.New()
	var lastVisited *TreeNode
	node := root
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		// 查看栈顶元素
		peek := stack.Back().Value.(*TreeNode)

		if peek.Right != nil && peek.Right != lastVisited {
			node = peek.Right
		} else {
			fmt.Println(peek.Val)
			lastVisited = peek
			stack.Remove(stack.Back())
		}
	}
}

// 二叉搜索树的删除操作

func (bst *binarySearchTree) remove(num int) {
	cur := bst.root
	// 若树为空，直接提前返回
	if cur == nil {
		return
	}
	// 待删除之前的节点位置
	var pre *TreeNode = nil
	// 循环查找，越过叶节点后跳出
	for cur != nil {
		if cur.Var == num {
			break
		}
		pre = cur
		if cur.Var.(int) < num {
			// 待删除节点在右子树中
			cur = cur.Right
		} else {
			// 待删除节点在左子树中
			cur = cur.Left
		}
	}
	// 若无待删除节点，则直接返回
	if cur == nil {
		return
	}
	// 子节点数为0或1
	if cur.Left == nil || cur.Right == nil {
		var child *TreeNode = nil
		if cur.Left != nil {
			child = cur.Left
		} else {
			child = cur.Right
		}
		// 删除节点cur
		if cur != bst.root {
			if pre.Left == cur {
				pre.Left = child
			} else {
				pre.Right = child
			}
		} else {
			// 若删除节点为根节点，则重新指定根节点
			bst.Right = child
		}
		// 子节点为2
	} else {
		// 获取中序遍历中待删除节点cur的下一个节点
		tmp := cur.Right
		for tmp.Left != nil {
			tmp = tmp.Left
		}
		// 递归删除节点tmp
		bst.remove(tmp.Var.(int))
		// 将tmp的值替换为cur
		cur.Var = tmp.Var
	}
}

// 递归删除方法

func (bst *binarySearchTree) Remove(var int) {
  bst.root = bst.removeNode(bst.root, var)
}

func (bst *binarySearchTree) removeNode(node *TreeNode, val int) *TreeNode {
  if node == nil {
    return node
  }
  if val < node.Var.(int) {
    node.Left = bst.removeNode(node.Left, val)
  } else if val > node.Var.(int) {
    node.Right = bst.removeNode(node.Right, val)
  } else {
    if node.Left == nil && node.Right == nil {
      return nil
    } else if node.Left == nil {
      return node.Right
    } else if node.Right == nil {
      return node.Left
    }
    minRight := bst.findMin(node.Right)
    node.Val = minRight.Val
    node.Right = bst.removeNode(node.Right, minRight.Val)
  }
  return node
}

func (bst *binarySearchTree) findMin(node *TreeNode) *TreeNode {
  if node.Left == nil {
    return node
  }
  return bst.findMin(node.Left)
}

