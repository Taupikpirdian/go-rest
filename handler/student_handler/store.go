package student_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"try/go-rest/http_request"
)

func (s_handler *StudentHandler) StoreDataBuku(w http.ResponseWriter, r *http.Request) {
	var (
		req     http_request.RequestStudent
		decoder = json.NewDecoder(r.Body)
	)

	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}

	fmt.Println(req)

	w.WriteHeader(200)
	fmt.Fprintf(w, "SUCCESS INSERT DATA")
	return
}
