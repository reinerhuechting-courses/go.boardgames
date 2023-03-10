package tictactoe

import "github.com/reinerhuechting-courses/go.boardgames/gameboard"

// Liefert true, falls der angegebene Spieler das Spiel gewonnen hat.
func PlayerWins(board gameboard.Board, player string) bool {
	return board.AnyRowContainsOnly(player) ||
		board.AnyColumnContainsOnly(player) ||
		board.AnyDiagContainsOnly(player)
}

// Liefert true, falls das Spiel mit dem Ergebnis "unentschieden" beendet ist.
func Draw(board gameboard.Board) bool {
	return board.EntryCount("X")+board.EntryCount("O") == 9 &&
		!PlayerWins(board, "X") &&
		!PlayerWins(board, "O")
}

// Liefert true, falls das Spiel beendet ist (egal, mit welchem Ergebnis).
func GameRunning(board gameboard.Board) bool {
	return !PlayerWins(board, "X") &&
		!PlayerWins(board, "O") &&
		!Draw(board)
}

// Liefert den Spieler, der als nächstes am Zug ist.
func NextPlayer(currentPlayer string) string {
	if currentPlayer == "X" {
		return "O"
	}
	return "X"
}
