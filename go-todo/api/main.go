package main

import (
	//"bufio"
	"fmt"
	"net/http"
	//"os"
	//"strings"
)

func main(){
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/tasks/add", AddTaskHandler)

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Todo API Running")
	})
	fmt.Println("Server running on: 4000")
	http.ListenAndServe(":4000", nil)
}

