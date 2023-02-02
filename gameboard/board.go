package gameboard

type Board []Row

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
