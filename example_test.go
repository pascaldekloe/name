package name_test

import (
	"fmt"
	"github.com/pascaldekloe/name"
)

func ExampleSnakeCase() {
	fmt.Println(name.SnakeCase("CamelToSnake, snake_to_snake: Anything goes!"))
	// Output: camel_to_snake_snake_to_snake_anything_goes
}

func ExampleDelimit() {
	fmt.Println(name.Sep("*All Hype is aGoodThing (TM)", '-'))
	fmt.Println(name.Sep("WebCrawler#socketTimeout", '.'))
	// Output:
	// all-hype-is-a-good-thing-TM
	// web.crawler.socket.timeout
}
