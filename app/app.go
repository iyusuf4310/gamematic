package app

import (
	"encoding/json"
	"fmt"
	"k/golang/gamematic/domain"
	"k/golang/gamematic/logger"
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

	//DB Client
	dbClient := getDBClient()
	//Wiring player
	playerRepositoryDb := domain.NewPlayerRepositoryDb(dbClient)
	ph := PlayerHandlers{service.NewPlayerService(playerRepositoryDb)}
	//Wiring team
	teamRepositoryDb := domain.NewTeamRepositoryDb(dbClient)
	th := TeamHandlers{service.NewTeamService(teamRepositoryDb)}
	//Wiring Coach
	coachRepositoryDb := domain.NewCoachRepositoryDB(dbClient)
	ch := CoachHandlers{service.NewCoachService(coachRepositoryDb)}

	//Define Muxroutes
	router := mux.NewRouter()

	// Team routes
	router.HandleFunc("/teams", th.getAllTeams).Methods(http.MethodGet)
	router.HandleFunc("/teams/{name}", th.getTeam).Methods(http.MethodGet)
	router.HandleFunc("/teams/newteam", th.newTeam).Methods(http.MethodPost)
	router.HandleFunc("/teams/{team_id:[0-9]+}", th.deleteTeam).Methods(http.MethodDelete)
	router.HandleFunc("/teams/updateteam", th.newTeam).Methods(http.MethodPut)

	// Player routes
	router.HandleFunc("/players", ph.getAllPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/newplayer", ph.newPlayer).Methods(http.MethodPost)
	router.HandleFunc("/players/{player_id:[0-9]+}", ph.getPlayer).Methods(http.MethodGet)

	//Coach routes
	router.HandleFunc("/coaches", ch.GetAllCoaches).Methods(http.MethodGet)
	router.HandleFunc("/coaches/newcoach", ch.NewCoach).Methods(http.MethodPost)
	router.HandleFunc("/coaches/{coach_id:[0-9]+}", ch.GetCoach).Methods(http.MethodGet)

	//Start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting Gamematic server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

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
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
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
