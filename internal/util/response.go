package util

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, data interface{}) {
	response, err := json.Marshal(map[string]interface{}{
		"data": data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

func WriteError(w http.ResponseWriter, code int, errMessage string) {
	response, err := json.Marshal(map[string]interface{}{
		"error": map[string]interface{}{
			"message": errMessage,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

func WriteErrorf(w http.ResponseWriter, code int, data error) {
	response, err := json.Marshal(map[string]interface{}{
		"error": map[string]interface{}{
			"message": data.Error(),
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}
