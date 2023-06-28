package rsloyaltypos5

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
	"unicode/utf16"
	_ "unicode/utf16"
	"unicode/utf8"
	_ "unicode/utf8"
)

type GetMessagesRequest struct {
	XMLName xml.Name `xml:"GetMessages"`
	Xmlns   string   `xml:"xmlns,attr"`
	Cheque  string   `xml:"cheque"`
}

type GetMessagesResponse struct {
	XMLName           xml.Name `xml:"GetMessagesResponse"`
	GetMessagesResult string   `xml:"GetMessagesResult"`
}

type Messages2 struct {
	XMLName  xml.Name `xml:"Messages"`
	Messages []Msg2   `xml:"Msg"`
}

type Msg2 struct {
	XMLName   xml.Name `xml:"Msg"`
	MessageID string   `xml:"MessageID,attr"`
	Device    string   `xml:"Device,attr"`
	Body      string   `xml:"Body,attr"`
}

func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

func (pos *RSLoyaltyPOS5) GetMessages(cheque string) (Messages2, error) {

	// declare variables
	var request = &GetMessagesRequest{
		Xmlns:  "http://tempuri.org/",
		Cheque: cheque,
	}
	var response = &GetMessagesResponse{}

	// exec request
	err := pos.soap(request, "GetMessages", response)

	// check execution errors
	if err != nil {
		return Messages2{}, err
	}
	//fmt.Printf("res: %s\n", response.GetMessagesResult)

	//
	var result = Messages2{}
	err = xml.Unmarshal([]byte(strings.Replace(response.GetMessagesResult, "<?xml version=\"1.0\" encoding=\"utf-16\"?>", "", -1)), &result)
	if err != nil {
		return Messages2{}, err
	}

	return result, nil

}
