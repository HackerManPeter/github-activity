package internal

import (
	"flag"
	"fmt"
)

type Arguments struct {
	Limit    uint
	Username string
}

func ReceiveArguments() (*Arguments, error) {
	usernamePtr := flag.String("username", "", "User's github username")
	limitPtr := flag.Uint("limit", 5, "Number of events per page")

	flag.Parse()

	if *usernamePtr == "" {
		return nil, fmt.Errorf("github username must be passed")
	}

	return &Arguments{
		Limit:    *limitPtr,
		Username: *usernamePtr,
	}, nil

}
