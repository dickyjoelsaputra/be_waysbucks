package handlers

import (
	productdto "be_waysbucks/dto/product"
	dto "be_waysbucks/dto/result"
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

var path_file_product = os.Getenv("PATH_FILE")

type handlerProduct struct {
  ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
  return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  products, err := h.ProductRepository.FindProducts()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }
// Create Embed Path File on Image property here ...
  for i, p := range products {
  products[i].Image = path_file_product + p.Image
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: products}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  var product models.Product
  product, err := h.ProductRepository.GetProduct(id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }
	// Create Embed Path File on Image property here ...
  product.Image = path_file_product + product.Image

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseProduct(product)}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // get data user token
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  dataContex := r.Context().Value("dataFile") // add this code
  filename := dataContex.(string) // add this code

  price, _ := strconv.Atoi(r.FormValue("price"))
  request := productdto.ProductRequest{
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

  product := models.Product{
    Title: request.Title,
    Price: request.Price,
    Image: request.Image,
    UserID: userId,
  }

  product, err = h.ProductRepository.CreateProduct(product)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  product, _ = h.ProductRepository.GetProduct(product.ID)

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: product}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")

  // get data user token
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  dataContex := r.Context().Value("dataFile") // add this code
  filename := dataContex.(string) // add this code

  price, _ := strconv.Atoi(r.FormValue("price"))
  request := productdto.ProductRequest{
    Title:        r.FormValue("title"),
    Price:        price,
    Image:        filename,
    UserID:       userId, 
  }

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

	product, err := h.ProductRepository.GetProduct(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
  if request.Title != "" {
    product.Title = request.Title
  }
  
  if request.Price != 0 {
    product.Price = request.Price
  }

  if request.Image != "" {
    product.Image = request.Image
  }

  if request.UserID != 0 {
    product.UserID = request.UserID
  }

  data, err := h.ProductRepository.UpdateProduct(product,id)
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

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  id, _ := strconv.Atoi(mux.Vars(r)["id"])


  product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

  deleteproduct, err := h.ProductRepository.DeleteProduct(product,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseDelete(deleteproduct)}
    json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(u models.Product) productdto.ProductResponse {
  return productdto.ProductResponse{
    ID:       u.ID,
    Title:    u.Title,
    Price:    u.Price,
    Image:    u.Image,
  }
}

func convertResponseDelete(u models.Product) productdto.ProductResponseDelete {
  return productdto.ProductResponseDelete{
    ID: u.ID,
  }
}


