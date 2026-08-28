package main

import (
	"fmt"
	"math/rand"
	"math"
)

var Name string
var Level int
var Num int
var Limit int
var Max int


func init() {
	fmt.Println("Добро поаловать в игру УГАДАЙКА 🤓")
		
	GetName()
}

func main() {
	
	for {

		GetLevel()

		CreateNum()
	
		StartGame()

		if !Continue() {break}
		
	}
}

func GetName() {
	
	fmt.Println("Как я могу к тебе обращаться?")
	fmt.Print("Введите имя: ")
	
	fmt.Scan(&Name)
	
	fmt.Println()
}

func GetLevel() {
	
	fmt.Println("В нашей игре есть три уровня:")
	fmt.Println("1. Easy: число от 1 до 50 и  утебя будет 15 попыток")
	fmt.Println("2. Medium: число от 1 до 100 и у тебя будет 10 попыток")
	fmt.Println("3. Hard: чесло от 1 до 200 и у тебя будет 5 попыток")
	
	fmt.Print("Какой уровень выберешь ты? (введи число 1/2/3) ")
	
	fmt.Scan(&Level)

}

func CreateNum() {
	var level_limit = map[int]int {
		1: 50,
		2: 100,
		3: 200,
	}

	Max = level_limit[Level]

	Num = rand.Intn(Max)

	fmt.Println()
	fmt.Println("Число загадано!")
	fmt.Println("Приступим")

} 

func StartGame() {

	var level_limit = map[int]int {
		1: 15,
		2: 10,
		3: 5,
	}

	Limit = level_limit[Level]

	var all_user_answers []int
	var user_answer int
	var right bool
	attempt_counter := 0
	
	
	for attempt_counter < Limit {
		
		attempt_counter++

		fmt.Printf("Попытка №%d. Введите число: ", attempt_counter)
		fmt.Scan(&user_answer)

		all_user_answers = append(all_user_answers, user_answer)

		right = CheckAnswer(user_answer)

		if right {
			fmt.Println("Ты победил 🥳")
			fmt.Printf("Количество попыток: %d.", attempt_counter)
			fmt.Println()
			fmt.Printf("Твои ответы: %v", all_user_answers)
			break
		}
	}

	if !right{
		fmt.Printf("%s , попытки закончились", Name)
		fmt.Println()
		fmt.Printf("Вот твои ответы: %v", all_user_answers)
		fmt.Println()
		fmt.Printf("Правильный ответ: %d", Num)
	}
}

func CheckAnswer(answer int) bool {

	difference  := float64(answer - Num)

	switch {
	case answer == Num:
		fmt.Println("Совершенно верно!!!😺")
	case int(math.Abs(difference)) <= 5:
		fmt.Println("Горячо 🥵")
	case int(math.Abs(difference)) <= 15:
		fmt.Println("Тепло 🤭")
	default:
		fmt.Println("Хлолдно 🥶")
	}

	if answer > Num {
		fmt.Println("Загаданное число меньше 👇")
		fmt.Println()
	} else if answer < Num {
		fmt.Println("Загаданное число больше 👆")
		fmt.Println()
	}

	return answer == Num
}

func Continue() bool {
	
	fmt.Printf("%s, хочешь поиграть ещё? (да/нет)", Name)
	var contin string
	fmt.Scan(&contin)
	
	return true
}
	
