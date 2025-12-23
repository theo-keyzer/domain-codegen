package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"maps"
)

type Kp interface {
	DoIts(glob *GlobT, va []string, lno string) int
	GetVar(glob *GlobT, va []string, lno string) (bool, interface{})
	GetLineNo() string
	TypeName() string
}

type KpCanon struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	ItsLink [] *KpLink 
	ItsSection [] *KpSection 
	Childs [] Kp
}

func (me KpCanon) TypeName() string {
    return me.Comp
}
func (me KpCanon) GetLineNo() string {
	return me.LineNo
}

func loadCanon(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCanon)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApCanon)
	st.LineNo = lno
	st.Comp = "Canon";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["name"].(string)
	st.Names["_key"] = "name"
	act.index["Canon_" + name] = st.Me;
	st.MyName = name
	act.ApCanon = append(act.ApCanon, st)
	return 0
}

func (me KpCanon) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if (va[0] == "Link_concept" && len(va) > 1) { // one.unit:25, go-struct-rio.act:765
		for _, st := range glob.Dats.ApLink {
			if (st.Kconceptp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Link_relation" && len(va) > 1) { // one.unit:26, go-struct-rio.act:765
		for _, st := range glob.Dats.ApLink {
			if (st.Krelationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Objective_canonical_form" && len(va) > 1) { // one.unit:44, go-struct-rio.act:765
		for _, st := range glob.Dats.ApObjective {
			if (st.Kcanonical_formp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Memory_canonical_form" && len(va) > 1) { // one.unit:68, go-struct-rio.act:765
		for _, st := range glob.Dats.ApMemory {
			if (st.Kcanonical_formp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Thought_canonical_form" && len(va) > 1) { // one.unit:95, go-struct-rio.act:765
		for _, st := range glob.Dats.ApThought {
			if (st.Kcanonical_formp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Memo_canonical_form" && len(va) > 1) { // one.unit:112, go-struct-rio.act:765
		for _, st := range glob.Dats.ApMemo {
			if (st.Kcanonical_formp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // one.unit:8, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Canon > one.unit:8, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Canon > one.unit:8, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpCanon) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Link" { // one.unit:20, go-struct-rio.act:744
		for _, st := range me.ItsLink {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Section" { // one.unit:28, go-struct-rio.act:744
		for _, st := range me.ItsSection {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if (va[0] == "Link_concept") { // one.unit:25, go-struct-rio.act:654
		for _, st := range glob.Dats.ApLink {
			if (st.Kconceptp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Link_relation") { // one.unit:26, go-struct-rio.act:654
		for _, st := range glob.Dats.ApLink {
			if (st.Krelationp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Objective_canonical_form") { // one.unit:44, go-struct-rio.act:654
		for _, st := range glob.Dats.ApObjective {
			if (st.Kcanonical_formp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Memory_canonical_form") { // one.unit:68, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemory {
			if (st.Kcanonical_formp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Thought_canonical_form") { // one.unit:95, go-struct-rio.act:654
		for _, st := range glob.Dats.ApThought {
			if (st.Kcanonical_formp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Memo_canonical_form") { // one.unit:112, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemo {
			if (st.Kcanonical_formp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Canon %s,%s > one.unit:8, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpLink struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kconceptp int
	Krelationp int
}

func (me KpLink) TypeName() string {
    return me.Comp
}
func (me KpLink) GetLineNo() string {
	return me.LineNo
}

func loadLink(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpLink)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApLink)
	st.LineNo = lno
	st.Comp = "Link";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kconceptp = -1
	st.Krelationp = -1
	st.Kparentp = len( act.ApCanon ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Link has no Canon parent\n") ;
		return 1
	}
	st.Parent = act.ApCanon[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Link under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApCanon[ len( act.ApCanon )-1 ].Childs = append(act.ApCanon[ len( act.ApCanon )-1 ].Childs, st)
	act.ApCanon[ len( act.ApCanon )-1 ].ItsLink = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsLink, st)	// one.unit:8, go-struct-rio.act:467
	act.ApLink = append(act.ApLink, st)
	return 0
}

func (me KpLink) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "concept" { // one.unit:25, go-struct-rio.act:680
		if (me.Kconceptp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kconceptp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "relation" { // one.unit:26, go-struct-rio.act:680
		if (me.Krelationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Krelationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // one.unit:8, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:20, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApLink[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Link > one.unit:20, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Link > one.unit:20, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpLink) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // one.unit:8, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCanon[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "concept" {
		if me.Kconceptp >= 0 {
			st := glob.Dats.ApCanon[ me.Kconceptp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "relation" {
		if me.Krelationp >= 0 {
			st := glob.Dats.ApCanon[ me.Krelationp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Link %s,%s > one.unit:20, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpSection struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
}

func (me KpSection) TypeName() string {
    return me.Comp
}
func (me KpSection) GetLineNo() string {
	return me.LineNo
}

func loadSection(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpSection)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApSection)
	st.LineNo = lno
	st.Comp = "Section";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApCanon ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Section has no Canon parent\n") ;
		return 1
	}
	st.Parent = act.ApCanon[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Section under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApCanon[ len( act.ApCanon )-1 ].Childs = append(act.ApCanon[ len( act.ApCanon )-1 ].Childs, st)
	act.ApCanon[ len( act.ApCanon )-1 ].ItsSection = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsSection, st)	// one.unit:8, go-struct-rio.act:467
	name,_ := st.Names["name"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Section_" + name	// one.unit:32, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApSection = append(act.ApSection, st)
	return 0
}

func (me KpSection) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if (va[0] == "parent") { // one.unit:8, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:28, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSection[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Section > one.unit:28, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Section > one.unit:28, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpSection) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // one.unit:8, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCanon[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
		if va[0] == "level" && len(va) > 1 && me.Kparentp >= 0 { // one.unit:33, go-struct-rio.act:284
			pos := 0
			v, ok := me.Names["level"].(string)
			if ok {
				pos, _ = strconv.Atoi(v)
			}
			if va[1] == "down" && pos > 0 {
				pst := glob.Dats.ApCanon[me.Kparentp]
				isin := false
				for _, st := range pst.ItsSection {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["level"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						break
					}
					if pos2 == pos {
						if len(va) > 2 {
							return st.DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, st)
					}
				}
				return 0
			}
			if va[1] == "up" && pos > 0 {
				pst := glob.Dats.ApCanon[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsSection {
					pos2 := 0
					v2, ok2 := st.Names["level"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == pos && st.Me != me.Me {
						prev = st.Me
						isin = true
						continue
					}
					if pos2 == (pos - 1) {
						isin = false
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApSection[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApSection[prev])
					}
				}
				return 0
			}
			if va[1] == "left" && pos > 0 {
				pst := glob.Dats.ApCanon[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsSection {
					pos2 := 0
					v2, ok2 := st.Names["level"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						prev = st.Me
						isin = true
						continue
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApSection[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApSection[prev])
					}
				}
				return 0
			}
			if va[1] == "right" && pos > 0 {
				pst := glob.Dats.ApCanon[me.Kparentp]
				isin := false
				for _, st := range pst.ItsSection {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["level"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 <= pos {
						break
					}
					if pos2 == (pos + 1) {
						if len(va) > 2 {
							ret := st.DoIts(glob, va[2:], lno)
							if ret != 0 {
								return ret
							}
							continue
						}
						ret := GoAct(glob, st)
						if ret != 0 {
							return ret
						}
					}
				}
				return 0
			}
		}
	        fmt.Printf("?No its %s for Section %s,%s > one.unit:28, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpObjective struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kcanonical_formp int
	ItsMemory [] *KpMemory 
	ItsThought [] *KpThought 
	Childs [] Kp
}

func (me KpObjective) TypeName() string {
    return me.Comp
}
func (me KpObjective) GetLineNo() string {
	return me.LineNo
}

func loadObjective(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpObjective)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApObjective)
	st.LineNo = lno
	st.Comp = "Objective";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonical_formp = -1
	name,_ := st.Names["objective_id"].(string)
	st.Names["_key"] = "objective_id"
	act.index["Objective_" + name] = st.Me;
	st.MyName = name
	act.ApObjective = append(act.ApObjective, st)
	return 0
}

func (me KpObjective) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "canonical_form" { // one.unit:44, go-struct-rio.act:680
		if (me.Kcanonical_formp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonical_formp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Memo_objective_id" && len(va) > 1) { // one.unit:113, go-struct-rio.act:765
		for _, st := range glob.Dats.ApMemo {
			if (st.Kobjective_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // one.unit:41, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Objective > one.unit:41, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Objective > one.unit:41, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpObjective) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Memory" { // one.unit:61, go-struct-rio.act:744
		for _, st := range me.ItsMemory {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Thought" { // one.unit:78, go-struct-rio.act:744
		for _, st := range me.ItsThought {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "canonical_form" {
		if me.Kcanonical_formp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonical_formp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Memo_objective_id") { // one.unit:113, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemo {
			if (st.Kobjective_idp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Objective %s,%s > one.unit:41, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpMemory struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kinvalidated_byp int
	Kcanonical_formp int
}

func (me KpMemory) TypeName() string {
    return me.Comp
}
func (me KpMemory) GetLineNo() string {
	return me.LineNo
}

func loadMemory(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpMemory)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApMemory)
	st.LineNo = lno
	st.Comp = "Memory";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kinvalidated_byp = -1
	st.Kcanonical_formp = -1
	st.Kparentp = len( act.ApObjective ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Memory has no Objective parent\n") ;
		return 1
	}
	st.Parent = act.ApObjective[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Memory under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApObjective[ len( act.ApObjective )-1 ].Childs = append(act.ApObjective[ len( act.ApObjective )-1 ].Childs, st)
	act.ApObjective[ len( act.ApObjective )-1 ].ItsMemory = append(act.ApObjective[ len( act.ApObjective )-1 ].ItsMemory, st)	// one.unit:41, go-struct-rio.act:467
	name,_ := st.Names["memory_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Memory_" + name	// one.unit:62, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApMemory = append(act.ApMemory, st)
	return 0
}

func (me KpMemory) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "invalidated_by" { // one.unit:67, go-struct-rio.act:680
		if (me.Kinvalidated_byp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kinvalidated_byp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "canonical_form" { // one.unit:68, go-struct-rio.act:680
		if (me.Kcanonical_formp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonical_formp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // one.unit:41, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:61, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Memory > one.unit:61, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Memory > one.unit:61, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpMemory) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // one.unit:41, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApObjective[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "invalidated_by" {
		if me.Kinvalidated_byp >= 0 {
			st := glob.Dats.ApThought[ me.Kinvalidated_byp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canonical_form" {
		if me.Kcanonical_formp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonical_formp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Memo_memory_id") { // one.unit:114, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemo {
			if (st.Kmemory_idp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Memory %s,%s > one.unit:61, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpThought struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kcanonical_formp int
	ItsMemo [] *KpMemo 
	Childs [] Kp
}

func (me KpThought) TypeName() string {
    return me.Comp
}
func (me KpThought) GetLineNo() string {
	return me.LineNo
}

func loadThought(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpThought)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApThought)
	st.LineNo = lno
	st.Comp = "Thought";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonical_formp = -1
	st.Kparentp = len( act.ApObjective ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Thought has no Objective parent\n") ;
		return 1
	}
	st.Parent = act.ApObjective[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Thought under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApObjective[ len( act.ApObjective )-1 ].Childs = append(act.ApObjective[ len( act.ApObjective )-1 ].Childs, st)
	act.ApObjective[ len( act.ApObjective )-1 ].ItsThought = append(act.ApObjective[ len( act.ApObjective )-1 ].ItsThought, st)	// one.unit:41, go-struct-rio.act:467
	name,_ := st.Names["thought_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Thought_" + name	// one.unit:79, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApThought = append(act.ApThought, st)
	return 0
}

func (me KpThought) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "canonical_form" { // one.unit:95, go-struct-rio.act:680
		if (me.Kcanonical_formp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonical_formp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // one.unit:41, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Memory_invalidated_by" && len(va) > 1) { // one.unit:67, go-struct-rio.act:765
		for _, st := range glob.Dats.ApMemory {
			if (st.Kinvalidated_byp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // one.unit:78, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Thought > one.unit:78, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Thought > one.unit:78, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpThought) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Memo" { // one.unit:108, go-struct-rio.act:744
		for _, st := range me.ItsMemo {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "parent" { // one.unit:41, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApObjective[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canonical_form" {
		if me.Kcanonical_formp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonical_formp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Memory_invalidated_by") { // one.unit:67, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemory {
			if (st.Kinvalidated_byp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
		if va[0] == "tree_linear" && len(va) > 1 && me.Kparentp >= 0 { // one.unit:80, go-struct-rio.act:284
			pos := 0
			v, ok := me.Names["tree_linear"].(string)
			if ok {
				pos, _ = strconv.Atoi(v)
			}
			if va[1] == "down" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				for _, st := range pst.ItsThought {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["tree_linear"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						break
					}
					if pos2 == pos {
						if len(va) > 2 {
							return st.DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, st)
					}
				}
				return 0
			}
			if va[1] == "up" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsThought {
					pos2 := 0
					v2, ok2 := st.Names["tree_linear"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == pos && st.Me != me.Me {
						prev = st.Me
						isin = true
						continue
					}
					if pos2 == (pos - 1) {
						isin = false
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApThought[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApThought[prev])
					}
				}
				return 0
			}
			if va[1] == "left" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsThought {
					pos2 := 0
					v2, ok2 := st.Names["tree_linear"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						prev = st.Me
						isin = true
						continue
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApThought[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApThought[prev])
					}
				}
				return 0
			}
			if va[1] == "right" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				for _, st := range pst.ItsThought {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["tree_linear"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 <= pos {
						break
					}
					if pos2 == (pos + 1) {
						if len(va) > 2 {
							ret := st.DoIts(glob, va[2:], lno)
							if ret != 0 {
								return ret
							}
							continue
						}
						ret := GoAct(glob, st)
						if ret != 0 {
							return ret
						}
					}
				}
				return 0
			}
		}
		if va[0] == "tree_parallel" && len(va) > 1 && me.Kparentp >= 0 { // one.unit:81, go-struct-rio.act:284
			pos := 0
			v, ok := me.Names["tree_parallel"].(string)
			if ok {
				pos, _ = strconv.Atoi(v)
			}
			if va[1] == "down" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				for _, st := range pst.ItsThought {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["tree_parallel"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						break
					}
					if pos2 == pos {
						if len(va) > 2 {
							return st.DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, st)
					}
				}
				return 0
			}
			if va[1] == "up" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsThought {
					pos2 := 0
					v2, ok2 := st.Names["tree_parallel"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == pos && st.Me != me.Me {
						prev = st.Me
						isin = true
						continue
					}
					if pos2 == (pos - 1) {
						isin = false
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApThought[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApThought[prev])
					}
				}
				return 0
			}
			if va[1] == "left" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsThought {
					pos2 := 0
					v2, ok2 := st.Names["tree_parallel"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 == (pos - 1) {
						prev = st.Me
						isin = true
						continue
					}
					if st.Me == me.Me && isin {
						if len(va) > 2 {
							return glob.Dats.ApThought[prev].DoIts(glob, va[2:], lno)
						}
						return GoAct(glob, glob.Dats.ApThought[prev])
					}
				}
				return 0
			}
			if va[1] == "right" && pos > 0 {
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				for _, st := range pst.ItsThought {
					if st.Me == me.Me {
						isin = true
						continue
					}
					if !isin {
						continue
					}
					pos2 := 0
					v2, ok2 := st.Names["tree_parallel"].(string)
					if ok2 {
						pos2, _ = strconv.Atoi(v2)
					}
					if pos2 == 0 {
						continue
					}
					if pos2 <= pos {
						break
					}
					if pos2 == (pos + 1) {
						if len(va) > 2 {
							ret := st.DoIts(glob, va[2:], lno)
							if ret != 0 {
								return ret
							}
							continue
						}
						ret := GoAct(glob, st)
						if ret != 0 {
							return ret
						}
					}
				}
				return 0
			}
		}
	        fmt.Printf("?No its %s for Thought %s,%s > one.unit:78, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpMemo struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kcanonical_formp int
	Kobjective_idp int
	Kmemory_idp int
	Kcomputationp int
}

func (me KpMemo) TypeName() string {
    return me.Comp
}
func (me KpMemo) GetLineNo() string {
	return me.LineNo
}

func loadMemo(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpMemo)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApMemo)
	st.LineNo = lno
	st.Comp = "Memo";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonical_formp = -1
	st.Kobjective_idp = -1
	st.Kmemory_idp = -1
	st.Kcomputationp = -1
	st.Kparentp = len( act.ApThought ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Memo has no Thought parent\n") ;
		return 1
	}
	st.Parent = act.ApThought[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Memo under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApThought[ len( act.ApThought )-1 ].Childs = append(act.ApThought[ len( act.ApThought )-1 ].Childs, st)
	act.ApThought[ len( act.ApThought )-1 ].ItsMemo = append(act.ApThought[ len( act.ApThought )-1 ].ItsMemo, st)	// one.unit:78, go-struct-rio.act:467
	name,_ := st.Names["memo_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Memo_" + name	// one.unit:109, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApMemo = append(act.ApMemo, st)
	return 0
}

func (me KpMemo) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "canonical_form" { // one.unit:112, go-struct-rio.act:680
		if (me.Kcanonical_formp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonical_formp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "objective_id" { // one.unit:113, go-struct-rio.act:680
		if (me.Kobjective_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kobjective_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "memory_id" { // one.unit:114, go-struct-rio.act:680
		if (me.Kmemory_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kmemory_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "computation" { // one.unit:115, go-struct-rio.act:680
		if (me.Kcomputationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOps[ me.Kcomputationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // one.unit:78, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:108, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMemo[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Memo > one.unit:108, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Memo > one.unit:108, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpMemo) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // one.unit:78, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApThought[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canonical_form" {
		if me.Kcanonical_formp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonical_formp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "objective_id" {
		if me.Kobjective_idp >= 0 {
			st := glob.Dats.ApObjective[ me.Kobjective_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "memory_id" {
		if me.Kmemory_idp >= 0 {
			st := glob.Dats.ApMemory[ me.Kmemory_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "computation" {
		if me.Kcomputationp >= 0 {
			st := glob.Dats.ApOps[ me.Kcomputationp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Memo %s,%s > one.unit:108, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpOps struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
}

func (me KpOps) TypeName() string {
    return me.Comp
}
func (me KpOps) GetLineNo() string {
	return me.LineNo
}

func loadOps(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpOps)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApOps)
	st.LineNo = lno
	st.Comp = "Ops";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["ops_id"].(string)
	st.Names["_key"] = "ops_id"
	act.index["Ops_" + name] = st.Me;
	st.MyName = name
	act.ApOps = append(act.ApOps, st)
	return 0
}

func (me KpOps) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if (va[0] == "Memo_computation" && len(va) > 1) { // one.unit:115, go-struct-rio.act:765
		for _, st := range glob.Dats.ApMemo {
			if (st.Kcomputationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // one.unit:125, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApOps[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Ops > one.unit:125, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Ops > one.unit:125, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpOps) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "data" {
		val := me.Names["data"]
		return( GoAct(glob, val) )
	}
	if (va[0] == "Memo_computation") { // one.unit:115, go-struct-rio.act:654
		for _, st := range glob.Dats.ApMemo {
			if (st.Kcomputationp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Ops %s,%s > one.unit:125, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpComp struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	ItsElement [] *KpElement 
	Childs [] Kp
}

func (me KpComp) TypeName() string {
    return me.Comp
}
func (me KpComp) GetLineNo() string {
	return me.LineNo
}

func loadComp(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpComp)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApComp)
	st.LineNo = lno
	st.Comp = "Comp";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = -1
	name,_ := st.Names["name"].(string)
	st.Names["_key"] = "name"
	act.index["Comp_" + name] = st.Me;
	st.MyName = name
	act.ApComp = append(act.ApComp, st)
	return 0
}

func (me KpComp) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "parent" { // one.unit:140, go-struct-rio.act:680
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kparentp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Comp_parent" && len(va) > 1) { // one.unit:140, go-struct-rio.act:765
		for _, st := range glob.Dats.ApComp {
			if (st.Kparentp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Element_comp" && len(va) > 1) { // one.unit:168, go-struct-rio.act:765
		for _, st := range glob.Dats.ApElement {
			if (st.Kcompp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // one.unit:134, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Comp > one.unit:134, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Comp > one.unit:134, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpComp) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Element" { // one.unit:150, go-struct-rio.act:744
		for _, st := range me.ItsElement {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "parent" {
		if me.Kparentp >= 0 {
			st := glob.Dats.ApComp[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Comp_parent") { // one.unit:140, go-struct-rio.act:654
		for _, st := range glob.Dats.ApComp {
			if (st.Kparentp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	if (va[0] == "Element_comp") { // one.unit:168, go-struct-rio.act:654
		for _, st := range glob.Dats.ApElement {
			if (st.Kcompp == me.Me) {
				if len(va) > 1 {
					ret := st.DoIts(glob, va[1:], lno)
					if (ret != 0) {
						return(ret)
					}
					continue
				}
				ret := GoAct(glob, st)
				if (ret != 0) {
					return(ret)
				}
			}
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Comp %s,%s > one.unit:134, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpElement struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kcompp int
	ItsOpt [] *KpOpt 
	Childs [] Kp
}

func (me KpElement) TypeName() string {
    return me.Comp
}
func (me KpElement) GetLineNo() string {
	return me.LineNo
}

func loadElement(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpElement)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApElement)
	st.LineNo = lno
	st.Comp = "Element";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcompp = -1
	st.Kparentp = len( act.ApComp ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Element has no Comp parent\n") ;
		return 1
	}
	st.Parent = act.ApComp[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Element under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApComp[ len( act.ApComp )-1 ].Childs = append(act.ApComp[ len( act.ApComp )-1 ].Childs, st)
	act.ApComp[ len( act.ApComp )-1 ].ItsElement = append(act.ApComp[ len( act.ApComp )-1 ].ItsElement, st)	// one.unit:134, go-struct-rio.act:467
	name,_ := st.Names["name"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Element_" + name	// one.unit:152, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApElement = append(act.ApElement, st)
	return 0
}

func (me KpElement) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if va[0] == "comp" { // one.unit:168, go-struct-rio.act:680
		if (me.Kcompp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kcompp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // one.unit:134, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:150, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApElement[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Element > one.unit:150, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Element > one.unit:150, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpElement) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Opt" { // one.unit:182, go-struct-rio.act:744
		for _, st := range me.ItsOpt {
			if len(va) > 1 {
				ret := st.DoIts(glob, va[1:], lno)
				if (ret != 0) {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "parent" { // one.unit:134, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApComp[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "comp" {
		if me.Kcompp >= 0 {
			st := glob.Dats.ApComp[ me.Kcompp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Element %s,%s > one.unit:150, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpOpt struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
}

func (me KpOpt) TypeName() string {
    return me.Comp
}
func (me KpOpt) GetLineNo() string {
	return me.LineNo
}

func loadOpt(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpOpt)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApOpt)
	st.LineNo = lno
	st.Comp = "Opt";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApElement ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Opt has no Element parent\n") ;
		return 1
	}
	st.Parent = act.ApElement[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Opt under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApElement[ len( act.ApElement )-1 ].Childs = append(act.ApElement[ len( act.ApElement )-1 ].Childs, st)
	act.ApElement[ len( act.ApElement )-1 ].ItsOpt = append(act.ApElement[ len( act.ApElement )-1 ].ItsOpt, st)	// one.unit:150, go-struct-rio.act:467
	name,_ := st.Names["name"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Opt_" + name	// one.unit:188, go-struct-rio.act:515
	act.index[s] = st.Me;
	st.MyName = name
	act.ApOpt = append(act.ApOpt, st)
	return 0
}

func (me KpOpt) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	if (va[0] == "parent") { // one.unit:150, go-struct-rio.act:644
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApElement[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // one.unit:182, go-struct-rio.act:228
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApOpt[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Opt > one.unit:182, go-struct-rio.act:235?", va[0], lno, me.LineNo)
		return false, msg
	}
	if va[0] == "payload" {
		tmp := maps.Clone(me.Names)
		delete(tmp, "kMe")
		delete(tmp, "kComp")
		jsonData, _ := json.MarshalIndent(tmp, "   ", "  ")
		return true, string(jsonData)
	}
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Opt > one.unit:182, go-struct-rio.act:247?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]]
	return true,rr
}

func (me KpOpt) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // one.unit:150, go-struct-rio.act:629
		if me.Kparentp >= 0 {
			st := glob.Dats.ApElement[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Opt %s,%s > one.unit:182, go-struct-rio.act:273?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpActor struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kname string
	Kcomp string
	Kattr string
	Keq string
	Kvalue string
	Childs [] Kp
}

func (me KpActor) TypeName() string {
    return me.Comp
}
func (me KpActor) GetLineNo() string {
	return me.LineNo
}

func loadActor(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpActor)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApActor)
	st.LineNo = lno
	st.Comp = "Actor";
	st.Flags = flag;
	st.Kname = cnv( st.Names["name"] )
	st.Kcomp = cnv( st.Names["comp"] )
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	act.index["Actor_" + st.Kname] = st.Me;
	act.ApActor = append(act.ApActor, st)
	return 0
}

func (me KpActor) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Actor > act.unit:2, go-struct-rio.act:77?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}

func (me KpActor) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Childs" { // one.unit:20, go-struct-rio.act:685
		for _, st := range me.Childs {
			ret := GoAct(glob, st)
			if (ret != 0) {
				return(ret)
			}
		}
		return(0)
	}
	return(0)
}

type KpAll struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhat string
	Kactor string
	Kattr string
	Keq string
	Kvalue string
	Kargs string
	Kactorp int
}

func (me KpAll) TypeName() string {
    return me.Comp
}
func (me KpAll) GetLineNo() string {
	return me.LineNo
}

func loadAll(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpAll)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApAll)
	st.LineNo = lno
	st.Comp = "All";
	st.Flags = flag;
	st.Kwhat = cnv( st.Names["what"] )
	st.Kactor = cnv( st.Names["actor"] )
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	st.Kargs = cnv( st.Names["args"] )
	st.Kactorp = -1
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " All has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApAll = append(act.ApAll, st)
	return 0
}


func (me KpAll) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,All > act.unit:27, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpAll) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpDu struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kactor string
	Kattr string
	Keq string
	Kvalue string
	Kargs string
	Kactorp int
}

func (me KpDu) TypeName() string {
    return me.Comp
}
func (me KpDu) GetLineNo() string {
	return me.LineNo
}

func loadDu(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpDu)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApDu)
	st.LineNo = lno
	st.Comp = "Du";
	st.Flags = flag;
	st.Kactor = cnv( st.Names["actor"] )
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	st.Kargs = cnv( st.Names["args"] )
	st.Kactorp = -1
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Du has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApDu = append(act.ApDu, st)
	return 0
}


func (me KpDu) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Du > act.unit:43, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpDu) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpNew struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhere string
	Kwhat string
	Kline string
}

func (me KpNew) TypeName() string {
    return me.Comp
}
func (me KpNew) GetLineNo() string {
	return me.LineNo
}

func loadNew(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpNew)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApNew)
	st.LineNo = lno
	st.Comp = "New";
	st.Flags = flag;
	st.Kwhere = cnv( st.Names["where"] )
	st.Kwhat = cnv( st.Names["what"] )
	st.Kline = cnv( st.Names["line"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " New has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApNew = append(act.ApNew, st)
	return 0
}


func (me KpNew) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,New > act.unit:58, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpNew) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpRefs struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhere string
}

func (me KpRefs) TypeName() string {
    return me.Comp
}
func (me KpRefs) GetLineNo() string {
	return me.LineNo
}

func loadRefs(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpRefs)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApRefs)
	st.LineNo = lno
	st.Comp = "Refs";
	st.Flags = flag;
	st.Kwhere = cnv( st.Names["where"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Refs has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApRefs = append(act.ApRefs, st)
	return 0
}


func (me KpRefs) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Refs > act.unit:68, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpRefs) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpVar struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kattr string
	Keq string
	Kvalue string
}

func (me KpVar) TypeName() string {
    return me.Comp
}
func (me KpVar) GetLineNo() string {
	return me.LineNo
}

func loadVar(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpVar)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApVar)
	st.LineNo = lno
	st.Comp = "Var";
	st.Flags = flag;
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Var has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApVar = append(act.ApVar, st)
	return 0
}


func (me KpVar) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Var > act.unit:76, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpVar) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpIts struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhat string
	Kactor string
	Kattr string
	Keq string
	Kvalue string
	Kargs string
	Kactorp int
}

func (me KpIts) TypeName() string {
    return me.Comp
}
func (me KpIts) GetLineNo() string {
	return me.LineNo
}

func loadIts(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpIts)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApIts)
	st.LineNo = lno
	st.Comp = "Its";
	st.Flags = flag;
	st.Kwhat = cnv( st.Names["what"] )
	st.Kactor = cnv( st.Names["actor"] )
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	st.Kargs = cnv( st.Names["args"] )
	st.Kactorp = -1
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Its has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApIts = append(act.ApIts, st)
	return 0
}


func (me KpIts) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Its > act.unit:86, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpIts) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpC struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kdesc string
}

func (me KpC) TypeName() string {
    return me.Comp
}
func (me KpC) GetLineNo() string {
	return me.LineNo
}

func loadC(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpC)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApC)
	st.LineNo = lno
	st.Comp = "C";
	st.Flags = flag;
	st.Kdesc = cnv( st.Names["desc"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " C has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApC = append(act.ApC, st)
	return 0
}


func (me KpC) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,C > act.unit:102, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpC) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpCs struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kdesc string
}

func (me KpCs) TypeName() string {
    return me.Comp
}
func (me KpCs) GetLineNo() string {
	return me.LineNo
}

func loadCs(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCs)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApCs)
	st.LineNo = lno
	st.Comp = "Cs";
	st.Flags = flag;
	st.Kdesc = cnv( st.Names["desc"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Cs has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApCs = append(act.ApCs, st)
	return 0
}


func (me KpCs) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Cs > act.unit:110, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpCs) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpOut struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhat string
	Kdesc string
}

func (me KpOut) TypeName() string {
    return me.Comp
}
func (me KpOut) GetLineNo() string {
	return me.LineNo
}

func loadOut(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpOut)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApOut)
	st.LineNo = lno
	st.Comp = "Out";
	st.Flags = flag;
	st.Kwhat = cnv( st.Names["what"] )
	st.Kdesc = cnv( st.Names["desc"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Out has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApOut = append(act.ApOut, st)
	return 0
}


func (me KpOut) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Out > act.unit:118, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpOut) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpIn struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kflag string
}

func (me KpIn) TypeName() string {
    return me.Comp
}
func (me KpIn) GetLineNo() string {
	return me.LineNo
}

func loadIn(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpIn)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApIn)
	st.LineNo = lno
	st.Comp = "In";
	st.Flags = flag;
	st.Kflag = cnv( st.Names["flag"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " In has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApIn = append(act.ApIn, st)
	return 0
}


func (me KpIn) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,In > act.unit:134, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpIn) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpBreak struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kwhat string
	Kactor string
	Kcheck string
}

func (me KpBreak) TypeName() string {
    return me.Comp
}
func (me KpBreak) GetLineNo() string {
	return me.LineNo
}

func loadBreak(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpBreak)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApBreak)
	st.LineNo = lno
	st.Comp = "Break";
	st.Flags = flag;
	st.Kwhat = cnv( st.Names["what"] )
	st.Kactor = cnv( st.Names["actor"] )
	st.Kcheck = cnv( st.Names["check"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Break has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApBreak = append(act.ApBreak, st)
	return 0
}


func (me KpBreak) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Break > act.unit:142, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpBreak) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpAdd struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kpath string
	Kdata string
}

func (me KpAdd) TypeName() string {
    return me.Comp
}
func (me KpAdd) GetLineNo() string {
	return me.LineNo
}

func loadAdd(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpAdd)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApAdd)
	st.LineNo = lno
	st.Comp = "Add";
	st.Flags = flag;
	st.Kpath = cnv( st.Names["path"] )
	st.Kdata = cnv( st.Names["data"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Add has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApAdd = append(act.ApAdd, st)
	return 0
}


func (me KpAdd) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Add > act.unit:158, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpAdd) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpThis struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kpath string
	Kactor string
	Kattr string
	Keq string
	Kvalue string
	Kargs string
	Kactorp int
}

func (me KpThis) TypeName() string {
    return me.Comp
}
func (me KpThis) GetLineNo() string {
	return me.LineNo
}

func loadThis(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpThis)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApThis)
	st.LineNo = lno
	st.Comp = "This";
	st.Flags = flag;
	st.Kpath = cnv( st.Names["path"] )
	st.Kactor = cnv( st.Names["actor"] )
	st.Kattr = cnv( st.Names["attr"] )
	st.Keq = cnv( st.Names["eq"] )
	st.Kvalue = cnv( st.Names["value"] )
	st.Kargs = cnv( st.Names["args"] )
	st.Kactorp = -1
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " This has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApThis = append(act.ApThis, st)
	return 0
}


func (me KpThis) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,This > act.unit:185, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpThis) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
type KpReplace struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kpath string
	Kwith string
	Kmatch string
}

func (me KpReplace) TypeName() string {
    return me.Comp
}
func (me KpReplace) GetLineNo() string {
	return me.LineNo
}

func loadReplace(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpReplace)
	st.Names = names
	st.MyName = ""
	st.Parent = ""
	st.Me = len(act.ApReplace)
	st.LineNo = lno
	st.Comp = "Replace";
	st.Flags = flag;
	st.Kpath = cnv( st.Names["path"] )
	st.Kwith = cnv( st.Names["with"] )
	st.Kmatch = cnv( st.Names["match"] )
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Replace has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApReplace = append(act.ApReplace, st)
	return 0
}


func (me KpReplace) GetVar(glob *GlobT, va []string, lno string) (bool, interface{}) {
	r := me.Names[va[0]]
	if r == nil { 
		rr := fmt.Sprintf("?%s?:%s,%s,Replace > act.unit:202, go-struct-rio.act:152?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := cnv( me.Names[va[0]] )
	return true,rr
}
func (me KpReplace) DoIts(glob *GlobT, va []string, lno string) int {
	return 0
}
