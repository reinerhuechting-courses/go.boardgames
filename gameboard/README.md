# Package `gameboard``

Dieses Package enthält Datentypen und Funktionen für den Umgang mit Spielfeldern.

## Konzept

Es gibt zwei zentrale Datentypen `Row` und `Board`.
Eine Zeile (`Row`) ist eine Liste von Strings und ein Spielfeld (`Board`) ist dann
eine Liste solcher Zeilen.

Insgesamt ist ein `Board` also ein 2D-Array aus Strings, allerdings sind die Datentypen
thematisch passend benannt.
Diese Benennung passiert in den beiden Dateien [`row.go`](row.go) und [`board.go`](board.go)
jeweils ganz oben in Form der Anweisungen `type Row []string` bzw. `type Board []Row`.

Neben einer Umbenennung haben diese Anweisungen aber noch einen weiteren Effekt:
Sie definieren zwei neue Datentypen `Row` und `Board`, denen wir anschließend sog.
*Methoden* zuordnen können. Dies sind Funktionen, die Informationen über die Objekte
bereitstellen können.

Für `Row` und `Board` definieren wir in den beiden Dateien eine ganze Reihe an Methoden,
mit denen relevante Informationen über den Zustand des Spielfelds abgefragt werden können.
Diese Methoden sollen hinterher von den Spielen genutzt werden, um die Spiellogik umzusetzen.
