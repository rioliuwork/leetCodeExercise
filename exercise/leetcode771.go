package exercise

/**
 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

示例 1:

输入: J = "aA", S = "aAAbbbb"
输出: 3
示例 2:

输入: J = "z", S = "ZZ"
输出: 0
注意:

S 和 J 最多含有50个字母。
 J 中的字符不重复。
*/
func NumJewelsInStones(jewels string, stones string) int {
	jMap := make(map[rune]int)
	jRunes := []rune(jewels)
	sRunes := []rune(stones)
	for _, r := range jRunes {
		jMap[r] = 0
		for _, s := range sRunes {
			if s == r {
				jMap[r]++
			}
		}
	}
	sum := 0
	for r, _ := range jMap {
		sum += jMap[r]
	}
	return sum
}

/**
官方解答 方法一:暴力
*/
func numJewelsInStones1(jewels string, stones string) int {
	jewelsCount := 0
	for _, s := range stones {
		for _, j := range jewels {
			if s == j {
				jewelsCount++
				break
			}
		}
	}
	return jewelsCount
}

/**
官方解答 方法二:哈希集合
*/
func numJewelsInStones2(jewels string, stones string) int {
	jewelsCount := 0
	jewelsSet := map[byte]bool{}
	for i := range jewels {
		jewelsSet[jewels[i]] = true
	}
	for i := range stones {
		if jewelsSet[stones[i]] {
			jewelsCount++
		}
	}
	return jewelsCount
}
