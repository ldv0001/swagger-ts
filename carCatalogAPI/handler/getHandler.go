package handler

import (
	"catalog/service"
	"encoding/json"
	"log"
	"net/http"
)

// @Summary GetCars
// @Description get cars from db
// @Produce json
// @Param regNum query string false "get car by regNum"
// @Param mark query string false "get car by mark"
// @Param model query string false "get car by model"
// @Param year query integer false "get car by year"
// @Param name query string false "get car by owner's name"
// @Param surname query string false "get car by owner's surname"
// @Param patronymic query integer false "get car by owners patronymic"
// Success 200 {integer} integer
// Failure 404 {integer} integer
// @Router /  [get]
func GetCars(w http.ResponseWriter, r *http.Request) {
	cars, status := service.GetCars(r)
	b, err := json.Marshal(cars)
	if err != nil {
		log.Println(b)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(b)

}
