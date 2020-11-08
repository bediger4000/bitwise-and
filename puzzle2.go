package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
This problem was asked by Yahoo.

Write a function that returns the bitwise AND of all integers
between M and N, inclusive.
*/

func main() {
	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	tmp := m
	if m > n {
		m = n
		n = tmp
	}

	and := rangeBitwiseAnd(uint(m), uint(n))

	fmt.Printf("m %064b\n", m)
	fmt.Printf("n %064b\n", n)
	fmt.Printf("  %064b\n", and)
}

func rangeBitwiseAnd(m, n uint) uint {
	var a uint = 0
	for m != n {
		m >>= 1
		n >>= 1
		a++
	}
	return m << a
}
