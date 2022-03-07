package letcode

import (
	"fmt"
	"testing"
)

func TestBraceExpansionII(t *testing.T) {
	fmt.Printf("%v\n", splitOr("a{x,ia,o}w"))
	fmt.Printf("%v\n", braceExpansionII("{a{x,ia,o}w,{n,{g,{u,o}},{a,{x,ia,o},w}},er}"))

	if len(braceExpansionII("{a,b}{c,{d,e}}")) != 6 {
		fmt.Printf("%v", braceExpansionII("{a,b}{c,{d,e}}"))
		t.Error("测试失败")
	}

	if len(braceExpansionII("{{a,z},a{b,c},{ab,z}}")) != 4 {
		fmt.Printf("%v", braceExpansionII("{{a,z},a{b,c},{ab,z}}"))
		t.Error("测试失败")
	}


	fmt.Printf("%v\n", braceExpansionII("a{x,ia,o}w"))
	fmt.Printf("%v\n", braceExpansionII("{n,{g,{u,o}}"))
	fmt.Printf("%v\n", braceExpansionII("{a{x,ia,o}w,{n,{g,{u,o}},{a,{x,ia,o},w}},er}"))

}
