package storage

type Storage interface {
	UpdateGauge(name string, value float64)
	UpdateCounter(name string, value int64)
	GetGauge(name string) (float64, bool)
	GetCounter(name string) (int64, bool)
	GetAllMetrics() map[string]interface{}
}

type MemStorage struct {
	gauges   map[string]float64
	counters map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]float64),
		counters: make(map[string]int64),
	}
}

func (ms *MemStorage) UpdateGauge(name string, value float64) {
	ms.gauges[name] = value
}

func (ms *MemStorage) UpdateCounter(name string, value int64) {
	ms.counters[name] += value
}

func (ms *MemStorage) GetGauge(name string) (float64, bool) {
	value, ok := ms.gauges[name]
	return value, ok
}

func (ms *MemStorage) GetCounter(name string) (int64, bool) {
	value, ok := ms.counters[name]
	return value, ok
}

func (ms *MemStorage) GetAllMetrics() map[string]interface{} {
	allMetrics := make(map[string]interface{})
	for name, value := range ms.gauges {
		allMetrics[name] = value
	}
	for name, value := range ms.counters {
		allMetrics[name] = value
	}
	return allMetrics
}
