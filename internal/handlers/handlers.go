package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/grnsv/metrics/internal/storage"
)

var store = storage.NewMemStorage()

func NewRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", HandleUpdateMetric)
	r.Get("/value/{type}/{name}", HandleGetMetricValue)
	r.Get("/", HandleGetAllMetrics)
	return r
}

func HandleUpdateMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	metricType := storage.MetricType(chi.URLParam(r, "type"))
	metricName := chi.URLParam(r, "name")
	metricValue := chi.URLParam(r, "value")

	if metricName == "" {
		http.Error(w, "Metric name is required", http.StatusNotFound)
		return
	}

	switch metricType {
	case storage.GaugeType:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "Invalid gauge value", http.StatusBadRequest)
			return
		}
		store.UpdateGauge(metricName, value)

	case storage.CounterType:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "Invalid counter value", http.StatusBadRequest)
			return
		}
		store.UpdateCounter(metricName, value)

	default:
		http.Error(w, "Invalid metric type", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetMetricValue(w http.ResponseWriter, r *http.Request) {
	metricType := storage.MetricType(chi.URLParam(r, "type"))
	metricName := chi.URLParam(r, "name")

	var value interface{}
	var ok bool

	switch metricType {
	case storage.GaugeType:
		value, ok = store.GetGauge(metricName)
	case storage.CounterType:
		value, ok = store.GetCounter(metricName)
	default:
		http.Error(w, "Invalid metric type", http.StatusBadRequest)
		return
	}

	if !ok {
		http.Error(w, "Metric not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%v", value)
}

func HandleGetAllMetrics(w http.ResponseWriter, r *http.Request) {
	allMetrics := store.GetAllMetrics()
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Metrics</title>
	</head>
	<body>
		<h1>Metrics</h1>
		<table border="1">
			<tr>
				<th>Name</th>
				<th>Value</th>
			</tr>
			{{range $name, $value := .}}
			<tr>
				<td>{{$name}}</td>
				<td>{{$value}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`
	t, err := template.New("metrics").Parse(tmpl)
	if err != nil {
		http.Error(w, "Error generating HTML", http.StatusInternalServerError)
		return
	}
	t.Execute(w, allMetrics)
}
