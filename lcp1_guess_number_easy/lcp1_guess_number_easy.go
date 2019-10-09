package main

import "fmt"

func game(guess []int, answer []int) int {
	cnt := 0
	if guess[0] == answer[0] {
		cnt++
	}
	if guess[1] == answer[0] {
		cnt++
	}
	if guess[1] == answer[1] {
		cnt++
	}

	return cnt

}

func main() {
	fmt.Println(game([]int{1, 2, 3}, []int{1, 1, 1}))
}
