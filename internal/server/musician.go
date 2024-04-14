package server

import (
	db "innoventes-test/db/sqlc"
	"innoventes-test/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) GetMusicianById(w http.ResponseWriter, r *http.Request) {
	musicianParamsId := chi.URLParam(r, "musicianId")

	musicianUUId, err := uuid.Parse(musicianParamsId)
	if err != nil {
		log.Printf("error parsing musician id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid musician id")
		return
	}

	musician, err := s.db.GetMusiciansById(r.Context(), musicianUUId)
	if err != nil {
		log.Printf("error getting musician by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting musician by musician id")
		return
	}

	data := map[string]interface{}{
		"data": musician,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetMusicByMusicianId(w http.ResponseWriter, r *http.Request) {
	musicianParamsId := chi.URLParam(r, "musicianId")

	musicianUUId, err := uuid.Parse(musicianParamsId)
	if err != nil {
		log.Printf("error parsing music id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid music id")
		return
	}

	allMusic, err := s.db.GetMusicByMusicianId(r.Context(), musicianUUId)
	if err != nil {
		log.Printf("error getting music by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by musician id")
		return
	}

	data := map[string]interface{}{
		"data": allMusic,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetMusicianByAlbumId(w http.ResponseWriter, r *http.Request) {
	albumIdParams := chi.URLParam(r, "albumId")

	albumUUID, err := uuid.Parse(albumIdParams)
	if err != nil {
		log.Printf("error parsing album id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid album id")
		return
	}

	allAlbumsByMusician, err := s.db.GetMusicianByAlbumId(r.Context(), albumUUID)
	if err != nil {
		log.Printf("error getting album by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting album by musician id")
		return
	}

	data := map[string]interface{}{
		"data": allAlbumsByMusician,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) AddNewMusician(w http.ResponseWriter, r *http.Request) {
	var newMusicianParams db.AddNewMusicianParams
	err := util.UnmarshalBody(r, &newMusicianParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	newMusician, err := s.db.AddNewMusician(r.Context(), newMusicianParams)
	if err != nil {
		log.Printf("error creating new musician %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while creating new musician")
		return
	}

	data := map[string]interface{}{
		"data": newMusician,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) UpdateMusicianById(w http.ResponseWriter, r *http.Request) {
	var musicianParams db.UpdateMusicianParams
	err := util.UnmarshalBody(r, &musicianParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	updatedMusician, err := s.db.UpdateMusician(r.Context(), musicianParams)
	if err != nil {
		log.Printf("error while updating musician %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while updating musician")
		return
	}

	data := map[string]interface{}{
		"data": updatedMusician,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) DeleteMusicianById(w http.ResponseWriter, r *http.Request) {
	musicianId := chi.URLParam(r, "musicianId")

	musicianUUID, err := uuid.Parse(musicianId)
	if err != nil {
		log.Printf("error parsing musician id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid musician id")
		return
	}

	err = s.db.DeleteMusician(r.Context(), musicianUUID)

	if err != nil {
		log.Printf("something went wrong while deleting musician %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while deleting musician")
		return
	}

	data := map[string]interface{}{
		"data": "musician deleted successfully",
	}
	util.SendSuccessResponse(w, data)
}
