package main

import (
	"fmt"
	"time"

	"github.com/jedruniu/healthchecker/healthchecker"
)

func main() {
	hc := healthchecker.NewFileBasedHealthChecker("testFile.txt", 5*time.Second)
	apiHc := healthchecker.NewApiCallBasedHealthChecker("http://google.com", 5*time.Second)

	for {
		fmt.Println("file based: ", hc.IsHealthy())
		fmt.Println("endpoint based: ", apiHc.IsHealthy())
		time.Sleep(hc.GetInterval())
	}
}
