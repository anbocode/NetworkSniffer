package web

import (
	//"../network"
	//"../structures"
	"fmt"
	"net/http"
	"os"
)

func LandingPage() {
	var (
		err error = nil
	)

	http.HandleFunc("/interfaces", GetInterface)
	http.Handle("/web/", FileServer("/web/"))

	if http.ListenAndServe("127.0.0.1:8989", nil); err != nil {
		fmt.Println("Error | http.ListenAndServe() : ", err.Error())
		os.Exit(0)
	}
}
