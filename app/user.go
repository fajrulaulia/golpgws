package golpgws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type UserService struct {
	Name    string
	Router  *mux.Router
	Context context.Context
	Db      *pgx.Conn
}

func (api *UserService) Register() {
	api.Router.Handle("/api/signup", http.HandlerFunc(api.Signup)).Methods("POST")
	log.Println(api.Name, "registered")
}

func (api *UserService) Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"Email"`
	}
	var resp response

	resp.ID = 1
	resp.Username = "Fajrul Aulia"
	resp.Email = "Auliafajrul7@gmail.com"

	output, err := json.Marshal(resp)
	if err != nil {
		return
	}
	w.Write(output)
}
