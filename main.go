package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const Row = 3
const Column = 3
const Length = Row * Column

func horizontalChoices(c float64, v int, withCenter bool, r *[Length]int) {
	if c > 0 {
		(*r)[v-1] = 1
	}
	if withCenter {
		(*r)[v] = 1
	}
	if c < Column-1 {
		(*r)[v+1] = 1
	}
}

func choices(v int) [Length]int {
	c := math.Mod(float64(v), Column)
	r := math.Floor(float64(v) / Column)
	var choices [Length]int
	if r > 0 {
		horizontalChoices(c, v-Column, true, &choices)
	}
	horizontalChoices(c, v, false, &choices)
	if r < Row-1 {
		horizontalChoices(c, v+Column, true, &choices)
	}
	return choices
}

func dice(choices [Length]int) int {
	var r []int
	for i, c := range choices {
		if c < 1 {
			continue
		}
		r = append(r, i)
	}
	return r[rand.Intn(len(r))]
}

func format(choices []string) string {
	return strings.Join(choices, " → ")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	next := rand.Intn(Length)
	var result = []string{strconv.Itoa(next + 1)}
	for i := 0; i < 5; i++ {
		c := choices(next)
		next = dice(c)
		result = append(result, strconv.Itoa(next+1))
	}
	fmt.Println(format(result))
}
