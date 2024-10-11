package base

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}
	extracted := h.array[0]
	// 为维持完全二叉树结构，把堆的最后一个节点拿出来
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.heapifyDown(0)
	return extracted
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		childToCompare = l
		if r <= lastIndex && h.array[r] < h.array[l] {
			childToCompare = r
		}
		if h.array[index] > h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// 建堆操作
// 创建一个空堆然后遍历列表，一次对每个元素执行入堆操作，即先将元素添加至堆的尾部，再对改元素执行从底到顶的堆化
// 由于节点是从顶到底依次被添加进二叉树的，因此堆是自上而下构建的
// 另外一种更高效的建堆方法是：
// 1. 将列表所有元素原封不动地添加到堆中，此时堆的性质尚未得到满足
// 2. 倒序遍历堆，依次对每个非叶子节点执行从顶到至底的堆化
// 每当堆化一个节点后，以该节点为根节点的子树就形成一个合法的子堆。而由于是倒序遍历，因此堆是“自下而上”构建的
func newMinHeap(nums []int) *MinHeap {
	// 将列表元素原封不动添加进堆
	h := &MinHeap{array: nums}
	for i := parent(len(h.array) - 1); i >= 0; i-- {
		h.heapifyUp(i)
	}
	return h
}
