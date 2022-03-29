package enterprises

import (
	"context"
	"fmt"
	"github.com/onosproject/aether-graphql/internal/store"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
	"github.com/onosproject/aether-roc-api/pkg/utils"
	externalRef1 "github.com/onosproject/config-models/modelplugin/aether-2.0.0/aether_2_0_0"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/openconfig/gnmi/proto/gnmi"
	"time"
)

var log = logging.GetLogger("enterprises")

// TODO: we should be maintaining our own models here
//type Enterprise struct {
//	ID          string
//	Name        string
//	Description string
//	Sites       []*sites.Site
//}

// TODO: List should work for both GraphQL and V1
func List(ctx context.Context, target string) (*types.Enterprises, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	gnmiGet, err := utils.NewGnmiGetRequest("/aether/v2.0.0/{target}/enterprises", store.Target)
	if err != nil {
		return nil, err
	}
	log.Info("gnmi-get-request", "req", gnmiGet.String())
	gnmiVal, err := utils.GetResponseUpdate(store.Client.Get(ctx, gnmiGet))
	if err != nil {
		log.Errorw("gnmi-get-error", "err", err)
		return nil, err
	}
	if gnmiVal == nil {
		return nil, nil
	}

	gnmiJsonVal, ok := gnmiVal.Value.(*gnmi.TypedValue_JsonVal)
	if !ok {
		return nil, fmt.Errorf("unexpected-type-of-reply-from-server: %v", gnmiVal.Value)
	}

	var gnmiResponse externalRef1.Device
	if err = externalRef1.Unmarshal(gnmiJsonVal.JsonVal, &gnmiResponse); err != nil {
		return nil, fmt.Errorf("error-unmarshalling-gnmiResponse: %v", err)
	}

	mpd := store.ModelPluginDevice{
		Device: gnmiResponse,
	}

	rocEnterprises, err := mpd.ToEnterprises()
	if err != nil {
		return nil, fmt.Errorf("error-casting-gnmiResponse-to-enterprise: %v", err)
	}

	log.Infow("received-gnmi-get-response", "rocEnterprises", rocEnterprises)
	return rocEnterprises, nil
}
