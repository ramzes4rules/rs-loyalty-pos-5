package rsloyaltypos5

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type GetDiscounts struct {
	XMLName xml.Name `xml:"GetCardBalance"`
	Xmlns   string   `xml:"xmlns,attr"`
	Cheque  string   `xml:"cheque,omitempty"`
}

type GetDiscountsResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ GetCardBalanceResponse"`
	//GetCardBalanceResult string   `xml:"GetCardBalanceResult,omitempty"`
}

func (pos *RSLoyaltyPOS5) GetDiscounts(cheque string) error {

	fmt.Printf(cheque)

	var request = &GetDiscounts{
		Xmlns:  "http://tempuri.org/",
		Cheque: cheque,
	}
	var response = &GetCardBalanceResponse{}

	err := pos.soap(request, "GetDiscounts", response)
	if err != nil {
		return err
	}

	var result = Messages{}
	err = xml.Unmarshal([]byte(strings.Replace(response.GetCardBalanceResult, ",", ".", -1)), &result)
	if err != nil {
		return err
	}
	return nil
}
