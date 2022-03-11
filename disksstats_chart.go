package monbundle

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

type namesResult struct {
	Name string
}

/// удалить?
func getMetricsDevs() []string {
	var result []namesResult
	db := dbInst.GetDB()
	db.Raw("SELECT DISTINCT name FROM counter_metrics").Scan(&result)

	names := make([]string, 0)
	for _, item := range result {
		if !strings.Contains(item.Name, "dm-") {
			names = append(names, item.Name)
		}
	}
	return names
}

type devQueryResult struct {
	ReadBytes    []CounterMetric
	WrittenBytes []CounterMetric
}

func getDiskMetrics(devName string) *devQueryResult {
	db := dbInst.GetDB()
	var resultRead []CounterMetric
	var resultWritten []CounterMetric
	filter := fmt.Sprintf("disk_%s_read_bytes", devName)
	db.Debug().Where("name = ?", filter).Order("ts").Find(&resultRead)
	filter = fmt.Sprintf("disk_%s_written_bytes", devName)
	db.Debug().Where("name = ?", filter).Order("ts").Find(&resultWritten)
	return &devQueryResult{
		ReadBytes:    resultRead,
		WrittenBytes: resultWritten,
	}

}

type ChartData struct {
	Values []float64
	TS     []time.Time
}

func prepSeriesForChart(m []CounterMetric) *ChartData {
	chartDataTs := make([]time.Time, 0)
	chartDataVal := make([]float64, 0)

	for index := range m[:len(m)-1] {
		deltaTS := m[index+1].TS.Sub(m[index].TS)
		deltaVal := m[index+1].Value - m[index].Value
		var avgValPerSec float64 = float64(deltaVal / uint64(deltaTS.Seconds()))
		chartDataTs = append(chartDataTs, m[index].TS)
		chartDataVal = append(chartDataVal, avgValPerSec)
	}
	return &ChartData{
		Values: chartDataVal,
		TS:     chartDataTs,
	}
}

// disk_dm-1_read_bytes
func RenderDiskChart(devName string) []byte {
	devMetrics := getDiskMetrics(devName)
	readChartData := prepSeriesForChart(devMetrics.ReadBytes)
	writtenChartData := prepSeriesForChart(devMetrics.WrittenBytes)

	graph := chart.Chart{
		Width:  480,
		Height: 200,
		Title:  fmt.Sprintf("%s readbytes/sec", devName),
		TitleStyle: chart.Style{
			FontSize: 8.0,
		},
		XAxis: chart.XAxis{

			ValueFormatter: chart.TimeMinuteValueFormatter,
			Style: chart.Style{
				FontSize: 8.0,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name:    "readBytes/sec",
				XValues: readChartData.TS,
				YValues: readChartData.Values,
			},
			chart.TimeSeries{
				Name:    "writtenBytes/sec",
				XValues: writtenChartData.TS,
				YValues: writtenChartData.Values,
			},
		},
	}
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
