package name_test

import (
	"fmt"
	"github.com/pascaldekloe/name"
)

func ExampleSnakeCase() {
	fmt.Print(name.SnakeCase("CamelToSnake"))
	fmt.Print(", ", name.SnakeCase("snake_to_snake"))
	fmt.Print(": ", name.SnakeCase("Anything goes!"))
	// Output: camel_to_snake, snake_to_snake: anything_goes
}

func ExampleSep() {
	fmt.Print(name.Sep("*All Hype is aGoodThing (TM)", '-'))
	// Output: all-hype-is-a-good-thing-TM
}
