package main
 
import (
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
var users []User
var medicalRecords []MedicalRecord

func main() {
    router := mux.NewRouter()
    
	users = append(users, User{Id: "1", UserName: "ganeshRao", FirstName: "Ganesh", LastName: "Rao", Address: "Mahalakshmi Layout"})
    users = append(users, User{Id: "2", UserName: "kotiCheshte", FirstName: "Rajesh", LastName: "Venkataraman"})
	
	medicalRecords = append(medicalRecords, MedicalRecord{Id: "1", Title: "General Checkup", Symptoms: "Cough and Cold"})
	medicalRecords = append(medicalRecords, MedicalRecord{Id: "2", Title: "Bedhi", Symptoms: "Bedhi"})
	medicalRecords = append(medicalRecords, MedicalRecord{Id: "3", Title: "Knee ligament tear", Symptoms: "Knee ligament tear, thumba novu"})
		
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	
	router.HandleFunc("/medicalrecords", GetMedicalRecords).Methods("GET")
    router.HandleFunc("/medicalrecords/{id}", GetMedicalRecord).Methods("GET")
    router.HandleFunc("/medicalrecords/{id}", CreateMedicalRecord).Methods("POST")
    router.HandleFunc("/medicalrecords/{id}", DeleteMedicalRecord).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":1234", router))
}