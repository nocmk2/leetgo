package leetgo

import "strconv"

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
