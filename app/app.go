package app

import (
	"encoding/json"
	"fmt"
	"k/golang/gamematic/domain"
	"k/golang/gamematic/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" {
		log.Fatal("Environment variable not defined ....")
	}
}

func Start() {

	//sanity check
	sanityCheck()

	//Wiring player
	dbClient := getDBClient()
	playerRepositoryDb := domain.NewPlayerRepositoryDb(dbClient)
	teamRepositoryDb := domain.NewTeamRepositoryDb(dbClient)

	th := TeamHandlers{service.NewTeamService(teamRepositoryDb)}
	ph := PlayerHandlers{service.NewPlayerService(playerRepositoryDb)}

	//Define Muxroutes
	router := mux.NewRouter()

	// Team routes
	router.HandleFunc("/teams", th.getAllTeams).Methods(http.MethodGet)
	router.HandleFunc("/teams/{name}", th.getTeam).Methods(http.MethodGet)

	// Player routes
	router.HandleFunc("/players", ph.getAllPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/newplayer", ph.newPlayer).Methods(http.MethodPost)
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

func getDBClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf(dbUser + ":" + dbPassword + "@tcp(" + dbAddr + ":" + dbPort + ")/" + dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
