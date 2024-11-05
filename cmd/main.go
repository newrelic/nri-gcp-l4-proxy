package main

import (
	"encoding/json"
	"gcp-l4-proxy-monitoring/pkg"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//TODO: get from config:
	// - Time range (actually scheduling time)
	// - Project name
	// - JSON file path

	//TEST
	projectName := "projects/labs-team-333620"
	startTime := int64(1729593091)
	endTime := int64(1729593390)
	filePath := "/Users/asantaren/Desktop/key-default-compute-labs-team-333620-9101aa490097.json"

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)

	// log.Println("---- New connections count ----")
	newConnectionsMetrics, err := pkg.ReadNewConnectionsMetric(
		projectName,
		startTime,
		endTime,
	)
	if err != nil {
		log.Error("Error reading new connections metric = ", err)
		os.Exit(1)
	}
	// prettyPrint(&newConnectionsMetrics)

	// log.Println("---- Closed connections count ----")
	closedConnectionsMetrics, err := pkg.ReadClosedConnectionsMetric(
		projectName,
		startTime,
		endTime,
	)
	if err != nil {
		log.Error("Error reading closed connections metric = ", err)
		os.Exit(2)
	}
	// prettyPrint(&closedConnectionsMetrics)

	// log.Println("---- Egress bytes count ----")
	egressBytesMetrics, err := pkg.ReadEgressBytesMetric(
		projectName,
		startTime,
		endTime,
	)
	if err != nil {
		log.Error("Error reading egress bytes metric = ", err)
		os.Exit(3)
	}
	// prettyPrint(&egressBytesMetrics)

	// log.Println("---- Ingress bytes count ----")
	ingressBytesMetrics, err := pkg.ReadIngressBytesMetric(
		projectName,
		startTime,
		endTime,
	)
	if err != nil {
		log.Error("Error reading ingress bytes metric = ", err)
		os.Exit(4)
	}
	// prettyPrint(&ingressBytesMetrics)

	// log.Println("-------- NR INFRA OUTPUT --------")

	pkg.ExportData([]*pkg.DeltaCountMetrics{
		&newConnectionsMetrics,
		&closedConnectionsMetrics,
		&egressBytesMetrics,
		&ingressBytesMetrics,
	})
}

func prettyPrint(metrics *pkg.DeltaCountMetrics) {
	jcart, _ := json.MarshalIndent(metrics, "", "\t")
	log.Println("Response = \n", string(jcart))
}
