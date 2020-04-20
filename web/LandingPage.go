package web

import (
	"../network"
	"../structures"
	"fmt"
	"net/http"
	"os"
)

func handleFunc01(w http.ResponseWriter, r *http.Request) {
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

func LandingPage() {
	var (
		fileServer http.Handler = http.FileServer(http.Dir("static/"))
		err        error        = nil
	)

	http.HandleFunc("/interfaces", handleFunc01)
	http.Handle("/main/", http.StripPrefix("/main/", fileServer))

	if http.ListenAndServe("127.0.0.1:8989", nil); err != nil {
		fmt.Println("Error | http.ListenAndServe() : ", err.Error())
		os.Exit(0)
	}
}
