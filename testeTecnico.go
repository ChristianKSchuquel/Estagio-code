package main

import "fmt"

func main() {
	Questao1()
	fmt.Print("\n")
	Questao2()
	fmt.Print("\n")
	Questao3()
	fmt.Print("\n")
	Questao4()
	fmt.Print("\n")
	Questao5()
}

func Questao1() {
	var Indice int = 13
	var Soma int = 0
	var K int = 0

	for K < Indice {
		K = K + 1

		Soma = Soma + 1
	}

	// Soma: 13
	fmt.Printf("Soma: %v", Soma)
	return
}

func Questao2() {
	var Numero int = 732
	var indice int = 0
	var fib [18]int
	fib[0] = 0
	fib[1] = 1

	for fib[indice] <= Numero {
		if fib[indice] == Numero {
			fmt.Printf("\nO número 732 pertence a sequência")
			return
		}

		fib[indice+2] = fib[indice] + fib[indice+1]
		indice = indice + 1
	}

	fmt.Printf("\nO número 732 não pertence a sequência")
	return
}

func Questao3() {
	a := 9
	b := 128
	c := 49
	d := 100
	e := 13
	f := 200

	fmt.Printf("\nResposta A: %v", a)
	fmt.Printf("\nResposta B: %v", b)
	fmt.Printf("\nResposta C: %v", c)
	fmt.Printf("\nResposta D: %v", d)
	fmt.Printf("\nResposta E: %v", e)
	fmt.Printf("\nResposta F: %v", f)
	return
}

func Questao4() {
	fmt.Println("\nComo o caminhão andava a 80km/h podemos descobrir o tempo que levou para concluir o trajeto sem obstaculos, 100km / 80km/h = 1,25, adicionando os 10 minutos dos pedágios, 0.17 horas, temos 1,42h de viagem")
	fmt.Println("Sabendo o tempo de viagem do caminhão, podemos calcular sua velocidade média 100km / 1,42h = 70,6km/h")
	fmt.Println("Tendo como referência Ribeirão Preto temos: carro -> x1 = 0 + 110*t e caminhão -> x2 = 100km - 70,6*t")
	fmt.Println("Fazemos t = x1 / v1, e t = x2 - 100 / -v2")
	fmt.Println("Agora para resolver fazemos, x = y")
	fmt.Println("x / v1 = x - 100 / -v2")
	fmt.Println("-v2 * x = v1*x - v1*100")
	fmt.Println("x = 100 * 110 / 110 + 70,6")
	fmt.Println("x = 11000 / 180,6 = 60,908")
	return
}

func Questao5() {
	var input string = "Quero essa vaga de estágio!"

	runas := []rune(input)

	for x, y := 0, len(runas)-1; x < y; x, y = x+1, y-1 {
		runas[x], runas[y] = runas[y], runas[x]
	}

	fmt.Print(string(runas))
	return
}
