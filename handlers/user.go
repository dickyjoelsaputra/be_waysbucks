package handlers

import (
	dto "be_waysbucks/dto/result"
	userdto "be_waysbucks/dto/users"
	"be_waysbucks/models"
	"be_waysbucks/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
  UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
  return &handler{UserRepository}
} 

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  users, err := h.UserRepository.FindUsers()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  // userrespone := userdto.UserResponse{
  //   Name: ,     
  //   Email: users,
  //   Image: users,
  // }
  
var UserResponses []userdto.UserResponse
    for _, s := range users {
        UserResponse := userdto.UserResponse{
            ID:    s.ID,
            Name:  s.Name,
            Email: s.Email,
            Image: s.Image,
        }

        UserResponses = append(UserResponses, UserResponse)
    }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: UserResponses}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  data, err := h.UserRepository.DeleteUser(user,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }


  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

  func convertResponse(u models.User) userdto.UserDeleteResponse {
  return userdto.UserDeleteResponse{
    ID:       u.ID,
  }
  }