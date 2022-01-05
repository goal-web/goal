package database

import "github.com/qbhy/goal/contracts"

type TransactionException struct {
	contracts.Exception
}

type RollbackException struct {
	contracts.Exception
}

type BeginException struct {
	contracts.Exception
}
