package routes

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/antonioofdz/hmisProject/pkg/handlers"
	"github.com/antonioofdz/hmisProject/pkg/models"
)

func LoadRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/login", loginUserController).Methods("POST")
	router.Handle("/user", handlers.CheckToken(http.HandlerFunc(getUserByTokenController))).Methods("GET")

	log.Fatal(http.ListenAndServe(":5002", router))
}

// Controlador para logear a un usuario
func loginUserController(w http.ResponseWriter, req *http.Request) {
	var userDB *models.UserDBCredentials
	if err := json.NewDecoder(req.Body).Decode(&userDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error Parsing BODY! [/user/login]"))
	}

	data, err := handlers.GetUserCredentials(userDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Something bad happened! [/user/login]"))
	}
	json.NewEncoder(w).Encode(data)
}

// Controlador que obtiene un usuario por su Token
func getUserByTokenController(w http.ResponseWriter, req *http.Request) {
	data, err := handlers.GetUserUserByToken(req.Header.Get("token"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Something bad happened! [/user]"))
	}

	json.NewEncoder(w).Encode(data)
}