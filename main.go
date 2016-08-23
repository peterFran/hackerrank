package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Point struct {
	X int
	Y int
}

type Route struct {
	bot      Point
	points   []Point
	order    []int
	score    int
	commands []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// f, err := os.Open("/Users/peter/Projects/GO/src/github.com/peterFran/hacker/test3.txt")

	// f, err := os.Open("/Users/petermeckiffe/Projects/go/src/github.com/peterfran/hackerrank/test.txt")
	// check(err)
	// reader := bufio.NewReader(f)
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	board := []string{}
	// for finished := false; finished == false; {
	meta, _, _ := reader.ReadLine()
	x, _ := strconv.Atoi(string(meta[0]))
	y, _ := strconv.Atoi(string(meta[2]))

	board = append(board, string(line))
	for i := 1; i < 5; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		board = append(board, string(line))
	}
	winningRoute := calculateMoves(x, y, board)

	fmt.Println(winningRoute.commands[0])
	meta, _, err := reader.ReadLine()
	if err != nil {
		time.Sleep(1000)
	}
	meta, _, _ = reader.ReadLine()
	x, _ = strconv.Atoi(string(meta[0]))
	y, _ = strconv.Atoi(string(meta[2]))
	fmt.Println(winningRoute.commands[1])

	// }
}

func calculateMoves(posc, posr int, board []string) Route {
	start := Point{posc, posr}
	points := []Point{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if string(board[i][j]) == "d" {
				dp := Point{j, i}
				points = append(points, dp)
			}
		}
		i++
	}
	var order []int
	for k := 0; k < len(points); k++ {
		order = append(order, k)
	}
	// fmt.Println(order)
	winningRoute := Route{start, points, order, 0, []string{}}
	calculateScore(&winningRoute)

	var col [][]int
	permute(order, 0, len(order)-1, &col)
	// fmt.Println(col)
	for i := 0; i < len(col); i++ {
		route := Route{start, points, col[i], 0, []string{}}
		calculateScore(&route)
		// fmt.Println(route)
		if route.score < winningRoute.score {
			winningRoute = route
			// fmt.Println(winningRoute.score)
		}
	}
	// fmt.Println(winningRoute)
	generateRoute(&winningRoute)
	return winningRoute
}

func calculateScore(route *Route) {

	for i := -1; i < len(route.points)-1; i++ {
		if i == -1 {
			route.score += calculateDiff(route.bot, route.points[route.order[i+1]])
		} else {
			route.score += calculateDiff(route.points[route.order[i]], route.points[route.order[i+1]])
		}
	}
}

func generateRoute(route *Route) {
	for i := -1; i < len(route.order)-1; i++ {
		pointB := route.points[route.order[i+1]]
		// fmt.Println(pointB)
		for diff := calculateDiff(route.bot, pointB); diff > 0; diff = calculateDiff(route.bot, pointB) {
			move := nextMove(route.bot, pointB)
			route.commands = append(route.commands, move)
			movePoint(&route.bot, move)
		}
		route.commands = append(route.commands, "CLEAN")
	}
}

func calculateDiff(pointA Point, pointB Point) int {
	diffx := pointA.X - pointB.X
	diffy := pointA.Y - pointB.Y

	diffx = int(math.Abs(float64(diffx)))
	diffy = int(math.Abs(float64(diffy)))
	return diffx + diffy
}

func nextMove(pointA Point, pointB Point) string {
	diffx := pointA.X - pointB.X
	diffy := pointA.Y - pointB.Y
	if diffx > 0 {
		return "LEFT"
	} else if diffx < 0 {
		return "RIGHT"
	} else if diffy > 0 {
		return "UP"
	} else if diffy < 0 {
		return "DOWN"
	} else {
		return "CLEAN"
	}
}

func movePoint(point *Point, move string) {
	if move == "LEFT" {
		point.X -= 1
	} else if move == "RIGHT" {
		point.X += 1
	} else if move == "DOWN" {
		point.Y += 1
	} else {
		point.Y -= 1
	}
}

/* Function to print permutations of string
   This function takes three parameters:
   1. String
   2. Starting index of the string
   3. Ending index of the string. */
func permute(arr []int, l, r int, col *[][]int) {

	if l == r {
		temparr := make([]int, len(arr))

		copy(temparr, arr)
		*col = append(*col, temparr)
		// fmt.Println(arr)

	} else {
		for i := l; i <= r; i++ {
			arr[l], arr[i] = arr[i], arr[l]
			permute(arr, l+1, r, col)
			arr[l], arr[i] = arr[i], arr[l] // backtrack
		}
	}
}
