package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func ExampleYurtici() {

	yurtici, err := util.NewProviderByName("yurtici")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(yurtici.CreateCargo(nil))

}
