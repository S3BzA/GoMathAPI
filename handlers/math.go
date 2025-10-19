package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/S3BzA/GoMathAPI/mathops"
	"github.com/S3BzA/GoMathAPI/utils"
)

func Add(w http.ResponseWriter, r *http.Request) {
	raw := r.PathValue("nums")
	numbers, err := utils.StringsToFloats(strings.Split(raw, ","))
	if err != nil {
		errJSON, jerr := utils.JsonEncode(utils.MathError{Error: "invalid data type"})
		if jerr != nil {
			log.Printf("json encode error: %v", jerr)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if werr := utils.WriteJSON(w, http.StatusBadRequest, errJSON); werr != nil {
			log.Printf("response write error: %v", werr)
		}
		return
	}

	result := utils.MathResult{Result: mathops.Add(numbers...)}
	resultJSON, jerr := utils.JsonEncode(result)
	if jerr != nil {
		log.Printf("json encode error: %v", jerr)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if werr := utils.WriteJSON(w, http.StatusOK, resultJSON); werr != nil {
		log.Printf("response write error: %v", werr)
	}
}

func Sub(w http.ResponseWriter, r *http.Request) {

}

func Mul(w http.ResponseWriter, r *http.Request) {

}

func Div(w http.ResponseWriter, r *http.Request) {

}

func Pow(w http.ResponseWriter, r *http.Request) {

}

func Mod(w http.ResponseWriter, r *http.Request) {

}
