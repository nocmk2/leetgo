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

//func compress(chars []byte) int {
//	k := 0
//	for i, j := 0, 0; i < len(chars); i = j {
//		for j < len(chars) && chars[i] == chars[j] {
//			j++
//		}
//		chars[k] = chars[i]
//		k++
//		if j-i == 1 {
//			continue
//		}
//		for _, c := range []byte(strconv.Itoa(j - i)) {
//			chars[k] = c
//			k++
//		}
//	}
//	return k
//}
