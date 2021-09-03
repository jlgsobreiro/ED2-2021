package main

import (
	"fmt"
)

type Node struct {
	esquerda, direita *Node
	valor             int
}

func Largura(node Node) (largD, largE int) {
	if node.esquerda != nil {
		largD, largE = Largura(*node.esquerda)
		largE += 1
	}
	if node.direita != nil {
		largD, largE = Largura(*node.direita)
		largD += 1
	}
	return
}

func ImprimeArvore(node Node) {
	fmt.Print(" ", node.valor, " ")
	if node.esquerda != nil {
		for _, i := Largura(node); i > 0; i-- {
			fmt.Print("  ")
		}
		ImprimeArvore(*node.esquerda)
	}
	if node.direita != nil {
		for i, _ := Largura(node); i > 0; i-- {
			fmt.Print("  ")
		}
		ImprimeArvore(*node.direita)

	}
	fmt.Print("\n")
}

func main() {
	raiz := Node{nil, nil, 0}
	item111 := Node{nil, nil, 111}
	item221 := Node{nil, nil, 221}
	item211 := Node{&item221, nil, 211}
	item11 := Node{&item111, nil, 11}
	item12 := Node{nil, nil, 12}
	item21 := Node{&item211, nil, 21}
	item22 := Node{nil, nil, 22}
	item1 := Node{&item11, &item12, 1}
	item2 := Node{&item21, &item22, 2}
	raiz.direita = &item1
	raiz.esquerda = &item2
	//fmt.Println(Largura(raiz))
	ImprimeArvore(raiz)

}
