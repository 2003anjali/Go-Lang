// program to find sum of square of number in go lang
package main

import "fmt"

func main() {
	var i, n, sum int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	sum = 0
	for i = 1; i <= n; i++ {
		sum = sum + (i * i)
	}
	fmt.Printf("Sum of square of %d is %d", n, sum)
}
