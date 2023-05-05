package Cheque

type Item struct {
	ItemID  int64  `xml:"ItemID,attr"`  // Внутренний идентификатор товара (RS.Loyalty). В случае если этот код неизвестен, необходимо устанавливать значение 0.
	ItemUID string `xml:"ItemUID,attr"` // Уникальный идентификатор товара во внешней системе. 	Формат может быть любым. Максимум 50 символов.
	Barcode string `xml:"Barcode,attr"` // Штрих код товара.
}
