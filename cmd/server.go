package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/saihnii4/utdr-video-creator/v2/pkg/utdr/api"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("you're lost aren't ya?\n"))
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK\n"))
	})
	r.Route("/textbox", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("info or smth\n"))
		})
		r.Get("/generate", api.GenerateTextBox)
		r.Get("/generate_anim", api.GenerateAnimatedTextBox)
	})

	slog.Info("server running")
	err := http.ListenAndServe(":4729", r)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
