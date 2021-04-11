package main
import ("fmt" 
		"math/rand" 
		"strconv"
	)

func main(){
	fmt.Println("Hello, user. what is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Println("welcome " + name + " would you like to play a game?")
	var answer string
	fmt.Scanln(&answer)
	if answer == "yes" {
		fmt.Println("ok, let's play a game!")
		game()
	} else {
		fmt.Println("ok, maybe later then. bye!")
	}
}

func game() {
	var active bool= true
	for active {
		var correctAnswer int = rand.Intn(10)
		var scorrectAnswer = strconv.Itoa(correctAnswer)
		fmt.Println("guess a number between 1 and 10")
		var guess int
		fmt.Scanln(&guess)
		var sguess = strconv.Itoa(guess)
		if guess == correctAnswer {
			fmt.Println("you got it! " + scorrectAnswer + " was the right number!")
			break;
		} else {
			fmt.Println("sorry, " + sguess + " is not the right number!")
		}
	}
}
