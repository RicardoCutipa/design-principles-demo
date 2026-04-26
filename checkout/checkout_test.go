package checkout

import "testing"

func TestTotal_StacksRulesAndCapsAtSubtotal(t *testing.T) {
	cart := Cart{Subtotal: 12500} // $125.00
	got := Total(cart, TenPercentOver100{}, FiveDollarsOff{})

	// 10% of 12500 = 1250, plus 500 = 1750 => total 10750
	const want Money = 10750
	if got != want {
		t.Fatalf("Total() = %d, want %d", got, want)
	}

	// Cap behavior: discount cannot exceed subtotal.
	zero := Cart{Subtotal: 300}
	got = Total(zero, FiveDollarsOff{})
	if got != 0 {
		t.Fatalf("Total() with subtotal cap = %d, want 0", got)
	}
}

