package middleware

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// NewMetric associates prometheus.Collector based on Metric.Type
func NewMetric(m *ginprometheus.Metric, subsystem string) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case "counter_vec":
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "counter":
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "gauge_vec":
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "gauge":
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "histogram_vec":
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "histogram":
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case "summary_vec":
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case "summary":
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	}
	return metric
}

var defaultMetricsPath = "/metrics"

var defaultMetrics = []*ginprometheus.Metric{
	&ginprometheus.Metric{}, // TODO
}

type PrometheusMiddleware struct {
	MetricsPath string
	metricsList []*ginprometheus.Metric

	requests *prometheus.CounterVec
}

func newPrometheusMiddleware(subsystem string, metricsList []*ginprometheus.Metric) *PrometheusMiddleware {
	p := &PrometheusMiddleware{
		metricsList: metricsList,
		MetricsPath: defaultMetricsPath,
	}
	p.registerMetrics(subsystem)
	return p
}

func NewPrometheusMiddleware(subsystem string, metricsList ...[]*ginprometheus.Metric) *PrometheusMiddleware {
	metrics := metricsList[0]
	for _, metric := range defaultMetrics {
		metrics = append(metrics, metric)
	}
	return newPrometheusMiddleware(subsystem, metrics)
}

func (p *PrometheusMiddleware) registerMetrics(subsystem string) {
	for _, metricDef := range p.metricsList {
		metric := NewMetric(metricDef, subsystem)
		if err := prometheus.Register(metric); err != nil {
			log.WithError(err).Errorf("%s could not be registered in Prometheus", metricDef.Name)
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

		c.Next()
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (p *PrometheusMiddleware) SetMetricsPath(e *gin.Engine) {
	e.GET(p.MetricsPath, prometheusHandler())
}

func (p *PrometheusMiddleware) Use(e *gin.Engine) {
	e.Use(p.HandlerFunc())
	p.SetMetricsPath(e)
}
