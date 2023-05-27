package main

import (
	"TP-ihw4/tpmath"
	"fmt"
)

func main() {
	// 1 - Целые чиселки
	var set1 = []int{1, 0, 1, 1, 0, 3, 1, 4, 1, 3, 2, 4, 0, 0, 5, 3, 1, 1, 1, 4, 2, 1, 0, 0, 1, 2, 5, 2, 6, 3, 6, 2, 4,
		1, 2, 1, 1, 1, 3, 2, 2, 0, 0, 0, 3, 3, 7, 1, 2, 1}
	fmt.Println("Исходная выборка: ", set1)

	tpmath.Sort(set1)
	fmt.Println("Вариационный ряд: ", set1)
	fmt.Println("Эмпирическая функция распределения: ", tpmath.DistributionFunc(set1))

	fmt.Println()
	fmt.Println("Выборочные аналоги числовых карактеристик:")
	// Вычислить выборочные аналоги числовых карактеристик
	fmt.Println("Математическое ожидание: ", tpmath.Expect(set1))
	fmt.Println("Дисперсия: ", tpmath.Variance(set1))
	fmt.Println("Медиана: ", tpmath.Median(set1))
	fmt.Println("Ассиметрия: ", tpmath.Asymmetry(set1))
	fmt.Println("Эксцесса: ", tpmath.Excess(set1))
	fmt.Println("Вероятность Pr(X ∈ [1.72;2.28]):", tpmath.Probability(set1, 1.72, 2.28))

}

func PrettyOutputSet[T tpmath.Number](set []T) {
	var prev = set[0]
	var count = 0
	fmt.Print("[")
	for _, v := range set {
		if v != prev {
			fmt.Print(prev, " * ", count, ", ")
			count = 0
			prev = v
		}
		count++
	}
	fmt.Print(prev, " * ", count, ", ")
	fmt.Println("]")
}
