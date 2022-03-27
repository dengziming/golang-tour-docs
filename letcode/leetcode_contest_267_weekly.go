package letcode


func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	//
	n := 12
	var dfs func(i int, score int, target[] int)

	maxV := 0
	var bobArrows []int
	dfs = func(i int, score int, target[] int) {
		if i == n {
			if score > maxV {
				maxV = score
				bobArrows = append([]int{}, target...)
				// 剩下的也加回去
				if numArrows != 0 {
					bobArrows[0]+=numArrows
				}
			}
			return
		}
		// 选择 0 只
		dfs(i+1, score, append(target, 0))

		if numArrows >= aliceArrows[i]+1 {
			numArrows -= aliceArrows[i]+1
			dfs(i+1, score+i, append(target, aliceArrows[i]+1))
			numArrows += aliceArrows[i]+1
		}
	}
	dfs(0, 0, []int{})
	return bobArrows
}

func countCollisions(directions string) int {
	sum := 0

	// 没有碰撞的，往右边的车辆数
	right := 0
	for _, direction := range directions {
		if direction == 'R' {
			right = right +1
		} else {
			sum += right
			right = 0
		}
	}

	// 没有碰撞的，往左边的车辆数
	left := 0
	for i := len(directions) - 1; i >= 0; i-- {
		if directions[i] == 'L' {
			left = left +1
		} else {
			sum += left
			left = 0
		}
	}

	return sum
}
