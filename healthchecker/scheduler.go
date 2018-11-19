package healthchecker

import (
	"fmt"
	"time"
)

func Schedule(hc HealthChecker) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(hc.GetInterval()):
				fmt.Println(hc.IsHealthy())
			case <-stop:
				break
			}
		}
	}()
	return stop
}
