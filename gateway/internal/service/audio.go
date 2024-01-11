package service

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type audioService struct {
	portAS string
}

func newAudioService(PortAudio string) *audioService {
	return &audioService{portAS: PortAudio}
}

func (as *audioService) GetAudio(id string) ([]byte, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, LOCALHOST+as.portAS+"/audio/"+id, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("gateway:getAudio request is not 200")
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
