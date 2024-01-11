package service

import (
	"bell/internal/model"
	"bell/internal/repository"
	"context"
)

type Bell interface {
	Set(ctx context.Context) error
	GetAll(ctx context.Context) ([]model.BellInfo, error)
	GetByID(ctx context.Context, id string) (*model.BellInfo, error)
	Update(ctx context.Context, bell *model.BellInfo) error
	Remove(ctx context.Context, id string) error
}

type Services struct {
	B Bell
}

type Dependencies struct {
	*repository.Repositories
	URLAPI string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{B: newBellService(deps.URLAPI, deps.Repositories.B)}
}
