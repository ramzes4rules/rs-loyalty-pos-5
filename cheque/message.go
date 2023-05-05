package Cheque

type Message struct {
	MessageID int64      `xml:"MessageID,attr"` //Внутренний ID сообщения.
	Device    DeviceType `xml:"Device,attr"`    // Тип скидки:
	Body      string     `xml:"Body,attr"`      // Текст сообщения.
}
