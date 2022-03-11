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

var DiskDev = os.Getenv("DISK_DEV")
var NetDev = os.Getenv("NETSTAT_DEV")

func runNetStat() {
	for {
		netstat := monbundle.LoadNetDevStat(NetDev)
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
	if DiskDev == "" {
		log.Fatal("Export DISK_DEV env variable for disk metrics, ex: DISK_DEV=sda")
		os.Exit(1)
	}

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
