package gameboard

import "fmt"

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
