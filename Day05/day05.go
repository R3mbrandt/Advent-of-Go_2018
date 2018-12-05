package main

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/R3mbrandt/Advent-of-Go_2018/helper"
)

func chkDoubles(inp string) string {
	for i := 0; i < len(inp)-1; i++ {
		if (unicode.IsUpper(rune(inp[i])) && (unicode.ToLower(rune(inp[i])) == rune(inp[i+1]))) || (unicode.IsLower(rune(inp[i])) && (unicode.ToUpper(rune(inp[i])) == rune(inp[i+1]))) {
			inp = chkDoubles(inp[:i] + inp[i+2:])
			i--
		}
	}
	return inp
}

func rmvUnit(input, unit string) int {
	re := regexp.MustCompile("[" + unit + "]")
	t := re.ReplaceAllString(input, "")
	return len(chkDoubles(t))
}

func main() {
	//inp := "dabAcCaCBAcCcaDA"
	inp := helper.ReadInput("input_day05.txt")[0]
	out := chkDoubles(inp)
	fmt.Println("Part 1:", len(out))
	min := len(inp)
	for r := 'a'; r <= 'z'; r++ {
		cutset := string(r) + string(unicode.ToUpper(r))
		newlen := rmvUnit(inp, cutset)
		if newlen < min {
			min = newlen
		}
	}
	fmt.Println("Part 2:", min)
}
