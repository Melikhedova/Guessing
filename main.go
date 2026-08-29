package main

import (
	"fmt"
	"math/rand"
	"math"
	"github.com/fatih/color"
	"time"
	"encoding/json"
	"os"
)

var Name string
var Level int
var Num int
var Limit int
var Max int
var Res string
var Attempt_counter int

type Result struct {
	Name string
	Data string
	Res string
	Level string
	Count int
}

func init() {
	fmt.Println("\x1b[32mДобро поаловать\x1b[0m в игру \x1b[33mУГАДАЙКА\x1b[0m 🤓")
		
	GetName()
}

func main() {
	
	defer Bye()

	for {

		GetLevel()

		CreateNum()
	
		StartGame()

		ToJason()

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
	color.RGB(154, 205, 50).Println("1. Easy: число от 1 до 50 и  утебя будет 15 попыток")
	color.RGB(240, 230, 140).Println("2. Medium: число от 1 до 100 и у тебя будет 10 попыток")
	color.RGB(139, 0, 0).Println("3. Hard: чесло от 1 до 200 и у тебя будет 5 попыток")
	
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
	Attempt_counter = 0
	
	
	for Attempt_counter < Limit {
		
		Attempt_counter++

		fmt.Printf("Попытка №\x1b[31m%d\x1b[0m. Введите число: ", Attempt_counter)
		fmt.Scan(&user_answer)

		all_user_answers = append(all_user_answers, user_answer)

		right = CheckAnswer(user_answer)

		if right {
			Res = "Выиграл"
			color.Green("Ты победил 🥳")
			fmt.Printf("Количество попыток: %d.", Attempt_counter)
			fmt.Println()
			fmt.Printf("Твои ответы: %v", all_user_answers)
			fmt.Println()
			fmt.Println()
			break
		}
	}

	if !right{
		Res = "Проиграл"
		fmt.Printf("%s , попытки \x1b[31mзакончились\x1b[0m", Name)
		fmt.Println()
		fmt.Printf("Вот твои ответы: %v", all_user_answers)
		fmt.Println()
		fmt.Printf("Правильный ответ: %d", Num)
		fmt.Println()
	}
}

func CheckAnswer(answer int) bool {

	difference  := float64(answer - Num)

	switch {
	case answer == Num:
		fmt.Println("Совершенно \x1b[32mверно\x1b[0m!!!😺")
	case int(math.Abs(difference)) <= 5:
		color.Red("Горячо 🥵")
	case int(math.Abs(difference)) <= 15:
		color.Yellow("Тепло 🤭")
	default:
		color.Blue("Хлолдно 🥶")
	}

	if answer > Num {
		fmt.Println("Загаданное число \x1b[4mменьше\x1b[0m 👇")
		fmt.Println()
	} else if answer < Num {
		fmt.Println("Загаданное число \x1b[1m\x1b[4mбольше\x1b[0m 👆")
		fmt.Println()
	}

	return answer == Num
}

func Continue() bool {
	
	fmt.Printf("%s, хочешь поиграть ещё? (да/нет)  ", Name)
	var contin string
	fmt.Scan(&contin)
	if contin != "да"{return false}

	return true
}
	
func Bye(){
	fmt.Println("Хорошо!")
	fmt.Println("Пока 🤧")
}

func ToJason() {
	data := time.Now().Format("02.01.2006 15:04")

	var level string
	switch Level{
	case 1: level = "Легкий"
	case 2: level = "Следний"
	case 3: level = "Сложный"
	}

	out := Result{Name, data, Res, level, Attempt_counter}

	file, err := os.OpenFile("resalt.jsonl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err == nil {
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(out)
	}
}

