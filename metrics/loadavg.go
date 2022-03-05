package metrics

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

const loadAvgFile = "/proc/loadavg"

type Metric struct {
	gorm.Model
	Name  string
	Value float32
	TS    int32
}

func LoadAVG_1m() *Metric {

	content, err := ioutil.ReadFile(loadAvgFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	s := string(content)
	items := strings.Split(s, " ")
	la1m, _ := strconv.ParseFloat(items[0], 32)
	return &Metric{
		Name:  "loadavg_1m",
		Value: float32(la1m),
		TS:    int32(time.Now().Unix()),
	}
}
