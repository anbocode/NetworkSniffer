package network

import (
	"../structures"
	"container/list"
	"encoding/json"
	"fmt"
	"net"
)

type resultJSONStructure struct {
	Interfaces []structures.InterfaceStructure
}

func GetInterfaces() structures.ReturnStructure {
	var (
		interfaces              []net.Interface
		interfaceStructures     []structures.InterfaceStructure
		result                  resultJSONStructure = resultJSONStructure{}
		interfaceStructuresJSON []byte
		errorList               *list.List                 = list.New()
		returnStructure         structures.ReturnStructure = structures.ReturnStructure{}
		err                     error                      = nil
	)

	if interfaces, err = net.Interfaces(); err != nil {
		_ = errorList.PushBack(err.Error())
	} else {
		interfaceStructures = make([]structures.InterfaceStructure, len(interfaces))
		result.Interfaces = make([]structures.InterfaceStructure, len(interfaces))
	}

	for i := 0; i < len(interfaces); i++ {
		interfaceStructures[i].Index = interfaces[i].Index
		interfaceStructures[i].MTU = interfaces[i].MTU
		interfaceStructures[i].Name = interfaces[i].Name
		interfaceStructures[i].HardwareAddress = interfaces[i].HardwareAddr.String()
		result.Interfaces[i] = interfaceStructures[i]
	}

	if interfaceStructuresJSON, err = json.MarshalIndent(interfaceStructures, "", "  "); err != nil {
		_ = errorList.PushBack(err.Error())
	}

	returnStructure.Result = interfaceStructuresJSON

	if errorList.Len() != 0 {
		var (
			index int           = 0
			e     *list.Element = nil
		)
		returnStructure.Error = true
		returnStructure.ErrorList = make([]string, errorList.Len())
		for e = errorList.Front(); e != nil; e = e.Next() {
			returnStructure.ErrorList[index], _ = e.Value.(string)
			index++
		}
	}

	return returnStructure
}
