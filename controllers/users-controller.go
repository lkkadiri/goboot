package controllers

import (
    "strconv"
    "encoding/json"
    "net/http"
    "fmt"
    "github.com/lkkadiri/goboot/models"
    "github.com/jinzhu/gorm"
    "github.com/gorilla/mux"
)

type (
    // UserController represents the controller for operating on the User resource
    UserController struct{
      dbConn *gorm.DB
    }
)

func NewUserController(db *gorm.DB) *UserController {
    return &UserController{db}
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    u := [] models.User{}
    uc.dbConn.Find(&u)
    uj, _ := json.Marshal(u)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", uj)
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId, err := strconv.Atoi(vars["id"])
    if err != nil {
      fmt.Fprintf(w, "Cannot parse user id")
    }

    u := &models.User{}
    uc.dbConn.Find(u, userId)
    uj, _ := json.Marshal(u)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    u := &models.User{}
    json.NewDecoder(r.Body).Decode(u)
    uc.dbConn.Create(u)
    uj, _ := json.Marshal(u)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId, err := strconv.Atoi(vars["id"])
    if err != nil {
      fmt.Fprintf(w, "Cannot parse user id")
    }
    u := &models.User{}
    uc.dbConn.Delete(u, userId)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
}
