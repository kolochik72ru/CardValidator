package utils

import (
	"strconv"
	"strings"
)

// Функция для проверки корректности номера с помощью алгоритма Луна
func luhnCheck(number string) bool {
	sum := 0
	parity := len(number) % 2

	for i, char := range number {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false // Если символ не является цифрой, возвращаем false
		}

		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit = digit - 9
			}
		}

		sum += digit
	}

	return sum%10 == 0
}

func getPaymentSystem(cardNumber string) string {
	if strings.HasPrefix(cardNumber, "4") {
		return "Visa"
	} else if strings.HasPrefix(cardNumber, "5") {
		return "MasterCard"
	} else if strings.HasPrefix(cardNumber, "34") || strings.HasPrefix(cardNumber, "37") {
		return "American Express"
	} else if strings.HasPrefix(cardNumber, "6") {
		return "Discover"
	} else if strings.HasPrefix(cardNumber, "35") {
		return "JCB"
	} else if strings.HasPrefix(cardNumber, "22") {
		return "Mir"
	} else if strings.HasPrefix(cardNumber, "30") || strings.HasPrefix(cardNumber, "36") || strings.HasPrefix(cardNumber, "38") {
		return "Diners Club"
	} else {
		return "Unknown payment system"
	}
}

func CardCheck(cardNumber string) (string, bool) {

	paymentSystem := getPaymentSystem(cardNumber)

	if luhnCheck(cardNumber) {
		return paymentSystem, true
	} else {
		return paymentSystem, false
	}
}
