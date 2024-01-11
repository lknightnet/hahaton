package service

import (
	"bytes"
	"encoding/json"
	"gateway/internal/model"
)

const (
	LOCALHOST = "http://172.17.0.1"
)

type MethodsBell interface {
	Set() error
	GetAll() ([]model.BellInfo, error)
	GetByID(id string) (*model.BellInfo, error)
	Update(bell *model.BellInfo) error
	Remove(id string) error
}

type MethodTrans interface {
	GetTrans(id string) (*model.Chat, error)
}

type MethodsAudio interface {
	GetAudio(id string) ([]byte, error)
}

func convert(body any) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

type Services struct {
	MB MethodsBell
	MA MethodsAudio
	MT MethodTrans
}

type DependenciesServices struct {
	PortBell  string
	PortAudio string
	PortTrans string
}

func NewServices(deps *DependenciesServices) *Services {
	return &Services{
		MB: newBellService(deps.PortBell),
		MA: newAudioService(deps.PortAudio),
		MT: newTransService(deps.PortTrans),
	}
}
