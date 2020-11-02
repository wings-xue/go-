package mytree

import "math"

// 定义二叉树
type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
	Msg   int
}

// 		1
// 	  2		3
// 4      5	  6
// New func create tree like up
func New() *Tree {
	t1 := &Tree{Val: 1}
	t2 := &Tree{Val: 2}
	t3 := &Tree{Val: 3}
	t4 := &Tree{Val: 4}
	t5 := &Tree{Val: 5}
	t6 := &Tree{Val: 6}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t3.Left = t5
	t3.Right = t6
	return t1
}

// 遍历
// 前序（递归）
// TreePath save tree traversal result
var TreePath = make([]int, 0)

func preprocedural(t *Tree) {
	if t == nil {
		return
	}
	// 当前节点的路径
	TreePath = append(TreePath, t.Val)
	preprocedural(t.Left)
	preprocedural(t.Right)
}

// 		1
// 	  2		3
// 4      5	  6
// New func create tree
// 中序（递归）
func midsequent(t *Tree) {
	if t == nil {
		return
	}
	// 当前节点的路径
	midsequent(t.Left)
	TreePath = append(TreePath, t.Val)
	midsequent(t.Right)
}

// 后续（递归）
func subsequent(t *Tree) {
	if t == nil {
		return
	}
	// 当前节点的路径
	subsequent(t.Left)
	subsequent(t.Right)
	TreePath = append(TreePath, t.Val)
}

// 递归总结，如果一个东西需要被用到，那就把它存起来，重要的什么条件存进去以及取出来以后做什么

// 前序(非递归)
func preproceduralv2(t *Tree) []int {
	treePath := make([]int, 0)
	stack := make([]*Tree, 0)
	for len(stack) != 0 || t != nil {
		if t == nil {
			t = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		treePath = append(treePath, t.Val)
		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		t = t.Left
	}
	return treePath

}

// 中序(非递归)
func midsequentv2(t *Tree) []int {
	treePath := make([]int, 0)
	stack := make([]*Tree, 0)
	for len(stack) != 0 || t != nil {
		if t == nil {
			t = stack[len(stack)-1]
			treePath = append(treePath, t.Val)
			t = t.Right
			stack = stack[0 : len(stack)-1]
		} else {
			for t != nil {
				stack = append(stack, t)
				t = t.Left
			}

		}

	}
	return treePath
}

// 后续(非递归)
func subsequentv2(t *Tree) []int {
	treePath := make([]int, 0)

	stack := make([]*Tree, 0)

	for len(stack) != 0 || t != nil {
		if t != nil {
			for t != nil {
				stack = append(stack, t)
				t = t.Left
			}
		} else {
			t = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			if t.Msg == 1 {
				treePath = append(treePath, t.Val)
				t = nil
			} else {
				if t.Right != nil {
					t.Msg = 1
					stack = append(stack, t)
				} else {
					treePath = append(treePath, t.Val)
				}
				t = t.Right
			}

		}
	}

	return treePath
}

// 深度（分治)
func DFS(t *Tree) []int {
	treePath := make([]int, 0)
	if t == nil {
		return treePath
	}
	leftPath := DFS(t.Left)
	rightPath := DFS(t.Right)
	treePath = append(treePath, t.Val)
	treePath = append(treePath, leftPath...)
	treePath = append(treePath, rightPath...)
	return treePath
}

// 深度（非分治）
func DFSv2(t *Tree, result *[]int) {
	if t == nil {
		return
	}
	*result = append(*result, t.Val)
	DFSv2(t.Left, result)
	DFSv2(t.Right, result)
}

// 层次遍历
func BFS(t *Tree) []int {
	treePath := make([]int, 0)
	// 通过当前层的信息获取下一层的信息
	// 扫描一层,全部入栈
	// 重新扫描出栈
	if t == nil {
		return treePath
	}

	stack := make([]*Tree, 0)
	stack = append(stack, t)
	for len(stack) != 0 {
		// 开始扫描一层
		c := len(stack)
		for i := 0; i < c; i++ {
			t = stack[0]
			stack = stack[1:]
			treePath = append(treePath, t.Val)
			if t.Left != nil {
				stack = append(stack, t.Left)
			}
			if t.Right != nil {
				stack = append(stack, t.Right)
			}
		}
	}

	return treePath
}

// 分治法快速排序

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func helper(nums []int, left, right int) int {

	if left == right || right-left == 1 {
		return left
	}
	nums = nums[left:right]
	right = right - 1
	direct := true
	// 左右两个图标
	for left != right {
		if nums[left] > nums[right] {
			swap(nums, left, right)
			direct = !direct
		}

		if direct {
			right--
		} else {
			left++
		}

	}
	return left
}

func quickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i := helper(nums, 0, len(nums))
	quickSort(nums[0:i])
	quickSort(nums[i+1:])

}

// 组合两个已经排序过得列表
// TODO: 不应该返回err
func merge(n1 []int, n2 []int) ([]int, error) {

	n := make([]int, len(n1)+len(n2))
	p1 := 0
	p2 := 0
	p := 0
	for p < len(n) {
		// TODO:可以优化条件
		switch {
		case p1 == len(n1):
			n[p] = n2[p2]
			p++
			p2++
		case p2 == len(n2):
			n[p] = n1[p1]
			p++
			p1++
		default:
			if n1[p1] < n2[p2] {
				n[p] = n1[p1]
				p++
				p1++
			} else {
				n[p] = n2[p2]
				p++
				p2++
			}

		}

	}
	return n, nil

}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			swap(nums, 0, 1)
		}
		return nums
	}

	// 一个array分成两部分
	i := int(math.Ceil(float64(len(nums)) / float64(2)))
	n1 := nums[0:i]
	n2 := nums[i:]

	m1 := mergeSort(n1)
	m2 := mergeSort(n2)
	// 对两部分分别排序
	n, _ := merge(m1, m2)
	// 合并两个部分
	return n
}
