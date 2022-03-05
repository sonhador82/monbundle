package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"monbundle/metrics"
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("metric.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	var m []metrics.Metric

	db.Find(&m)

	var xVals = make([]float64, 5)
	var yVals = make([]float64, 5)

	for i, v := range m {
		fmt.Printf("i: %v, v: %v\n", i, v.Value)
		yVals = append(yVals, float64(v.Value))
		xVals = append(xVals, float64(i))
	}
	fmt.Println(m[0].Value)
	fmt.Println(len(m))
	fmt.Println(yVals)

	graph := chart.Chart{

		Series: []chart.Series{

			chart.ContinuousSeries{
				Style: chart.Style{

					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: xVals,
				YValues: yVals,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("test.png", buffer.Bytes(), 0644)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
