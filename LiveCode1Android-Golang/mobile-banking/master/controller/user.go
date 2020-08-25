package controller

import (
	"banking/master/model"
	"banking/master/usecase"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func UserController(r *mux.Router, service usecase.UserUseCase) {
	UserHandler := UserHandler{service}
	r.HandleFunc("/users", UserHandler.ListUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", UserHandler.UserById).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", UserHandler.DeleteById).Methods(http.MethodDelete)
	r.HandleFunc("/user/{id}", UserHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user", UserHandler.InsertUser).Methods(http.MethodPost)
}

func (s UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	User, err := s.UserUseCase.GetAllDataUser()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		byteOfUsers, err := json.Marshal(User)
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfUsers)
	}

}
func (s UserHandler) UserById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idUser := param["id"]
	User, err := s.UserUseCase.GetUserById(idUser)
	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Id Not Found"))
		}
	} else {
		byteOfUsers, err := json.Marshal(User)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfUsers)
	}
}

func (s UserHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	var User model.UserModel
	param := mux.Vars(r)
	idUser := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&User)
	err := s.UserUseCase.DeleteUserById(idUser)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Deleted"))
	}

}
func (s UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var User model.UserModel
	param := mux.Vars(r)
	idUser := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&User)
	err := s.UserUseCase.UpdateUserById(idUser, User)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Updated"))
	}
}
func (s UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var User model.UserModel
	_ = json.NewDecoder(r.Body).Decode(&User)
	err := s.UserUseCase.InsertUser(User)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Register Failed"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Register Success"))
	}
}
