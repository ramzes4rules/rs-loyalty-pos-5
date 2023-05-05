package Cheque

import (
	"encoding/xml"
	"fmt"
)

func (cheque *Cheque) Xml() string {
	var response, err = xml.MarshalIndent(cheque, "", "   ")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(response)
}
