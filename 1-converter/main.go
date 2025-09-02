package main

import (
	"fmt"
	"strings"
)

// Доступные валюты для конвертации
var availableCurrencies = []string{"USD", "EUR", "RUB", "GBP", "JPY"}

// Map для быстрой проверки доступности валют
var availableCurrenciesMap = map[string]bool{
	"USD": true,
	"EUR": true,
	"RUB": true,
	"GBP": true,
	"JPY": true,
}

// Курсы валют к USD (базовая валюта)
var currencyRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.85,
	"RUB": 75.0,
	"GBP": 0.73,
	"JPY": 110.0,
}

func main() {
	showMenu()

	// Ввод исходной валюты
	fromCurrency := inputCurrency("исходной")

	// Ввод суммы
	amount := inputAmount()

	// Ввод целевой валюты
	toCurrency := inputCurrency("целевой")

	// Расчёт и вывод результата
	result := convertCurrency(amount, fromCurrency, toCurrency)
	fmt.Printf("\nРезультат конвертации: %.2f %s = %.2f %s\n",
		amount, fromCurrency, result, toCurrency)
}

// showMenu отображает приветствие и доступные валюты
func showMenu() {
	fmt.Println("=== Калькулятор валют ===")
	fmt.Println("Доступные валюты:")
	for _, currency := range availableCurrencies {
		fmt.Printf("- %s\n", currency)
	}
	fmt.Println()
}

// inputCurrency запрашивает ввод валюты с валидацией
func inputCurrency(currencyType string) string {
	for {
		fmt.Printf("Введите %s валюту (например, USD): ", currencyType)
		var input string
		fmt.Scan(&input)

		// Приводим к верхнему регистру для унификации
		input = strings.ToUpper(strings.TrimSpace(input))

		// Проверяем, что введённая валюта доступна через map
		if availableCurrenciesMap[input] {
			return input
		}

		fmt.Printf("Ошибка: валюта '%s' не поддерживается. Попробуйте снова.\n", input)
		fmt.Println("Доступные валюты:", strings.Join(availableCurrencies, ", "))
	}
}

// inputAmount запрашивает ввод суммы с валидацией
func inputAmount() float64 {
	for {
		fmt.Print("Введите сумму для конвертации: ")
		var input string
		fmt.Scan(&input)

		// Пытаемся преобразовать в число
		var amount float64
		_, err := fmt.Sscanf(input, "%f", &amount)

		if err != nil || amount <= 0 {
			fmt.Println("Ошибка: введите корректное положительное число.")
			continue
		}

		return amount
	}
}

// convertCurrency рассчитывает конвертацию валют
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	// Если валюты одинаковые, возвращаем исходную сумму
	if fromCurrency == toCurrency {
		return amount
	}

	// Проверяем, что обе валюты существуют в map курсов
	fromRate, fromExists := currencyRates[fromCurrency]
	toRate, toExists := currencyRates[toCurrency]

	if !fromExists || !toExists {
		// В реальном приложении здесь должно быть логирование ошибки
		// или возврат ошибки, но для простоты возвращаем 0
		return 0
	}

	// Конвертируем через USD (базовая валюта)
	// Сначала в USD, затем в целевую валюту
	usdAmount := amount / fromRate // в USD
	result := usdAmount * toRate   // в целевую валюту

	return result
}
