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
	var active bool = true
	var hint bool = false
	var guessesLeft int = 8
	var correctAnswer int = rand.Intn(20)
	var scorrectAnswer = strconv.Itoa(correctAnswer)
	for active {
		var sguessesleft = strconv.Itoa(guessesLeft)
		fmt.Println("guess a number between 1 and 20")
		var guess int
		fmt.Scanln(&guess)
		var sguess = strconv.Itoa(guess)
		if guess == correctAnswer {
			fmt.Println("you got it! " + scorrectAnswer + " was the right number!")
			break;
		} else {
			if hint == true {
				if guess < correctAnswer{
					fmt.Println("that's not it. the number is higher. guess again.")
				}
				if guess > correctAnswer {
					fmt.Println("that's not it. the number is lower. guess aagin")
				}
			}
			if guessesLeft == 0 {
				fmt.Println("sorry, you didn't guess the number in time. you've lost.")
				break;
			}
			if guessesLeft <= 3 && hint == false{
				fmt.Println("you're alomost out of turns. would you like a hint?")
				var activateHints string
				fmt.Scanln(&activateHints)
				if activateHints == "yes"{
					hint = true
				} else {
					fmt.Println("okay, good luck")
				}
			}
			guessesLeft -= 1
			fmt.Println("sorry, " + sguess + " is not the right number! You have " + sguessesleft + " guesses left")
		}
	}
}
