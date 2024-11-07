// Export data using the NR Infra data model.

package pkg

import (
	"time"

	"github.com/newrelic/infra-integrations-sdk/v4/integration"
)

func ExportMetrics(entity *integration.Entity, data *L4ProxyMetrics) error {
	// Add metrics
	err := addMetrics(entity, &data.NewConn)
	if err != nil {
		return err
	}
	err = addMetrics(entity, &data.ClosedConn)
	if err != nil {
		return err
	}
	err = addMetrics(entity, &data.Egress)
	if err != nil {
		return err
	}
	err = addMetrics(entity, &data.Ingress)
	if err != nil {
		return err
	}

	// Define common attributes
	for key, val := range data.Attributes {
		entity.AddCommonDimension(key, val)
	}

	return nil
}

// Add metrics
func addMetrics(entity *integration.Entity, metrics *DeltaCountMetrics) error {
	for _, d := range metrics.Values {
		count, err := integration.Count(time.UnixMilli(d.Interval.EndTime), metrics.Name, float64(d.Value))
		if err != nil {
			return err
		}
		entity.AddMetric(count)
	}
	return nil
}
