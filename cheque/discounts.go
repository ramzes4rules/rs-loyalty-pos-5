package Cheque

type DiscountType string

const (
	DiscountTypePercent  DiscountType = "Percent"
	DiscountTypeAmount   DiscountType = "Amount"
	DiscountTypeTender   DiscountType = "Tender"
	DiscountTypeFixPrice DiscountType = "FixPrice"
)

type Discounts struct {
	Discounts []Discount `xml:"Discount"`
}
