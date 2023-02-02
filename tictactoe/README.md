# Package `tictactoe``

Dieses Package enthält Datentypen und Funktionen für den Umgang mit Spielfeldern.

## Konzept

Das Package ist in mehrere Dateien unterteilt, die unterschiedlichen Zwecken dienen.
Bei größeren Projekten wären diese Dateien möglicherweise sogar weitere Unterpackages,
hier soll diese Unterteilung zunächst ausreichen.

* [`logic.go`](logic.go)  
  Enthält Funktionen für die Spiellogik.
  I.W. sind dies Funktionen, die jeweils ein Spielfeld und einen Spieler erwarten.
  Das Spielfeld ist dabei vom Typ [`gameboard.Board`](../gameboard/board.go)
  und der Spieler ist einfach ein String, der das Zeichen des Spielers ("X" oder "O")
  enthält.
  Der Zweck dieser Funktionen ist es, den Spielzustand zu bestimmen, also
ob das Spiel noch läuft, ein Spieler gewonnen hat oder es unentschieden beendet ist.
