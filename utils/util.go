package utils

import (
	"encoding/json"
	"net/http"
)

func Message(rc string, status bool, message string) map[string]interface{} {
	var valueStatus string
	if status == false {
		valueStatus = "Failed"
	} else {
		valueStatus = "Success"
	}
	return map[string]interface{}{"rc": rc, "detail": valueStatus, "message": message}
}
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
