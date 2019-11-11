package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
)

type HTTP struct {
	ErrorToStatusCode func(error) int
	ErrorToMessage    func(error) string
}

func (HTTP) DecodeRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return fmt.Errorf("method not allowed")
	}

	ct := r.Header.Get("Content-Type") // application/json; charset=utf-8
	mt, _, _ := mime.ParseMediaType(ct)
	if mt != "application/json" {
		http.Error(w, "Unsupported MediaType", http.StatusUnsupportedMediaType)
		return fmt.Errorf("invalid mediatype")
	}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return fmt.Errorf("bad request")
	}
	return nil
}

func (t HTTP) EncodeResult(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		status := t.ErrorToStatusCode(err)
		if status == http.StatusInternalServerError {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{t.ErrorToMessage(err)})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
