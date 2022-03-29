package store

import (
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"os"
	"time"
)

var log = logging.GetLogger("GnmiClient")

var Client gnmi.GNMIClient

var Target string

func ConnectGNMI(address string) error {
	opts, err := certs.HandleCertPaths("", "", "", true)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	optsWithRetry := []grpc.DialOption{
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithInterval(100 * time.Millisecond))),
	}
	optsWithRetry = append(opts, optsWithRetry...)
	gnmiConn, err := grpc.Dial(address, optsWithRetry...)
	if err != nil {
		log.Errorw("cannot-connect-to-gnmi-server", "err", err, "address", address)
		return err
	}

	Client = gnmi.NewGNMIClient(gnmiConn)
	return nil
}

func SetTarget(target string) {
	Target = target
}
