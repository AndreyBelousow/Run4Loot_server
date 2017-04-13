package controllers

import "anime/database"

//Увеличение счетчика у пользователя
func Update(userid string, count int) string {
	database.Update(userid, count)
	return "updated " + userid
}

//Удаление пользователя
func Delete(userid string) string {
	database.Delete(userid)
	return "deleted " + userid
}

//Показать табличку с результатами
func ShowUsersTable() string {
	return database.Kek()
	return "here is table"
}
