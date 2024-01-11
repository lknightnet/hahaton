package http

import (
	"encoding/json"
	"gateway/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type transController struct {
	MT service.MethodTrans
}

func newTransController(mt service.MethodTrans) *transController {
	return &transController{MT: mt}
}

func NewTransRoutes(c *chi.Mux, mt service.MethodTrans) {

	ctrl := newTransController(mt)

	//Вызвать лишь единожды. Иначе сущностей будет слишком много. Начинает программу.
	c.Get("/trans/{id:[0-9]+}", ctrl.getTrans)
}

func (t *transController) getTrans(w http.ResponseWriter, r *http.Request) {
	transID := chi.URLParam(r, "id")

	trans, err := t.MT.GetTrans(transID)
	if err != nil {
		log.Println(err)
	}

	marsh, err := json.Marshal(trans)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(marsh)
	if err != nil {
		log.Println(err)
	}
}
