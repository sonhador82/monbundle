package main

import (
	"fmt"
	"log"
	"monbundle"
	"os"
	"sync"
	"time"
)

const scanFreq = time.Second * 5

func runNetStat() {
	netDev := os.Getenv("NETSTAT_DEV")
	for {
		netstat := monbundle.LoadNetDevStat(netDev)
		fmt.Println(netstat)
		monbundle.DbInst().UpdateCounterMetrics(netstat)
		time.Sleep(scanFreq)
	}
}

func runLA() {
	log.Println("Start LA scraping...")
	for {
		metrics := monbundle.LoadAvg()
		monbundle.DbInst().UpdateFloatMetrics(metrics)
		time.Sleep(scanFreq)
	}
}

func diskStats() {
	log.Println("Start DisksStats scraping...")
	for {
		metrics := monbundle.LoadDiskStats()
		monbundle.DbInst().UpdateCounterMetrics(metrics)
		time.Sleep(scanFreq)
	}
}

func main() {

	wg := sync.WaitGroup{}

	scrapers := []interface{}{runLA, diskStats}
	wg.Add(len(scrapers))
	for _, fn := range scrapers {
		go fn.(func())()
	}
	wg.Add(1)
	go monbundle.Serve()
	wg.Wait()
}
