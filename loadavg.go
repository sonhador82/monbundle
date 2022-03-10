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

func getSeries(name string) ([]time.Time, []float64) {
	dbConn := DbInst().GetDB()
	var m []FloatMetric
	dbConn.Debug().Where("name = ?", name).Find(&m)
	dataTS := make([]time.Time, 0)
	dataVal := make([]float64, 0)
	for _, v := range m {
		dataTS = append(dataTS, v.TS)
		dataVal = append(dataVal, v.Value)
	}
	return dataTS, dataVal
}

func RenderLAChart() []byte {
	data1mTS, data1mVal := getSeries("loadavg_1m")
	data5mTS, data5mVal := getSeries("loadavg_5m")
	data15mTS, data15mVal := getSeries("loadavg_15m")

	graph := chart.Chart{
		Height: 200,
		Width:  480,
		Title:  "LA 1m/5m/15m",
		TitleStyle: chart.Style{
			FontSize: 8.0,
		},
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeMinuteValueFormatter,
			Style: chart.Style{
				FontSize: 6.0,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name:    "1m",
				XValues: data1mTS,
				YValues: data1mVal,
			},
			chart.TimeSeries{
				Name:    "5m",
				XValues: data5mTS,
				YValues: data5mVal,
			},
			chart.TimeSeries{
				Name:    "15m",
				XValues: data15mTS,
				YValues: data15mVal,
			},
		},
	}

	// add legend
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return buffer.Bytes()
}
