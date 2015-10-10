package main

import (
	"fmt"
	"github.com/twlicate/vk"
	"log"
)

func main() {

	// For run this you need token and "messages" access
	// Paste token here:
	vkapi := vk.Api{"..."}

	params := make(map[string]string)
	params["count"] = "5"
	params["out"] = "1"
	json, e := vkapi.Request("messages.get", params)
	if e != nil {
		log.Fatalln(e)
	}

	// Do with json what necessary
	fmt.Println(string(json))
}
