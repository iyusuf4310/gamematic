package app

import (
	"encoding/json"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/service"
	"net/http"
)

//Team Handlers
type TeamHandlers struct {
	service service.TeamService
}

//Gets all Teams from the repository
func (th *TeamHandlers) getAllTeams(w http.ResponseWriter, r *http.Request) {

	teams, err := th.service.GetAllTeams()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, teams)
	}
}

//Gets all Players from the repository
func (h *TeamHandlers) newTeam(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTeamRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
	}
	team, appError := h.service.NewTeam(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
	} else {
		writeResponse(w, http.StatusOK, team)
	}
}
