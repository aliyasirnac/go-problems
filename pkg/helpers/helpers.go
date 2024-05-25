package helpers

import "time"

func SetInterval(f func(), second int64) {
	ticker := time.NewTicker(time.Duration(second) * time.Second)
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
