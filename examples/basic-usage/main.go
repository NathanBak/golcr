package main

import (
	"fmt"
	"log"

	golcr "github.com/NathanBak/golcr/pkg/client"
)

const (
	user = "PUT_USERNAME_HERE"
	pass = "PUT_PASSWORD_HERE"
)

func main() {

	myCreds := &creds{username: user, password: pass}

	lcrClient, err := golcr.NewClient(myCreds)
	if err != nil {
		log.Fatal(err)
	}

	orgs, err := lcrClient.GetOrgs()
	if err != nil {
		log.Fatal(err)
	}

	for _, org := range *orgs {
		fmt.Printf("%v (%v)", org.Name, org.SubOrgID)
	}
}

type creds struct {
	username string
	password string
}

func (c *creds) GetUserPass() (string, string) {
	return c.username, c.password
}
