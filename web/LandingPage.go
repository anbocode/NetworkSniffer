package web

import (
	"../network"
	"../structures"
	"fmt"
	"net/http"
	"os"
)

func LandingPage() {
	var (
		interfaces structures.ReturnStructure = structures.ReturnStructure{}
		err        error                      = nil
	)

	interfaces = network.GetInterfaces()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			fmt.Fprintf(w, string(interfaces.Result))
		default:
			fmt.Fprintf(w, `
				"Data": "Page broken"
			`)
		}
	})

	if http.ListenAndServe("127.0.0.1:8989", nil); err != nil {
		fmt.Println("Error | http.ListenAndServe() : ", err.Error())
		os.Exit(0)
	}
}
