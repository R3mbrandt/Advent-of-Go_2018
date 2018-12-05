package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/R3mbrandt/Advent-of-Go_2018/helper"
)

type shiftState struct {
	date  string
	state string
}

type guardSleeps struct {
	sumMinutes   int
	sleepMinutes map[int]int
}

func convertSortData(input []string) (data []shiftState) {
	re := regexp.MustCompile(`\[(.*)\] (.*)`)

	for _, l := range input {
		m := re.FindStringSubmatch(l)
		d := strings.Trim(m[1], "-")
		data = append(data, shiftState{d, m[2]})
	}
	sort.Slice(data, func(i, j int) bool { return data[i].date < data[j].date })
	return
}

func fillGuards(data []shiftState) map[int]*guardSleeps {
	aktGuard := -1
	guards := make(map[int]*guardSleeps)
	re := regexp.MustCompile(`#(\d+)`)

	//fmt.Println(guards)
	for i := 0; i < len(data); i++ {
		if i >= len(data) {
			break
		}
		m := re.FindStringSubmatch(data[i].state)

		if m == nil {
			break
		}
		aktGuard, _ = strconv.Atoi(m[1])
		//fmt.Println(i,data[i],m,aktGuard)
		if guards[aktGuard] == nil {
			guards[aktGuard] = &guardSleeps{0, make(map[int]int)}
		}
		for j := i + 1; j < len(data)-1; j += 2 {
			//fmt.Println("--",j,data[j],aktGuard)
			if re.FindStringSubmatch(data[j].state) != nil {
				i = j - 1
				break
			}
			d1 := data[j].date
			d2 := data[j+1].date
			asleep, _ := strconv.Atoi(d1[len(d1)-2:])
			wakeup, _ := strconv.Atoi(d2[len(d2)-2:])
			guards[aktGuard].sumMinutes += (wakeup - asleep)
			for m := asleep; m < wakeup; m++ {
				guards[aktGuard].sleepMinutes[m]++
			}
		}
	}
	return guards
}

func main() {
	input := helper.ReadInput("input_day04.txt")

	data := convertSortData(input)
	fmt.Println(len(data))
	fmt.Println("----")
	guards := fillGuards(data)
	mostsleepminuts := 0
	sleepguy := 0
	for k, v := range guards {
		if v.sumMinutes > mostsleepminuts {
			mostsleepminuts = v.sumMinutes
			sleepguy = k
		}
	}
	fmt.Println(sleepguy)

	minute := 0
	times := 0

	for m, v := range guards[sleepguy].sleepMinutes {
		if v > times {
			times = v
			minute = m
		}
	}

	fmt.Println(sleepguy, minute, sleepguy*minute)
	fmt.Println("------")

	for k, v := range guards {
		minute = 0
		times = 0
		for m, t := range v.sleepMinutes {
			if t > times {
				times = t
				minute = m
			}
		}
		fmt.Println(k, minute, times)
	}

}
