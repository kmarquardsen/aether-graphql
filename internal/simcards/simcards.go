package simcards

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
)

func List(ctx context.Context, site types.EnterprisesEnterpriseSite) []*model.SimCard {
	var simCards []*model.SimCard
	for _, sim := range *site.SimCard {
		d := &model.SimCard{
			ID:   sim.SimId,
			Name: *sim.DisplayName,
		}
		simCards = append(simCards, d)
	}
	return simCards
}
