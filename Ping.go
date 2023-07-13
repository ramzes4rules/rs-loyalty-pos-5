package rsloyaltypos5

import (
	"encoding/xml"
)

type PingRequest struct {
	XMLName xml.Name `xml:"Ping"`
	Xmlns   string   `xml:"xmlns,attr"`
}
type PingResponse struct {
	XMLName xml.Name `xml:"PingResponse"`
}

func (pos *RSLoyaltyPOS5) Ping() error {

	var request = &PingRequest{Xmlns: "http://tempuri.org/"}
	var response = &PingResponse{}

	err := pos.soap(request, "Ping", response)
	if err != nil {
		return err
	}

	return nil
}
