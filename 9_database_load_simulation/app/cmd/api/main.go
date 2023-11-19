package main

import (
	"net/http"
	app "prjctr/md/9_database_load_simulation"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("Starting...")

	storage := app.NewStorage()
	err := storage.AutoMigrate(&app.User{})
	if err != nil {
		logrus.WithError(err).
			Error("migrate failed")

		return
	}

	logrus.Debug("Migrated...")

	handler := app.NewUserHandler(storage)

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/users/rnd", handler.CreateRandom).Methods(http.MethodPost)
	r.HandleFunc("/users", handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/status", handler.Status).Methods(http.MethodGet)

	port := "80"
	logrus.Debug("Serving on port ", port)

	logrus.Fatal(http.ListenAndServe(":" + port, r))
}
