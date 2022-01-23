package utils

import (
	"encoding/json"
	"go-microservice/model"
	"net/http"
)

func Return(w http.ResponseWriter, status bool, code int, err error, data interface{}) {

	response := model.Response{
		Status: status,
		Code:   code,
		Error:  "",
		Data:   data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	_ = json.NewEncoder(w).Encode(response)
}
