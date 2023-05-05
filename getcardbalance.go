package rsloyaltypos5

import (
	"encoding/xml"
	"strings"
)

type GetCardBalance struct {
	XMLName            xml.Name `xml:"GetCardBalance"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber,omitempty"`
}

type GetCardBalanceResponse struct {
	XMLName              xml.Name `xml:"http://tempuri.org/ GetCardBalanceResponse"`
	GetCardBalanceResult string   `xml:"GetCardBalanceResult,omitempty"`
}

type Messages struct {
	XMLName xml.Name `xml:"Messages"`
	Balance Balance
	Msg     Msg
}

type Balance struct {
	Value float32 `xml:"Value,attr"`
}

type Msg struct {
	Device int    `xml:"Device,attr"`
	Body   string `xml:"Body,attr"`
}

func (pos *RSLoyaltyPOS5) GetCardBalance(number string) (*Messages, error) {

	var request = &GetCardBalance{
		Xmlns:              "http://tempuri.org/",
		DiscountCardNumber: number,
	}
	var response = &GetCardBalanceResponse{}

	err := pos.soap(request, "GetCardBalance", response)
	if err != nil {
		return nil, err
	}

	var result = Messages{}
	err = xml.Unmarshal([]byte(strings.Replace(response.GetCardBalanceResult, ",", ".", -1)), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
