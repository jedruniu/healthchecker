package main

import (
	"os"
	"os/signal"
	"time"

	h "github.com/jedruniu/healthchecker/healthchecker"
)

func main() {

	fileCheck := h.NewHealthCheck(
		"file based one",
		10,
		3,
		1*time.Second,
		h.NewFileBased("testFile.txt"),
	)

	apiCheck := h.NewHealthCheck(
		"google endpoint",
		10,
		3,
		10*time.Second,
		h.NewApiCallBased("http://google.com"),
	)

	redisCheck := h.NewHealthCheck(
		"get some key from redis",
		1,
		1,
		30*time.Second,
		h.NewRedisBased("some_key"),
	)

	h.Run(fileCheck)
	h.Run(apiCheck)
	h.Run(redisCheck)

	// TODO implement server to fetch data
	// TODO implement bash script checks

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
