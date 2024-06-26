package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	// music
	r.Route("/music", func(r chi.Router) {
		r.Get("/{musicId}", s.GetMusicWithId)
		r.Get("/album/{albumId}", s.GetMusicByAlbumId)
		r.Get("/", s.GetAllMusic)
		r.Put("/", s.CreateMusic)
		r.Patch("/", s.UpdateMusic)
		r.Delete("/{musicId}", s.DeleteMusicWithId)
	})

	// music album
	r.Route("/album", func(r chi.Router) {
		r.Get("/{albumId}", s.GetMusicAlbumWithId)
		r.Get("/", s.GetAllAlbums)
		r.Put("/", s.AddNewAlbum)
		r.Patch("/", s.UpdateAlbum)
		r.Delete("/{albumId}", s.DeleteAlbumById)
		r.Get("/musician/{musicianId}", s.GetAlbumsByMusicianId)
		r.Get("/musician/{musicianId}/sort/price", s.GetAlbumsByMusicianId)
		r.Get("/musician/all/{albumId}", s.GetAlbumMusicians)
		r.Post("/musician", s.AddMusicianToAlbum)
		r.Delete("/musician", s.RemoveMusicianFromAlbum)
		r.Patch("/musician", s.UpdateMusicianOfAlbum)
	})

	// musician
	r.Route("/musician", func(r chi.Router) {
		r.Get("/{musicianId}", s.GetMusicianById)
		r.Get("/music/{musicianId}", s.GetMusicByMusicianId)
		r.Get("/album/{musicianId}", s.GetMusicianByAlbumId)
		r.Get("/", s.GetAllMusicians)
		r.Put("/", s.AddNewMusician)
		r.Patch("/", s.UpdateMusicianById)
		r.Delete("/music/{musicId}", s.DeleteMusicianById)
	})

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

// func (s *Server) getMusicAlbum(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")

// 	albumId, err := uuid.Parse(id)
// 	if err != nil {
// 		http.Error(w, "invalid album id", http.StatusBadRequest)
// 		return
// 	}

// 	albums, err := s.db.GetMusicAlbumsById(r.Context(), albumId)

// 	if err != nil {
// 		log.Default().Printf("error fetching music albums with params %s", albumId)
// 		w.WriteHeader(422)
// 		w.Write([]byte(fmt.Sprintf("error fetching music albums with params %s", albumId)))
// 		return
// 	}

// 	w.Write([]byte(albums.AlbumName))
// }
