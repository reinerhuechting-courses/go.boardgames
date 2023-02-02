package gameboard

import "fmt"

func ExampleBoard_constructor() {
	board1 := NewBoard(3, 4, " ")

	fmt.Println(board1.AnyRowContainsOnly(" "))
	fmt.Println(board1.AnyRowContainsOnly("X"))
	fmt.Println(board1.AnyRowContainsOnly("O"))
	fmt.Println(len(board1))
	fmt.Println(len(board1[0]))

	// Output:
	// true
	// false
	// false
	// 3
	// 4
}

func ExampleBoard_numbered() {
	board1 := NewBoardNumbered(2, 3)
	for _, row := range board1 {
		fmt.Println(row)
	}
	fmt.Println()
	board2 := NewBoardNumbered(3, 3)
	for _, row := range board2 {
		fmt.Println(row)
	}

	// Output:
	// [1 2 3]
	// [4 5 6]
	//
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]

}

func ExampleBoard_Row() {
	board1 := Board{
		{"X", "X", "O"},
		{"O", "O", "O"},
		{"O", "X", "X"},
	}

	fmt.Println(board1.Row(0))
	fmt.Println(board1.Row(1))
	fmt.Println(board1.Row(2))

	// Output:
	// [X X O]
	// [O O O]
	// [O X X]
}

func ExampleBoard_Column() {
	board1 := Board{
		{"X", "X", "O"},
		{"O", "O", "O"},
		{"O", "X", "X"},
	}

	fmt.Println(board1.Column(0))
	fmt.Println(board1.Column(1))
	fmt.Println(board1.Column(2))

	// Output:
	// [X O O]
	// [X O X]
	// [O O X]
}

func ExampleBoard_PrincipalDiag1() {
	board1 := NewBoardNumbered(3, 3)

	fmt.Println(board1.PrincipalDiag1())

	// Output:
	// [1 5 9]
}

func ExampleBoard_PrincipalDiag2() {
	board1 := NewBoardNumbered(3, 3)

	fmt.Println(board1.PrincipalDiag2())

	// Output:
	// [3 5 7]
}

func ExampleBoard_AnyRowContainsOnly() {
	board1 := Board{
		{"X", "X", "O"},
		{"O", "O", "O"},
		{"O", "X", "X"},
	}

	fmt.Println(board1.AnyRowContainsOnly("X"))
	fmt.Println(board1.AnyRowContainsOnly("O"))
	fmt.Println(board1.AnyRowContainsOnly("P"))

	// Output:
	// false
	// true
	// false
}

func ExampleBoard_AnyColumnContainsOnly() {
	board1 := Board{
		{"O", "X", "P"},
		{"O", "O", "P"},
		{"O", "X", "P"},
	}

	fmt.Println(board1.AnyColumnContainsOnly("X"))
	fmt.Println(board1.AnyColumnContainsOnly("O"))
	fmt.Println(board1.AnyColumnContainsOnly("P"))

	// Output:
	// false
	// true
	// true
}

func ExampleBoard_AnyDiagContainsOnly() {
	board1 := Board{
		{"O", "X", "P"},
		{"O", "O", "P"},
		{"O", "X", "O"},
	}

	fmt.Println(board1.AnyDiagContainsOnly("X"))
	fmt.Println(board1.AnyDiagContainsOnly("O"))

	// Output:
	// false
	// true
}

func ExampleBoard_EntryCount() {
	board1 := Board{
		{"O", "X", "P"},
		{"O", "O", "P"},
		{"O", "X", "O"},
	}

	fmt.Println(board1.EntryCount("O"))
	fmt.Println(board1.EntryCount("X"))
	fmt.Println(board1.EntryCount("P"))

	// Output:
	// 5
	// 2
	// 2
}

func ExampleBoard_String() {
	board1 := NewBoardNumbered(3, 3)
	fmt.Println(board1)

	// Output:
	// +---+---+---+
	// | 1 | 2 | 3 |
	// +---+---+---+
	// | 4 | 5 | 6 |
	// +---+---+---+
	// | 7 | 8 | 9 |
	// +---+---+---+
}
