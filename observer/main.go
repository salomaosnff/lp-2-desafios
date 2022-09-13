package main

import (
	"fmt"
	"observable"
)

func main() {
	var messages = observable.Create()

	messages.Subscribe(func(message any) {
		fmt.Printf("Sub1 %s\n", message)
	})

	messages.Subscribe(func(message any) {
		fmt.Printf("Sub2 %s\n", message)
	})

	messages.Publish("Opa meu Chapa!")
}
