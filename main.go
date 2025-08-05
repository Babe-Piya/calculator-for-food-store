package main

import (
	"fmt"
)

func main() {
	fmt.Println(CalculateOrder(Order{
		RedSet:    0,
		GreenSet:  0,
		BlueSet:   0,
		YellowSet: 0,
		PinkSet:   2,
		PurpleSet: 0,
		OrangeSet: 0,
		IsMember:  false,
	}))

}
