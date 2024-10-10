package base

type TreeNode1 struct {
	Val    int
	Height int
	Left   *TreeNode1
	Right  *TreeNode1
}

func NewTreeNode1(v int) *TreeNode1 {
	return &TreeNode1{
		Val:    v,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}
}

func (t *aVLTree) height(n *TreeNode1) int {
	// 空节点高度为-1， 叶节点高度为0
	if n == nil {
		return 0
	}
	return n.Height
}

func (t *aVLTree) updateHeight(n *TreeNode1) {
	n.Height = 1 + max(t.height(n.Left), t.height(n.Right))
}

// 获取节点平衡因子

func (t *aVLTree) balanceFactor(n *TreeNode1) int {
	if n == nil {
		return 0
	}
	// 节点平衡因子 = 左子树高度 - 右子树高度
	return t.height(n.Left) - t.height(n.Right)
}

// 右旋操作

func (t *aVLTree) rightRotate(n *TreeNode1) *TreeNode1 {
	child := n.Left
	grandChild := child.Right
	// 以child为原点，将node向右旋转
	child.Right = n
	n.Left = grandChild
	// 更新节点高度
	t.updateHeight(n)
	t.updateHeight(child)
	return child
}

// 左旋操作

func (t *aVLTree) leftRotate(n *TreeNode1) *TreeNode1 {
	child := n.Right
	grandChild := child.Left
	// 以child为原点，将node向左旋转
	child.Left = n
	n.Right = grandChild
	// 更新节点高度
	t.updateHeight(n)
	t.updateHeight(child)
	return child
}

// 旋转操作

func (t *aVLTree) rotate(node *TreeNode1) *TreeNode1 {
	// 获取节点的平衡因子
	bf := t.balanceFactor(node)
	// 左偏树
	if bf > 1 {
		if t.balanceFactor(node.Left) >= 0 {
			// 右旋
			return t.rightRotate(node)
		} else {
			// 先左旋后右旋
			node.Left = t.leftRotate(node.Left)
			return t.rightRotate(node)
		}
	}
	// 右偏树
	if bf < -1 {
		if t.balanceFactor(node.Right) <= 0 {
			// 左旋
			return t.leftRotate(node)
		} else {
			// 先右旋后左旋
			node.Right = t.rightRotate(node.Right)
			return t.leftRotate(node)
		}
	}
	return node
}

// 插入节点
func (t *aVLTree) insert(val int) {
	t.root = t.insertHelper(t.root, val)
}

// 递归插入节点
// 自底向上旋转
func (t *aVLTree) insertHelper(node *TreeNode1, val int) *TreeNode1 {
	if node == nil {
		return NewTreeNode1(val)
	}
	// 查找插入位置并插入节点
	if val < node.Val {
		node.Left = t.insertHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.insertHelper(node.Right, val)
	} else {
		return node
	}
	// 更新节点高度
	t.updateHeight(node)
	return t.rotate(node)
}

func (t *aVLTree) remove(val int) {
	t.root = t.removeHelper(t.root, val)
}

func (t *aVLTree) removeHelper(node *TreeNode1, val int) *TreeNode1 {
	if node == nil {
		return nil
	}
	if val < node.Val {
		node.Left = t.removeHelper(node.Left, val)
	} else if val > node.Val {
		node.Right = t.removeHelper(node.Right, val)
	} else {
		if node.Left == nil || node.Right == nil {
			child := node.Left
			if node.Right != nil {
				child = node.Right
			}
			if child == nil {
				return nil
			} else {
				node = child
			}
		} else {
			temp := node.Right
			for temp.Left != nil {
				temp = temp.Left
			}
			node.Right = t.removeHelper(node.Right, temp.Val)
			node.Val = temp.Val
		}
	}
	t.updateHeight(node)
	return t.rotate(node)
}
