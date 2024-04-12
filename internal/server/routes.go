package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	// r.Get("/health", s.healthHandler)
	r.Get("/music-album/{id}", s.getMusicAlbum)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
// 	jsonResp, _ := json.Marshal(s.db.Health())
// 	_, _ = w.Write(jsonResp)
// }

func (s *Server) getMusicAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	albumId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid album id", http.StatusBadRequest)
		return
	}

	albums, err := s.db.GetMusicAlbumsById(r.Context(), albumId)

	if err != nil {
		log.Default().Printf("error fetching music albums with params %s", albumId)
		w.WriteHeader(422)
		w.Write([]byte(fmt.Sprintf("error fetching music albums with params %s", albumId)))
		return
	}

	w.Write([]byte(albums.AlbumName))
}
