package exercise

/**
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。



示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45


限制：

1 <= n <= 10000
*/
func sumNums(n int) int {
	result := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		result += n
		//逻辑与, A && B,只有A为true时,才会执行判断B;当A为false时,不需要再判断B的值,借此来控制递归的进行
		return n > 0 && sum(n-1)
	}
	sum(n)
	return result
}
