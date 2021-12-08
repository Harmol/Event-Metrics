package manager

import (
	"context"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/cache"

	"netease.com/kubediag/event-metrics/pkg/metric"
)

type EventManager struct {
	Cache  cache.Cache
	Log    logr.Logger
	Ticker *time.Ticker
	Metrics *metric.EventMetrics
}

// NewEventManager create EventManager and expose Metrics to prometheus
func NewEventManager(
	cache cache.Cache,
	log logr.Logger,
	duration time.Duration,
	metricsAddr string,
) *EventManager {
	em := EventManager{
		Cache:  cache,
		Log:    log,
		Ticker: time.NewTicker(duration),
		Metrics: metric.NewEventMetrics(),
	}

	// registry Event Metrics
	// and set handler for prometheus
	em.Metrics.RegistryMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(metricsAddr, nil)
	return &em
}

// GenerateMetrics get EventList and use it to generate metrics
func (em *EventManager) GenerateMetrics() error {
	ctx := context.Background()
	log := em.Log

	log.Info("generating Event Metrics.")
	var eventList corev1.EventList
	err := em.Cache.List(ctx, &eventList)
	if err != nil {
		log.Error(err, "Error in collecting EventList.")
		return err
	}

	em.Metrics.GenerateMetrics(eventList)

	log.Info("Generated Event Metrics.")
	return nil
}

// Start generate metrics everytime Ticker ticks.
func (em *EventManager) Start() {
	log := em.Log

	for ;; <-em.Ticker.C {
		log.Info("start to generate EventMetrics", "time", time.Now().String())
		if err := em.GenerateMetrics(); err != nil {
			log.Error(err, "Error in generating metrics.")
			return
		}
	}
}