package tree_array

import (
	"strconv"
	"testing"
)

func TestTree(t *testing.T) {
	tree := New(10)
	for i := 1; i <= 10; i++ {
		tree.Add(i, i)
	}
	if res := tree.GetSum(10); res != 55 {
		t.Error("get sum " + strconv.Itoa(10) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(1); res != 1 {
		t.Error("get sum " + strconv.Itoa(1) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(3); res != 6 {
		t.Error("get sum " + strconv.Itoa(3) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(5); res != 15 {
		t.Error("get sum " + strconv.Itoa(5) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(7); res != 28 {
		t.Error("get sum " + strconv.Itoa(7) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(9); res != 45 {
		t.Error("get sum " + strconv.Itoa(9) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(4); res != 10 {
		t.Error("get sum " + strconv.Itoa(4) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(6); res != 21 {
		t.Error("get sum " + strconv.Itoa(6) + " wrong result: " + strconv.Itoa(res))
	}
	if res := tree.GetSum(8); res != 36 {
		t.Error("get sum " + strconv.Itoa(8) + " wrong result: " + strconv.Itoa(res))
	}
}
