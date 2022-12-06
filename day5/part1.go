package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	stacks := [9][]string{}

	input := "        [C] [B] [H]                \n[W]     [D] [J] [Q] [B]            \n[P] [F] [Z] [F] [B] [L]            \n[G] [Z] [N] [P] [J] [S] [V]        \n[Z] [C] [H] [Z] [G] [T] [Z]     [C]\n[V] [B] [M] [M] [C] [Q] [C] [G] [H]\n[S] [V] [L] [D] [F] [F] [G] [L] [F]\n[B] [J] [V] [L] [V] [G] [L] [N] [J]\n 1   2   3   4   5   6   7   8   9 \n\nmove 5 from 4 to 7\nmove 8 from 5 to 9\nmove 6 from 2 to 8\nmove 7 from 7 to 9\nmove 1 from 7 to 4\nmove 2 from 7 to 4\nmove 9 from 8 to 4\nmove 16 from 9 to 7\nmove 1 from 3 to 8\nmove 15 from 4 to 5\nmove 3 from 9 to 5\nmove 2 from 3 to 5\nmove 1 from 8 to 7\nmove 3 from 1 to 7\nmove 5 from 3 to 5\nmove 13 from 7 to 2\nmove 5 from 7 to 1\nmove 7 from 2 to 6\nmove 2 from 7 to 8\nmove 3 from 6 to 5\nmove 2 from 8 to 2\nmove 2 from 6 to 1\nmove 11 from 1 to 7\nmove 2 from 2 to 9\nmove 8 from 6 to 5\nmove 2 from 9 to 6\nmove 3 from 6 to 4\nmove 1 from 4 to 7\nmove 22 from 5 to 6\nmove 13 from 6 to 9\nmove 5 from 2 to 7\nmove 6 from 5 to 8\nmove 13 from 7 to 2\nmove 2 from 4 to 6\nmove 5 from 6 to 3\nmove 2 from 7 to 5\nmove 3 from 3 to 6\nmove 2 from 6 to 2\nmove 8 from 2 to 4\nmove 2 from 4 to 7\nmove 2 from 2 to 9\nmove 5 from 4 to 5\nmove 2 from 3 to 2\nmove 1 from 5 to 4\nmove 6 from 5 to 9\nmove 1 from 7 to 3\nmove 1 from 5 to 9\nmove 5 from 5 to 1\nmove 1 from 6 to 8\nmove 1 from 5 to 8\nmove 4 from 6 to 9\nmove 8 from 8 to 9\nmove 1 from 3 to 6\nmove 4 from 1 to 7\nmove 3 from 6 to 4\nmove 7 from 2 to 6\nmove 27 from 9 to 8\nmove 3 from 4 to 7\nmove 6 from 8 to 1\nmove 1 from 4 to 6\nmove 1 from 2 to 7\nmove 7 from 6 to 3\nmove 1 from 4 to 3\nmove 4 from 1 to 6\nmove 1 from 9 to 2\nmove 1 from 2 to 4\nmove 1 from 4 to 5\nmove 3 from 9 to 4\nmove 5 from 7 to 8\nmove 2 from 5 to 6\nmove 4 from 6 to 9\nmove 10 from 8 to 3\nmove 2 from 4 to 7\nmove 3 from 1 to 7\nmove 2 from 9 to 6\nmove 6 from 3 to 1\nmove 7 from 3 to 4\nmove 2 from 1 to 9\nmove 4 from 1 to 9\nmove 1 from 3 to 6\nmove 1 from 3 to 8\nmove 2 from 9 to 5\nmove 2 from 5 to 3\nmove 3 from 3 to 1\nmove 1 from 4 to 6\nmove 5 from 7 to 6\nmove 2 from 3 to 4\nmove 2 from 8 to 1\nmove 9 from 4 to 7\nmove 4 from 9 to 3\nmove 2 from 8 to 3\nmove 1 from 1 to 4\nmove 1 from 6 to 2\nmove 1 from 2 to 9\nmove 6 from 3 to 5\nmove 2 from 1 to 3\nmove 1 from 3 to 2\nmove 1 from 2 to 9\nmove 8 from 6 to 8\nmove 2 from 6 to 3\nmove 1 from 1 to 2\nmove 7 from 7 to 9\nmove 13 from 8 to 6\nmove 1 from 2 to 8\nmove 6 from 9 to 3\nmove 1 from 1 to 6\nmove 2 from 8 to 5\nmove 5 from 3 to 4\nmove 2 from 8 to 1\nmove 8 from 5 to 2\nmove 4 from 3 to 2\nmove 5 from 8 to 4\nmove 2 from 9 to 4\nmove 4 from 4 to 7\nmove 10 from 2 to 6\nmove 1 from 2 to 9\nmove 24 from 6 to 1\nmove 17 from 1 to 8\nmove 1 from 9 to 2\nmove 2 from 4 to 9\nmove 10 from 7 to 4\nmove 1 from 2 to 5\nmove 5 from 9 to 1\nmove 1 from 7 to 6\nmove 12 from 8 to 6\nmove 1 from 7 to 5\nmove 2 from 5 to 6\nmove 16 from 6 to 8\nmove 12 from 1 to 6\nmove 2 from 1 to 7\nmove 9 from 6 to 2\nmove 2 from 4 to 1\nmove 1 from 1 to 5\nmove 7 from 4 to 6\nmove 13 from 8 to 2\nmove 5 from 8 to 2\nmove 2 from 7 to 3\nmove 2 from 4 to 9\nmove 1 from 5 to 4\nmove 3 from 9 to 8\nmove 2 from 4 to 2\nmove 2 from 3 to 8\nmove 1 from 1 to 5\nmove 1 from 4 to 8\nmove 6 from 2 to 7\nmove 1 from 5 to 8\nmove 1 from 6 to 2\nmove 7 from 6 to 8\nmove 1 from 6 to 2\nmove 24 from 2 to 1\nmove 10 from 8 to 3\nmove 4 from 8 to 2\nmove 4 from 7 to 1\nmove 5 from 2 to 9\nmove 1 from 6 to 2\nmove 10 from 3 to 1\nmove 2 from 7 to 3\nmove 2 from 3 to 7\nmove 2 from 7 to 9\nmove 35 from 1 to 5\nmove 28 from 5 to 6\nmove 2 from 2 to 7\nmove 19 from 6 to 4\nmove 3 from 1 to 2\nmove 3 from 2 to 5\nmove 23 from 4 to 7\nmove 2 from 6 to 8\nmove 4 from 7 to 6\nmove 3 from 5 to 6\nmove 13 from 7 to 4\nmove 2 from 5 to 6\nmove 2 from 9 to 4\nmove 5 from 6 to 3\nmove 6 from 4 to 5\nmove 1 from 4 to 8\nmove 4 from 4 to 6\nmove 5 from 9 to 7\nmove 2 from 8 to 7\nmove 5 from 3 to 2\nmove 4 from 5 to 2\nmove 5 from 2 to 9\nmove 4 from 8 to 4\nmove 1 from 9 to 8\nmove 2 from 2 to 6\nmove 4 from 4 to 2\nmove 3 from 2 to 3\nmove 3 from 5 to 1\nmove 2 from 3 to 2\nmove 3 from 1 to 4\nmove 1 from 9 to 4\nmove 5 from 4 to 9\nmove 2 from 4 to 3\nmove 5 from 6 to 8\nmove 1 from 9 to 7\nmove 2 from 6 to 3\nmove 1 from 4 to 5\nmove 1 from 9 to 4\nmove 6 from 8 to 6\nmove 2 from 3 to 6\nmove 2 from 9 to 4\nmove 2 from 3 to 9\nmove 1 from 3 to 1\nmove 17 from 6 to 4\nmove 1 from 1 to 8\nmove 1 from 6 to 5\nmove 1 from 9 to 2\nmove 11 from 4 to 6\nmove 9 from 4 to 5\nmove 7 from 9 to 4\nmove 2 from 5 to 2\nmove 1 from 4 to 9\nmove 5 from 2 to 1\nmove 1 from 2 to 9\nmove 4 from 4 to 9\nmove 4 from 1 to 5\nmove 1 from 1 to 7\nmove 1 from 8 to 9\nmove 8 from 7 to 8\nmove 4 from 7 to 4\nmove 9 from 5 to 2\nmove 2 from 4 to 1\nmove 11 from 6 to 8\nmove 2 from 4 to 3\nmove 2 from 4 to 8\nmove 1 from 1 to 4\nmove 3 from 2 to 8\nmove 1 from 1 to 3\nmove 3 from 3 to 9\nmove 8 from 9 to 6\nmove 1 from 4 to 8\nmove 2 from 9 to 3\nmove 5 from 6 to 9\nmove 7 from 5 to 6\nmove 2 from 3 to 4\nmove 5 from 7 to 9\nmove 2 from 4 to 5\nmove 2 from 2 to 3\nmove 10 from 9 to 5\nmove 2 from 6 to 3\nmove 6 from 2 to 7\nmove 10 from 5 to 3\nmove 6 from 7 to 1\nmove 2 from 1 to 7\nmove 4 from 3 to 9\nmove 3 from 8 to 2\nmove 2 from 7 to 5\nmove 19 from 8 to 7\nmove 4 from 5 to 9\nmove 4 from 9 to 8\nmove 1 from 2 to 5\nmove 3 from 6 to 8\nmove 1 from 5 to 9\nmove 5 from 9 to 7\nmove 6 from 3 to 8\nmove 1 from 3 to 8\nmove 2 from 3 to 2\nmove 23 from 7 to 6\nmove 10 from 8 to 4\nmove 4 from 4 to 9\nmove 4 from 2 to 6\nmove 1 from 3 to 8\nmove 4 from 8 to 4\nmove 31 from 6 to 4\nmove 9 from 4 to 5\nmove 8 from 5 to 3\nmove 1 from 6 to 7\nmove 2 from 5 to 7\nmove 4 from 9 to 2\nmove 21 from 4 to 8\nmove 4 from 2 to 9\nmove 3 from 3 to 9\nmove 2 from 7 to 9\nmove 11 from 4 to 9\nmove 1 from 8 to 5\nmove 1 from 5 to 9\nmove 9 from 9 to 3\nmove 3 from 1 to 5\nmove 2 from 5 to 8\nmove 11 from 3 to 6\nmove 4 from 6 to 3\nmove 2 from 8 to 3\nmove 10 from 9 to 6\nmove 22 from 8 to 9\nmove 1 from 1 to 8\nmove 4 from 6 to 3\nmove 2 from 7 to 6\nmove 3 from 8 to 3\nmove 14 from 3 to 2\nmove 1 from 3 to 4\nmove 1 from 2 to 4\nmove 2 from 9 to 1\nmove 1 from 5 to 7\nmove 1 from 3 to 2\nmove 14 from 6 to 5\nmove 13 from 5 to 2\nmove 1 from 5 to 6\nmove 1 from 7 to 9\nmove 8 from 9 to 4\nmove 2 from 6 to 7\nmove 23 from 2 to 4\nmove 2 from 1 to 4\nmove 2 from 2 to 5\nmove 1 from 5 to 1\nmove 1 from 7 to 2\nmove 1 from 5 to 9\nmove 16 from 9 to 5\nmove 1 from 2 to 4\nmove 13 from 5 to 3\nmove 1 from 1 to 4\nmove 1 from 7 to 1\nmove 1 from 5 to 3\nmove 2 from 5 to 7\nmove 2 from 7 to 1\nmove 9 from 3 to 2\nmove 2 from 1 to 7\nmove 1 from 1 to 9\nmove 19 from 4 to 2\nmove 1 from 9 to 7\nmove 1 from 7 to 8\nmove 23 from 2 to 8\nmove 2 from 7 to 2\nmove 12 from 4 to 5\nmove 12 from 5 to 1\nmove 5 from 2 to 9\nmove 2 from 2 to 7\nmove 5 from 8 to 1\nmove 3 from 9 to 4\nmove 1 from 2 to 8\nmove 1 from 2 to 4\nmove 4 from 8 to 1\nmove 2 from 3 to 1\nmove 2 from 7 to 5\nmove 1 from 4 to 9\nmove 8 from 4 to 7\nmove 13 from 8 to 6\nmove 1 from 3 to 1\nmove 13 from 6 to 7\nmove 13 from 7 to 6\nmove 7 from 1 to 4\nmove 5 from 7 to 3\nmove 3 from 4 to 3\nmove 13 from 6 to 1\nmove 3 from 8 to 6\nmove 8 from 3 to 8\nmove 12 from 1 to 8\nmove 1 from 3 to 5\nmove 6 from 1 to 7\nmove 3 from 6 to 8\nmove 1 from 3 to 8\nmove 1 from 9 to 2\nmove 3 from 5 to 6\nmove 1 from 7 to 3\nmove 8 from 7 to 1\nmove 2 from 6 to 2\nmove 3 from 4 to 3\nmove 2 from 9 to 2\nmove 6 from 8 to 9\nmove 5 from 2 to 5\nmove 2 from 3 to 4\nmove 5 from 5 to 4\nmove 1 from 3 to 9\nmove 8 from 4 to 5\nmove 1 from 6 to 8\nmove 2 from 1 to 4\nmove 1 from 1 to 4\nmove 3 from 1 to 5\nmove 3 from 1 to 6\nmove 7 from 1 to 9\nmove 2 from 6 to 9\nmove 1 from 3 to 5\nmove 17 from 8 to 7\nmove 17 from 7 to 6\nmove 5 from 5 to 2\nmove 5 from 2 to 1\nmove 13 from 6 to 2\nmove 1 from 1 to 4\nmove 5 from 5 to 1\nmove 1 from 1 to 5\nmove 10 from 9 to 1\nmove 13 from 1 to 8\nmove 13 from 8 to 4\nmove 5 from 6 to 7\nmove 8 from 1 to 7\nmove 1 from 1 to 3\nmove 12 from 2 to 6\nmove 1 from 3 to 8\nmove 6 from 6 to 2\nmove 2 from 5 to 1\nmove 5 from 2 to 5\nmove 2 from 5 to 9\nmove 12 from 4 to 2\nmove 1 from 6 to 2\nmove 15 from 2 to 1\nmove 1 from 8 to 6\nmove 2 from 7 to 3\nmove 2 from 4 to 2\nmove 1 from 2 to 9\nmove 1 from 2 to 6\nmove 7 from 7 to 3\nmove 1 from 4 to 1\nmove 17 from 1 to 2\nmove 3 from 6 to 4\nmove 1 from 3 to 8\nmove 3 from 9 to 6\nmove 4 from 6 to 3\nmove 13 from 2 to 9\nmove 3 from 2 to 8\nmove 2 from 5 to 1\nmove 6 from 8 to 2\nmove 1 from 6 to 2\nmove 3 from 2 to 7\nmove 3 from 1 to 6\nmove 2 from 9 to 8\nmove 6 from 9 to 8\nmove 8 from 9 to 3\nmove 7 from 7 to 4\nmove 20 from 3 to 7\nmove 4 from 6 to 8\nmove 1 from 8 to 6\nmove 2 from 6 to 4\nmove 3 from 2 to 1\nmove 2 from 9 to 6\nmove 9 from 8 to 6\nmove 3 from 1 to 9\nmove 9 from 4 to 8\nmove 1 from 5 to 6\nmove 3 from 4 to 2\nmove 1 from 5 to 3\nmove 8 from 6 to 4\nmove 4 from 9 to 3\nmove 10 from 8 to 6\nmove 5 from 2 to 3\nmove 3 from 6 to 4\nmove 10 from 3 to 1\nmove 11 from 4 to 1\nmove 1 from 8 to 2\nmove 2 from 4 to 2\nmove 1 from 4 to 9\nmove 10 from 6 to 3\nmove 21 from 1 to 5\nmove 2 from 2 to 7\nmove 1 from 9 to 6\nmove 1 from 6 to 3\nmove 1 from 6 to 7\nmove 11 from 5 to 6\nmove 1 from 2 to 8\nmove 1 from 5 to 9\nmove 11 from 6 to 3\nmove 1 from 8 to 4\nmove 1 from 4 to 1\nmove 3 from 5 to 7\nmove 1 from 1 to 5\nmove 5 from 5 to 8\nmove 23 from 7 to 9\nmove 5 from 8 to 4\nmove 1 from 5 to 2\nmove 12 from 3 to 4\nmove 6 from 3 to 6\nmove 1 from 5 to 2\nmove 8 from 9 to 2\nmove 1 from 7 to 8\nmove 2 from 7 to 9\nmove 4 from 3 to 5\nmove 1 from 5 to 9\nmove 1 from 6 to 5\nmove 4 from 6 to 5\nmove 3 from 2 to 1\nmove 3 from 1 to 3\nmove 8 from 9 to 1\nmove 4 from 2 to 9\nmove 1 from 9 to 7\nmove 14 from 4 to 8\nmove 3 from 3 to 4\nmove 1 from 5 to 8\nmove 2 from 8 to 6\nmove 2 from 6 to 7\nmove 4 from 4 to 3\nmove 12 from 9 to 1\nmove 1 from 3 to 2\nmove 6 from 8 to 2\nmove 1 from 7 to 1\nmove 5 from 2 to 3\nmove 21 from 1 to 3\nmove 5 from 5 to 4\nmove 1 from 8 to 5\nmove 2 from 2 to 7\nmove 1 from 6 to 1\nmove 2 from 9 to 2\nmove 1 from 2 to 9\nmove 1 from 1 to 5\nmove 4 from 3 to 5\nmove 7 from 8 to 1\nmove 6 from 1 to 9\nmove 1 from 2 to 5\nmove 6 from 9 to 7\nmove 8 from 3 to 4\nmove 2 from 4 to 8\nmove 1 from 1 to 6\nmove 10 from 3 to 9\nmove 12 from 4 to 2\nmove 1 from 8 to 1"
	lines := strings.Split(input, "\n")
	i := 0

	for i = 0; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
		for j := range lines[i] {
			if unicode.IsLetter(rune(lines[i][j])) {
				stacks[(j-1)/4] = append(stacks[(j-1)/4], string(lines[i][j]))
			}
		}

	}
	i++
	for ; i < len(lines); i++ {
		a := strings.Split(lines[i], " ")
		dst, _ := strconv.Atoi(a[5])
		src, _ := strconv.Atoi(a[3])
		dst--
		src--
		num, _ := strconv.Atoi(a[1])
		fmt.Println(src, dst, lines[i], i)
		for j := 0; j < num; j++ {
			stacks[dst] = append([]string{stacks[src][0]}, stacks[dst]...)
			stacks[src] = stacks[src][1:]
		}
	}
	ans := ""
	for k := 0; k < 9; k++ {
		if len(stacks[k]) != 0 {
			ans += stacks[k][0]
		}
	}
	fmt.Println(ans)
}
func print(stacks [9][]string, lable string) {
	//print the stacks vertically next to each other
	fmt.Println(lable, "----------------")

	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			if len(stacks[j]) > i {
				fmt.Print(stacks[j][i])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
