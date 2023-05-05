package rsloyaltypos5

import (
	"fmt"
	Cheque "github.com/ramzes4rules/rs.loyalty.pos.5/cheque"
	"testing"
)

var pos = RSLoyaltyPOS5{
	url:      "http://afanasiy-test.retailloyalty.ru:8383/RSLoyaltyStoreService",
	login:    "rsl_reserv",
	password: "rsl_reserv",
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

func TestRSLoyaltyPOS5_GetCardBalance(t *testing.T) {
	balance, err := pos.GetCardBalance(ValidCard)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Printf("\tBalance: %5.2f\n", balance.Balance.Value)
}

func TestRSLoyaltyPOS5_GetDiscounts(t *testing.T) {
	cheque := Cheque.Cheque{}
	err := pos.GetDiscounts(cheque.Test())
	if err != nil {
		t.Errorf(err.Error())
	}
}
