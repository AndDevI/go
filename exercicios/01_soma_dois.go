package exercicios

import (
	"bufio"
	"fmt"
	"os"
)

const instruction string = "Digite 2 números para somar"
const notNumber string = "Você não digitou um número"

func Somar() {
	input := bufio.NewReader(os.Stdin)
	var number1 int
	var number2 int

	fmt.Println(instruction)

	if _, err := fmt.Fscan(input, &number1, &number2); err != nil {
		fmt.Println(notNumber)
		
		return
	}

	fmt.Println(number1 + number2)
}
