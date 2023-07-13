package rsloyaltypos5

import "encoding/xml"

type IsCouponValidRequest struct {
	XMLName      xml.Name `xml:"IsCouponValid"`
	Xmlns        string   `xml:"xmlns,attr"`
	CouponNumber string   `xml:"couponNumber"`
}

type IsCouponValidResponse struct {
	XMLName             xml.Name `xml:"IsCouponValidResponse"`
	IsCouponValidResult bool     `xml:"IsCouponValidResult"`
}

// IsCouponValid returns true/false or exception in case of error
// Возможные варианты исключений:
// •	"Номер купона передан как пустая строка"
// •	"Купон не найден"

func (pos *RSLoyaltyPOS5) IsCouponValid(couponNumber string) (*bool, error) {

	//
	var request = &IsCouponValidRequest{
		Xmlns:        "http://tempuri.org/",
		CouponNumber: couponNumber,
	}
	var response = &IsCouponValidResponse{}

	err := pos.soap(request, "IsCouponValid", response)
	if err != nil {
		return nil, err
	}

	return &response.IsCouponValidResult, nil
}
