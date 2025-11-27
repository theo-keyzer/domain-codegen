package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"strings"
)

type Component struct {
	Type       string         `json:"_type"`
	Name       string         `json:"key"`
	Fields     map[string]any `json:"fields"`
	LineNumber int            `json:"line_number"`
}

var (
	// Simple key = value
	simpleRegex = regexp.MustCompile(`^(\w+)\s*=\s*(.+)$`)
	// key = [a, b, c]
	arrayRegex = regexp.MustCompile(`^(\w+)\s*=\s*\[(.*)\]$`)
	// key = """ or key = ~~~
	multiStartRegex = regexp.MustCompile(`^(\w+)\s*=\s*(("""|~~~))\s*$`)
)

func ParseRioFile(path string) ([]Component, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var components []Component
	var current *Component
	inMulti := false
	multiKey := ""
	delimiter := ""
	var multiLines []string

	scanner := bufio.NewScanner(file)
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		raw := scanner.Text()
		trimmed := strings.TrimSpace(raw)

		// Skip blank lines and full-line comments outside blocks
		if trimmed == "" || strings.HasPrefix(trimmed, "//") && !inMulti && current == nil {
			continue
		}

		// === MULTI-LINE HANDLING ===
		if inMulti {
			if strings.TrimSpace(trimmed) == delimiter {
				content := strings.Join(multiLines, "\n")
				if strings.HasSuffix(content, "\n") {
					content = content[:len(content)-1]
				}

				if delimiter == `"""` {
var js any
    if err := json.Unmarshal([]byte(content), &js); err != nil {
        // Optional: log error or panic in dev
        // log.Printf("Invalid JSON in %s at line %d: %v", multiKey, lineNo, err)
        current.Fields[multiKey] = map[string]any{"_error": "invalid json, use ~~~", "_raw": content}
    } else {
        current.Fields[multiKey] = js
    }				} else {
					current.Fields[multiKey] = content
				}

				inMulti = false
				multiLines = nil
				continue
			}
			multiLines = append(multiLines, raw)
			continue
		}

		// === MULTI-LINE START ===
		if m := multiStartRegex.FindStringSubmatch(trimmed); m != nil && current != nil {
			multiKey = m[1]
			delimiter = m[2]
			inMulti = true
			multiLines = nil
			continue
		}

		// === COMPONENT START: Type "name" {  or  Type name { ===
		if strings.Contains(raw, "{") {
			openIdx := strings.Index(raw, "{")
			before := strings.TrimSpace(raw[:openIdx])
			after := strings.TrimSpace(raw[openIdx+1:])

			parts := strings.Fields(before)
			if len(parts) < 1 {
				continue
			}

			compType := parts[0]
			compName := ""
			if len(parts) > 1 {
				// Everything after type is the name (supports spaces + quotes)
				nameRaw := strings.Join(parts[1:], " ")
				compName = strings.Trim(nameRaw, `"`)
			}

			current = &Component{
				Type:       compType,
				Name:       compName,
				Fields:     map[string]any{}, //{"key": compName},
				LineNumber: lineNo,
			}

			if after != "" && after != "}" {
				parseField(current, after)
			}
			continue
		}

		// === COMPONENT END ===
		if trimmed == "}" && current != nil {
			components = append(components, *current)
			current = nil
			continue
		}

		// === INSIDE COMPONENT: normal field ===
		if current != nil {
			noComment := stripLineComments(raw)
			if strings.TrimSpace(noComment) != "" {
				parseField(current, noComment)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return components, nil
}

// Safely strips // comments only when not inside quotes
func stripLineComments(line string) string {
	inQuote := false
	escaped := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		if escaped {
			escaped = false
			continue
		}
		if c == '\\' {
			escaped = true
			continue
		}
		if c == '"' {
			inQuote = !inQuote
			continue
		}
		if !inQuote && i+1 < len(line) && c == '/' && line[i+1] == '/' {
			return strings.TrimSpace(line[:i])
		}
	}
	return strings.TrimSpace(line)
}

func parseField(c *Component, line string) {
	line = strings.TrimSpace(line)

	// Array: tags = [a, b, "c d"]
	if m := arrayRegex.FindStringSubmatch(line); m != nil {
		key := m[1]
		content := strings.TrimSpace(m[2])

		// Try as JSON array first
		var arr []any
		if json.Unmarshal([]byte("["+content+"]"), &arr) == nil {
			c.Fields[key] = arr
			return
		}
		// Fallback: split on commas (simple but safe)
		items := splitCommaSeparated(content)
		var simple []string
		for _, s := range items {
			s = strings.TrimSpace(strings.Trim(s, `"`))
			if s != "" {
				simple = append(simple, s)
			}
		}
		c.Fields[key] = simple
		return
	}

	// Simple key = value
	if m := simpleRegex.FindStringSubmatch(line); m != nil {
		key := m[1]
		val := strings.TrimSpace(m[2])
		val = strings.Trim(val, `"`)

		var js any
		if json.Unmarshal([]byte(val), &js) == nil {
			c.Fields[key] = js
		} else {
			c.Fields[key] = val
		}
	}
}

// Simple comma splitter that respects quotes (good enough for 99% cases)
func splitCommaSeparated(s string) []string {
	var parts []string
	var current strings.Builder
	inQuote := false

	for _, r := range s {
		if r == '"' {
			inQuote = !inQuote
			current.WriteRune(r)
			continue
		}
		if r == ',' && !inQuote {
			parts = append(parts, current.String())
			current.Reset()
			continue
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	return parts
}
