package gameboard

import "fmt"

type Board []Row

// Konstruktor für ein Spielfeld mit gegebener Größe und initialer Füllung.
func NewBoard(height, width int, startContent string) Board {
	board := Board{}
	for i := 0; i < height; i++ {
		board = append(board, NewRow(width, startContent))
	}
	return board
}

// Konstruktor für ein Spielfeld mit gegebener Größe, dessen Felder nummeriert sind.
func NewBoardNumbered(height, width int) Board {
	board := NewBoard(height, width, "")
	entry := 1
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			board[row][col] = fmt.Sprintf("%d", entry)
			entry++
		}
	}
	return board
}

// Prüft, ob eine der Zeilen des Spielfelds ausschließlich
// den durch entry gegebenen String enthält.
func (board Board) AnyRowContainsOnly(entry string) bool {
	for _, row := range board {
		if row.ContainsOnly(entry) {
			return true
		}
	}
	return false
}
