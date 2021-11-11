package app

import (
	"k/golang/gamematic/service"

	"net/http"
)

//Player Handlers
type PlayerHandlers struct {
	service service.PlayerService
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
