package rsloyaltypos5

import (
	"encoding/xml"
	"strings"
)

type SubtractBonus45Request struct {
	XMLName            xml.Name `xml:"SubtractBonus45"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber"`
	Amount             float64  `xml:"amount"`
	Cheque             string   `xml:"cheque"`
}

type SubtractBonus45Response struct {
	XMLName               xml.Name `xml:"SubtractBonus45Response"`
	SubtractBonus45Result string   `xml:"SubtractBonus45Result"`
}

type SubtractedChequelines struct {
	XMLName     xml.Name                `xml:"ChequeLines"`
	ChequeLines []SubtracktedChequeline `xml:"ChequeLine"`
}

type SubtracktedChequeline struct {
	XMLName      xml.Name `xml:"ChequeLine"`
	ChequeLineNo int      `xml:"ChequeLineNo,attr"`
	Amount       float64  `xml:"Amount,attr"`
}

func (pos *RSLoyaltyPOS5) SubtractBonus45(discountCardNumber string, amount float64, cheque string) ([]SubtracktedChequeline, error) {

	//
	var request = &SubtractBonus45Request{
		Xmlns:              "http://tempuri.org/",
		Cheque:             cheque,
		Amount:             amount,
		DiscountCardNumber: discountCardNumber,
	}
	var response = &SubtractBonus45Response{}

	// run request
	err := pos.soap(request, "SubtractBonus45", response)

	// check execution errors
	if err != nil {
		//fmt.Printf("Error: %v", err)
		return []SubtracktedChequeline{}, err
	}

	var res = SubtractedChequelines{}
	err = xml.Unmarshal([]byte(strings.Replace(response.SubtractBonus45Result, "<?xml version=\"1.0\" encoding=\"utf-16\"?>", "", -1)), &res)
	if err != nil {
		return []SubtracktedChequeline{}, err
	}
	return res.ChequeLines, nil

}
