package name

import "testing"

var goldenCamelSnakes = map[string]string{
	"name": "name",
	"ID":   "ID",

	"loFi": "lo_fi",
	"HiFi": "hi_fi",

	"rawHTTPBody": "raw_HTTP_body",

	// single outer
	"vitaminC": "vitamin_C",
	"TCell":    "T_cell",

	// double outer
	"masterDB": "master_DB",
	"IOBounds": "IO_bounds",

	// tripple outer
	"mainAPI": "main_API",
	"TCPConn": "TCP_conn",

	// numbers
	"b2b":  "b2b",
	"4x4":  "4x4",
	"No5":  "no5",
	"DB2":  "DB2",
	"3M":   "3M",
	"7Up":  "7_up",
	"20th": "20th",
}

func TestCamelToSnake(t *testing.T) {
	for camel, snake := range goldenCamelSnakes {
		if got := SnakeCase(camel); got != snake {
			t.Errorf("ToSnake(%q) = %q, want %q", camel, got, snake)
		}
	}
}

func TestSnakeToSnake(t *testing.T) {
	for _, s := range goldenCamelSnakes {
		if got := SnakeCase(s); got != s {
			t.Errorf("ToSnake(%q) = %q", s, got)
		}
	}
}
