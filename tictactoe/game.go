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
	// Hilfsvariable für die Prüfung, ob ein gültiger Zug eingegeben wurde.
	moveValid := false

	// In einer Schleife nach Eingaben fragen, bis eine gültige Eingabe gemacht wurde:
	for !moveValid {

		// Spielfeld ausgeben:
		fmt.Println(board)

		// Spieler nach seinem Zug fragen:
		fmt.Printf("Spieler %s, du bist an der Reihe.\n", currentPlayer)
		fmt.Print("Bitte wähle ein freies Feld für deinen Zug: ")
		var input int
		fmt.Scanln(&input)

		// Ziel-Position des Zuges berechnen:
		input--
		row := input / 3
		column := input % 3

		// Gültigkeit der Eingabe prüfen:
		inputValid := input >= 0 && input <= 8
		// Prüfen, ob die Eingabe erlaubt ist:
		moveValid = inputValid && board[row][column] != "X" && board[row][column] != "O"

		// Zug ausführen oder Fehlermeldung ausgeben:
		if moveValid {
			board[row][column] = currentPlayer
		} else {
			fmt.Println("Diese Eingabe ist nicht erlaubt!")
			fmt.Println()
		}
	}
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
