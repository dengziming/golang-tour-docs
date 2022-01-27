package letcode

// 思路，如果 words 没有重复，只需要判断 s 从某个点开始接下来字符串 split 后是否和 words 一致
// 如果 words 有重复，那需要判断
func findSubstring(s string, words []string) []int {
	var result []int
	step := len(words[0])
	span := len(words) * step

	counts := makeMap(words)

	for i := range s {
		if i + span <= len(s) {
			substring := makeMap(makeStringArray(s[i:], step, len(words)))
			if equals(counts, substring) {
				result = append(result, i)
			}
		}
	}
	return result
}

func equals(a map[string]int, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for ak, av := range a {
		if bv,ok:=b[ak]; ok && av == bv {
			// nothing
		} else {
			return false
		}
	}
	return true
}

func makeMap(words []string) map[string]int {
	counts := make(map[string]int)
	for _,word := range words {
		if _,ok:=counts[word];ok {
			counts[word] = counts[word] + 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func makeStringArray(s string, step int, count int) []string {
	result := make([]string, count)
	for index := 0; index < count; index ++ {
		result[index] = s[index * step: (index + 1) * step]
	}
	return result
}
