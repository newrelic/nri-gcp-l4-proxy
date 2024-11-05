// Export data using the NR Infra data model.

package pkg

import (
	"time"

	"github.com/newrelic/infra-integrations-sdk/v4/integration"
	log "github.com/sirupsen/logrus"
)

const integrationName = "gcp_l4_proxy_metrics"
const integrationVersion = "0.1.0"
const entityName = "gcp:l4_proxy"
const entityType = "LoadBalancer"
const entityDisplay = "Google Cloud L4 Proxy Load Balancer Metrics"

func ExportData(data *L4ProxyMetrics) error {

	// Create integration
	i, err := integration.New(integrationName, integrationVersion)
	if err != nil {
		log.Error("Error creating Nr Infra integration", err)
		return err
	}

	// Create entity
	entity, err := i.NewEntity(entityName, entityType, entityDisplay)
	if err != nil {
		log.Error("Error creating entity", err)
		return err
	}

	// Add metrics
	addMetrics(entity, &data.NewConn)
	addMetrics(entity, &data.ClosedConn)
	addMetrics(entity, &data.Egress)
	addMetrics(entity, &data.Ingress)

	// Define common attributes
	for key, val := range data.Attributes {
		entity.AddCommonDimension(key, val)
	}

	//TODO: define inventory with load balancer metadata provided in the API response

	i.AddEntity(entity)

	err = i.Publish()
	if err != nil {
		log.Error("Error publishing", err)
		return err
	}

	return nil
}

// Add metrics
func addMetrics(entity *integration.Entity, metrics *DeltaCountMetrics) {
	for _, d := range metrics.Values {
		count, _ := integration.Count(time.UnixMilli(d.Interval.EndTime), metrics.Name, float64(d.Value))
		entity.AddMetric(count)
	}
}
