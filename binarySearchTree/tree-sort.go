package binarySearchTree

func TreeSort(k []int) []int {
	if len(k) <= 1 {
		return k
	}
	return CreateBST(k...).ToSlice()
}
