package http

import (
	"bytes"
	"gateway/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type audioController struct {
	MA service.MethodsAudio
}

func newAudioController(ma service.MethodsAudio) *audioController {
	return &audioController{MA: ma}
}

func NewAudioRoutes(c *chi.Mux, ma service.MethodsAudio) {

	ctrl := newAudioController(ma)

	//Вызвать лишь единожды. Иначе сущностей будет слишком много. Начинает программу.
	c.Get("/audio/{id:[0-9]+}", ctrl.getAudio)
}

func (b *audioController) getAudio(w http.ResponseWriter, r *http.Request) {
	audioID := chi.URLParam(r, "id")
	body, err := b.MA.GetAudio(audioID)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	audioChunks := bytes.NewBuffer(body)

	if _, err := audioChunks.WriteTo(w); err != nil {
		log.Println(err)
	}
}
