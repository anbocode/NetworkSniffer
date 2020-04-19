package web

import (
	"../network"
	"../structures"
	"fmt"
	"net/http"
)

func LandingPage() {
	var (
		interfaces structures.ReturnStructure = structures.ReturnStructure{}
	)

	interfaces = network.GetInterfaces()
	fmt.Println(interfaces)
}
