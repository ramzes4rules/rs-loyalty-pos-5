package rsloyaltypos5

import (
	"fmt"
	Cheque "github.com/ramzes4rules/rs-loyalty-pos-5/cheque"
	"testing"
	"time"
)

var pos = RSLoyaltyPOS5{
	Url: "http://afanasiy-test.retailloyalty.ru:8383/RSLoyaltyStoreService",
	//Login:    "rsl_reserv",
	Login: "00308301463784",
	//Password: "rsl_reserv",
	Password: "00308301463784",
}
var ValidCard = "7963101205527879"
var InvalidCard = "1234567890"

// TestRSLoyaltyPOS5_Ping tests method Ping, if success return no error
func TestRSLoyaltyPOS5_Ping(t *testing.T) {
	err := pos.Ping()
	if err != nil {
		t.Errorf("%v\n", err)
	}
}

// TestRSLoyaltyPOS5_IsOnline
func TestRSLoyaltyPOS5_IsOnline(t *testing.T) {
	online, err := pos.IsOnline()
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	fmt.Printf("\tIsOnline: %t\n", *online)
}

// TestRSLoyaltyPOS5_IsCardValid validate card by number
func TestRSLoyaltyPOS5_IsCardValid(t *testing.T) {
	valid, err := pos.IsCardValid(ValidCard)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	fmt.Printf("\tIsCardValid: %t\n", *valid)
}

func TestRSLoyaltyPOS5_IsCouponValid(t *testing.T) {
	valid, err := pos.IsCouponValid("9800200000043")
	if err != nil {
		t.Errorf("\tError: %v\n\r", err)
		return
	}
	fmt.Printf("\tCoupon valid: %t\n\r", *valid)
}

func TestRSLoyaltyPOS5_GetCardBalance(t *testing.T) {
	balance, err := pos.GetCardBalance("8641212605463261")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("\tBalance: %5.2f\n", balance.Balance.Value)
}

func TestRSLoyaltyPOS5_GetDiscounts(t *testing.T) {
	//
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("000001")
	cheque.AddCoupon("980001")
	cheque.AddCoupon("980002")
	cheque.AddLine("49000123", 100.0, 2, []Cheque.Coupon{{CouponNo: "99001"}, {CouponNo: "99002"}})
	cheque.AddLine("49000123", 200.0, 2, nil)
	cheque.AddLine("49000123", 300.0, 2, nil)
	//fmt.Println(cheque.Xml())

	dis, err := pos.GetDiscounts(cheque.Xml())
	//err := pos.GetDiscounts("<?xml version=\"1.0\" encoding=\"utf-16\"?>\n\t<Cheque xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" StoreID=\"100\" ShiftNo=\"54\" ChequeUID=\"00100001-0054-0000-0501-683105466661\" ChequeNo=\"5\" OpenTime=\"2023-05-03T12:17:46.66+03:00\" CloseTime=\"0001-01-01T00:00:00\" Amount=\"204.00\" SubtractedBonus=\"0\" PositionCount=\"2\" Status=\"Open\" ChequeType=\"Sale\">\n\t\t<DiscountCard DiscountCardID=\"0\" DiscountCardNo=\"2977397293309\" SubtractAmount=\"0\" BonusCard=\"true\" EnteredAsPhoneNumber=\"false\" SubtractedBonus=\"0\"/>\n\t\t<ChequeLines>\n\t\t\t<ChequeLine ChequeLineNo=\"1\" NoPayBonus=\"false\" NoAddBonus=\"false\" NoDiscounts=\"false\" Price=\"102.00\" Quantity=\"1.000\" Amount=\"102.00\" MinAmount=\"0\" MinPrice=\"0\" MaxDiscount=\"100.00\" BonusDiscount=\"0\">\n\t\t\t\t<Item ItemID=\"0\" ItemUID=\"81020900001\"/>\n\t\t\t\t<Discounts/>\n\t\t\t</ChequeLine>\n\t\t\t<ChequeLine ChequeLineNo=\"2\" NoPayBonus=\"false\" NoAddBonus=\"false\" NoDiscounts=\"false\" Price=\"102.00\" Quantity=\"1.000\" Amount=\"102.00\" MinAmount=\"0\" MinPrice=\"0\" MaxDiscount=\"100.00\" BonusDiscount=\"0\">\n\t\t\t\t<Item ItemID=\"0\" ItemUID=\"81020900001\"/>\n\t\t\t\t<Discounts/>\n\t\t\t</ChequeLine>\n\t\t</ChequeLines>\n\t\t<Discounts/>\n\t\t<Payments/>\n\t\t<Messages/>\n\t</Cheque>\n")
	if err != nil {
		t.Errorf(err.Error())
	}

	// print result
	if len(dis.ChequeLines.ChequeLines) > 0 {
		fmt.Printf("\tGot messages: %d\n", len(dis.ChequeLines.ChequeLines))
		fmt.Printf("\tChequeLineNo: %d\n", dis.ChequeLines.ChequeLines[0].ChequeLineNo)
	} else {
		fmt.Printf("\tDon't got any discounts\n")
	}
}

func TestRSLoyaltyPOS5_GetMessages(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard(ValidCard)
	cheque.AddLine("49000123", 100.0, 2, nil)
	cheque.AddLine("1442", 200.0, 2, nil)
	cheque.AddLine("49000123", 300.0, 2, nil)

	msgs, err := pos.GetMessages(cheque.Xml())

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// print result
	if len(msgs.Messages) > 0 {
		fmt.Printf("\tGot messages: %d\n", len(msgs.Messages))
		fmt.Printf("\tMessageID: %s\n", msgs.Messages[0].MessageID)
	} else {
		fmt.Printf("\tDon't got any messages\n")
	}
}

func TestRSLoyaltyPOS5_GetCardDiscountAmountString(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("0000199000089")
	cheque.AddLine("49000123", 100.0, 2, nil)
	cheque.AddLine("1442", 200.0, 2, nil)
	cheque.AddLine("15633", 300.0, 2, nil)

	amountString, err := pos.GetCardDiscountAmountString("0000199000089", cheque.Xml())
	if err != nil {
		fmt.Printf("\tError: %v\n", err)
		return
	}
	fmt.Printf("\tAmount=%s\n", amountString)
}

func TestRSLoyaltyPOS5_GetCardDiscountAmount(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("0000199000089")
	cheque.AddLine("49000123", 100.0, 2, nil)
	cheque.AddLine("1442", 200.0, 2, nil)
	cheque.AddLine("15633", 300.0, 2, nil)

	//var tch = strings.Replace(cheque.Xml(), "\"", "&quot;", -1)
	amountString, err := pos.GetCardDiscountAmount("0000199000089", cheque.Xml())
	if err != nil {
		fmt.Printf("\tError: %v\n", err)
		return
	}
	fmt.Printf("\tAmount=%f\n", amountString)
}

func TestRSLoyaltyPOS5_SubtractBonus45(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("8641212605463261")
	cheque.AddLine("1442", 200.0, 2, nil)

	subs, err := pos.SubtractBonus45("8641212605463261", 2, cheque.Xml())
	if err != nil {
		fmt.Printf("\tError: %v\n", err)
		return
	}
	for _, line := range subs {
		fmt.Printf("\tLine number: %d, amount: %f\n", line.ChequeLineNo, line.Amount)
	}
}

func TestRSLoyaltyPOS5_CancelSubtractBonus(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("8641212605463261")
	cheque.AddLine("1442", 200.0, 2, nil)
	err := pos.CancelSubtractBonus("8641212605463261", 2, cheque.Xml())
	if err != nil {
		fmt.Printf("\tError: %v\n\r", err)
		return
	}
	fmt.Printf("\tSuccess!\n\r")
}

func TestRSLoyaltyPOS5_Accrual(t *testing.T) {
	cheque := Cheque.New("1", time.Now())
	cheque.AddDiscountCard("8641212605463261")
	cheque.AddLine("1442", 200.0, 2, nil)

	accrual, err := pos.Accrual(cheque.Xml())

	if err != nil {
		fmt.Printf("\tError: %v\n\r", err)
		return
	}

	fmt.Printf("\tAccrual: %s\n\r", *accrual)

}
