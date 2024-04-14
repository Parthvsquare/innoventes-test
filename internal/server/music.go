package server

import (
	db "innoventes-test/db/sqlc"
	"innoventes-test/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) GetAllMusic(w http.ResponseWriter, r *http.Request) {

	allMusic, err := s.db.GetAllMusic(r.Context())
	if err != nil {
		log.Printf("error getting music %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music")
		return
	}

	data := map[string]interface{}{
		"data": allMusic,
	}
	util.SendSuccessResponse(w, data)
}

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

	newMusic, err := s.db.AddNewMusic(r.Context(), createMusic)
	if err != nil {
		log.Printf("error creating music by %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error creating music")
		return
	}

	data := map[string]interface{}{
		"data": newMusic,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	var updateMusic db.UpdateMusicParams
	err := util.UnmarshalBody(r, &updateMusic)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	updatedMusic, err := s.db.UpdateMusic(r.Context(), updateMusic)
	if err != nil {
		log.Printf("error updating music by %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error updating music")
		return
	}

	data := map[string]interface{}{
		"data": updatedMusic,
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
		log.Printf("error deleting music by music id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error deleting music by music id")
		return
	}

	data := map[string]interface{}{
		"data": "music deleted successfully",
	}
	util.SendSuccessResponse(w, data)
}
