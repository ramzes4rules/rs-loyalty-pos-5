package Cheque

type DiscountCard struct {
	DiscountCardID       int     `xml:"DiscountCardID,attr"`
	DiscountCardNo       string  `xml:"DiscountCardNo,attr"`
	SubtractAmount       float32 `xml:"SubtractAmount,attr"`
	BonusCard            bool    `xml:"BonusCard,attr"`
	EnteredAsPhoneNumber bool    `xml:"EnteredAsPhoneNumber,attr"`
	SubtractedBonus      float32 `xml:"SubtractedBonus,attr"`
}
