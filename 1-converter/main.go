package main

import "fmt"

func main() {
	// Считываем ввод пользователя: базовая валюта, целевая валюта и сумма
	baseCurrency, targetCurrency, amount := readUserInput()

	// Вызываем заглушку функции конвертации (реализация будет добавлена позже)
	result := convertCurrency(amount, baseCurrency, targetCurrency)

	// Выводим результат конвертации
	fmt.Println(result)
}

// readUserInput считывает из консоли исходную валюту, целевую валюту и сумму
func readUserInput() (string, string, float64) {
	var baseCurrency string
	var targetCurrency string
	var amount float64
	fmt.Println("Введите базовую валюту: ")
	fmt.Scan(&baseCurrency)
	fmt.Println("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	fmt.Println("Введите сумму: ")
	fmt.Scan(&amount)
	return baseCurrency, targetCurrency, amount
}

// convertCurrency — заглушка функции расчёта конвертации валют
// Принимает сумму и коды валют (исходная и целевая), возвращает рассчитанную сумму
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	// TODO: Реализовать конвертацию на основе курсов
	return 0
}
