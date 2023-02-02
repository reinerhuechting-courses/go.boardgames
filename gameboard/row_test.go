package gameboard

import "fmt"

func ExampleRow_constructor() {
	board1 := NewRow(3, "X")

	fmt.Println(board1.ContainsOnly(" "))
	fmt.Println(board1.ContainsOnly("X"))
	fmt.Println(board1.ContainsOnly("O"))
	fmt.Println(len(board1))
	fmt.Println(board1[0])
	fmt.Println(board1[1])
	fmt.Println(board1[2])

	// Output:
	// false
	// true
	// false
	// 3
	// X
	// X
	// X
}

func ExampleRow_ContainsOnly_xrow() {
	row1 := Row{"X", "X", "X"}

	fmt.Println(row1.ContainsOnly("X"))
	fmt.Println(row1.ContainsOnly("O"))
	fmt.Println(row1.ContainsOnly("P"))
	fmt.Println(row1.ContainsOnly(""))

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleRow_ContainsOnly_mixedrow() {
	row1 := Row{"X", "X", "O"}

	fmt.Println(row1.ContainsOnly("X"))
	fmt.Println(row1.ContainsOnly("O"))
	fmt.Println(row1.ContainsOnly("P"))
	fmt.Println(row1.ContainsOnly(""))

	// Output:
	// false
	// false
	// false
	// false
}

func ExampleRow_ContainsOnly_emptyrow() {
	row1 := Row{}

	fmt.Println(row1.ContainsOnly("X"))
	fmt.Println(row1.ContainsOnly("O"))
	fmt.Println(row1.ContainsOnly("P"))
	fmt.Println(row1.ContainsOnly(""))

	// Output:
	// true
	// true
	// true
	// true
}

func ExampleRow_String() {
	row1 := Row{"X", "X", "O"}
	fmt.Println(row1)
	row2 := Row{" ", " ", " "}
	fmt.Println(row2)

	// Output:
	// | X | X | O |
	// |   |   |   |
}
