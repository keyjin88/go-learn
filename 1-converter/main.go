package main

import "fmt"

func main() {
	var eurusdCourse float64 = 1.32
	var rubusdCourse float64 = 82.5
	//Рассчитать курс рубля к евро и вывести в консоль
	var rubeurCourse float64 = rubusdCourse * eurusdCourse
	fmt.Println(rubeurCourse)
}
