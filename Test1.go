// Test1 "math/rand" "time" var a int
package main

import (
	"fmt",
	"math/rand"
)

func main() {
	// программа для регистрацыя пользователей на экзамен
	var a int
	var b int             // Выбор тип записи
	var h string          // выбор тип экзамена
	var j string          // обезательные предметы
	var a1, b1, c1 string // необезательные предметы
	var uoi string        // подверждение участие
	var pass1, pass2 int
	var name string
	var city string
	var phone string

	fmt.Println("Добро пожаловать выбирите действие.1 Запись на экзамен,2.Регистрацыя на экзамен")
	fmt.Scan(&b)

	

	if b == 1 {
		fmt.Println("Выберите тип экзамена")
		fmt.Scan(&h)
		fmt.Println("Выбирите основные предметы до 2!")
		fmt.Scan(&j)
		fmt.Println("Выбирите доп предметы до 3!")
		fmt.Scan(&a1, &b1, &c1)
		fmt.Println("Введите ФИО,номер телефона, город")
		fmt.Scan(&name, &phone, &city)
		fmt.Println("Введите паспортные данные")
		fmt.Scan(&pass1, &pass2)
		fmt.Println("Анкета предметы", j, a1, b1, c1, "Тип эзамена", h, "Личные данные", name, phone, city, pass1, pass2)
		fmt.Println("Подвердите анкету")
		fmt.Scan(&uoi)
		if uoi == "Да" {
			fmt.Println("Регистрацыя прошло успешно ожидаем вас в пунктах здачи экзамена")
		} else {
			fmt.Println("Перезаполните данные")
			fmt.Println("Выберите тип экзамена")
			fmt.Scan(&h)
			fmt.Println("Выбирите основные предметы до 2!")
			fmt.Scan(&j)
			fmt.Println("Выбирите доп предметы до 3!")
			fmt.Scan(&a1, &b1, &c1)
			fmt.Println("Введите ФИО,номер телефона, город")
			fmt.Scan(&name, &phone, &city)
			fmt.Println("Введите паспортные данные")
			fmt.Scan(&pass1, &pass2)
			fmt.Println("Анкета предметы", j, a1, b1, c1, "Тип эзамена", h, "Личные данные", name, phone, city, pass1, pass2)
		}

	}

}
