package usecase

import (
	"context"

	"github.com/tigerbig/spatial-data-plateform/internal/domain/collection"
	"github.com/tigerbig/spatial-data-plateform/internal/domain/repository"
)

type SpatialUseCases struct {
	repo repository.SpatialRepositorys
}

func NewSpatialUseCases(repo repository.SpatialRepositorys) *SpatialUseCases {
	return &SpatialUseCases{
		repo: repo,
	}
}

func (useCase *SpatialUseCases) Create(ctx context.Context, feature *collection.Features) (*collection.Features, error) {
	return useCase.repo.Create(ctx, feature)
}

func (useCase *SpatialUseCases) GetAll(ctx context.Context) ([]collection.Features, error) {
	return useCase.repo.FindAll(ctx)
}
