package main

import (
	"anime/controllers"

	"strconv"

	"fmt"

	"net/http"

	"github.com/go-martini/martini"
)

func main() {

	m := martini.Classic()

	//Обращение к табличке с результатами
	m.Get("/table", controllers.ShowUsersTable)

	//Увеличение счетчика у пользователя, стоит заменить на POST
	m.Get("/update", func(req *http.Request) {
		requrl := req.URL.Query()
		userid := requrl.Get("id")
		count, err := strconv.Atoi(requrl.Get("count"))
		if err != nil {
			fmt.Println(err)
		}
		controllers.Update(userid, count)
		fmt.Println("__________________________________________")
		fmt.Println("UPDATED id:", userid, " count:", count)
		fmt.Println("__________________________________________")
	})

	//Удаление пользователя
	m.Get("/delete", func(req *http.Request) {
		requrl := req.URL.Query()
		userid := requrl.Get("id")
		controllers.Delete(userid)
		fmt.Println("__________________________________________")
		fmt.Println("DELETED id:", userid)
		fmt.Println("__________________________________________")
	})

	m.Run()
}
