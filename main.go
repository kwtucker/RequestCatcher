package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	setupResponse(&w, r)

	fmt.Println("Parsing Request")
	fmt.Println(r.URL.String())

	query := r.URL.Query()
	respByt, err := json.MarshalIndent(query, "", "   ")
	if err != nil {
		fmt.Println("Query Fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByt)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func main() {
	router := httprouter.New()
	router.PUT("/*id", Index)
	router.POST("/*id", Index)
	router.GET("/*id", Index)

	log.Fatal(http.ListenAndServe("localhost:1234", router))
}
