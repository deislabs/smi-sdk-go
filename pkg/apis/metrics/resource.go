package metrics

import (
	"path"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewTrafficMetrics constructs a TrafficMetrics with all the defaults
func NewTrafficMetrics(obj *v1.ObjectReference) *TrafficMetrics {
	var selfLink string

	// If Namespace is empty, it is assumed that this is a non-namespaced resource.
	if obj.Namespace == "" {
		selfLink = path.Join(baseURL(), getKindName(obj.Kind), obj.Name)
	} else {
		selfLink = path.Join(
			baseURL(),
			"namespaces",
			obj.Namespace,
			getKindName(obj.Kind),
			obj.Name)
	}

	metrics := []*Metric{}
	for _, m := range AvailableMetrics {
		n := *m
		metrics = append(metrics, &n)
	}

	return &TrafficMetrics{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TrafficMetrics",
			APIVersion: APIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			CreationTimestamp: metav1.Now(),
			Name:              obj.Name,
			Namespace:         obj.Namespace,
			SelfLink:          selfLink,
		},
		Resource: obj,
		Metrics:  metrics,
	}
}

// TrafficMetrics provide the metrics for a specific resource
type TrafficMetrics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	*Interval `json:",inline"`

	Resource *v1.ObjectReference `json:"resource"`
	Edge     *Edge               `json:"edge"`
	Metrics  []*Metric           `json:"metrics"`
}

// Get returns a metric associated with a name
func (t *TrafficMetrics) Get(name string) *Metric {
	for _, metric := range t.Metrics {
		if metric.Name == name {
			return metric
		}
	}

	return nil
}