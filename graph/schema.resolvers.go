package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/onosproject/aether-graphql/internal/devicegroups"
	"github.com/onosproject/aether-graphql/internal/devices"
	"github.com/onosproject/aether-graphql/internal/slices"
	"github.com/onosproject/aether-graphql/internal/smallcells"
	"github.com/onosproject/aether-graphql/internal/store"

	"github.com/onosproject/aether-graphql/graph/generated"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-graphql/internal/enterprises"
	"github.com/onosproject/aether-graphql/internal/simcards"
	"github.com/onosproject/aether-graphql/internal/sites"
)

func (r *queryResolver) Enterprises(ctx context.Context) ([]*model.Enterprise, error) {
	var ents []*model.Enterprise

	rocEnterprises, err := enterprises.List(ctx, "connectivity-service-v2")
	if err != nil {
		return nil, err
	}

	for _, re := range *rocEnterprises.Enterprise {
		var apps []*model.Application
		for _, ra := range *re.Application {
			app := &model.Application{
				ID:   ra.ApplicationId,
				Name: *ra.DisplayName,
			}
			apps = append(apps, app)
		}

		e := &model.Enterprise{
			ID:           re.EnterpriseId,
			Name:         *re.DisplayName,
			Description:  *re.Description,
			Applications: apps,
			Sites:        sites.List(ctx, re),
		}
		ents = append(ents, e)
	}

	return ents, nil
}

func (r *queryResolver) Applications(ctx context.Context) ([]*model.Application, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sites(ctx context.Context) ([]*model.Site, error) {
	var s []*model.Site

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		s = append(s, sites.List(ctx, re)...)
	}
	return s, nil
}

func (r *queryResolver) SmallCells(ctx context.Context) ([]*model.SmallCell, error) {
	var results []*model.SmallCell

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		for _, rs := range *re.Site {
			results = append(results, smallcells.List(ctx, rs)...)
		}
	}
	return results, nil
}

func (r *queryResolver) Slices(ctx context.Context) ([]*model.Slice, error) {
	var results []*model.Slice

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		for _, rs := range *re.Site {
			results = append(results, slices.List(ctx, rs)...)
		}
	}
	return results, nil
}

func (r *queryResolver) DeviceGroups(ctx context.Context) ([]*model.DeviceGroup, error) {
	var results []*model.DeviceGroup

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		for _, rs := range *re.Site {
			results = append(results, devicegroups.List(ctx, rs)...)
		}
	}
	return results, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	var results []*model.Device

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		for _, rs := range *re.Site {
			results = append(results, devices.List(ctx, rs)...)
		}
	}
	return results, nil
}

func (r *queryResolver) SimCards(ctx context.Context) ([]*model.SimCard, error) {
	var results []*model.SimCard

	rocEnterprises, err := enterprises.List(ctx, store.Target)
	if err != nil {
		return nil, err
	}
	for _, re := range *rocEnterprises.Enterprise {
		for _, rs := range *re.Site {
			results = append(results, simcards.List(ctx, rs)...)
		}
	}
	return results, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
