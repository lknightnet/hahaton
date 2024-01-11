package service

import (
	"bell/internal/model"
	"bell/internal/repository"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type bellService struct {
	urlAPI string
	db     repository.Bell
}

func (b *bellService) Set(ctx context.Context) error {
	for i := 0; i < 6; i++ {
		resp, err := http.Get(b.urlAPI)
		if err != nil {
			return errors.Wrap(err, "Bell/Set: fail to GET request to API server")
		}

		bellInfo := &model.BellInfo{}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Bell/Set: fail to read response body")
		}

		err = json.Unmarshal(body, &bellInfo)
		if err != nil {
			return errors.Wrap(err, "Bell/Set: fail to json unmarshal BellInfo")
		}

		err = b.db.Set(ctx, bellInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *bellService) GetAll(ctx context.Context) ([]model.BellInfo, error) {
	return b.db.GetAll(ctx)
}

func (b *bellService) GetByID(ctx context.Context, id string) (*model.BellInfo, error) {
	return b.db.GetByID(ctx, id)
}

func (b *bellService) Update(ctx context.Context, bell *model.BellInfo) error {
	return b.db.Update(ctx, bell)
}

func (b *bellService) Remove(ctx context.Context, id string) error {
	return b.db.Remove(ctx, id)
}

func newBellService(urlAPI string, db repository.Bell) *bellService {
	return &bellService{urlAPI: urlAPI, db: db}
}
