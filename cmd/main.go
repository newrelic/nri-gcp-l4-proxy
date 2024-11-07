package main

import (
	"encoding/json"
	"gcp-l4-proxy-monitoring/pkg"
	"os"
	"time"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/v4/args"
	"github.com/newrelic/infra-integrations-sdk/v4/integration"
	log "github.com/sirupsen/logrus"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Name      string `default:"" help:"Google Cloud project, organization or folder name. Example: 'projects/my-project-555555'"`
	FilePath  string `default:"" help:"Service account JSON file path, used for JWT authentication."`
	Since     int    `default:"0" help:"Time frame of the request in seconds, starting from now. If set, start_time and end_time will be ignored."`
	StartTime int    `default:"0" help:"Start time in UNIX epoch, seconds."`
	EndTime   int    `default:"0" help:"End time in UNIX epoch, seconds."`
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
	// Create integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error("Error creating Nr Infra integration = ", err)
		os.Exit(1)
	}

	// Build args
	projectName := args.Name
	filePath := args.FilePath
	var startTime int64
	var endTime int64

	if args.Since > 0 {
		endTime = time.Now().Unix()
		startTime = endTime - int64(args.Since)
	} else if args.EndTime > 0 && args.StartTime > 0 {
		endTime = int64(args.EndTime)
		startTime = int64(args.StartTime)
	} else {
		log.Error("Either start/end times or time_frame must be set and bigger than zero.")
		os.Exit(2)
	}

	// Required by Google Cloud Monitoring library to perform the JWT authentication
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)

	// Create entity
	entity, err := i.NewEntity(entityName, entityType, entityDisplay)
	if err != nil {
		log.Error("Error creating entity", err)
		os.Exit(3)
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
			os.Exit(4)
		}

		closedConnectionsMetrics, err := pkg.ReadClosedConnectionsMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading closed connections metric = ", err)
			os.Exit(5)
		}

		egressBytesMetrics, err := pkg.ReadEgressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading egress bytes metric = ", err)
			os.Exit(6)
		}

		ingressBytesMetrics, err := pkg.ReadIngressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading ingress bytes metric = ", err)
			os.Exit(7)
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
		os.Exit(8)
	}
}

func prettyPrint(metrics *pkg.DeltaCountMetrics) {
	jcart, _ := json.MarshalIndent(metrics, "", "\t")
	log.Println("Response = \n", string(jcart))
}
