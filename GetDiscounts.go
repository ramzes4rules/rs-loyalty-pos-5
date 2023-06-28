package rsloyaltypos5

import (
	"encoding/xml"
	Cheque "github.com/ramzes4rules/rs-loyalty-pos-5/cheque"
	"strings"
)

type GetDiscounts struct {
	XMLName xml.Name `xml:"GetDiscounts"`
	Xmlns   string   `xml:"xmlns,attr"`
	Cheque  string   `xml:"cheque,omitempty"`
}

type GetDiscountsResponse struct {
	XMLName            xml.Name `xml:"http://tempuri.org/ GetDiscountsResponse"`
	GetDiscountsResult string   `xml:"GetDiscountsResult,omitempty"`
}

type ChequeDiscounts struct {
	XMLName     xml.Name          `xml:"ChequeDiscounts"`
	ChequeLines IncomeChequeLines `xml:"ChequeLines"`
}

type IncomeChequeLines struct {
	XMLName     xml.Name           `xml:"ChequeLines"`
	ChequeLines []IncomeChequeLine `xml:"ChequeLine"`
}

type IncomeChequeLine struct {
	XMLName      xml.Name        `xml:"ChequeLine"`
	ChequeLineNo int             `xml:"ChequeLineNo"`
	TotalAmount  float32         `xml:"TotalAmount"`
	Discounts    IncomeDiscounts `xml:"Discounts"`
}

type IncomeDiscounts struct {
	XMLName   xml.Name         `xml:"Discounts"`
	Discounts []IncomeDiscount `xml:"Discount"`
}

type IncomeDiscount struct {
	XMLName    xml.Name            `xml:"Discount"`
	DiscountID int                 `xml:"DiscountID,attr"`
	Type       Cheque.DiscountType `xml:"Type,attr"`
	Percent    float32             `xml:"Percent,attr"`
	Amount     float32             `xml:"Amount,attr"`
}

//<Cheque TotalAmount="0">

//<?xml version="1.0" encoding="utf-16"?>
//<ChequeDiscounts xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
//<ChequeLines>
//<ChequeLine ChequeLineNo="1" TotalAmount="0">
//<Discounts>
//<Discount DiscountID="8" Type="Percent" Percent="0" Amount="51"/>
//</Discounts>
//</ChequeLine>
//<ChequeLine ChequeLineNo="2" TotalAmount="0">
//<Discounts>
//<Discount DiscountID="8" Type="Percent" Percent="0" Amount="51"/>
//</Discounts>
//</ChequeLine>
//</ChequeLines>
//</ChequeDiscounts>

func (pos *RSLoyaltyPOS5) GetDiscounts(cheque string) (ChequeDiscounts, error) {

	// creating instance
	var request = &GetDiscounts{
		Xmlns:  "http://tempuri.org/",
		Cheque: cheque,
	}
	var response = &GetDiscountsResponse{}

	// exec request
	err := pos.soap(request, "GetDiscounts", response)

	// check execution error
	if err != nil {
		return ChequeDiscounts{}, err
	}

	//fmt.Printf(response.GetDiscountsResult)

	var result = ChequeDiscounts{}
	err = xml.Unmarshal([]byte(strings.Replace(response.GetDiscountsResult, "<?xml version=\"1.0\" encoding=\"utf-16\"?>", "", -1)), &result)
	if err != nil {
		return ChequeDiscounts{}, err
	}

	return result, nil
}
