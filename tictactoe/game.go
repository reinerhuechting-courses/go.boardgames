package tictactoe

import (
	"fmt"

	"github.com/reinerhuechting-courses/go.boardgames/gameboard"
)

func Run() {
	// Initialisierung des Spiels
	currentPlayer := "X"
	board := gameboard.NewBoardNumbered(3, 3)

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
	fmt.Println(board)
	fmt.Printf("Spieler %s, du bist an der Reihe.\n", currentPlayer)
	fmt.Print("Bitte wähle ein freies Feld für deinen Zug: ")
	var input int
	fmt.Scanln(&input)
	input--
	row := input / 3
	column := input % 3
	board[row][column] = currentPlayer
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
