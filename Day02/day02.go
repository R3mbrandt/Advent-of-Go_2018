package main

import (
	"fmt"

	"github.com/R3mbrandt/Advent-of-Go_2018/helper"
)

func mapContains(m map[rune]int, x int) bool {
	for _, v := range m {
		if x == v {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	var threes, twos int

	for _, l := range input {
		wordMap := make(map[rune]int)
		for _, r := range l {
			wordMap[r]++
		}
		if mapContains(wordMap, 2) {
			twos++
		}
		if mapContains(wordMap, 3) {
			threes++
		}
	}
	return threes * twos
}

func part2(input []string) string {
	var commonLetters []rune

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			word1 := input[i]
			word2 := input[j]
			commonLetters = nil
			diffCount := 0
			for p, r := range word1 {
				if r == rune(word2[p]) {
					commonLetters = append(commonLetters, r)
					continue
				} else {
					diffCount++
				}
			}
			if diffCount == 1 {
				return string(commonLetters)
			}
		}
	}
	return ""
}

func main() {
	input := helper.ReadInput("input_day02.txt")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
