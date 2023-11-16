package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userHandler struct {
	storage *gorm.DB
}

func (h *userHandler) Status(w http.ResponseWriter, r *http.Request) {
	var count int64

	res := h.storage.WithContext(r.Context()).Model(User{}).Count(&count)
	if err := res.Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.WithError(res.Error).
				Error("record not found")

			w.WriteHeader(http.StatusNotFound)
			return
		default:
			logrus.WithError(res.Error).
				Error("get users failed")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"count": count,
	})
	if err != nil {
		logrus.WithError(err).
			Error("marshal response body failed")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	// get date of birth from query params
	dateOfBirthRaw := r.URL.Query().Get("bd")
	dateOfBirth, err := time.Parse(time.DateOnly, dateOfBirthRaw)
	if err != nil {
		logrus.WithError(err).
			WithField("bd", dateOfBirthRaw).
			Error("parse date of birth failed")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get users from storage
	var users []User
	res := h.storage.WithContext(r.Context()).Find(&users, "date_of_birth = ?", dateOfBirth)
	if err := res.Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logrus.WithError(res.Error).
				Error("record not found")

			w.WriteHeader(http.StatusNotFound)
			return
		default:
			logrus.WithError(res.Error).
				Error("get users failed")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	logrus.WithField("count", res.RowsAffected).
		Info("users found")

	// return users
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		logrus.WithError(err).
			Error("marshal response body failed")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		logrus.WithError(err).
			Error("unmarshal request body failed")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := h.storage.WithContext(r.Context()).Create(&u)

	if res.Error != nil {
		logrus.WithError(res.Error).
			Error("create model failed")

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.WithField("user", u).
		Debug("model created")

	w.WriteHeader(http.StatusOK)
}

func NewUserHandler(storage *gorm.DB) *userHandler {
	return &userHandler{
		storage: storage,
	}
}
