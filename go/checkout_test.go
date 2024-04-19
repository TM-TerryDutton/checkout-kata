package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckout(t *testing.T) {
	testCases := []struct {
		description   string
		catalogue     ProductPrices
		specialOffers *SpecialOffers
		scannedItems  []string
		expectedTotal int
	}{
		{
			description:   "Scans product",
			catalogue:     map[string]int{"A": 50},
			scannedItems:  []string{"A"},
			expectedTotal: 50,
		},
		{
			description:   "Scans multiple products",
			catalogue:     map[string]int{"A": 50},
			scannedItems:  []string{"A", "A", "A"},
			expectedTotal: 150,
		},
		{
			description:   "Scans different products",
			catalogue:     map[string]int{"A": 50, "B": 30},
			scannedItems:  []string{"A", "B"},
			expectedTotal: 80,
		},
		{
			description: "Applies special offers",
			catalogue:   map[string]int{"A": 50},
			specialOffers: &SpecialOffers{
				"A": SpecialOffer{quantity: 3, price: 130},
			},
			scannedItems:  []string{"A", "A", "A", "A"},
			expectedTotal: 180,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			checkout := NewCheckout(tc.catalogue, tc.specialOffers)

			for _, item := range tc.scannedItems {
				checkout.Scan(item)
			}
			total := checkout.GetTotalPrice()

			assert.Equal(t, tc.expectedTotal, total)
		})
	}
}
