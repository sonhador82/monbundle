package metrics

import (
	"time"

	"gorm.io/gorm"
)

type CounterMetric struct {
	ID    uint64
	Name  string
	Value uint64
	TS    time.Time
}

type Metric struct {
	gorm.Model
	Name  string
	Value float32
	TS    int32
}
