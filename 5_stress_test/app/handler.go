package app

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type sendEventHandler struct {
	storage *gorm.DB
}

func (h *sendEventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var m SendEventModel

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		logrus.WithError(err).
			Error("unmarshal request body failed")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := h.storage.WithContext(r.Context()).Create(&m)

	if res.Error != nil {
		logrus.WithError(res.Error).
			Error("create model failed")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.WithField("model", m).
		Debug("model created")

	w.WriteHeader(http.StatusOK)
}

func NewHandler(storage *gorm.DB) *sendEventHandler {
	return &sendEventHandler{
		storage: storage,
	}
}
