package handlers

import (
	dto "be_waysbucks/dto/result"
	topingdto "be_waysbucks/dto/toping"
	"be_waysbucks/models"
	"encoding/json"
	"fmt"
	"strconv"

	"be_waysbucks/repositories"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

var path_file_toping = os.Getenv("PATH_FILE")

type handlerToping struct {
  TopingRepository repositories.TopingRepository
}

func HandlerToping(TopingRepository repositories.TopingRepository) *handlerToping {
  return &handlerToping{TopingRepository}
}

func (h *handlerToping) FindTopings(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  topings, err := h.TopingRepository.FindTopings()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }
// Create Embed Path File on Image property here ...
  for i, p := range topings {
  topings[i].Image = path_file_toping + p.Image
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: topings}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) GetToping(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  var toping models.Toping
  toping, err := h.TopingRepository.GetToping(id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }
	// Create Embed Path File on Image property here ...
  toping.Image = path_file_toping + toping.Image

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseToping(toping)}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) CreateToping(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // get data user token
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  dataContex := r.Context().Value("dataFile") // add this code
  filename := dataContex.(string) // add this code

  price, _ := strconv.Atoi(r.FormValue("price"))
  request := topingdto.TopingRequest{
    Title:        r.FormValue("title"),
    Price:        price,
    Image:        filename,
  }

  fmt.Println(request)

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  toping := models.Toping{
    Title: request.Title,
    Price: request.Price,
    Image: request.Image,
    UserID: userId,
  }

  toping, err = h.TopingRepository.CreateToping(toping)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  toping, _ = h.TopingRepository.GetToping(toping.ID)

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: toping}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerToping) UpdateToping(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  // get data user token
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  dataContex := r.Context().Value("dataFile") // add this code
  filename := dataContex.(string) // add this code

  price, _ := strconv.Atoi(r.FormValue("price"))
  request := topingdto.TopingRequest{
    Title:        r.FormValue("title"),
    Price:        price,
    Image:        filename,
    UserID:       userId, 
  }

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

	toping, err := h.TopingRepository.GetToping(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
  if request.Title != "" {
    toping.Title = request.Title
  }
  
  if request.Price != 0 {
    toping.Price = request.Price
  }

  if request.Image != "" {
    toping.Image = request.Image
  }

  if request.UserID != 0 {
    toping.UserID = request.UserID
  }

  data, err := h.TopingRepository.UpdateToping(toping,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: data}
  json.NewEncoder(w).Encode(response)

}

func (h *handlerToping) DeleteToping(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  id, _ := strconv.Atoi(mux.Vars(r)["id"])


  toping, err := h.TopingRepository.GetToping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

  deletetoping, err := h.TopingRepository.DeleteToping(toping,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseDeleteToping(deletetoping)}
    json.NewEncoder(w).Encode(response)
}

func convertResponseToping(u models.Toping) topingdto.TopingResponse {
  return topingdto.TopingResponse{
    ID:       u.ID,
    Title:    u.Title,
    Price:    u.Price,
    Image:    u.Image,
  }
}

func convertResponseDeleteToping(u models.Toping) topingdto.TopingResponseDelete {
  return topingdto.TopingResponseDelete{
    ID: u.ID,
  }
}


