package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

const PORT = 8010

func getContractSchema() {
	data, err := os.ReadFile("./contracts/people.yml")
	if err != nil {
		fmt.Printf("err readfile: %+v\n", err)
	}

	var schema map[string]interface{}
	err = yaml.Unmarshal(data, &schema)
	if err != nil {
		fmt.Printf("err unmarshal: %+v\n", err)
	}
	fmt.Printf("contract: %+v\n", schema)
}

func validatePeople(w http.ResponseWriter, req *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %+v\n", data)
	if v, ok := data["name"]; ok {
		fmt.Print(v)
	} else {
		fmt.Print("no name")
	}
	getContractSchema()
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
