package Cheque

import (
	"encoding/xml"
	"fmt"
)

func (cheque *Cheque) Xml() string {
	//var response, err = xml.MarshalIndent(cheque, "", "   ")
	var response, err = xml.Marshal(cheque)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return "<?xml version=\"1.0\" encoding=\"utf-16\"?>" + string(response)
}
