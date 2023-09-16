package main

import "fmt"

func main() {
	fmt.Println("Welcome to my game")

	/*name := "Tim"
	age := 21
	fmt.Printf("Hello %v, you are %v, how are you", name, age) */

	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name)
	fmt.Printf("Hello %v, welcome to the game \n", name)

	fmt.Printf("Enter your age: \n")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("you can play")
	} else {
		fmt.Println("you can't play")
		return
	}
	score := 0
	num_questions := 2

	fmt.Printf("What is better, php or python?")
	var answer string
	fmt.Scan(&answer)
	fmt.Println(&answer)

	if answer == "php" {
		fmt.Println("maybe yes, maybe not")
		score++
	} else {
		fmt.Println("maybe yes, maybe not again")
	}

	fmt.Printf("Does it laravel is good for frontend? ")
	var answer1 string
	fmt.Scan(&answer1)
	if answer1 == "yes" {
		fmt.Println("correct")
		score++
	} else if answer1 == "no" {
		fmt.Println("incorrect")
	}
	fmt.Printf("You scored %v out of %v \n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored  %v%%.", percent)
}
