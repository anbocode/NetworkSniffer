package web

import (
	"../network"
	"../structures"
	"fmt"
	"net/http"
)

func GetInterface(w http.ResponseWriter, r *http.Request) {
	var (
		interfaces         structures.ReturnStructure = structures.ReturnStructure{}
		unrecognisedMethod string                     = `
			"Data": "Page broken"
		`
	)

	interfaces = network.GetInterfaces()

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, string(interfaces.Result))
	default:
		fmt.Fprintf(w, unrecognisedMethod)
	}
}
