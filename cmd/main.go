package main

import (
	"gcp-l4-proxy-monitoring/pkg"
	"os"
	"time"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/v4/args"
	"github.com/newrelic/infra-integrations-sdk/v4/integration"

	"github.com/newrelic/infra-integrations-sdk/v4/log"
)

const (
	IntegrationName    = "gcp_l4_proxy_metrics"
	IntegrationVersion = "0.1.0"
	EntityName         = "gcp:l4_proxy"
	EntityType         = "LoadBalancer"
	EntityDisplay      = "Google Cloud L4 Proxy Load Balancer Metrics"
)

const (
	ErrNewIntegration         = 1
	ErrNewEntity              = 2
	ErrArgTimes               = 3
	ErrArgName                = 4
	ErrArgFilePath            = 5
	ErrImportMetricNewConn    = 6
	ErrImportMetricClosedConn = 7
	ErrImportMetricEgress     = 8
	ErrImportMetricIngress    = 9
	ErrPublish                = 10
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

func main() {
	// Create integration
	i, err := integration.New(IntegrationName, IntegrationVersion, integration.Args(&args))
	if err != nil {
		log.Error("Error creating Nr Infra integration = ", err)
		os.Exit(ErrNewIntegration)
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
		log.Error("Either parameters 'start_time' / 'end_time' or 'since' must be defined and be bigger than zero.")
		os.Exit(ErrArgTimes)
	}

	if projectName == "" {
		log.Error("Parameter 'name' must be defined.")
		os.Exit(ErrArgName)
	}

	if filePath == "" {
		log.Error("File path must be defined.")
		os.Exit(ErrArgFilePath)
	}

	// Required by Google Cloud Monitoring library to perform the JWT authentication
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)

	// Create entity
	entity, err := i.NewEntity(EntityName, EntityType, EntityDisplay)
	if err != nil {
		log.Error("Error creating entity", err)
		os.Exit(ErrNewEntity)
	}

	// Create Metrics
	if args.All() || args.Metrics {
		// All responses contain the same resource labels. Just using the first one.
		newConnectionsMetrics, resourceLabels, err := pkg.ReadNewConnectionsMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading new connections metric = ", err)
			os.Exit(ErrImportMetricNewConn)
		}

		closedConnectionsMetrics, _, err := pkg.ReadClosedConnectionsMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading closed connections metric = ", err)
			os.Exit(ErrImportMetricClosedConn)
		}

		egressBytesMetrics, _, err := pkg.ReadEgressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading egress bytes metric = ", err)
			os.Exit(ErrImportMetricEgress)
		}

		ingressBytesMetrics, _, err := pkg.ReadIngressBytesMetric(
			projectName,
			startTime,
			endTime,
		)
		if err != nil {
			log.Error("Error reading ingress bytes metric = ", err)
			os.Exit(ErrImportMetricIngress)
		}

		pkg.ExportMetrics(entity, &pkg.L4ProxyMetrics{
			NewConn:    newConnectionsMetrics,
			ClosedConn: closedConnectionsMetrics,
			Egress:     egressBytesMetrics,
			Ingress:    ingressBytesMetrics,
			Attributes: resourceLabels,
		})
	}

	i.AddEntity(entity)

	err = i.Publish()
	if err != nil {
		log.Error("Error publishing = ", err)
		os.Exit(ErrPublish)
	}
}
