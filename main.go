package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	hc := healthchecker.NewFileBasedHealthChecker("testFile.txt", 1*time.Second)
	apiHc := healthchecker.NewApiCallBasedHealthChecker("http://google.com", 5*time.Second)

	healthchecker.Schedule(hc)
	healthchecker.Schedule(apiHc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}
