package main

import (
	"fmt"
)

type Node struct {
	esquerda, direita *Node
	valor             string
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

func InsereValorNaArvore(node *Node, valor string) *Node {
	if node == nil {
		node = &Node{nil, nil, valor}
	} else if node.valor == "" {
		node.valor = valor
	} else {
		if node.valor < valor {
			node.esquerda = InsereValorNaArvore(node.esquerda, valor)

		} else if node.valor > valor {
			node.direita = InsereValorNaArvore(node.direita, valor)
		} else {
			fmt.Print("Valor já existe\n")
		}
	}
	return node
}

func ProcuraValorNaArvore(node *Node, valor string) *Node {
	if node.valor == valor {
		return node
	} else {
		if node.valor < valor {
			resp := ProcuraValorNaArvore(node.esquerda, valor)
			if resp.valor == valor {
				return resp
			}
		} else {
			resp := ProcuraValorNaArvore(node.direita, valor)
			if resp.valor == valor {
				return resp
			}
		}
	}
	return nil
}

func main() {
	raiz := Node{nil, nil, ""}
	InsereValorNaArvore(&raiz, "c")
	InsereValorNaArvore(&raiz, "d")
	InsereValorNaArvore(&raiz, "e")
	InsereValorNaArvore(&raiz, "f")
	InsereValorNaArvore(&raiz, "m")
	InsereValorNaArvore(&raiz, "n")
	InsereValorNaArvore(&raiz, "o")
	InsereValorNaArvore(&raiz, "g")
	InsereValorNaArvore(&raiz, "h")
	InsereValorNaArvore(&raiz, "w")
	InsereValorNaArvore(&raiz, "x")
	InsereValorNaArvore(&raiz, "i")
	InsereValorNaArvore(&raiz, "j")
	InsereValorNaArvore(&raiz, "k")
	InsereValorNaArvore(&raiz, "l")
	InsereValorNaArvore(&raiz, "p")
	InsereValorNaArvore(&raiz, "q")
	InsereValorNaArvore(&raiz, "r")
	InsereValorNaArvore(&raiz, "s")
	InsereValorNaArvore(&raiz, "t")
	InsereValorNaArvore(&raiz, "u")
	InsereValorNaArvore(&raiz, "v")
	InsereValorNaArvore(&raiz, "a")
	InsereValorNaArvore(&raiz, "b")
	InsereValorNaArvore(&raiz, "y")
	InsereValorNaArvore(&raiz, "z")

	/*item111 := Node{nil, nil, 111}
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
	*/
	fmt.Println(Largura(raiz))
	fmt.Print(raiz)
	ImprimeArvore(raiz)
	fmt.Print(ProcuraValorNaArvore(&raiz, "q"))

}
