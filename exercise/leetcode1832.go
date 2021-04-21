package exercise

import "strings"

/**

全字母句 指包含英语字母表中每个字母至少一次的句子。

给你一个仅由小写英文字母组成的字符串 sentence ，请你判断 sentence 是否为 全字母句 。

如果是，返回 true ；否则，返回 false 。



示例 1：

输入：sentence = "thequickbrownfoxjumpsoverthelazydog"
输出：true
解释：sentence 包含英语字母表中每个字母至少一次。
示例 2：

输入：sentence = "leetcode"
输出：false


提示：

1 <= sentence.length <= 1000
sentence 由小写英语字母组成
*/
func CheckIfPangram(sentence string) bool {
	letterMap := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0,
		"r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	for _, letter := range strings.Split(sentence, "") {
		letterMap[letter]++
	}
	for letter := range letterMap {
		if letterMap[letter] == 0 {
			return false
		}
	}
	return true
}

/**
位运算解法: int型有32位，使用了其中的26位，每一位对应一个字母，然后使用了或运算，比如出现了一个字母a，那么(letter - 'a') 的ascII值为0，此时1向左移动零位，
即第一位为1，以此类推 —> 第五行代码的逻辑。 而0x3ffffff的意思是，一个f在16进制中表示15，15转换为2进制为1111，即4位，所以6个f一共是24位；
3转换为2进制为11，即2位，所以24+2=26位，也就代表了26个字母； ^表示异或运算，1 ^ 1 = 0，也就是说如果res前26位都为1的话，res ^ 0x3ffffff 的值将为 0；
故而返回true，否则返回false。
*/
func checkIfPangram1(sentence string) bool {
	var res int32 = 0
	for _, letter := range []rune(sentence) {
		res |= 1 << (letter - 'a')
		if (res ^ 0x3ffffff) == 0 {
			return true
		}
	}
	return false
}
