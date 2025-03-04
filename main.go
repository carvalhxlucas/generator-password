package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	password := generatePassword(12)
	fmt.Println(password)
}

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	allChars := lowerCase + upperCase + numbers + special

	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
		special[rand.Intn(len(numbers))],
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
