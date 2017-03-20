package main

import (
    "encoding/json"
    "net/http"
 
    "github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range users {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&User{})
}
 
func GetUsers(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(users)
}
 
func CreateUser(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var user User
    _ = json.NewDecoder(req.Body).Decode(&user)
    user.Id = params["id"]
    users = append(users, user)
    json.NewEncoder(w).Encode(users)
}
 
func DeleteUser(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range users {
        if item.Id == params["id"] {
            users = append(users[:index], users[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(users)
}

func GetMedicalRecords(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(medicalRecords)
}

func GetMedicalRecord(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range medicalRecords {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&MedicalRecord{})
}
 
func CreateMedicalRecord(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var medicalRecord MedicalRecord
    _ = json.NewDecoder(req.Body).Decode(&medicalRecord)
    medicalRecord.Id = params["id"]
    medicalRecords = append(medicalRecords, medicalRecord)
    json.NewEncoder(w).Encode(medicalRecords)
}
 
func DeleteMedicalRecord(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range medicalRecords {
        if item.Id == params["id"] {
            medicalRecords = append(medicalRecords[:index], medicalRecords[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(medicalRecords)
}

func GetPatientHistory(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range patientHistory {
        if item.PatientId == params["patientid"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&PatientHistory{})
}