package monbundle

import (
	"log"
	"monbundle/metrics"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func UpdateMetric(m *metrics.Metric) {
	db, err := gorm.Open(sqlite.Open("metric.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	// migrate
	db.AutoMigrate(&metrics.Metric{})

	// create
	db.Create(&m)

}
