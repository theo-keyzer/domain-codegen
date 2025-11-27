package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"maps"
)

type Kp interface {
	DoIts(glob *GlobT, va []string, lno string) int
	GetVar(glob *GlobT, va []string, lno string) (bool, string)
	GetLineNo() string
	TypeName() string
}

type KpActor struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kname string
	Kcomp string
	Kattr string
	Keq string
	Kvalue string
	Childs [] Kp
}

func loadActor(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpActor)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApActor)
	st.LineNo = lno
	st.Comp = "Actor";
	st.Flags = flag;
	p,st.Kname = getw(ln,p)
	p,st.Kcomp = getw(ln,p)
	p,st.Kattr = getw(ln,p)
	p,st.Keq = getw(ln,p)
	p,st.Kvalue = getws(ln,p)
	act.index["Actor_" + st.Kname] = st.Me;
	act.ApActor = append(act.ApActor, st)
	return 0
}

type KpAll struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kparentp int
	Kwhat string
	Kactor string
	Kargs string
	Kactorp int
}

func loadAll(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpAll)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApAll)
	st.LineNo = lno
	st.Comp = "All";
	st.Flags = flag;
	p,st.Kwhat = getw(ln,p)
	p,st.Kactor = getw(ln,p)
	p,st.Kargs = getws(ln,p)
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

type KpDu struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kparentp int
	Kactor string
	Kargs string
	Kactorp int
}

func loadDu(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpDu)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApDu)
	st.LineNo = lno
	st.Comp = "Du";
	st.Flags = flag;
	p,st.Kactor = getw(ln,p)
	p,st.Kargs = getws(ln,p)
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

type KpNew struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kparentp int
	Kwhere string
	Kwhat string
	Kline string
}

func loadNew(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpNew)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApNew)
	st.LineNo = lno
	st.Comp = "New";
	st.Flags = flag;
	p,st.Kwhere = getw(ln,p)
	p,st.Kwhat = getw(ln,p)
	p,st.Kline = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " New has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApNew = append(act.ApNew, st)
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
	Kparentp int
	Kwhere string
}

func loadRefs(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpRefs)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApRefs)
	st.LineNo = lno
	st.Comp = "Refs";
	st.Flags = flag;
	p,st.Kwhere = getw(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Refs has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApRefs = append(act.ApRefs, st)
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
	Kparentp int
	Kattr string
	Keq string
	Kvalue string
}

func loadVar(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpVar)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApVar)
	st.LineNo = lno
	st.Comp = "Var";
	st.Flags = flag;
	p,st.Kattr = getw(ln,p)
	p,st.Keq = getw(ln,p)
	p,st.Kvalue = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Var has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApVar = append(act.ApVar, st)
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
	Kparentp int
	Kwhat string
	Kactor string
	Kargs string
	Kactorp int
}

func loadIts(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpIts)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApIts)
	st.LineNo = lno
	st.Comp = "Its";
	st.Flags = flag;
	p,st.Kwhat = getw(ln,p)
	p,st.Kactor = getw(ln,p)
	p,st.Kargs = getws(ln,p)
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

type KpC struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kparentp int
	Kdesc string
}

func loadC(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpC)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApC)
	st.LineNo = lno
	st.Comp = "C";
	st.Flags = flag;
	p,st.Kdesc = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " C has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApC = append(act.ApC, st)
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
	Kparentp int
	Kdesc string
}

func loadCs(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpCs)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCs)
	st.LineNo = lno
	st.Comp = "Cs";
	st.Flags = flag;
	p,st.Kdesc = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Cs has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApCs = append(act.ApCs, st)
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
	Kparentp int
	Kwhat string
	Kpad string
	Kdesc string
}

func loadOut(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpOut)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApOut)
	st.LineNo = lno
	st.Comp = "Out";
	st.Flags = flag;
	p,st.Kwhat = getw(ln,p)
	p,st.Kpad = getw(ln,p)
	p,st.Kdesc = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Out has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApOut = append(act.ApOut, st)
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
	Kparentp int
	Kflag string
}

func loadIn(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpIn)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApIn)
	st.LineNo = lno
	st.Comp = "In";
	st.Flags = flag;
	p,st.Kflag = getw(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " In has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApIn = append(act.ApIn, st)
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
	Kparentp int
	Kwhat string
	Kpad string
	Kactor string
	Kcheck string
}

func loadBreak(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpBreak)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApBreak)
	st.LineNo = lno
	st.Comp = "Break";
	st.Flags = flag;
	p,st.Kwhat = getw(ln,p)
	p,st.Kpad = getw(ln,p)
	p,st.Kactor = getw(ln,p)
	p,st.Kcheck = getw(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Break has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApBreak = append(act.ApBreak, st)
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
	Kparentp int
	Kpath string
	Kdata string
}

func loadAdd(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpAdd)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApAdd)
	st.LineNo = lno
	st.Comp = "Add";
	st.Flags = flag;
	p,st.Kpath = getw(ln,p)
	p,st.Kdata = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Add has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApAdd = append(act.ApAdd, st)
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
	Kparentp int
	Kpath string
	Kactor string
	Kargs string
	Kactorp int
}

func loadThis(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpThis)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApThis)
	st.LineNo = lno
	st.Comp = "This";
	st.Flags = flag;
	p,st.Kpath = getw(ln,p)
	p,st.Kactor = getw(ln,p)
	p,st.Kargs = getws(ln,p)
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

type KpReplace struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Kparentp int
	Kpath string
	Kpad string
	Kwith string
	Kpad2 string
	Kmatch string
}

func loadReplace(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpReplace)
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApReplace)
	st.LineNo = lno
	st.Comp = "Replace";
	st.Flags = flag;
	p,st.Kpath = getw(ln,p)
	p,st.Kpad = getw(ln,p)
	p,st.Kwith = getw(ln,p)
	p,st.Kpad2 = getw(ln,p)
	p,st.Kmatch = getws(ln,p)
	st.Kparentp = len(act.ApActor)-1;
	if (st.Kparentp < 0 ) { 
		print(lno + " Replace has no Actor parent\n") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApReplace = append(act.ApReplace, st)
	return 0
}

type KpAgent struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ktoolsp int
	ItsObjective [] *KpObjective 
	ItsMemory [] *KpMemory 
	Childs [] Kp
}

func (me KpAgent) TypeName() string {
    return me.Comp
}
func (me KpAgent) GetLineNo() string {
	return me.LineNo
}

func loadAgent(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpAgent)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApAgent)
	st.LineNo = lno
	st.Comp = "Agent";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ktoolsp = -1
	name,_ := st.Names["agent_id"].(string)
	st.Names["_key"] = "agent_id"
	act.index["Agent_" + name] = st.Me;
	st.MyName = name
	act.ApAgent = append(act.ApAgent, st)
	return 0
}

func (me KpAgent) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "tools" { // cog.unit:17, go-struct-rio.act:621
		if (me.Ktoolsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApTool[ me.Ktoolsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Citation_agent" && len(va) > 1) { // cog.unit:90, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCitation {
			if (st.Kagentp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Action_agent" && len(va) > 1) { // cog.unit:111, go-struct-rio.act:706
		for _, st := range glob.Dats.ApAction {
			if (st.Kagentp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Session_agent_id" && len(va) > 1) { // session.unit:14, go-struct-rio.act:706
		for _, st := range glob.Dats.ApSession {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Snapshot_agent_id" && len(va) > 1) { // session.unit:35, go-struct-rio.act:706
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Restoration_agent_id" && len(va) > 1) { // session.unit:53, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRestoration {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Restoration_agent_id2" && len(va) > 1) { // session.unit:56, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRestoration {
			if (st.Kagent_id2p == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Cycle_agent_id" && len(va) > 1) { // session.unit:89, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCycle {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Recommendation_agent_id" && len(va) > 1) { // session.unit:114, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenExecution_agent_id" && len(va) > 1) { // session.unit:234, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Kagent_idp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // cog.unit:14, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Agent > cog.unit:14, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Agent > cog.unit:14, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpAgent) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Objective" { // cog.unit:24, go-struct-rio.act:685
		for _, st := range me.ItsObjective {
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
	if va[0] == "Memory" { // cog.unit:40, go-struct-rio.act:685
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
	if va[0] == "tools" {
		if me.Ktoolsp >= 0 {
			st := glob.Dats.ApTool[ me.Ktoolsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Citation_agent") { // cog.unit:90, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCitation {
			if (st.Kagentp == me.Me) {
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
	if (va[0] == "Action_agent") { // cog.unit:111, go-struct-rio.act:595
		for _, st := range glob.Dats.ApAction {
			if (st.Kagentp == me.Me) {
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
	if (va[0] == "Session_agent_id") { // session.unit:14, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSession {
			if (st.Kagent_idp == me.Me) {
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
	if (va[0] == "Snapshot_agent_id") { // session.unit:35, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Kagent_idp == me.Me) {
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
	if (va[0] == "Restoration_agent_id") { // session.unit:53, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Kagent_idp == me.Me) {
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
	if (va[0] == "Restoration_agent_id2") { // session.unit:56, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Kagent_id2p == me.Me) {
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
	if (va[0] == "Cycle_agent_id") { // session.unit:89, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Kagent_idp == me.Me) {
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
	if (va[0] == "Recommendation_agent_id") { // session.unit:114, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kagent_idp == me.Me) {
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
	if (va[0] == "GenExecution_agent_id") { // session.unit:234, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Kagent_idp == me.Me) {
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
	        fmt.Printf("?No its %s for Agent %s,%s > cog.unit:14, go-struct-rio.act:222?", va[0], lno, me.LineNo)
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
	Kparentp int
	Kparent_objp int
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
	st.Kparent_objp = -1
	st.Kparentp = len( act.ApAgent ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Objective has no Agent parent\n") ;
		return 1
	}
	st.Parent = act.ApAgent[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Objective under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApAgent[ len( act.ApAgent )-1 ].Childs = append(act.ApAgent[ len( act.ApAgent )-1 ].Childs, st)
	act.ApAgent[ len( act.ApAgent )-1 ].ItsObjective = append(act.ApAgent[ len( act.ApAgent )-1 ].ItsObjective, st)	// cog.unit:14, go-struct-rio.act:416
	name,_ := st.Names["objective_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Objective_" + name	// cog.unit:25, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApObjective = append(act.ApObjective, st)
	return 0
}

func (me KpObjective) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "parent_obj" { // cog.unit:26, go-struct-rio.act:621
		if (me.Kparent_objp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kparent_objp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Objective_parent_obj" && len(va) > 1) { // cog.unit:26, go-struct-rio.act:706
		for _, st := range glob.Dats.ApObjective {
			if (st.Kparent_objp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Memory_objective" && len(va) > 1) { // cog.unit:47, go-struct-rio.act:706
		for _, st := range glob.Dats.ApMemory {
			if (st.Kobjectivep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // cog.unit:24, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Objective > cog.unit:24, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Objective > cog.unit:24, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpObjective) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Thought" { // cog.unit:65, go-struct-rio.act:685
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
	if va[0] == "parent" { // cog.unit:14, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApAgent[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "parent_obj" {
		if me.Kparent_objp >= 0 {
			st := glob.Dats.ApObjective[ me.Kparent_objp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Objective_parent_obj") { // cog.unit:26, go-struct-rio.act:595
		for _, st := range glob.Dats.ApObjective {
			if (st.Kparent_objp == me.Me) {
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
	if (va[0] == "Memory_objective") { // cog.unit:47, go-struct-rio.act:595
		for _, st := range glob.Dats.ApMemory {
			if (st.Kobjectivep == me.Me) {
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
	if (va[0] == "Snapshot_objectives") { // session.unit:36, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Kobjectivesp == me.Me) {
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
	if (va[0] == "Cycle_objective") { // session.unit:90, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Kobjectivep == me.Me) {
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
	if (va[0] == "Recommendation_objective") { // session.unit:117, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kobjectivep == me.Me) {
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
	        fmt.Printf("?No its %s for Objective %s,%s > cog.unit:24, go-struct-rio.act:222?", va[0], lno, me.LineNo)
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
	Kembeddingp int
	Kobjectivep int
	Kinvalidated_byp int
	ItsMemorySource [] *KpMemorySource 
	Childs [] Kp
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
	st.Kembeddingp = -1
	st.Kobjectivep = -1
	st.Kinvalidated_byp = -1
	st.Kparentp = len( act.ApAgent ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Memory has no Agent parent\n") ;
		return 1
	}
	st.Parent = act.ApAgent[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Memory under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApAgent[ len( act.ApAgent )-1 ].Childs = append(act.ApAgent[ len( act.ApAgent )-1 ].Childs, st)
	act.ApAgent[ len( act.ApAgent )-1 ].ItsMemory = append(act.ApAgent[ len( act.ApAgent )-1 ].ItsMemory, st)	// cog.unit:14, go-struct-rio.act:416
	name,_ := st.Names["memory_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Memory_" + name	// cog.unit:41, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApMemory = append(act.ApMemory, st)
	return 0
}

func (me KpMemory) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "embedding" { // cog.unit:46, go-struct-rio.act:621
		if (me.Kembeddingp >= 0 && len(va) > 1) {
			return( glob.Dats.ApVector[ me.Kembeddingp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "objective" { // cog.unit:47, go-struct-rio.act:621
		if (me.Kobjectivep >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kobjectivep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "invalidated_by" { // cog.unit:48, go-struct-rio.act:621
		if (me.Kinvalidated_byp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kinvalidated_byp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // cog.unit:40, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Memory > cog.unit:40, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Memory > cog.unit:40, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpMemory) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "MemorySource" { // cog.unit:57, go-struct-rio.act:685
		for _, st := range me.ItsMemorySource {
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
	if va[0] == "parent" { // cog.unit:14, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApAgent[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "embedding" {
		if me.Kembeddingp >= 0 {
			st := glob.Dats.ApVector[ me.Kembeddingp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "objective" {
		if me.Kobjectivep >= 0 {
			st := glob.Dats.ApObjective[ me.Kobjectivep ]
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
	if (va[0] == "Citation_memory") { // cog.unit:91, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCitation {
			if (st.Kmemoryp == me.Me) {
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
	if (va[0] == "Action_result_mem") { // cog.unit:112, go-struct-rio.act:595
		for _, st := range glob.Dats.ApAction {
			if (st.Kresult_memp == me.Me) {
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
	if (va[0] == "Snapshot_memories") { // session.unit:39, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Kmemoriesp == me.Me) {
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
	if (va[0] == "Restoration_source_memory") { // session.unit:54, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Ksource_memoryp == me.Me) {
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
	if (va[0] == "Restoration_restored_as") { // session.unit:57, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Krestored_asp == me.Me) {
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
	if (va[0] == "Cycle_baseline") { // session.unit:93, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Kbaselinep == me.Me) {
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
	if (va[0] == "Cycle_outcome") { // session.unit:97, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Koutcomep == me.Me) {
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
	if (va[0] == "Recommendation_outcome") { // session.unit:115, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Koutcomep == me.Me) {
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
	        fmt.Printf("?No its %s for Memory %s,%s > cog.unit:40, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpMemorySource struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kimplp int
	Kcanonp int
}

func (me KpMemorySource) TypeName() string {
    return me.Comp
}
func (me KpMemorySource) GetLineNo() string {
	return me.LineNo
}

func loadMemorySource(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpMemorySource)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApMemorySource)
	st.LineNo = lno
	st.Comp = "MemorySource";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kimplp = -1
	st.Kcanonp = -1
	st.Kparentp = len( act.ApMemory ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " MemorySource has no Memory parent\n") ;
		return 1
	}
	st.Parent = act.ApMemory[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " MemorySource under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApMemory[ len( act.ApMemory )-1 ].Childs = append(act.ApMemory[ len( act.ApMemory )-1 ].Childs, st)
	act.ApMemory[ len( act.ApMemory )-1 ].ItsMemorySource = append(act.ApMemory[ len( act.ApMemory )-1 ].ItsMemorySource, st)	// cog.unit:40, go-struct-rio.act:416
	act.ApMemorySource = append(act.ApMemorySource, st)
	return 0
}

func (me KpMemorySource) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "impl" { // cog.unit:58, go-struct-rio.act:621
		if (me.Kimplp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "canon" { // cog.unit:59, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:40, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // cog.unit:57, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMemorySource[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,MemorySource > cog.unit:57, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,MemorySource > cog.unit:57, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpMemorySource) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // cog.unit:40, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApMemory[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "impl" {
		if me.Kimplp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for MemorySource %s,%s > cog.unit:57, go-struct-rio.act:222?", va[0], lno, me.LineNo)
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
	Kprevp int
	Kbranchesp int
	Kparent_branchp int
	ItsCitation [] *KpCitation 
	ItsAction [] *KpAction 
	ItsThoughtSource [] *KpThoughtSource 
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
	st.Kprevp = -1
	st.Kbranchesp = -1
	st.Kparent_branchp = -1
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
	act.ApObjective[ len( act.ApObjective )-1 ].ItsThought = append(act.ApObjective[ len( act.ApObjective )-1 ].ItsThought, st)	// cog.unit:24, go-struct-rio.act:416
	name,_ := st.Names["thought_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Thought_" + name	// cog.unit:66, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApThought = append(act.ApThought, st)
	return 0
}

func (me KpThought) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "prev" { // cog.unit:69, go-struct-rio.act:621
		if (me.Kprevp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kprevp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "branches" { // cog.unit:70, go-struct-rio.act:621
		if (me.Kbranchesp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kbranchesp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "parent_branch" { // cog.unit:71, go-struct-rio.act:621
		if (me.Kparent_branchp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kparent_branchp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:24, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Thought_prev" && len(va) > 1) { // cog.unit:69, go-struct-rio.act:706
		for _, st := range glob.Dats.ApThought {
			if (st.Kprevp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Thought_branches" && len(va) > 1) { // cog.unit:70, go-struct-rio.act:706
		for _, st := range glob.Dats.ApThought {
			if (st.Kbranchesp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Thought_parent_branch" && len(va) > 1) { // cog.unit:71, go-struct-rio.act:706
		for _, st := range glob.Dats.ApThought {
			if (st.Kparent_branchp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // cog.unit:65, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Thought > cog.unit:65, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Thought > cog.unit:65, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpThought) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Citation" { // cog.unit:88, go-struct-rio.act:685
		for _, st := range me.ItsCitation {
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
	if va[0] == "Action" { // cog.unit:107, go-struct-rio.act:685
		for _, st := range me.ItsAction {
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
	if va[0] == "ThoughtSource" { // cog.unit:125, go-struct-rio.act:685
		for _, st := range me.ItsThoughtSource {
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
	if va[0] == "parent" { // cog.unit:24, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApObjective[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "prev" {
		if me.Kprevp >= 0 {
			st := glob.Dats.ApThought[ me.Kprevp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "branches" {
		if me.Kbranchesp >= 0 {
			st := glob.Dats.ApThought[ me.Kbranchesp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "parent_branch" {
		if me.Kparent_branchp >= 0 {
			st := glob.Dats.ApThought[ me.Kparent_branchp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Memory_invalidated_by") { // cog.unit:48, go-struct-rio.act:595
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
	if (va[0] == "Thought_prev") { // cog.unit:69, go-struct-rio.act:595
		for _, st := range glob.Dats.ApThought {
			if (st.Kprevp == me.Me) {
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
	if (va[0] == "Thought_branches") { // cog.unit:70, go-struct-rio.act:595
		for _, st := range glob.Dats.ApThought {
			if (st.Kbranchesp == me.Me) {
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
	if (va[0] == "Thought_parent_branch") { // cog.unit:71, go-struct-rio.act:595
		for _, st := range glob.Dats.ApThought {
			if (st.Kparent_branchp == me.Me) {
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
	if (va[0] == "Snapshot_thoughts") { // session.unit:37, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Kthoughtsp == me.Me) {
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
	if (va[0] == "Cycle_hypothesis") { // session.unit:91, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Khypothesisp == me.Me) {
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
	if (va[0] == "Cycle_intervention") { // session.unit:95, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Kinterventionp == me.Me) {
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
		if va[0] == "level" && len(va) > 1 && me.Kparentp >= 0 { // cog.unit:67, go-struct-rio.act:233
			pos := 0
			v, ok := me.Names["level"].(string)
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
				pst := glob.Dats.ApObjective[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsThought {
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
	        fmt.Printf("?No its %s for Thought %s,%s > cog.unit:65, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCitation struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kagentp int
	Kmemoryp int
}

func (me KpCitation) TypeName() string {
    return me.Comp
}
func (me KpCitation) GetLineNo() string {
	return me.LineNo
}

func loadCitation(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCitation)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCitation)
	st.LineNo = lno
	st.Comp = "Citation";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kagentp = -1
	st.Kmemoryp = -1
	st.Kparentp = len( act.ApThought ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Citation has no Thought parent\n") ;
		return 1
	}
	st.Parent = act.ApThought[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Citation under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApThought[ len( act.ApThought )-1 ].Childs = append(act.ApThought[ len( act.ApThought )-1 ].Childs, st)
	act.ApThought[ len( act.ApThought )-1 ].ItsCitation = append(act.ApThought[ len( act.ApThought )-1 ].ItsCitation, st)	// cog.unit:65, go-struct-rio.act:416
	name,_ := st.Names["citation_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Citation_" + name	// cog.unit:89, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApCitation = append(act.ApCitation, st)
	return 0
}

func (me KpCitation) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "agent" { // cog.unit:90, go-struct-rio.act:621
		if (me.Kagentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagentp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "memory" { // cog.unit:91, go-struct-rio.act:621
		if (me.Kmemoryp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kmemoryp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:65, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // cog.unit:88, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCitation[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Citation > cog.unit:88, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Citation > cog.unit:88, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCitation) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // cog.unit:65, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApThought[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent" {
		if me.Kagentp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "memory" {
		if me.Kmemoryp >= 0 {
			st := glob.Dats.ApMemory[ me.Kmemoryp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Citation %s,%s > cog.unit:88, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpAction struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ktoolp int
	Kagentp int
	Kresult_memp int
}

func (me KpAction) TypeName() string {
    return me.Comp
}
func (me KpAction) GetLineNo() string {
	return me.LineNo
}

func loadAction(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpAction)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApAction)
	st.LineNo = lno
	st.Comp = "Action";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ktoolp = -1
	st.Kagentp = -1
	st.Kresult_memp = -1
	st.Kparentp = len( act.ApThought ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Action has no Thought parent\n") ;
		return 1
	}
	st.Parent = act.ApThought[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Action under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApThought[ len( act.ApThought )-1 ].Childs = append(act.ApThought[ len( act.ApThought )-1 ].Childs, st)
	act.ApThought[ len( act.ApThought )-1 ].ItsAction = append(act.ApThought[ len( act.ApThought )-1 ].ItsAction, st)	// cog.unit:65, go-struct-rio.act:416
	name,_ := st.Names["action_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Action_" + name	// cog.unit:108, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApAction = append(act.ApAction, st)
	return 0
}

func (me KpAction) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "tool" { // cog.unit:109, go-struct-rio.act:621
		if (me.Ktoolp >= 0 && len(va) > 1) {
			return( glob.Dats.ApTool[ me.Ktoolp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent" { // cog.unit:111, go-struct-rio.act:621
		if (me.Kagentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagentp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "result_mem" { // cog.unit:112, go-struct-rio.act:621
		if (me.Kresult_memp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kresult_memp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:65, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // cog.unit:107, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApAction[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Action > cog.unit:107, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Action > cog.unit:107, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpAction) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // cog.unit:65, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApThought[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "tool" {
		if me.Ktoolp >= 0 {
			st := glob.Dats.ApTool[ me.Ktoolp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent" {
		if me.Kagentp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "result_mem" {
		if me.Kresult_memp >= 0 {
			st := glob.Dats.ApMemory[ me.Kresult_memp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Action %s,%s > cog.unit:107, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpThoughtSource struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kimplp int
	Kcanonp int
}

func (me KpThoughtSource) TypeName() string {
    return me.Comp
}
func (me KpThoughtSource) GetLineNo() string {
	return me.LineNo
}

func loadThoughtSource(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpThoughtSource)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApThoughtSource)
	st.LineNo = lno
	st.Comp = "ThoughtSource";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kimplp = -1
	st.Kcanonp = -1
	st.Kparentp = len( act.ApThought ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " ThoughtSource has no Thought parent\n") ;
		return 1
	}
	st.Parent = act.ApThought[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " ThoughtSource under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApThought[ len( act.ApThought )-1 ].Childs = append(act.ApThought[ len( act.ApThought )-1 ].Childs, st)
	act.ApThought[ len( act.ApThought )-1 ].ItsThoughtSource = append(act.ApThought[ len( act.ApThought )-1 ].ItsThoughtSource, st)	// cog.unit:65, go-struct-rio.act:416
	act.ApThoughtSource = append(act.ApThoughtSource, st)
	return 0
}

func (me KpThoughtSource) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "impl" { // cog.unit:126, go-struct-rio.act:621
		if (me.Kimplp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "canon" { // cog.unit:127, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // cog.unit:65, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // cog.unit:125, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApThoughtSource[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,ThoughtSource > cog.unit:125, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,ThoughtSource > cog.unit:125, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpThoughtSource) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // cog.unit:65, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApThought[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "impl" {
		if me.Kimplp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for ThoughtSource %s,%s > cog.unit:125, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpTool struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
}

func (me KpTool) TypeName() string {
    return me.Comp
}
func (me KpTool) GetLineNo() string {
	return me.LineNo
}

func loadTool(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpTool)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApTool)
	st.LineNo = lno
	st.Comp = "Tool";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["tool_id"].(string)
	st.Names["_key"] = "tool_id"
	act.index["Tool_" + name] = st.Me;
	st.MyName = name
	act.ApTool = append(act.ApTool, st)
	return 0
}

func (me KpTool) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "Agent_tools" && len(va) > 1) { // cog.unit:17, go-struct-rio.act:706
		for _, st := range glob.Dats.ApAgent {
			if (st.Ktoolsp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Action_tool" && len(va) > 1) { // cog.unit:109, go-struct-rio.act:706
		for _, st := range glob.Dats.ApAction {
			if (st.Ktoolp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // cog.unit:133, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApTool[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Tool > cog.unit:133, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Tool > cog.unit:133, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpTool) DoIts(glob *GlobT, va []string, lno string) int {
	if (va[0] == "Agent_tools") { // cog.unit:17, go-struct-rio.act:595
		for _, st := range glob.Dats.ApAgent {
			if (st.Ktoolsp == me.Me) {
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
	if (va[0] == "Action_tool") { // cog.unit:109, go-struct-rio.act:595
		for _, st := range glob.Dats.ApAction {
			if (st.Ktoolp == me.Me) {
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
	        fmt.Printf("?No its %s for Tool %s,%s > cog.unit:133, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpVector struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
}

func (me KpVector) TypeName() string {
    return me.Comp
}
func (me KpVector) GetLineNo() string {
	return me.LineNo
}

func loadVector(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpVector)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApVector)
	st.LineNo = lno
	st.Comp = "Vector";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["vector_id"].(string)
	st.Names["_key"] = "vector_id"
	act.index["Vector_" + name] = st.Me;
	st.MyName = name
	act.ApVector = append(act.ApVector, st)
	return 0
}

func (me KpVector) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "Memory_embedding" && len(va) > 1) { // cog.unit:46, go-struct-rio.act:706
		for _, st := range glob.Dats.ApMemory {
			if (st.Kembeddingp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // cog.unit:142, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApVector[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Vector > cog.unit:142, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Vector > cog.unit:142, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpVector) DoIts(glob *GlobT, va []string, lno string) int {
	if (va[0] == "Memory_embedding") { // cog.unit:46, go-struct-rio.act:595
		for _, st := range glob.Dats.ApMemory {
			if (st.Kembeddingp == me.Me) {
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
	        fmt.Printf("?No its %s for Vector %s,%s > cog.unit:142, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpSession struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kagent_idp int
	Kagent_id2p int
	Kparent_sessionp int
	ItsSnapshot [] *KpSnapshot 
	ItsRestoration [] *KpRestoration 
	ItsCrossRef [] *KpCrossRef 
	ItsCycle [] *KpCycle 
	ItsRecommendation [] *KpRecommendation 
	ItsMetricS [] *KpMetricS 
	ItsArtifact [] *KpArtifact 
	ItsGenExecution [] *KpGenExecution 
	Childs [] Kp
}

func (me KpSession) TypeName() string {
    return me.Comp
}
func (me KpSession) GetLineNo() string {
	return me.LineNo
}

func loadSession(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpSession)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApSession)
	st.LineNo = lno
	st.Comp = "Session";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kagent_idp = -1
	st.Kagent_id2p = -1
	st.Kparent_sessionp = -1
	name,_ := st.Names["session_id"].(string)
	st.Names["_key"] = "session_id"
	act.index["Session_" + name] = st.Me;
	st.MyName = name
	act.ApSession = append(act.ApSession, st)
	return 0
}

func (me KpSession) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "agent_id" { // session.unit:14, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "parent_session" { // session.unit:16, go-struct-rio.act:621
		if (me.Kparent_sessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparent_sessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Session_parent_session" && len(va) > 1) { // session.unit:16, go-struct-rio.act:706
		for _, st := range glob.Dats.ApSession {
			if (st.Kparent_sessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Snapshot_session" && len(va) > 1) { // session.unit:34, go-struct-rio.act:706
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Restoration_source_session" && len(va) > 1) { // session.unit:52, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRestoration {
			if (st.Ksource_sessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Restoration_session" && len(va) > 1) { // session.unit:55, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRestoration {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Cycle_session" && len(va) > 1) { // session.unit:88, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCycle {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Recommendation_session" && len(va) > 1) { // session.unit:113, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Pattern_discovered_in" && len(va) > 1) { // session.unit:154, go-struct-rio.act:706
		for _, st := range glob.Dats.ApPattern {
			if (st.Kdiscovered_inp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Evolution_sessions" && len(va) > 1) { // session.unit:179, go-struct-rio.act:706
		for _, st := range glob.Dats.ApEvolution {
			if (st.Ksessionsp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenExecution_session" && len(va) > 1) { // session.unit:233, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenLearning_session" && len(va) > 1) { // session.unit:290, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenLearning {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Implementation_session" && len(va) > 1) { // canon.unit:71, go-struct-rio.act:706
		for _, st := range glob.Dats.ApImplementation {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "CanonVersion_session" && len(va) > 1) { // canon.unit:102, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Ksessionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // session.unit:12, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Session > session.unit:12, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Session > session.unit:12, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpSession) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Snapshot" { // session.unit:30, go-struct-rio.act:685
		for _, st := range me.ItsSnapshot {
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
	if va[0] == "Restoration" { // session.unit:50, go-struct-rio.act:685
		for _, st := range me.ItsRestoration {
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
	if va[0] == "CrossRef" { // session.unit:64, go-struct-rio.act:685
		for _, st := range me.ItsCrossRef {
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
	if va[0] == "Cycle" { // session.unit:85, go-struct-rio.act:685
		for _, st := range me.ItsCycle {
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
	if va[0] == "Recommendation" { // session.unit:107, go-struct-rio.act:685
		for _, st := range me.ItsRecommendation {
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
	if va[0] == "MetricS" { // session.unit:192, go-struct-rio.act:685
		for _, st := range me.ItsMetricS {
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
	if va[0] == "Artifact" { // session.unit:204, go-struct-rio.act:685
		for _, st := range me.ItsArtifact {
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
	if va[0] == "GenExecution" { // session.unit:230, go-struct-rio.act:685
		for _, st := range me.ItsGenExecution {
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
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "parent_session" {
		if me.Kparent_sessionp >= 0 {
			st := glob.Dats.ApSession[ me.Kparent_sessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Session_parent_session") { // session.unit:16, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSession {
			if (st.Kparent_sessionp == me.Me) {
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
	if (va[0] == "Snapshot_session") { // session.unit:34, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSnapshot {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "Restoration_source_session") { // session.unit:52, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Ksource_sessionp == me.Me) {
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
	if (va[0] == "Restoration_session") { // session.unit:55, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRestoration {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "Cycle_session") { // session.unit:88, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCycle {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "Recommendation_session") { // session.unit:113, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "Pattern_discovered_in") { // session.unit:154, go-struct-rio.act:595
		for _, st := range glob.Dats.ApPattern {
			if (st.Kdiscovered_inp == me.Me) {
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
	if (va[0] == "Evolution_sessions") { // session.unit:179, go-struct-rio.act:595
		for _, st := range glob.Dats.ApEvolution {
			if (st.Ksessionsp == me.Me) {
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
	if (va[0] == "GenExecution_session") { // session.unit:233, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "GenLearning_session") { // session.unit:290, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenLearning {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "Implementation_session") { // canon.unit:71, go-struct-rio.act:595
		for _, st := range glob.Dats.ApImplementation {
			if (st.Ksessionp == me.Me) {
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
	if (va[0] == "CanonVersion_session") { // canon.unit:102, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Ksessionp == me.Me) {
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
	        fmt.Printf("?No its %s for Session %s,%s > session.unit:12, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpSnapshot struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksessionp int
	Kagent_idp int
	Kobjectivesp int
	Kthoughtsp int
	Kagent_id2p int
	Kmemoriesp int
}

func (me KpSnapshot) TypeName() string {
    return me.Comp
}
func (me KpSnapshot) GetLineNo() string {
	return me.LineNo
}

func loadSnapshot(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpSnapshot)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApSnapshot)
	st.LineNo = lno
	st.Comp = "Snapshot";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksessionp = -1
	st.Kagent_idp = -1
	st.Kobjectivesp = -1
	st.Kthoughtsp = -1
	st.Kagent_id2p = -1
	st.Kmemoriesp = -1
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Snapshot has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Snapshot under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsSnapshot = append(act.ApSession[ len( act.ApSession )-1 ].ItsSnapshot, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["snapshot_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Snapshot_" + name	// session.unit:31, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApSnapshot = append(act.ApSnapshot, st)
	return 0
}

func (me KpSnapshot) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "session" { // session.unit:34, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id" { // session.unit:35, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "objectives" { // session.unit:36, go-struct-rio.act:621
		if (me.Kobjectivesp >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kobjectivesp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "thoughts" { // session.unit:37, go-struct-rio.act:621
		if (me.Kthoughtsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kthoughtsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "memories" { // session.unit:39, go-struct-rio.act:621
		if (me.Kmemoriesp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kmemoriesp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:30, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSnapshot[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Snapshot > session.unit:30, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Snapshot > session.unit:30, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpSnapshot) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "objectives" {
		if me.Kobjectivesp >= 0 {
			st := glob.Dats.ApObjective[ me.Kobjectivesp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "thoughts" {
		if me.Kthoughtsp >= 0 {
			st := glob.Dats.ApThought[ me.Kthoughtsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "memories" {
		if me.Kmemoriesp >= 0 {
			st := glob.Dats.ApMemory[ me.Kmemoriesp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Snapshot %s,%s > session.unit:30, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpRestoration struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksource_sessionp int
	Kagent_idp int
	Ksource_memoryp int
	Ksessionp int
	Kagent_id2p int
	Krestored_asp int
}

func (me KpRestoration) TypeName() string {
    return me.Comp
}
func (me KpRestoration) GetLineNo() string {
	return me.LineNo
}

func loadRestoration(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpRestoration)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApRestoration)
	st.LineNo = lno
	st.Comp = "Restoration";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksource_sessionp = -1
	st.Kagent_idp = -1
	st.Ksource_memoryp = -1
	st.Ksessionp = -1
	st.Kagent_id2p = -1
	st.Krestored_asp = -1
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Restoration has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Restoration under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsRestoration = append(act.ApSession[ len( act.ApSession )-1 ].ItsRestoration, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["restoration_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Restoration_" + name	// session.unit:51, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApRestoration = append(act.ApRestoration, st)
	return 0
}

func (me KpRestoration) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "source_session" { // session.unit:52, go-struct-rio.act:621
		if (me.Ksource_sessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksource_sessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id" { // session.unit:53, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "source_memory" { // session.unit:54, go-struct-rio.act:621
		if (me.Ksource_memoryp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Ksource_memoryp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "session" { // session.unit:55, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id2" { // session.unit:56, go-struct-rio.act:621
		if (me.Kagent_id2p >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_id2p ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "restored_as" { // session.unit:57, go-struct-rio.act:621
		if (me.Krestored_asp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Krestored_asp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:50, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApRestoration[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Restoration > session.unit:50, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Restoration > session.unit:50, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpRestoration) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "source_session" {
		if me.Ksource_sessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksource_sessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "source_memory" {
		if me.Ksource_memoryp >= 0 {
			st := glob.Dats.ApMemory[ me.Ksource_memoryp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id2" {
		if me.Kagent_id2p >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_id2p ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "restored_as" {
		if me.Krestored_asp >= 0 {
			st := glob.Dats.ApMemory[ me.Krestored_asp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Restoration %s,%s > session.unit:50, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCrossRef struct {
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

func (me KpCrossRef) TypeName() string {
    return me.Comp
}
func (me KpCrossRef) GetLineNo() string {
	return me.LineNo
}

func loadCrossRef(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCrossRef)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCrossRef)
	st.LineNo = lno
	st.Comp = "CrossRef";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " CrossRef has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " CrossRef under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsCrossRef = append(act.ApSession[ len( act.ApSession )-1 ].ItsCrossRef, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["xref_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_CrossRef_" + name	// session.unit:65, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApCrossRef = append(act.ApCrossRef, st)
	return 0
}

func (me KpCrossRef) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:64, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCrossRef[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,CrossRef > session.unit:64, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,CrossRef > session.unit:64, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCrossRef) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for CrossRef %s,%s > session.unit:64, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCycle struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksessionp int
	Kagent_idp int
	Kobjectivep int
	Khypothesisp int
	Kagent_id2p int
	Kbaselinep int
	Kobjective2p int
	Kinterventionp int
	Kagent_id3p int
	Koutcomep int
	ItsCanonCycle [] *KpCanonCycle 
	Childs [] Kp
}

func (me KpCycle) TypeName() string {
    return me.Comp
}
func (me KpCycle) GetLineNo() string {
	return me.LineNo
}

func loadCycle(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCycle)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCycle)
	st.LineNo = lno
	st.Comp = "Cycle";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksessionp = -1
	st.Kagent_idp = -1
	st.Kobjectivep = -1
	st.Khypothesisp = -1
	st.Kagent_id2p = -1
	st.Kbaselinep = -1
	st.Kobjective2p = -1
	st.Kinterventionp = -1
	st.Kagent_id3p = -1
	st.Koutcomep = -1
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Cycle has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Cycle under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsCycle = append(act.ApSession[ len( act.ApSession )-1 ].ItsCycle, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["cycle_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Cycle_" + name	// session.unit:86, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApCycle = append(act.ApCycle, st)
	return 0
}

func (me KpCycle) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "session" { // session.unit:88, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id" { // session.unit:89, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "objective" { // session.unit:90, go-struct-rio.act:621
		if (me.Kobjectivep >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kobjectivep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "hypothesis" { // session.unit:91, go-struct-rio.act:621
		if (me.Khypothesisp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Khypothesisp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "baseline" { // session.unit:93, go-struct-rio.act:621
		if (me.Kbaselinep >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Kbaselinep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "intervention" { // session.unit:95, go-struct-rio.act:621
		if (me.Kinterventionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApThought[ me.Kinterventionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "outcome" { // session.unit:97, go-struct-rio.act:621
		if (me.Koutcomep >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Koutcomep ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:85, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCycle[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Cycle > session.unit:85, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Cycle > session.unit:85, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCycle) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "CanonCycle" { // session.unit:133, go-struct-rio.act:685
		for _, st := range me.ItsCanonCycle {
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
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "objective" {
		if me.Kobjectivep >= 0 {
			st := glob.Dats.ApObjective[ me.Kobjectivep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "hypothesis" {
		if me.Khypothesisp >= 0 {
			st := glob.Dats.ApThought[ me.Khypothesisp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "baseline" {
		if me.Kbaselinep >= 0 {
			st := glob.Dats.ApMemory[ me.Kbaselinep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "intervention" {
		if me.Kinterventionp >= 0 {
			st := glob.Dats.ApThought[ me.Kinterventionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "outcome" {
		if me.Koutcomep >= 0 {
			st := glob.Dats.ApMemory[ me.Koutcomep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Pattern_evidence") { // session.unit:155, go-struct-rio.act:595
		for _, st := range glob.Dats.ApPattern {
			if (st.Kevidencep == me.Me) {
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
	if (va[0] == "CanonVersion_evidence") { // canon.unit:103, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kevidencep == me.Me) {
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
	        fmt.Printf("?No its %s for Cycle %s,%s > session.unit:85, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpRecommendation struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksessionp int
	Kagent_idp int
	Koutcomep int
	Kagent_id2p int
	Kobjectivep int
	Kcanonp int
	Kimplp int
}

func (me KpRecommendation) TypeName() string {
    return me.Comp
}
func (me KpRecommendation) GetLineNo() string {
	return me.LineNo
}

func loadRecommendation(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpRecommendation)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApRecommendation)
	st.LineNo = lno
	st.Comp = "Recommendation";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksessionp = -1
	st.Kagent_idp = -1
	st.Koutcomep = -1
	st.Kagent_id2p = -1
	st.Kobjectivep = -1
	st.Kcanonp = -1
	st.Kimplp = -1
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Recommendation has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Recommendation under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsRecommendation = append(act.ApSession[ len( act.ApSession )-1 ].ItsRecommendation, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["rec_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Recommendation_" + name	// session.unit:112, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApRecommendation = append(act.ApRecommendation, st)
	return 0
}

func (me KpRecommendation) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "session" { // session.unit:113, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id" { // session.unit:114, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "outcome" { // session.unit:115, go-struct-rio.act:621
		if (me.Koutcomep >= 0 && len(va) > 1) {
			return( glob.Dats.ApMemory[ me.Koutcomep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "objective" { // session.unit:117, go-struct-rio.act:621
		if (me.Kobjectivep >= 0 && len(va) > 1) {
			return( glob.Dats.ApObjective[ me.Kobjectivep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "canon" { // session.unit:119, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "impl" { // session.unit:123, go-struct-rio.act:621
		if (me.Kimplp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:107, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApRecommendation[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Recommendation > session.unit:107, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Recommendation > session.unit:107, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpRecommendation) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "outcome" {
		if me.Koutcomep >= 0 {
			st := glob.Dats.ApMemory[ me.Koutcomep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "objective" {
		if me.Kobjectivep >= 0 {
			st := glob.Dats.ApObjective[ me.Kobjectivep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "impl" {
		if me.Kimplp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Recommendation %s,%s > session.unit:107, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCanonCycle struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kimplementationsp int
	Kbaseline_canonp int
	Knew_versionp int
}

func (me KpCanonCycle) TypeName() string {
    return me.Comp
}
func (me KpCanonCycle) GetLineNo() string {
	return me.LineNo
}

func loadCanonCycle(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCanonCycle)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCanonCycle)
	st.LineNo = lno
	st.Comp = "CanonCycle";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kimplementationsp = -1
	st.Kbaseline_canonp = -1
	st.Knew_versionp = -1
	st.Kparentp = len( act.ApCycle ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " CanonCycle has no Cycle parent\n") ;
		return 1
	}
	st.Parent = act.ApCycle[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " CanonCycle under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApCycle[ len( act.ApCycle )-1 ].Childs = append(act.ApCycle[ len( act.ApCycle )-1 ].Childs, st)
	act.ApCycle[ len( act.ApCycle )-1 ].ItsCanonCycle = append(act.ApCycle[ len( act.ApCycle )-1 ].ItsCanonCycle, st)	// session.unit:85, go-struct-rio.act:416
	act.ApCanonCycle = append(act.ApCanonCycle, st)
	return 0
}

func (me KpCanonCycle) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "implementations" { // session.unit:138, go-struct-rio.act:621
		if (me.Kimplementationsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplementationsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "baseline_canon" { // session.unit:139, go-struct-rio.act:621
		if (me.Kbaseline_canonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kbaseline_canonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "new_version" { // session.unit:140, go-struct-rio.act:621
		if (me.Knew_versionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanonVersion[ me.Knew_versionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:85, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCycle[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:133, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCanonCycle[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,CanonCycle > session.unit:133, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,CanonCycle > session.unit:133, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCanonCycle) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:85, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCycle[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "implementations" {
		if me.Kimplementationsp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplementationsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "baseline_canon" {
		if me.Kbaseline_canonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kbaseline_canonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "new_version" {
		if me.Knew_versionp >= 0 {
			st := glob.Dats.ApCanonVersion[ me.Knew_versionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for CanonCycle %s,%s > session.unit:133, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpPattern struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kdiscovered_inp int
	Kevidencep int
	ItsPatternSource [] *KpPatternSource 
	Childs [] Kp
}

func (me KpPattern) TypeName() string {
    return me.Comp
}
func (me KpPattern) GetLineNo() string {
	return me.LineNo
}

func loadPattern(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpPattern)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApPattern)
	st.LineNo = lno
	st.Comp = "Pattern";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kdiscovered_inp = -1
	st.Kevidencep = -1
	name,_ := st.Names["pattern_id"].(string)
	st.Names["_key"] = "pattern_id"
	act.index["Pattern_" + name] = st.Me;
	st.MyName = name
	act.ApPattern = append(act.ApPattern, st)
	return 0
}

func (me KpPattern) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "discovered_in" { // session.unit:154, go-struct-rio.act:621
		if (me.Kdiscovered_inp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kdiscovered_inp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "evidence" { // session.unit:155, go-struct-rio.act:621
		if (me.Kevidencep >= 0 && len(va) > 1) {
			return( glob.Dats.ApCycle[ me.Kevidencep ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Evolution_motivation" && len(va) > 1) { // session.unit:178, go-struct-rio.act:706
		for _, st := range glob.Dats.ApEvolution {
			if (st.Kmotivationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenPatternRef_pattern" && len(va) > 1) { // gen-high.unit:127, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenPatternRef {
			if (st.Kpatternp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // session.unit:150, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApPattern[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Pattern > session.unit:150, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Pattern > session.unit:150, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpPattern) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "PatternSource" { // session.unit:166, go-struct-rio.act:685
		for _, st := range me.ItsPatternSource {
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
	if va[0] == "discovered_in" {
		if me.Kdiscovered_inp >= 0 {
			st := glob.Dats.ApSession[ me.Kdiscovered_inp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "evidence" {
		if me.Kevidencep >= 0 {
			st := glob.Dats.ApCycle[ me.Kevidencep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Evolution_motivation") { // session.unit:178, go-struct-rio.act:595
		for _, st := range glob.Dats.ApEvolution {
			if (st.Kmotivationp == me.Me) {
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
	if (va[0] == "GenPatternRef_pattern") { // gen-high.unit:127, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenPatternRef {
			if (st.Kpatternp == me.Me) {
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
	        fmt.Printf("?No its %s for Pattern %s,%s > session.unit:150, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpPatternSource struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kimplp int
	Kcanonp int
}

func (me KpPatternSource) TypeName() string {
    return me.Comp
}
func (me KpPatternSource) GetLineNo() string {
	return me.LineNo
}

func loadPatternSource(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpPatternSource)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApPatternSource)
	st.LineNo = lno
	st.Comp = "PatternSource";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kimplp = -1
	st.Kcanonp = -1
	st.Kparentp = len( act.ApPattern ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " PatternSource has no Pattern parent\n") ;
		return 1
	}
	st.Parent = act.ApPattern[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " PatternSource under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApPattern[ len( act.ApPattern )-1 ].Childs = append(act.ApPattern[ len( act.ApPattern )-1 ].Childs, st)
	act.ApPattern[ len( act.ApPattern )-1 ].ItsPatternSource = append(act.ApPattern[ len( act.ApPattern )-1 ].ItsPatternSource, st)	// session.unit:150, go-struct-rio.act:416
	act.ApPatternSource = append(act.ApPatternSource, st)
	return 0
}

func (me KpPatternSource) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "impl" { // session.unit:167, go-struct-rio.act:621
		if (me.Kimplp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "canon" { // session.unit:168, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:150, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApPattern[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:166, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApPatternSource[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,PatternSource > session.unit:166, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,PatternSource > session.unit:166, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpPatternSource) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:150, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApPattern[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "impl" {
		if me.Kimplp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for PatternSource %s,%s > session.unit:166, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpEvolution struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kmotivationp int
	Ksessionsp int
}

func (me KpEvolution) TypeName() string {
    return me.Comp
}
func (me KpEvolution) GetLineNo() string {
	return me.LineNo
}

func loadEvolution(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpEvolution)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApEvolution)
	st.LineNo = lno
	st.Comp = "Evolution";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kmotivationp = -1
	st.Ksessionsp = -1
	name,_ := st.Names["evolution_id"].(string)
	st.Names["_key"] = "evolution_id"
	act.index["Evolution_" + name] = st.Me;
	st.MyName = name
	act.ApEvolution = append(act.ApEvolution, st)
	return 0
}

func (me KpEvolution) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "motivation" { // session.unit:178, go-struct-rio.act:621
		if (me.Kmotivationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApPattern[ me.Kmotivationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "sessions" { // session.unit:179, go-struct-rio.act:621
		if (me.Ksessionsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // session.unit:174, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApEvolution[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Evolution > session.unit:174, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Evolution > session.unit:174, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpEvolution) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "motivation" {
		if me.Kmotivationp >= 0 {
			st := glob.Dats.ApPattern[ me.Kmotivationp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "sessions" {
		if me.Ksessionsp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Evolution %s,%s > session.unit:174, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpMetricS struct {
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

func (me KpMetricS) TypeName() string {
    return me.Comp
}
func (me KpMetricS) GetLineNo() string {
	return me.LineNo
}

func loadMetricS(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpMetricS)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApMetricS)
	st.LineNo = lno
	st.Comp = "MetricS";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " MetricS has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " MetricS under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsMetricS = append(act.ApSession[ len( act.ApSession )-1 ].ItsMetricS, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["metric_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_MetricS_" + name	// session.unit:193, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApMetricS = append(act.ApMetricS, st)
	return 0
}

func (me KpMetricS) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:192, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMetricS[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,MetricS > session.unit:192, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,MetricS > session.unit:192, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpMetricS) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for MetricS %s,%s > session.unit:192, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpArtifact struct {
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

func (me KpArtifact) TypeName() string {
    return me.Comp
}
func (me KpArtifact) GetLineNo() string {
	return me.LineNo
}

func loadArtifact(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpArtifact)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApArtifact)
	st.LineNo = lno
	st.Comp = "Artifact";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Artifact has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Artifact under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsArtifact = append(act.ApSession[ len( act.ApSession )-1 ].ItsArtifact, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["artifact_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Artifact_" + name	// session.unit:205, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApArtifact = append(act.ApArtifact, st)
	return 0
}

func (me KpArtifact) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:204, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Artifact > session.unit:204, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Artifact > session.unit:204, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpArtifact) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Artifact %s,%s > session.unit:204, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenExecution struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kprogramp int
	Ksessionp int
	Kagent_idp int
	ItsGenValidation [] *KpGenValidation 
	ItsGenLearning [] *KpGenLearning 
	ItsGenDebug [] *KpGenDebug 
	Childs [] Kp
}

func (me KpGenExecution) TypeName() string {
    return me.Comp
}
func (me KpGenExecution) GetLineNo() string {
	return me.LineNo
}

func loadGenExecution(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenExecution)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenExecution)
	st.LineNo = lno
	st.Comp = "GenExecution";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kprogramp = -1
	st.Ksessionp = -1
	st.Kagent_idp = -1
	st.Kparentp = len( act.ApSession ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenExecution has no Session parent\n") ;
		return 1
	}
	st.Parent = act.ApSession[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenExecution under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSession[ len( act.ApSession )-1 ].Childs = append(act.ApSession[ len( act.ApSession )-1 ].Childs, st)
	act.ApSession[ len( act.ApSession )-1 ].ItsGenExecution = append(act.ApSession[ len( act.ApSession )-1 ].ItsGenExecution, st)	// session.unit:12, go-struct-rio.act:416
	name,_ := st.Names["execution_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenExecution_" + name	// session.unit:231, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenExecution = append(act.ApGenExecution, st)
	return 0
}

func (me KpGenExecution) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "program" { // session.unit:232, go-struct-rio.act:621
		if (me.Kprogramp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kprogramp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "session" { // session.unit:233, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "agent_id" { // session.unit:234, go-struct-rio.act:621
		if (me.Kagent_idp >= 0 && len(va) > 1) {
			return( glob.Dats.ApAgent[ me.Kagent_idp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:12, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:230, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenExecution[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenExecution > session.unit:230, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenExecution > session.unit:230, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenExecution) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "GenValidation" { // session.unit:259, go-struct-rio.act:685
		for _, st := range me.ItsGenValidation {
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
	if va[0] == "GenLearning" { // session.unit:287, go-struct-rio.act:685
		for _, st := range me.ItsGenLearning {
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
	if va[0] == "GenDebug" { // session.unit:306, go-struct-rio.act:685
		for _, st := range me.ItsGenDebug {
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
	if va[0] == "parent" { // session.unit:12, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSession[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "program" {
		if me.Kprogramp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kprogramp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "agent_id" {
		if me.Kagent_idp >= 0 {
			st := glob.Dats.ApAgent[ me.Kagent_idp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "GenLearning_baseline") { // session.unit:291, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenLearning {
			if (st.Kbaselinep == me.Me) {
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
	        fmt.Printf("?No its %s for GenExecution %s,%s > session.unit:230, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenValidation struct {
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

func (me KpGenValidation) TypeName() string {
    return me.Comp
}
func (me KpGenValidation) GetLineNo() string {
	return me.LineNo
}

func loadGenValidation(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenValidation)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenValidation)
	st.LineNo = lno
	st.Comp = "GenValidation";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenExecution ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenValidation has no GenExecution parent\n") ;
		return 1
	}
	st.Parent = act.ApGenExecution[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenValidation under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs, st)
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenValidation = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenValidation, st)	// session.unit:230, go-struct-rio.act:416
	name,_ := st.Names["validation_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenValidation_" + name	// session.unit:260, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenValidation = append(act.ApGenValidation, st)
	return 0
}

func (me KpGenValidation) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // session.unit:230, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenExecution[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:259, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenValidation[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenValidation > session.unit:259, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenValidation > session.unit:259, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenValidation) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:230, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenExecution[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenValidation %s,%s > session.unit:259, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenLearning struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksessionp int
	Kbaselinep int
	Kapplied_to_programp int
}

func (me KpGenLearning) TypeName() string {
    return me.Comp
}
func (me KpGenLearning) GetLineNo() string {
	return me.LineNo
}

func loadGenLearning(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenLearning)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenLearning)
	st.LineNo = lno
	st.Comp = "GenLearning";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksessionp = -1
	st.Kbaselinep = -1
	st.Kapplied_to_programp = -1
	st.Kparentp = len( act.ApGenExecution ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenLearning has no GenExecution parent\n") ;
		return 1
	}
	st.Parent = act.ApGenExecution[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenLearning under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs, st)
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenLearning = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenLearning, st)	// session.unit:230, go-struct-rio.act:416
	name,_ := st.Names["learning_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenLearning_" + name	// session.unit:288, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenLearning = append(act.ApGenLearning, st)
	return 0
}

func (me KpGenLearning) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "session" { // session.unit:290, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "baseline" { // session.unit:291, go-struct-rio.act:621
		if (me.Kbaselinep >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenExecution[ me.Kbaselinep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "applied_to_program" { // session.unit:295, go-struct-rio.act:621
		if (me.Kapplied_to_programp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kapplied_to_programp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // session.unit:230, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenExecution[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:287, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenLearning[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenLearning > session.unit:287, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenLearning > session.unit:287, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenLearning) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:230, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenExecution[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "baseline" {
		if me.Kbaselinep >= 0 {
			st := glob.Dats.ApGenExecution[ me.Kbaselinep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "applied_to_program" {
		if me.Kapplied_to_programp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kapplied_to_programp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenLearning %s,%s > session.unit:287, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenDebug struct {
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

func (me KpGenDebug) TypeName() string {
    return me.Comp
}
func (me KpGenDebug) GetLineNo() string {
	return me.LineNo
}

func loadGenDebug(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenDebug)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenDebug)
	st.LineNo = lno
	st.Comp = "GenDebug";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenExecution ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenDebug has no GenExecution parent\n") ;
		return 1
	}
	st.Parent = act.ApGenExecution[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenDebug under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].Childs, st)
	act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenDebug = append(act.ApGenExecution[ len( act.ApGenExecution )-1 ].ItsGenDebug, st)	// session.unit:230, go-struct-rio.act:416
	name,_ := st.Names["debug_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenDebug_" + name	// session.unit:307, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenDebug = append(act.ApGenDebug, st)
	return 0
}

func (me KpGenDebug) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // session.unit:230, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenExecution[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // session.unit:306, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenDebug[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenDebug > session.unit:306, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenDebug > session.unit:306, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenDebug) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // session.unit:230, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenExecution[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenDebug %s,%s > session.unit:306, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
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
	ItsCanonMeta [] *KpCanonMeta 
	ItsCanonVersion [] *KpCanonVersion 
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

func (me KpCanon) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "MemorySource_canon" && len(va) > 1) { // cog.unit:59, go-struct-rio.act:706
		for _, st := range glob.Dats.ApMemorySource {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "ThoughtSource_canon" && len(va) > 1) { // cog.unit:127, go-struct-rio.act:706
		for _, st := range glob.Dats.ApThoughtSource {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Recommendation_canon" && len(va) > 1) { // session.unit:119, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "CanonCycle_baseline_canon" && len(va) > 1) { // session.unit:139, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCanonCycle {
			if (st.Kbaseline_canonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "PatternSource_canon" && len(va) > 1) { // session.unit:168, go-struct-rio.act:706
		for _, st := range glob.Dats.ApPatternSource {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Link_concept" && len(va) > 1) { // canon.unit:30, go-struct-rio.act:706
		for _, st := range glob.Dats.ApLink {
			if (st.Kconceptp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Link_relation" && len(va) > 1) { // canon.unit:31, go-struct-rio.act:706
		for _, st := range glob.Dats.ApLink {
			if (st.Krelationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Implementation_canon" && len(va) > 1) { // canon.unit:68, go-struct-rio.act:706
		for _, st := range glob.Dats.ApImplementation {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "CanonVersion_canon" && len(va) > 1) { // canon.unit:96, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kcanonp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Library_canons" && len(va) > 1) { // canon.unit:127, go-struct-rio.act:706
		for _, st := range glob.Dats.ApLibrary {
			if (st.Kcanonsp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // canon.unit:14, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Canon > canon.unit:14, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Canon > canon.unit:14, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCanon) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Link" { // canon.unit:26, go-struct-rio.act:685
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
	if va[0] == "Section" { // canon.unit:33, go-struct-rio.act:685
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
	if va[0] == "CanonMeta" { // canon.unit:47, go-struct-rio.act:685
		for _, st := range me.ItsCanonMeta {
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
	if va[0] == "CanonVersion" { // canon.unit:90, go-struct-rio.act:685
		for _, st := range me.ItsCanonVersion {
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
	if (va[0] == "MemorySource_canon") { // cog.unit:59, go-struct-rio.act:595
		for _, st := range glob.Dats.ApMemorySource {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "ThoughtSource_canon") { // cog.unit:127, go-struct-rio.act:595
		for _, st := range glob.Dats.ApThoughtSource {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "Recommendation_canon") { // session.unit:119, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "CanonCycle_baseline_canon") { // session.unit:139, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonCycle {
			if (st.Kbaseline_canonp == me.Me) {
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
	if (va[0] == "PatternSource_canon") { // session.unit:168, go-struct-rio.act:595
		for _, st := range glob.Dats.ApPatternSource {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "Link_concept") { // canon.unit:30, go-struct-rio.act:595
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
	if (va[0] == "Link_relation") { // canon.unit:31, go-struct-rio.act:595
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
	if (va[0] == "Implementation_canon") { // canon.unit:68, go-struct-rio.act:595
		for _, st := range glob.Dats.ApImplementation {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "CanonVersion_canon") { // canon.unit:96, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kcanonp == me.Me) {
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
	if (va[0] == "Library_canons") { // canon.unit:127, go-struct-rio.act:595
		for _, st := range glob.Dats.ApLibrary {
			if (st.Kcanonsp == me.Me) {
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
	        fmt.Printf("?No its %s for Canon %s,%s > canon.unit:14, go-struct-rio.act:222?", va[0], lno, me.LineNo)
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
	act.ApCanon[ len( act.ApCanon )-1 ].ItsLink = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsLink, st)	// canon.unit:14, go-struct-rio.act:416
	act.ApLink = append(act.ApLink, st)
	return 0
}

func (me KpLink) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "concept" { // canon.unit:30, go-struct-rio.act:621
		if (me.Kconceptp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kconceptp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "relation" { // canon.unit:31, go-struct-rio.act:621
		if (me.Krelationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Krelationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // canon.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // canon.unit:26, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApLink[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Link > canon.unit:26, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Link > canon.unit:26, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpLink) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // canon.unit:14, go-struct-rio.act:570
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
	        fmt.Printf("?No its %s for Link %s,%s > canon.unit:26, go-struct-rio.act:222?", va[0], lno, me.LineNo)
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
	act.ApCanon[ len( act.ApCanon )-1 ].ItsSection = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsSection, st)	// canon.unit:14, go-struct-rio.act:416
	name,_ := st.Names["name"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Section_" + name	// canon.unit:37, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApSection = append(act.ApSection, st)
	return 0
}

func (me KpSection) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // canon.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // canon.unit:33, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSection[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Section > canon.unit:33, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Section > canon.unit:33, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpSection) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // canon.unit:14, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCanon[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
		if va[0] == "level" && len(va) > 1 && me.Kparentp >= 0 { // canon.unit:38, go-struct-rio.act:233
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
	        fmt.Printf("?No its %s for Section %s,%s > canon.unit:33, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCanonMeta struct {
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

func (me KpCanonMeta) TypeName() string {
    return me.Comp
}
func (me KpCanonMeta) GetLineNo() string {
	return me.LineNo
}

func loadCanonMeta(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCanonMeta)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCanonMeta)
	st.LineNo = lno
	st.Comp = "CanonMeta";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApCanon ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " CanonMeta has no Canon parent\n") ;
		return 1
	}
	st.Parent = act.ApCanon[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " CanonMeta under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApCanon[ len( act.ApCanon )-1 ].Childs = append(act.ApCanon[ len( act.ApCanon )-1 ].Childs, st)
	act.ApCanon[ len( act.ApCanon )-1 ].ItsCanonMeta = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsCanonMeta, st)	// canon.unit:14, go-struct-rio.act:416
	act.ApCanonMeta = append(act.ApCanonMeta, st)
	return 0
}

func (me KpCanonMeta) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // canon.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // canon.unit:47, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCanonMeta[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,CanonMeta > canon.unit:47, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,CanonMeta > canon.unit:47, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCanonMeta) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // canon.unit:14, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCanon[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for CanonMeta %s,%s > canon.unit:47, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpImplementation struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kcanonp int
	Ksessionp int
}

func (me KpImplementation) TypeName() string {
    return me.Comp
}
func (me KpImplementation) GetLineNo() string {
	return me.LineNo
}

func loadImplementation(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpImplementation)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApImplementation)
	st.LineNo = lno
	st.Comp = "Implementation";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonp = -1
	st.Ksessionp = -1
	name,_ := st.Names["impl_id"].(string)
	st.Names["_key"] = "impl_id"
	act.index["Implementation_" + name] = st.Me;
	st.MyName = name
	act.ApImplementation = append(act.ApImplementation, st)
	return 0
}

func (me KpImplementation) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "canon" { // canon.unit:68, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "session" { // canon.unit:71, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "MemorySource_impl" && len(va) > 1) { // cog.unit:58, go-struct-rio.act:706
		for _, st := range glob.Dats.ApMemorySource {
			if (st.Kimplp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "ThoughtSource_impl" && len(va) > 1) { // cog.unit:126, go-struct-rio.act:706
		for _, st := range glob.Dats.ApThoughtSource {
			if (st.Kimplp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Recommendation_impl" && len(va) > 1) { // session.unit:123, go-struct-rio.act:706
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kimplp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "CanonCycle_implementations" && len(va) > 1) { // session.unit:138, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCanonCycle {
			if (st.Kimplementationsp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "PatternSource_impl" && len(va) > 1) { // session.unit:167, go-struct-rio.act:706
		for _, st := range glob.Dats.ApPatternSource {
			if (st.Kimplp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "CanonVersion_implementations" && len(va) > 1) { // canon.unit:101, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kimplementationsp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // canon.unit:62, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Implementation > canon.unit:62, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Implementation > canon.unit:62, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpImplementation) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "MemorySource_impl") { // cog.unit:58, go-struct-rio.act:595
		for _, st := range glob.Dats.ApMemorySource {
			if (st.Kimplp == me.Me) {
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
	if (va[0] == "ThoughtSource_impl") { // cog.unit:126, go-struct-rio.act:595
		for _, st := range glob.Dats.ApThoughtSource {
			if (st.Kimplp == me.Me) {
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
	if (va[0] == "Recommendation_impl") { // session.unit:123, go-struct-rio.act:595
		for _, st := range glob.Dats.ApRecommendation {
			if (st.Kimplp == me.Me) {
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
	if (va[0] == "CanonCycle_implementations") { // session.unit:138, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonCycle {
			if (st.Kimplementationsp == me.Me) {
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
	if (va[0] == "PatternSource_impl") { // session.unit:167, go-struct-rio.act:595
		for _, st := range glob.Dats.ApPatternSource {
			if (st.Kimplp == me.Me) {
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
	if (va[0] == "CanonVersion_implementations") { // canon.unit:101, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kimplementationsp == me.Me) {
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
	        fmt.Printf("?No its %s for Implementation %s,%s > canon.unit:62, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCanonVersion struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kcanonp int
	Kpreviousp int
	Kimplementationsp int
	Ksessionp int
	Kevidencep int
}

func (me KpCanonVersion) TypeName() string {
    return me.Comp
}
func (me KpCanonVersion) GetLineNo() string {
	return me.LineNo
}

func loadCanonVersion(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCanonVersion)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCanonVersion)
	st.LineNo = lno
	st.Comp = "CanonVersion";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonp = -1
	st.Kpreviousp = -1
	st.Kimplementationsp = -1
	st.Ksessionp = -1
	st.Kevidencep = -1
	st.Kparentp = len( act.ApCanon ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " CanonVersion has no Canon parent\n") ;
		return 1
	}
	st.Parent = act.ApCanon[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " CanonVersion under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApCanon[ len( act.ApCanon )-1 ].Childs = append(act.ApCanon[ len( act.ApCanon )-1 ].Childs, st)
	act.ApCanon[ len( act.ApCanon )-1 ].ItsCanonVersion = append(act.ApCanon[ len( act.ApCanon )-1 ].ItsCanonVersion, st)	// canon.unit:14, go-struct-rio.act:416
	name,_ := st.Names["version"].(string)
	s := strconv.Itoa(st.Kparentp) + "_CanonVersion_" + name	// canon.unit:95, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApCanonVersion = append(act.ApCanonVersion, st)
	return 0
}

func (me KpCanonVersion) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "canon" { // canon.unit:96, go-struct-rio.act:621
		if (me.Kcanonp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // canon.unit:97, go-struct-rio.act:621
		if (me.Kpreviousp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanonVersion[ me.Kpreviousp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "implementations" { // canon.unit:101, go-struct-rio.act:621
		if (me.Kimplementationsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApImplementation[ me.Kimplementationsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "session" { // canon.unit:102, go-struct-rio.act:621
		if (me.Ksessionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSession[ me.Ksessionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "evidence" { // canon.unit:103, go-struct-rio.act:621
		if (me.Kevidencep >= 0 && len(va) > 1) {
			return( glob.Dats.ApCycle[ me.Kevidencep ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // canon.unit:14, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // canon.unit:90, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCanonVersion[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,CanonVersion > canon.unit:90, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,CanonVersion > canon.unit:90, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCanonVersion) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // canon.unit:14, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApCanon[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "canon" {
		if me.Kcanonp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "previous" {
		if me.Kpreviousp >= 0 {
			st := glob.Dats.ApCanonVersion[ me.Kpreviousp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "implementations" {
		if me.Kimplementationsp >= 0 {
			st := glob.Dats.ApImplementation[ me.Kimplementationsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "session" {
		if me.Ksessionp >= 0 {
			st := glob.Dats.ApSession[ me.Ksessionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "evidence" {
		if me.Kevidencep >= 0 {
			st := glob.Dats.ApCycle[ me.Kevidencep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "CanonCycle_new_version") { // session.unit:140, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonCycle {
			if (st.Knew_versionp == me.Me) {
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
	if (va[0] == "CanonVersion_previous") { // canon.unit:97, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCanonVersion {
			if (st.Kpreviousp == me.Me) {
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
	        fmt.Printf("?No its %s for CanonVersion %s,%s > canon.unit:90, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpLibrary struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kcanonsp int
}

func (me KpLibrary) TypeName() string {
    return me.Comp
}
func (me KpLibrary) GetLineNo() string {
	return me.LineNo
}

func loadLibrary(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpLibrary)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApLibrary)
	st.LineNo = lno
	st.Comp = "Library";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kcanonsp = -1
	name,_ := st.Names["lib_id"].(string)
	st.Names["_key"] = "lib_id"
	act.index["Library_" + name] = st.Me;
	st.MyName = name
	act.ApLibrary = append(act.ApLibrary, st)
	return 0
}

func (me KpLibrary) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "canons" { // canon.unit:127, go-struct-rio.act:621
		if (me.Kcanonsp >= 0 && len(va) > 1) {
			return( glob.Dats.ApCanon[ me.Kcanonsp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // canon.unit:118, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApLibrary[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Library > canon.unit:118, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Library > canon.unit:118, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpLibrary) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "canons" {
		if me.Kcanonsp >= 0 {
			st := glob.Dats.ApCanon[ me.Kcanonsp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Library %s,%s > canon.unit:118, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenProgram struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ksupersedesp int
	ItsGenBootstrap [] *KpGenBootstrap 
	ItsGenInput [] *KpGenInput 
	ItsGenStrategy [] *KpGenStrategy 
	ItsGenPatternRef [] *KpGenPatternRef 
	ItsGenOutputSpec [] *KpGenOutputSpec 
	ItsGenMetric [] *KpGenMetric 
	ItsGenEvolution [] *KpGenEvolution 
	Childs [] Kp
}

func (me KpGenProgram) TypeName() string {
    return me.Comp
}
func (me KpGenProgram) GetLineNo() string {
	return me.LineNo
}

func loadGenProgram(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenProgram)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenProgram)
	st.LineNo = lno
	st.Comp = "GenProgram";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksupersedesp = -1
	name,_ := st.Names["program_id"].(string)
	st.Names["_key"] = "program_id"
	act.index["GenProgram_" + name] = st.Me;
	st.MyName = name
	act.ApGenProgram = append(act.ApGenProgram, st)
	return 0
}

func (me KpGenProgram) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "supersedes" { // gen-high.unit:21, go-struct-rio.act:621
		if (me.Ksupersedesp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Ksupersedesp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "GenExecution_program" && len(va) > 1) { // session.unit:232, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Kprogramp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenLearning_applied_to_program" && len(va) > 1) { // session.unit:295, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenLearning {
			if (st.Kapplied_to_programp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenProgram_supersedes" && len(va) > 1) { // gen-high.unit:21, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenProgram {
			if (st.Ksupersedesp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenBootstrap_superseded_by" && len(va) > 1) { // gen-high.unit:46, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenBootstrap {
			if (st.Ksuperseded_byp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenEvolution_previous_version" && len(va) > 1) { // gen-high.unit:190, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenEvolution {
			if (st.Kprevious_versionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "GenStage_program" && len(va) > 1) { // gen-high.unit:221, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenStage {
			if (st.Kprogramp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // gen-high.unit:13, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenProgram > gen-high.unit:13, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenProgram > gen-high.unit:13, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenProgram) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "GenBootstrap" { // gen-high.unit:40, go-struct-rio.act:685
		for _, st := range me.ItsGenBootstrap {
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
	if va[0] == "GenInput" { // gen-high.unit:60, go-struct-rio.act:685
		for _, st := range me.ItsGenInput {
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
	if va[0] == "GenStrategy" { // gen-high.unit:80, go-struct-rio.act:685
		for _, st := range me.ItsGenStrategy {
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
	if va[0] == "GenPatternRef" { // gen-high.unit:125, go-struct-rio.act:685
		for _, st := range me.ItsGenPatternRef {
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
	if va[0] == "GenOutputSpec" { // gen-high.unit:147, go-struct-rio.act:685
		for _, st := range me.ItsGenOutputSpec {
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
	if va[0] == "GenMetric" { // gen-high.unit:167, go-struct-rio.act:685
		for _, st := range me.ItsGenMetric {
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
	if va[0] == "GenEvolution" { // gen-high.unit:188, go-struct-rio.act:685
		for _, st := range me.ItsGenEvolution {
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
	if va[0] == "supersedes" {
		if me.Ksupersedesp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Ksupersedesp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "GenExecution_program") { // session.unit:232, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenExecution {
			if (st.Kprogramp == me.Me) {
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
	if (va[0] == "GenLearning_applied_to_program") { // session.unit:295, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenLearning {
			if (st.Kapplied_to_programp == me.Me) {
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
	if (va[0] == "GenProgram_supersedes") { // gen-high.unit:21, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenProgram {
			if (st.Ksupersedesp == me.Me) {
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
	if (va[0] == "GenBootstrap_superseded_by") { // gen-high.unit:46, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenBootstrap {
			if (st.Ksuperseded_byp == me.Me) {
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
	if (va[0] == "GenEvolution_previous_version") { // gen-high.unit:190, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenEvolution {
			if (st.Kprevious_versionp == me.Me) {
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
	if (va[0] == "GenStage_program") { // gen-high.unit:221, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenStage {
			if (st.Kprogramp == me.Me) {
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
	        fmt.Printf("?No its %s for GenProgram %s,%s > gen-high.unit:13, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenBootstrap struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Ksuperseded_byp int
}

func (me KpGenBootstrap) TypeName() string {
    return me.Comp
}
func (me KpGenBootstrap) GetLineNo() string {
	return me.LineNo
}

func loadGenBootstrap(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenBootstrap)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenBootstrap)
	st.LineNo = lno
	st.Comp = "GenBootstrap";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksuperseded_byp = -1
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenBootstrap has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenBootstrap under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenBootstrap = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenBootstrap, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["bootstrap_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenBootstrap_" + name	// gen-high.unit:41, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenBootstrap = append(act.ApGenBootstrap, st)
	return 0
}

func (me KpGenBootstrap) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "superseded_by" { // gen-high.unit:46, go-struct-rio.act:621
		if (me.Ksuperseded_byp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Ksuperseded_byp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:40, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenBootstrap[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenBootstrap > gen-high.unit:40, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenBootstrap > gen-high.unit:40, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenBootstrap) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "superseded_by" {
		if me.Ksuperseded_byp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Ksuperseded_byp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenBootstrap %s,%s > gen-high.unit:40, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenInput struct {
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

func (me KpGenInput) TypeName() string {
    return me.Comp
}
func (me KpGenInput) GetLineNo() string {
	return me.LineNo
}

func loadGenInput(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenInput)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenInput)
	st.LineNo = lno
	st.Comp = "GenInput";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenInput has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenInput under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenInput = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenInput, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["input_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenInput_" + name	// gen-high.unit:61, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenInput = append(act.ApGenInput, st)
	return 0
}

func (me KpGenInput) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:60, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenInput[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenInput > gen-high.unit:60, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenInput > gen-high.unit:60, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenInput) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenInput %s,%s > gen-high.unit:60, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenStrategy struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	ItsGenSearchSpace [] *KpGenSearchSpace 
	Childs [] Kp
}

func (me KpGenStrategy) TypeName() string {
    return me.Comp
}
func (me KpGenStrategy) GetLineNo() string {
	return me.LineNo
}

func loadGenStrategy(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenStrategy)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenStrategy)
	st.LineNo = lno
	st.Comp = "GenStrategy";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenStrategy has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenStrategy under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenStrategy = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenStrategy, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["strategy_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenStrategy_" + name	// gen-high.unit:81, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenStrategy = append(act.ApGenStrategy, st)
	return 0
}

func (me KpGenStrategy) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:80, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenStrategy[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenStrategy > gen-high.unit:80, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenStrategy > gen-high.unit:80, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenStrategy) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "GenSearchSpace" { // gen-high.unit:105, go-struct-rio.act:685
		for _, st := range me.ItsGenSearchSpace {
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
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenStrategy %s,%s > gen-high.unit:80, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenSearchSpace struct {
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

func (me KpGenSearchSpace) TypeName() string {
    return me.Comp
}
func (me KpGenSearchSpace) GetLineNo() string {
	return me.LineNo
}

func loadGenSearchSpace(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenSearchSpace)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenSearchSpace)
	st.LineNo = lno
	st.Comp = "GenSearchSpace";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenStrategy ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenSearchSpace has no GenStrategy parent\n") ;
		return 1
	}
	st.Parent = act.ApGenStrategy[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenSearchSpace under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenStrategy[ len( act.ApGenStrategy )-1 ].Childs = append(act.ApGenStrategy[ len( act.ApGenStrategy )-1 ].Childs, st)
	act.ApGenStrategy[ len( act.ApGenStrategy )-1 ].ItsGenSearchSpace = append(act.ApGenStrategy[ len( act.ApGenStrategy )-1 ].ItsGenSearchSpace, st)	// gen-high.unit:80, go-struct-rio.act:416
	name,_ := st.Names["space_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenSearchSpace_" + name	// gen-high.unit:106, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenSearchSpace = append(act.ApGenSearchSpace, st)
	return 0
}

func (me KpGenSearchSpace) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // gen-high.unit:80, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenStrategy[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:105, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenSearchSpace[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenSearchSpace > gen-high.unit:105, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenSearchSpace > gen-high.unit:105, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenSearchSpace) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:80, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenStrategy[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenSearchSpace %s,%s > gen-high.unit:105, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenPatternRef struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kpatternp int
}

func (me KpGenPatternRef) TypeName() string {
    return me.Comp
}
func (me KpGenPatternRef) GetLineNo() string {
	return me.LineNo
}

func loadGenPatternRef(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenPatternRef)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenPatternRef)
	st.LineNo = lno
	st.Comp = "GenPatternRef";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kpatternp = -1
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenPatternRef has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenPatternRef under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenPatternRef = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenPatternRef, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["ref_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenPatternRef_" + name	// gen-high.unit:126, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenPatternRef = append(act.ApGenPatternRef, st)
	return 0
}

func (me KpGenPatternRef) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "pattern" { // gen-high.unit:127, go-struct-rio.act:621
		if (me.Kpatternp >= 0 && len(va) > 1) {
			return( glob.Dats.ApPattern[ me.Kpatternp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:125, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenPatternRef[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenPatternRef > gen-high.unit:125, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenPatternRef > gen-high.unit:125, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenPatternRef) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "pattern" {
		if me.Kpatternp >= 0 {
			st := glob.Dats.ApPattern[ me.Kpatternp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenPatternRef %s,%s > gen-high.unit:125, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenOutputSpec struct {
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

func (me KpGenOutputSpec) TypeName() string {
    return me.Comp
}
func (me KpGenOutputSpec) GetLineNo() string {
	return me.LineNo
}

func loadGenOutputSpec(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenOutputSpec)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenOutputSpec)
	st.LineNo = lno
	st.Comp = "GenOutputSpec";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenOutputSpec has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenOutputSpec under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenOutputSpec = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenOutputSpec, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["spec_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenOutputSpec_" + name	// gen-high.unit:148, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenOutputSpec = append(act.ApGenOutputSpec, st)
	return 0
}

func (me KpGenOutputSpec) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:147, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenOutputSpec[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenOutputSpec > gen-high.unit:147, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenOutputSpec > gen-high.unit:147, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenOutputSpec) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenOutputSpec %s,%s > gen-high.unit:147, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenMetric struct {
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

func (me KpGenMetric) TypeName() string {
    return me.Comp
}
func (me KpGenMetric) GetLineNo() string {
	return me.LineNo
}

func loadGenMetric(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenMetric)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenMetric)
	st.LineNo = lno
	st.Comp = "GenMetric";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenMetric has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenMetric under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenMetric = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenMetric, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["metric_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenMetric_" + name	// gen-high.unit:168, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenMetric = append(act.ApGenMetric, st)
	return 0
}

func (me KpGenMetric) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:167, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenMetric[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenMetric > gen-high.unit:167, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenMetric > gen-high.unit:167, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenMetric) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenMetric %s,%s > gen-high.unit:167, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenEvolution struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kprevious_versionp int
}

func (me KpGenEvolution) TypeName() string {
    return me.Comp
}
func (me KpGenEvolution) GetLineNo() string {
	return me.LineNo
}

func loadGenEvolution(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenEvolution)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenEvolution)
	st.LineNo = lno
	st.Comp = "GenEvolution";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kprevious_versionp = -1
	st.Kparentp = len( act.ApGenProgram ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenEvolution has no GenProgram parent\n") ;
		return 1
	}
	st.Parent = act.ApGenProgram[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenEvolution under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].Childs, st)
	act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenEvolution = append(act.ApGenProgram[ len( act.ApGenProgram )-1 ].ItsGenEvolution, st)	// gen-high.unit:13, go-struct-rio.act:416
	name,_ := st.Names["evolution_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenEvolution_" + name	// gen-high.unit:189, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenEvolution = append(act.ApGenEvolution, st)
	return 0
}

func (me KpGenEvolution) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "previous_version" { // gen-high.unit:190, go-struct-rio.act:621
		if (me.Kprevious_versionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kprevious_versionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // gen-high.unit:13, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // gen-high.unit:188, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenEvolution[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenEvolution > gen-high.unit:188, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenEvolution > gen-high.unit:188, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenEvolution) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:13, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "previous_version" {
		if me.Kprevious_versionp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kprevious_versionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for GenEvolution %s,%s > gen-high.unit:188, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenPipeline struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	ItsGenStage [] *KpGenStage 
	Childs [] Kp
}

func (me KpGenPipeline) TypeName() string {
    return me.Comp
}
func (me KpGenPipeline) GetLineNo() string {
	return me.LineNo
}

func loadGenPipeline(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenPipeline)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenPipeline)
	st.LineNo = lno
	st.Comp = "GenPipeline";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["pipeline_id"].(string)
	st.Names["_key"] = "pipeline_id"
	act.index["GenPipeline_" + name] = st.Me;
	st.MyName = name
	act.ApGenPipeline = append(act.ApGenPipeline, st)
	return 0
}

func (me KpGenPipeline) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "previous" { // gen-high.unit:207, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenPipeline[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenPipeline > gen-high.unit:207, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenPipeline > gen-high.unit:207, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenPipeline) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "GenStage" { // gen-high.unit:218, go-struct-rio.act:685
		for _, st := range me.ItsGenStage {
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
	        fmt.Printf("?No its %s for GenPipeline %s,%s > gen-high.unit:207, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpGenStage struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kprogramp int
	Kparallel_withp int
}

func (me KpGenStage) TypeName() string {
    return me.Comp
}
func (me KpGenStage) GetLineNo() string {
	return me.LineNo
}

func loadGenStage(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpGenStage)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApGenStage)
	st.LineNo = lno
	st.Comp = "GenStage";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kprogramp = -1
	st.Kparallel_withp = -1
	st.Kparentp = len( act.ApGenPipeline ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " GenStage has no GenPipeline parent\n") ;
		return 1
	}
	st.Parent = act.ApGenPipeline[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " GenStage under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApGenPipeline[ len( act.ApGenPipeline )-1 ].Childs = append(act.ApGenPipeline[ len( act.ApGenPipeline )-1 ].Childs, st)
	act.ApGenPipeline[ len( act.ApGenPipeline )-1 ].ItsGenStage = append(act.ApGenPipeline[ len( act.ApGenPipeline )-1 ].ItsGenStage, st)	// gen-high.unit:207, go-struct-rio.act:416
	name,_ := st.Names["stage_id"].(string)
	s := strconv.Itoa(st.Kparentp) + "_GenStage_" + name	// gen-high.unit:219, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApGenStage = append(act.ApGenStage, st)
	return 0
}

func (me KpGenStage) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "program" { // gen-high.unit:221, go-struct-rio.act:621
		if (me.Kprogramp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenProgram[ me.Kprogramp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "parallel_with" { // gen-high.unit:225, go-struct-rio.act:621
		if (me.Kparallel_withp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenStage[ me.Kparallel_withp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // gen-high.unit:207, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApGenPipeline[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "GenStage_parallel_with" && len(va) > 1) { // gen-high.unit:225, go-struct-rio.act:706
		for _, st := range glob.Dats.ApGenStage {
			if (st.Kparallel_withp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // gen-high.unit:218, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApGenStage[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,GenStage > gen-high.unit:218, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,GenStage > gen-high.unit:218, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpGenStage) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // gen-high.unit:207, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApGenPipeline[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "program" {
		if me.Kprogramp >= 0 {
			st := glob.Dats.ApGenProgram[ me.Kprogramp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "parallel_with" {
		if me.Kparallel_withp >= 0 {
			st := glob.Dats.ApGenStage[ me.Kparallel_withp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "GenStage_parallel_with") { // gen-high.unit:225, go-struct-rio.act:595
		for _, st := range glob.Dats.ApGenStage {
			if (st.Kparallel_withp == me.Me) {
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
	        fmt.Printf("?No its %s for GenStage %s,%s > gen-high.unit:218, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpProject struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kdomainp int
	Kmodelp int
	Kstrategyp int
	Khardwarep int
}

func (me KpProject) TypeName() string {
    return me.Comp
}
func (me KpProject) GetLineNo() string {
	return me.LineNo
}

func loadProject(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpProject)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApProject)
	st.LineNo = lno
	st.Comp = "Project";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kdomainp = -1
	st.Kmodelp = -1
	st.Kstrategyp = -1
	st.Khardwarep = -1
	name,_ := st.Names["project"].(string)
	st.Names["_key"] = "project"
	act.index["Project_" + name] = st.Me;
	st.MyName = name
	act.ApProject = append(act.ApProject, st)
	return 0
}

func (me KpProject) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "domain" { // omni.unit:13, go-struct-rio.act:621
		if (me.Kdomainp >= 0 && len(va) > 1) {
			return( glob.Dats.ApDomain[ me.Kdomainp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "model" { // omni.unit:14, go-struct-rio.act:621
		if (me.Kmodelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "strategy" { // omni.unit:15, go-struct-rio.act:621
		if (me.Kstrategyp >= 0 && len(va) > 1) {
			return( glob.Dats.ApStrategy[ me.Kstrategyp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "hardware" { // omni.unit:16, go-struct-rio.act:621
		if (me.Khardwarep >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Khardwarep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // omni.unit:8, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApProject[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Project > omni.unit:8, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Project > omni.unit:8, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpProject) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "domain" {
		if me.Kdomainp >= 0 {
			st := glob.Dats.ApDomain[ me.Kdomainp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model" {
		if me.Kmodelp >= 0 {
			st := glob.Dats.ApModel[ me.Kmodelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "strategy" {
		if me.Kstrategyp >= 0 {
			st := glob.Dats.ApStrategy[ me.Kstrategyp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "hardware" {
		if me.Khardwarep >= 0 {
			st := glob.Dats.ApHardware[ me.Khardwarep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Project %s,%s > omni.unit:8, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpDomain struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
}

func (me KpDomain) TypeName() string {
    return me.Comp
}
func (me KpDomain) GetLineNo() string {
	return me.LineNo
}

func loadDomain(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpDomain)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApDomain)
	st.LineNo = lno
	st.Comp = "Domain";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["name"].(string)
	st.Names["_key"] = "name"
	act.index["Domain_" + name] = st.Me;
	st.MyName = name
	act.ApDomain = append(act.ApDomain, st)
	return 0
}

func (me KpDomain) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "Project_domain" && len(va) > 1) { // omni.unit:13, go-struct-rio.act:706
		for _, st := range glob.Dats.ApProject {
			if (st.Kdomainp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:19, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApDomain[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Domain > omni.unit:19, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Domain > omni.unit:19, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpDomain) DoIts(glob *GlobT, va []string, lno string) int {
	if (va[0] == "Project_domain") { // omni.unit:13, go-struct-rio.act:595
		for _, st := range glob.Dats.ApProject {
			if (st.Kdomainp == me.Me) {
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
	        fmt.Printf("?No its %s for Domain %s,%s > omni.unit:19, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpHardware struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparent_hwp int
	Kemulationp int
	Knoise_modelp int
}

func (me KpHardware) TypeName() string {
    return me.Comp
}
func (me KpHardware) GetLineNo() string {
	return me.LineNo
}

func loadHardware(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpHardware)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApHardware)
	st.LineNo = lno
	st.Comp = "Hardware";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparent_hwp = -1
	st.Kemulationp = -1
	st.Knoise_modelp = -1
	name,_ := st.Names["hardware"].(string)
	st.Names["_key"] = "hardware"
	act.index["Hardware_" + name] = st.Me;
	st.MyName = name
	act.ApHardware = append(act.ApHardware, st)
	return 0
}

func (me KpHardware) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "parent_hw" { // omni.unit:31, go-struct-rio.act:621
		if (me.Kparent_hwp >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Kparent_hwp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "emulation" { // omni.unit:32, go-struct-rio.act:621
		if (me.Kemulationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Kemulationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "noise_model" { // omni.unit:33, go-struct-rio.act:621
		if (me.Knoise_modelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApConstraint[ me.Knoise_modelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Project_hardware" && len(va) > 1) { // omni.unit:16, go-struct-rio.act:706
		for _, st := range glob.Dats.ApProject {
			if (st.Khardwarep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Hardware_parent_hw" && len(va) > 1) { // omni.unit:31, go-struct-rio.act:706
		for _, st := range glob.Dats.ApHardware {
			if (st.Kparent_hwp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Hardware_emulation" && len(va) > 1) { // omni.unit:32, go-struct-rio.act:706
		for _, st := range glob.Dats.ApHardware {
			if (st.Kemulationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Model_hardware" && len(va) > 1) { // omni.unit:41, go-struct-rio.act:706
		for _, st := range glob.Dats.ApModel {
			if (st.Khardwarep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_hardware" && len(va) > 1) { // omni.unit:79, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Khardwarep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Kernel_hardware" && len(va) > 1) { // omni.unit:121, go-struct-rio.act:706
		for _, st := range glob.Dats.ApKernel {
			if (st.Khardwarep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Constraint_target_hw" && len(va) > 1) { // omni.unit:162, go-struct-rio.act:706
		for _, st := range glob.Dats.ApConstraint {
			if (st.Ktarget_hwp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Metric_target_hw" && len(va) > 1) { // omni.unit:172, go-struct-rio.act:706
		for _, st := range glob.Dats.ApMetric {
			if (st.Ktarget_hwp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Fusion_hardware" && len(va) > 1) { // omni.unit:188, go-struct-rio.act:706
		for _, st := range glob.Dats.ApFusion {
			if (st.Khardwarep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:26, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Hardware > omni.unit:26, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Hardware > omni.unit:26, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpHardware) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent_hw" {
		if me.Kparent_hwp >= 0 {
			st := glob.Dats.ApHardware[ me.Kparent_hwp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "emulation" {
		if me.Kemulationp >= 0 {
			st := glob.Dats.ApHardware[ me.Kemulationp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "noise_model" {
		if me.Knoise_modelp >= 0 {
			st := glob.Dats.ApConstraint[ me.Knoise_modelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Project_hardware") { // omni.unit:16, go-struct-rio.act:595
		for _, st := range glob.Dats.ApProject {
			if (st.Khardwarep == me.Me) {
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
	if (va[0] == "Hardware_parent_hw") { // omni.unit:31, go-struct-rio.act:595
		for _, st := range glob.Dats.ApHardware {
			if (st.Kparent_hwp == me.Me) {
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
	if (va[0] == "Hardware_emulation") { // omni.unit:32, go-struct-rio.act:595
		for _, st := range glob.Dats.ApHardware {
			if (st.Kemulationp == me.Me) {
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
	if (va[0] == "Model_hardware") { // omni.unit:41, go-struct-rio.act:595
		for _, st := range glob.Dats.ApModel {
			if (st.Khardwarep == me.Me) {
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
	if (va[0] == "Op_hardware") { // omni.unit:79, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Khardwarep == me.Me) {
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
	if (va[0] == "Kernel_hardware") { // omni.unit:121, go-struct-rio.act:595
		for _, st := range glob.Dats.ApKernel {
			if (st.Khardwarep == me.Me) {
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
	if (va[0] == "Constraint_target_hw") { // omni.unit:162, go-struct-rio.act:595
		for _, st := range glob.Dats.ApConstraint {
			if (st.Ktarget_hwp == me.Me) {
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
	if (va[0] == "Metric_target_hw") { // omni.unit:172, go-struct-rio.act:595
		for _, st := range glob.Dats.ApMetric {
			if (st.Ktarget_hwp == me.Me) {
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
	if (va[0] == "Fusion_hardware") { // omni.unit:188, go-struct-rio.act:595
		for _, st := range glob.Dats.ApFusion {
			if (st.Khardwarep == me.Me) {
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
	        fmt.Printf("?No its %s for Hardware %s,%s > omni.unit:26, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpModel struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Khardwarep int
	Ksearch_spacep int
	Kconfigp int
	ItsLayer [] *KpLayer 
	ItsTensor [] *KpTensor 
	ItsOp [] *KpOp 
	Childs [] Kp
}

func (me KpModel) TypeName() string {
    return me.Comp
}
func (me KpModel) GetLineNo() string {
	return me.LineNo
}

func loadModel(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpModel)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApModel)
	st.LineNo = lno
	st.Comp = "Model";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Khardwarep = -1
	st.Ksearch_spacep = -1
	st.Kconfigp = -1
	name,_ := st.Names["model"].(string)
	st.Names["_key"] = "model"
	act.index["Model_" + name] = st.Me;
	st.MyName = name
	act.ApModel = append(act.ApModel, st)
	return 0
}

func (me KpModel) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "hardware" { // omni.unit:41, go-struct-rio.act:621
		if (me.Khardwarep >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Khardwarep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "search_space" { // omni.unit:42, go-struct-rio.act:621
		if (me.Ksearch_spacep >= 0 && len(va) > 1) {
			return( glob.Dats.ApSearchSpace[ me.Ksearch_spacep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "config" { // omni.unit:43, go-struct-rio.act:621
		if (me.Kconfigp >= 0 && len(va) > 1) {
			return( glob.Dats.ApConfig[ me.Kconfigp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Project_model" && len(va) > 1) { // omni.unit:14, go-struct-rio.act:706
		for _, st := range glob.Dats.ApProject {
			if (st.Kmodelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Block_model" && len(va) > 1) { // omni.unit:60, go-struct-rio.act:706
		for _, st := range glob.Dats.ApBlock {
			if (st.Kmodelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Arg_model" && len(va) > 1) { // omni.unit:104, go-struct-rio.act:706
		for _, st := range glob.Dats.ApArg {
			if (st.Kmodelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "SearchSpace_target_model" && len(va) > 1) { // omni.unit:137, go-struct-rio.act:706
		for _, st := range glob.Dats.ApSearchSpace {
			if (st.Ktarget_modelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Constraint_model" && len(va) > 1) { // omni.unit:163, go-struct-rio.act:706
		for _, st := range glob.Dats.ApConstraint {
			if (st.Kmodelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Checkpoint_model" && len(va) > 1) { // omni.unit:180, go-struct-rio.act:706
		for _, st := range glob.Dats.ApCheckpoint {
			if (st.Kmodelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "ControlFlow_model1" && len(va) > 1) { // omni.unit:196, go-struct-rio.act:706
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel1p == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "ControlFlow_model2" && len(va) > 1) { // omni.unit:198, go-struct-rio.act:706
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel2p == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "ControlFlow_model3" && len(va) > 1) { // omni.unit:200, go-struct-rio.act:706
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel3p == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:36, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Model > omni.unit:36, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Model > omni.unit:36, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpModel) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Layer" { // omni.unit:47, go-struct-rio.act:685
		for _, st := range me.ItsLayer {
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
	if va[0] == "Tensor" { // omni.unit:63, go-struct-rio.act:685
		for _, st := range me.ItsTensor {
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
	if va[0] == "Op" { // omni.unit:72, go-struct-rio.act:685
		for _, st := range me.ItsOp {
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
	if va[0] == "hardware" {
		if me.Khardwarep >= 0 {
			st := glob.Dats.ApHardware[ me.Khardwarep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "search_space" {
		if me.Ksearch_spacep >= 0 {
			st := glob.Dats.ApSearchSpace[ me.Ksearch_spacep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "config" {
		if me.Kconfigp >= 0 {
			st := glob.Dats.ApConfig[ me.Kconfigp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Project_model") { // omni.unit:14, go-struct-rio.act:595
		for _, st := range glob.Dats.ApProject {
			if (st.Kmodelp == me.Me) {
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
	if (va[0] == "Block_model") { // omni.unit:60, go-struct-rio.act:595
		for _, st := range glob.Dats.ApBlock {
			if (st.Kmodelp == me.Me) {
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
	if (va[0] == "Arg_model") { // omni.unit:104, go-struct-rio.act:595
		for _, st := range glob.Dats.ApArg {
			if (st.Kmodelp == me.Me) {
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
	if (va[0] == "SearchSpace_target_model") { // omni.unit:137, go-struct-rio.act:595
		for _, st := range glob.Dats.ApSearchSpace {
			if (st.Ktarget_modelp == me.Me) {
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
	if (va[0] == "Constraint_model") { // omni.unit:163, go-struct-rio.act:595
		for _, st := range glob.Dats.ApConstraint {
			if (st.Kmodelp == me.Me) {
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
	if (va[0] == "Checkpoint_model") { // omni.unit:180, go-struct-rio.act:595
		for _, st := range glob.Dats.ApCheckpoint {
			if (st.Kmodelp == me.Me) {
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
	if (va[0] == "ControlFlow_model1") { // omni.unit:196, go-struct-rio.act:595
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel1p == me.Me) {
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
	if (va[0] == "ControlFlow_model2") { // omni.unit:198, go-struct-rio.act:595
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel2p == me.Me) {
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
	if (va[0] == "ControlFlow_model3") { // omni.unit:200, go-struct-rio.act:595
		for _, st := range glob.Dats.ApControlFlow {
			if (st.Kmodel3p == me.Me) {
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
	        fmt.Printf("?No its %s for Model %s,%s > omni.unit:36, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpLayer struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kparent_layerp int
}

func (me KpLayer) TypeName() string {
    return me.Comp
}
func (me KpLayer) GetLineNo() string {
	return me.LineNo
}

func loadLayer(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpLayer)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApLayer)
	st.LineNo = lno
	st.Comp = "Layer";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparent_layerp = -1
	st.Kparentp = len( act.ApModel ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Layer has no Model parent\n") ;
		return 1
	}
	st.Parent = act.ApModel[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Layer under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApModel[ len( act.ApModel )-1 ].Childs = append(act.ApModel[ len( act.ApModel )-1 ].Childs, st)
	act.ApModel[ len( act.ApModel )-1 ].ItsLayer = append(act.ApModel[ len( act.ApModel )-1 ].ItsLayer, st)	// omni.unit:36, go-struct-rio.act:416
	name,_ := st.Names["layer"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Layer_" + name	// omni.unit:51, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApLayer = append(act.ApLayer, st)
	return 0
}

func (me KpLayer) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "parent_layer" { // omni.unit:52, go-struct-rio.act:621
		if (me.Kparent_layerp >= 0 && len(va) > 1) {
			return( glob.Dats.ApLayer[ me.Kparent_layerp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // omni.unit:36, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Layer_parent_layer" && len(va) > 1) { // omni.unit:52, go-struct-rio.act:706
		for _, st := range glob.Dats.ApLayer {
			if (st.Kparent_layerp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_layer" && len(va) > 1) { // omni.unit:87, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Klayerp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:47, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApLayer[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Layer > omni.unit:47, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Layer > omni.unit:47, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpLayer) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // omni.unit:36, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApModel[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "parent_layer" {
		if me.Kparent_layerp >= 0 {
			st := glob.Dats.ApLayer[ me.Kparent_layerp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Layer_parent_layer") { // omni.unit:52, go-struct-rio.act:595
		for _, st := range glob.Dats.ApLayer {
			if (st.Kparent_layerp == me.Me) {
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
	if (va[0] == "Op_layer") { // omni.unit:87, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Klayerp == me.Me) {
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
	        fmt.Printf("?No its %s for Layer %s,%s > omni.unit:47, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpBlock struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kmodelp int
}

func (me KpBlock) TypeName() string {
    return me.Comp
}
func (me KpBlock) GetLineNo() string {
	return me.LineNo
}

func loadBlock(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpBlock)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApBlock)
	st.LineNo = lno
	st.Comp = "Block";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kmodelp = -1
	name,_ := st.Names["block"].(string)
	st.Names["_key"] = "block"
	act.index["Block_" + name] = st.Me;
	st.MyName = name
	act.ApBlock = append(act.ApBlock, st)
	return 0
}

func (me KpBlock) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "model" { // omni.unit:60, go-struct-rio.act:621
		if (me.Kmodelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // omni.unit:55, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApBlock[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Block > omni.unit:55, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Block > omni.unit:55, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpBlock) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "model" {
		if me.Kmodelp >= 0 {
			st := glob.Dats.ApModel[ me.Kmodelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Block %s,%s > omni.unit:55, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpTensor struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kproducerp int
	Kdistributionp int
}

func (me KpTensor) TypeName() string {
    return me.Comp
}
func (me KpTensor) GetLineNo() string {
	return me.LineNo
}

func loadTensor(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpTensor)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApTensor)
	st.LineNo = lno
	st.Comp = "Tensor";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kproducerp = -1
	st.Kdistributionp = -1
	st.Kparentp = len( act.ApModel ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Tensor has no Model parent\n") ;
		return 1
	}
	st.Parent = act.ApModel[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Tensor under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApModel[ len( act.ApModel )-1 ].Childs = append(act.ApModel[ len( act.ApModel )-1 ].Childs, st)
	act.ApModel[ len( act.ApModel )-1 ].ItsTensor = append(act.ApModel[ len( act.ApModel )-1 ].ItsTensor, st)	// omni.unit:36, go-struct-rio.act:416
	name,_ := st.Names["tensor"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Tensor_" + name	// omni.unit:67, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApTensor = append(act.ApTensor, st)
	return 0
}

func (me KpTensor) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "producer" { // omni.unit:68, go-struct-rio.act:621
		if (me.Kproducerp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Kproducerp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "distribution" { // omni.unit:69, go-struct-rio.act:621
		if (me.Kdistributionp >= 0 && len(va) > 1) {
			return( glob.Dats.ApEnergyFunction[ me.Kdistributionp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // omni.unit:36, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Op_predicate" && len(va) > 1) { // omni.unit:85, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Kpredicatep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:63, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApTensor[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Tensor > omni.unit:63, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Tensor > omni.unit:63, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpTensor) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // omni.unit:36, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApModel[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "producer" {
		if me.Kproducerp >= 0 {
			st := glob.Dats.ApOp[ me.Kproducerp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "distribution" {
		if me.Kdistributionp >= 0 {
			st := glob.Dats.ApEnergyFunction[ me.Kdistributionp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Op_predicate") { // omni.unit:85, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Kpredicatep == me.Me) {
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
	if (va[0] == "Arg_tensor") { // omni.unit:105, go-struct-rio.act:595
		for _, st := range glob.Dats.ApArg {
			if (st.Ktensorp == me.Me) {
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
	        fmt.Printf("?No its %s for Tensor %s,%s > omni.unit:63, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpOp struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Khardwarep int
	Kenergy_fnp int
	Ksearch_spacep int
	Kstrategyp int
	Kpredicatep int
	Knext_opp int
	Klayerp int
	ItsArg [] *KpArg 
	ItsControlFlow [] *KpControlFlow 
	Childs [] Kp
}

func (me KpOp) TypeName() string {
    return me.Comp
}
func (me KpOp) GetLineNo() string {
	return me.LineNo
}

func loadOp(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpOp)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApOp)
	st.LineNo = lno
	st.Comp = "Op";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Khardwarep = -1
	st.Kenergy_fnp = -1
	st.Ksearch_spacep = -1
	st.Kstrategyp = -1
	st.Kpredicatep = -1
	st.Knext_opp = -1
	st.Klayerp = -1
	st.Kparentp = len( act.ApModel ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Op has no Model parent\n") ;
		return 1
	}
	st.Parent = act.ApModel[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Op under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApModel[ len( act.ApModel )-1 ].Childs = append(act.ApModel[ len( act.ApModel )-1 ].Childs, st)
	act.ApModel[ len( act.ApModel )-1 ].ItsOp = append(act.ApModel[ len( act.ApModel )-1 ].ItsOp, st)	// omni.unit:36, go-struct-rio.act:416
	name,_ := st.Names["op"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Op_" + name	// omni.unit:76, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApOp = append(act.ApOp, st)
	return 0
}

func (me KpOp) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "hardware" { // omni.unit:79, go-struct-rio.act:621
		if (me.Khardwarep >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Khardwarep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "energy_fn" { // omni.unit:80, go-struct-rio.act:621
		if (me.Kenergy_fnp >= 0 && len(va) > 1) {
			return( glob.Dats.ApEnergyFunction[ me.Kenergy_fnp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "search_space" { // omni.unit:81, go-struct-rio.act:621
		if (me.Ksearch_spacep >= 0 && len(va) > 1) {
			return( glob.Dats.ApSearchSpace[ me.Ksearch_spacep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "strategy" { // omni.unit:82, go-struct-rio.act:621
		if (me.Kstrategyp >= 0 && len(va) > 1) {
			return( glob.Dats.ApStrategy[ me.Kstrategyp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "predicate" { // omni.unit:85, go-struct-rio.act:621
		if (me.Kpredicatep >= 0 && len(va) > 1) {
			return( glob.Dats.ApTensor[ me.Kpredicatep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "next_op" { // omni.unit:86, go-struct-rio.act:621
		if (me.Knext_opp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Knext_opp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "layer" { // omni.unit:87, go-struct-rio.act:621
		if (me.Klayerp >= 0 && len(va) > 1) {
			return( glob.Dats.ApLayer[ me.Klayerp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // omni.unit:36, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if (va[0] == "Tensor_producer" && len(va) > 1) { // omni.unit:68, go-struct-rio.act:706
		for _, st := range glob.Dats.ApTensor {
			if (st.Kproducerp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_next_op" && len(va) > 1) { // omni.unit:86, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Knext_opp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:72, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Op > omni.unit:72, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Op > omni.unit:72, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpOp) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Arg" { // omni.unit:98, go-struct-rio.act:685
		for _, st := range me.ItsArg {
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
	if va[0] == "ControlFlow" { // omni.unit:191, go-struct-rio.act:685
		for _, st := range me.ItsControlFlow {
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
	if va[0] == "parent" { // omni.unit:36, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApModel[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "hardware" {
		if me.Khardwarep >= 0 {
			st := glob.Dats.ApHardware[ me.Khardwarep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "energy_fn" {
		if me.Kenergy_fnp >= 0 {
			st := glob.Dats.ApEnergyFunction[ me.Kenergy_fnp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "search_space" {
		if me.Ksearch_spacep >= 0 {
			st := glob.Dats.ApSearchSpace[ me.Ksearch_spacep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "strategy" {
		if me.Kstrategyp >= 0 {
			st := glob.Dats.ApStrategy[ me.Kstrategyp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "predicate" {
		if me.Kpredicatep >= 0 {
			st := glob.Dats.ApTensor[ me.Kpredicatep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "next_op" {
		if me.Knext_opp >= 0 {
			st := glob.Dats.ApOp[ me.Knext_opp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "layer" {
		if me.Klayerp >= 0 {
			st := glob.Dats.ApLayer[ me.Klayerp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Tensor_producer") { // omni.unit:68, go-struct-rio.act:595
		for _, st := range glob.Dats.ApTensor {
			if (st.Kproducerp == me.Me) {
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
	if (va[0] == "Op_next_op") { // omni.unit:86, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Knext_opp == me.Me) {
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
	if (va[0] == "Constraint_target_op") { // omni.unit:164, go-struct-rio.act:595
		for _, st := range glob.Dats.ApConstraint {
			if (st.Ktarget_opp == me.Me) {
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
	        fmt.Printf("?No its %s for Op %s,%s > omni.unit:72, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpArg struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kmodelp int
	Ktensorp int
}

func (me KpArg) TypeName() string {
    return me.Comp
}
func (me KpArg) GetLineNo() string {
	return me.LineNo
}

func loadArg(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpArg)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApArg)
	st.LineNo = lno
	st.Comp = "Arg";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kmodelp = -1
	st.Ktensorp = -1
	st.Kparentp = len( act.ApOp ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Arg has no Op parent\n") ;
		return 1
	}
	st.Parent = act.ApOp[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Arg under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApOp[ len( act.ApOp )-1 ].Childs = append(act.ApOp[ len( act.ApOp )-1 ].Childs, st)
	act.ApOp[ len( act.ApOp )-1 ].ItsArg = append(act.ApOp[ len( act.ApOp )-1 ].ItsArg, st)	// omni.unit:72, go-struct-rio.act:416
	name,_ := st.Names["arg"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Arg_" + name	// omni.unit:102, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApArg = append(act.ApArg, st)
	return 0
}

func (me KpArg) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "model" { // omni.unit:104, go-struct-rio.act:621
		if (me.Kmodelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "tensor" { // omni.unit:105, go-struct-rio.act:621
		if (me.Ktensorp >= 0 && len(va) > 1) {
			return( glob.Dats.ApTensor[ me.Ktensorp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // omni.unit:72, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // omni.unit:98, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApArg[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Arg > omni.unit:98, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Arg > omni.unit:98, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpArg) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // omni.unit:72, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApOp[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model" {
		if me.Kmodelp >= 0 {
			st := glob.Dats.ApModel[ me.Kmodelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "tensor" {
		if me.Ktensorp >= 0 {
			st := glob.Dats.ApTensor[ me.Ktensorp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Arg %s,%s > omni.unit:98, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpConfig struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kschedulep int
}

func (me KpConfig) TypeName() string {
    return me.Comp
}
func (me KpConfig) GetLineNo() string {
	return me.LineNo
}

func loadConfig(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpConfig)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApConfig)
	st.LineNo = lno
	st.Comp = "Config";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kschedulep = -1
	name,_ := st.Names["config"].(string)
	st.Names["_key"] = "config"
	act.index["Config_" + name] = st.Me;
	st.MyName = name
	act.ApConfig = append(act.ApConfig, st)
	return 0
}

func (me KpConfig) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "schedule" { // omni.unit:112, go-struct-rio.act:621
		if (me.Kschedulep >= 0 && len(va) > 1) {
			return( glob.Dats.ApStrategy[ me.Kschedulep ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Model_config" && len(va) > 1) { // omni.unit:43, go-struct-rio.act:706
		for _, st := range glob.Dats.ApModel {
			if (st.Kconfigp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:107, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApConfig[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Config > omni.unit:107, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Config > omni.unit:107, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpConfig) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "schedule" {
		if me.Kschedulep >= 0 {
			st := glob.Dats.ApStrategy[ me.Kschedulep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Model_config") { // omni.unit:43, go-struct-rio.act:595
		for _, st := range glob.Dats.ApModel {
			if (st.Kconfigp == me.Me) {
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
	        fmt.Printf("?No its %s for Config %s,%s > omni.unit:107, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpKernel struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Khardwarep int
}

func (me KpKernel) TypeName() string {
    return me.Comp
}
func (me KpKernel) GetLineNo() string {
	return me.LineNo
}

func loadKernel(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpKernel)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApKernel)
	st.LineNo = lno
	st.Comp = "Kernel";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Khardwarep = -1
	name,_ := st.Names["kernel"].(string)
	st.Names["_key"] = "kernel"
	act.index["Kernel_" + name] = st.Me;
	st.MyName = name
	act.ApKernel = append(act.ApKernel, st)
	return 0
}

func (me KpKernel) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "hardware" { // omni.unit:121, go-struct-rio.act:621
		if (me.Khardwarep >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Khardwarep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // omni.unit:115, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApKernel[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Kernel > omni.unit:115, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Kernel > omni.unit:115, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpKernel) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "hardware" {
		if me.Khardwarep >= 0 {
			st := glob.Dats.ApHardware[ me.Khardwarep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Kernel %s,%s > omni.unit:115, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpEnergyFunction struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
}

func (me KpEnergyFunction) TypeName() string {
    return me.Comp
}
func (me KpEnergyFunction) GetLineNo() string {
	return me.LineNo
}

func loadEnergyFunction(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpEnergyFunction)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApEnergyFunction)
	st.LineNo = lno
	st.Comp = "EnergyFunction";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	name,_ := st.Names["energy_fn"].(string)
	st.Names["_key"] = "energy_fn"
	act.index["EnergyFunction_" + name] = st.Me;
	st.MyName = name
	act.ApEnergyFunction = append(act.ApEnergyFunction, st)
	return 0
}

func (me KpEnergyFunction) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "Tensor_distribution" && len(va) > 1) { // omni.unit:69, go-struct-rio.act:706
		for _, st := range glob.Dats.ApTensor {
			if (st.Kdistributionp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_energy_fn" && len(va) > 1) { // omni.unit:80, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Kenergy_fnp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:124, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApEnergyFunction[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,EnergyFunction > omni.unit:124, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,EnergyFunction > omni.unit:124, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpEnergyFunction) DoIts(glob *GlobT, va []string, lno string) int {
	if (va[0] == "Tensor_distribution") { // omni.unit:69, go-struct-rio.act:595
		for _, st := range glob.Dats.ApTensor {
			if (st.Kdistributionp == me.Me) {
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
	if (va[0] == "Op_energy_fn") { // omni.unit:80, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Kenergy_fnp == me.Me) {
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
	        fmt.Printf("?No its %s for EnergyFunction %s,%s > omni.unit:124, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpSearchSpace struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ktarget_modelp int
	ItsDimension [] *KpDimension 
	Childs [] Kp
}

func (me KpSearchSpace) TypeName() string {
    return me.Comp
}
func (me KpSearchSpace) GetLineNo() string {
	return me.LineNo
}

func loadSearchSpace(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpSearchSpace)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApSearchSpace)
	st.LineNo = lno
	st.Comp = "SearchSpace";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ktarget_modelp = -1
	name,_ := st.Names["space"].(string)
	st.Names["_key"] = "space"
	act.index["SearchSpace_" + name] = st.Me;
	st.MyName = name
	act.ApSearchSpace = append(act.ApSearchSpace, st)
	return 0
}

func (me KpSearchSpace) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "target_model" { // omni.unit:137, go-struct-rio.act:621
		if (me.Ktarget_modelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Ktarget_modelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Model_search_space" && len(va) > 1) { // omni.unit:42, go-struct-rio.act:706
		for _, st := range glob.Dats.ApModel {
			if (st.Ksearch_spacep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_search_space" && len(va) > 1) { // omni.unit:81, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Ksearch_spacep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Strategy_search_space" && len(va) > 1) { // omni.unit:152, go-struct-rio.act:706
		for _, st := range glob.Dats.ApStrategy {
			if (st.Ksearch_spacep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:132, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSearchSpace[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,SearchSpace > omni.unit:132, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,SearchSpace > omni.unit:132, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpSearchSpace) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Dimension" { // omni.unit:140, go-struct-rio.act:685
		for _, st := range me.ItsDimension {
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
	if va[0] == "target_model" {
		if me.Ktarget_modelp >= 0 {
			st := glob.Dats.ApModel[ me.Ktarget_modelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Model_search_space") { // omni.unit:42, go-struct-rio.act:595
		for _, st := range glob.Dats.ApModel {
			if (st.Ksearch_spacep == me.Me) {
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
	if (va[0] == "Op_search_space") { // omni.unit:81, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Ksearch_spacep == me.Me) {
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
	if (va[0] == "Strategy_search_space") { // omni.unit:152, go-struct-rio.act:595
		for _, st := range glob.Dats.ApStrategy {
			if (st.Ksearch_spacep == me.Me) {
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
	        fmt.Printf("?No its %s for SearchSpace %s,%s > omni.unit:132, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpDimension struct {
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

func (me KpDimension) TypeName() string {
    return me.Comp
}
func (me KpDimension) GetLineNo() string {
	return me.LineNo
}

func loadDimension(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpDimension)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApDimension)
	st.LineNo = lno
	st.Comp = "Dimension";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kparentp = len( act.ApSearchSpace ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Dimension has no SearchSpace parent\n") ;
		return 1
	}
	st.Parent = act.ApSearchSpace[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " Dimension under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApSearchSpace[ len( act.ApSearchSpace )-1 ].Childs = append(act.ApSearchSpace[ len( act.ApSearchSpace )-1 ].Childs, st)
	act.ApSearchSpace[ len( act.ApSearchSpace )-1 ].ItsDimension = append(act.ApSearchSpace[ len( act.ApSearchSpace )-1 ].ItsDimension, st)	// omni.unit:132, go-struct-rio.act:416
	name,_ := st.Names["dimension"].(string)
	s := strconv.Itoa(st.Kparentp) + "_Dimension_" + name	// omni.unit:144, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApDimension = append(act.ApDimension, st)
	return 0
}

func (me KpDimension) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // omni.unit:132, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSearchSpace[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // omni.unit:140, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApDimension[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Dimension > omni.unit:140, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Dimension > omni.unit:140, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpDimension) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // omni.unit:132, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSearchSpace[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Dimension %s,%s > omni.unit:140, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpStrategy struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ksearch_spacep int
	Kfitnessp int
}

func (me KpStrategy) TypeName() string {
    return me.Comp
}
func (me KpStrategy) GetLineNo() string {
	return me.LineNo
}

func loadStrategy(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpStrategy)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApStrategy)
	st.LineNo = lno
	st.Comp = "Strategy";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ksearch_spacep = -1
	st.Kfitnessp = -1
	name,_ := st.Names["strategy"].(string)
	st.Names["_key"] = "strategy"
	act.index["Strategy_" + name] = st.Me;
	st.MyName = name
	act.ApStrategy = append(act.ApStrategy, st)
	return 0
}

func (me KpStrategy) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "search_space" { // omni.unit:152, go-struct-rio.act:621
		if (me.Ksearch_spacep >= 0 && len(va) > 1) {
			return( glob.Dats.ApSearchSpace[ me.Ksearch_spacep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "fitness" { // omni.unit:153, go-struct-rio.act:621
		if (me.Kfitnessp >= 0 && len(va) > 1) {
			return( glob.Dats.ApMetric[ me.Kfitnessp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Project_strategy" && len(va) > 1) { // omni.unit:15, go-struct-rio.act:706
		for _, st := range glob.Dats.ApProject {
			if (st.Kstrategyp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Op_strategy" && len(va) > 1) { // omni.unit:82, go-struct-rio.act:706
		for _, st := range glob.Dats.ApOp {
			if (st.Kstrategyp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Config_schedule" && len(va) > 1) { // omni.unit:112, go-struct-rio.act:706
		for _, st := range glob.Dats.ApConfig {
			if (st.Kschedulep == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:147, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApStrategy[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Strategy > omni.unit:147, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Strategy > omni.unit:147, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpStrategy) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "search_space" {
		if me.Ksearch_spacep >= 0 {
			st := glob.Dats.ApSearchSpace[ me.Ksearch_spacep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "fitness" {
		if me.Kfitnessp >= 0 {
			st := glob.Dats.ApMetric[ me.Kfitnessp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Project_strategy") { // omni.unit:15, go-struct-rio.act:595
		for _, st := range glob.Dats.ApProject {
			if (st.Kstrategyp == me.Me) {
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
	if (va[0] == "Op_strategy") { // omni.unit:82, go-struct-rio.act:595
		for _, st := range glob.Dats.ApOp {
			if (st.Kstrategyp == me.Me) {
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
	if (va[0] == "Config_schedule") { // omni.unit:112, go-struct-rio.act:595
		for _, st := range glob.Dats.ApConfig {
			if (st.Kschedulep == me.Me) {
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
	        fmt.Printf("?No its %s for Strategy %s,%s > omni.unit:147, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpConstraint struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ktarget_hwp int
	Kmodelp int
	Ktarget_opp int
}

func (me KpConstraint) TypeName() string {
    return me.Comp
}
func (me KpConstraint) GetLineNo() string {
	return me.LineNo
}

func loadConstraint(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpConstraint)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApConstraint)
	st.LineNo = lno
	st.Comp = "Constraint";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ktarget_hwp = -1
	st.Kmodelp = -1
	st.Ktarget_opp = -1
	name,_ := st.Names["constraint_id"].(string)
	st.Names["_key"] = "constraint_id"
	act.index["Constraint_" + name] = st.Me;
	st.MyName = name
	act.ApConstraint = append(act.ApConstraint, st)
	return 0
}

func (me KpConstraint) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "target_hw" { // omni.unit:162, go-struct-rio.act:621
		if (me.Ktarget_hwp >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Ktarget_hwp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "model" { // omni.unit:163, go-struct-rio.act:621
		if (me.Kmodelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "target_op" { // omni.unit:164, go-struct-rio.act:621
		if (me.Ktarget_opp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Ktarget_opp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Hardware_noise_model" && len(va) > 1) { // omni.unit:33, go-struct-rio.act:706
		for _, st := range glob.Dats.ApHardware {
			if (st.Knoise_modelp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:157, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApConstraint[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Constraint > omni.unit:157, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Constraint > omni.unit:157, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpConstraint) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "target_hw" {
		if me.Ktarget_hwp >= 0 {
			st := glob.Dats.ApHardware[ me.Ktarget_hwp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model" {
		if me.Kmodelp >= 0 {
			st := glob.Dats.ApModel[ me.Kmodelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "target_op" {
		if me.Ktarget_opp >= 0 {
			st := glob.Dats.ApOp[ me.Ktarget_opp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Hardware_noise_model") { // omni.unit:33, go-struct-rio.act:595
		for _, st := range glob.Dats.ApHardware {
			if (st.Knoise_modelp == me.Me) {
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
	        fmt.Printf("?No its %s for Constraint %s,%s > omni.unit:157, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpMetric struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Ktarget_hwp int
}

func (me KpMetric) TypeName() string {
    return me.Comp
}
func (me KpMetric) GetLineNo() string {
	return me.LineNo
}

func loadMetric(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpMetric)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApMetric)
	st.LineNo = lno
	st.Comp = "Metric";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Ktarget_hwp = -1
	name,_ := st.Names["metric"].(string)
	st.Names["_key"] = "metric"
	act.index["Metric_" + name] = st.Me;
	st.MyName = name
	act.ApMetric = append(act.ApMetric, st)
	return 0
}

func (me KpMetric) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "target_hw" { // omni.unit:172, go-struct-rio.act:621
		if (me.Ktarget_hwp >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Ktarget_hwp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Strategy_fitness" && len(va) > 1) { // omni.unit:153, go-struct-rio.act:706
		for _, st := range glob.Dats.ApStrategy {
			if (st.Kfitnessp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // omni.unit:167, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApMetric[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Metric > omni.unit:167, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Metric > omni.unit:167, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpMetric) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "target_hw" {
		if me.Ktarget_hwp >= 0 {
			st := glob.Dats.ApHardware[ me.Ktarget_hwp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Strategy_fitness") { // omni.unit:153, go-struct-rio.act:595
		for _, st := range glob.Dats.ApStrategy {
			if (st.Kfitnessp == me.Me) {
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
	        fmt.Printf("?No its %s for Metric %s,%s > omni.unit:167, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpCheckpoint struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kmodelp int
}

func (me KpCheckpoint) TypeName() string {
    return me.Comp
}
func (me KpCheckpoint) GetLineNo() string {
	return me.LineNo
}

func loadCheckpoint(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpCheckpoint)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApCheckpoint)
	st.LineNo = lno
	st.Comp = "Checkpoint";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kmodelp = -1
	name,_ := st.Names["checkpoint_id"].(string)
	st.Names["_key"] = "checkpoint_id"
	act.index["Checkpoint_" + name] = st.Me;
	st.MyName = name
	act.ApCheckpoint = append(act.ApCheckpoint, st)
	return 0
}

func (me KpCheckpoint) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "model" { // omni.unit:180, go-struct-rio.act:621
		if (me.Kmodelp >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodelp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // omni.unit:175, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApCheckpoint[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Checkpoint > omni.unit:175, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Checkpoint > omni.unit:175, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpCheckpoint) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "model" {
		if me.Kmodelp >= 0 {
			st := glob.Dats.ApModel[ me.Kmodelp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Checkpoint %s,%s > omni.unit:175, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpFusion struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Khardwarep int
}

func (me KpFusion) TypeName() string {
    return me.Comp
}
func (me KpFusion) GetLineNo() string {
	return me.LineNo
}

func loadFusion(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpFusion)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApFusion)
	st.LineNo = lno
	st.Comp = "Fusion";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Khardwarep = -1
	name,_ := st.Names["fusion"].(string)
	st.Names["_key"] = "fusion"
	act.index["Fusion_" + name] = st.Me;
	st.MyName = name
	act.ApFusion = append(act.ApFusion, st)
	return 0
}

func (me KpFusion) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "hardware" { // omni.unit:188, go-struct-rio.act:621
		if (me.Khardwarep >= 0 && len(va) > 1) {
			return( glob.Dats.ApHardware[ me.Khardwarep ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "previous" { // omni.unit:183, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApFusion[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,Fusion > omni.unit:183, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,Fusion > omni.unit:183, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpFusion) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "hardware" {
		if me.Khardwarep >= 0 {
			st := glob.Dats.ApHardware[ me.Khardwarep ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Fusion %s,%s > omni.unit:183, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpControlFlow struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]any
	Kparentp int
	Kmodel1p int
	Kmodel2p int
	Kmodel3p int
}

func (me KpControlFlow) TypeName() string {
    return me.Comp
}
func (me KpControlFlow) GetLineNo() string {
	return me.LineNo
}

func loadControlFlow(act *ActT, ln string, pos int, lno string, flag []string, names map[string]any) int {
	st := new(KpControlFlow)
	st.Names = names
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApControlFlow)
	st.LineNo = lno
	st.Comp = "ControlFlow";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	st.Kmodel1p = -1
	st.Kmodel2p = -1
	st.Kmodel3p = -1
	st.Kparentp = len( act.ApOp ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " ControlFlow has no Op parent\n") ;
		return 1
	}
	st.Parent = act.ApOp[st.Kparentp].MyName
	par,ok := st.Names["parent"].(string)
	if ok && par != st.Parent {
		print(lno + " ControlFlow under wrong parent " + st.Parent + ", " +  par + "\n") ;
	}
	act.ApOp[ len( act.ApOp )-1 ].Childs = append(act.ApOp[ len( act.ApOp )-1 ].Childs, st)
	act.ApOp[ len( act.ApOp )-1 ].ItsControlFlow = append(act.ApOp[ len( act.ApOp )-1 ].ItsControlFlow, st)	// omni.unit:72, go-struct-rio.act:416
	name,_ := st.Names["control"].(string)
	s := strconv.Itoa(st.Kparentp) + "_ControlFlow_" + name	// omni.unit:195, go-struct-rio.act:464
	act.index[s] = st.Me;
	st.MyName = name
	act.ApControlFlow = append(act.ApControlFlow, st)
	return 0
}

func (me KpControlFlow) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "model1" { // omni.unit:196, go-struct-rio.act:621
		if (me.Kmodel1p >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodel1p ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "model2" { // omni.unit:198, go-struct-rio.act:621
		if (me.Kmodel2p >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodel2p ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "model3" { // omni.unit:200, go-struct-rio.act:621
		if (me.Kmodel3p >= 0 && len(va) > 1) {
			return( glob.Dats.ApModel[ me.Kmodel3p ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // omni.unit:72, go-struct-rio.act:585
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApOp[ me.Kparentp ].GetVar(glob, va[1:], lno) );
		}
	}
	if va[0] == "previous" { // omni.unit:191, go-struct-rio.act:178
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApControlFlow[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	if len(va) > 1 {
		msg := fmt.Sprintf("?%s.?:%s,%s,ControlFlow > omni.unit:191, go-struct-rio.act:185?", va[0], lno, me.LineNo)
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
		rr := fmt.Sprintf("?%s?:%s,%s,ControlFlow > omni.unit:191, go-struct-rio.act:197?", va[0], lno, me.LineNo) 
		return false,rr
	}
	rr := me.Names[va[0]].(string)
	return true,rr
}

func (me KpControlFlow) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // omni.unit:72, go-struct-rio.act:570
		if me.Kparentp >= 0 {
			st := glob.Dats.ApOp[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model1" {
		if me.Kmodel1p >= 0 {
			st := glob.Dats.ApModel[ me.Kmodel1p ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model2" {
		if me.Kmodel2p >= 0 {
			st := glob.Dats.ApModel[ me.Kmodel2p ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "model3" {
		if me.Kmodel3p >= 0 {
			st := glob.Dats.ApModel[ me.Kmodel3p ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for ControlFlow %s,%s > omni.unit:191, go-struct-rio.act:222?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

