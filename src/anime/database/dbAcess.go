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

func Kek() string {

	//Сюда собираем хтмл страничку
	page := "<html><head></head><body>"

	db, err := sql.Open("mysql", "root:18091996@/anime?charset=utf8")
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select * from testint order by labels desc")

	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			page += columns[i] + ": " + value + "<br>"
		}
		page += "<br>"
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	defer db.Close()

	page += "</body></html>"
	return page
}
