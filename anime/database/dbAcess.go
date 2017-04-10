package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db = new(sql.DB)
var stmtIns = new(sql.Stmt)

//Подключение к базе
func OnStart() {
	db, err := sql.Open("mysql", "root:18091996@/anime?charset=utf8")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

//Увеличение счетчика
func Update(userid string, count int) {
	stmtIns, err := db.Prepare("insert into testint values (?, ?)")
	_, err = stmtIns.Exec(1488, count)
	if err != nil {
		panic(err.Error())
	}
	//db.Query("insert into users (char,int) values(", userid, count, ")")
}

//Удаление пользователя
func Delete(userid string) {
	db.Query("delete from users where userid=", userid)
}
