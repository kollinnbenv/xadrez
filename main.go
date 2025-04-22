package main

import "fmt"

var tabuleiro_range [8][8]string // Tabuleiro Global

func inicializarTabuleiro() {
	tabuleiro_range = [8][8]string{
		{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"},
		{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{"♟", "♟", "♟", "♟", "♟", "♟", "♟", "♟"},
		{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
	}
}

func moverPeca(origemX, origemY, destinoX, destinoY int) {
	tabuleiro_range[destinoX][destinoY] = tabuleiro_range[origemX][origemY]
	tabuleiro_range[origemX][origemY] = " "
}

func imprimirTabuleiro() {
	fmt.Println("  a b c d e f g h")
	fmt.Println("  ----------------")
	for i, linha := range tabuleiro_range {
		fmt.Printf("%d| ", 8-i)
		for _, casa := range linha {
			fmt.Printf("%s ", casa)
		}
		fmt.Printf("| %d\n", 8-i)
	}
	fmt.Println("  ----------------")
	fmt.Println("  a b c d e f g h")
}

func main() {
	inicializarTabuleiro()
	fmt.Println("Tabuleiro antes do movimento:")
	imprimirTabuleiro()

	moverPeca(1, 4, 3, 4) // Move um peão

	fmt.Println("\nTabuleiro após o movimento:")
	imprimirTabuleiro()
}
