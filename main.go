package main

import (
	"fmt"

	"github.com/vitor7001/simulator/application/route"
)

func main() {
	route := route.Route{
		ID:        "1",
		ClientID:  "1",
		Positions: nil,
	}

	route.LoadPositions()

	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[1])
}
