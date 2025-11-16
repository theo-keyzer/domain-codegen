package main

import (
	"fmt"
	"strconv"
)

type Kp interface {
	DoIts(glob *GlobT, va []string, lno string) int
	GetVar(glob *GlobT, va []string, lno string) (bool, string)
	GetLineNo() string
}

type KpComp struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]string
	Kparentp int
	ItsElement [] *KpElement 
	Childs [] Kp
}

func (me KpComp) GetLineNo() string {
	return me.LineNo
}

func loadComp(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpComp)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApComp)
	st.LineNo = lno
	st.Comp = "Comp";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["name"] = getw(ln,p)
	p,st.Names["nop"] = getw(ln,p)
	p,st.Names["parent"] = getw(ln,p)
	p,st.Names["find"] = getw(ln,p)
	p,st.Names["doc"] = getws(ln,p)
	st.Kparentp = -1
	name,_ := st.Names["name"]
	act.index["Comp_" + name] = st.Me;
	st.MyName = name
	act.ApComp = append(act.ApComp, st)
	return 0
}

func (me KpComp) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "parent" { // unit.unit:8, g_struct.act:580
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kparentp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "Comp_parent" && len(va) > 1) { // unit.unit:8, g_struct.act:659
		for _, st := range glob.Dats.ApComp {
			if (st.Kparentp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Element_comp" && len(va) > 1) { // unit.unit:35, g_struct.act:659
		for _, st := range glob.Dats.ApElement {
			if (st.Kcompp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // unit.unit:2, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Comp > unit.unit:2, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpComp) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Element" { // unit.unit:18, g_struct.act:637
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
				return( st.DoIts(glob, va[1:], me.LineNo) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if (va[0] == "Comp_parent") { // unit.unit:8, g_struct.act:553
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
	if (va[0] == "Element_comp") { // unit.unit:35, g_struct.act:553
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
	        fmt.Printf("?No its %s for Comp %s,%s > unit.unit:2, g_struct.act:200?", va[0], lno, me.LineNo)
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
	Names map[string]string
	Kparentp int
	Kcompp int
	ItsOpt [] *KpOpt 
	Childs [] Kp
}

func (me KpElement) GetLineNo() string {
	return me.LineNo
}

func loadElement(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpElement)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApElement)
	st.LineNo = lno
	st.Comp = "Element";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["name"] = getw(ln,p)
	p,st.Names["type"] = getw(ln,p)
	p,st.Names["comp"] = getw(ln,p)
	p,st.Names["check"] = getw(ln,p)
	p,st.Names["doc"] = getws(ln,p)
	st.Kcompp = -1
	st.Kparentp = len( act.ApComp ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Element has no Comp parent") ;
		return 1
	}
	st.Parent = act.ApComp[st.Kparentp].MyName
	act.ApComp[ len( act.ApComp )-1 ].Childs = append(act.ApComp[ len( act.ApComp )-1 ].Childs, st)
	act.ApComp[ len( act.ApComp )-1 ].ItsElement = append(act.ApComp[ len( act.ApComp )-1 ].ItsElement, st)	// unit.unit:2, g_struct.act:390
	name,_ := st.Names["name"]
	s := strconv.Itoa(st.Kparentp) + "_Element_" + name	// unit.unit:20, g_struct.act:436
	act.index[s] = st.Me;
	act.ApElement = append(act.ApElement, st)
	return 0
}

func (me KpElement) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "comp" { // unit.unit:35, g_struct.act:580
		if (me.Kcompp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kcompp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // unit.unit:2, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApComp[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // unit.unit:18, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApElement[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Element > unit.unit:18, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpElement) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Opt" { // unit.unit:48, g_struct.act:637
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
	if va[0] == "parent" { // unit.unit:2, g_struct.act:527
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
				return( st.DoIts(glob, va[1:], me.LineNo) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Element %s,%s > unit.unit:18, g_struct.act:200?", va[0], lno, me.LineNo)
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
	Names map[string]string
	Kparentp int
}

func (me KpOpt) GetLineNo() string {
	return me.LineNo
}

func loadOpt(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpOpt)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApOpt)
	st.LineNo = lno
	st.Comp = "Opt";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["name"] = getw(ln,p)
	p,st.Names["pad"] = getw(ln,p)
	p,st.Names["doc"] = getws(ln,p)
	st.Kparentp = len( act.ApElement ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Opt has no Element parent") ;
		return 1
	}
	st.Parent = act.ApElement[st.Kparentp].MyName
	act.ApElement[ len( act.ApElement )-1 ].Childs = append(act.ApElement[ len( act.ApElement )-1 ].Childs, st)
	act.ApElement[ len( act.ApElement )-1 ].ItsOpt = append(act.ApElement[ len( act.ApElement )-1 ].ItsOpt, st)	// unit.unit:18, g_struct.act:390
	name,_ := st.Names["name"]
	s := strconv.Itoa(st.Kparentp) + "_Opt_" + name	// unit.unit:54, g_struct.act:436
	act.index[s] = st.Me;
	act.ApOpt = append(act.ApOpt, st)
	return 0
}

func (me KpOpt) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // unit.unit:18, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApElement[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // unit.unit:48, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApOpt[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Opt > unit.unit:48, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpOpt) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // unit.unit:18, g_struct.act:527
		if me.Kparentp >= 0 {
			st := glob.Dats.ApElement[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Opt %s,%s > unit.unit:48, g_struct.act:200?", va[0], lno, me.LineNo)
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
		print(lno + " All has no Actor parent") ;
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
		print(lno + " Du has no Actor parent") ;
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
		print(lno + " New has no Actor parent") ;
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
		print(lno + " Refs has no Actor parent") ;
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
		print(lno + " Var has no Actor parent") ;
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
		print(lno + " Its has no Actor parent") ;
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
		print(lno + " C has no Actor parent") ;
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
		print(lno + " Cs has no Actor parent") ;
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
		print(lno + " Out has no Actor parent") ;
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
		print(lno + " In has no Actor parent") ;
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
		print(lno + " Break has no Actor parent") ;
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
		print(lno + " Add has no Actor parent") ;
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
		print(lno + " This has no Actor parent") ;
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
		print(lno + " Replace has no Actor parent") ;
		return 1
	}
	act.ApActor[ len( act.ApActor )-1 ].Childs = append(act.ApActor[ len( act.ApActor )-1 ].Childs, st)
	act.ApReplace = append(act.ApReplace, st)
	return 0
}

type KpArtifact struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]string
	ItsLink [] *KpLink 
	ItsO [] *KpO 
	ItsSection [] *KpSection 
	Childs [] Kp
}

func (me KpArtifact) GetLineNo() string {
	return me.LineNo
}

func loadArtifact(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpArtifact)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApArtifact)
	st.LineNo = lno
	st.Comp = "Artifact";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["name"] = getw(ln,p)
	p,st.Names["category"] = getw(ln,p)
	p,st.Names["topic_type"] = getw(ln,p)
	p,st.Names["doc"] = getws(ln,p)
	name,_ := st.Names["name"]
	act.index["Artifact_" + name] = st.Me;
	st.MyName = name
	act.ApArtifact = append(act.ApArtifact, st)
	return 0
}

func (me KpArtifact) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "Link_concept" && len(va) > 1) { // artifact.unit:21, g_struct.act:659
		for _, st := range glob.Dats.ApLink {
			if (st.Kconceptp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if (va[0] == "Link_relation" && len(va) > 1) { // artifact.unit:22, g_struct.act:659
		for _, st := range glob.Dats.ApLink {
			if (st.Krelationp == me.Me) {
				return (st.GetVar(glob, va[1:], lno) )
			}
		}
	}
	if va[0] == "previous" { // artifact.unit:6, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Artifact > artifact.unit:6, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpArtifact) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "Link" { // artifact.unit:16, g_struct.act:637
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
	if va[0] == "O" { // artifact.unit:24, g_struct.act:637
		for _, st := range me.ItsO {
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
	if va[0] == "Section" { // artifact.unit:31, g_struct.act:637
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
	if (va[0] == "Link_concept") { // artifact.unit:21, g_struct.act:553
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
	if (va[0] == "Link_relation") { // artifact.unit:22, g_struct.act:553
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
	        fmt.Printf("?No its %s for Artifact %s,%s > artifact.unit:6, g_struct.act:200?", va[0], lno, me.LineNo)
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
	Names map[string]string
	Kparentp int
	Kconceptp int
	Krelationp int
}

func (me KpLink) GetLineNo() string {
	return me.LineNo
}

func loadLink(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpLink)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApLink)
	st.LineNo = lno
	st.Comp = "Link";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["concept"] = getw(ln,p)
	p,st.Names["relation"] = getw(ln,p)
	st.Kconceptp = -1
	st.Krelationp = -1
	st.Kparentp = len( act.ApArtifact ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Link has no Artifact parent") ;
		return 1
	}
	st.Parent = act.ApArtifact[st.Kparentp].MyName
	act.ApArtifact[ len( act.ApArtifact )-1 ].Childs = append(act.ApArtifact[ len( act.ApArtifact )-1 ].Childs, st)
	act.ApArtifact[ len( act.ApArtifact )-1 ].ItsLink = append(act.ApArtifact[ len( act.ApArtifact )-1 ].ItsLink, st)	// artifact.unit:6, g_struct.act:390
	act.ApLink = append(act.ApLink, st)
	return 0
}

func (me KpLink) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if va[0] == "concept" { // artifact.unit:21, g_struct.act:580
		if (me.Kconceptp >= 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Kconceptp ].GetVar(glob, va[1:], lno) )
		}
	}
	if va[0] == "relation" { // artifact.unit:22, g_struct.act:580
		if (me.Krelationp >= 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Krelationp ].GetVar(glob, va[1:], lno) )
		}
	}
	if (va[0] == "parent") { // artifact.unit:6, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // artifact.unit:16, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApLink[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Link > artifact.unit:16, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpLink) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // artifact.unit:6, g_struct.act:527
		if me.Kparentp >= 0 {
			st := glob.Dats.ApArtifact[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "concept" {
		if me.Kconceptp >= 0 {
			st := glob.Dats.ApArtifact[ me.Kconceptp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], me.LineNo) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	if va[0] == "relation" {
		if me.Krelationp >= 0 {
			st := glob.Dats.ApArtifact[ me.Krelationp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], me.LineNo) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for Link %s,%s > artifact.unit:16, g_struct.act:200?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpO struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]string
	Kparentp int
}

func (me KpO) GetLineNo() string {
	return me.LineNo
}

func loadO(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpO)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApO)
	st.LineNo = lno
	st.Comp = "O";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["note"] = getws(ln,p)
	st.Kparentp = len( act.ApArtifact ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " O has no Artifact parent") ;
		return 1
	}
	st.Parent = act.ApArtifact[st.Kparentp].MyName
	act.ApArtifact[ len( act.ApArtifact )-1 ].Childs = append(act.ApArtifact[ len( act.ApArtifact )-1 ].Childs, st)
	act.ApArtifact[ len( act.ApArtifact )-1 ].ItsO = append(act.ApArtifact[ len( act.ApArtifact )-1 ].ItsO, st)	// artifact.unit:6, g_struct.act:390
	act.ApO = append(act.ApO, st)
	return 0
}

func (me KpO) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // artifact.unit:6, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // artifact.unit:24, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApO[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,O > artifact.unit:24, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpO) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // artifact.unit:6, g_struct.act:527
		if me.Kparentp >= 0 {
			st := glob.Dats.ApArtifact[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for O %s,%s > artifact.unit:24, g_struct.act:200?", va[0], lno, me.LineNo)
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
	Names map[string]string
	Kparentp int
	ItsD [] *KpD 
	Childs [] Kp
}

func (me KpSection) GetLineNo() string {
	return me.LineNo
}

func loadSection(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpSection)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApSection)
	st.LineNo = lno
	st.Comp = "Section";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["name"] = getw(ln,p)
	p,st.Names["level"] = getw(ln,p)
	p,st.Names["note"] = getws(ln,p)
	st.Kparentp = len( act.ApArtifact ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " Section has no Artifact parent") ;
		return 1
	}
	st.Parent = act.ApArtifact[st.Kparentp].MyName
	act.ApArtifact[ len( act.ApArtifact )-1 ].Childs = append(act.ApArtifact[ len( act.ApArtifact )-1 ].Childs, st)
	act.ApArtifact[ len( act.ApArtifact )-1 ].ItsSection = append(act.ApArtifact[ len( act.ApArtifact )-1 ].ItsSection, st)	// artifact.unit:6, g_struct.act:390
	name,_ := st.Names["name"]
	s := strconv.Itoa(st.Kparentp) + "_Section_" + name	// artifact.unit:36, g_struct.act:436
	act.index[s] = st.Me;
	act.ApSection = append(act.ApSection, st)
	return 0
}

func (me KpSection) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // artifact.unit:6, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApArtifact[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // artifact.unit:31, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApSection[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,Section > artifact.unit:31, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpSection) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "D" { // artifact.unit:40, g_struct.act:637
		for _, st := range me.ItsD {
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
	if va[0] == "parent" { // artifact.unit:6, g_struct.act:527
		if me.Kparentp >= 0 {
			st := glob.Dats.ApArtifact[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
		if va[0] == "level" && len(va) > 1 && me.Kparentp >= 0 { // artifact.unit:37, g_struct.act:211
			pos := 0
			v, ok := me.Names["level"]
			if ok {
				pos, _ = strconv.Atoi(v)
			}
			if va[1] == "down" && pos > 0 {
				pst := glob.Dats.ApArtifact[me.Kparentp]
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
					v2, ok2 := st.Names["level"]
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
				pst := glob.Dats.ApArtifact[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsSection {
					pos2 := 0
					v2, ok2 := st.Names["level"]
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
				pst := glob.Dats.ApArtifact[me.Kparentp]
				isin := false
				prev := 0
				for _, st := range pst.ItsSection {
					pos2 := 0
					v2, ok2 := st.Names["level"]
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
				pst := glob.Dats.ApArtifact[me.Kparentp]
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
					v2, ok2 := st.Names["level"]
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
	        fmt.Printf("?No its %s for Section %s,%s > artifact.unit:31, g_struct.act:200?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

type KpD struct {
	Kp
	Me int
	MyName string
	Parent string
	LineNo string
	Comp string
	Flags [] string
	Names map[string]string
	Kparentp int
}

func (me KpD) GetLineNo() string {
	return me.LineNo
}

func loadD(act *ActT, ln string, pos int, lno string, flag []string) int {
	p := pos
	st := new(KpD)
	st.Names = make(map[string]string) 
      st.MyName = ""
      st.Parent = ""
	st.Me = len(act.ApD)
	st.LineNo = lno
	st.Comp = "D";
	st.Flags = flag;
	st.Names["kComp"] = st.Comp
	st.Names["kMe"] = strconv.Itoa(st.Me)
	st.Names["_lno"] = lno
	p,st.Names["note"] = getws(ln,p)
	st.Kparentp = len( act.ApSection ) - 1;
	st.Names["kParentp"] = strconv.Itoa(st.Kparentp)
	if (st.Kparentp < 0 ) { 
		print(lno + " D has no Section parent") ;
		return 1
	}
	st.Parent = act.ApSection[st.Kparentp].MyName
	act.ApSection[ len( act.ApSection )-1 ].Childs = append(act.ApSection[ len( act.ApSection )-1 ].Childs, st)
	act.ApSection[ len( act.ApSection )-1 ].ItsD = append(act.ApSection[ len( act.ApSection )-1 ].ItsD, st)	// artifact.unit:31, g_struct.act:390
	act.ApD = append(act.ApD, st)
	return 0
}

func (me KpD) GetVar(glob *GlobT, va []string, lno string) (bool, string) {
	if (va[0] == "parent") { // artifact.unit:31, g_struct.act:542
		if (me.Kparentp >= 0 && len(va) > 1) {
			return( glob.Dats.ApSection[ me.Kparentp ].GetVar(glob, va[1:], me.LineNo) );
		}
	}
	if va[0] == "previous" { // artifact.unit:40, g_struct.act:171
		if (me.Me > 0 && len(va) > 1) {
			return( glob.Dats.ApD[ me.Me - 1 ].GetVar(glob, va[1:], lno) )
		}
	}
	r,ok := me.Names[va[0]]
	if !ok { r = fmt.Sprintf("?%s?:%s,%s,D > artifact.unit:40, g_struct.act:178?", va[0], lno, me.LineNo) }
	return ok,r
}

func (me KpD) DoIts(glob *GlobT, va []string, lno string) int {
	if va[0] == "parent" { // artifact.unit:31, g_struct.act:527
		if me.Kparentp >= 0 {
			st := glob.Dats.ApSection[ me.Kparentp ]
			if len(va) > 1 {
				return( st.DoIts(glob, va[1:], lno) )
			}
			return( GoAct(glob, st) )
		}
		return(0)
	}
	        fmt.Printf("?No its %s for D %s,%s > artifact.unit:40, g_struct.act:200?", va[0], lno, me.LineNo)
		glob.RunErrs += 1
	return(0)
}

