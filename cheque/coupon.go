package Cheque

type Coupon struct {
	CouponNo string `xml:"CouponNo,attr"` // Номер купона. Должен быть заведен в системе RS.Loyalty.
}
