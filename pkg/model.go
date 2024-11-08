package pkg

import (
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
)

// L4 proxy metrics
type L4ProxyMetrics struct {
	NewConn    DeltaCountMetrics
	ClosedConn DeltaCountMetrics
	Egress     DeltaCountMetrics
	Ingress    DeltaCountMetrics
	Attributes map[string]string
}

// Delta count metric.
type DeltaCountMetrics struct {
	Name   string
	Values []DeltaCountMetricValue
}

// Delta count metric.
type DeltaCountMetricValue struct {
	Value    int64
	Interval TimeInterval
}

// Delta time.
type TimeInterval struct {
	// Unix epoch in millis
	StartTime int64
	// Unix epoch in millis
	EndTime int64
}

func makeTimeInterval(interval *monitoringpb.TimeInterval) TimeInterval {
	return TimeInterval{
		StartTime: interval.StartTime.Seconds*1_000 + int64(interval.StartTime.Nanos)/1_000_000,
		EndTime:   interval.EndTime.Seconds*1_000 + int64(interval.EndTime.Nanos)/1_000_000,
	}
}

func MakeDeltaCountMetricValue(point *monitoringpb.Point) DeltaCountMetricValue {
	return DeltaCountMetricValue{
		Interval: makeTimeInterval(point.GetInterval()),
		Value:    point.GetValue().GetInt64Value(),
	}
}
