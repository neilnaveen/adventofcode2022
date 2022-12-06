package main

import (
	"fmt"
	"strings"
)

type point struct {
	a, b int
}

func main() {
	input := "A Y\nB X\nC Z"
	enemy := "ABC"
	me := "XYZ"
	m := map[point]int{point{1, 0}: 6, point{2, 1}: 6, point{0, 2}: 6, point{0, 0}: 3, point{1, 1}: 3, point{2, 2}: 3}
	score := 0
	for _, i := range strings.Split(input, "\n") {
		mval, eval := 0, 0
		for j := range enemy {
			if enemy[j] == i[0] {
				eval = j
			}
		}
		for j := range me {
			if me[j] == i[2] {
				mval = j
			}
		}
		score += m[point{mval, eval}] + mval + 1
	}
	fmt.Println(score)

}
