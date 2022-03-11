package monbundle

import (
	"fmt"
	"strings"
)

type namesResult struct {
	Name string
}

func GetMetricsDevs() {
	var result []namesResult
	db := dbInst.GetDB()
	db.Raw("SELECT DISTINCT name FROM counter_metrics").Scan(&result)

	names := make([]string, 0)
	for _, item := range result {
		if !strings.Contains(item.Name, "dm-") {
			names = append(names, item.Name)
		}
	}
	fmt.Println(names)

}

// func getSeries() {
// 	const name = "disk_sdc_read_bytes"
// 	dbConn := monbundle.DbInst().GetDB()
// 	var m []monbundle.CounterMetric
// 	dbConn.Debug().Where("name = ?", name).Order("ts").Find(&m)
// 	fmt.Println(m[0])
// 	fmt.Println(m[len(m)-1])
// 	// делаем дельту? берем текущий + следующий и делим дельту на кол-во секунд в промежутке?
// 	fmt.Printf("len %v\n", len(m))

// 	chartDataTs := make([]time.Time, 0)
// 	chartDataVal := make([]float64, 0)

// 	for index := range m[:len(m)-1] {
// 		deltaTS := m[index+1].TS.Sub(m[index].TS)
// 		deltaVal := m[index+1].Value - m[index].Value
// 		var avgValPerSec float64 = float64(deltaVal / uint64(deltaTS.Seconds()))
// 		fmt.Printf("DeltaTS: %v, DeltaVal: %v\n", deltaTS.Seconds(), deltaVal)
// 		fmt.Printf("avgPerSec: %v\n", avgValPerSec)
// 		chartDataTs = append(chartDataTs, m[index].TS)
// 		chartDataVal = append(chartDataVal, avgValPerSec)
// 	}

// 	graph := chart.Chart{
// 		Height: 200,
// 		Width:  480,
// 		Title:  "Disk readbytes/sec",
// 		XAxis: chart.XAxis{
// 			ValueFormatter: chart.TimeMinuteValueFormatter,
// 			Style: chart.Style{
// 				FontSize: 6.0,
// 			},
// 		},
// 		Series: []chart.Series{
// 			chart.TimeSeries{
// 				Name:    "readbytes/sec",
// 				XValues: chartDataTs,
// 				YValues: chartDataVal,
// 			},
// 		},
// 	}

// 	buffer := bytes.NewBuffer([]byte{})
// 	err := graph.Render(chart.PNG, buffer)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}

// 	os.WriteFile("test.png", buffer.Bytes(), 0644)
// 	//return buffer.Bytes()

// }
