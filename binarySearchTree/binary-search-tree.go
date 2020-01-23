package binarySearchTree

type BST struct {
	Left   *BST
	Key    int
	Right  *BST
	Parent *BST
}

func CreateBST(keys ...int) *BST {
	if len(keys) == 0 {
		panic("Cannot create empty BST")
	}

	root := &BST{nil, keys[0], nil, nil}
	for _, elem := range keys[1:] {
		root.Insert(elem)
	}
	return root
}

func (n *BST) leftOrRight(k int) *BST {
	if k < n.Key {
		return n.Left
	} else if k > n.Key {
		return n.Right
	} else {
		return nil
	}
}

func (n *BST) Insert(k int) {
	if k < n.Key {
		if n.Left != nil {
			n.Left.Insert(k)
		} else {
			n.Left = &BST{nil, k, nil, n}
		}
	} else if k > n.Key {
		if n.Right != nil {
			n.Right.Insert(k)
		} else {
			n.Right = &BST{nil, k, nil, n}
		}
	}
}

func (n *BST) find(k int) *BST {
	if n.Key == k {
		return n
	} else if child := n.leftOrRight(k); child != nil {
		return child.find(k)
	} else {
		return nil
	}
}

func (n *BST) findMin() *BST {
	for n != nil && n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *BST) children() []*BST {
	children := make([]*BST, 0, 2)
	if n.Left != nil {
		children = append(children, n.Left)
	}
	if n.Right != nil {
		children = append(children, n.Right)
	}
	return children
}

func (n *BST) Contains(k int) bool {
	found := n.find(k)
	return found != nil
}

func (n *BST) Delete(k int) bool {
	found := n.find(k)
	if found == nil {
		return false
	}

	parent := found.Parent
	children := found.children()

	if len(children) == 0 {
		if parent == nil {
			panic("Cannot delete root of 1-element tree")
		} else if parent.Left == found {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return true
	} else if len(children) == 1 {
		replacement := children[0]
		found.Key = replacement.Key
		found.Right = replacement.Right
		found.Left = replacement.Left
		return true
	} else {
		minRight := found.Right.findMin()
		found.Key = minRight.Key
		minRight.Delete(minRight.Key)
		return true
	}
}

func (n *BST) Traverse(fn func(k int)) {
	if n.Left != nil {
		n.Left.Traverse(fn)
	}
	fn(n.Key)
	if n.Right != nil {
		n.Right.Traverse(fn)
	}
}

func (n *BST) ToSlice() []int {
	a := make([]int, 0)
	n.Traverse(func(k int) { a = append(a, k) })
	return a
}

func (n *BST) BreadthFirstSearch(k int) bool {
	queue := make([]*BST, 1)
	queue[0] = n
	for _, elem := range queue {
		if elem.Key == k {
			return true
		} else {
			if elem.Left != nil {
				queue = append(queue, elem.Left)
			}
			if elem.Right != nil {
				queue = append(queue, elem.Right)
			}
		}
	}
	return false
}
