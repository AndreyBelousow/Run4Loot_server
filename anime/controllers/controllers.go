package controllers

import (
	"anime/database"
	//"anime/view"
)

//KEK SHREK
func KekShrek() string {
	return "KEK SHREK"
}

//Увеличение счетчика у пользователя
func Update(userid int, count int) {
	database.Update(userid, count)
}

//Удаление пользователя
func Delete(userid int) {
	database.Delete(userid)
}

//Показать табличку с результатами
func ShowUsersTable() string {
	return "ТУТ ТИПА ХТМЛ ТАБЛИЧКА С РЕЗУЛЬТАТАМИ"
}

//Подключение к базе
func StartBase() {
	database.OnStart()
}
