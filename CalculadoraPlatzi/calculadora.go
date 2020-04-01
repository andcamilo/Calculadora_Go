package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type calc struct {}

func (calc)operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println("el resultado es: ",operador1+operador2)
		return operador1+operador2
	case "-":
		fmt.Println("el resultado es: ",operador1-operador2)
		return operador1-operador2
	case "*":
		fmt.Println("el resultado es: ",operador1*operador2)
		return operador1*operador2
	case "/":
		fmt.Println("el resultado es: ",operador1/operador2)
		return operador1/operador2
	default:
		fmt.Println("Operación no soportada")
		return 0



	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador

}

func LeerEntrada() string {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()	
}

func main() {
	
	entrada := LeerEntrada()
	operador := LeerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador)) 


	
}