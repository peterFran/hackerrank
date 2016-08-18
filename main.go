package main

import (
	"bufio"
	"fmt"
	"github.com/fighterlyt/permutation"
	"math"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

type Route struct {
	points []Point
	score  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var points []Point

	// f, err := os.Open("/Users/peter/Projects/GO/src/github.com/peterFran/hacker/test.txt")

	f, err := os.Open("/Users/petermeckiffe/Projects/go/src/github.com/peterfran/hackerrank/test.txt")
	check(err)
	reader := bufio.NewReader(f)
	meta, _, _ := reader.ReadLine()
	x, _ := strconv.Atoi(string(meta[0]))
	y, _ := strconv.Atoi(string(meta[2]))
	start := Point{x, y}
	points = append(points, start)

	i := 0
	for line, _, err := reader.ReadLine(); ; line, _, err = reader.ReadLine() {
		if err != nil {
			break
		}
		fmt.Println(string(line))
		for j := 0; j < len(line); j++ {

			if string(line[j]) == "d" {
				dp := Point{j, i}
				points = append(points, dp)
			}

		}
		i++
	}
	perms := range int[]{len(points)-1}
	fmt.Println(points)
	route := Route{points, 0}
	calculateScore(&route)
	fmt.Println(route.score)
}

func calculateScore(route *Route) {
	for i := 0; i < len(route.points)-1; i++ {
		diffx := route.points[i].X - route.points[i+1].X
		diffy := route.points[i].Y - route.points[i+1].Y

		diffx = int(math.Abs(float64(diffx)))
		diffy = int(math.Abs(float64(diffy)))
		fmt.Println(diffx, diffy)
		route.score += diffx
		route.score += diffy
	}
}
