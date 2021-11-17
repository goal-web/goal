package signal

import (
	"github.com/qbhy/goal/contracts"
	"os"
)

var (
	SIGNAL_RECEIVED = contracts.EventName("SIGNAL_RECEIVED")
)

type SignalReceived struct {
	Signal os.Signal
}

func (this *SignalReceived) Name() contracts.EventName {
	return SIGNAL_RECEIVED
}
