package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DirtyPoint struct{
	X int
	Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Open("/Users/peter/Projects/GO/src/github.com/peterFran/hacker/test.txt")
	check(err)
	reader := bufio.NewReader(f)
	meta, _, _ := reader.ReadLine()
	x, _ := strconv.Atoi(string(meta[0]))
	y, _ := strconv.Atoi(string(meta[2]))
	// start = Coordinates(x, y)

	fmt.Println(x, y)
	for line,hasMore,_ := reader.ReadLine() ; hasMore == true ; {
		for i:=0 ; i < len(line) ; i++{
			
		}
	}
	N := len(line)
	for i := 0 ; i < 
	
	fmt.Println(string(N))
}

// func createGrid(dimensions int) {

// }
