package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Баженов Дмитрий Александрович")
	// calculator(1.35, 0.98)
	// calculatorArray(1.35, 0.98)

	fox := NewFox("Алиса", 4, "рыжий")
	fmt.Println(fox.GetView())
	fmt.Println(fox.Hunt())
	fmt.Println(fox.Speak())
}

// // func calculator(a, b float64) {
// // 	fmt.Println("x\tЧислитель\tЗнаменатель\tРезультат")
// // 	fmt.Println("-----------------------------------------------")

// // 	for i := 1.14; i < 4.24; i += 0.62 {
// // 		numerator := math.Pow((a*i)+b, 1.0/3)
// // 		denominator := math.Pow(math.Log(i), 2)
// // 		result := numerator / denominator

// // 		fmt.Printf("%.2f\t%.6f\t%.6f\t%.6f\n", i, numerator, denominator, result)
// // 	}
// // }

// func calculatorArray(a, b float64) {
// 	var arrX [5]float64 = [5]float64{0.35, 1.28, 3.51, 5.21, 4.16}
// 	fmt.Println("x\tЧислитель\tЗнаменатель\tРезультат")
// 	fmt.Println("-----------------------------------------------")

// 	for i := 0; i < len(arrX); i++ {
// 		numerator := math.Pow((a*arrX[i])+b, 1.0/3)
// 		denominator := math.Pow(math.Log(arrX[i]), 2)
// 		result := numerator / denominator
// 		fmt.Printf("%.2f\t%.6f\t%.6f\t%.6f\n", arrX[i], numerator, denominator, result)
// 	}
// }

type Fox struct {
	Name  string
	Age   int
	Color string
}

func NewFox(name string, age int, color string) Fox {
	return Fox{
		Name:  name,
		Age:   age,
		Color: color,
	}
}

func (f Fox) GetView() string {
	return `
    /\   /\
   /  \_/  \
  /         \
 |   \   /   |
  \_       _/
    |     |
    |     |
` + fmt.Sprintf("Лиса %s\nВозраст: %d\nОкрас: %s", f.Name, f.Age, f.Color)
}

func (f Fox) Hunt() string {
	return fmt.Sprintf("%s тихо крадётся за добычей...", f.Name)
}

func (f Fox) Speak() string {
	return "Фыр-фыр! Что-то говорит лиса " + f.Name
}
