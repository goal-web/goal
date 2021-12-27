package signal

import (
	"os"
)

type SignalReceived struct {
	Signal os.Signal
}

func (this *SignalReceived) Event() string {
	return "SIGNAL_RECEIVED"
}
