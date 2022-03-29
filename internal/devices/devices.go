package devices

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
)

func List(ctx context.Context, site types.EnterprisesEnterpriseSite) []*model.Device {
	var devices []*model.Device
	for _, rd := range *site.Device {
		d := &model.Device{
			ID:   rd.DeviceId,
			Name: *rd.DisplayName,
			//SimCard: rd.SimCard, // TODO simcards.Get()
		}
		devices = append(devices, d)
	}
	return devices
}
