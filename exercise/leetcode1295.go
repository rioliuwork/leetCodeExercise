package exercise

import (
	"math"
	"strconv"
)

/**
给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。



示例 1：

输入：nums = [12,345,2,6,7896]
输出：2
解释：
12 是 2 位数字（位数为偶数）
345 是 3 位数字（位数为奇数）
2 是 1 位数字（位数为奇数）
6 是 1 位数字 位数为奇数）
7896 是 4 位数字（位数为偶数）
因此只有 12 和 7896 是位数为偶数的数字
示例 2：

输入：nums = [555,901,482,1771]
输出：1
解释：
只有 1771 是位数为偶数的数字。


提示：

1 <= nums.length <= 500
1 <= nums[i] <= 10^5
*/
func FindNumbers(nums []int) int {
	total := 0
	var a int
	var temp int
	for _, num := range nums {
		a = 10
		temp = 1
		for temp != 0 {
			temp = num / a
			if temp > 0 && temp < 10 {
				total++
			}
			a *= 100
		}
	}
	return total
}

/**
官方解答：方法一：枚举 + 字符串
*/
func findNumbers1(nums []int) int {
	total := 0
	for _, num := range nums {
		// a&1等价于a%2;用于判断数字的奇偶
		if len(strconv.Itoa(num))&1 == 0 {
			total++
		}
	}
	return total
}

/**
官方解答:方法二：枚举 + 数学
	我们也可以使用语言内置的以 10 为底的对数函数 log10() 来得到整数 x 包含的数字个数。
*/
func findNumbers2(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		if int(math.Log10(float64(nums[i]))+1)%2 == 0 {
			total++
		}
	}
	return total
}
