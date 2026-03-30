package repository

import (
	"context"

	"github.com/tigerbig/spatial-data-plateform/internal/domain/collection"
)

type SpatialRepositorys interface {
	Create(ctx context.Context, feature *collection.Features) (*collection.Features, error)
	FindAll(ctx context.Context) ([]collection.Features, error)
}
