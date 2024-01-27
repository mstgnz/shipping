package main

import (
	"log"

	"github.com/mstgnz/shipping/util"
)

func main() {

	ptt, err := util.NewProviderByName("ptt")

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(ptt.CreateCargo(nil))

}
