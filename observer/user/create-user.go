package user

import (
	"observable"
)

func CreateUser(name string, obs observable.Observable) {
	var user = &User{
		Name: name,
	}

	obs.Publish(user)
}
