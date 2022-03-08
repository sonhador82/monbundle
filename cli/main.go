package main

import (
	"fmt"
	"monbundle"
	"monbundle/metrics"
	"sync"
	"time"
)

const scanFreq = time.Second * 5

func runNetStat() {
	for {
		netstat := metrics.LoadNetDevStat("wlp2s0")
		fmt.Println(netstat)
		monbundle.DbInst().UpdateCounterMetrics(netstat)
		time.Sleep(scanFreq)
	}
}

func runLA() {
	for {
		la1m := metrics.LoadAVG_1m()
		la5m := metrics.LoadAvg_5m()
		fmt.Printf("%v\n", la1m)
		fmt.Printf("%v\n", la5m)
		monbundle.DbInst().UpdateMetric(la1m)
		monbundle.DbInst().UpdateMetric(la5m)
		time.Sleep(scanFreq)
	}
}

func dataGc() {
	for {
		fmt.Println("Cleanup db")
		monbundle.DbInst().CleanUpMetrics()
		fmt.Println("After cleanup")
		time.Sleep(time.Hour)
	}
}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(3)
	go runLA()
	go runNetStat()
	go dataGc()
	wg.Wait()
}
