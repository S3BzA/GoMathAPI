package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/S3BzA/GoMathAPI/mathops"
	"github.com/S3BzA/GoMathAPI/utils"
)

func Sum(w http.ResponseWriter, r *http.Request) {
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
	
	result := utils.MathResult{Result: mathops.Sum(numbers...)}
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

func Mul(w http.ResponseWriter, r *http.Request) {
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

	result := utils.MathResult{Result: mathops.Multiply(numbers...)}
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

func Div(w http.ResponseWriter, r *http.Request) {
	raw := r.PathValue("nums")
	numbers, err := utils.StringsToFloats(strings.Split(raw, ","))

	if len(numbers) != 2 {
		errJSON, jerr := utils.JsonEncode(utils.MathError{Error: "need two operands"})
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

	result := utils.MathResult{Result: mathops.Divide(numbers[0], numbers[1])}
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

func Pow(w http.ResponseWriter, r *http.Request) {
	raw := r.PathValue("nums")
	numbers, err := utils.StringsToFloats(strings.Split(raw, ","))

	if len(numbers) != 2 {
		errJSON, jerr := utils.JsonEncode(utils.MathError{Error: "need two operands"})
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

	result := utils.MathResult{Result: mathops.Power(numbers[0], int64(numbers[1]))}
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

func Mod(w http.ResponseWriter, r *http.Request) {
	raw := r.PathValue("nums")
	numbers, err := utils.StringsToInts(strings.Split(raw, ","))

	if len(numbers) != 2 {
		errJSON, jerr := utils.JsonEncode(utils.MathError{Error: "need two operands"})
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

	result := utils.MathResult{Result: float64(mathops.Modulo(numbers[0], numbers[1]))}
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
