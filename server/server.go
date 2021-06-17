package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"model/product"
	"net/http"
	"strconv"
	"strings"
)

var path string
var splitted []string
var length int

func main() {
	http.HandleFunc("/products", controllerProducts)
	http.HandleFunc("/products/", controllerProducts)
	http.HandleFunc("/products/new", controllerProducts)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func pathChecker(r *http.Request) {
	path = r.URL.Path[1:]
	splitted = strings.Split(path, "/")
	length = len(splitted)
}

func controllerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	pathChecker(r)

	if length == 1 {
		if r.Method != "GET" {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}
		productsIndex(w)
	} else if length == 2 {
		if r.Method == "GET" {
			productsShow(w)
		} else if r.Method == "POST" {
			productsCreate(w, r)
		} else if r.Method == "PUT" || r.Method == "PATCH" {
			productsUpdate(w, r)
		} else if r.Method == "DELETE" {
			productsDelete(w, r)
		} else {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
	}
}

func productsIndex(w http.ResponseWriter) {
	products, err := model.Index()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func productsShow(w http.ResponseWriter) {
	strId := splitted[length-1]
	intId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	product, err := model.Show(intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(result)
	return
}

func productsCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))

	var product *model.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(product)

	result, err := model.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, _ := json.Marshal(result)
	w.Write(resp)
	return
}

func productsUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))

	var product *model.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result, err := model.Update(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, _ := json.Marshal(result)
	w.Write(resp)
	return
}

func productsDelete(w http.ResponseWriter, r *http.Request) {
	strId := splitted[length-1]
	intId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	success, err := model.Delete(intId)
	if success {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("deleted"))
		return
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
