package checkout

type Money int64 // cents

type Cart struct {
	Subtotal Money
}

type PricingRule interface {
	Discount(cart Cart) Money
}

func Total(cart Cart, rules ...PricingRule) Money {
	var discount Money
	for _, rule := range rules {
		discount += rule.Discount(cart)
	}
	if discount > cart.Subtotal {
		discount = cart.Subtotal
	}
	return cart.Subtotal - discount
}

