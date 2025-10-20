package handlers

import (
	"net/http"

	"github.com/S3BzA/GoMathAPI/database"
)

func Create(w http.ResponseWriter, r *http.Request) {

}

func Read(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB("user","password123","localhost","5432","mydatabase")
}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}