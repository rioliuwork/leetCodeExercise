package exercise

import "math"

/**
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/
func Divide(dividend int, divisor int) int {
	var result int
	var flag, flag1, flag2 int
	if dividend < 0 {
		flag1 = 1
	}
	if divisor < 0 {
		flag2 = 1
	}
	flag = flag1 ^ flag2
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return dividend * -1
		}
		return math.MaxInt32
	}
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	//for dividend >= divisor {
	//	dividend -= divisor
	//	result++
	//}

	result = digui(result, dividend, divisor, divisor)

	if flag == 1 {
		result *= -1
	}
	return result
}

func digui(result int, dividend int, divisor int, base int) int {
	for {
		if dividend < base {
			return 0
		}
		temp := divisor
		divisor += divisor
		if result != 0 {
			result = result + result
		} else {
			result++
		}
		if dividend < divisor {
			if base > (dividend - temp) {
				return result
			}
			return result + digui(0, dividend-temp, base, base)
		}
	}
}

func add(a int, b int) int {
	if b == 0 {
		return a
	}

	sum := a ^ b
	carry := (a & b) << 1
	return add(sum, carry)
}

func sub(a int, b int) int {
	return add(a, add(^b, 1))
}

func mul(a int, b int) int {
	res := 0
	for i := 1; i != 0; {
		if b&i != 0 {
			res = add(res, a)
		}
		i <<= 1
		a <<= 1
	}
	return res
}

//有问题,计算有错
func Div(a int, b int) int {
	sign := 1
	if a&(1<<31) != 0 {
		a = ^sub(a, 1)
		sign ^= 1
	}
	if b&(1<<31) != 0 {
		b = ^sub(b, 1)
		sign ^= 1
	}

	res := 0
	for i := 0x8000; i != 0; i >>= 1 {
		if ((a >> i) & 0xFFFF) >= b {
			res = add(res, (1<<i)&0xFFFF)
			a = sub(a, (b<<i)&0xFFFF)
		}
	}

	if sign == 0 {
		res = ^sub(res, 1)
	}
	return res
}
