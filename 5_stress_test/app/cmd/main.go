package main

import (
	"net/http"
	app "prjctr/md/5_stress_test"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	storage := app.NewStorage()
	storage.AutoMigrate(&app.SendEventModel{})

	handler := app.NewHandler(storage)

	r := mux.NewRouter()
	r.HandleFunc("/events", handler.Handle).Methods(http.MethodPost)

	logrus.Fatal(http.ListenAndServe(":80", r))
}