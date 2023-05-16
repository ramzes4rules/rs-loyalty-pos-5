package Cheque

import "encoding/xml"

type Lines struct {
	XMLName     xml.Name `xml:"ChequeLines"`
	ChequeLines []Line   `xml:"ChequeLine"`
}

type Line struct {
	XMLName       xml.Name  `xml:"ChequeLine"`
	ChequeLineNo  int       `xml:"ChequeLineNo,attr"`  //
	NoPayBonus    bool      `xml:"NoPayBonus,attr"`    // Флаг признака не оплачивать бонусами эту позицию
	NoAddBonus    bool      `xml:"NoAddBonus,attr"`    // Флаг признака не начислять бонусы за эту позицию
	NoDiscounts   bool      `xml:"NoDiscounts,attr"`   // Флаг признака не использовать скидки для этой позиции
	Price         float32   `xml:"Price,attr"`         // Цена товара по позиции
	Quantity      float32   `xml:"Quantity,attr"`      // Количество товара по позиции
	Amount        float32   `xml:"Amount,attr"`        // Итоговая сумма позиции БЕЗ учета скидки бонусами, но с учетом процентных скидок.
	MinAmount     float32   `xml:"MinAmount,attr"`     // Минимальная сумма по позиции
	MinPrice      float32   `xml:"MinPrice,attr"`      // Минимальная цена для позиции
	MaxDiscount   float32   `xml:"MaxDiscount,attr"`   // Максимальная скидка по позиции в процентах
	BonusDiscount float32   `xml:"BonusDiscount,attr"` // Сумма скидки бонусами по позиции примененная после списания бонусов
	Item          Item      `xml:"Item"`               // Товар
	Discounts     Discounts `xml:"Discounts"`          // Описание данных по скидкам для этой позиции
	Coupon        []Coupon  `xml:"Coupon"`             // Примененный купон
}

//func (line *ChequeLine) AddCoupon(number string) {
//	line.Coupon = append(line.Coupon, Coupon{CouponNo: number})
//}
