package main

import (
	"encoding/json"
	"net/http"
)

type Buku struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul"`
	Pengarang string `json:"pengarang"`
}

var perpustakaan []Buku

func GetBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(perpustakaan)
}

func CreateBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var params Buku
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	perpustakaan = append(perpustakaan, params)
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(perpustakaan)
}

func main() {
	mux := http.NewServeMux()
	perpustakaan = []Buku{
		{ID: 1, Judul: "One Piece", Pengarang: "Eichiro Oda"},
		{ID: 2, Judul: "Conan", Pengarang: "Aoyama Gosho"},
	}
	mux.HandleFunc("/books", GetBooks)
	mux.HandleFunc("/create-books", CreateBooks)
	http.ListenAndServe(":8000", mux)
}
