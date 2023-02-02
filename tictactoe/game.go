package tictactoe

import (
	"fmt"

	"github.com/reinerhuechting-courses/go.boardgames/gameboard"
)

func Run() {
	// Initialisierung des Spiels
	currentPlayer := "X"
	board := gameboard.NewBoard(3, 3, " ")

	// Game Loop
	for GameRunning(board) {
		board = MakeMove(board, currentPlayer)
		currentPlayer = NextPlayer(currentPlayer)
	}

	// Auswertung
	PrintResult(board)
}

// Fragt einen Zug vom angegebenen Spieler ab und führt diesen aus.
func MakeMove(board gameboard.Board, currentPlayer string) gameboard.Board {
	// TODO
	return board
}

// Gibt das Spielergebnis auf die Konsole aus.
func PrintResult(board gameboard.Board) {
	fmt.Println("Das Spiel ist beendet.")
	if PlayerWins(board, "X") {
		fmt.Println("Spieler X gewinnt.")
	}
	if PlayerWins(board, "O") {
		fmt.Println("Spieler O gewinnt.")
	}
	if Draw(board) {
		fmt.Println("Es gibt keinen Gewinner.")
	}
}
