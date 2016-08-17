package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var points []Point

	// f, err := os.Open("/Users/peter/Projects/GO/src/github.com/peterFran/hackerrank/test.txt")

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
	fmt.Println(points)
}
