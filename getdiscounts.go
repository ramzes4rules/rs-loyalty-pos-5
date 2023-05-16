package rsloyaltypos5

import (
	"encoding/xml"
	"fmt"
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
	XMLName xml.Name `xml:"ChequeDiscounts"`
	//ChequeLines ChequeLines `xml:"ChequeLines"`
}

//type ChequeLine struct {
//	XMLName xml.Name `xml:"ChequeLine"`
//	ChequeLine
//}

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

func (pos *RSLoyaltyPOS5) GetDiscounts(cheque string) error {

	//
	var request = &GetDiscounts{Xmlns: "http://tempuri.org/", Cheque: cheque}
	var response = &GetDiscountsResponse{}

	//
	err := pos.soap(request, "GetDiscounts", response)
	if err != nil {
		return err
	}

	fmt.Printf(response.GetDiscountsResult)
	//var result = Messages{}
	//err = xml.Unmarshal(response, &result)
	//if err != nil {
	//	return err
	//}
	return nil
}
