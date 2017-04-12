package main

import (
	"os"

	"anime/controllers"

	"strconv"

	"fmt"

	"net/http"

	"github.com/go-martini/martini"

	"bufio"
)

func main() {

	m := martini.Classic()

	canWork := false

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ALLOW ACESS? [y/n]")

	flag := true
	for flag {
		char, _, err := reader.ReadRune()
		if err != nil {
		}
		switch char {
		case 'y':
			fmt.Println("ACESS ALLOWED")
			canWork = true
			flag = false
			break
		case 'n':
			fmt.Println("ACESS DENIED")
			canWork = false
			flag = false
			break
		}
	}

	//Обращение к табличке с результатами
	m.Get("/table", controllers.ShowUsersTable)

	//Увеличение счетчика у пользователя
	m.Get("/update", func(res http.ResponseWriter, req *http.Request) {
		if canWork {
			requrl := req.URL.Query()
			userid := requrl.Get("id")
			count, err := strconv.Atoi(requrl.Get("count"))
			if err != nil {
				fmt.Println(err)
			}
			controllers.Update(userid, count)
			res.WriteHeader(200)
			fmt.Println("__________________________________________")
			fmt.Println("UPDATED id:", userid, " count:", count)
			fmt.Println("__________________________________________")
		}
		res.WriteHeader(1488)
		fmt.Println("__________________________________________")
		fmt.Println("ACESS DENIED")
		fmt.Println("__________________________________________")
	})

	//Удаление пользователя
	m.Get("/delete", func(res http.ResponseWriter, req *http.Request) {
		if canWork {
			requrl := req.URL.Query()
			userid := requrl.Get("id")
			controllers.Delete(userid)
			res.WriteHeader(200)
			fmt.Println("__________________________________________")
			fmt.Println("DELETED id:", userid)
			fmt.Println("__________________________________________")
		}
		res.WriteHeader(1488)
		fmt.Println("__________________________________________")
		fmt.Println("ACESS DENIED")
		fmt.Println("__________________________________________")
	})

	m.Run()
}
