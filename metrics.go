package monbundle

import (
	"time"
)

type CounterMetric struct {
	ID    uint64
	Name  string
	Value uint64
	TS    time.Time
}

type FloatMetric struct {
	ID    uint64
	Name  string
	Value float64
	TS    time.Time
}
