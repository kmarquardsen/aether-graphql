package smallcells

import (
	"context"
	"github.com/onosproject/aether-graphql/graph/model"
	"github.com/onosproject/aether-roc-api/pkg/aether_2_1_0/types"
)

func List(ctx context.Context, site types.EnterprisesEnterpriseSite) []*model.SmallCell {
	var smallCells []*model.SmallCell
	for _, sc := range *site.SmallCell {
		d := &model.SmallCell{
			ID:   sc.SmallCellId,
			Name: *sc.DisplayName,
		}
		smallCells = append(smallCells, d)
	}
	return smallCells
}
