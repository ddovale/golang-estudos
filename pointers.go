package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature = "shark"
	var pointer *string = &creature

	fmt.Println("Creature =", creature)
	fmt.Println("Pointer =", pointer)

	//Com desreferenciação
	fmt.Println("*Pointer =", *pointer)

	//Alterando o valor armazenado no endereço a qual
	//o ponteiro aponta
	*pointer = "jellyfish"
	fmt.Println("*Pointer =", *pointer)

	//Isso significa que também alteramos o valor da variável
	//creature, pois o ponteiro aponta para o endereço da
	//variável creature
	fmt.Println("Creature =", creature)

	fmt.Println("-------------------")

	//Passando variáveis por valor e por referência para funções
	var creature2 Creature = Creature{Species: "shark"}

	fmt.Printf("1)Original: %+v\n", creature2)
	changeCreatureValue(creature2)
	fmt.Printf("3)Original: %+v\n", creature2)
	changeCreatureReference(&creature2)
	fmt.Printf("5)Original: %+v\n", creature2)
}

//Recebe por valor, logo não altera a variável original
func changeCreatureValue(creature Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2)Value: %+v\n", creature)
}

//Recebe por referência, logo altera a variável original
func changeCreatureReference(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("4)Reference: %+v\n", creature)
}
