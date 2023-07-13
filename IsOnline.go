package rsloyaltypos5

import "encoding/xml"

type IsOnlineRequest struct {
	XMLName xml.Name `xml:"IsOnline"`
	Xmlns   string   `xml:"xmlns,attr"`
}

type IsOnlineResponse struct {
	XMLName xml.Name `xml:"IsOnlineResponse"`
	Result  bool     `xml:"IsOnlineResult"`
}

// IsOnline checking the availability of the service RS.Loyalty.Store from the client side
func (pos *RSLoyaltyPOS5) IsOnline() (*bool, error) {

	var request = &IsOnlineRequest{Xmlns: "http://tempuri.org/"}
	var response = &IsOnlineResponse{}

	err := pos.soap(request, "IsOnline", response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}
