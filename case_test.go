package name

import "testing"

var goldenCases = []struct{ snake, lowerCamel, upperCamel string }{
	{"name", "name", "Name"},

	{"", "", ""},
	{"I", "i", "I"},
	{"ID", "iD", "ID"},
	{"wi_fi", "wiFi", "WiFi"},

	// single outer abbreviation
	{"vitamin_C", "vitaminC", "VitaminC"},
	{"T_cell", "tCell", "TCell"},

	// double outer abbreviation
	{"master_DB", "masterDB", "MasterDB"},
	{"IO_bounds", "iOBounds", "IOBounds"},

	// tripple outer abbreviation
	{"main_API", "mainAPI", "MainAPI"},
	{"TCP_conn", "tCPConn", "TCPConn"},

	// inner abbreviation
	{"raw_URL_query", "rawURLQuery", "RawURLQuery"},

	// numbers
	{"4x4", "4x4", "4x4"},
	{"no5", "no5", "No5"},
	{"DB2", "dB2", "DB2"},
	{"3M", "3M", "3M"},
	{"7_up", "7Up", "7Up"},
	{"20th", "20th", "20th"},
}

func TestSnakeToSnake(t *testing.T) {
	for _, golden := range goldenCases {
		s := golden.snake
		if got := SnakeCase(s); got != s {
			t.Errorf("%q: got %q", s, got)
		}
	}
}

func TestLowerCamelToLowerCamel(t *testing.T) {
	for _, golden := range goldenCases {
		s := golden.lowerCamel
		if got := CamelCase(s, false); got != s {
			t.Errorf("%q: got %q", s, got)
		}
	}
}

func TestUpperCamelToUpperCamel(t *testing.T) {
	for _, golden := range goldenCases {
		s := golden.upperCamel
		if got := CamelCase(s, true); got != s {
			t.Errorf("%q: got %q", s, got)
		}
	}
}

func TestSnakeToLowerCamel(t *testing.T) {
	for _, golden := range goldenCases {
		snake, want := golden.snake, golden.lowerCamel
		if got := CamelCase(snake, false); got != want {
			t.Errorf("%q: got %q, want %q", snake, got, want)
		}
	}
}

func TestSnakeToUpperCamel(t *testing.T) {
	for _, golden := range goldenCases {
		snake, want := golden.snake, golden.upperCamel
		if got := CamelCase(snake, true); got != want {
			t.Errorf("%q: got %q, want %q", snake, got, want)
		}
	}
}

func TestLowerCamelToSnake(t *testing.T) {
	for _, golden := range goldenCases {
		camel, want := golden.lowerCamel, golden.snake
		if got := SnakeCase(camel); got != want {
			t.Errorf("%q: got %q, want %q", camel, got, want)
		}
	}
}

func TestLowerCamelToUpperCamel(t *testing.T) {
	for _, golden := range goldenCases {
		camel, want := golden.lowerCamel, golden.upperCamel
		if got := CamelCase(camel, true); got != want {
			t.Errorf("%q: got %q, want %q", camel, got, want)
		}
	}
}

func TestUpperCamelToSnake(t *testing.T) {
	for _, golden := range goldenCases {
		camel, want := golden.upperCamel, golden.snake
		if got := SnakeCase(camel); got != want {
			t.Errorf("%q: got %q, want %q", camel, got, want)
		}
	}
}

func TestUpperCamelToLowerCamel(t *testing.T) {
	for _, golden := range goldenCases {
		camel, want := golden.upperCamel, golden.lowerCamel
		if got := CamelCase(camel, false); got != want {
			t.Errorf("%q: got %q, want %q", camel, got, want)
		}
	}
}

func BenchmarkCases(b *testing.B) {
	for _, sample := range []string{"a2B", "foo-bar", "ProcessHelperFactoryConfig#defaultIDBuilder"} {
		b.Run(sample, func(b *testing.B) {
			b.Run("CamelCase", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					CamelCase(sample, true)
				}
			})
			b.Run("snake_case", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					SnakeCase(sample)
				}
			})
		})
	}
}
