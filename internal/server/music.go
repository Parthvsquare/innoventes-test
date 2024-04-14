package server

import (
	db "innoventes-test/db/sqlc"
	"innoventes-test/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) GetMusicWithId(w http.ResponseWriter, r *http.Request) {
	musicParamsId := chi.URLParam(r, "musicId")

	musicUUId, err := uuid.Parse(musicParamsId)
	if err != nil {
		log.Printf("error parsing music id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid music id")
		return
	}

	musicWithGivenId, err := s.db.GetMusicById(r.Context(), musicUUId)
	if err != nil {
		log.Printf("error getting music by music id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by music id")
		return
	}

	data := map[string]interface{}{
		"data": musicWithGivenId,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetMusicByAlbumId(w http.ResponseWriter, r *http.Request) {
	albumParamsId := chi.URLParam(r, "albumId")

	albumUUId, err := uuid.Parse(albumParamsId)
	if err != nil {
		log.Printf("error parsing album id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid album id")
		return
	}

	musicWithGivenAlbumId, err := s.db.GetMusicByAlbumId(r.Context(), albumUUId)
	if err != nil {
		log.Printf("error getting music by album id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by album id")
		return
	}

	data := map[string]interface{}{
		"data": musicWithGivenAlbumId,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) CreateMusic(w http.ResponseWriter, r *http.Request) {
	var createMusic db.AddNewMusicParams
	err := util.UnmarshalBody(r, &createMusic)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	musicWithGivenAlbumId, err := s.db.AddNewMusic(r.Context(), createMusic)
	if err != nil {
		log.Printf("error getting music by album id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by album id")
		return
	}

	data := map[string]interface{}{
		"data": musicWithGivenAlbumId,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) DeleteMusicWithId(w http.ResponseWriter, r *http.Request) {
	musicParamsId := chi.URLParam(r, "musicId")

	musicUUId, err := uuid.Parse(musicParamsId)
	if err != nil {
		log.Printf("error parsing music id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid music id")
		return
	}

	err = s.db.DeleteMusic(r.Context(), musicUUId)
	if err != nil {
		log.Printf("error getting music by music id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by music id")
		return
	}

	data := map[string]interface{}{
		"data": "music deleted successfully",
	}
	util.SendSuccessResponse(w, data)
}
