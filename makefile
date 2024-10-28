# Makefile for Chat Assistant in Go using Fyne

# Nazwa projektu
APP_NAME = chat-assistant

# Ścieżka do głównego źródła
SRC = pkg/main.go

# Flagi kompilatora
GO_FLAGS = -v

.PHONY: all clean build run deps

# Domyślny cel - kompilacja i uruchomienie aplikacji
all: deps build run

# Zainstalowanie zależności
deps:
	@echo "Pobieranie zależności..."
	go mod tidy

# Kompilacja aplikacji
build:
	@echo "Budowanie aplikacji..."
	go build $(GO_FLAGS) -o $(APP_NAME) $(SRC)

# Uruchomienie aplikacji
run: build
	@echo "Uruchamianie aplikacji..."
	./$(APP_NAME)

# Czyste usunięcie plików binarnych
clean:
	@echo "Czyszczenie plików..."
	rm -f $(APP_NAME)
