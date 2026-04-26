package checkout

type TenPercentOver100 struct{}

func (TenPercentOver100) Discount(c Cart) Money {
	if c.Subtotal < 10000 { // $100.00
		return 0
	}
	return c.Subtotal / 10
}

type FiveDollarsOff struct{}

func (FiveDollarsOff) Discount(c Cart) Money {
	const five Money = 500
	if c.Subtotal < five {
		return c.Subtotal
	}
	return five
}

