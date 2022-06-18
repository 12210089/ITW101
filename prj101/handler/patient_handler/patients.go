package patienthandler

import (
	"encoding/json"
	"net/http"
	"prj101/model/patientdomain"
	"prj101/services/patientservice"

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

func AddPatient(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "add courser")
	var ptn patientdomain.Patient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ptn); err != nil {
		// fmt.Fprintf(w, "handle error")
		response, _ := json.Marshal(map[string]string{"error": "invalid json body"})
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}
	defer r.Body.Close()
	saveErr := patientservice.PatientCreate(ptn)
	if saveErr != nil {
		respondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{"status": "Patient added"})
}

// getting patient by id
func GetPatientByCid(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	course, getErr := patientservice.PatientRead(cid)
	if getErr != nil {
		respondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, course)
}

// updating
func UpdatePatientByCid(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	var ptn patientdomain.Patient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ptn); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid jason body")
		return
	}
	defer r.Body.Close()
	ptn.CID = cid
	result, err := patientservice.UpdatePatient(ptn)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}

// deleting
func DeletePatientByCid(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	if err := patientservice.DeletePatient(cid); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})

}
