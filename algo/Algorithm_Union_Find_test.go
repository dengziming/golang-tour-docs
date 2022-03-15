package algo

import "testing"

func createId() []int {
	p := make([]int, 10)
	for i := 0; i < len(p); i++ {
		p[i] = i
	}
	return p
}

func createWeight() []int {
	p := make([]int, 10)
	for i := 0; i < len(p); i++ {
		p[i] = 1
	}
	return p
}

func TestQuickFindUF(t *testing.T) {
	uf := QuickFindUF{createId()}

	uf.union(1, 3)
	uf.union(3,8)
	uf.union(8,4)
	uf.union(4,9)

	if !uf.connected(1, 9){
		t.Error("测试失败")
	}

	if uf.connected(1, 7){
		t.Error("测试失败")
	}
}

func TestQuickUnionUF(t *testing.T) {
	uf := QuickUnionUF{createId()}

	uf.union(1, 3)
	uf.union(3,8)
	uf.union(8,4)
	uf.union(4,9)

	if !uf.connected(1, 9){
		t.Error("测试失败")
	}

	if uf.connected(1, 7){
		t.Error("测试失败")
	}
}

func TestWeightQuickUnionUF(t *testing.T) {
	uf := WeightQuickUnionUF{createId(), createWeight()}

	uf.union(1, 3)
	uf.union(3,8)
	uf.union(8,4)
	uf.union(4,9)

	if !uf.connected(1, 9){
		t.Error("测试失败")
	}

	if uf.connected(1, 7){
		t.Error("测试失败")
	}
}
