package main

import (
	"encoding/json"
	"gcp-l4-proxy-monitoring/pkg"
	"os"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/v4/args"
	"github.com/newrelic/infra-integrations-sdk/v4/integration"
	log "github.com/sirupsen/logrus"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
}

var (
	args argumentList
)

const (
	integrationName    = "gcp_l4_proxy_metrics"
	integrationVersion = "0.1.0"
	entityName         = "gcp:l4_proxy"
	entityType         = "LoadBalancer"
	entityDisplay      = "Google Cloud L4 Proxy Load Balancer Metrics"
)

func main() {
	//TODO: get from config:
	// - Time range (actually scheduling time)
	// - Project name
	// - JSON file path
	projectName := "projects/labs-team-333620"
	startTime := int64(1729593091)
	endTime := int64(1729593390)
	filePath := "/Users/asantaren/Desktop/key-default-compute-labs-team-333620-9101aa490097.json"

	// Required by Google Cloud Monitoring library to perform tje JWT authentication
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)

	// Create integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error("Error creating Nr Infra integration", err)
		os.Exit(1)
	}

	// Create entity
	entity, err := i.NewEntity(entityName, entityType, entityDisplay)
	if err != nil {
		log.Error("Error creating entity", err)
		os.Exit(2)
	}

	// Create Metrics
	if args.All() || args.Metrics {
		newConnectionsMetrics, err := pkg.ReadNewConnectionsMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading new connections metric = ", err)
			os.Exit(3)
		}

		closedConnectionsMetrics, err := pkg.ReadClosedConnectionsMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading closed connections metric = ", err)
			os.Exit(4)
		}

		egressBytesMetrics, err := pkg.ReadEgressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading egress bytes metric = ", err)
			os.Exit(5)
		}

		ingressBytesMetrics, err := pkg.ReadIngressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading ingress bytes metric = ", err)
			os.Exit(6)
		}

		pkg.ExportData(entity, &pkg.L4ProxyMetrics{
			NewConn:    newConnectionsMetrics,
			ClosedConn: closedConnectionsMetrics,
			Egress:     egressBytesMetrics,
			Ingress:    ingressBytesMetrics,
			Attributes: map[string]string{}, //TODO: populate attributes from the response
		})
	}

	i.AddEntity(entity)

	err = i.Publish()
	if err != nil {
		log.Error("Error publishing", err)
		os.Exit(7)
	}
}

func prettyPrint(metrics *pkg.DeltaCountMetrics) {
	jcart, _ := json.MarshalIndent(metrics, "", "\t")
	log.Println("Response = \n", string(jcart))
}
