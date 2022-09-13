package main

import (
	"user"
)

func main() {
	var userObserver = *user.CreateObserver()

	user.CreateUser("João", userObserver)
}
