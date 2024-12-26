package storage

type MetricType string

const (
	GaugeType   MetricType = "gauge"
	CounterType MetricType = "counter"
)

type Metric interface {
	GetType() MetricType
	GetName() string
	GetValue() interface{}
}

type GaugeMetric struct {
	Name  string
	Value float64
}

func (m GaugeMetric) GetType() MetricType {
	return GaugeType
}

func (m GaugeMetric) GetName() string {
	return m.Name
}

func (m GaugeMetric) GetValue() interface{} {
	return m.Value
}

type CounterMetric struct {
	Name  string
	Value int64
}

func (m CounterMetric) GetType() MetricType {
	return CounterType
}

func (m CounterMetric) GetName() string {
	return m.Name
}

func (m CounterMetric) GetValue() interface{} {
	return m.Value
}
