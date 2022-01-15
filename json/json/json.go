package json

import "net/http"

func Json(res http.ResponseWriter, jsonStruct interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	jsonData := []byte(`{"status":"OK"}`)
	res.Write(jsonData)
}
