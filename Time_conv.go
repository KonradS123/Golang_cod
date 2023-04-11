package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "12:00:00AM"
	var time string
	h, _ := strconv.Atoi(s[0:2])
	fmt.Print(h, "\n")
	t := int(s[8])
	if t == 80 {
		 if h == 12 && t == 80 {
			h = h
		} else {
			h += 12
		}
	} else {
		if h == 12 {
			h = 00
		}
	}
	h1 := strconv.Itoa(h)
	if h < 10 || h == 0 {
		time = "0" + h1 + ":" + s[3:5] + ":" + s[6:8]
	} else {
		time = h1 + ":" + s[3:5] + ":" + s[6:8]
	}
	fmt.Print(time)
}
