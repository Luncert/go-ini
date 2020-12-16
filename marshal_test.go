package go_ini

import (
	"fmt"
)

func ExampleMarshal() {
	cfg := NewConfig()
	cfg.AddSection(
		NewSection("a").
			AddVariable(NewVariable("name", "Joy")).
			AddVariable(NewVariable("age", 18))).
		AddSection(
			NewSection("b").
				AddVariable(NewVariable("single", false)).
				AddVariable(NewVariable("money", 813.01))).
		AddSection(
			NewSection("c").
				AddVariable(NewVariable("contact", []string{"12345", "123456"})))

	fmt.Println(Marshal(cfg))

	// Output:
	// [a]
	// age = 18
	// name = "Joy"
	//
	// [b]
	// money = 813.010000
	// single = no
	//
	// [c]
	// contact = "12345", "123456"
	//
}
