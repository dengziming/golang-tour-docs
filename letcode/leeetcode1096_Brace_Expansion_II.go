package letcode

import "sort"

func braceExpansionII(expression string) []string {
	var bfs func(expr string) []string

	exit := func(expr string) bool {
		for i := range expr {
			if expr[i] == ',' {
				return false
			}
		}
		return true
	}

	bfs = func(expression string) []string {
		if exit(expression) {
			return []string{expression}
		}
		exprs := splitAnd(expression)

		var currentList = []string{""}
		for i := 0; i < len(exprs); i++ {
			var nextList []string
			expr := splitOr(exprs[i])
			for _, a := range currentList {
				for _, token := range expr {
					for _, b := range bfs(token) {
						nextList = append(nextList, a + b)
					}
				}
			}
			if i == len(exprs)-1 {
				return nextList
			} else {
				currentList = nextList
			}
		}

		return nil
	}

	list := bfs(expression)
	result := make([]string, 0)
	m := make(map[string]bool)
	for _, e := range list {
		if !m[e] {
			m[e] = true
			result = append(result, e)
		}

	}
	sort.Strings(result)
	return result
}

// 拆分成多个 expr，每个 expr 连接在一起
func splitAnd(expr string) []string {
	result := make([]string, 0)

	tmp := 0
	flag := 0 // 上一个段
	for i := range expr {
		if expr[i] == '{' {
			if tmp == 0 && flag != i {
				// 从 0 到 1，需要检查是否可以分段
				result = append(result, expr[flag:i])
				flag = i
			}
			tmp++
		} else if expr[i] == '}' {
			tmp--
			// 从 1 到 0 需要分段
			if tmp == 0 {
				// 注意这里去掉了首尾 { } result = append(result, expr[flag + 1:i])
				result = append(result, expr[flag:i + 1])
				flag = i + 1
			}
		} else if expr[i] == ',' {

		}
	}
	// 最后完还要检查
	if flag != len(expr) {
		result = append(result, expr[flag:])
	}
	return result
}

// 拆分成多个 expr，每个 expr 是任选一个的关系
func splitOr(expression string) []string {
	expr := expression
	if expression[0] == '{' {
		expr = expression[1:len(expression)-1]
	}
	result := make([]string, 0)
	tmp := 0
	flag := 0 // 上一个段
	for i := range expr {
		if expr[i] == '{' {
			tmp++
		} else if expr[i] == '}' {
			tmp--
		} else if expr[i] == ',' && tmp == 0 {
			if flag != i {
				result = append(result, expr[flag: i])
			}
			flag = i + 1
		}
	}
	if flag != len(expr) {
		result = append(result, expr[flag:])
	}
	return result
}
