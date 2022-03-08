package letcode

func processQueries(queries []int, m int) []int {
	p := make([]int, m)
	for i := range p {
		p[i] = i + 1
	}
	var r []int
	for i := range queries {
		for j := range p {
			if queries[i] == p[j] {
				r = append(r, j)
				t:=[]int{p[j]}
				p = append(t,append(p[:j],p[j+1:]...)...)
				break
			}
		}
	}
	return r
}
