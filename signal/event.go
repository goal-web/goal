package signal

import (
	"os"
)

type Received struct {
	Signal os.Signal
}

func (this *Received) Event() string {
	return "SIGNAL_RECEIVED"
}

func (this *Received) IsSync() bool {
	return true
}
