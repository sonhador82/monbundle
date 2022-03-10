package monbundle

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const statFile = "/proc/diskstats"
const blockDevsDir = "/sys/block"

type BlockDev struct {
	Name       string
	SectorSize uint
}

func getSectorSize(devName string) int {
	raw, err := os.ReadFile(fmt.Sprintf("/sys/block/%s/queue/hw_sector_size", devName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	secSize, err := strconv.Atoi(strings.TrimSpace(string(raw)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return secSize
}

func findBlockDevs() []BlockDev {
	devs, err := os.ReadDir(blockDevsDir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	prepedDevs := make([]BlockDev, 0)
	for _, blockDev := range devs {
		secSize := getSectorSize(blockDev.Name())
		prepedDevs = append(prepedDevs, BlockDev{
			Name:       blockDev.Name(),
			SectorSize: uint(secSize),
		})
	}
	return prepedDevs
}

func LoadDiskStats() []*CounterMetric {
	blockDevs := findBlockDevs()

	rawData, err := os.ReadFile(statFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	metrics := make([]*CounterMetric, 0)
	dataLn := strings.Split(string(rawData), "\n")
	for _, ln := range dataLn {
		flds := strings.Fields(ln)
		if len(flds) > 0 {
			dskName := flds[2]
			dskSectRead, _ := strconv.ParseUint(flds[5], 10, 64)
			dskSectWrite, _ := strconv.ParseUint(flds[9], 10, 64)
			for _, dev := range blockDevs {
				if dev.Name == dskName {
					metrics = append(metrics, &CounterMetric{
						Name:  fmt.Sprintf("disk_%s_read_bytes", dev.Name),
						Value: dskSectRead * uint64(dev.SectorSize),
						TS:    time.Now(),
					})
					metrics = append(metrics, &CounterMetric{
						Name:  fmt.Sprintf("disk_%s_written_bytes", dev.Name),
						Value: dskSectWrite * uint64(dev.SectorSize),
						TS:    time.Now(),
					})
				}
			}
		}

	}
	return metrics
}
