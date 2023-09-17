package main

import "fmt"

func main() {
	fmt.Println("Números divisiveis por 3 entre 1 e 100")

	for i := 1; i <= 100; i++ {
		// Verificar se i é divisível por 3 (resto da divisão é igual a 0)

		if i%3 == 0 {
			fmt.Print(i)
		}
	}

	// Segunda parte do Desafio
	for n := 1; n <= 100; n++ {
		if n%3 == 0 {
			fmt.Println("Pin")
		} else {
			fmt.Println("Pan")
		}

	}
}
