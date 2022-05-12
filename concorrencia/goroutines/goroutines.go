package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrência é diferente de paralelismo

	// paralelismo acontece quando existe 2 ou mais tarefas executando exatamente ao mesmo tempo,
	// possivel somente com processadores com mais de 1 núcleo

	// tarefas concorrete não necessáriamente estão executando ao mesmo tempo,
	// mas uma tarefa não espera outra tarefa terminar para poder processar.
	// As tarefas ou processos, vão "revezando".
	// Um tarefa executa um pouco e pararia, para outra tarefa processar um pouco, parar, e voltar para a primeira.

	fmt.Println("gotoutines")
	go escrever("Olá Mundo!") // gotoutine
	escrever("Programando em Golang!")

}

func escrever(texto string) {
	// loop infinito
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
