package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/bad", getBadRequest).Methods(http.MethodGet)
	muxRouter.HandleFunc("/name/{PARAM}", getHelloRequest).Methods(http.MethodGet)
	muxRouter.HandleFunc("/data", postDataFunc).Methods(http.MethodPost)
	muxRouter.HandleFunc("/headers", postHeadersFunc).Methods(http.MethodPost)
	muxRouter.HandleFunc("/", defaultFunc)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), muxRouter); err != nil {
		log.Fatal(err)
	}

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

func getHelloRequest(w http.ResponseWriter, r *http.Request) {
	var urlParam = mux.Vars(r)["PARAM"]
	w.Write([]byte(fmt.Sprintf("Hello, %s!", urlParam)))
}

func postDataFunc(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("I got message:\n%s", requestBody)))
}

func postHeadersFunc(w http.ResponseWriter, r *http.Request) {
	var sum = 0
	for k, v := range r.Header {
		fmt.Println(k)
		fmt.Println(v)

		if strings.EqualFold(k, "a") || strings.EqualFold(k, "b") {
			value, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatal(err)
			}
			sum += value
		}
	}
	fmt.Println(sum)
	w.Header().Set("a+b", strconv.Itoa(sum))
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
