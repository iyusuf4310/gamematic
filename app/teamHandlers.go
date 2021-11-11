package app

import (
	"g/go/allsports/service"

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
