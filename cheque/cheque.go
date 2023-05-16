package Cheque

import "time"

type Type string

const (
	ChequeTypeSale   Type = "Sale"
	ChequeTypeReturn Type = "Return"
)

type Status string

const (
	StatusClosed    Status = "Closed "
	StatusOpen      Status = "Open"
	StatusCancelled Status = "Cancelled"
)

type Cheque struct {
	XmlnsXsd        string         `xml:"xmlns:xsd,attr"`       //:xsd="http://www.w3.org/2001/XMLSchema"
	XmlnsXsi        string         `xml:"xmlns:xsi,attr"`       //:xsi="http://www.w3.org/2001/XMLSchema-instance"
	StoreID         int32          `xml:"StoreID,attr"`         //
	ShiftNo         string         `xml:"ShiftNo,attr"`         // Номер смены. Не используется.
	ChequeUID       string         `xml:"ChequeUID,attr"`       // Обязательный уникальный идентификатор чека
	ChequeNo        string         `xml:"ChequeNo,attr"`        // Номер чека на кассе
	OpenTime        time.Time      `xml:"OpenTime,attr"`        // Время открытия чека
	CloseTime       time.Time      `xml:"CloseTime,attr"`       // Время закрытия чека
	Amount          float32        `xml:"Amount,attr"`          // Сумма чека
	SubtractedBonus float32        `xml:"SubtractedBonus,attr"` // Текущий вычет бонусов из чека в денежных единицах
	PositionCount   int32          `xml:"PositionCount,attr"`   // Количество позиций в чеке
	Status          Status         `xml:"Status,attr"`          // Статус чека
	ChequeType      Type           `xml:"ChequeType,attr"`      // Тип чека
	DiscountCard    []DiscountCard `xml:"DiscountCard"`         // Описание данных по дисконтным картам
	Coupon          []Coupon       `xml:"Coupon"`               // Описание данных купонов
	ChequeLines     Lines          `xml:"ChequeLines"`
	Discounts       Discounts      `xml:"Discounts"` // Описание данных по скидкам для этой позиции
	Messages        Messages       `xml:"Messages"`
	//Payments
	//SaleCheque
}

func random() {

}

func New(number string, datetime time.Time) Cheque {
	var cheque = Cheque{}

	cheque.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	cheque.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	cheque.ChequeNo = number
	cheque.ChequeUID = "0987654321"
	cheque.Status = StatusOpen
	cheque.OpenTime = datetime
	cheque.ChequeType = ChequeTypeSale
	cheque.ShiftNo = "001"
	cheque.StoreID = 1

	return cheque
}

func (cheque *Cheque) AddDiscountCard(number string) {
	cheque.DiscountCard = append(cheque.DiscountCard, DiscountCard{
		DiscountCardNo:       number,
		SubtractAmount:       0,
		BonusCard:            true,
		EnteredAsPhoneNumber: false,
		SubtractedBonus:      0,
	})
}
func (cheque *Cheque) AddCoupon(number string) {
	cheque.Coupon = append(cheque.Coupon, Coupon{CouponNo: number})
}

func (cheque *Cheque) AddLine(item string, price float32, quantity float32, coupons []Coupon) {
	var line = Line{}
	line.ChequeLineNo = len(cheque.ChequeLines.ChequeLines) + 1
	line.Item.ItemUID = item
	line.Price = price
	line.Quantity = quantity
	line.Amount = price * quantity
	line.MaxDiscount = 100
	for _, coupon := range coupons {
		line.Coupon = append(line.Coupon, Coupon{CouponNo: coupon.CouponNo})
	}

	cheque.PositionCount++
	cheque.Amount += line.Amount

	cheque.ChequeLines.ChequeLines = append(cheque.ChequeLines.ChequeLines, line)

}
