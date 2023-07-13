package rsloyaltypos5

import "encoding/xml"

type AccrualRequest struct {
	XMLName xml.Name `xml:"Accrual"`
	Xmlns   string   `xml:"xmlns,attr"`
	Cheque  string   `xml:"cheque"`
}

type AccrualResponse struct {
	XMLName       xml.Name `xml:"AccrualResponse"`
	AccrualResult string   `xml:"AccrualResult"`
}

// Accrual closes cheque and returns slip message (*string, error)
func (pos *RSLoyaltyPOS5) Accrual(cheque string) (*string, error) {

	//
	var request = &AccrualRequest{Xmlns: "http://tempuri.org/", Cheque: cheque}
	var response = &AccrualResponse{}

	// run request
	err := pos.soap(request, "Accrual", response)

	// check execution errors
	if err != nil {
		return nil, err
	}

	return &response.AccrualResult, nil
}
