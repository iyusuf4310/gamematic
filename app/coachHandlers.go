package app

import (
	"encoding/json"
	"k/golang/gamematic/dto"
	"k/golang/gamematic/service"
	"net/http"

	"github.com/gorilla/mux"
)

//Coach Handlers
type CoachHandlers struct {
	service service.CoachService
}

func (ch *CoachHandlers) GetCoach(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["coach_id"]
	team, err := ch.service.GetCoach(name)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, team)
	}
}

//Gets all Teams from the repository
func (th *CoachHandlers) GetAllCoaches(w http.ResponseWriter, r *http.Request) {

	teams, err := th.service.GetAllCoaches()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, teams)
	}
}

//Gets all Players from the repository
func (h *CoachHandlers) NewCoach(w http.ResponseWriter, r *http.Request) {
	var request dto.NewCoachRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error)
	} else {
	}
	team, appError := h.service.NewCoach(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
	} else {
		writeResponse(w, http.StatusOK, team)
	}
}
