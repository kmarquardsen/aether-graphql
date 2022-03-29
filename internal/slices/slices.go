package slices

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
)

func List(ctx context.Context, site types.EnterprisesEnterpriseSite) []*model.Slice {
	var slices []*model.Slice
	for _, slice := range *site.Slice {
		d := &model.Slice{
			ID:   slice.SliceId,
			Name: *slice.DisplayName,
		}
		slices = append(slices, d)
	}
	return slices
}
