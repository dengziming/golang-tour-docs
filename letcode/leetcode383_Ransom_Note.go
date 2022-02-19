package letcode

func canConstruct(ransomNote string, magazine string) bool {
	counts1 := make(map[int32]int)
	counts2 := make(map[int32]int)

	for _, char := range ransomNote {
		counts1[char]++
	}
	for _, char := range magazine {
		counts2[char]++
	}

	for char := range counts1 {
		if counts1[char] - counts2[char] > 0 {
			return false
		}
	}
	return true
}
