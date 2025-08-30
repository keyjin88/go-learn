package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Функция для вычисления среднего значения
func calculateAverage(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Функция для вычисления суммы
func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Функция для вычисления медианы
func calculateMedian(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Создаем копию слайса для сортировки
	sortedNumbers := make([]float64, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Float64s(sortedNumbers)

	length := len(sortedNumbers)
	if length%2 == 0 {
		// Четное количество элементов - берем среднее двух центральных
		mid := length / 2
		return (sortedNumbers[mid-1] + sortedNumbers[mid]) / 2
	} else {
		// Нечетное количество элементов - берем центральный
		return sortedNumbers[length/2]
	}
}

// Функция для парсинга строки чисел
func parseNumbers(numbersStr string) ([]float64, error) {
	// Убираем пробелы и разбиваем по запятым
	numbersStr = strings.TrimSpace(numbersStr)
	if numbersStr == "" {
		return []float64{}, nil
	}

	parts := strings.Split(numbersStr, ",")
	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("не удалось преобразовать '%s' в число: %v", part, err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

// Функция для отображения главного меню
func showMenu() {
	fmt.Println("\n=== КАЛЬКУЛЯТОР СТАТИСТИЧЕСКИХ ОПЕРАЦИЙ ===")
	fmt.Println("1. AVG - Вычислить среднее значение")
	fmt.Println("2. SUM - Вычислить сумму")
	fmt.Println("3. MED - Вычислить медиану")
	fmt.Println("4. Выход")
	fmt.Print("Выберите операцию (1-4): ")
}

// Функция для получения чисел от пользователя
func getNumbersFromUser() ([]float64, error) {
	fmt.Print("Введите числа через запятую (например: 2,10,9): ")
	
	var input string
	fmt.Scanln(&input)
	
	numbers, err := parseNumbers(input)
	if err != nil {
		return nil, err
	}
	
	if len(numbers) == 0 {
		return nil, fmt.Errorf("не указано ни одного числа")
	}
	
	return numbers, nil
}

// Функция для выполнения вычислений
func performCalculation(operation string, numbers []float64) {
	fmt.Printf("\nВведенные числа: %v\n", numbers)
	
	var result float64
	switch operation {
	case "AVG":
		result = calculateAverage(numbers)
		fmt.Printf("Среднее значение: %.2f\n", result)
	case "SUM":
		result = calculateSum(numbers)
		fmt.Printf("Сумма: %.2f\n", result)
	case "MED":
		result = calculateMedian(numbers)
		fmt.Printf("Медиана: %.2f\n", result)
	}
}

func main() {
	fmt.Println("Добро пожаловать в калькулятор статистических операций!")

	for {
		showMenu()

		var choice string
		fmt.Scanln(&choice)
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("\n--- ВЫЧИСЛЕНИЕ СРЕДНЕГО ЗНАЧЕНИЯ ---")
			numbers, err := getNumbersFromUser()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}
			performCalculation("AVG", numbers)

		case "2":
			fmt.Println("\n--- ВЫЧИСЛЕНИЕ СУММЫ ---")
			numbers, err := getNumbersFromUser()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}
			performCalculation("SUM", numbers)

		case "3":
			fmt.Println("\n--- ВЫЧИСЛЕНИЕ МЕДИАНЫ ---")
			numbers, err := getNumbersFromUser()
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}
			performCalculation("MED", numbers)

		case "4":
			fmt.Println("\nСпасибо за использование калькулятора! До свидания!")
			return

		default:
			fmt.Println("Неверный выбор! Пожалуйста, введите число от 1 до 4.")
		}

		// Пауза перед следующим циклом
		fmt.Print("\nНажмите Enter для продолжения...")
		fmt.Scanln() // Очищаем буфер ввода
	}
}
