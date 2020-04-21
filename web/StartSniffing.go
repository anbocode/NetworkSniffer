package web

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func StartSniffing(w http.ResponseWriter, r *http.Request) {
	var (
		unrecognisedMethod string = `
			"Data": "Page broken"
		`
	)

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":

		var (
			interfaceNumber int
		)

		r.ParseForm()
		interfaceNumber, _ = strconv.Atoi(r.Form["interfaceNumber"][0])

		fmt.Println("Interface number receivec : ", interfaceNumber)
		fmt.Fprintf(w, `{
			"interface": "Sniffing will start soon"
		}`)
	default:
		fmt.Fprintf(w, unrecognisedMethod)
	}
}
