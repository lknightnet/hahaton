package repository

import (
	"bell/internal/model"
	"bell/pkg/database/postgres"
	"context"
)

type Bell interface {
	Set(ctx context.Context, bell *model.BellInfo) error
	GetAll(ctx context.Context) ([]model.BellInfo, error)
	GetByID(ctx context.Context, id string) (*model.BellInfo, error)
	Update(ctx context.Context, bell *model.BellInfo) error
	Remove(ctx context.Context, id string) error
}

type Repositories struct {
	B Bell
}


func NewRepositories(db *postgres.Postgres) *Repositories {
	return &Repositories{B: newBellRepository(db)}
}
