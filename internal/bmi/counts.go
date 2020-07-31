package bmi

import (
	"expvar"
	"sync"
)

var lock sync.RWMutex

// Counts keep track of served calculations
type Counts struct {
	Total        uint64
	Calculations map[string]uint64
	Errors       uint64
	Average      float64
	LastCalculation string
}

var counts Counts

func NewCounts() Counts {
	return Counts{
		Calculations: make(map[string]uint64),
	}
}

func init() {
	counts = NewCounts()
}

func (c *Counts) register(bmi float64) {
	lock.Lock()
	defer lock.Unlock()

	level := DescribeLevel(bmi)
	c.Calculations[level]++

	// Update total number of calculations
	c.Total++

	// Update running average
	c.Average = c.Average + (bmi-c.Average)/float64(c.Total)
}

func (c *Counts) registerError() {
	lock.Lock()
	defer lock.Unlock()

	c.Errors++
}

func GetCounts() Counts {
	lock.RLock()
	defer lock.RUnlock()

	var calculations map[string]uint64 = make(map[string]uint64)
	for k, v := range counts.Calculations {
		calculations[k] = v
	}

	var lastCalculation string = "n/a"
	v := expvar.Get("LastCalculation")
	vs, ok := v.(*expvar.String); if ok {
		lastCalculation = vs.Value()
	}
	return Counts{
		Total:        counts.Total,
		Calculations: calculations,
		Errors:       counts.Errors,
		Average:      counts.Average,
		LastCalculation: lastCalculation,
	}
}
