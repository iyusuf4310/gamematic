package app

import (
	"encoding/json"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/service"
	"net/http"

	"github.com/gorilla/mux"
)

//Player Handlers
type PlayerHandlers struct {
	service service.PlayerService
}

func (ch *PlayerHandlers) getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["player_id"]
	player, err := ch.service.GetPlayer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, player)
	}
}

//Gets all Players from the repository
func (ph *PlayerHandlers) getAllPlayers(w http.ResponseWriter, r *http.Request) {

	players, err := ph.service.GetAllPlayers()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, players)
	}

}

//Gets all Players from the repository
func (h *PlayerHandlers) newPlayer(w http.ResponseWriter, r *http.Request) {
	var request dto.NewPlayerRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
	}
	player, appError := h.service.NewPlayer(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
	} else {
		writeResponse(w, http.StatusOK, player)
	}

}
