package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//* Обработка паники
	defer func() { // defer - отложенный вызов
		r := recover() // обработка паники

		if r != nil {
			fmt.Println("Recover", r)
		}
	}() // вызов неименованной функции

	fmt.Println("___Калькулятор индекса массы тела___")

	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			// fmt.Println(err)
			panic("Не заданы параметры для расчёта")
		}
		outputResult(IMT)

		isRepeatCalculation := checkRepeatCalculation()

		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	show := fmt.Sprintf("Ваш индекс массы тела: %.0f\n", imt)
	fmt.Print(show)

	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userKg, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}

	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	return  IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	
	fmt.Print("Введите свой рост (в сантиметрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)

	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string

	fmt.Print("Вы хотите сделать ещё расчёт (y/n): ")
	fmt.Scan(&userChoice)

	if userChoice == "y" || userChoice == "Y" {
		return  true
	}

	return false
}
