package exercicios

import (
	"bufio"
	"fmt"
	"os"
)

const instruction2 string = "Digite um número"

func EvenOrOdd() {
	input := bufio.NewReader(os.Stdin)
	var number int
	var pair string = "par"
	var odd string = "ímpar"

	fmt.Println(instruction2)

	if _, err := fmt.Fscan(input, &number); err != nil {
		fmt.Println(notNumber)

		return
	}

	if number%2 == 0 {
		fmt.Println(pair)

		return
	}

	fmt.Println(odd)
}
