// mrl.go
// MRL (Minimal Reasoning Language) interpreter using participle
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// -----------------------------------------------------------------------------
// AST Definitions for Participle
// -----------------------------------------------------------------------------

type Program struct {
	Expressions []*Expression `@@*`
}

type Expression struct {
	List       *ListExpr   `  @@`
	Dict       *DictExpr   `| @@`
	Array      *ArrayExpr  `| @@`
	Primitive  *Primitive  `| @@`
}

type ListExpr struct {
	Elements []*Expression `"(" @@* ")"`
}

type DictExpr struct {
	Pairs []*KeyValue `"{" @@* "}"`
}

type KeyValue struct {
	Key   string      `(@Ident | @String | @Number) ":"`
	Value *Expression `@@`
}

type ArrayExpr struct {
	Elements []*Expression `"[" @@* "]"`
}

type Primitive struct {
	Number   *float64 `  @Number`
	String   *string  `| @String`
	Boolean  *bool    `| (@"true" | @"false")`
	Symbol   *string  `| @Symbol`
	Variable *string  `| @Variable`
	Ident    *string  `| @Ident`
}

// -----------------------------------------------------------------------------
// Runtime Value Types
// -----------------------------------------------------------------------------

type Value interface {
	String() string
	IsTruthy() bool
}

type Number float64
func (n Number) String() string { return strconv.FormatFloat(float64(n), 'g', -1, 64) }
func (n Number) IsTruthy() bool { return n != 0 }

type Str string
func (s Str) String() string { return string(s) }
func (s Str) IsTruthy() bool { return len(s) > 0 }

type Bool bool
func (b Bool) String() string { return strconv.FormatBool(bool(b)) }
func (b Bool) IsTruthy() bool { return bool(b) }

type Symbol string
func (s Symbol) String() string { return string(s) }
func (s Symbol) IsTruthy() bool { return true }

type List []Value
func (l List) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i, v := range l {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(v.String())
	}
	sb.WriteString(")")
	return sb.String()
}
func (l List) IsTruthy() bool { return len(l) > 0 }

type Dict map[string]Value
func (d Dict) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	i := 0
	for k, v := range d {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(k + ":" + v.String())
		i++
	}
	sb.WriteString("}")
	return sb.String()
}
func (d Dict) IsTruthy() bool { return len(d) > 0 }

type Array []Value
func (a Array) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range a {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(v.String())
	}
	sb.WriteString("]")
	return sb.String()
}
func (a Array) IsTruthy() bool { return len(a) > 0 }

type MRLFunction struct {
	Name   string
	Params []string
	Body   *Expression
	Env    *Env
}

func (f MRLFunction) String() string {
	return fmt.Sprintf("<fn:%s(%v)>", f.Name, f.Params)
}
func (f MRLFunction) IsTruthy() bool { return true }

// -----------------------------------------------------------------------------
// Environment
// -----------------------------------------------------------------------------

type Env struct {
	vars   map[string]Value
	funcs  map[string]*MRLFunction
	parent *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{
		vars:   make(map[string]Value),
		funcs:  make(map[string]*MRLFunction),
		parent: parent,
	}
}

func (e *Env) GetVar(name string) (Value, bool) {
	v, ok := e.vars[name]
	if ok {
		return v, true
	}
	if e.parent != nil {
		return e.parent.GetVar(name)
	}
	return nil, false
}

func (e *Env) SetVar(name string, val Value) {
	e.vars[name] = val
}

func (e *Env) GetFunc(name string) (*MRLFunction, bool) {
	f, ok := e.funcs[name]
	if ok {
		return f, true
	}
	if e.parent != nil {
		return e.parent.GetFunc(name)
	}
	return nil, false
}

func (e *Env) SetFunc(name string, fn *MRLFunction) {
	e.funcs[name] = fn
}

// -----------------------------------------------------------------------------
// Lexer Definition
// -----------------------------------------------------------------------------

var mrlLexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Comment", Pattern: `;[^\n]*`},
	{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
	{Name: "Number", Pattern: `[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`},
	{Name: "String", Pattern: `"[^"]*"`},
	{Name: "Variable", Pattern: `\?[a-zA-Z_][a-zA-Z0-9_]*`},
	{Name: "Symbol", Pattern: `@[a-zA-Z_][a-zA-Z0-9_]*|[+\-*/<>=]|←|→|λ`},
	{Name: "Ident", Pattern: `[a-zA-Z_][a-zA-Z0-9_]*`},
	{Name: "Punct", Pattern: `[(){}[\]:,]`},
})

// -----------------------------------------------------------------------------
// Parser Setup
// -----------------------------------------------------------------------------

var parser = participle.MustBuild[Program](
	participle.Lexer(mrlLexer),
	participle.Elide("Comment", "Whitespace"),
	participle.UseLookahead(2),
)

// -----------------------------------------------------------------------------
// AST to Value Conversion
// -----------------------------------------------------------------------------

func astToValue(expr *Expression, env *Env) (Value, error) {
	if expr == nil {
		return nil, fmt.Errorf("nil expression")
	}

	if expr.List != nil {
		return evalList(expr.List, env)
	}

	if expr.Dict != nil {
		result := make(Dict)
		for _, pair := range expr.Dict.Pairs {
			val, err := astToValue(pair.Value, env)
			if err != nil {
				return nil, err
			}
			// Clean up key - remove quotes if it's a string
			key := pair.Key
			if len(key) >= 2 && key[0] == '"' && key[len(key)-1] == '"' {
				key = key[1 : len(key)-1]
			}
			result[key] = val
		}
		return result, nil
	}

	if expr.Array != nil {
		var result Array
		for _, elem := range expr.Array.Elements {
			val, err := astToValue(elem, env)
			if err != nil {
				return nil, err
			}
			result = append(result, val)
		}
		return result, nil
	}

	if expr.Primitive != nil {
		return evalPrimitive(expr.Primitive, env)
	}

	return nil, fmt.Errorf("unknown expression type")
}

func evalPrimitive(p *Primitive, env *Env) (Value, error) {
	if p.Number != nil {
		return Number(*p.Number), nil
	}
	if p.String != nil {
		// Remove quotes
		s := *p.String
		if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
			s = s[1 : len(s)-1]
		}
		return Str(s), nil
	}
	if p.Boolean != nil {
		return Bool(*p.Boolean), nil
	}
	if p.Variable != nil {
		val, ok := env.GetVar(*p.Variable)
		if !ok {
			return nil, fmt.Errorf("unbound variable: %s", *p.Variable)
		}
		return val, nil
	}
	if p.Symbol != nil {
		return Symbol(*p.Symbol), nil
	}
	if p.Ident != nil {
		// Identifiers are treated as symbols (for function names, etc.)
		return Symbol(*p.Ident), nil
	}
	return nil, fmt.Errorf("unknown primitive")
}

func evalList(list *ListExpr, env *Env) (Value, error) {
	if len(list.Elements) == 0 {
		return List{}, nil
	}

	// Get first element
	first, err := astToValue(list.Elements[0], env)
	if err != nil {
		return nil, err
	}

	// Check if first is a symbol (primitive or function)
	if sym, ok := first.(Symbol); ok {
		return handleOperation(string(sym), list.Elements[1:], env)
	}

	// Otherwise evaluate all elements as a list
	var result List
	result = append(result, first)
	for _, elem := range list.Elements[1:] {
		val, err := astToValue(elem, env)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

// -----------------------------------------------------------------------------
// Operation Handlers
// -----------------------------------------------------------------------------

func handleOperation(op string, args []*Expression, env *Env) (Value, error) {
	switch op {
	case "@b": // Bind
		return handleBind(args, env)
	case "@g": // Get
		return handleGet(args, env)
	case "@i": // If
		return handleIf(args, env)
	case "@t": // Try
		return handleTry(args, env)
	case "@s": // Synthesize
		return handleSynthesize(args, env)
	case "@o": // Output
		return handleOutput(args, env)
	case "@f": // Fail
		return handleFail(args, env)
	case "@w", "@web_search": // Web search
		return handleWebSearch(args, env)
	case "@c", "@code": // Code execution
		return handleCodeExec(args, env)
	case "@def": // Function definition
		return handleDefine(args, env)
	case "@foreach": // Foreach loop
		return handleForeach(args, env)
	case "@map": // Map
		return handleMap(args, env)
	case "@filter": // Filter
		return handleFilter(args, env)
	case "@collect": // Collect
		return handleCollect(args, env)
	case "@coh": // Coherence update
		return handleCoherence(args, env)
	case "@refl": // Reflection
		return handleReflection(args, env)
	case "+", "-", "*", "/": // Arithmetic
		return handleArithmetic(op, args, env)
	case ">", "<", "=", ">=", "<=": // Comparison
		return handleComparison(op, args, env)
	default:
		// Check if it's a user-defined function
		if fn, ok := env.GetFunc(op); ok {
			return callFunction(fn, args, env)
		}
		return nil, fmt.Errorf("unknown operation: %s", op)
	}
}

func handleBind(args []*Expression, env *Env) (Value, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("@b expects: @b ?var ← expr")
	}

	// Get variable name - should be in Primitive.Variable
	var varName string
	if args[0].Primitive != nil && args[0].Primitive.Variable != nil {
		varName = *args[0].Primitive.Variable
	} else {
		// Try to get as symbol
		varExpr, err := astToValue(args[0], env)
		if err != nil {
			return nil, err
		}
		if sym, ok := varExpr.(Symbol); ok && strings.HasPrefix(string(sym), "?") {
			varName = string(sym)
		} else {
			return nil, fmt.Errorf("first argument to @b must be a variable")
		}
	}

	// Check for arrow (it should be a symbol)
	arrow, err := astToValue(args[1], env)
	if err != nil {
		return nil, err
	}
	if sym, ok := arrow.(Symbol); !ok || (string(sym) != "←" && string(sym) != "<-") {
		return nil, fmt.Errorf("@b expects ← after variable, got: %v", arrow)
	}

	// Evaluate the value
	val, err := astToValue(args[2], env)
	if err != nil {
		return nil, err
	}

	env.SetVar(varName, val)
	return val, nil
}

func handleGet(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@g expects: @g source key")
	}

	source, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	key, err := astToValue(args[1], env)
	if err != nil {
		return nil, err
	}

	// Handle dict access
	if dict, ok := source.(Dict); ok {
		keyStr := key.String()
		if val, exists := dict[keyStr]; exists {
			return val, nil
		}
		return Str(""), nil
	}

	// Handle array/list access
	if list, ok := source.(Array); ok {
		if num, ok := key.(Number); ok {
			idx := int(num)
			if idx >= 0 && idx < len(list) {
				return list[idx], nil
			}
		}
		return nil, fmt.Errorf("index out of bounds")
	}

	return nil, fmt.Errorf("@g: cannot access %T with key %v", source, key)
}

func handleIf(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@i needs condition and then-branch")
	}

	cond, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	if cond.IsTruthy() {
		return astToValue(args[1], env)
	} else if len(args) >= 3 {
		return astToValue(args[2], env)
	}

	return Str(""), nil
}

func handleTry(args []*Expression, env *Env) (Value, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("@t needs at least one expression")
	}

	result, err := astToValue(args[0], env)
	if err != nil {
		// If there's a fallback, try it
		if len(args) >= 2 {
			return astToValue(args[1], env)
		}
		return nil, err
	}
	return result, nil
}

func handleSynthesize(args []*Expression, env *Env) (Value, error) {
	if len(args) == 0 {
		return Str("[SYNTHESIS: empty]"), nil
	}

	var parts []string
	for _, arg := range args {
		val, err := astToValue(arg, env)
		if err != nil {
			return nil, err
		}
		parts = append(parts, val.String())
	}

	return Str("[SYNTHESIS: " + strings.Join(parts, " ") + "]"), nil
}

func handleOutput(args []*Expression, env *Env) (Value, error) {
	if len(args) == 0 {
		return Str(""), nil
	}

	val, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	fmt.Printf("OUTPUT: %s\n", val.String())
	return val, nil
}

func handleFail(args []*Expression, env *Env) (Value, error) {
	msg := "assertion failed"
	if len(args) > 0 {
		val, _ := astToValue(args[0], env)
		msg = val.String()
	}
	return nil, fmt.Errorf("@f: %s", msg)
}

func handleWebSearch(args []*Expression, env *Env) (Value, error) {
	query := "unknown"
	if len(args) > 0 {
		val, err := astToValue(args[0], env)
		if err == nil {
			query = val.String()
		}
	}
	return Str(fmt.Sprintf("[web: searching for '%s']", query)), nil
}

func handleCodeExec(args []*Expression, env *Env) (Value, error) {
	code := ""
	if len(args) > 0 {
		val, err := astToValue(args[0], env)
		if err == nil {
			code = val.String()
		}
	}
	return Str(fmt.Sprintf("[code: would execute '%s']", code)), nil
}

func handleDefine(args []*Expression, env *Env) (Value, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("@def expects: @def name [params] body")
	}

	// Get function name - can be either Symbol or Ident
	nameVal, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}
	name, ok := nameVal.(Symbol)
	if !ok {
		return nil, fmt.Errorf("function name must be a symbol or identifier")
	}

	// Get parameters - need to extract from AST directly without evaluation
	params := []string{}
	if args[1].Array != nil {
		// Parameters are in an array expression
		for _, paramExpr := range args[1].Array.Elements {
			if paramExpr.Primitive != nil && paramExpr.Primitive.Variable != nil {
				params = append(params, *paramExpr.Primitive.Variable)
			} else if paramExpr.Primitive != nil && paramExpr.Primitive.Symbol != nil {
				// Also accept symbols that start with ?
				sym := *paramExpr.Primitive.Symbol
				if strings.HasPrefix(sym, "?") {
					params = append(params, sym)
				}
			}
		}
	} else {
		return nil, fmt.Errorf("@def expects parameters as an array: [?param1 ?param2 ...]")
	}

	// Store function
	fn := &MRLFunction{
		Name:   string(name),
		Params: params,
		Body:   args[2],
		Env:    env,
	}
	env.SetFunc(string(name), fn)

	return fn, nil
}

func callFunction(fn *MRLFunction, args []*Expression, env *Env) (Value, error) {
	if len(args) != len(fn.Params) {
		return nil, fmt.Errorf("function %s expects %d arguments, got %d",
			fn.Name, len(fn.Params), len(args))
	}

	// Create new environment
	fnEnv := NewEnv(fn.Env)

	// Evaluate and bind arguments
	for i, param := range fn.Params {
		val, err := astToValue(args[i], env)
		if err != nil {
			return nil, err
		}
		fnEnv.SetVar(param, val)
	}

	// Execute function body
	return astToValue(fn.Body, fnEnv)
}

func handleForeach(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@foreach expects collection and body")
	}

	collection, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	var items []Value
	if arr, ok := collection.(Array); ok {
		items = []Value(arr)
	} else if list, ok := collection.(List); ok {
		items = []Value(list)
	} else {
		return nil, fmt.Errorf("@foreach requires array or list")
	}

	var results Array
	for _, item := range items {
		loopEnv := NewEnv(env)
		loopEnv.SetVar("?it", item)

		for _, bodyExpr := range args[1:] {
			result, err := astToValue(bodyExpr, loopEnv)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
	}

	return results, nil
}

func handleMap(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@map expects collection and function")
	}

	collection, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	var items []Value
	if arr, ok := collection.(Array); ok {
		items = []Value(arr)
	} else if list, ok := collection.(List); ok {
		items = []Value(list)
	}

	var results Array
	for _, item := range items {
		mapEnv := NewEnv(env)
		mapEnv.SetVar("?it", item)
		result, err := astToValue(args[1], mapEnv)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func handleFilter(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@filter expects collection and predicate")
	}

	collection, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	var items []Value
	if arr, ok := collection.(Array); ok {
		items = []Value(arr)
	} else if list, ok := collection.(List); ok {
		items = []Value(list)
	}

	var results Array
	for _, item := range items {
		filterEnv := NewEnv(env)
		filterEnv.SetVar("?it", item)
		pred, err := astToValue(args[1], filterEnv)
		if err != nil {
			continue
		}
		if pred.IsTruthy() {
			results = append(results, item)
		}
	}

	return results, nil
}

func handleCollect(args []*Expression, env *Env) (Value, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("@collect expects collection and item")
	}

	collection, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}

	item, err := astToValue(args[1], env)
	if err != nil {
		return nil, err
	}

	var results Array
	if arr, ok := collection.(Array); ok {
		results = arr
	}
	results = append(results, item)

	// Update variable
	if len(args) > 0 {
		if expr := args[0]; expr.Primitive != nil && expr.Primitive.Variable != nil {
			env.SetVar(*expr.Primitive.Variable, results)
		}
	}

	return results, nil
}

func handleCoherence(args []*Expression, env *Env) (Value, error) {
	return Str("[coherence: updated]"), nil
}

func handleReflection(args []*Expression, env *Env) (Value, error) {
	return Str("[reflection: generated]"), nil
}

func handleArithmetic(op string, args []*Expression, env *Env) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("%s expects 2 arguments", op)
	}

	a, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}
	b, err := astToValue(args[1], env)
	if err != nil {
		return nil, err
	}

	// Special case: string concatenation with +
	if op == "+" {
		// If either operand is a string, do string concatenation
		if _, isStrA := a.(Str); isStrA {
			return Str(a.String() + b.String()), nil
		}
		if _, isStrB := b.(Str); isStrB {
			return Str(a.String() + b.String()), nil
		}
	}

	na := toNumber(a)
	nb := toNumber(b)

	switch op {
	case "+":
		return Number(na + nb), nil
	case "-":
		return Number(na - nb), nil
	case "*":
		return Number(na * nb), nil
	case "/":
		if nb == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return Number(na / nb), nil
	}

	return nil, fmt.Errorf("unknown arithmetic op: %s", op)
}

func handleComparison(op string, args []*Expression, env *Env) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("%s expects 2 arguments", op)
	}

	a, err := astToValue(args[0], env)
	if err != nil {
		return nil, err
	}
	b, err := astToValue(args[1], env)
	if err != nil {
		return nil, err
	}

	na := toNumber(a)
	nb := toNumber(b)

	switch op {
	case ">":
		return Bool(na > nb), nil
	case "<":
		return Bool(na < nb), nil
	case "=":
		return Bool(na == nb), nil
	case ">=":
		return Bool(na >= nb), nil
	case "<=":
		return Bool(na <= nb), nil
	}

	return nil, fmt.Errorf("unknown comparison op: %s", op)
}

func toNumber(v Value) float64 {
	switch x := v.(type) {
	case Number:
		return float64(x)
	case Str:
		f, _ := strconv.ParseFloat(string(x), 64)
		return f
	case Bool:
		if x {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// -----------------------------------------------------------------------------
// REPL
// -----------------------------------------------------------------------------

func printHelp() {
	fmt.Println(`
MRL Interpreter v1.0 (with Participle)
=======================================

Core Primitives:
  @b ?var ← expr              Bind variable
  @g source key               Get/extract data
  @i cond then [else]         If-then-else
  @t expr [alt: fallback]     Try with fallback
  @s [text...]                Synthesize/summarize
  @o expr                     Output result
  @f [message]                Fail/assert
  @w query                    Web search (stub)
  @c code                     Code execution (stub)

Functions:
  @def name [params] body     Define function
  (name arg1 arg2...)         Call function

Collections:
  @foreach collection body    Iterate over items
  @map collection fn          Map function over items
  @filter collection pred     Filter items by predicate
  @collect ?list item         Append item to list

Operators:
  + - * /                     Arithmetic
  > < = >= <=                 Comparisons

Meta-Cognitive:
  @coh state updates          Update coherence
  @refl state                 Reflection

Examples:
  (@b ?x ← 5)
  (@b ?y ← 10)
  (@o (+ ?x ?y))              ; OUTPUT: 15

  (@def add [?a ?b] (+ ?a ?b))
  (add 3 4)                   ; → 7

  (@b ?nums ← [1 2 3 4 5])
  (@filter ?nums (> ?it 2))   ; → [3 4 5]

Commands:
  :q, :quit    Exit
  :h, :help    Show help
  :funcs       List defined functions
`)
}

func listFunctions(env *Env) {
	if len(env.funcs) == 0 {
		fmt.Println("No functions defined.")
		return
	}
	fmt.Println("Defined functions:")
	for name, fn := range env.funcs {
		fmt.Printf("  %s(%v)\n", name, fn.Params)
	}
}

func main() {
	fmt.Println("MRL Interpreter v1.0 (with Participle)")
	fmt.Println("Type :h for help\n")

	global := NewEnv(nil)

	// Pre-define some utility functions
	global.SetFunc("double", &MRLFunction{
		Name:   "double",
		Params: []string{"?x"},
		Body: &Expression{
			List: &ListExpr{
				Elements: []*Expression{
					{Primitive: &Primitive{Symbol: stringPtr("*")}},
					{Primitive: &Primitive{Variable: stringPtr("?x")}},
					{Primitive: &Primitive{Number: float64Ptr(2)}},
				},
			},
		},
		Env: global,
	})

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("mrl> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Handle meta commands
		switch line {
		case ":q", ":quit":
			fmt.Println("Goodbye!")
			return
		case ":h", ":help":
			printHelp()
			continue
		case ":funcs":
			listFunctions(global)
			continue
		}

		// Parse and evaluate
		program, err := parser.ParseString("", line)
		if err != nil {
			fmt.Printf("PARSE ERROR: %v\n", err)
			continue
		}

		var lastResult Value
		for _, expr := range program.Expressions {
			result, err := astToValue(expr, global)
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				break // Stop on first error
			} else if result != nil {
				lastResult = result
			}
		}
		
		// Only print the last result
		if lastResult != nil {
			fmt.Printf("→ %s\n", lastResult.String())
		}
	}
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}
