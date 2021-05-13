package exercise

import "math"

/**
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。


示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0


提示：

-231 <= x <= 231 - 1
*/
func Reverse(x int) int {
	var nums []int
	symbol := 1
	if x < 0 {
		symbol = -1
		x *= -1
	}
	for x != 0 {
		nums = append(nums, x%10)
		x /= 10
	}
	result := 0
	for i := len(nums) - 1; i >= 0; i-- {
		temp := 1
		for j := 0; j < i; j++ {
			temp *= 10
		}
		if math.MaxInt32-(nums[len(nums)-1-i]*temp) < result || math.MinInt32+(nums[len(nums)-1-i]*temp) > result*-1 {
			return 0
		}
		result += nums[len(nums)-1-i] * temp
	}
	return result * symbol
}

/**
官方更优代码
*/
func reverse(x int) int {
	result := 0
	for x != 0 {
		if result > math.MaxInt32/10 || result < math.MinInt32/10 {
			return 0
		}
		result = result*10 + (x % 10)
		x /= 10
	}
	return result
}
