package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = 8010

func validatePeople(w http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Print(data)
	if v, ok := data["name"]; ok {
		fmt.Print(v)
	} else {
		fmt.Print("no name")
	}
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//
	// // Do something with the Person struct...
	// fmt.Fprintf(w, "Person: %q", p)
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("test")
}

func createServer() {
	http.HandleFunc("GET /test", test)
	http.HandleFunc("POST /peoples", validatePeople)
	log.Printf("Server is running at :%v", PORT)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", PORT), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Closed.")
	} else if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

func main() {
	createServer()
	// a := make([]string, 0, 10)
	// a = append(a, "A")
	// ab := append(a, "B")
	// a := "A"
	// a = a + "A"
	// ab := a + "B"
	// fmt.Println(ab)
	// ac := append(a, "C")
	// ac := a + "C"
	// fmt.Println(ab, ac)
	// abe := append(ab, "E")
	// abe := ab + "E"
	// fmt.Println(abe)
}
