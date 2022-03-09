package main

import (
	"log"
	"monbundle"
	"os"
)

func main() {
	chartData := monbundle.RenderLAChart()
	err := os.WriteFile("la.png", chartData, 0755)
	if err != nil {
		log.Fatal("error writing chart")
		os.Exit(1)
	}
}

// var m []monbundle.CounterMetric
// var m5m []monbundle.CounterMetric

// db.Debug().Where("name = ?", "loadavg_5m").Find(&m5m)
// fmt.Println(m5m)

// db.Debug().Where("name = ?", "loadavg_1m").Find(&m)

// xVals := [100]time.Time{}
// yVals := [100]float64{}

// for i, v := range m {
// 	if i < 100 {
// 		fmt.Printf("i: %v, v: %v\n", i, v.Value)
// 		yVals[i] = float64(v.Value)
// 		xVals[i] = time.Unix(int64(v.TS), 0)

// 	}
// 	//		yVals = append(yVals, float64(v.Value))
// 	//		xVals = append(xVals, time.Unix(int64(v.TS), 0))
// }

// x5mVals := [100]time.Time{}
// y5mVals := [100]float64{}
// for i, v := range m5m {
// 	if i < 100 {
// 		y5mVals[i] = float64(v.Value)
// 		x5mVals[i] = time.Unix(int64(v.TS), 0)
// 	}
// }

// graph := chart.Chart{
// 	Title: "LA 1/5/15 minute",
// 	TitleStyle: chart.Style{
// 		FontSize: 8.0,
// 	},

// 	Width:  600,
// 	Height: 150,
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

// // graph := chart.Chart{

// // 	Series: []chart.Series{

// // 		chart.ContinuousSeries{
// // 			Style: chart.Style{

// // 				StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
// // 				FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
// // 			},
// // 			XValues: xVals,
// // 			YValues: yVals,
// // 		},
// // 	},
// // }

// buffer := bytes.NewBuffer([]byte{})
// err = graph.Render(chart.PNG, buffer)
// if err != nil {
// 	log.Fatal(err)
// 	os.Exit(1)
// }

// err = ioutil.WriteFile("test.png", buffer.Bytes(), 0644)
// if err != nil {
// 	log.Fatalln(err)
// 	os.Exit(1)
// }

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "image/png")
// 	_, err := w.Write(buffer.Bytes())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// })

// log.Fatal(http.ListenAndServe(":8000", nil))
