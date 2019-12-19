package main

import (
	"fmt"
	"os"

	"chatbot-notifier/internal/option"
)

func main() {
	if err := option.Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
