package api

import (
	"fmt"
	"io"
	"net/http"
)

var data string

func InputDump(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Cannot read body...", http.StatusBadRequest)
				return
			}

			data = string(body)
			fmt.Fprintf(w, "Json has been stored...\n")

		case http.MethodGet:
			fmt.Fprint(w, data)
			fmt.Println(data)
		
	}

}