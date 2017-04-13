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

	//ДА да, именно так в го пишется while(true)
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

	//Переключение режима работы
	m.Get("/switch", func(res http.ResponseWriter, req *http.Request) string {
		requrl := req.URL.Query()
		pass := requrl.Get("pass")
		if pass == "kek" {
			canWork = !canWork
			if canWork {
				fmt.Println("__________________________________________")
				fmt.Println("SERVER  IS ONLINE NOW")
				fmt.Println("__________________________________________")
				return "SWITCHED ONLINE"
			} else {
				fmt.Println("__________________________________________")
				fmt.Println("SERVER  IS OFFLINE NOW")
				fmt.Println("__________________________________________")
				return "SWITCHED OFFLINE"
			}
		} else {
			return "INCORRECT PASSWORD"
		}
	})

	//Обращение к табличке с результатами
	m.Get("/table", func(res http.ResponseWriter) string {
		if canWork {
			return controllers.ShowUsersTable()
		} else {
			return "SERVER OFFLINE"
		}
	})

	//Увеличение счетчика у пользователя
	m.Get("/update", func(res http.ResponseWriter, req *http.Request) string {
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
			return "UPDATED:" + userid
		} else {
			res.WriteHeader(1488)
			fmt.Println("__________________________________________")
			fmt.Println("ACESS DENIED")
			fmt.Println("__________________________________________")
			return "SERVER OFFLINE"
		}
	})

	//Удаление пользователя
	m.Get("/delete", func(res http.ResponseWriter, req *http.Request) string {
		if canWork {
			requrl := req.URL.Query()
			userid := requrl.Get("id")
			controllers.Delete(userid)
			res.WriteHeader(200)
			fmt.Println("__________________________________________")
			fmt.Println("DELETED id:", userid)
			fmt.Println("__________________________________________")
			return "DELETED:" + userid

		} else {
			res.WriteHeader(1488)
			fmt.Println("__________________________________________")
			fmt.Println("ACESS DENIED")
			fmt.Println("__________________________________________")
			return "SERVER OFFLINE"
		}
	})

	m.Run()
}
