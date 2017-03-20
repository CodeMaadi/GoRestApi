package main
 
import (
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)
 
var users []User
 
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