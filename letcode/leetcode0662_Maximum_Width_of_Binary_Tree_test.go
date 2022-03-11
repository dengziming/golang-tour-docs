package letcode

import "testing"

func TestWidthOfBinaryTree(t *testing.T) {
	if widthOfBinaryTree(&TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil},&TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{9, nil, nil}}}) != 4 {
		t.Error("测试失败")
	}
}
