package main

import (
	"strings"
	"fmt"
//	"strconv"
)

type ActT struct {
	index       map[string]int
	ApComp [] *KpComp
	ApElement [] *KpElement
	ApOpt [] *KpOpt
	ApActor [] *KpActor
	ApAll [] *KpAll
	ApDu [] *KpDu
	ApNew [] *KpNew
	ApRefs [] *KpRefs
	ApVar [] *KpVar
	ApIts [] *KpIts
	ApC [] *KpC
	ApCs [] *KpCs
	ApOut [] *KpOut
	ApIn [] *KpIn
	ApBreak [] *KpBreak
	ApAdd [] *KpAdd
	ApThis [] *KpThis
	ApReplace [] *KpReplace
	ApArtifact [] *KpArtifact
	ApLink [] *KpLink
	ApO [] *KpO
	ApSection [] *KpSection
	ApD [] *KpD
}

func refs(act *ActT) int {
	errs := 0
	v := ""
	//p := -1
	res := 0
	err := false
	for _, st := range act.ApComp {

//  unit.unit:8, g_run.act:169

		v, _ = st.Names["parent"]
		err, res = fnd3(act, "Comp_" + v, v, "ref:Comp.parent:Comp." + v,  ".", st.LineNo, "unit.unit:8, g_run.act:173" );
		st.Kparentp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApElement {

//  unit.unit:35, g_run.act:169

		v, _ = st.Names["comp"]
		err, res = fnd3(act, "Comp_" + v, v, "ref:Element.comp:Comp." + v,  ".", st.LineNo, "unit.unit:35, g_run.act:173" );
		st.Kcompp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApAll {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:34, g_run.act:159" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApDu {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:46, g_run.act:159" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApIts {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:87, g_run.act:159" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThis {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:186, g_run.act:159" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApLink {

//  artifact.unit:21, g_run.act:169

		v, _ = st.Names["concept"]
		err, res = fnd3(act, "Artifact_" + v, v, "ref:Link.concept:Artifact." + v,  "+", st.LineNo, "artifact.unit:21, g_run.act:173" );
		st.Kconceptp = res
		if (err == false) {
			errs += 1
		}
//  artifact.unit:22, g_run.act:169

		v, _ = st.Names["relation"]
		err, res = fnd3(act, "Artifact_" + v, v, "ref:Link.relation:Artifact." + v,  "+", st.LineNo, "artifact.unit:22, g_run.act:173" );
		st.Krelationp = res
		if (err == false) {
			errs += 1
		}
	}
	return(errs)
}

func DoAll(glob *GlobT, va []string, lno string) int {
	if va[0] == "Comp" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Comp_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApComp[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApComp[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApComp {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Element" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Element_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApElement[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApElement[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApElement {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Opt" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Opt_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApOpt[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApOpt[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApOpt {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Actor" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Actor_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApActor[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApActor[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApActor {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Artifact" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Artifact_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApArtifact[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApArtifact[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApArtifact {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Link" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Link_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApLink[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApLink[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApLink {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "O" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["O_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApO[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApO[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApO {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "Section" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Section_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApSection[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApSection[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApSection {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	if va[0] == "D" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["D_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApD[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApD[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApD {
			if len(va) > 2 {
				ret := st.DoIts(glob, va[2:], lno)
				if ret != 0 {
					return(ret)
				}
				continue
			}
			ret := GoAct(glob, st)
			if ret != 0 {
				return(ret)
			}
		}
		return(0)
	}
	fmt.Printf("?No all %s cmd ?%s? > g_run.act:54", va[0], lno);
	return 0;
}

func Load(act *ActT, toks string, ln string, pos int, lno string) int {
	errs := 0
	ss := strings.Split(toks,".")
	tok := ss[0]
	flag := ss[1:]
	if tok == "Comp" { errs += loadComp(act,ln,pos,lno,flag) }
	if tok == "Element" { errs += loadElement(act,ln,pos,lno,flag) }
	if tok == "Opt" { errs += loadOpt(act,ln,pos,lno,flag) }
	if tok == "Actor" { errs += loadActor(act,ln,pos,lno,flag) }
	if tok == "All" { errs += loadAll(act,ln,pos,lno,flag) }
	if tok == "Du" { errs += loadDu(act,ln,pos,lno,flag) }
	if tok == "New" { errs += loadNew(act,ln,pos,lno,flag) }
	if tok == "Refs" { errs += loadRefs(act,ln,pos,lno,flag) }
	if tok == "Var" { errs += loadVar(act,ln,pos,lno,flag) }
	if tok == "Its" { errs += loadIts(act,ln,pos,lno,flag) }
	if tok == "C" { errs += loadC(act,ln,pos,lno,flag) }
	if tok == "Cs" { errs += loadCs(act,ln,pos,lno,flag) }
	if tok == "Out" { errs += loadOut(act,ln,pos,lno,flag) }
	if tok == "In" { errs += loadIn(act,ln,pos,lno,flag) }
	if tok == "Break" { errs += loadBreak(act,ln,pos,lno,flag) }
	if tok == "Add" { errs += loadAdd(act,ln,pos,lno,flag) }
	if tok == "This" { errs += loadThis(act,ln,pos,lno,flag) }
	if tok == "Replace" { errs += loadReplace(act,ln,pos,lno,flag) }
	if tok == "Artifact" { errs += loadArtifact(act,ln,pos,lno,flag) }
	if tok == "Link" { errs += loadLink(act,ln,pos,lno,flag) }
	if tok == "O" { errs += loadO(act,ln,pos,lno,flag) }
	if tok == "Section" { errs += loadSection(act,ln,pos,lno,flag) }
	if tok == "D" { errs += loadD(act,ln,pos,lno,flag) }
	return errs
}

