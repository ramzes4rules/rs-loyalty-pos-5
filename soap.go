package rsloyaltypos5

import (
	b64 "encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Fault struct {
	Faultcode   string `xml:"faultcode,omitempty"`
	Faultstring string `xml:"faultstring,omitempty"`
}

type RequestEnvelope struct {
	XMLName xml.Name    `xml:"soap:Envelope"`
	Xmlns   string      `xml:"xmlns:soap,attr"`
	Body    RequestBody `xml:"soap:Body"`
}
type RequestBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Request any
}

type ResponseEnvelope struct {
	XMLName xml.Name     `xml:"Envelope"`
	Xmlns   string       `xml:"xmlns:soap,attr"`
	Body    ResponseBody `xml:"s:Body"`
}
type ResponseBody struct {
	XMLName  xml.Name `xml:"s:Body"`
	Response any      //`xml:"AboutResponse"`
}

func (pos *RSLoyaltyPOS5) soap(request any, method string, response any) error {

	//
	var fault = Fault{}
	var client = &http.Client{}

	//
	var Envelope = RequestEnvelope{Xmlns: "http://schemas.xmlsoap.org/soap/envelope/"}
	Envelope.Body.Request = request
	req1, err := xml.Marshal(Envelope)
	if err != nil {
		return fmt.Errorf("request marshaling failed: %v", err)
	}
	text := "<?xml version=\"1.0\" encoding=\"utf-8\"?>" + string(req1)
	payload := strings.NewReader(text)

	//
	req, err := http.NewRequest("POST", pos.Url, payload)
	if err != nil {
		return fmt.Errorf("request creation failed: %v", err)
	}

	req.Header.Add("soapAction", fmt.Sprintf("http://tempuri.org/IRSLoyaltyService/%s", method))
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	//fmt.Println(b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", pos.login, pos.password))))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", pos.Login, pos.Password)))))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request execution failed: %v", err)
	}

	//
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("response reading failed: %v", err)
	}
	//fmt.Println(string(body))

	var r = strings.Replace(string(body),
		"<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\"><s:Body>", "", -1)
	r = strings.Replace(r, "</s:Body></s:Envelope>", "", -1)
	//fmt.Println(fmt.Sprintf("Получен результат: '%s'", r))

	// Парсим ответ
	//fmt.Printf("StatusCode: %d\n", res.StatusCode)
	switch res.StatusCode {
	case 400, 401:
		return errors.New(res.Status)
	case 500:
		err = xml.Unmarshal([]byte(r), &fault)
		if err != nil {
			return fmt.Errorf("fault unmarshalling failed: %v", err)
		}
		return fmt.Errorf("got error response: %v", fault.Faultstring)
	default:
		err = xml.Unmarshal([]byte(r), &response)
		if err != nil {
			return fmt.Errorf("response unmarshalling failed: %v", err)
		}
		return nil
	}
}
