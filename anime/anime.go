package main

import (
	"anime/controllers"

	"strconv"

	"fmt"

	"log"

	"net/http"

	"github.com/go-martini/martini"
)

func main() {

	//Режим работы, продакшн или нет
	//prod := false

	m := martini.Classic()

	//if prod {
	controllers.StartBase()

	//Обращение к табличке с результатами
	m.Get("/table", controllers.ShowUsersTable)

	//Увеличение счетчика у пользователя, стоит заменить на POST
	m.Get("/", func(req *http.Request) {
		requrl := req.URL.Query()
		userid := requrl.Get("id")
		count, err := strconv.Atoi(requrl.Get("count"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(userid)
		fmt.Println(count)

		controllers.Update(userid, count)
	})

	//Удаление пользователя
	m.Delete("/", func(req *http.Request) {
		requrl := req.URL.Query()
		userid := requrl.Get("id")
		controllers.Delete(userid)
	})
	//} else {
	//	m.Get("/table")
	//}

	m.Run()
}
