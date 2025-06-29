package server

import (
	"fmt"
	"net/http"
	"parce/data"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func ServerStart() {
	// Регистрируем обработчики для разных маршрутов
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Запускаем сервер
	fmt.Printf("Starting server at port %s", data.ServerPort)
	fmt.Println(data.ServerPort)
	err := http.ListenAndServe(":"+data.ServerPort, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
