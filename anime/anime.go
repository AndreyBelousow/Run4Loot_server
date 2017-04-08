package main

import (
	"anime/controllers"

	"github.com/go-martini/martini"
)

func main() {

	m := martini.Classic()
	controllers.StartBase()

	//Вызов кека со шреком
	m.Get("/kekshrek", controllers.KekShrek)

	//Обращение к табличке с результатами
	m.Get("/table", controllers.ShowUsersTable)

	//Увеличение счетчика у пользователя
	m.Post("/", controllers.Update)

	//Удаление пользователя
	m.Delete("/", controllers.Delete)

	m.Run()
}
