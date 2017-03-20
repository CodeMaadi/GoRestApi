package main
 
import (
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
var users []User

func main() {
    router := mux.NewRouter()
    users = append(users, User{Id: "1", UserName: "ganeshRao", FirstName: "Ganesh", LastName: "Rao", Address: "Mahalakshmi Layout"})
    users = append(users, User{Id: "2", UserName: "kotiCheshte", FirstName: "Rajesh", LastName: "Venkataraman"})
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":1234", router))
}