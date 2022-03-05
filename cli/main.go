package main

import (
	"fmt"
	"monbundle"
	"monbundle/metrics"
	"sync"
	"time"
)

const scanFreq = time.Second * 5

func runLA() {
	for {
		la1m := metrics.LoadAVG_1m()
		fmt.Printf("%v\n", la1m)
		monbundle.UpdateMetric(la1m)
		time.Sleep(scanFreq)
	}
}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go runLA()
	wg.Wait()
}
