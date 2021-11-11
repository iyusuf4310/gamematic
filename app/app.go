package app

import (
	"encoding/json"
	"g/go/allsports/domain"
	"g/go/allsports/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//Wiring player
	th := TeamHandlers{service.NewTeamService(domain.NewTeamRepositoryDb())}
	ph := PlayerHandlers{service.NewPlayerService(domain.NewPlayerRepositoryDb())}

	//Define Muxroutes
	router := mux.NewRouter()

	// Team routes
	router.HandleFunc("/teams", th.getAllTeams).Methods(http.MethodGet)
	router.HandleFunc("/teams/{name}", th.getTeam).Methods(http.MethodGet)

	// Player routes
	router.HandleFunc("/players", ph.getAllPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/{player_id:[0-9]+}", ph.getPlayer).Methods(http.MethodGet)

	//Start server
	log.Fatal(http.ListenAndServe("localhost:8013", router))
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

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
