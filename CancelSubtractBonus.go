package rsloyaltypos5

import (
	"encoding/xml"
)

type CancelSubtractBonusRequest struct {
	XMLName            xml.Name `xml:"CancelSubtractBonus"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber"`
	Amount             float64  `xml:"amount"`
	Cheque             string   `xml:"cheque"`
}

type CancelSubtractBonusResponse struct {
	XMLName                   xml.Name `xml:"CancelSubtractBonusResponse"`
	CancelSubtractBonusResult string   `xml:"CancelSubtractBonusResult"`
}

func (pos *RSLoyaltyPOS5) CancelSubtractBonus(discountCardNumber string, amount float64, cheque string) error {
	//
	var request = &CancelSubtractBonusRequest{
		Xmlns:              "http://tempuri.org/",
		Cheque:             cheque,
		Amount:             amount,
		DiscountCardNumber: discountCardNumber,
	}
	var response = &CancelSubtractBonusResponse{}

	// run request
	err := pos.soap(request, "CancelSubtractBonus", response)

	// check execution errors
	if err != nil {
		//fmt.Printf("Error: %v", err)
		return err
	}
	//fmt.Printf("R: %s\n", response.CancelSubtractBonusResult)

	return nil
}
