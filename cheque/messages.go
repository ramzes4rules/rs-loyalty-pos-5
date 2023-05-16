package Cheque

type DeviceType string

const (
	DeviceTypeCheque          DeviceType = "Cheque"
	DeviceTypeCustomerDisplay DeviceType = "CustomerDisplay"
	DeviceTypeCashierDisplay  DeviceType = "CashierDisplay"
)

type Messages struct {
	Messages []Message `xml:"Message"`
}

type Message struct {
	MessageID int64      `xml:"MessageID,attr"` //Внутренний ID сообщения.
	Device    DeviceType `xml:"Device,attr"`    // Тип скидки:
	Body      string     `xml:"Body,attr"`      // Текст сообщения.
}
