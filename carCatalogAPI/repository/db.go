package repository

import (
	"catalog/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
// conn = "user=postgres password=root dbname=swaggerDB sslmode=disable"

)

func GetAllCars() []model.Car {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	conn := os.Getenv("connToDB")

	cars := []model.Car{}
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
	}
	query := `select c.reg_num, c.mark, c.model, c.year,p.name,p.surname, p.patronymic from cars c 
	join people p on c.people_id = p.id `

	result, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	for result.Next() {
		car := model.Car{}
		people := model.People{}
		err := result.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year,
			&people.Name, &people.Surname, &people.Patronymic)
		if err != nil {
			log.Println(err)
		}
		car.Owner = &people
		cars = append(cars, car)
	}
	return cars

}

func AddCarsToDB(cars []model.Car) int {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	conn := os.Getenv("connToDB")
	db, err := sql.Open("postgres", conn)
	status := 201

	if err != nil {
		log.Println(err)
	}
	tx, err := db.Begin()
	id := 0
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()

	for _, v := range cars {

		result, err := tx.Prepare("insert into people(name, surname, patronymic) values($1,$2,$3)returning id")
		if err != nil {
			log.Println(err)
		}
		err = result.QueryRow(v.Owner.Name, v.Owner.Surname, v.Owner.Patronymic).Scan(&id)
		if err != nil {
			log.Println(err)
		}

		result, err = tx.Prepare("insert into cars(reg_num, mark, model,year,people_id) values($1,$2,$3,$4,$5)")
		if err != nil {
			log.Println(err)
		}
		_, err = result.Exec(v.RegNum, v.Mark, v.Model, v.Year, id)
		if err != nil {
			log.Println(err)
		}
	}
	err = tx.Commit()

	if err != nil {
		status = http.StatusNotFound
		log.Println(err)

	}

	defer db.Close()

	return status
}

func GetCarOnRegNum(regNum string) model.Car {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	conn := os.Getenv("connToDB")
	car := model.Car{}
	people := model.People{}
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
	}
	query := `select c.reg_num, c.mark, c.model, c.year,p.name,p.surname, p.patronymic from cars c 
	join people p on c.people_id = p.id where c.reg_num = $1`

	result, err := db.Query(query, regNum)
	if err != nil {
		log.Println(err)
	}
	for result.Next() {
		err := result.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year,
			&people.Name, &people.Surname, &people.Patronymic)
		if err != nil {
			log.Println(err)
		}
		car.Owner = &people
	}
	return car
}

func UpdateCar(car model.Car) int {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	conn := os.Getenv("connToDB")
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Println(err)
	}
	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
	}

	defer tx.Rollback()

	updatePeopleQuery := `update people set name=$1, surname=$2, patronymic=$3 
	                      from cars as c where c.people_id = id and c.reg_num = $4`
	updateCarQuery := `update cars set mark=$1, model=$2, year=$3
	                     where  reg_num = $4`

	result, err := tx.Exec(updatePeopleQuery, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic, car.RegNum)
	if err != nil {
		log.Println(err)
	}
	result, err = tx.Exec(updateCarQuery, car.Mark, car.Model, car.Year, car.RegNum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	return http.StatusCreated

}

func Delete(regNum string) int {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}
	conn := os.Getenv("connToDB")
	fmt.Println("in delete db")
	status := 200

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
	}
	queryGetPeopleId := `select people_id from cars where reg_num = $1`
	queryDeleteCars := `delete from cars where reg_num =$1`
	queryDeletePeople := `delete from people where id =$1`
	id := 0

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()

	result, err := tx.Query(queryGetPeopleId, regNum)
	if err != nil {
		log.Println("Delete func: ", err)
	}
	for result.Next() {
		err = result.Scan(&id)
		if err != nil {
			log.Println(err)
		}
	}

	_, err = tx.Exec(queryDeleteCars, regNum)
	if err != nil {
		log.Println(err)
	}
	_, err = tx.Exec(queryDeletePeople, id)
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	if id == 0 {
		status = 404
	}
	defer db.Close()
	return status
}
