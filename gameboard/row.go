package gameboard

type Row []string

// Prüft, ob die Zeile ausschließlich den durch entry gegebenen String enthält.
func (row Row) ContainsOnly(entry string) bool {
	for _, e := range row {
		if e != entry {
			return false
		}
	}
	return true
}
