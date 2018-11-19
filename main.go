package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	hc := healthchecker.NewFileBasedHealthCheck("file based one", "testFile.txt", 1*time.Second)
	apiHc := healthchecker.NewApiCallBasedHealthCheck("hit google", "http://google.com", 5*time.Second)
	redisHc := healthchecker.NewRedisBasedHealthCheck("get from redis", "some_key", 2*time.Second)

	healthchecker.Run(hc)
	healthchecker.Run(apiHc)
	healthchecker.Run(redisHc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
