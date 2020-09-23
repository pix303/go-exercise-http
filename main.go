package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	resp, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		return err
	}

	dataResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var users []User
	err = json.Unmarshal(dataResp, &users)
	if err != nil {
		return err
	}

	for _, u := range users {
		fmt.Println(u)
	}

	dataToSave, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data.json", dataToSave, 0777)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id       int
	Name     string
	Username string
	Address  Address
}

type Address struct {
	Street string
	Suite  string
	Loc    GeoLoc `json:"geo"`
}

type GeoLoc struct {
	Lat string
	Lng string
}

func (u User) ToString() string {
	return fmt.Sprintf("ID: %d \t %s\t\t(%s)", u.Id, u.Name, u.Username)
}
