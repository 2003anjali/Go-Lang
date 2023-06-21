package main

import (
	"fmt"
	"strconv"
) 
func ispalindrome(n int)bool{   //(n int)-> input type , bool -> return type
	str:= strconv.Itoa(n) //integer to ascii
	length:=len(str)
	for i:=0;i<length/2;i++{
		if str[i]!=str[length-i-1]{
			return false
		}
	}
	return true
}
func largestpalindromeproduct() (int,int,int){
	largestpalindrome:=0
	var multiplicand1, multiplicand2 int
	for i:=99;i>=10;i-- { //since we want the largest so we are starting from the last 
		for j:=i;j>=10   ;j--{
			product :=i*j 
			if product< largestpalindrome{
				break
			}
			if ispalindrome(product)&&product>largestpalindrome{
				largestpalindrome = product
				multiplicand1 = i 
				multiplicand2 = j 
			}
		}
	}
	return largestpalindrome , multiplicand1, multiplicand2
}

func main(){
	result, multiplicand1, multiplicand2 := largestpalindromeproduct()
	fmt.Println("The largest palindrome product is ",result)
	fmt.Printf("The multiplicands are : %d and %d ", multiplicand1,multiplicand2)

}