// Import L4 Proxy data from the Google Cloud Monitoring API.

package pkg

import (
	"context"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/api/iterator"
)

// Import data from Google Cloud Metric API.
// name: project, organization or folder name
// filter: filter applied to data
// startTime: star time in seconds (unix epoch)
// endTime: star time in seconds (unix epoch)
// return: time series data or error
func importMetrics(metricName string, reqName string, filter string, startTime int64, endTime int64) (DeltaCountMetrics, map[string]string, error) {
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		// Could not crate metric client
		return DeltaCountMetrics{}, nil, err
	}

	interval := monitoringpb.TimeInterval{
		EndTime: &googlepb.Timestamp{
			Seconds: endTime,
		},
		StartTime: &googlepb.Timestamp{
			Seconds: startTime,
		},
	}

	req := &monitoringpb.ListTimeSeriesRequest{
		Name:     reqName,
		Filter:   filter,
		Interval: &interval,
	}

	it := c.ListTimeSeries(ctx, req)
	metricValues := []DeltaCountMetricValue{}
	var resourceLabels map[string]string
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// Could not read next data block
			return DeltaCountMetrics{}, nil, err
		}
		if len(metricValues) == 0 {
			resourceLabels = resp.Resource.Labels
		}
		for _, point := range resp.Points {
			metricValue := MakeDeltaCountMetricValue(point)
			metricValues = append(metricValues, metricValue)
		}
	}
	metrics := DeltaCountMetrics{
		Name:   metricName,
		Values: metricValues,
	}

	return metrics, resourceLabels, nil
}

// Read l4_proxy/tcp/new_connections_count metric.
func ReadNewConnectionsMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, map[string]string, error) {
	return importMetrics(
		"gcp.l4_proxy.new_connections",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/tcp/new_connections_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/tcp/closed_connections_count metric.
func ReadClosedConnectionsMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, map[string]string, error) {
	return importMetrics(
		"gcp.l4_proxy.closed_connections",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/tcp/closed_connections_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/egress_bytes_count metric.
func ReadEgressBytesMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, map[string]string, error) {
	return importMetrics(
		"gcp.l4_proxy.egress_bytes",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/egress_bytes_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/ingress_bytes_count metric.
func ReadIngressBytesMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, map[string]string, error) {
	return importMetrics(
		"gcp.l4_proxy.ingress_bytes",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/ingress_bytes_count\"",
		startTime,
		endTime,
	)
}
