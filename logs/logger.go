package logs

import (
	"github.com/apex/log"
	"goal/contracts"
)

func WithFields(fields contracts.Fields) contracts.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(fields))}
}

func WithError(err error) contracts.Logger {
	return &ApexLogger{Entry: log.WithError(err)}
}

func WithException(exception contracts.Exception) contracts.Logger {
	return &ApexLogger{Entry: log.WithError(exception).WithFields(log.Fields(exception.Fields()))}
}

func Default() contracts.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(contracts.Fields{}))}
}
