package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Bold    = "\033[1m"
)

var wordsEasy = []string{"chat", "chien", "maison", "pomme", "livre"}
var wordsMedium = []string{"ordinateur", "programmation", "internet", "fichier", "logiciel"}
var wordsHard = []string{"algorithme", "cryptographie", "concurrentiel", "synchronisation", "virtualisation"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		printWithColor("=== Menu du Pendu ===", Cyan, true)
		printWithColor("1. Jouer", Green, false)
		printWithColor("2. Quitter", Red, false)
		fmt.Print("Choisissez une option : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			selectDifficulty(scanner)
		case "2":
			printWithColor("Merci d'avoir joué ! À bientôt.", Magenta, true)
			return
		default:
			printWithColor("Option invalide, veuillez réessayer.", Red, true)
		}
	}
}

func selectDifficulty(scanner *bufio.Scanner) {
	clearScreen()
	printWithColor("=== Sélection du niveau de difficulté ===", Cyan, true)
	printWithColor("1. Facile", Green, false)
	printWithColor("2. Moyen", Yellow, false)
	printWithColor("3. Difficile", Red, false)
	fmt.Print("Choisissez un niveau : ")
	scanner.Scan()
	levelChoice := scanner.Text()

	switch levelChoice {
	case "1":
		play(scanner, wordsEasy)
	case "2":
		play(scanner, wordsMedium)
	case "3":
		play(scanner, wordsHard)
	default:
		printWithColor("Option invalide, veuillez réessayer.", Red, true)
		selectDifficulty(scanner)
	}
}

func play(scanner *bufio.Scanner, wordList []string) {
	word := pickRandomWord(wordList)
	guesses := ""
	remainingAttempts := 9

	for remainingAttempts > 0 {
		clearScreen()
		displayHangman(remainingAttempts)
		displayWord(word, guesses)
		if isWordComplete(word, guesses) {
			printWithColor("Félicitations, vous avez trouvé le mot !", Green, true)
			break
		}

		fmt.Printf("\nTentatives restantes : %s%d%s\n", Yellow, remainingAttempts, Reset)
		fmt.Print("Entrez une lettre : ")
		scanner.Scan()
		letter := strings.ToLower(scanner.Text())

		if len(letter) != 1 || !isLetter(letter) {
			printWithColor("Entrée invalide, veuillez entrer une seule lettre.", Red, true)
			continue
		}

		if strings.Contains(guesses, letter) {
			printWithColor("Vous avez déjà deviné cette lettre.", Yellow, true)
			continue
		}

		guesses += letter

		if !strings.Contains(word, letter) {
			printWithColor("Mauvaise lettre !", Red, true)
			remainingAttempts--
		} else {
			printWithColor("Bonne lettre !", Green, true)
		}
	}

	if remainingAttempts == 0 {
		clearScreen()
		displayHangman(remainingAttempts)
		printWithColor(fmt.Sprintf("Désolé, vous avez perdu. Le mot était : %s", word), Red, true)
	}
}

func pickRandomWord(wordList []string) string {
	rand.Seed(time.Now().UnixNano())
	return wordList[rand.Intn(len(wordList))]
}

func displayWord(word, guesses string) {
	for _, letter := range word {
		if strings.ContainsRune(guesses, letter) {
			fmt.Printf("%s%c%s ", Green, letter, Reset)
		} else {
			fmt.Printf("%s_ %s", Yellow, Reset)
		}
	}
	fmt.Println()
}

func isWordComplete(word, guesses string) bool {
	for _, letter := range word {
		if !strings.ContainsRune(guesses, letter) {
			return false
		}
	}
	return true
}

func isLetter(input string) bool {
	if len(input) != 1 {
		return false
	}
	r := rune(input[0])
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func displayHangman(attemptsLeft int) {
	states := []string{

		Red + `
		` + Reset,

		Red + `

			
			_________
				  
		` + Reset,
		Red + `

			|       
			|       
			|       
			|       
			|________ 
			  
		` + Reset,
		Red + `
			_________
			|       
			|       
			|       
			|       
			|________
		` + Reset,
		Yellow + `

			_________
			|       |
			|       O
			|       
			|       
			|________
		   	   
		` + Reset,
		Yellow + `
		
			_________
			|       |
			|       O
			|       |
			|       
			|________

		` + Reset,
		Yellow + `
			   
			_________
			|       |
			|       O
			|      /|
			|       
			|________

		` + Reset,
		Blue + `
		  	_________
			|       |
			|       O
			|      /|\
			|       
			|________
	  
		` + Reset,
		Cyan + `	
			_________
			|       |
			|       O
			|      /|\
			|      / 
			|________ 
		` + Reset,
		White + `
			_________
			|       |
			|       O
			|      /|\
			|      / \
			|________ 	
		
	 
		` + Reset,
	}

	fmt.Println(states[9-attemptsLeft])
}

func printWithColor(text, color string, bold bool) {
	if bold {
		fmt.Println(Bold + color + text + Reset)
	} else {
		fmt.Println(color + text + Reset)
	}
}
