package swagger

func getCar() map[string]Car {
	cars := make(map[string]Car)
	cars["X123XX150"] = Car{RegNum: "X123XX150",
		Mark:  "Lada",
		Model: "Vesta",
		Year:  2002,
		Owner: &People{Name: "Ivan", Surname: "Ivanov", Patronymic: "Ivanovich"}}

	cars["X123XX151"] = Car{RegNum: "X123XX151",
		Mark:  "Lada",
		Model: "Granta",
		Year:  2011,
		Owner: &People{Name: "Ivan", Surname: "Petrov", Patronymic: "Petrovich"}}

	cars["X123XX152"] = Car{RegNum: "X123XX152",
		Mark:  "Lada",
		Model: "Largus",
		Year:  2012,
		Owner: &People{Name: "Petr", Surname: "Ivanov", Patronymic: "Ivanovich"}}
	cars["X123XX153"] = Car{RegNum: "X123XX153",
		Mark:  "GAZ",
		Model: "66",
		Year:  1970,
		Owner: &People{Name: "Platon", Surname: "Ivanov", Patronymic: "Sokratovich"}}
	return cars
}
