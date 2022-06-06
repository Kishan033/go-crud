package userServices

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres/models"
	"go-postgres/repositories/userrepository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ErrInvalidArgument = errors.New("Invalid Argument")
)

type Service interface {
	Add(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type service struct {
	userRepo userrepository.Repository
}

func NewService(userRepo userrepository.Repository) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Add(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	insertID := svc.userRepo.Add(user)

	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func (svc *service) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("User not found")
	}
	user, err := svc.userRepo.Get(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func (svc *service) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := svc.userRepo.List()

	if err != nil {
		log.Fatalf("Unable to get user list. %v", err)
	}

	json.NewEncoder(w).Encode(users)
}

func (svc *service) Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := svc.userRepo.Edit(int64(id), user)

	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func (svc *service) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := svc.userRepo.Delete(int64(id))

	msg := fmt.Sprintf("User deleted successfully. Total rows/record deleted %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
