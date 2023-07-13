package rsloyaltypos5

import "encoding/xml"

type GetCardDiscountAmountRequest struct {
	XMLName            xml.Name `xml:"GetCardDiscountAmount"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber"`
	Cheque             string   `xml:"cheque"`
}

type GetCardDiscountAmountResponse struct {
	XMLName           xml.Name `xml:"GetCardDiscountAmountResponse"`
	GetMessagesResult float64  `xml:"GetCardDiscountAmountResult"`
}

func (pos *RSLoyaltyPOS5) GetCardDiscountAmount(discountCardNumber string, cheque string) (float64, error) {

	// declare variables
	var request = &GetCardDiscountAmountRequest{
		Xmlns:              "http://tempuri.org/",
		DiscountCardNumber: discountCardNumber,
		Cheque:             cheque,
	}
	var response = &GetCardDiscountAmountResponse{}

	// run request
	err := pos.soap(request, "GetCardDiscountAmount", response)

	// check execution errors
	if err != nil {
		return 0, err
	} else {
		return response.GetMessagesResult, nil
	}
}
