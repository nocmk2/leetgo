package leetgo

/*
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

示例 1：

输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
示例 2：

输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
示例 3：

输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

*/
func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		// 和1比较 如果最右边第一位是1的话才会返回1，0的话返回0,这样配合右移就可以数1的个数
		if num&1 == 1 {
			count++
		}
		num >>= 1 // 右移  直到所有的1都移出去后 那么num==0跳出循环
	}
	return count

}
