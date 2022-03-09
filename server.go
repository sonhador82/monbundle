package monbundle

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
)

func prepGlobChart() []byte {
	chartData := RenderLAChart()
	buffer := bytes.NewBuffer(chartData)

	img, format, err := image.Decode(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("format: %s", format)

	bigImg := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	draw.Draw(bigImg, image.Rect(0, 0, 480, 200), img, image.Point{0, 0}, draw.Over)

	var bbuf []byte
	buf := bytes.NewBuffer(bbuf)
	err = png.Encode(buf, bigImg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return buf.Bytes()
}

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		chartBuf := prepGlobChart()

		w.Header().Add("Content-Type", "image/png")
		_, err := w.Write(chartBuf)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
