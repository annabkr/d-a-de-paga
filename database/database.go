package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/annabkr/paydayz/app"
	"github.com/annabkr/paydayz/model"
	l "github.com/annabkr/paydayz/utils/logger"
)

func BeginTransaction() (pgx.Tx, error) {
	l.Info("Beginning INSERT transaction")  
	l.Info(fmt.Sprintf("Context: %+v", context.Background()))
	tx, err := app.GetPool().Begin(context.Background())
	if err != nil {
		return tx, errors.Wrap(err, "problem beginning transaction")
	}
	defer func(){
		err := tx.Rollback(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		l.Info("Rolling back transaction")
	}()

	return tx, nil
}

func Insert(t model.Transaction) error { 
	sql := getInsertQuery(t)
	l.Info(fmt.Sprintf("Executing query: %s", sql))
	tx, err := BeginTransaction()
	if err != nil {
		l.Info("Problem starting INSERT transaction")
		return err 
	}
	l.Info("INSERT transaction has begun and will now execute")
	
	tag, err := tx.Exec(context.Background(), sql)
	l.Info("INSERT transaction execution complete")
	if !tag.Insert(){
		msg := "Insert failed"
		if err != nil {
			return errors.Wrap(err, msg)
		}
		err = errors.New(msg)
	}
	l.Info(fmt.Sprintf("Result of transaction: \n%+v \n %+v\n", tag, err))
	return err
}

func getInsertQuery(t model.Transaction) string { 
	return fmt.Sprintf(`INSERT INTO transactions (id, amount, source) VALUES('%s', %.2f, '%s')`, t.ID.String(), t.Amount, t.Source.String())
}
