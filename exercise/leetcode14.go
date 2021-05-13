package exercise

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
*/
func longestCommonPrefix(strs []string) string {
	com := strs[0]
	for _, str := range strs {
		if len(str) < len(com) {
			com, str = str, com
		}
		for i := 0; i < len(com); i++ {
			if str[i] != com[i] {
				com = com[:i]
			}
		}
	}
	return com
}
