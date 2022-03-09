package monbundle

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance interface {
	UpdateFloatMetrics(m []*FloatMetric)
	UpdateCounterMetrics(m []*CounterMetric)
	GetDB() *gorm.DB
}

type dbInstance struct {
	db *gorm.DB
}

func (dbInstance *dbInstance) GetDB() *gorm.DB {
	return dbInstance.db
}

func (dbInstance *dbInstance) UpdateFloatMetrics(m []*FloatMetric) {
	// migrate
	dbInstance.db.AutoMigrate(&FloatMetric{})

	dbInstance.db.Create(&m)
}

func (dbInstance *dbInstance) UpdateCounterMetrics(m []*CounterMetric) {
	// migrate
	dbInstance.db.AutoMigrate(&CounterMetric{})

	dbInstance.db.Create(&m)
}

// func (dbInstance *dbInstance) CleanUpMetrics() {
// 	tsToDelete := time.Now().AddDate(0, 0, -1)
// 	fmt.Printf("ts: %v\n", int32(tsToDelete.Unix()))
// 	dbInstance.db.Debug().Unscoped().Where("ts < ?", int32(tsToDelete.Unix())).Delete(&metrics.Metric{})
// }

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
