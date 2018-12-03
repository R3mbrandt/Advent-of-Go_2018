package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/R3mbrandt/Advent-of-Go_2018/helper"
)

type patch struct {
	id     int
	startx int
	starty int
	width  int
	length int
}

type inch struct {
	x int
	y int
}

func ProcessData(input []string) []patch {
	patches := make([]patch, 0, len(input))
	r := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
	for _, l := range input {
		m := r.FindStringSubmatch(l)
		id, _ := strconv.Atoi(m[1])
		startx, _ := strconv.Atoi(m[2])
		starty, _ := strconv.Atoi(m[3])
		width, _ := strconv.Atoi(m[4])
		length, _ := strconv.Atoi(m[5])
		patches = append(patches, patch{id, startx, starty, width, length})
	}
	return patches
}

func main() {
	var part1 int
	var exists = struct{}{}
	fabric := make(map[inch][]int)
	claims := make(map[int]struct{})

	input := helper.ReadInput("input_day03.txt")
	data := ProcessData(input)

	for _, d := range data {
		for x := 0; x < d.width; x++ {
			for y := 0; y < d.length; y++ {
				key := inch{d.startx + x, d.starty + y}
				fabric[key] = append(fabric[key], d.id)
			}
		}
		claims[d.id] = exists
	}
	for _, v := range fabric {
		if len(v) > 1 {
			part1++
			for _, c := range v {
				delete(claims, c)
			}
		}
	}
	fmt.Println("Part 1:", part1)
	for key := range claims {
		fmt.Println("Part 2:", key) //should only be one!
	}
}
