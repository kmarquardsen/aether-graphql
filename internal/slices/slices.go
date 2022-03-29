package slices

import (
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_0_0/types"
)

func List(site types.EnterprisesEnterpriseSite) []*model.Slice {
	slices := []*model.Slice{}
	for _, slice := range *site.Slice {
		d := &model.Slice{
			ID:   slice.SliceId,
			Name: *slice.DisplayName,
		}
		slices = append(slices, d)
	}
	return slices
}
