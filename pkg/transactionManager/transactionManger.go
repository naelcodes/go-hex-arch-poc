package transactionManager

import (
	"context"
	"errors"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type TransactionManager struct {
	transaction *ent.Tx
	database    *ent.Client
	context     context.Context
}

func NewTransactionManager(context context.Context, client *ent.Client) *TransactionManager {
	return &TransactionManager{context: context, database: client}
}

func (tm *TransactionManager) Begin() error {
	transaction, err := tm.database.Tx(tm.context)
	if err != nil {
		return CustomErrors.ServiceError(err, "Starting Transaction")
	}
	tm.transaction = transaction
	return nil
}

func (tm *TransactionManager) Commit() error {
	return tm.transaction.Commit()
}

func (tm *TransactionManager) Rollback() error {
	if tm.transaction == nil {
		return nil
	}
	return tm.transaction.Rollback()
}

func (tm *TransactionManager) GetTransaction() *ent.Tx {
	return tm.transaction
}

func (tm *TransactionManager) CatchError() error {
	utils.Logger.Info(fmt.Sprintf("[TransactionManager - CatchError] Transaction: %v", tm.transaction))
	if err := recover(); err != nil {
		rollBackErr := tm.Rollback()
		err, _ := err.(error)
		if rollBackErr != nil {
			return CustomErrors.UnknownError(errors.Join(rollBackErr, err))
		}
		return err
	}
	return nil
}
