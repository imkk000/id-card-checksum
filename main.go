package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("fname :", randomName())
	fmt.Println("lname :", randomName())
	fmt.Println("idcard:", randomIDCard())
	fmt.Println("phone :", randomPhoneNumber())
	fmt.Println("email :", randomEmail())
}

func randomName() string {
	return randomLowerLetter(5 + rand.Intn(7))
}

func randomEmail() string {
	return fmt.Sprintf("%s@%s.com", randomLowerLetter(5+rand.Intn(7)), randomLowerLetter(5))
}

func randomLowerLetter(n int) string {
	var sb strings.Builder
	for range n {
		sb.WriteRune(rune(rand.Intn(26)) + 'a')
	}
	return sb.String()
}

func randomIDCard() string {
	var sb strings.Builder
	sum := 13
	sb.WriteRune('1')
	for i := 12; i >= 2; i-- {
		n := rand.Intn(10)
		sum += n * i
		sb.WriteRune(rune(n) + '0')
	}
	sum = (11 - (sum % 11)) % 10
	sb.WriteRune(rune(sum) + '0')
	return sb.String()
}

func randomPhoneNumber() string {
	var sb strings.Builder
	sb.WriteString("08")
	for range 8 {
		sb.WriteRune(rune(rand.Intn(10)) + '0')
	}
	return sb.String()
}
