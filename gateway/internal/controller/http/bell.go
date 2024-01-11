package http

import (
	"encoding/json"
	"gateway/internal/model"
	"gateway/internal/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type bellController struct {
	M service.MethodsBell
}

func newBellController(m service.MethodsBell) *bellController {
	return &bellController{M: m}
}

func NewBellRoutes(c *chi.Mux, m service.MethodsBell) {

	ctrl := newBellController(m)

	//Вызвать лишь единожды. Иначе сущностей будет слишком много. Начинает программу.
	c.Get("/set", ctrl.set)

	//Отдает все сущности, сколько бы их там не было.
	c.Get("/get/all", ctrl.getAll)
	//Отдает только одну сущность, ID, которой было написано.
	c.Get("/get/{id:[0-9]+}", ctrl.getByID)

	c.Post("/update", ctrl.update)

	//Удаляет одну сущность, ID, которой было написано.
	c.Post("/remove/{id:[0-9]+}", ctrl.remove)
}

func (b *bellController) set(w http.ResponseWriter, r *http.Request) {
	err := b.M.Set()
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (b *bellController) getAll(w http.ResponseWriter, r *http.Request) {
	bells, err := b.M.GetAll()
	if err != nil {
		log.Println(err)
	}

	marsh, err := json.Marshal(bells)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(marsh)
	if err != nil {
		log.Println(err)
	}
}

func (b *bellController) getByID(w http.ResponseWriter, r *http.Request) {
	bellID := chi.URLParam(r, "id")

	bell, err := b.M.GetByID(bellID)
	if err != nil {
		log.Println(err)
	}

	marsh, err := json.Marshal(bell)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marsh)
	if err != nil {
		log.Println(err)
	}
}

func (b *bellController) update(w http.ResponseWriter, r *http.Request) {
	var body model.BellInfo

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
	}

	err = b.M.Update(&body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (b *bellController) remove(w http.ResponseWriter, r *http.Request) {
	bellID := chi.URLParam(r, "id")
	err := b.M.Remove(bellID)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
