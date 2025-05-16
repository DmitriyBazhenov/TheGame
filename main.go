package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Баженов Дмитрий Александрович")
	calculator(1.35, 0.98)
}

func calculator(a, b float64) {
	fmt.Println("x\tЧислитель\tЗнаменатель\tРезультат")
	fmt.Println("-----------------------------------------------")

	for i := 1.14; i < 4.24; i += 0.62 {
		numerator := math.Pow((a*i)+b, 1.0/3)
		denominator := math.Pow(math.Log(i), 2)
		result := numerator / denominator

		fmt.Printf("%.2f\t%.6f\t%.6f\t%.6f\n", i, numerator, denominator, result)
	}
}
