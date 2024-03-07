package controller

import (
	"encoding/json"
	"net/http"

	m "latihan_uts/model"
)

func SendSuccessDetailSongResponse(w http.ResponseWriter, data []m.DetailPlaylistSong) {
	w.Header().Set("Content-Type", "application/json")
	var response m.DetailSongResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = data

	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = code
	response.Message = message

	json.NewEncoder(w).Encode(response)
}
