package http

import (
	"bell/internal/model"
	"bell/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type bellController struct {
	bell service.Bell
}

func newBellController(bell service.Bell) *bellController {
	return &bellController{bell: bell}
}

//TODO: EDIT ERROR

func NewBellRoutes(r *mux.Router, bell service.Bell) {

	ctrl := newBellController(bell)

	//Вызвать лишь единожды. Иначе сущностей будет слишком много. Начинает программу.
	r.HandleFunc("/set", ctrl.set).Methods(http.MethodGet)

	//Отдает все сущности, сколько бы их там не было.
	r.HandleFunc("/get/all", ctrl.getAll).Methods(http.MethodGet)
	//Отдает только одну сущность, ID, которой было написано.
	r.HandleFunc("/get/{id:[0-9]+}", ctrl.getByID).Methods(http.MethodGet)

	//Обновляет одну сущность, ID, которой было написано.
	r.HandleFunc("/update", ctrl.update).Methods(http.MethodPost)

	//Удаляет одну сущность, ID, которой было написано.
	r.HandleFunc("/remove/{id:[0-9]+}", ctrl.remove).Methods(http.MethodPost)
}

func (b *bellController) set(w http.ResponseWriter, r *http.Request) {
	err := b.bell.Set(r.Context())
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (b *bellController) getAll(w http.ResponseWriter, r *http.Request) {
	bells, err := b.bell.GetAll(r.Context())
	if err != nil {
		log.Println(err)
	}

	marsh, err := json.Marshal(bells)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marsh)
	if err != nil {
		log.Println(err)
	}
}

func (b *bellController) getByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bell, err := b.bell.GetByID(r.Context(), vars["id"])
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

	err = b.bell.Update(r.Context(), &body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (b *bellController) remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := b.bell.Remove(r.Context(), vars["id"])
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
