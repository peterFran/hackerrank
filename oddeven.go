package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T int
	f, _ := os.Open("/Users/petermeckiffe/Projects/go/src/github.com/peterfran/hackerrank/test8.txt")
	scanner := bufio.NewReader(f)
	// scanner := bufio.NewReader(os.Stdin)
	line, _, _ := scanner.ReadLine()
	T, _ = strconv.Atoi(string(line))
	for i := 0; i < T; i++ {
		line, _, _ = scanner.ReadLine()

		str := string(line)
		fmt.Println(len(line))
		fmt.Println(len(str))
		strEven := ""
		strOdd := ""
		for j := 0; j < len(str); j++ {
			if j%2 == 0 {
				strEven += string(str[j])
			} else {
				strOdd += string(str[j])
			}
		}
		fmt.Printf("%s %s\n", strEven, strOdd)
	}
}
