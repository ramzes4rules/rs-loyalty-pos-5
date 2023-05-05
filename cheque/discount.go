package Cheque

type Discount struct {
	DiscountID int32        `xml:"DiscountID,attr"` // Внутренний ID скидки.
	Type       DiscountType `xml:"Type,attr"`       // Тип скидки:
	Percent    float32      `xml:"Percent,attr"`    // Величина процента скидки
	Amount     float32      `xml:"Amount,attr"`     // Величина скидки в единицах оплаты
}
