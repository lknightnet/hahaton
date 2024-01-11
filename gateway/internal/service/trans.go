package service

import (
	"encoding/json"
	"gateway/internal/model"
	"io"
	"log"
	"net/http"
	"time"
)

type transService struct {
	portTrans string
}

func newTransService(portTrans string) *transService {
	return &transService{portTrans: portTrans}
}

func (t *transService) GetTrans(id string) (*model.Chat, error) {
	client := &http.Client{Timeout: time.Hour}

	req, err := http.NewRequest(http.MethodGet, LOCALHOST+t.portTrans+"/transcription/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	trans := &model.Chat{
		Messages: make([]model.Message, 0),
	}

	err = json.Unmarshal(body, &trans)
	if err != nil {
		return nil, err
	}

	log.Println(trans)
	return trans, nil
}
