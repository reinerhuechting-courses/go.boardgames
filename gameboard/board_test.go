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
