package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/grnsv/metrics/internal/storage"
)

var (
	metrics   []storage.Metric
	pollCount int64
	client    = resty.New()
)

func collectMetrics() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	metrics = append(metrics,
		storage.GaugeMetric{Name: "Alloc", Value: float64(m.Alloc)},
		storage.GaugeMetric{Name: "BuckHashSys", Value: float64(m.BuckHashSys)},
		storage.GaugeMetric{Name: "Frees", Value: float64(m.Frees)},
		storage.GaugeMetric{Name: "GCCPUFraction", Value: m.GCCPUFraction},
		storage.GaugeMetric{Name: "GCSys", Value: float64(m.GCSys)},
		storage.GaugeMetric{Name: "HeapAlloc", Value: float64(m.HeapAlloc)},
		storage.GaugeMetric{Name: "HeapIdle", Value: float64(m.HeapIdle)},
		storage.GaugeMetric{Name: "HeapInuse", Value: float64(m.HeapInuse)},
		storage.GaugeMetric{Name: "HeapObjects", Value: float64(m.HeapObjects)},
		storage.GaugeMetric{Name: "HeapReleased", Value: float64(m.HeapReleased)},
		storage.GaugeMetric{Name: "HeapSys", Value: float64(m.HeapSys)},
		storage.GaugeMetric{Name: "LastGC", Value: float64(m.LastGC)},
		storage.GaugeMetric{Name: "Lookups", Value: float64(m.Lookups)},
		storage.GaugeMetric{Name: "MCacheInuse", Value: float64(m.MCacheInuse)},
		storage.GaugeMetric{Name: "MCacheSys", Value: float64(m.MCacheSys)},
		storage.GaugeMetric{Name: "MSpanInuse", Value: float64(m.MSpanInuse)},
		storage.GaugeMetric{Name: "MSpanSys", Value: float64(m.MSpanSys)},
		storage.GaugeMetric{Name: "Mallocs", Value: float64(m.Mallocs)},
		storage.GaugeMetric{Name: "NextGC", Value: float64(m.NextGC)},
		storage.GaugeMetric{Name: "NumForcedGC", Value: float64(m.NumForcedGC)},
		storage.GaugeMetric{Name: "NumGC", Value: float64(m.NumGC)},
		storage.GaugeMetric{Name: "OtherSys", Value: float64(m.OtherSys)},
		storage.GaugeMetric{Name: "PauseTotalNs", Value: float64(m.PauseTotalNs)},
		storage.GaugeMetric{Name: "StackInuse", Value: float64(m.StackInuse)},
		storage.GaugeMetric{Name: "StackSys", Value: float64(m.StackSys)},
		storage.GaugeMetric{Name: "Sys", Value: float64(m.Sys)},
		storage.GaugeMetric{Name: "TotalAlloc", Value: float64(m.TotalAlloc)},
		storage.CounterMetric{Name: "PollCount", Value: pollCount},
		storage.GaugeMetric{Name: "RandomValue", Value: rand.Float64()},
	)

	pollCount++
}

func sendMetrics() {
	for _, metric := range metrics {
		url := fmt.Sprintf("http://%s/update/%s/%s/%v",
			config.ServerURL.String(), metric.GetType(), metric.GetName(), metric.GetValue())
		_, err := client.R().
			SetHeader("Content-Type", "text/plain").
			Post(url)
		if err != nil {
			log.Printf("Error sending request: %v", err)
		}
	}

	metrics = metrics[:0]
}

func main() {
	if err := parseVars(); err != nil {
		log.Fatalf("Agent failed: %v", err)
	}
	go func() {
		for {
			collectMetrics()
			time.Sleep(config.PollInterval.Duration)
		}
	}()

	for {
		sendMetrics()
		time.Sleep(config.ReportInterval.Duration)
	}
}
