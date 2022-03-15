package algo

type UF interface {
	connected(p, q int) bool
	union(p, q int)
}

type QuickFindUF struct {
	id []int
}

func (uf QuickFindUF) connected(p, q int) bool {
	return uf.id[p] == uf.id[q]
}

func (uf QuickFindUF) union(p, q int) {
	pid := uf.id[p]
	qid := uf.id[q]

	for i := range uf.id {
		if uf.id[i] == pid {
			uf.id[i] = qid
		}
	}
}

type QuickUnionUF struct {
	id []int
}

func (uf QuickUnionUF) root(i int) int {
	for uf.id[i] != i {
		i = uf.id[i]
	}
	return i
}

func (uf QuickUnionUF) connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf QuickUnionUF) union(p, q int) {
	uf.id[p] = uf.id[q]
}

// WeightQuickUnionUF weighted
type WeightQuickUnionUF struct {
	id []int
	weight []int
}

func (uf WeightQuickUnionUF) root(i int) int {
	for uf.id[i] != i {
		// this can be optimized by compress path
		i = uf.id[i]
	}
	return i
}

func (uf WeightQuickUnionUF) connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

func (uf WeightQuickUnionUF) union(p, q int) {
	rp, rQ := uf.id[p], uf.id[q]

	if uf.weight[rp] > uf.weight[rQ] {
		uf.id[rQ] = rp
		uf.weight[rp] += uf.weight[rQ]
	} else {
		uf.id[rp] = rQ
		uf.weight[rQ] += uf.weight[rp]
	}
}
