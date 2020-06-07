package main

import "fmt"

func main() {
	//res := generateParenthesis(2)
	res := generateParenthesis(1)
	fmt.Print(res)
}

/*
回溯法，也是深度优先搜索算法
*/
func generateParenthesis(n int) []string {
	var res []string = make([]string, 0)
	dfs(n, n, "", &res)
	return res
}

func dfs(n1 int, n2 int, str string, res *[]string) {
	if n1 == 0 && n2 == 0 {
		*res = append(*res, str)
		return
	}

	if n1 > 0 {
		str = str + "("
		n1 = n1 - 1
		dfs(n1, n2, str, res)
		n1 = n1 + 1
		str = str[0 : len(str)-1]
		// 上面五句可以直接写成dfs(n1 - 1, n2, str + "(")，但是这样写对理解会造成困扰，因为其通过传参默认包含了回溯操作
		// 但是，写成一句话的形式可以大幅提高性能
	}

	if n1 < n2 {
		str = str + ")"
		n2 = n2 - 1
		dfs(n1, n2, str, res)
		n2 = n2 + 1
		str = str[0 : len(str)-1]
	}
}
