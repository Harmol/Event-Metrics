package manager

import (
	"context"
	"net/http"
	"time"

	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"netease.com/kubediag/event-metrics/pkg/metric"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EventManager struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	Ticker *time.Ticker
	Metrics *metric.EventMetrics
}

// NewEventManager create EventManager and expose Metrics to prometheus
func NewEventManager(
	cli client.Client,
	log logr.Logger,
	scheme *runtime.Scheme,
	duration time.Duration,
	metricsAddr string,
) *EventManager {
	em := EventManager{
		Client: cli,
		Log:    log,
		Scheme: scheme,
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
	err := em.Client.List(ctx, &eventList)
	if err != nil {
		log.Error(err, "Error in collecting EventList.")
		return err
	}

	em.Metrics.GenerateMetrics(eventList)

	log.Info("Generated Event Metrics.")
	return nil
}

// GenerateMetricsWithTicker generate metrics everytime Ticker ticks.
func (em *EventManager) GenerateMetricsWithTicker() {
	log := em.Log
	log.WithName("Ticker")

	log.Info("Start to generate Event Metrics with Ticker.", "time", time.Now().String())
	if err := em.GenerateMetrics(); err != nil {
		log.Error(err, "Error in generating metrics.")
		return
	}

	go func() {
		for {
			tick := <-em.Ticker.C
			log.Info("Time received from Ticker: %v", tick.String())

			if err := em.GenerateMetrics(); err != nil {
				log.Error(err, "Error in generating metrics.")
				return
			}
		}
	}()
}