package letcode

func partitionLabels(s string) (partition []int){

	lastPos := make(map[string]int)
	for i := 0; i < len(s); i++ {
		lastPos[s[i:i+1]] = i
	}

	start, end := 0, 0
	for i := range s {
		if lastPos[s[i:i+1]] > end {
			end = lastPos[s[i:i+1]]
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return

}
