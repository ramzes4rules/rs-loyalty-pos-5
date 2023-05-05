package Cheque

import "time"

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
	ChequeLines     ChequeLines    `xml:"ChequeLines"`
	Discounts       Discounts      `xml:"Discounts"` // Описание данных по скидкам для этой позиции
	Messages        Messages       `xml:"Messages"`
	//Payments
	//SaleCheque
}
