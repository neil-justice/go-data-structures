package binarySearchTree

type TreeSet struct {
	bst *BST
}

func (set *TreeSet) Insert(k int) {
	set.bst.Insert(k)
}

func (set *TreeSet) Contains(k int) bool {
	return set.bst.Contains(k)
}

func (set *TreeSet) Delete(k int) bool {
	return set.bst.Delete(k)
}

func (set *TreeSet) ForEach(fn func(int)) {
	set.bst.Traverse(fn)
}
