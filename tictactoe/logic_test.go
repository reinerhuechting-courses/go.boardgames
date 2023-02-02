package tictactoe

import (
	"fmt"

	"github.com/reinerhuechting-courses/go.boardgames/gameboard"
)

func ExamplePlayerWins_x_row() {
	board1 := gameboard.Board{
		{" ", "O", " "},
		{"X", "X", "X"},
		{"X", "O", " "},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// true
	// false
	// false
	// false
}

func ExamplePlayerWins_x_col() {
	board1 := gameboard.Board{
		{"O", "O", "X"},
		{" ", " ", "X"},
		{"X", "O", "X"},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// true
	// false
	// false
	// false
}

func ExamplePlayerWins_x_diag1() {
	board1 := gameboard.Board{
		{"X", "O", " "},
		{" ", "X", " "},
		{"X", "O", "X"},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// true
	// false
	// false
	// false
}

func ExamplePlayerWins_x_diag2() {
	board1 := gameboard.Board{
		{"O", "O", "X"},
		{" ", "X", " "},
		{"X", "O", "X"},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleDraw() {
	board1 := gameboard.Board{
		{"O", "X", "O"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// false
	// false
	// true
	// false
}

func ExampleGameRunning() {
	board1 := gameboard.Board{
		{"O", "X", "O"},
		{"O", " ", "O"},
		{"X", "O", "X"},
	}
	fmt.Println(PlayerWins(board1, "X"))
	fmt.Println(PlayerWins(board1, "O"))
	fmt.Println(Draw(board1))
	fmt.Println(GameRunning(board1))

	// Output:
	// false
	// false
	// false
	// true
}
