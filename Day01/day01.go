package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/R3mbrandt/Advent-of-Go_2018/helper"
)

//processInput converts an array of strings, to an array of ints
func processInput(lines []string) (data []int) {
	for _, l := range lines {
		value, err := strconv.Atoi(l)
		if err != nil {
			log.Fatalln("Can't parse input data", err)
		}
		data = append(data, value)
	}
	return
}

//part1 part one of the puzzle -- just sum up the whole array (maybe there is a higher level function in golang?)
func part1(data []int) (result int) {
	for _, i := range data {
		result += i
	}
	return
}

//part2 part two of the puzzle -- use a "set" (in golang it is a map[int]struct{} - construction), to save all frequencies
//the tricky part is, that you have to loop over the list again and again until you find a frequency (sum), which is already in the set!
func part2(data []int) (result int) {
	var exists = struct{}{}
	set := make(map[int]struct{})
	set[0] = exists
	for {
		for _, i := range data {
			result += i
			if _, ok := set[result]; ok {
				return
			}
			set[result] = exists
		}
	}
}

//main Programm starts here
func main() {
	input := helper.ReadInput("input_day01.txt")
	data := processInput(input)

	sum := part1(data)
	fmt.Println("Part 1:", sum)

	freq := part2(data)
	fmt.Println("Part 2:", freq)
}
