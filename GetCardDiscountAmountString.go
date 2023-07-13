package rsloyaltypos5

import (
	"encoding/xml"
)

type GetCardDiscountAmountStringRequest struct {
	XMLName            xml.Name `xml:"GetCardDiscountAmountString"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber"`
	Cheque             string   `xml:"cheque"`
}

type GetCardDiscountAmountStringResponse struct {
	XMLName           xml.Name `xml:"GetCardDiscountAmountStringResponse"`
	GetMessagesResult string   `xml:"GetCardDiscountAmountStringResult"`
}

func (pos *RSLoyaltyPOS5) GetCardDiscountAmountString(discountCardNumber string, cheque string) (string, error) {

	// declare variables
	var request = &GetCardDiscountAmountStringRequest{
		Xmlns:              "http://tempuri.org/",
		Cheque:             cheque,
		DiscountCardNumber: discountCardNumber,
	}
	var response = &GetCardDiscountAmountStringResponse{}

	// run request
	err := pos.soap(request, "GetCardDiscountAmountString", response)

	// check execution errors
	if err != nil {
		return "", err
	} else {
		return response.GetMessagesResult, nil
	}
}
