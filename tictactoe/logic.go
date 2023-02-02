package tictactoe

import "github.com/reinerhuechting-courses/go.boardgames/gameboard"

// Liefert true, falls der angegebene Spieler das Spiel gewonnen hat.
func PlayerWins(board gameboard.Board, player string) bool {
	return board.AnyRowContainsOnly(player) ||
		board.AnyColumnContainsOnly(player) ||
		board.AnyDiagContainsOnly(player)
}

// Liefert true, falls das Spielmit dem Ergebnis "unentschieden" beendet ist.
func Draw(board gameboard.Board) bool {
	return board.EntryCount(" ") == 0 &&
		!PlayerWins(board, "X") &&
		!PlayerWins(board, "O")
}
