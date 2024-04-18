package handler

import (
	db "catalog/repository"
	"catalog/service"
	"net/http"
)

// @Summary AddCars
// @Description add cars to db
// @Accept json
// @Produce json
// @Param input body []string false "regNums"
// Success 201 {integer} integer
// Failure 404 {integer} integer
// @Router /add  [post]
func AddCars(w http.ResponseWriter, r *http.Request) {
	regNums := service.AddNewRegNums(r)
	cars := service.GetFromSwaggerAPI(regNums["regNum"])
	status := db.AddCarsToDB(cars)
	w.WriteHeader(status)
}

// @Summary Update
// @Description update cars in db
// @Produce json
// @Param regNum query string true "update car by regNum"
// @Param mark query string false "update car by mark"
// @Param model query string false "update car by model"
// @Param year query integer false "update car by year"
// @Param name query string false "update car by owner's name"
// @Param surname query string false "update car by owner's surname"
// @Param patronymic query integer false "update car by owners patronymic"
// Success 201 {integer} integer
// Failure 404 {integer} integer
// @Router /update  [get]
func Update(w http.ResponseWriter, r *http.Request) {
	regNum := r.FormValue("regNum")
	if regNum == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	status := service.UpdateCar(r)
	w.WriteHeader(status)
}

// @Summary Delete
// @Description delete car from db
// @Produce json
// @Param regNum query string true "delete car by regNum"
// Success 200 {integer} integer
// Failure 404 {integer} integer
// @Router /delete  [get]
func Delete(w http.ResponseWriter, r *http.Request) {
	regNum := r.FormValue("regNum")
	status := service.Delete(regNum)
	w.WriteHeader(status)
}
