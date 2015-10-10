package main

import (
	"encoding/json"
	"fmt"
	"github.com/twlicate/vk"
	"log"
)

type user struct {
	Id      int    `json:"uid"`
	Name    string `json:"first_name"`
	Surname string `json:"last_name"`
}

func main() {

	var vkapi vk.Api

	// Open methods does not require token
	// vkapi.Token = "..."

	params := make(map[string]string)
	params["user_ids"] = "1"

	result, e := vkapi.Request("users.get", params)
	if e != nil {
		log.Fatalln(e)
	}

	// Remove [] from json
	result = result[1 : len(result)-1]

	var u user
	json.Unmarshal(result, &u)
	fmt.Printf("Id: %d, Name: %s, Surname: %s", u.Id, u.Name, u.Surname)
}
