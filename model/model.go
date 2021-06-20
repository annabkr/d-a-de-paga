package model

import "github.com/google/uuid"

// Schema of the transaction table
type Transaction struct {
	ID     uuid.UUID      `json:"id"`
	Amount float64    `json:"amount"`
	Source SourceType `json:"source"`
}

type SourceType string

func (s SourceType) String() string{
	return string(s)
}

const (
	Paycheck SourceType = "paycheck"
	Stocks   SourceType = "stocks"
	Shopping SourceType = "shopping"
	Groceries SourceType = "groceries"
	Clothes SourceType = "clothes"
	Invalid SourceType = "invalid"
)

var ValidSourceTypes = []SourceType{Paycheck, Stocks, Shopping, Groceries, Clothes}

func GetSourceType(s string) SourceType {
	for _, source := range ValidSourceTypes{
		if source.String() == s {
			return source
		}
	}
	return Invalid
}

func IsValid(s SourceType) bool{
	return s != Invalid
}
