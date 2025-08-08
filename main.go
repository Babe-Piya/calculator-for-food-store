package main

import (
	"fmt"
)

func main() {
	order := NewOrder()
	fmt.Println(order.CalculateOrder(OrderRequest{
		Items: []Item{{
			Menu:  "Orange Set",
			Count: 4,
		}},
		IsMember: false,
	}))

}
