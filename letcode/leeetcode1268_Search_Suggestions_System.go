package letcode

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	// 需要找到 >= word 的三个数
	// 我们找第一个大于等于 word 的下标，下标往后三个数进行判断即可
	binarySearch := func(left, right int, word string) int {
		for left != right {
			mid := (left+right)/2
			// final : [3,5] 5
			// final : [3] 5
			if products[mid] >= word {
				right = mid
			} else {
				left = mid+1
			}
		}
		return right
	}

	commonPrefix := func(s1, search string) bool {
		if len(s1) < len(search){
			return false
		}
		return search == s1[:len(search)]
	}

	result := make([][]string, len(searchWord))
	left, right := 0, len(products) - 1
	for i := 0; i < len(searchWord); i++ {
		result[i] = make([]string, 0)
		idx := binarySearch(left, right, searchWord[:i+1])
		// 小优化
		left = idx

		for j := idx; j < len(products) && j < idx + 3; j++ {
			if commonPrefix(products[j], searchWord[:i+1]) {
				result[i] = append(result[i], products[j])
			} else {
				// 小优化
				break
			}
		}
	}
	return result
}
