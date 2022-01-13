package utils

import "time"

func SetInterval(second int, callback func(), onClose func()) chan bool {
	closeChan := make(chan bool, 1)
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(second))
		for {
			select {
			case <-ticker.C:
				callback()
			case <-closeChan:
				onClose()
				return
			}
		}
	}()
	return closeChan
}
