package go_ini

import (
	"fmt"
)

func ExampleNewConfig() {
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

	fmt.Println(cfg.ToString())
	// Output:
	// Config {
	//  Sections = [
	//    Section(a) {
	//      Variables = {
	//        age  = 18
	//        name = "Joy"
	//      }
	//    }
	//    Section(b) {
	//      Variables = {
	//        money  = 813.010000
	//        single = no
	//      }
	//    }
	//    Section(c) {
	//      Variables = {
	//        contact = ["12345", "123456"]
	//      }
	//    }
	//  ]
	// }
}
