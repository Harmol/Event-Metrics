package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type EventMetrics struct {
	kubeEventInfo *prometheus.GaugeVec
	kubeEventSource *prometheus.GaugeVec
	kubeEventInvolvedObject *prometheus.GaugeVec
	kubeEventCount *prometheus.GaugeVec
}

func NewEventMetrics() *EventMetrics {
	kubeEventInfo := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_event_info",
			Help: "Information about Event",
		},
		[]string{
			"event_name",
			"event_namespace",
			"event_kind",
			"event_first_time_stamp",
			"event_type",
			"event_reason",
			"event_reporting_component",
			"event_reporting_instance",
			"event_uid",
		})

	kubeEventSource := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_event_source",
			Help: "Source of event",
		},
		[]string{
			"event_name",
			"event_namespace",
			"source_component",
			"source_host",
			"event_uid",
		})

	kubeEventInvolvedObject := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_event_involved_object",
			Help: "Kubernetes object that involved the event",
		},
		[]string{
			"event_name",
			"event_namespace",
			"involved_object_kind",
			"involved_object_name",
			"involved_object_namespace",
			"involved_object_uid",
			"event_uid",
		})

	kubeEventCount := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kube_event_count",
			Help: "Count label of event",
		},
		[]string{
			"event_name",
			"event_namespace",
			"event_uid",
		})

	em := &EventMetrics{
		kubeEventInfo: kubeEventInfo,
		kubeEventSource: kubeEventSource,
		kubeEventInvolvedObject: kubeEventInvolvedObject,
		kubeEventCount: kubeEventCount,
	}
	return em
}


// RegistryMetrics registry designed metrics for prometheus
func (em *EventMetrics) RegistryMetrics() {
	metrics.Registry.MustRegister(
		em.kubeEventInfo,
		em.kubeEventSource,
		em.kubeEventInvolvedObject,
		em.kubeEventCount,
		)
}

// GenerateMetrics generate metrics with given EventList
func (em *EventMetrics) GenerateMetrics(eventList corev1.EventList) {
	for _, event := range eventList.Items {
		em.kubeEventInfo.WithLabelValues(
			event.Name,
			event.Namespace,
			event.Kind,
			event.FirstTimestamp.String(),
			event.Type,
			event.Reason,
			event.ReportingController,
			event.ReportingInstance,
			string(event.UID),
		).Set(1)

		em.kubeEventSource.WithLabelValues(
			event.Name,
			event.Namespace,
			event.Source.Component,
			event.Source.Host,
			string(event.UID),
		).Set(1)

		em.kubeEventInvolvedObject.WithLabelValues(
			event.Name,
			event.Namespace,
			event.InvolvedObject.Kind,
			event.InvolvedObject.Name,
			event.InvolvedObject.Namespace,
			string(event.InvolvedObject.UID),
			string(event.UID),
		).Set(1)

		em.kubeEventCount.WithLabelValues(
			event.Name,
			event.Namespace,
			string(event.UID),
		).Set(float64(event.Count))
	}
}