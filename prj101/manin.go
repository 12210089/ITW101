package main

import (
	"net/http"
	patienthandler "prj101/handler/patient_handler"
	userdata "prj101/handler/user_details_handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/register", userdata.Register).Methods("POST")
	r.HandleFunc("/user/login/{email}", userdata.Login).Methods("GET")

	// patient
	r.HandleFunc("/patients/add", patienthandler.AddPatient).Methods("POST")
	r.HandleFunc("/patients/get/{cid}", patienthandler.GetPatientByCid).Methods("GET")
	r.HandleFunc("/patients/update/{cid}", patienthandler.UpdatePatientByCid).Methods("PUT")
	r.HandleFunc("/patients/delete/{cid}", patienthandler.DeletePatientByCid).Methods("DELETE")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/").Handler(fs)
	http.ListenAndServe(":8001", r)

}
