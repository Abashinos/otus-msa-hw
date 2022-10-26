package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Metric struct {
	Namespace   string
	Name        string
	Help        string
	ConstLabels prometheus.Labels
	LabelNames  []string
	Buckets     []float64

	MetricCollector prometheus.Collector

	ID   string
	Type string
}

// NewMetric associates prometheus.Collector based on Metric.Type
func NewMetric(m *Metric, subsystem string, labels map[string]string) prometheus.Collector {
	var collector prometheus.Collector
	m.ConstLabels = labels
	switch m.Type {
	case "counter_vec":
		collector = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
			m.LabelNames,
		)
	case "counter":
		collector = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
		)
	case "gauge_vec":
		collector = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
			m.LabelNames,
		)
	case "gauge":
		collector = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
		)
	case "histogram_vec":
		collector = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
				Buckets:     m.Buckets,
			},
			m.LabelNames,
		)
	case "histogram":
		collector = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
				Buckets:     m.Buckets,
			},
		)
	case "summary_vec":
		collector = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
			m.LabelNames,
		)
	case "summary":
		collector = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem:   subsystem,
				Namespace:   m.Namespace,
				Name:        m.Name,
				Help:        m.Help,
				ConstLabels: m.ConstLabels,
			},
		)
	}
	return collector
}

var defaultMetricsPath = "/metrics"

var reqTotal = &Metric{
	ID:         "reqTotal",
	Name:       "requests_total",
	Help:       "Total requests",
	Type:       "counter_vec",
	LabelNames: []string{"method", "path", "status_code"},
}
var reqDuration = &Metric{
	ID:   "reqDuration",
	Name: "request_duration_seconds",
	Help: "Duration of HTTP requests",
	Type: "histogram_vec",
	Buckets: []float64{
		0.000000001, // 1ns
		0.000000002,
		0.000000005,
		0.00000001, // 10ns
		0.00000002,
		0.00000005,
		0.0000001, // 100ns
		0.0000002,
		0.0000005,
		0.000001, // 1µs
		0.000002,
		0.000005,
		0.00001, // 10µs
		0.00002,
		0.00005,
		0.0001, // 100µs
		0.0002,
		0.0005,
		0.001, // 1ms
		0.002,
		0.005,
		0.01, // 10ms
		0.02,
		0.05,
		0.1, // 100 ms
		0.2,
		0.5,
		1.0, // 1s
		2.0,
		5.0,
		10.0, // 10s
		15.0,
		20.0,
		30.0,
	},
	LabelNames: []string{"method", "path"},
}

var defaultMetrics = []*Metric{
	reqTotal,
	reqDuration,
	// TODO
}

type PrometheusMiddleware struct {
	MetricsPath string
	metricsList []*Metric

	requests *prometheus.CounterVec
}

func newPrometheusMiddleware(subsystem string,
	metricsList []*Metric, labels map[string]string) *PrometheusMiddleware {
	p := &PrometheusMiddleware{
		metricsList: metricsList,
		MetricsPath: defaultMetricsPath,
	}
	p.registerMetrics(subsystem, labels)
	return p
}

func NewPrometheusMiddleware(subsystem string, labels map[string]string) *PrometheusMiddleware {
	return newPrometheusMiddleware(subsystem, defaultMetrics, labels)
}

func (p *PrometheusMiddleware) SetMetricsPath(path string) {
	p.MetricsPath = path
}

func (p *PrometheusMiddleware) registerMetrics(subsystem string, labels map[string]string) {
	if labels == nil {
		labels = make(map[string]string)
	}
	for _, metricDef := range p.metricsList {
		metric := NewMetric(metricDef, subsystem, labels)
		if err := prometheus.Register(metric); err != nil {
			log.WithError(err).Errorf("%s could not be registered in Prometheus", metricDef.Name)
		} else {
			log.Printf("Successfully registered metric '%s' in Prometheus", metricDef.Name)
		}
		metricDef.MetricCollector = metric
	}
}

func (p *PrometheusMiddleware) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == p.MetricsPath {
			c.Next()
			return
		}
		method := c.Request.Method
		path := c.Request.URL.Path

		start := time.Now()
		c.Next()
		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9

		status := strconv.Itoa(c.Writer.Status())
		reqTotal.MetricCollector.(*prometheus.CounterVec).WithLabelValues(method, path, status).Inc()
		if c.Writer.Status() < 400 {
			reqDuration.MetricCollector.(*prometheus.HistogramVec).WithLabelValues(method, path).Observe(elapsed)
		}
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (p *PrometheusMiddleware) RegisterMetricsRoute(e *gin.Engine) {
	e.GET(p.MetricsPath, prometheusHandler())
}

func (p *PrometheusMiddleware) Use(e *gin.Engine) {
	e.Use(p.HandlerFunc())
	p.RegisterMetricsRoute(e)
}
