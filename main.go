package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type RBody struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type RBodySum struct {
	Arr []int `json:"array"`
}

func checkCT(ct string, w http.ResponseWriter) {
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
		if mediaType != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Calculator Api")
	log.Println("endpoint / called")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var body RBody

	ct := r.Header.Get("Content-Type")
	checkCT(ct, w)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := body.Number1 + body.Number2
	log.Printf("Hasil penjumlahan adalah: %d\n", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": ` + strconv.Itoa(result) + `}`))
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	var body RBody

	ct := r.Header.Get("Content-Type")
	checkCT(ct, w)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := body.Number1 - body.Number2
	log.Printf("Hasil pengurangan adalah: %d\n", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": ` + strconv.Itoa(result) + `}`))

}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	var body RBody

	ct := r.Header.Get("Content-Type")
	checkCT(ct, w)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := body.Number1 / body.Number2
	log.Printf("Hasil pembagian adalah: %d\n", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": ` + strconv.Itoa(result) + `}`))

}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	var body RBody

	ct := r.Header.Get("Content-Type")
	checkCT(ct, w)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := body.Number1 * body.Number2
	log.Printf("Hasil perkalian adalah: %d\n", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": ` + strconv.Itoa(result) + `}`))

}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	var body RBodySum

	ct := r.Header.Get("Content-Type")
	checkCT(ct, w)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := 0
	for _, v := range body.Arr {
		result += v
	}
	log.Printf("Hasil Penjumlahan Array adalah: %d", result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result": ` + strconv.Itoa(result) + `}`))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/subtract", subtractHandler)
	mux.HandleFunc("/divide", divideHandler)
	mux.HandleFunc("/multiply", multiplyHandler)
	mux.HandleFunc("/sum", sumHandler)

	log.Println("starting server on :6969")

	err := http.ListenAndServe(":6969", mux)
	log.Fatal(err)
}
