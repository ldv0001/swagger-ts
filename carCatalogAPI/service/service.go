package service

import (
	"catalog/model"
	db "catalog/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetCars(r *http.Request) ([]model.Car, int) {
	cars := []model.Car{}
	regNum := r.FormValue("regNum")
	if regNum != "" {
		car, status := FilterRegNum(regNum)
		cars = append(cars, car)
		return cars, status
	}
	cars = db.GetAllCars()

	mark := r.FormValue("mark")
	carModel := r.FormValue("model")
	year := r.FormValue("year")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	patronymic := r.FormValue("patronymic")

	if mark != "" {
		cars = FilterCars(cars, mark, func(car model.Car) bool { return car.Mark == mark })
	}
	if year != "" {
		i, err := strconv.Atoi(year)
		if err != nil {
			log.Println(err)
		}
		cars = FilterCars(cars, mark, func(car model.Car) bool { return car.Year == int32(i) })
	}
	if carModel != "" {
		cars = FilterCars(cars, carModel, func(car model.Car) bool { return car.Model == carModel })
	}
	if name != "" {
		cars = FilterCars(cars, name, func(car model.Car) bool { return car.Owner.Name == name })
	}
	if surname != "" {
		cars = FilterCars(cars, surname, func(car model.Car) bool { return car.Owner.Surname == surname })
	}
	if patronymic != "" {
		cars = FilterCars(cars, patronymic, func(car model.Car) bool { return car.Owner.Patronymic == patronymic })
	}
	return cars, http.StatusOK

}

func AddNewRegNums(r *http.Request) map[string][]string {
	b := make([]byte, 1024)
	n, _ := r.Body.Read(b)
	var regNums map[string][]string
	err := json.Unmarshal(b[:n], &regNums)
	if err != nil {
		log.Println(err)
	}
	return regNums
}

func GetFromSwaggerAPI(regNums []string) []model.Car {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	swaggerAPI := os.Getenv("swaggerAPI")

	cars := make([]model.Car, 0, 10)
	for _, v := range regNums {

		car := model.Car{}
		request := fmt.Sprintf(swaggerAPI, v)
		conn, err := http.Get(request)
		if err != nil {
			log.Print(err)
		}

		if conn.StatusCode == 404 {
			continue
		}

		b, err := io.ReadAll(conn.Body)
		if err != nil {
			log.Println(err)
		}

		err = json.Unmarshal(b, &car)
		if err != nil {
			log.Println(err)
		}
		cars = append(cars, car)
		conn.Body.Close()
	}
	return cars
}

func UpdateCar(r *http.Request) int {
	mark := r.FormValue("mark")
	carModel := r.FormValue("model")
	year := r.FormValue("year")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	patronymic := r.FormValue("patronymic")

	car := db.GetCarOnRegNum(r.FormValue("regNum"))
	if mark != "" {
		car.Mark = mark
	}
	if year != "" {
		i, err := strconv.Atoi(year)
		if err != nil {
			log.Println(err)
		}
		car.Year = int32(i)
	}
	if carModel != "" {
		car.Model = carModel
	}
	if name != "" {
		car.Owner.Name = name
	}
	if surname != "" {
		car.Owner.Surname = surname
	}
	if patronymic != "" {
		car.Owner.Patronymic = patronymic
	}
	status := db.UpdateCar(car)
	fmt.Println(status)
	return status

}

func Delete(regNum string) int {
	status := db.Delete(regNum)
	return status
}
