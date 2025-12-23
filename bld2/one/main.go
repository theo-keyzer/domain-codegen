package main

import (
	"encoding/json"
	"fmt"
	"os"
	"log"
	"strings"
	"github.com/alecthomas/participle/v2"
)

// GlobT represents global state
type GlobT struct {
	LoadErrs  int
	RunErrs   int
	Acts      ActT
	Dats      ActT
	Winp      int
	Wins      []WinT
	Collect   map[string]interface{}
	OutOn     bool
	InOn      bool
	Ins       strings.Builder
}


// KpExtra represents extra key-value pairs
type KpExtra struct {
	Kp
	Names map[string]string
}
func (me KpExtra) GetVar(glob *GlobT, s []string, ln string) (bool, interface{}) {
	r,ok := me.Names[s[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s, Command line arguments", s[0], ln) }
	return ok,r
}
func (me KpExtra) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
func (me KpExtra) GetLineNo() string {
	return "Command line arguments"
}

func main() {
	args := os.Args[1:] // Skip the program name
	if len(args) < 2 {
		fmt.Println(args)
		return
	}

	glob := new(GlobT)
	glob.Winp = -1
	glob.OutOn = true
	glob.Collect = make(map[string]interface{})
	// Load files and check for errors
	glob.LoadErrs += loadFilesh(args[0], &glob.Acts)
	glob.LoadErrs += loadFilesh(args[1], &glob.Dats)

	if len(glob.Acts.ApActor) > 0 {
		kp := &KpExtra{
			Names: make(map[string]string),
		}
		
		// Store args in kp.Names
		for i, arg := range args {
			kp.Names[fmt.Sprint(i)] = arg
		}
		NewAct(glob, glob.Acts.ApActor[0].Kname, "", "run:1", "", "", "", true)
		GoAct(glob, kp)
	}

	if glob.LoadErrs > 0 || glob.RunErrs > 0 {
//		fmt.Println("Errors", glob.LoadErrs, glob.RunErrs)
		println("Errors", glob.LoadErrs, glob.RunErrs)
		os.Exit(1)
	}
}

func loadFilesh(files string, act *ActT) int {
	act.index = make(map[string]int) 
	errs := 0
	fileList := strings.Split(files, ",")
	
	for _, file := range fileList {
		components := loadRio(file)
		/*
		components, err := ParseRioFile(file)
		if err != nil {
			fmt.Printf("Error reading h file %s: %v\n", file, err)
			errs = errs + 1
			continue
		}
		*/
		for _, comp := range components {
			lno := fmt.Sprintf("%s:%d", file, comp.LineNumber)
			Loadh(act, comp.Type, "", 7, lno,comp.Fields)
		}
		
	}
	
	errs += refs(act)
	return errs
}

func loadRio(path string) []Component {
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

	// Pass filename to the legacy converter
	components := rioFile.ToLegacyComponents(path)
	return components

}

func cnv(ss interface{}) string {
	if ss == nil {
		return ""
	}
    	if ret, ok := ss.(string); ok {
    		return ret
    	}
	bytes, err := json.Marshal(ss)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func cnv2(ss interface{}) string {
	if ss == nil {
		return ""
	}
    	if _, ok := ss.([]interface{}); ok {
    		ss = ""
    	}
	return ss.(string)
}

// fnd3 finds a value in the index
func fnd3(act *ActT, s string, f string, msg string, chk string, lno string, olno string) (bool, int) {
	if v, exists := act.index[s]; exists {
		return true, v
	}
	
	if chk == "*" || chk == "?" {
		return true, -1
	}
	if f == chk {
		return true, -1
	}
	if f == "" && chk == "." {
		return true, -1
	}
	
	fmt.Printf("%s %s (%s) not found %s (%s) > %s\n", msg,f, s, lno,chk,olno)
	return false, -1
}

// fnd2 finds a value in the index
func fnd2(act *ActT, s string, f string, chk string, lno string, olno string) (bool, int) {
	if v, exists := act.index[s]; exists {
		return true, v
	}
	
	if chk == "*" || chk == "?" {
		return true, -1
	}
	if f == chk {
		return true, -1
	}
	
	fmt.Printf("%s (%s) not found %s (%s) > %s\n", f, s, lno,chk,olno)
	return false, -1
}

// getName gets a value from a map by name
func getName(m map[string]string, n string) string {
	if v, ok := m[n]; ok {
		return v
	}
	return ""
}


