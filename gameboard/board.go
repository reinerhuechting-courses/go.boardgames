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

// Liefert die i-te Zeile des Spielfelds.
func (board Board) Row(i int) Row {
	return board[i]
}

// Liefert die i-te Spalte des Spielfelds.
func (board Board) Column(i int) Row {
	result := Row{}
	for _, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Liefert die Hauptdiagonale des Spielfelds von links oben nach rechts unten.
func (board Board) PrincipalDiag1() Row {
	result := Row{}
	for i, row := range board {
		result = append(result, row[i])
	}
	return result
}

// Liefert die Hauptdiagonale des Spielfelds von links unten nach rechts oben.
func (board Board) PrincipalDiag2() Row {
	result := Row{}
	for i, row := range board {
		result = append(result, row[len(board)-i-1])
	}
	return result
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

// Prüft, ob eine der Spalten des Spielfelds ausschließlich
// den durch entry gegebenen String enthält.
func (board Board) AnyColumnContainsOnly(entry string) bool {
	for col := 0; col < len(board[0]); col++ {
		if board.Column(col).ContainsOnly(entry) {
			return true
		}
	}
	return false
}

// Prüft, ob eine der Hauptdiagonalen des Spielfelds ausschließlich
// den durch entry gegebenen String enthält.
func (board Board) AnyDiagContainsOnly(entry string) bool {
	return board.PrincipalDiag1().ContainsOnly(entry) ||
		board.PrincipalDiag2().ContainsOnly(entry)
}

func (board Board) EntryCount(entry string) int {
	count := 0
	for _, row := range board {
		count += row.EntryCount(entry)
	}
	return count
}
