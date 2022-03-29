package sites

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-graphql/internal/devices"
	"github.com/onosproject/aether-graphql/internal/simcards"
	"github.com/onosproject/aether-graphql/internal/slices"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
)

func List(ctx context.Context, enterprise types.EnterprisesEnterprise) []*model.Site {
	var sites []*model.Site
	for _, rs := range *enterprise.Site {
		simCards := simcards.List(ctx, rs)
		s := &model.Site{
			ID:       rs.SiteId,
			Name:     *rs.DisplayName,
			Devices:  devices.List(ctx, rs),
			SimCards: simCards,
			Slices:   slices.List(rs),
		}
		sites = append(sites, s)
	}

	return sites
}
