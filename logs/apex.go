package logs

import (
	"github.com/apex/log"
	"github.com/goal-web/contracts"
)

type ApexLogger struct {
	Entry *log.Entry
}

func (this *ApexLogger) WithFields(m contracts.Fields) contracts.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithFields(log.Fields(m)),
		}
	}

	this.Entry = this.Entry.WithFields(log.Fields(m))

	return this
}

func (this *ApexLogger) WithField(key string, value interface{}) contracts.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithField(key, value),
		}
	}

	this.Entry = this.Entry.WithField(key, value)

	return this
}

func (this *ApexLogger) WithError(err error) contracts.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithError(err),
		}
	}

	this.Entry = this.Entry.WithError(err)

	return this
}

func (this *ApexLogger) WithException(err contracts.Exception) contracts.Logger {
	if this == nil || this.Entry == nil {
		this = &ApexLogger{
			Entry: log.WithError(err).WithFields(log.Fields(err.Fields())),
		}
	}

	this.Entry = this.Entry.WithError(err).WithFields(log.Fields(err.Fields()))

	return this
}

func (this ApexLogger) Info(msg string) {
	this.Entry.Info(msg)
}

func (this ApexLogger) Warn(msg string) {
	this.Entry.Warn(msg)
}

func (this ApexLogger) Debug(msg string) {
	this.Entry.Debug(msg)
}

func (this ApexLogger) Error(msg string) {
	this.Entry.Error(msg)
}

func (this ApexLogger) Fatal(msg string) {
	this.Entry.Fatal(msg)
}
