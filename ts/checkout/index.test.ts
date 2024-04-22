import { Checkout, SpecialOffer, type ProductPrices, type SpecialOffers } from ".";

describe('Checkout', () => {
  const testCases: {
    description: string;
    products: ProductPrices;
    toScan: string,
    specialOffers?: SpecialOffers;
    expectedTotal: number;
  }[] = [
    {
      description: "Scans items",
      products: { "a": 50 },
      toScan: "aaa",
      expectedTotal: 150
    },
    {
      description: "Scans different items",
      products: { "a": 50, "b": 25 },
      toScan: "aab",
      expectedTotal: 125
    },
    {
      description: "Applies discounts",
      products: { "a": 50, "b": 25 },
      specialOffers: { "a": new SpecialOffer(3, 130) },
      toScan: "aaab",
      expectedTotal: 155
    },
  ];
  
  it.each(testCases)(`$description`, (tc) => {
    const checkout = new Checkout(tc.products, tc.specialOffers);
    tc.toScan.split('').forEach((sku) => checkout.scan(sku));
    const total = checkout.getTotalPrice();
    expect(total).toEqual(tc.expectedTotal);
  });
});