package monbundle

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const netStatFile = "/proc/net/dev"

func LoadNetDevStat(iface string) []*CounterMetric {
	data, err := os.ReadFile(netStatFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var recBytes, transBytes int64

	infoLines := strings.Split(string(data), "\n")
	for _, line := range infoLines {
		devInfo := strings.Fields(line)
		if len(devInfo) > 0 && strings.Contains(devInfo[0], iface) {
			fmt.Println(iface, devInfo)
			parsed, err := strconv.Atoi(devInfo[1])
			if err != nil {
				parsed = 0
			}
			recBytes = int64(parsed)
			parsed, err = strconv.Atoi(devInfo[9])
			if err != nil {
				parsed = 0
			}
			transBytes = int64(parsed)
		}
	}
	return []*CounterMetric{
		{
			Name:  fmt.Sprintf("netstat_%s_recv_bytes", iface),
			Value: uint64(recBytes),
			TS:    time.Now(),
		},
		{
			Name:  fmt.Sprintf("netstat_%s_trans_bytes", iface),
			Value: uint64(transBytes),
			TS:    time.Now(),
		},
	}
}
