package devicegroups

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
)

func List(ctx context.Context, site types.EnterprisesEnterpriseSite) []*model.DeviceGroup {
	var deviceGroups []*model.DeviceGroup
	for _, dg := range *site.DeviceGroup {
		d := &model.DeviceGroup{
			ID:   dg.DeviceGroupId,
			Name: *dg.DisplayName,
			//Devices: TODO
		}
		deviceGroups = append(deviceGroups, d)
	}
	return deviceGroups
}
