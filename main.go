package main

import (
	"errors"
	"fmt"
	//"os/user"
	//"weak"
	//"net/http"
)

/*func challenge ()bool {
	username := "admin"
	password := "1234"

	if username == "admin" && password == "1234" {
		fmt.Println("login success")
		return true
	} else{
		fmt.Println("login failed")
		return false
	}

}*/

/*func authentication() string {
	role := "admin"

	switch role {
	case "admin" :
		return"Ful access"
	case "user":
		return "Limited access"
	case "guest":
		return "Read only"
		default:
		return "Read only"

	}
}

func add(a int, b int){
	fmt.Println(a + b)
}

func greet() {
	fmt.Println("this is anika")
}
func greetUser(name string) {
	fmt.Println("Hello", name)
}
func addd(a int, b int) int {
	return a + b
}
func sumNumbers(nums []int) int {
	sum := 0 

	for _, n := range nums {
		sum += n
		}
		return sum 

	}

//p1..
func isEven(num int) bool {
	return num % 2 == 0
}
func max(a int, b int )int {
	if a > b {
		return a
	}
	return b
}
func login(username string, password string) bool {
	if username == "admin" && password == "1234" {
		return true
	}
	return false
}
//multiple return value
func calculate (a int, b int )(int, int,) {
	sum := a + b
	diff := a - b
	return sum, diff 
}

// real world pattern(value , error..)
func divide(a int, b int)(int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
//basic multiple value return..
func checkAccount(id int) (int, bool) {
	if id == 100 {
		return 5000, true 
	}
	return 0, false 
}
// real go errors..
func checkAccount2(id int)(int, error) {
	if id != 100{
		return 0, errors.New("account not found")
	}
	return 5000, nil
}

//challenge..
func checkProduct(name string) (int, error) {
	if name == "laptop" {
		return 50000, nil 
	}
	if name == "mobile" {
		return  20000, nil 
	}
	return 0, errors.New("product not found")
}
func findUser(id int)(string, error){
	if id == 1 {
		return "Anika",  nil
	}
	if id == 2 {
		return "Rahim", nil 
	}
	if id == 3{
		return "Rani", nil
	}
	return "", errors.New("user not found")
}
// step 2 test....
//test 1
func temparetureCheck2(temp int) string{
	if temp >= 30{
		return "Hot"
	}else if temp >= 20{
		return "Normal"
	}else{
		return "Cold"
	}
}
//test 2..
func gradeCheck(grade string)string {
	//grade = "B"
	switch grade{
	case "A" :
		return "Excellent"
	case "B" :
		return "Good"
	case "C" :
		return "Average"
	default:
		return "Fail"
	}
}
//test 3...
func sumSlice(numbers[]int) int {
	sum :=0
	for _, numberr := range numbers {
		sum += numberr
	}
	return sum
}


//test 4.... 
func printUntilSeven(){
	for i := 1; i <=10; i++{
		if i == 7 {
			break
		}
	
	fmt.Println(i)
}
}
//test 5.... 
func login1(username, password string) (bool, string) {
	if username == "admin" && password == "1234" { 
		return true,"login successful" } 
		return false, "invalid Credentials" }



// test 6.... 
func getStatus1(message string, code int){
	message = "Ok"
	code = 200
	return
}
//challange...
func checkNumber(num int)(string, bool) {
	if num % 2 == 0 {
		return "Even", true
	}
	return "Odd", false 
}*/

//===============step C===============//
//topic Array and Slice===============//




func main(){
	//==================step 3========//
	//slice=========================
	numbers := []int{10, 20, 30}
	numbers

	//array=========================
	var books [11]string 
	books[0] = "Golang Basics"
	books[1] = "Web development"
	books[2] = "Data Structures"
	books[3] = "Algorithms"
	books[4] = "Networking"
	books[5] = "go Development"
	books[6] = "Web development"
	books[7] = "Web development"
	books[8] = "Web development"
	books[9] = "Web development"
	books[10] = "Web development"
	fmt.Println(books)
	fmt.Println(books[10])
	fmt.Println(len(books))
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i])
	}
	//cleen way
	for index, value := range books {
		fmt.Println(index, value)
	}
}
	//chalange
	/*numbers := []int{1,2,3,4,5}
	for _, n := range numbers{
	result6, ok :=checkNumber(n)
	fmt.Println(n, "->", result6, ok)
	}

	// step 4..
	printUntilSeven()
	//test 5

	ok, message := login1("admin", "1234")
	fmt.Println(ok, message)
	// test 6
	//msg, statusCode := getStatus1()
	//fmt.Println(msg, statusCode)
	temparetureCheck2(20)
	gradeCheck("B")
	nums := []int{2, 4, 6, 8}
	result5 := sumSlice(nums)
	fmt.Println("Sum", result5)
	///////
	user, err := findUser(3)
	if err != nil {
		fmt.Println("error2", err)
		return
	}
	fmt.Println("user",user)
	//chalange ..
	price, err := checkProduct("mobile")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("price",price)

	
	// real go errors..
	balence2, err := checkAccount2(200)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Balance", balence2)

	//basic multiple value return..
	balance, valid := checkAccount(100)
	fmt.Println(balance)
	fmt.Println(valid)
	//real world pattern(value, error..)
	result3, err := divide(10,2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result3:", result3)

//multiple return value

s, d := calculate(10,5)
fmt.Println("sum:", s)
fmt.Println("Diff:", d)

	number :=10
	result := isEven(number)
	fmt.Println("Is number even?", result)
	numbars :=[] int {10, 20, 30}
	total := sumNumbers(numbars)
	fmt.Println(total)
	add(20, 20)
	greet()
	greetUser("Anika")
	greetUser("Rahim")
	result2 := addd(10, 20)
	fmt.Println(result2)
}


//range loop...........

/*students := []string{"Rahim", "Karim", "Salam"}

for index, name := range students {
	fmt.Println(index, name)
}

for _,name := range students {
	fmt.println(name)
}
// sum print...........

numbers := []int{ 10, 20, 30, 40}

sum := 0

for _, number := range numbers {
	sum += numbers

}
fmt.Println(sum)
}


// slice creating..........

/*users := []string{"user1","user2","user3"}
for _, user := range users {
	fmt.Println("Sending email to :", user)
}
	// mini challenge 2

	for i := 1; i <= 20; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	//mini challenge
	for i := 1; i <= 20; i++ {
		if i == 7 {
			break
		}
		fmt.Println(i)
	}

	for i :=1; i <=20; i++ {
		if i % 2 == 0  {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	for i :=10; i >= 1; i-- {
	
		fmt.Println(i)
	}

	i := 2

	for i <= 20 {
		fmt.Println(i)
		i++
	}

	
	for {
		var input string
		fmt.Scanln(&input)

		if input == "exit"{
			break
		}
	}

	fruits := [] string{"Apple", "banana", "Mango"}

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	//authentication()
	/*drink := "modhu"
	if drink == "tea" {
		fmt.Println("Tea")
	}else if drink == "coffie" {
		fmt.Println("coffie")
	}else if drink == "juice" {
		fmt.Println("Juice")
	}else {
		fmt.Println("Nothing")
	}

	drinkk := "Water"

	switch drinkk {
	case "tea":
		fmt.Println("Tea")
	case "coffe": 
	fmt.Println("coffee")
	case "juice" :
		fmt.Println("juice")
	default: 
	fmt.Println(("Nothing"))
	}

	day := "friday"

	switch day {
	case "satarday":
		fmt.Println("satarday")
	case "sunday" :
		fmt.Println("sunday")
	case "monday":
		fmt.Println("monday")
	default:
		fmt.Println("Normal day")
	}

	day = "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Holiday")
	default:
		fmt.Println("working day")
	}

	temperature := 25

	switch {
	case temperature >= 30:
		fmt.Println("Hot")
	case temperature >= 20:
		fmt.Println("Normal")
	default:
		fmt.Print("cold")
	} 

	status := 404

	switch status {
	case 200:
		fmt.Println("ok")
	case 401:
		fmt.Println("Unauthorized")
	case 404:
		fmt.Println("not found")
	case 500:
		fmt.Println("Server Error")
	default:
		fmt.Println("Unknown status")
	}*/
//practice 1
/*for i :=1; i <= 5; i++ {
		fmt.Println(i)
	}


	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}*/ 

	/*for i := 10; i <=1; i++ {
		fmt.Println(i)
	} 

	
	number1 := 3

	switch number1 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}
//practice 2
    grade := "B"
	switch grade{
	case "A" :
		fmt.Println("Excellent")
	case "B" :
		fmt.Println("Good")
	case "C" :
		fmt.Println("Average")
		default :
		fmt.Println("Fail")
	}
//practice 3

agee := 75

switch {
case agee >= 60 :
	fmt.Println("Senior")
case agee >= 18 :
	fmt.Println("Adult")
default:
	fmt.Println("Child")
}

//mini challenge



	



	//challenge()
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
		fmt.Println(" update Child ")
	}
	//for {
		//fmt.Println("Hello")
	//}
	







/*func helloHandler(w http.ResponsWriter, *http.Request){
	fmt.Fprintf(w, "hello i am anika . welcome my go web server")
}

func main() {
	http.Handlefunc("\" helloHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}*/
