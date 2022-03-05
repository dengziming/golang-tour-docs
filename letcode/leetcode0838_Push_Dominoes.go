package letcode

func pushDominoes(dominoes string) string {

	all := make([]int, 0)
	for idx, char := range dominoes {
		if char == 'L' {
			all = append(all, idx)
		} else if char == 'R' {
			all = append(all, idx)
		}
	}

	if len(all) == 0 {
		return dominoes
	}

	s := []byte(dominoes)
	// 第一个区间
	if dominoes[all[0]] == 'L' {
		for j := 0; j <= all[0] - 1; j++ {
			s[j] = 'L'
		}
	}
	for i := 0; i < len(all); i++ {
		// 最后一个区间
		if i + 1 == len(all) {
			if dominoes[all[i]] == 'R' {
				for j := all[i] + 1; j < len(dominoes); j++ {
					s[j] = 'R'
				}
			}
			continue
		}

		left := all[i] + 1
		right :=  all[i+1] - 1
		if dominoes[all[i]] == 'R' {
			//  / /
			if dominoes[all[i+1]] == 'R' {
				for j := left; j <= right; j++ {
					s[j] = 'R'
				}
			} else { // / \
				for j := left; j < (left + right +1)/2; j++ {
					s[j] = 'R'
				}
				for j := (left + right)/2 + 1; j <= right; j++ {
					s[j] = 'L'
				}
			}
		} else if dominoes[all[i+1]] == 'L'{
			for j := left; j <= right; j++ {
				s[j] = 'L'
			}
		}
	}
	return string(s)
}
