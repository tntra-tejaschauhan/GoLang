package main

import "fmt"
import (
	"main/concurrency"
)
var a int = 10
var b int = 33
func mainto() {

	fmt.Println("Hi There")
}
func printSum() {
	// Maint02()
	fmt.Println("The sum is ", a+b)
}

func main() {
	var nums [4]int 
	mainto()
	printSum()
	concurrency.Main1()
	nums[0] = 1
	fmt.Println(len(nums))  
	fmt.Println(nums[0])
	fmt.Println(nums)
}
