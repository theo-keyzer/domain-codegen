package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	rioLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Comment", `//.*`},
		{"Whitespace", `\s+`},
		// Ident allows dots, colons, and underscores
		{"Ident", `[.a-zA-Z_][a-zA-Z0-9_:.]*`},
		// JSON-eligible quotes
		{"JsonQuote", `(?s)'''.*?'''|'.*?'`},
		// Raw/Literal quotes
		{"RawQuote", `(?s)"""Selection"""|"""(?s).*?"""|".*?"|(?s)~~~.*?~~~|~.*?~`},
		{"Punct", `[][={},]`}, // Added comma to Punct
	})
)

type RioFile struct {
	Components []ComponentDef `@@*`
}

type ComponentDef struct {
	Pos    lexer.Position
	Type   string     `@Ident`
	Fields []FieldDef `"{" @@* "}"`
}

type FieldDef struct {
	Key   string   `@Ident "="`
	Value ValueDef `@@`
}

type ValueDef struct {
	// Captures one or more items separated by commas
	Items []SingleValue `@@ ( "," @@ )*`
}

type SingleValue struct {
	Json *string `( @JsonQuote | @Ident )`
	Raw  *string `| @RawQuote`
}

type Component struct {
	Type       string         `json:"_type"`
	Name       string         `json:"key"`
	Parent     string         `json:"parent,omitempty"`
	Filename   string         `json:"filename"`
	Fields     map[string]any `json:"fields"`
	LineNumber int            `json:"line_number"`
}

func (f *RioFile) ToLegacyComponents(filename string) []Component {
	var result []Component
	for _, c := range f.Components {
		comp := Component{
			Type:       c.Type,
			Filename:   filename,
			Fields:     make(map[string]any),
			LineNumber: c.Pos.Line,
		}

		for _, field := range c.Fields {
			var processedValues []any
			for _, item := range field.Value.Items {
				var val any
				if item.Raw != nil {
					val = cleanValue(*item.Raw, false, true)
				} else if item.Json != nil {
					val = cleanValue(*item.Json, true, false)
				}
				processedValues = append(processedValues, val)
			}

			// If only one item was provided, don't return it as a list
			var finalVal any
			if len(processedValues) == 1 {
				finalVal = processedValues[0]
			} else {
				finalVal = processedValues
			}

			switch field.Key {
			case "key":
				comp.Name = fmt.Sprint(finalVal)
			case "parent":
				comp.Parent = fmt.Sprint(finalVal)
			default:
				comp.Fields[field.Key] = finalVal
			}
		}
		result = append(result, comp)
	}
	return result
}

func cleanValue(s string, detectJson bool, isRaw bool) any {
	isTripleTilde := strings.HasPrefix(s, "~~~")
	unquoted := unquote(s)

	if isTripleTilde && strings.HasPrefix(unquoted, "\n") {
		unquoted = unquoted[1:]
	}

	if detectJson {
		var js any
		if err := json.Unmarshal([]byte(unquoted), &js); err == nil {
			return js
		}
	}
	return unquoted
}

func unquote(s string) string {
	if len(s) < 2 {
		return s
	}
	if len(s) >= 6 {
		for _, d := range []string{`"""`, `'''`, `~~~`} {
			if s[:3] == d && s[len(s)-3:] == d {
				return s[3 : len(s)-3]
			}
		}
	}
	for _, d := range []string{`"`, `'`, `~`} {
		if s[:1] == d && s[len(s)-1:] == d {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func test_main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <file.rio>")
	}
	path := os.Args[1]

	parser, err := participle.Build[RioFile](
		participle.Lexer(rioLexer),
		participle.Elide("Comment", "Whitespace"),
	)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rioFile, err := parser.Parse(path, file)
	if err != nil {
		log.Fatal(err)
	}

	components := rioFile.ToLegacyComponents(path)
	output, _ := json.MarshalIndent(components, "", "  ")
	fmt.Println(string(output))
}
