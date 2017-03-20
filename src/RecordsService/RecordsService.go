package main
 
import (
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
var users []User
var medicalRecords []MedicalRecord
var patientHistory []PatientHistory

func main() {
    router := mux.NewRouter()
    
	users = append(users, User{Id: "1", UserName: "ganeshRao", FirstName: "Ganesh", LastName: "Rao", Address: "Mahalakshmi Layout"})
    users = append(users, User{Id: "2", UserName: "kotiCheshte", FirstName: "Rajesh", LastName: "Venkataraman"})

    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	var mr1 MedicalRecord
	mr1 = MedicalRecord{Id: "1", Title: "General Checkup", Symptoms: "Cough and Cold"}
	
	var mr2 MedicalRecord
	mr2 = MedicalRecord{Id: "2", Title: "Bedhi", Symptoms: "Bedhi"}

	var mr3 MedicalRecord
	mr3 = MedicalRecord{Id: "3", Title: "Knee ligament tear", Symptoms: "Knee ligament tear, thumba novu"}
	
	medicalRecords = append(medicalRecords, mr1)
	medicalRecords = append(medicalRecords, mr2)
	medicalRecords = append(medicalRecords, mr3)
	
	router.HandleFunc("/medicalrecords", GetMedicalRecords).Methods("GET")
    router.HandleFunc("/medicalrecords/{id}", GetMedicalRecord).Methods("GET")
    router.HandleFunc("/medicalrecords/{id}", CreateMedicalRecord).Methods("POST")
    router.HandleFunc("/medicalrecords/{id}", DeleteMedicalRecord).Methods("DELETE")

	// Bug below: No two patients should share a medical record. This is for testing the Handlers.
	patientHistory = append(patientHistory, PatientHistory{PatientId: "1", MedicalRecords: medicalRecords})
	patientHistory = append(patientHistory, PatientHistory{PatientId: "2", MedicalRecords: medicalRecords})
	
    router.HandleFunc("/patients/{patientid}/history", GetPatientHistory).Methods("GET")

	// TODO: Implement GetPatientHistoryByMedicalRecordId
	log.Fatal(http.ListenAndServe(":1234", router))
}