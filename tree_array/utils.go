package tree_array

type Tree struct {
	arr []int
}

func New(length int) *Tree {
	return &Tree{
		arr: make([]int, length+1),
	}
}

func lowbit(x int) int {
	return x & (-x)
}

func bitindex(x int) int {
	return x & (x - 1)
}

func (t *Tree) Add(n int, value int) {
	for i := n; i <= len(t.arr); i += lowbit(i) {
		t.arr[i] += value
	}
}

func (t *Tree) GetSum(n int) (sum int) {
	for n > 0 {
		sum += t.arr[n]
		n = bitindex(n)
	}
	return
}
