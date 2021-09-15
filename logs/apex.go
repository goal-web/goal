package logs

import (
	"github.com/apex/log"
	"qbhy/contracts"
)

type ApexLogger struct {
	Entry *log.Entry
}

func (a *ApexLogger) WithFields(m contracts.Fields) contracts.Logger {
	if a == nil || a.Entry == nil {
		a = &ApexLogger{
			Entry: log.WithFields(log.Fields(m)),
		}
	}

	a.Entry.WithFields(log.Fields(m))

	return a
}

func (a *ApexLogger) WithError(err error) contracts.Logger {
	if a == nil || a.Entry == nil {
		a = &ApexLogger{
			Entry: log.WithError(err),
		}
	}

	a.Entry.WithError(err)

	return a
}

func (a *ApexLogger) WithException(err contracts.Exception) contracts.Logger {
	if a == nil || a.Entry == nil {
		a = &ApexLogger{
			Entry: log.WithError(err).WithFields(log.Fields(err.Fields())),
		}
	}

	a.Entry.WithError(err).WithFields(log.Fields(err.Fields()))

	return a
}

func (a ApexLogger) Info(msg string) {
	a.Entry.Info(msg)
}

func (a ApexLogger) Warn(msg string) {
	a.Entry.Warn(msg)
}

func (a ApexLogger) Debug(msg string) {
	a.Entry.Debug(msg)
}

func (a ApexLogger) Error(msg string) {
	a.Entry.Error(msg)
}

func (a ApexLogger) Fatal(msg string) {
	a.Entry.Fatal(msg)
}
