package main

import "fmt"

func main() {
	var n int
	var a [10]int
	var sum int
	fmt.Printf("Enter size of array ( 1 - 9) : ")
	fmt.Scan(&n)

	fmt.Printf("Enter array elements : ")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < n; i++ {
		sum = sum + a[i]
	}
	fmt.Println("Sum of array is ", sum)
}
