package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gerador de Senhas")
	w.Resize(fyne.NewSize(400, 200))

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Sua senha aparecerá aqui")

	generateButton := widget.NewButton("Gerar Senha", func() {
		password := generatePassword(12)
		entry.SetText(password)
	})

	w.SetContent(container.NewVBox(
		entry,
		generateButton,
	))

	w.ShowAndRun()
}

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
		special[rand.Intn(len(special))],
	}

	password := make([]byte, length-len(mandatory))
	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	password = append(password, mandatory...)
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}
