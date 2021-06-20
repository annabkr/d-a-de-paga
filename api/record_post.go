package api

import (
	"encoding/json"
	"net/http"
	"pkg/errors"

	"github.com/google/uuid"

	db "github.com/annabkr/paydayz/database"
	"github.com/annabkr/paydayz/model"
	e "github.com/annabkr/paydayz/utils/errors"
)

type TransactionRep struct {
	Amount float64 `json:"amount"`
	Source string `json:"source"`
}

func postRecord(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var transactionRep TransactionRep
	err := decoder.Decode(&transactionRep)
	if err != nil {
		return errors.Wrap(err, "unable to decode request")
	}

	source := model.GetSourceType(transactionRep.Source)
	if !model.IsValid(source){
		return e.NewBadRequestError("invalid source type", nil)
	}

	transaction := model.Transaction{
		ID: uuid.New(),
		Amount: transactionRep.Amount,
		Source: source,
	}

	err = db.Insert(transaction)
	if err != nil {
		return err
	}
	return nil
}

// type Transaction struct {
// 	ID     int64      `json:"id"`
// 	Amount float64    `json:"float"`
// 	Source SourceType `json:"source"`
// }


// CREATE TABLE transactions(
//     id SERIAL PRIMARY KEY,
//     amount FLOAT NOT NULL,
//     source TEXT NOT NULL
// )