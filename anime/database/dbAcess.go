package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db = new(sql.DB)

//подключение к базе
func OnStart() {
	db, err := sql.Open("mysql", "JESUS:18091996@/anime?charset=utf8")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

//Увеличение счетчика
func Update(userid int, count int) {
	db.Query("update users set labels=", count, "where userid=", userid)
}

//Удаление пользователя
func Delete(userid int) {
	db.Query("delete from users where userid=", userid)
}
