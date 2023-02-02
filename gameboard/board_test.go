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
