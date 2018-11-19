package main

import (
	"fmt"
	"time"

	"github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	hc := healthchecker.NewFileBasedHealthChecker("testFile.txt", 5*time.Second)

	for {
		fmt.Println(hc.isHealthy())
		time.Sleep(hc.getInterval())
	}
}
