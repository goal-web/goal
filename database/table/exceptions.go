package table

import "github.com/goal-web/contracts"

type CreateException struct {
	contracts.Exception
}

type InsertException struct {
	contracts.Exception
}

type UpdateException struct {
	contracts.Exception
}

type DeleteException struct {
	contracts.Exception
}

type SelectException struct {
	contracts.Exception
}

type NotFoundException struct {
	contracts.Exception
}
