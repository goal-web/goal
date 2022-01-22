package database

import (
	"github.com/goal-web/contracts"
)

type ConnectionErrorCode int

const (
	DbDriverDontExist ConnectionErrorCode = iota
	DbConnectionDontExist
)

type DBConnectionException struct {
	error
	Connection string
	Code       ConnectionErrorCode
	fields     contracts.Fields
}

func (this DBConnectionException) Fields() contracts.Fields {
	this.fields["Code"] = this.Code
	return this.fields
}
