package handlers

import (
	orderdto "be_waysbucks/dto/order"
	dto "be_waysbucks/dto/result"
	"be_waysbucks/models"
	"be_waysbucks/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerOrder struct {
  OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
  return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // get data user token
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))
  
  request := new(orderdto.OrderRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  uuser, err := h.OrderRepository.GetUserOrder(userId)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  uuuser := models.UserResponse{
    ID: userId,
    Name: uuser.Name,
    Email: uuser.Email,
  }

  // fmt.Println("user id", userId)
  // fmt.Println("product id",request.ProductId)
  // fmt.Println("quantiti",request.Quantity)
  // fmt.Print("toping id",request.TopingId)
  // done

  product, err := h.OrderRepository.GetProductOrder(request.ProductId)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  var topingid []int
  for _, v := range request.TopingId  {
    topingid = append(topingid, int(v))
  }

  // fmt.Print("product id",product , "topingid" , t)
  // t = [1,2]
  // done 

  topings , err := h.OrderRepository.GetTopingOrder(topingid)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  // fmt.Println("topings order ", topings)
  // done
  
  var topingharga int
    for _, v := range request.TopingId {
    topings , err := h.OrderRepository.GetToping(v)
    if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
    }
    topingharga += topings.Price
  }
  TotalHarga := (topingharga + product.Price) * request.Quantity

  // fmt.Println(TotalHarga) Total Harga

  Order := models.Order{
    Qty: request.Quantity,
    UserID: userId,
    User: uuuser,
    TotalPrice: TotalHarga,
    Product: product,
    TopingID: topingid,
    Toping: topings,
    ProductID: product.ID,
  }


  orders , err := h.OrderRepository.CreateOrder(Order)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  fmt.Println(orders)

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Status: http.StatusOK, Data: Order}
  json.NewEncoder(w).Encode(response)
}

// DELETE ORDER ==================================================================

func (h *handlerOrder) DeleteOrder(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
  id, _ := strconv.Atoi(mux.Vars(r)["id"])


  order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

  deleteorder, err := h.OrderRepository.DeleteOrder(order,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Status: http.StatusOK, Data: convertResponseDeleteOrder(deleteorder)}
    json.NewEncoder(w).Encode(response)
}

// GET ORDER ===============================================================================

func (h *handlerOrder) GetOrder(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  
  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  fmt.Println(userId)

  order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

  // produkid := order.Product
  // fmt.Println(order)

  // product, err := h.OrderRepository.GetProductOrder(produkid.ID)
  // if err != nil {
  //   w.WriteHeader(http.StatusInternalServerError)
  //   response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
  //   json.NewEncoder(w).Encode(response)
  //   return
  // }


  // topings , err := h.OrderRepository.GetTopingOrder(order.TopingID)
  // if err != nil {
  //   w.WriteHeader(http.StatusInternalServerError)
  //   response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
  //   json.NewEncoder(w).Encode(response)
  //   return
  // }

  //   Order := models.Order{
  //     ID: order.ID,
  //     UserID: order.UserID,
  //     User: order.User,
  //     Qty: order.Qty,
  //     TotalPrice: order.TotalPrice,
  //     ProductID: product.ID,
  //     Product: product,
  //     // TopingID: order.TopingID,
  //     // Toping: topings,
  // }


    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Status: http.StatusOK, Data: order}
    json.NewEncoder(w).Encode(response)

}

func (h *handlerOrder) MyTransaction(w http.ResponseWriter, r *http.Request) {
  userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
  userId := int(userInfo["id"].(float64))

  fmt.Println(userId)

    order, err := h.OrderRepository.GetOrderByUserID(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Status: http.StatusOK, Data: order}
    json.NewEncoder(w).Encode(response)

}




func convertResponseDeleteOrder(u models.Order) orderdto.OrderResponseDelete {
  return orderdto.OrderResponseDelete{
    ID: u.ID,
  }
}

