package gameboard

import "fmt"

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
