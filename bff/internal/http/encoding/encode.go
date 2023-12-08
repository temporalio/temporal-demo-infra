package encoding

import (
	"encoding/json"
	"net/http"
)

func EncodeJSONResponseBody(w http.ResponseWriter, obj interface{}, statusCode int) error {
	w.Header().Add("Content-Type", "application/json")
	out, err := json.Marshal(&obj)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}
