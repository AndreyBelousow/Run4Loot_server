package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"strconv"
)

//Увеличение счетчика
func Update(userid string, count int) {
	db, err := sql.Open("mysql", "root:18091996@/anime?charset=utf8")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	kek, err := strconv.Atoi(userid)
	stmtIns, err := db.Prepare("insert into testint values (?, ?)")
	_, err = stmtIns.Exec(kek, count)
	stmtUpd, err := db.Prepare("update testint set labels = ? where userid = ?")
	_, err = stmtUpd.Exec(count, kek)
	defer db.Close()
}

//Удаление пользователя
func Delete(userid string) {
	db, err := sql.Open("mysql", "root:18091996@/anime?charset=utf8")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	kek, err := strconv.Atoi(userid)
	stmtDel, err := db.Prepare("delete from testint where userid=?")
	_, err = stmtDel.Exec(kek)
	defer db.Close()
}
