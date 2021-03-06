package leetgo

import "strconv"

/*
给定一组字符，使用原地算法将其压缩。

压缩后的长度必须始终小于或等于原数组长度。

数组的每个元素应该是长度为1 的字符（不是 int 整数类型）。

在完成原地修改输入数组后，返回数组的新长度。



进阶：
你能否仅使用O(1) 空间解决问题？
*/

func compress(chars []byte) int {
	t, i := 0, 0
	for t < len(chars) && i < len(chars) {
		chars[t] = chars[i]
		t++
		temp := i
		for i < len(chars) && chars[i] == chars[t-1] {
			i++
		}
		if i-temp > 1 {
			for _, e := range strconv.Itoa(i - temp) {
				chars[t] = byte(e)
				t++
			}
		}
	}
	return t
}

func compress1(chars []byte) int {
	write := 0
	for slow, fast := 0, 0; slow < len(chars); slow = fast {
		for fast != len(chars) && chars[slow] == chars[fast] {
			fast++
		}
		chars[write] = chars[slow]
		write++
		if fast-slow == 1 {
			continue
		}
		for _, c := range []byte(strconv.Itoa(fast - slow)) {
			chars[write] = c
			write++
		}
	}
	return write
}
