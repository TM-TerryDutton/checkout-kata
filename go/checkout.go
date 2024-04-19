package checkout

type SpecialOffer struct {
	quantity, price int
}
type SpecialOffers map[string]SpecialOffer
type ProductPrices map[string]int

type Checkout struct {
	products      ProductPrices
	scannedItems  map[string]int
	specialOffers SpecialOffers
}

func (s SpecialOffer) Apply(quantity int) (total int, remaining int) {
	offerCount := quantity / s.quantity
	total = offerCount * s.price
	remaining = quantity - (offerCount * s.quantity)
	return
}

func (c Checkout) ApplySpecialOffer(sku string, quantity int) (total int, remaining int) {
	remaining = quantity
	offer, exists := c.specialOffers[sku]
	if exists {
		total, remaining = offer.Apply(quantity)
	}
	return
}

func (c *Checkout) Scan(sku string) {
	c.scannedItems[sku]++
}

func (c Checkout) GetTotalPrice() int {
	total := 0

	for sku, quantity := range c.scannedItems {
		price := c.products[sku]
		offerTotal, remaining := c.ApplySpecialOffer(sku, quantity)
		total += offerTotal
		total += remaining * price
	}
	return total
}

func NewCheckout(products ProductPrices, specialOffers *SpecialOffers) Checkout {
	checkout := Checkout{
		products:     products,
		scannedItems: map[string]int{},
	}
	if specialOffers != nil {
		checkout.specialOffers = *specialOffers
	}
	return checkout
}
