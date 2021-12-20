package app

import (
	"encoding/json"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Team Handlers
type TeamHandlers struct {
	service service.TeamService
}

func (ch *TeamHandlers) getTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	team, err := ch.service.GetTeam(name)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, team)
	}
}

func (ch *TeamHandlers) deleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["team_id"]

	teamId, _ := strconv.Atoi(id)

	err := ch.service.DeleteTeam(teamId)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, "Successfully Deleted team: "+id)
	}
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

//Create new team
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

//Update Team
func (h *TeamHandlers) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	var request dto.NewTeamRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
	}
	team, appError := h.service.UpdateTeam(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
	} else {
		writeResponse(w, http.StatusOK, team)
	}

}