package exercise

/**
给你一个整数数组 nums 。

如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

返回好数对的数目。



示例 1：

输入：nums = [1,2,3,1,1,3]
输出：4
解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
示例 2：

输入：nums = [1,1,1,1]
输出：6
解释：数组中的每组数字都是好数对
示例 3：

输入：nums = [1,2,3]
输出：0


提示：

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/
func NumIdenticalPairs(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				total++
			}
		}
	}
	return total
}

/**
官方解法二：组合计数
思路与算法

用哈希表统计每个数在序列中出现的次数，假设数字 kk 在序列中出现的次数为 vv，那么满足题目中所说的 {\rm nums}[i] = {\rm nums}[j] = k(i < j)nums[i]=nums[j]=k(i<j) 的 (i, j)(i,j) 的数量就是 \frac{v(v - 1)}{2}
2
v(v−1)
​
 ，即 kk 这个数值对答案的贡献是 \frac{v(v - 1)}{2}
2
v(v−1)
​
 。我们只需要把所有数值的贡献相加，即可得到答案。。
--------------------------------------
a1=(v-1);d=-1;an=0;总项数为v的等差数列求和公式：v*(v-1)/2
*/
func NumIdenticalPairs1(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}
