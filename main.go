package main

import(
	"fmt"
	//"net/http"
)

func challenge ()bool {
	username := "admin"
	password := "1234"

	if username == "admin" && password == "1234" {
		fmt.Println("login success")
		return true
	} else{
		fmt.Println("login failed")
		return false
	}

}
func main(){
	challenge()
	number := 7

	if number % 2 == 0 { 
		fmt.Println("Even")
	}else {
		fmt.Println("Odd")
	}

	tempareture := 35

	if tempareture >= 30{
		fmt.Println("Hot")
	} else if tempareture >= 20 {
		fmt.Println("normal")
	}else{
		fmt.Println("cold")
	}

	balance := 500

	if balance >= 1000{
		fmt.Println("you can withdraw")
	}else {
		fmt.Println("Insufficient balance")
	}

	age := 15

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13{
		fmt.Println("Teen")
	} else {
		fmt.Println("Child ")
	}
	}








/*func helloHandler(w http.ResponsWriter, *http.Request){
	fmt.Fprintf(w, "hello i am anika . welcome my go web server")
}

func main() {
	http.Handlefunc("\" helloHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}*/
