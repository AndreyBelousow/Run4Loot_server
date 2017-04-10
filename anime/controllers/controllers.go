package controllers

import (
	"anime/database"

	//"anime/view"
)

//Увеличение счетчика у пользователя
func Update(userid string, count int) {
	database.Update(userid, count)
}

//Удаление пользователя
func Delete(userid string) {
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
