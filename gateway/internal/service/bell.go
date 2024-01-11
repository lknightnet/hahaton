package service

import (
	"encoding/json"
	"errors"
	"gateway/internal/model"
	"io"
	"net/http"
)

type bellService struct {
	portBell string
}

func (b *bellService) Set() error {
	client := &http.Client{
		Timeout: 0,
	}

	req, err := http.NewRequest(http.MethodGet, LOCALHOST+b.portBell+"/set", nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("gateway:set request is not 200")
	}
	return nil
}

func (b *bellService) GetAll() ([]model.BellInfo, error) {
	client := &http.Client{
		Timeout: 0,
	}

	req, err := http.NewRequest(http.MethodGet, LOCALHOST+b.portBell+"/get/all", nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("gateway:getAll request is not 200")
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bells := make([]model.BellInfo, 0)

	err = json.Unmarshal(respBody, &bells)
	if err != nil {
		return nil, err
	}

	return bells, nil
}

func (b *bellService) GetByID(id string) (*model.BellInfo, error) {
	client := &http.Client{
		Timeout: 0,
	}

	req, err := http.NewRequest(http.MethodGet, LOCALHOST+b.portBell+"/get/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("gateway:getByID request is not 200")
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bell := &model.BellInfo{}

	err = json.Unmarshal(respBody, &bell)
	if err != nil {
		return nil, err
	}

	return bell, nil
}

func (b *bellService) Update(bell *model.BellInfo) error {
	client := &http.Client{
		Timeout: 0,
	}

	buf, err := convert(bell)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, LOCALHOST+b.portBell+"/update", buf)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("gateway:update request is not 200")
	}

	return nil
}

func (b *bellService) Remove(id string) error {
	client := &http.Client{
		Timeout: 0,
	}

	req, err := http.NewRequest(http.MethodPost, LOCALHOST+b.portBell+"/remove/"+id, nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("gateway:remove request is not 200")
	}

	return nil
}

var _ MethodsBell = &bellService{}

func newBellService(portBell string) *bellService {
	return &bellService{portBell: portBell}
}
