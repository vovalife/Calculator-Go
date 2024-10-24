package main

import "fmt"

func main() {
	// объявление констант
	const (
		a1 = "*"
		a2 = "/"
		a3 = "-"
		a4 = "+"
		a5 = "%"
	)
	// объявление перемен
	var a, b float32
	var h string

	fmt.Scan(&a, &b, &h)

	switch h {
	case a1:
		fmt.Println(a * b)
	case a2:
		if b == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Println(a / b)
		}
	case a3:
		fmt.Println(a - b)
	case a4:
		fmt.Println(a + b)
	case a5:
		fmt.Println(int(a) % int(b))

	}

}
