package main

import "fmt"

func main() {
	var arr = [...]int{44, 53, 31, 27, 77, 60, 66, 77, 26, 36}
	var ilosc int = 0
	var max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	for _, value := range arr {
		if max == value {
			ilosc += 1
		}
	}
	fmt.Print(arr, "\n")
	fmt.Print(ilosc, "\n")
	fmt.Print(max)
}
