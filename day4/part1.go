package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	sum := 0
	for _, i := range strings.Split(input, "\n") {
		firstRange, secondRange := strings.Split(i, ",")[0], strings.Split(i, ",")[1]
		firstLeft, _ := strconv.Atoi(strings.Split(firstRange, "-")[0])
		firstRight, _ := strconv.Atoi(strings.Split(firstRange, "-")[1])
		secondLeft, _ := strconv.Atoi(strings.Split(secondRange, "-")[0])
		secondRight, _ := strconv.Atoi(strings.Split(secondRange, "-")[1])
		if (secondRight >= firstRight && secondLeft <= firstLeft) || (firstRight >= secondRight && firstLeft <= secondLeft) {
			fmt.Println(i)
			sum++
		}
	}
	fmt.Println(sum)
}
