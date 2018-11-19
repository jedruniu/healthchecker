package healthchecker

import (
	"fmt"
	"time"
)

func Run(hc HealthChecker) {
	go func() {
		for {
			select {
			case <-time.After(hc.GetInterval()):
				fmt.Println(hc.IsHealthy())
			}
		}
	}()
}
