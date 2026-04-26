package main

import (
	"fmt"

	"github.com/RicardoCutipa/design-principles-demo/checkout"
)

func main() {
	cart := checkout.Cart{Subtotal: 12500} // $125.00
	total := checkout.Total(cart, checkout.TenPercentOver100{}, checkout.FiveDollarsOff{})
	fmt.Printf("Subtotal: %d cents, Total after discounts: %d cents\n", cart.Subtotal, total)
}

