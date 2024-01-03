package main

import (
	"encoding/json"
	"log"
)

const json_raw string = `
[
	{
		"login_name": "proxyman",
		"last_login": "3 days ago",
		"email": "proxyman@test.test"	
	},
	{
		"login_name": "dancer",
		"last_login": "15 mins ago",
		"email": "dancer@test.test"	
	}
]
`

type Credentials struct {
	Login      string `json:"login_name"`
	Last_login string `json:"last_login"`
	Email      string `json:email`
}

func main() {
	// Read json - Unmarshal
	log.Println("######### JSON Unmarshal ############")
	var creds []Credentials
	err := json.Unmarshal([]byte(json_raw), &creds)
	if err != nil {
		log.Println("Error during JSON read: ", err)
	}
	log.Println(creds)

	// Marshal a.k.a Create JSON from variables
	log.Println("######### JSON Marshal ############")
	var cred1 Credentials
	cred1.Login = "TesterJack"
	cred1.Last_login = "Never"
	cred1.Email = "TesterJack@test.test"

	var cred2 Credentials
	cred2.Login = "DuaLipa"
	cred2.Last_login = "1 mins ago"
	cred2.Email = "DuaLipa@test.test"

	var dev_team []Credentials
	dev_team = append(dev_team, cred1)
	dev_team = append(dev_team, cred2)

	newCreds, err := json.MarshalIndent(dev_team, "", "  ")
	if err != nil {
		log.Println("Error during marshal: ", err)
	}
	log.Println(string(newCreds))
}
