package tictactoe

import "github.com/reinerhuechting-courses/go.boardgames/gameboard"

// Liefert true, falls der angegebene Spieler das Spiel gewonnen hat.
func PlayerWins(board gameboard.Board, player string) bool {
	return board.AnyRowContainsOnly(player) ||
		board.AnyColumnContainsOnly(player) ||
		board.AnyDiagContainsOnly(player)
}
