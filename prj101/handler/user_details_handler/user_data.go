package userdata

import (
	"encoding/json"
	"fmt"
	"prj101/model/userdomain"
	"prj101/services/userservices"

	"net/http"

	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user userdomain.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json Body")
		return
	}
	defer r.Body.Close()
	saveErr := userservices.UserRegCreate(user)
	if saveErr != nil {
		respondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{"status": "user added"})
}

// fetching the data to compare while logging in
func Login(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	fmt.Println(email)
	userDetails, getErr := userservices.ReadUserEmail(email)
	if getErr != nil {
		respondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, userDetails)
}
