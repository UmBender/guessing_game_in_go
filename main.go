package main

import (
	"fmt"
	"math/rand"
)

func main() {
	menuInicial()

	valorChute :=  0
	numeroChutes := 0
	objetivo := rand.Intn(100) + 1

	for {
		if objetivo == valorChute {
			break
		}
		_, err := fmt.Scan(&valorChute)
		if err != nil {
			fmt.Println("Ocorreu um erro na leitura", err)
			continue
		}
		if valorChute > objetivo {
			fmt.Println("Que pena, seu chute foi muito alto!")
		}else if valorChute < objetivo {
			fmt.Println("Que pena, seu chute foi muito baixo!")
		}else {
			fmt.Println("Parabéns! Você acertou o número.")
		}
		numeroChutes++

	}
	fmt.Println("O número de chutes foi de", numeroChutes)
}

func menuInicial()  {
	fmt.Println("Bem vindo ao jogo do chute")
	fmt.Println("Os valores possíveis estão entre 1 e 100")
	
}


