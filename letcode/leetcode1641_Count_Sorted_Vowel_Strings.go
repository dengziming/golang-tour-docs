package letcode

func countVowelStrings(n int) int {
	
	chars := []string{"a", "e", "i", "o", "u"}
	
	result := []string{""}
	for i := 0; i < n; i++ {
		tmp := make([]string, 0)
		for _, pre := range result {
			for _, char := range chars {
				if len(pre) == 0 || char >= pre[len(pre)-1:] {
					tmp = append(tmp, char)
				}
			}
		}
		result = tmp
	}
	return len(result)
}
