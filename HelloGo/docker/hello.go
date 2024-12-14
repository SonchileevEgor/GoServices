package main

import (
	// "errors"
	"fmt"
	// "strconv"
	// "io"
	// "net/http"
	// "regexp"
	// "strings"
)

// Сигнатура функции
// func add (i int, j int) int {
// 	return i+j;
// }
// func sub (i int, j int) int {
// 	return i-j;
// }
// func mul (i int, j int) int {
// 	return i*j;
// }
// func div (i int, j int) int {
// 	return i/j;
// }

// var opMap = map[string] func(int, int) int {
// 	"+": add,
// 	"-": sub,
// 	"*": mul,
// 	"/": div,
// }
// Сигнатура функции

func main() {
	fmt.Println("Hello, GoLang!")
	// var a, b int
	// fmt.Scanln(&a, &b)
	// fmt.Println(a + b)
	// var c int
	// if a > b {
	// 	c = 1
	// } else {
	// 	c = -1
	// }
	// fmt.Println(c)

	//parser
	// var response, err = http.Get("https://pro-bodyart.ru/listCoaches")
	// var respBody, errorBody = io.ReadAll(response.Body)
	// if err != nil{
	// 	fmt.Println("error in response")
	// }
	// if (errorBody != nil) {
	// 	fmt.Println("error in reading response body")
	// }
	// var strBody = string(respBody)
	// var matchedSliv = regexp.MustCompile("Почта:.+u")
	// var emails = matchedSliv.FindAll([]byte(strBody), -1)
	// for _, element := range emails {
	// 	fmt.Println(strings.Replace(string(element), "</b>", "", 1))
	// }
	//parser

	// fmt.Print(add(2, 3));
	
	// expressions := [][]string {
	// 	{"2", "+", "3"},
	// 	{"2", "-", "3"},
	// 	{"2", "*", "3"},
	// 	{"2", "/", "3"},
	// 	{"two", "+", "three"},
	// 	{"2", "%", "3"},
	// }

	// for _, expression := range expressions {
	// 	if len(expression) != 3 {
	// 		fmt.Println("Invalid expression", expression)
	// 	}
	// 	p1, err := strconv.Atoi(expression[0])
	// 	if (err != nil) {
	// 		fmt.Println("error", err);
	// 		continue
	// 	}
	// 	op := expression[1]
	// 	opFunc, ok := opMap[op]
	// 	if !ok {
	// 		fmt.Println("unsupported operator", op)
	// 		continue
	// 	}

	// 	p2, err := strconv.Atoi(expression[2])
	// 	if (err != nil) {
	// 		fmt.Println("error", err);
	// 		continue
	// 	}
	// 	result := opFunc(p1,p2)
	// 	fmt.Println(result)
	// }

	
}
