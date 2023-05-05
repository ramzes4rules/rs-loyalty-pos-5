package Cheque

import (
	"encoding/xml"
	"fmt"
	"time"
)

func (cheque Cheque) Test() string {

	// Заполняем атрибуты чека
	cheque.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	cheque.XmlnsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	cheque.StoreID = 0
	cheque.ShiftNo = "001/001"
	cheque.ChequeUID = "qwerty"
	cheque.ChequeNo = "34563"
	cheque.OpenTime = time.Now()
	cheque.CloseTime = time.Now()
	cheque.Amount = 1342.00
	cheque.SubtractedBonus = 32.0
	cheque.PositionCount = 3
	cheque.Status = StatusOpen
	cheque.ChequeType = ChequeTypeSale
	cheque.DiscountCard = append(cheque.DiscountCard, DiscountCard{
		DiscountCardNo:       "3846656766",
		SubtractAmount:       10,
		BonusCard:            false,
		EnteredAsPhoneNumber: true,
	})
	cheque.Coupon = append(cheque.Coupon, Coupon{CouponNo: "9900003221"})

	// Добавляем строки чека
	var cl = ChequeLine{
		ChequeLineNo:  1,
		NoPayBonus:    false,
		NoAddBonus:    false,
		NoDiscounts:   false,
		Price:         100,
		Quantity:      4,
		Amount:        360,
		MinAmount:     0,
		MinPrice:      0,
		MaxDiscount:   100,
		BonusDiscount: 0,
		Item:          Item{ItemID: 0, ItemUID: "{D8FDA393-A6C0-4BA8-BF85-20269E58FC6A}", Barcode: "4612345234412"},
		Discounts: Discounts{Discounts: []Discount{{
			DiscountID: 1133,
			Type:       DiscountTypePercent,
			Percent:    0,
			Amount:     9,
		}, {
			DiscountID: 1134,
			Type:       DiscountTypeAmount,
			Percent:    0,
			Amount:     1,
		}}},
		Coupon: []Coupon{{CouponNo: "9900000043212"}, {CouponNo: "9900000043211"}},
	}
	var cls []ChequeLine
	cls = append(cls, cl)
	var clss ChequeLines
	clss.ChequeLines = cls
	cheque.ChequeLines = clss

	// Заполянем дисконт по чеку
	cheque.Discounts = Discounts{Discounts: []Discount{{
		DiscountID: 1022,
		Type:       DiscountTypeFixPrice,
		Percent:    0,
		Amount:     10,
	}, {
		DiscountID: 1021,
		Type:       DiscountTypeAmount,
		Percent:    0,
		Amount:     5,
	}}}

	// Заполянем сообщения в чеке
	cheque.Messages = Messages{Messages: []Message{{
		MessageID: 556,
		Device:    DeviceTypeCheque,
		Body:      "Сообщение на слипе",
	}, {
		MessageID: 567,
		Device:    DeviceTypeCashierDisplay,
		Body:      "Сообщение на экране кассира",
	}}}

	// Сериализуем чек в xml-формате
	var response, err = xml.MarshalIndent(cheque, "", "   ")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Возвращаем чек обратно
	return string(response)
}
