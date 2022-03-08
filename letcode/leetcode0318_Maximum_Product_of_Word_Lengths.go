package letcode

func maxProduct2(words []string) int {

	product := func(word1, word2 string) int {
		counts := make([]bool, 26)
		for i := range word1 {
			counts[word1[i]-'a']=true
		}
		for i := range word2 {
			if counts[word2[i]-'a'] {
				return 0
			}
		}
		return len(word1) * len(word2)
	}

	products := func(word string, words []string) int {
		max := 0
		for i := range words {
			tmp := product(word, words[i])
			if tmp > max {
				max = tmp
			}
		}
		return max
	}

	max := 0
	for i := 0; i < len(words) -1; i++ {
		tmp := products(words[i], words[i+1:])
		if tmp > max {
			max = tmp
		}
	}
	return max
}
