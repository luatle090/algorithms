package main

import (
	"fmt"

	"github.com/algorithms/leetcode"
)

func main() {
	leetcode.CanConstruct("aa", "baa")

	fmt.Println()
	root := leetcode.InitTreeNodeTest()
	min := leetcode.Dfs(root)
	fmt.Println(min)
}
