package main

import (
	"net/http"
	"encoding/json"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
)

func main() {
	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
		Concise: true,
		// Tags: map[string]string{
		// 	"version": "v1.0-81aa4244d9fc8076a",
		// 	"env":     "dev",
		// },
	})

	// Service
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.RealIP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		bs, _ := json.Marshal(r.Header)
		w.Write([]byte(bs))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("oh no")
	})

	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		w.Header().Add("Content-Type", "text/plain")
		oplog.Info().Msg("info here")
		w.Write([]byte("info here"))
	})

	r.Get("/warn", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		oplog.Warn().Msg("warn here")
		w.WriteHeader(400)
		w.Write([]byte("warn here"))
	})

	r.Get("/err", func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		oplog.Error().Msg("err here")
		w.WriteHeader(500)
		w.Write([]byte("err here"))
	})

	r.Get("/azis/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte(`{"status":"ok"}`))
	})
	r.Get("/azis/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte(`pong`))
	})

	http.ListenAndServe(":8080", r)
}
