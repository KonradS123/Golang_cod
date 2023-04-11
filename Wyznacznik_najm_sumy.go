package main

import "fmt"

func main() {

	var sum int
	var ar = []int{}
	var arr = [...]int{7, 69, 2, 221, 8974}
	for i := 0; i < len(arr); i++ {
		for b := 0; b < len(arr); b++ {
			sum = sum + arr[b]
			fmt.Print("\n", "b=", sum)
		}
		sum = sum - arr[i]
		ar = append(ar, sum)
		fmt.Print("\n", "i=", sum)
		sum = 0
	}
	fmt.Print("\n", "poza ", sum)
	fmt.Print("\n", ar)
	var max int = 0
	var min int = 0
	//var b int = 0
	min = ar[0]
	max = ar[0]
	for _, value := range ar {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	fmt.Print(max, min)
}
