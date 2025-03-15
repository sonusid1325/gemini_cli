package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term" // Correct package for getting terminal size
)

func Init() {
	var text string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("prompt").
				Value(&text),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(text) == "bye" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Println("Error getting terminal size:", err)
		width = 80
	}

	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Width(width - 4).
		MaxWidth(width - 200)

	output := fmt.Sprintf("Gemini:\n%s", text)
	borderedOutput := style.Render(output)

	fmt.Println(borderedOutput)
}
