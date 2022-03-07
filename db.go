package monbundle

import (
	"fmt"
	"log"
	"monbundle/metrics"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance interface {
	UpdateMetric(m *metrics.Metric) error
	CleanUpMetrics()
}

type dbInstance struct {
	db *gorm.DB
}

func (dbInstance *dbInstance) UpdateMetric(m *metrics.Metric) error {
	// migrate
	dbInstance.db.AutoMigrate(&metrics.Metric{})

	// create
	dbInstance.db.Create(&m)

	return nil
}

func (dbInstance *dbInstance) CleanUpMetrics() {
	tsToDelete := time.Now().AddDate(0, 0, -1)
	fmt.Printf("ts: %v\n", int32(tsToDelete.Unix()))
	dbInstance.db.Debug().Unscoped().Where("ts < ?", int32(tsToDelete.Unix())).Delete(&metrics.Metric{})
}

func newDbInstance() DbInstance {
	db, err := gorm.Open(sqlite.Open("metric.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to init db. ", err)
		os.Exit(1)
	}
	return &dbInstance{db: db}
}

var dbInst = newDbInstance()

func DbInst() DbInstance { return dbInst }
