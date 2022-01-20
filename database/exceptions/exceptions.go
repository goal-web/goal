package exceptions

import "github.com/goal-web/contracts"

type TransactionException struct {
	contracts.Exception
}

type RollbackException struct {
	contracts.Exception
}

type BeginException struct {
	contracts.Exception
}
