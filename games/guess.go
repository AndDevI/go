package games

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const title string = "Jogo da advinhação"
const instruction string = "Um número entre 0 e 100 será sorteado, tente acertar!!"
const instructionAttempts string = "Digite um número"
const notNumber string = "A sua tentativa não é número"
const greaterNumberDrawn string = "Errou! O número sorteado é maior que: %v\n"
const lessNumberDrawn string = "Errou! O número sorteado é menor que: %v\n"
const gameOver string = "Infelizmente, você não acertou o número que era: %d\n Esses foram suas tentativas: %v\n"
const congratulations string = "Parabéns, você acertou!! O número que era: %d\n Você acertou em: %v tentativas\n"

const attemptsMax int = 10
const quantityDrawn int = 101

func Game() {
	printHeader()

	scanner := bufio.NewScanner(os.Stdin)

	attempts := [attemptsMax]int64{}
	numberDrawn := drawNumber()

	for i := range attempts {
		printInstructionAttempts()

		input, ok := readLine(scanner)
		if !ok {
			return
		}

		attemptInt, ok := parseAttempt(input)
		if !ok {
			printNotNumber()

			return
		}

		if handleResult(attemptInt, numberDrawn, i) {
			return
		}

		attempts[i] = attemptInt
	}

	printGameOver(numberDrawn, attempts)
}

func printHeader() {
	fmt.Println(title)
	fmt.Println(instruction)
}

func printInstructionAttempts() {
	fmt.Println(instructionAttempts)
}

func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}

	return strings.TrimSpace(scanner.Text()), true
}

func parseAttempt(value string) (int64, bool) {
	attempt, err := strconv.ParseInt(value, 10, 64)

	return attempt, err == nil
}

func printNotNumber() {
	fmt.Println(notNumber)
}

func handleResult(attempt, numberDrawn int64, i int) bool {
	switch {
	case attempt < numberDrawn:
		fmt.Printf(greaterNumberDrawn, attempt)
	case attempt > numberDrawn:
		fmt.Printf(lessNumberDrawn, attempt)
	case attempt == numberDrawn:
		fmt.Printf(congratulations, numberDrawn, i+1)

		return true
	}

	return false
}

func drawNumber() int64 {
	return rand.Int64N(int64(quantityDrawn))
}

func printGameOver(numberDrawn int64, attempts [attemptsMax]int64) {
	fmt.Printf(gameOver, numberDrawn, attempts)
}
