package main

// lalalal
import (
	"fmt"
)

var count = 0

var n, m int //lola
var l = 5

func lol() int {
	count++
	count++
	count++
	count++
	r := 6
	r = 8
	var k = 10
	var d = 5
	return r + d + k
}

func main() {
	count++
	count++
	count++
	count++
	flag := 1
	n = 3
	m = 4
	for m > 0 {
		count++
		count++
		n = 2*n + 1
		m = m / 2
	}
	if flag == 1 {
		fmt.Printf("EQUAL")
	} else {
		count++
		count++
		m = 6
		n = 8
		fmt.Printf("NOT EQUAL")
	}
	fmt.Printf("Count of Assignments = " + fmt.Sprint(count))

}

/*
=> EQUAL */
