package rsloyaltypos5

import "encoding/xml"

type IsCardValidRequest struct {
	XMLName            xml.Name `xml:"IsCardValid"`
	Xmlns              string   `xml:"xmlns,attr"`
	DiscountCardNumber string   `xml:"discountCardNumber"`
}

type IsCardValidResponse struct {
	XMLName           xml.Name `xml:"IsCardValidResponse"`
	IsCardValidResult bool     `xml:"IsCardValidResult"`
}

func (pos *RSLoyaltyPOS5) IsCardValid(number string) (*bool, error) {

	var request = &IsCardValidRequest{
		Xmlns:              "http://tempuri.org/",
		DiscountCardNumber: number,
	}
	var response = &IsCardValidResponse{}

	err := pos.soap(request, "IsCardValid", response)
	if err != nil {
		return nil, err
	}

	return &response.IsCardValidResult, nil
}
