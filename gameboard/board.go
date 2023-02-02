package gameboard

type Board []Row

// Konstruktor für ein Spielfeld mit gegebener Größe und initialer Füllung.
func NewBoard(height, width int, startContent string) Board {
	board := Board{}
	for i := 0; i < height; i++ {
		board = append(board, NewRow(width, startContent))
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
