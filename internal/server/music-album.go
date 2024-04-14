package server

import (
	db "innoventes-test/db/sqlc"
	"innoventes-test/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	allAlbums, err := s.db.GetAllMusicAlbums(r.Context())
	if err != nil {
		log.Printf("error getting albums by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting albums by musician id")
		return
	}

	data := map[string]interface{}{
		"data": allAlbums,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetAlbumsByMusicianId(w http.ResponseWriter, r *http.Request) {
	musicianParamsId := chi.URLParam(r, "musicianId")

	musicianUUId, err := uuid.Parse(musicianParamsId)
	if err != nil {
		log.Printf("error parsing musician id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid musician id")
		return
	}

	allAlbums, err := s.db.GetAlbumsByMusicianId(r.Context(), musicianUUId)
	if err != nil {
		log.Printf("error getting albums by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting albums by musician id")
		return
	}

	data := map[string]interface{}{
		"data": allAlbums,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetAlbumsByMusicianIdSorted(w http.ResponseWriter, r *http.Request) {
	musicianParamsId := chi.URLParam(r, "musicianId")

	musicianUUId, err := uuid.Parse(musicianParamsId)
	if err != nil {
		log.Printf("error parsing musician id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid musician id")
		return
	}

	allAlbums, err := s.db.GetAlbumsByMusicianIdSorted(r.Context(), musicianUUId)
	if err != nil {
		log.Printf("error getting albums by musician id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting albums by musician id")
		return
	}

	data := map[string]interface{}{
		"data": allAlbums,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetAlbumMusicians(w http.ResponseWriter, r *http.Request) {
	albumIdParams := chi.URLParam(r, "albumId")

	albumUUID, err := uuid.Parse(albumIdParams)
	if err != nil {
		log.Printf("error parsing album id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid album id")
		return
	}

	allAlbumMusicians, err := s.db.GetMusicianForAlbum(r.Context(), albumUUID)
	if err != nil {
		log.Printf("error getting album musicians %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting album musicians")
		return
	}

	data := map[string]interface{}{
		"data": allAlbumMusicians,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) GetMusicAlbumWithId(w http.ResponseWriter, r *http.Request) {
	albumIdParams := chi.URLParam(r, "albumId")

	albumUUID, err := uuid.Parse(albumIdParams)
	if err != nil {
		log.Printf("error parsing album id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid album id")
		return
	}

	albumByAlbumId, err := s.db.GetAlbumByAlbumId(r.Context(), albumUUID)
	if err != nil {
		log.Printf("error getting album by album id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting album by album id")
		return
	}

	data := map[string]interface{}{
		"data": albumByAlbumId,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) AddNewAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbumParams db.AddNewAlbumParams
	err := util.UnmarshalBody(r, &newAlbumParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	newAlbum, err := s.db.AddNewAlbum(r.Context(), newAlbumParams)
	if err != nil {
		log.Printf("error creating new album %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while creating new album")
		return
	}

	data := map[string]interface{}{
		"data": newAlbum,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var albumParams db.UpdateAlbumParams
	err := util.UnmarshalBody(r, &albumParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	udpatedAlbum, err := s.db.UpdateAlbum(r.Context(), albumParams)
	if err != nil {
		log.Printf("error while updating album %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while updating album")
		return
	}

	data := map[string]interface{}{
		"data": udpatedAlbum,
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) DeleteAlbumById(w http.ResponseWriter, r *http.Request) {
	albumId := chi.URLParam(r, "albumId")

	albumUUID, err := uuid.Parse(albumId)
	if err != nil {
		log.Printf("error parsing album id %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, "invalid album id")
		return
	}

	err = s.db.DeleteAlbum(r.Context(), albumUUID)
	if err != nil {
		log.Printf("error getting music by music id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error getting music by music id")
		return
	}

	data := map[string]interface{}{
		"data": "album deleted successfully",
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) AddMusicianToAlbum(w http.ResponseWriter, r *http.Request) {
	var albumParams db.AddMusicianToAlbumParams
	err := util.UnmarshalBody(r, &albumParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	err = s.db.AddMusicianToAlbum(r.Context(), albumParams)
	if err != nil {
		log.Printf("error while adding musicians to album %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while adding musicians to album")
		return
	}

	data := map[string]interface{}{
		"data": "Added musician to album successfully",
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) RemoveMusicianFromAlbum(w http.ResponseWriter, r *http.Request) {
	var removeMusicianToAlbumParams db.DeleteMusicianFromAlbumParams
	err := util.UnmarshalBody(r, &removeMusicianToAlbumParams)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	err = s.db.DeleteMusicianFromAlbum(r.Context(), removeMusicianToAlbumParams)
	if err != nil {
		log.Printf("error getting music by music id %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error removing album musician from album")
		return
	}

	data := map[string]interface{}{
		"data": "removed musician from album successfully",
	}
	util.SendSuccessResponse(w, data)
}

func (s *Server) UpdateMusicianOfAlbum(w http.ResponseWriter, r *http.Request) {
	var updateMusicianOfAlbum db.UpdateMusicianOfAlbumParams
	err := util.UnmarshalBody(r, &updateMusicianOfAlbum)

	if err != nil {
		log.Printf("error un marshalling request body %v", err)
		util.SendErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	err = s.db.UpdateMusicianOfAlbum(r.Context(), updateMusicianOfAlbum)
	if err != nil {
		log.Printf("error while updating musicians to album %v", err)
		util.SendErrorResponse(w, http.StatusInternalServerError, "error while updating musicians to album")
		return
	}

	data := map[string]interface{}{
		"data": "updated musician to album successfully",
	}
	util.SendSuccessResponse(w, data)
}
