// Import L4 Proxy data from the Google Cloud Monitoring API.

package pkg

import (
	"context"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

// Import data from Google Cloud Metric API.
// name: project, organization or folder name
// filter: filter applied to data
// startTime: star time in seconds (unix epoch)
// endTime: star time in seconds (unix epoch)
// return: time series data or error
func importData(metricName string, reqName string, filter string, startTime int64, endTime int64) (DeltaCountMetrics, error) {
	ctx := context.Background()
	c, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Error("Could not crate metric client: ", err)
		return DeltaCountMetrics{}, err
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
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Error("Could not read next data block: ", err)
			return DeltaCountMetrics{}, err
		}

		// // Pretty print response
		// jcart, _ := json.MarshalIndent(resp, "", "\t")
		// log.Println("Response = \n", string(jcart))

		for _, point := range resp.Points {
			metricValue := FromPointToMetricValue(point)
			metricValues = append(metricValues, metricValue)
		}
	}
	metrics := DeltaCountMetrics{
		Name:       metricName,
		Values:     metricValues,
		Attributes: map[string]any{}, //TODO: populate attributes from the response
	}

	return metrics, nil
}

// Read l4_proxy/tcp/new_connections_count metric.
func ReadNewConnectionsMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, error) {
	return importData(
		"new_connections",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/tcp/new_connections_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/tcp/closed_connections_count metric.
func ReadClosedConnectionsMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, error) {
	return importData(
		"closed_connections",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/tcp/closed_connections_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/egress_bytes_count metric.
func ReadEgressBytesMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, error) {
	return importData(
		"egress_bytes",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/egress_bytes_count\"",
		startTime,
		endTime,
	)
}

// Read l4_proxy/ingress_bytes_count metric.
func ReadIngressBytesMetric(name string, startTime int64, endTime int64) (DeltaCountMetrics, error) {
	return importData(
		"ingress_bytes",
		name,
		"metric.type = \"loadbalancing.googleapis.com/l4_proxy/ingress_bytes_count\"",
		startTime,
		endTime,
	)
}
