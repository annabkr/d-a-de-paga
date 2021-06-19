package model

// Schema of the transaction table
type Transaction struct {
	ID     int64      `json:"id"`
	Amount float64    `json:"float"`
	Source SourceType `json:"source"`
}

type SourceType string

const (
	PayCheck SourceType = "paycheck"
	Stocks   SourceType = "stocks"
)
