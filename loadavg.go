package monbundle

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

const loadAvgFile = "/proc/loadavg"

func loadContent() []byte {
	content, err := ioutil.ReadFile(loadAvgFile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return content
}

func LoadAvg() []*FloatMetric {
	rawData := string(loadContent())
	items := strings.Split(rawData, " ")
	la1m, _ := strconv.ParseFloat(items[0], 64)
	la5m, _ := strconv.ParseFloat(items[1], 64)
	la15m, _ := strconv.ParseFloat(items[2], 64)
	return []*FloatMetric{
		{
			Name:  "loadavg_1m",
			Value: la1m,
			TS:    time.Now(),
		},
		{
			Name:  "loadavg_5m",
			Value: la5m,
			TS:    time.Now(),
		},
		{
			Name:  "loadavg_15m",
			Value: la15m,
			TS:    time.Now(),
		},
	}
}

func getSeries(name string)

func RenderLAChart() []byte {
	db_conn := DbInst().GetDB()
	var m1m []FloatMetric
	var m5m []db.FloatMetric
	var m15m []db.FloatMetric

	db_conn.Debug().Where("name = ?", "loadavg_1m").Find(&m1m)
	data1mTS := make([]time.Time, 0)
	data1mVal := make([]float64, 0)
	for _, v := range m1m {
		data1mTS = append(data1mTS, v.TS)
	}
	for _, v := range m1m {
		data1mVal = append(data1mVal, v.Value)
	}

	graph := chart.Chart{
		Title: "LA 1m/5m/15m",
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: data1mTS,
				YValues: data1mVal,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return buffer.Bytes()
}

// graph := chart.Chart{
// 	Title: "LA 1/5/15 minutes",
// 	XAxis: chart.XAxis{
// 		ValueFormatter: chart.TimeMinuteValueFormatter,
// 	},
// 	Series: []chart.Series{
// 		chart.TimeSeries{
// 			XValues: xVals[:100],
// 			YValues: yVals[:100],
// 		},
// 		chart.TimeSeries{
// 			XValues: x5mVals[:100],
// 			YValues: y5mVals[:100],
// 		},
// 	},
// }
