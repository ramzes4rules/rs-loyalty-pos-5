package Cheque

import (
	"math/rand"
)

type UidType string

const (
	GUID   UidType = "GUID"
	Number UidType = "Number"
)

func NewUID(UidType UidType) string {

	const set1 = "0123456789"
	const set2 = "ABCDEF"
	const set3 = "abcdef"

	switch UidType {
	case GUID:
		var set = set1 + set3
		var uid = ""
		for i := 1; i <= 8; i++ {
			uid += string(set[rand.Intn(len(set)-1)])
		}
		uid += "-"
		for i := 1; i <= 4; i++ {
			uid += string(set[rand.Intn(len(set)-1)])
		}
		uid += "-"
		for i := 1; i <= 4; i++ {
			uid += string(set[rand.Intn(len(set)-1)])
		}
		uid += "-"
		for i := 1; i <= 4; i++ {
			uid += string(set[rand.Intn(len(set)-1)])
		}
		uid += "-"
		for i := 1; i <= 12; i++ {
			uid += string(set[rand.Intn(len(set)-1)])
		}
		return uid

	case Number:
		var number = ""
		for i := 1; i <= 5; i++ {
			number += string(set1[rand.Intn(len(set1)-1)])
		}
		return number

	default:
		return ""
	}
}
