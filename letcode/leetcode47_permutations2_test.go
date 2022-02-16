package letcode
import "testing"

func TestPermuteUnique(t *testing.T) {
	if len(permuteUnique([]int{1})) != 1 {
		t.Error("测试失败")
	}

	if len(permuteUnique([]int{1,2})) != 2 {
		t.Error("测试失败")
	}

	if len(permuteUnique([]int{1,2,3})) != 6 {
		t.Error("测试失败")
	}

	if len(permuteUnique([]int{1,1,2})) != 3 {
		t.Error("测试失败")
	}
}