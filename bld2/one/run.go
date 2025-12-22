package main

import (
	"strings"
	"fmt"
	"strconv"
)

type ActT struct {
	index       map[string]int
	ApCanon [] *KpCanon
	ApLink [] *KpLink
	ApSection [] *KpSection
	ApObjective [] *KpObjective
	ApMemory [] *KpMemory
	ApThought [] *KpThought
	ApMemo [] *KpMemo
	ApOps [] *KpOps
	ApComp [] *KpComp
	ApElement [] *KpElement
	ApOpt [] *KpOpt
	ApFunc [] *KpFunc
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
}

func refs(act *ActT) int {
	errs := 0
	v := ""
	p := -1
	res := 0
	err := false
	for _, st := range act.ApLink {

//  one.unit:25, go-run-rio.act:179

		v, _ = st.Names["concept"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Link.concept:Canon." + v,  "+", st.LineNo, "one.unit:25, go-run-rio.act:183" );
		st.Kconceptp = res
		if (err == false) {
			errs += 1
		}
//  one.unit:26, go-run-rio.act:179

		v, _ = st.Names["relation"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Link.relation:Canon." + v,  "+", st.LineNo, "one.unit:26, go-run-rio.act:183" );
		st.Krelationp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApObjective {

//  one.unit:44, go-run-rio.act:179

		v, _ = st.Names["canonical_form"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Objective.canonical_form:Canon." + v,  "*", st.LineNo, "one.unit:44, go-run-rio.act:183" );
		st.Kcanonical_formp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemory {

//  one.unit:67, go-run-rio.act:207

		v, _ = st.Names["invalidated_by"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Thought_" + v,v, "ref_link:Memory.invalidated_by:Objective." + st.Parent + ".Thought." + v,  "*", st.LineNo, "one.unit:67, go-run-rio.act:210" );
		st.Kinvalidated_byp = res
		if (err == false) {
			errs += 1
		}
//  one.unit:68, go-run-rio.act:179

		v, _ = st.Names["canonical_form"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Memory.canonical_form:Canon." + v,  "*", st.LineNo, "one.unit:68, go-run-rio.act:183" );
		st.Kcanonical_formp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThought {

//  one.unit:95, go-run-rio.act:179

		v, _ = st.Names["canonical_form"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Thought.canonical_form:Canon." + v,  "*", st.LineNo, "one.unit:95, go-run-rio.act:183" );
		st.Kcanonical_formp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemo {

//  one.unit:112, go-run-rio.act:179

		v, _ = st.Names["canonical_form"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Memo.canonical_form:Canon." + v,  "*", st.LineNo, "one.unit:112, go-run-rio.act:183" );
		st.Kcanonical_formp = res
		if (err == false) {
			errs += 1
		}
//  one.unit:115, go-run-rio.act:179

		v, _ = st.Names["computation"].(string)
		err, res = fnd3(act, "Ops_" + v, v, "ref:Memo.computation:Ops." + v,  "*", st.LineNo, "one.unit:115, go-run-rio.act:183" );
		st.Kcomputationp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApComp {

//  one.unit:140, go-run-rio.act:179

		v, _ = st.Names["parent"].(string)
		err, res = fnd3(act, "Comp_" + v, v, "ref:Comp.parent:Comp." + v,  ".", st.LineNo, "one.unit:140, go-run-rio.act:183" );
		st.Kparentp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApElement {

//  one.unit:168, go-run-rio.act:179

		v, _ = st.Names["comp"].(string)
		err, res = fnd3(act, "Comp_" + v, v, "ref:Element.comp:Comp." + v,  ".", st.LineNo, "one.unit:168, go-run-rio.act:183" );
		st.Kcompp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApAll {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:49, go-run-rio.act:169" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApDu {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:61, go-run-rio.act:169" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApIts {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:102, go-run-rio.act:169" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThis {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:202, go-run-rio.act:169" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemo {

//  one.unit:113, go-run-rio.act:271
	p = st.Me
	p = act.ApMemo[p].Kparentp
	p = act.ApThought[p].Kparentp
	if p >= 0 {
		st.Kobjective_idp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Memo.objective_id unresolved from ref:Memo.canonical_form:Canon.x %s (-) > one.unit:113, go-run-rio.act:284\n", st.LineNo)
		errs += 1
	}
//  one.unit:114, go-run-rio.act:234

	if st.Kobjective_idp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Memo.memory_id unresolved from up_copy:Memo.objective_id:Objective %s > one.unit:114, go-run-rio.act:238", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApObjective[st.Kobjective_idp].MyName
		v, _ = st.Names["memory_id"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kobjective_idp) + "_Memory_" + v, v, "ref_child:Memo.memory_id:Objective." + parent + "." + v + " from up_copy:Memo.objective_id", "*", st.LineNo, "one.unit:114, go-run-rio.act:245")
		st.Kmemory_idp = res
		if !err {
			errs += 1
		}
	}
	}
	return(errs)
}

func DoAll(glob *GlobT, va []string, lno string) int {
	if va[0] == "Canon" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Canon_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCanon[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCanon[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCanon {
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
	if va[0] == "Objective" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Objective_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApObjective[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApObjective[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApObjective {
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
	if va[0] == "Memory" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Memory_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApMemory[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApMemory[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApMemory {
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
	if va[0] == "Thought" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Thought_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApThought[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApThought[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApThought {
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
	if va[0] == "Memo" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Memo_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApMemo[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApMemo[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApMemo {
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
	if va[0] == "Ops" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Ops_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApOps[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApOps[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApOps {
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
	if va[0] == "Func" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Func_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApFunc[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApFunc[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApFunc {
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
	if va[0] == "All" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["All_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApAll[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApAll[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApAll {
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
	if va[0] == "Du" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Du_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApDu[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApDu[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApDu {
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
	if va[0] == "New" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["New_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApNew[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApNew[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApNew {
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
	if va[0] == "Refs" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Refs_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApRefs[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApRefs[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApRefs {
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
	if va[0] == "Var" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Var_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApVar[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApVar[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApVar {
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
	if va[0] == "Its" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Its_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApIts[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApIts[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApIts {
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
	if va[0] == "C" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["C_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApC[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApC[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApC {
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
	if va[0] == "Cs" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Cs_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCs[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCs[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCs {
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
	if va[0] == "Out" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Out_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApOut[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApOut[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApOut {
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
	if va[0] == "In" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["In_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApIn[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApIn[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApIn {
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
	if va[0] == "Break" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Break_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApBreak[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApBreak[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApBreak {
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
	if va[0] == "Add" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Add_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApAdd[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApAdd[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApAdd {
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
	if va[0] == "This" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["This_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApThis[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApThis[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApThis {
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
	if va[0] == "Replace" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Replace_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApReplace[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApReplace[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApReplace {
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
	fmt.Printf("?No all %s cmd ?%s? > Command line arguments, go-run-rio.act:42", va[0], lno);
	return 0;
}

func Load(act *ActT, toks string, ln string, pos int, lno string) int {
	errs := 0
	return errs
}

func Loadh(act *ActT, toks string, ln string, pos int, lno string, nm map[string]any) int {
	errs := 0
	ss := strings.Split(toks,".")
	tok := ss[0]
	flag := ss[1:]
	if tok == "Canon" { errs += loadCanon(act,ln,pos,lno,flag,nm) }
	if tok == "Link" { errs += loadLink(act,ln,pos,lno,flag,nm) }
	if tok == "Section" { errs += loadSection(act,ln,pos,lno,flag,nm) }
	if tok == "Objective" { errs += loadObjective(act,ln,pos,lno,flag,nm) }
	if tok == "Memory" { errs += loadMemory(act,ln,pos,lno,flag,nm) }
	if tok == "Thought" { errs += loadThought(act,ln,pos,lno,flag,nm) }
	if tok == "Memo" { errs += loadMemo(act,ln,pos,lno,flag,nm) }
	if tok == "Ops" { errs += loadOps(act,ln,pos,lno,flag,nm) }
	if tok == "Comp" { errs += loadComp(act,ln,pos,lno,flag,nm) }
	if tok == "Element" { errs += loadElement(act,ln,pos,lno,flag,nm) }
	if tok == "Opt" { errs += loadOpt(act,ln,pos,lno,flag,nm) }
	if tok == "Func" { errs += loadFunc(act,ln,pos,lno,flag,nm) }
	if tok == "Actor" { errs += loadActor(act,ln,pos,lno,flag,nm) }
	if tok == "All" { errs += loadAll(act,ln,pos,lno,flag,nm) }
	if tok == "Du" { errs += loadDu(act,ln,pos,lno,flag,nm) }
	if tok == "New" { errs += loadNew(act,ln,pos,lno,flag,nm) }
	if tok == "Refs" { errs += loadRefs(act,ln,pos,lno,flag,nm) }
	if tok == "Var" { errs += loadVar(act,ln,pos,lno,flag,nm) }
	if tok == "Its" { errs += loadIts(act,ln,pos,lno,flag,nm) }
	if tok == "C" { errs += loadC(act,ln,pos,lno,flag,nm) }
	if tok == "Cs" { errs += loadCs(act,ln,pos,lno,flag,nm) }
	if tok == "Out" { errs += loadOut(act,ln,pos,lno,flag,nm) }
	if tok == "In" { errs += loadIn(act,ln,pos,lno,flag,nm) }
	if tok == "Break" { errs += loadBreak(act,ln,pos,lno,flag,nm) }
	if tok == "Add" { errs += loadAdd(act,ln,pos,lno,flag,nm) }
	if tok == "This" { errs += loadThis(act,ln,pos,lno,flag,nm) }
	if tok == "Replace" { errs += loadReplace(act,ln,pos,lno,flag,nm) }
	return errs
}

