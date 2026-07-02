package api

import (
	"fmt"
	"io"
	"net/http"
)

var data string

func InputDump(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method...", http.StatusBadRequest)
	} else {
		body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Cannot read body...", http.StatusBadRequest)
				return
			}

			data = string(body)
			fmt.Fprint(w, data)
			fmt.Println(data)
	}
}