package gameboard

import "fmt"

type Row []string

// Konstruktor für eine Zeile mit der angegebenen Breite und initialer Füllung.
func NewRow(width int, startContent string) Row {
	row := Row{}
	for i := 0; i < width; i++ {
		row = append(row, startContent)
	}
	return row
}

// Prüft, ob die Zeile ausschließlich den durch entry gegebenen String enthält.
func (row Row) ContainsOnly(entry string) bool {
	for _, e := range row {
		if e != entry {
			return false
		}
	}
	return true
}

// Liefert die Anzahl der Felder in der Zeile, die mit entry belegt sind.
func (row Row) EntryCount(entry string) int {
	count := 0
	for _, e := range row {
		if e == entry {
			count++
		}
	}
	return count
}

// Liefert eine String-Repräsentation der Zeile.
// Diese Repräsentation ist für das Zeichnen des Spielfelds geeignet.
func (row Row) String() string {
	result := "|"
	for _, e := range row {
		result += fmt.Sprintf(" %s |", e)
	}
	return result
}
