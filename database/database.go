package database

import (
	"context"
	"fmt"
	"log"
	"pkg/errors"
	"time"

	"github.com/annabkr/paydayz/app"
	"github.com/annabkr/paydayz/model"
	"github.com/jackc/pgx/v4"
)


func BeginTransaction() (pgx.Tx, error) {
	ctx := context.Background()
	tx, err := app.GetPool().Begin(ctx)
	if err != nil {
		return tx, err
	}
	defer func(){
		err := tx.Rollback(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return tx, nil
}

func Insert(t model.Transaction) error {
	sql := getInsertQuery(t)
	tx, err := BeginTransaction()
	if err != nil {
		return err 
	}
	tag, err := tx.Exec(context.Background(), sql, time.Now())
	if !tag.Insert(){
		msg := "Insert failed"
		if err != nil {
			return errors.Wrap(err, msg)
		}
		err = errors.New(msg)
	}
	return err
}

func getInsertQuery(t model.Transaction) string {
	return fmt.Sprintf(`INSERT INTO transactions (id, amount, source)
	VALUES(%d, %f, '%s')`, t.ID, t.Amount, t.Source)
}
