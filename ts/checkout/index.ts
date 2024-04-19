export type Products = { [sku: string]: number };
export type SpecialOffers = { [sku: string]: SpecialOffer };
type SpecialOfferResult = { totalPrice: number, remaining: number };

export class SpecialOffer {
  constructor(
    private qualifyingQuantity: number, 
    private price: number
  ) {}

  public apply(quantity: number): SpecialOfferResult {
    const totalPrice = Math.floor(quantity / this.qualifyingQuantity) * this.price;
    const remaining = quantity % this.qualifyingQuantity;
    return { totalPrice, remaining  };
  }
}

export class Checkout {
  private scannedItems: Products = {}

  constructor(
    private products: Products, 
    private specialOffers: SpecialOffers = {}
  ) {}

  public scan(sku: string) {
    this.scannedItems[sku] = (this.scannedItems[sku] || 0) + 1;
  }

  public getTotalPrice (): number {
    let total = 0;
    for (const sku in this.scannedItems) {
      const price = this.products[sku];
      const quantity = this.scannedItems[sku];
      const { totalPrice, remaining } = this.applySpecialOffer(sku, quantity);
      total += totalPrice + (price * remaining);
    }
    return total;
  }

  private applySpecialOffer (sku: string, quantity: number): SpecialOfferResult {
    const specialOffer = this.specialOffers[sku];
    if (specialOffer) return specialOffer.apply(quantity);
    return { totalPrice: 0, remaining: quantity };
  }
}