package Cheque

type Status string

const (
	StatusClosed    Status = "Closed "
	StatusOpen      Status = "Open"
	StatusCancelled Status = "Cancelled"
)
