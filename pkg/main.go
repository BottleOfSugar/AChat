package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Funkcja do generowania odpowiedzi asystenta
func getAssistantResponse(input string) string {
	lowerInput := strings.ToLower(input)
	switch {
	case strings.Contains(lowerInput, "hej"):
		return "Hej! Jak mogę Ci pomóc?"
	case strings.Contains(lowerInput, "pogoda"):
		return "Dzisiaj jest piękna pogoda, świeci słońce!"
	case strings.Contains(lowerInput, "czas"):
		return "Aktualnie nie znam czasu, ale mogę pomóc w innych sprawach!"
	default:
		return "Nie jestem pewien, jak odpowiedzieć na to pytanie."
	}
}

func main() {
	// Tworzenie aplikacji i okna w rozdzielczości HD
	a := app.New()
	w := a.NewWindow("AChat 1.0 BETA ")
	w.Resize(fyne.NewSize(1920, 1080))

	// Pole do wpisywania wiadomości
	input := widget.NewEntry()
	input.SetPlaceHolder("Wpisz swoją wiadomość...")
	input.Resize(fyne.NewSize(1600, 40)) // Ustawienie szerokości pola

	// Pole tekstowe do wyświetlania czatu
	chatContent := widget.NewLabel("Asystent: Witaj! Jak mogę pomóc?\n")

	// Kontener na wiadomości
	chatContainer := container.NewVBox(chatContent)

	// Przycisk "Wyślij" z funkcją obsługującą wiadomości użytkownika
	sendButton := widget.NewButton("Wyślij", func() {
		userMessage := input.Text
		if userMessage == "" {
			return
		}
		chatContainer.Add(widget.NewLabel(fmt.Sprintf("Ty: %s", userMessage)))
		assistantResponse := getAssistantResponse(userMessage)
		chatContainer.Add(widget.NewLabel(fmt.Sprintf("Asystent: %s", assistantResponse)))

		input.SetText("") // Wyczyść pole
	})

	// Rozmieszczenie elementów z szerszym polem do wpisywania wiadomości
	w.SetContent(container.NewBorder(nil, container.NewHBox(input, sendButton), nil, nil, chatContainer))

	w.ShowAndRun()
}
