// main
package main

import (
	"fmt"
	//	"io"
	"log"
	"strconv"
	"strings"
	//	"mystruct"
	"net/http"
	//	"stringutil"
)

func cal(problem, opertion string) int {
	subStrings := strings.Split(problem, opertion)
	num1, _ := strconv.Atoi(subStrings[0])
	num2, _ := strconv.Atoi(subStrings[1])
	if strings.EqualFold(opertion, "+") {
		return num1 + num2
	} else if strings.EqualFold(opertion, "-") {
		return num1 - num2
	} else if strings.EqualFold(opertion, "/") {
		return num1 / num2
	} else {
		return num1 * num2
	}
	return 0
}

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[1:]
	var result = 0
	if strings.Contains(path, "+") {
		result = cal(path, "+")
	} else if strings.Contains(path, "*") {
		result = cal(path, "*")
	} else if strings.Contains(path, "-") {
		result = cal(path, "-")
	} else if strings.Contains(path, "/") {
		result = cal(path, "/")
	}
	fmt.Println("path: %v ", path)
	fmt.Println("request.URL.String(): %v ", request.URL.String())
	fmt.Fprintf(w, "%v = %v", path, result)
	//	io.WriteString(w, "hello, Xiao gang")
}

func main() {
	http.HandleFunc("/", SimpleServer)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
