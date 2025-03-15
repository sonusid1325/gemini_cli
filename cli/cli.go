package cli

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Init() {
	var text string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What's your favorite programming language?").
				Value(&text),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You chose: %s\n", text)
}
