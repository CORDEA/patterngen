package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const Row = 3
const Column = 3

func horizontalChoices(c float64, v float64, withCenter bool) []float64 {
	var result []float64
	if c > 0 {
		result = append(result, v-1)
	}
	if withCenter {
		result = append(result, v)
	}
	if c < Column-1 {
		result = append(result, v+1)
	}
	return result
}

func choices(v float64) []float64 {
	c := math.Mod(v, Column)
	r := math.Floor(v / Column)
	var choices []float64
	if r > 0 {
		choices = append(
			choices,
			horizontalChoices(c, v-Column, true)...,
		)
	}
	choices = append(
		choices,
		horizontalChoices(c, v, false)...,
	)
	if r < Row-1 {
		choices = append(
			choices,
			horizontalChoices(c, v+Column, true)...,
		)
	}
	return choices
}

func dice(choices []float64) float64 {
	return choices[rand.Intn(len(choices))]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	next := float64(rand.Intn(9))
	var result = []float64{next + 1}
	for i := 0; i < 5; i++ {
		c := choices(next)
		next = dice(c)
		result = append(result, next+1)
	}
	fmt.Printf("%v\n", result)
}
