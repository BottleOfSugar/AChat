# Asystent Czatowy w Go z użyciem Fyne

Ten projekt jest prostym asystentem czatu stworzonym w języku Go z wykorzystaniem biblioteki **Fyne** do tworzenia GUI. Aplikacja uruchamia okno w rozdzielczości HD (1920x1080) z szerokim paskiem do wpisywania wiadomości oraz przyciskiem "Wyślij". Na podstawie wprowadzonych danych użytkownika, asystent generuje odpowiedzi na podstawie kilku predefiniowanych tematów.

## Funkcje

- **Interfejs czatu**: Umożliwia użytkownikowi komunikację z asystentem za pomocą prostego interfejsu tekstowego.
- **Odpowiedzi asystenta**: Asystent rozpoznaje kilka tematów, takich jak "Hej", "pogoda" i "czas" i generuje odpowiedzi zależnie od wpisanej frazy.
- **Responsywne pole wiadomości**: Pasek do wpisywania wiadomości jest szeroki, co pozwala na komfortowe wpisywanie dłuższych wiadomości.

## Jak uruchomić

1. Upewnij się, że masz zainstalowane [Go](https://golang.org/doc/install) oraz [Fyne](https://developer.fyne.io/).
2. Skopiuj kod źródłowy do pliku `.go`.
3. Uruchom aplikację poleceniem:
   ```bash
   go run nazwapliku.go
