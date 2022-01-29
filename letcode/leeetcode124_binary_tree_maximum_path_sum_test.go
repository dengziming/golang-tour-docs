package letcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	if maxPathSum(&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}) !=6 {
		t.Error("测试失败")
	}

	if maxPathSum(&TreeNode{-10, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}) !=42 {
		t.Error("测试失败")
	}

	if maxPathSum(&TreeNode{-3, nil, nil}) != -3 {
		t.Error("测试失败")
	}

	// [5,4,8,11,null,13,4,7,2,null,null,null,1]
	if maxPathSum(&TreeNode{
		5,
		&TreeNode{4,
			&TreeNode{11,
				&TreeNode{7, nil, nil},
				&TreeNode{2, nil, nil}}, nil},
		&TreeNode{8,
			&TreeNode{13, nil, nil},
			&TreeNode{4,
				nil,
				&TreeNode{1, nil, nil}}}}) != 48 {
		t.Error("测试失败")
	}
}