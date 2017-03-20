package main
 
import (
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)

type User struct {
    Id string `json:"Id,omitempty"`
	UserName string `json:"UserName,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	MiddleName string `json:"MiddleName,omitempty"`
	LastName string `json:"LastName,omitempty"`
	UserImage string `json:"UserImage,omitempty"`
	UserLink string `json:"UserLink,omitempty"`
	Address string `json:"Address,omitempty"`
	City string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
	ZipCode string `json:"ZipCode,omitempty"`
	Country string `json:"Country,omitempty"`
	AadharID string `json:"AadharId,omitempty"`
	PanCardNumber string `json:"PanCardNumber,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Email string `json:"Email,omitempty"`
	Role string `json:"Role,omitempty"`
}
 
var users []User
 
func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range users {
        if item.Id == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&User{})
}
 
func GetUsersEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(users)
}
 
func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var user User
    _ = json.NewDecoder(req.Body).Decode(&user)
    user.Id = params["id"]
    users = append(users, user)
    json.NewEncoder(w).Encode(users)
}
 
func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
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
    router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
    router.HandleFunc("/users/{id}", GetUserEndpoint).Methods("GET")
    router.HandleFunc("/users/{id}", CreateUserEndpoint).Methods("POST")
    router.HandleFunc("/users/{id}", DeleteUserEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":1234", router))
}