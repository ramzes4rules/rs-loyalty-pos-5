package Cheque

import (
	"math/rand"
)

type UidType string

const (
	GUID UidType = "GUID"
)

func NewUID(UidType UidType) string {

	const set1 = "0123456789"
	const set2 = "ABCDEF"
	const set3 = "abcdef"

	switch UidType {
	case GUID:
		var set = set1 + set2
		var uid = ""

		//rand.Seed(time.Now().UnixNano())
		//06B71CCC-FDB7-46AE-AE6C-5F49380F1CFA

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
	default:
		return ""
	}
}
