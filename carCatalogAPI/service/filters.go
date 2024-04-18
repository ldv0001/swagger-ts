package service

import (
	"catalog/model"
	db "catalog/repository"
	"net/http"
)

type filter func(car model.Car) bool

func FilterRegNum(regNum string) (model.Car, int) {
	car := db.GetCarOnRegNum(regNum)
	if car.RegNum == "" {
		return car, http.StatusNotFound
	}
	return car, http.StatusOK
}

func FilterCars(cars []model.Car, parameter string, filter filter) []model.Car {
	nCars := make([]model.Car, 0, len(cars))
	for _, v := range cars {
		if filter(v) {
			nCars = append(nCars, v)
		}
	}
	return nCars
}
