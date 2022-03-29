package prom

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	promApi "github.com/prometheus/client_golang/api"
	promApiV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
	"time"
)

var log = logging.GetLogger("prometheus")

// PrometheusClient -
type PrometheusClient interface {
	Init() error
	Query(query string) (promModel.Value, error)
}

// PrometheusConnection -
type PrometheusConnection struct {
	Address string
	client  promApi.Client
	v1api   promApiV1.API
}

// Init -
func (a *PrometheusConnection) Init() error {
	log.Infof("Initializing new PrometheusConnection")

	var err error
	a.client, err = promApi.NewClient(promApi.Config{
		Address: a.Address,
	})
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}

	a.v1api = promApiV1.NewAPI(a.client)

	return nil
}

// Query executes a query
func (a *PrometheusConnection) Query(query string) (promModel.Value, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, warnings, err := a.v1api.Query(ctx, query, time.Now())
	if err != nil {
		return nil, fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		fmt.Printf("Prometheus Warnings: %v\n", warnings)
	}

	// result is a Value, which is an interface to ValueType and a String() method
	// Can cast to:
	//    Matrix, Vector, *Scalar, *String
	return result, nil
}
