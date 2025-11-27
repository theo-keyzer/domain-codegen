package main

import (
	"strings"
	"fmt"
	"strconv"
)

type ActT struct {
	index       map[string]int
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
	ApAgent [] *KpAgent
	ApObjective [] *KpObjective
	ApMemory [] *KpMemory
	ApMemorySource [] *KpMemorySource
	ApThought [] *KpThought
	ApCitation [] *KpCitation
	ApAction [] *KpAction
	ApThoughtSource [] *KpThoughtSource
	ApTool [] *KpTool
	ApVector [] *KpVector
	ApSession [] *KpSession
	ApSnapshot [] *KpSnapshot
	ApRestoration [] *KpRestoration
	ApCrossRef [] *KpCrossRef
	ApCycle [] *KpCycle
	ApRecommendation [] *KpRecommendation
	ApCanonCycle [] *KpCanonCycle
	ApPattern [] *KpPattern
	ApPatternSource [] *KpPatternSource
	ApEvolution [] *KpEvolution
	ApMetricS [] *KpMetricS
	ApArtifact [] *KpArtifact
	ApGenExecution [] *KpGenExecution
	ApGenValidation [] *KpGenValidation
	ApGenLearning [] *KpGenLearning
	ApGenDebug [] *KpGenDebug
	ApCanon [] *KpCanon
	ApLink [] *KpLink
	ApSection [] *KpSection
	ApCanonMeta [] *KpCanonMeta
	ApImplementation [] *KpImplementation
	ApCanonVersion [] *KpCanonVersion
	ApLibrary [] *KpLibrary
	ApGenProgram [] *KpGenProgram
	ApGenBootstrap [] *KpGenBootstrap
	ApGenInput [] *KpGenInput
	ApGenStrategy [] *KpGenStrategy
	ApGenSearchSpace [] *KpGenSearchSpace
	ApGenPatternRef [] *KpGenPatternRef
	ApGenOutputSpec [] *KpGenOutputSpec
	ApGenMetric [] *KpGenMetric
	ApGenEvolution [] *KpGenEvolution
	ApGenPipeline [] *KpGenPipeline
	ApGenStage [] *KpGenStage
	ApProject [] *KpProject
	ApDomain [] *KpDomain
	ApHardware [] *KpHardware
	ApModel [] *KpModel
	ApLayer [] *KpLayer
	ApBlock [] *KpBlock
	ApTensor [] *KpTensor
	ApOp [] *KpOp
	ApArg [] *KpArg
	ApConfig [] *KpConfig
	ApKernel [] *KpKernel
	ApEnergyFunction [] *KpEnergyFunction
	ApSearchSpace [] *KpSearchSpace
	ApDimension [] *KpDimension
	ApStrategy [] *KpStrategy
	ApConstraint [] *KpConstraint
	ApMetric [] *KpMetric
	ApCheckpoint [] *KpCheckpoint
	ApFusion [] *KpFusion
	ApControlFlow [] *KpControlFlow
}

func refs(act *ActT) int {
	errs := 0
	v := ""
	p := -1
	res := 0
	err := false
	for _, st := range act.ApAll {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:34, go-run-rio.act:170" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApDu {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:46, go-run-rio.act:170" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApIts {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:87, go-run-rio.act:170" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThis {

		err, res = fnd2(act, "Actor_" + st.Kactor, st.Kactor,  ".", st.LineNo, "act.unit:186, go-run-rio.act:170" );
		st.Kactorp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApAgent {

//  cog.unit:17, go-run-rio.act:180

		v, _ = st.Names["tools"].(string)
		err, res = fnd3(act, "Tool_" + v, v, "ref:Agent.tools:Tool." + v,  "*", st.LineNo, "cog.unit:17, go-run-rio.act:184" );
		st.Ktoolsp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApObjective {

//  cog.unit:26, go-run-rio.act:208

		v, _ = st.Names["parent_obj"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Objective_" + v,v, "ref_link:Objective.parent_obj:Agent." + st.Parent + ".Objective." + v,  "*", st.LineNo, "cog.unit:26, go-run-rio.act:211" );
		st.Kparent_objp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemory {

//  cog.unit:46, go-run-rio.act:180

		v, _ = st.Names["embedding"].(string)
		err, res = fnd3(act, "Vector_" + v, v, "ref:Memory.embedding:Vector." + v,  "*", st.LineNo, "cog.unit:46, go-run-rio.act:184" );
		st.Kembeddingp = res
		if (err == false) {
			errs += 1
		}
//  cog.unit:47, go-run-rio.act:208

		v, _ = st.Names["objective"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Objective_" + v,v, "ref_link:Memory.objective:Agent." + st.Parent + ".Objective." + v,  "*", st.LineNo, "cog.unit:47, go-run-rio.act:211" );
		st.Kobjectivep = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemorySource {

//  cog.unit:58, go-run-rio.act:180

		v, _ = st.Names["impl"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:MemorySource.impl:Implementation." + v,  "+", st.LineNo, "cog.unit:58, go-run-rio.act:184" );
		st.Kimplp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThought {

//  cog.unit:69, go-run-rio.act:208

		v, _ = st.Names["prev"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Thought_" + v,v, "ref_link:Thought.prev:Objective." + st.Parent + ".Thought." + v,  "*", st.LineNo, "cog.unit:69, go-run-rio.act:211" );
		st.Kprevp = res
		if (err == false) {
			errs += 1
		}
//  cog.unit:70, go-run-rio.act:208

		v, _ = st.Names["branches"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Thought_" + v,v, "ref_link:Thought.branches:Objective." + st.Parent + ".Thought." + v,  "*", st.LineNo, "cog.unit:70, go-run-rio.act:211" );
		st.Kbranchesp = res
		if (err == false) {
			errs += 1
		}
//  cog.unit:71, go-run-rio.act:208

		v, _ = st.Names["parent_branch"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Thought_" + v,v, "ref_link:Thought.parent_branch:Objective." + st.Parent + ".Thought." + v,  "*", st.LineNo, "cog.unit:71, go-run-rio.act:211" );
		st.Kparent_branchp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApAction {

//  cog.unit:109, go-run-rio.act:180

		v, _ = st.Names["tool"].(string)
		err, res = fnd3(act, "Tool_" + v, v, "ref:Action.tool:Tool." + v,  "+", st.LineNo, "cog.unit:109, go-run-rio.act:184" );
		st.Ktoolp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApThoughtSource {

//  cog.unit:126, go-run-rio.act:180

		v, _ = st.Names["impl"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:ThoughtSource.impl:Implementation." + v,  "+", st.LineNo, "cog.unit:126, go-run-rio.act:184" );
		st.Kimplp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApSession {

//  session.unit:14, go-run-rio.act:180

		v, _ = st.Names["agent_id"].(string)
		err, res = fnd3(act, "Agent_" + v, v, "ref:Session.agent_id:Agent." + v,  "+", st.LineNo, "session.unit:14, go-run-rio.act:184" );
		st.Kagent_idp = res
		if (err == false) {
			errs += 1
		}
//  session.unit:16, go-run-rio.act:180

		v, _ = st.Names["parent_session"].(string)
		err, res = fnd3(act, "Session_" + v, v, "ref:Session.parent_session:Session." + v,  "*", st.LineNo, "session.unit:16, go-run-rio.act:184" );
		st.Kparent_sessionp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApRestoration {

//  session.unit:52, go-run-rio.act:180

		v, _ = st.Names["source_session"].(string)
		err, res = fnd3(act, "Session_" + v, v, "ref:Restoration.source_session:Session." + v,  "+", st.LineNo, "session.unit:52, go-run-rio.act:184" );
		st.Ksource_sessionp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApRecommendation {

//  session.unit:119, go-run-rio.act:180

		v, _ = st.Names["canon"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Recommendation.canon:Canon." + v,  "+", st.LineNo, "session.unit:119, go-run-rio.act:184" );
		st.Kcanonp = res
		if (err == false) {
			errs += 1
		}
//  session.unit:123, go-run-rio.act:180

		v, _ = st.Names["impl"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:Recommendation.impl:Implementation." + v,  "*", st.LineNo, "session.unit:123, go-run-rio.act:184" );
		st.Kimplp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApCanonCycle {

//  session.unit:138, go-run-rio.act:180

		v, _ = st.Names["implementations"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:CanonCycle.implementations:Implementation." + v,  "*", st.LineNo, "session.unit:138, go-run-rio.act:184" );
		st.Kimplementationsp = res
		if (err == false) {
			errs += 1
		}
//  session.unit:139, go-run-rio.act:180

		v, _ = st.Names["baseline_canon"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:CanonCycle.baseline_canon:Canon." + v,  "+", st.LineNo, "session.unit:139, go-run-rio.act:184" );
		st.Kbaseline_canonp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApPattern {

//  session.unit:154, go-run-rio.act:180

		v, _ = st.Names["discovered_in"].(string)
		err, res = fnd3(act, "Session_" + v, v, "ref:Pattern.discovered_in:Session." + v,  "*", st.LineNo, "session.unit:154, go-run-rio.act:184" );
		st.Kdiscovered_inp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApPatternSource {

//  session.unit:167, go-run-rio.act:180

		v, _ = st.Names["impl"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:PatternSource.impl:Implementation." + v,  "*", st.LineNo, "session.unit:167, go-run-rio.act:184" );
		st.Kimplp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApEvolution {

//  session.unit:178, go-run-rio.act:180

		v, _ = st.Names["motivation"].(string)
		err, res = fnd3(act, "Pattern_" + v, v, "ref:Evolution.motivation:Pattern." + v,  "*", st.LineNo, "session.unit:178, go-run-rio.act:184" );
		st.Kmotivationp = res
		if (err == false) {
			errs += 1
		}
//  session.unit:179, go-run-rio.act:180

		v, _ = st.Names["sessions"].(string)
		err, res = fnd3(act, "Session_" + v, v, "ref:Evolution.sessions:Session." + v,  "*", st.LineNo, "session.unit:179, go-run-rio.act:184" );
		st.Ksessionsp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenExecution {

//  session.unit:232, go-run-rio.act:180

		v, _ = st.Names["program"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenExecution.program:GenProgram." + v,  "+", st.LineNo, "session.unit:232, go-run-rio.act:184" );
		st.Kprogramp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenLearning {

//  session.unit:295, go-run-rio.act:180

		v, _ = st.Names["applied_to_program"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenLearning.applied_to_program:GenProgram." + v,  "*", st.LineNo, "session.unit:295, go-run-rio.act:184" );
		st.Kapplied_to_programp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApLink {

//  canon.unit:30, go-run-rio.act:180

		v, _ = st.Names["concept"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Link.concept:Canon." + v,  "+", st.LineNo, "canon.unit:30, go-run-rio.act:184" );
		st.Kconceptp = res
		if (err == false) {
			errs += 1
		}
//  canon.unit:31, go-run-rio.act:180

		v, _ = st.Names["relation"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Link.relation:Canon." + v,  "+", st.LineNo, "canon.unit:31, go-run-rio.act:184" );
		st.Krelationp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApImplementation {

//  canon.unit:68, go-run-rio.act:180

		v, _ = st.Names["canon"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Implementation.canon:Canon." + v,  "+", st.LineNo, "canon.unit:68, go-run-rio.act:184" );
		st.Kcanonp = res
		if (err == false) {
			errs += 1
		}
//  canon.unit:71, go-run-rio.act:180

		v, _ = st.Names["session"].(string)
		err, res = fnd3(act, "Session_" + v, v, "ref:Implementation.session:Session." + v,  "*", st.LineNo, "canon.unit:71, go-run-rio.act:184" );
		st.Ksessionp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApCanonVersion {

//  canon.unit:101, go-run-rio.act:180

		v, _ = st.Names["implementations"].(string)
		err, res = fnd3(act, "Implementation_" + v, v, "ref:CanonVersion.implementations:Implementation." + v,  "*", st.LineNo, "canon.unit:101, go-run-rio.act:184" );
		st.Kimplementationsp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApLibrary {

//  canon.unit:127, go-run-rio.act:180

		v, _ = st.Names["canons"].(string)
		err, res = fnd3(act, "Canon_" + v, v, "ref:Library.canons:Canon." + v,  ".", st.LineNo, "canon.unit:127, go-run-rio.act:184" );
		st.Kcanonsp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenProgram {

//  gen-high.unit:21, go-run-rio.act:180

		v, _ = st.Names["supersedes"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenProgram.supersedes:GenProgram." + v,  "*", st.LineNo, "gen-high.unit:21, go-run-rio.act:184" );
		st.Ksupersedesp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenBootstrap {

//  gen-high.unit:46, go-run-rio.act:180

		v, _ = st.Names["superseded_by"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenBootstrap.superseded_by:GenProgram." + v,  "*", st.LineNo, "gen-high.unit:46, go-run-rio.act:184" );
		st.Ksuperseded_byp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenPatternRef {

//  gen-high.unit:127, go-run-rio.act:180

		v, _ = st.Names["pattern"].(string)
		err, res = fnd3(act, "Pattern_" + v, v, "ref:GenPatternRef.pattern:Pattern." + v,  "+", st.LineNo, "gen-high.unit:127, go-run-rio.act:184" );
		st.Kpatternp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenEvolution {

//  gen-high.unit:190, go-run-rio.act:180

		v, _ = st.Names["previous_version"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenEvolution.previous_version:GenProgram." + v,  "*", st.LineNo, "gen-high.unit:190, go-run-rio.act:184" );
		st.Kprevious_versionp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApGenStage {

//  gen-high.unit:221, go-run-rio.act:180

		v, _ = st.Names["program"].(string)
		err, res = fnd3(act, "GenProgram_" + v, v, "ref:GenStage.program:GenProgram." + v,  "+", st.LineNo, "gen-high.unit:221, go-run-rio.act:184" );
		st.Kprogramp = res
		if (err == false) {
			errs += 1
		}
//  gen-high.unit:225, go-run-rio.act:208

		v, _ = st.Names["parallel_with"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_GenStage_" + v,v, "ref_link:GenStage.parallel_with:GenPipeline." + st.Parent + ".GenStage." + v,  "*", st.LineNo, "gen-high.unit:225, go-run-rio.act:211" );
		st.Kparallel_withp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApProject {

//  omni.unit:13, go-run-rio.act:180

		v, _ = st.Names["domain"].(string)
		err, res = fnd3(act, "Domain_" + v, v, "ref:Project.domain:Domain." + v,  "*", st.LineNo, "omni.unit:13, go-run-rio.act:184" );
		st.Kdomainp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:14, go-run-rio.act:180

		v, _ = st.Names["model"].(string)
		err, res = fnd3(act, "Model_" + v, v, "ref:Project.model:Model." + v,  "+", st.LineNo, "omni.unit:14, go-run-rio.act:184" );
		st.Kmodelp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:15, go-run-rio.act:180

		v, _ = st.Names["strategy"].(string)
		err, res = fnd3(act, "Strategy_" + v, v, "ref:Project.strategy:Strategy." + v,  "*", st.LineNo, "omni.unit:15, go-run-rio.act:184" );
		st.Kstrategyp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:16, go-run-rio.act:180

		v, _ = st.Names["hardware"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Project.hardware:Hardware." + v,  "*", st.LineNo, "omni.unit:16, go-run-rio.act:184" );
		st.Khardwarep = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApHardware {

//  omni.unit:31, go-run-rio.act:180

		v, _ = st.Names["parent_hw"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Hardware.parent_hw:Hardware." + v,  "*", st.LineNo, "omni.unit:31, go-run-rio.act:184" );
		st.Kparent_hwp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:32, go-run-rio.act:180

		v, _ = st.Names["emulation"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Hardware.emulation:Hardware." + v,  "*", st.LineNo, "omni.unit:32, go-run-rio.act:184" );
		st.Kemulationp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:33, go-run-rio.act:180

		v, _ = st.Names["noise_model"].(string)
		err, res = fnd3(act, "Constraint_" + v, v, "ref:Hardware.noise_model:Constraint." + v,  "*", st.LineNo, "omni.unit:33, go-run-rio.act:184" );
		st.Knoise_modelp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApModel {

//  omni.unit:41, go-run-rio.act:180

		v, _ = st.Names["hardware"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Model.hardware:Hardware." + v,  "*", st.LineNo, "omni.unit:41, go-run-rio.act:184" );
		st.Khardwarep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:42, go-run-rio.act:180

		v, _ = st.Names["search_space"].(string)
		err, res = fnd3(act, "SearchSpace_" + v, v, "ref:Model.search_space:SearchSpace." + v,  "*", st.LineNo, "omni.unit:42, go-run-rio.act:184" );
		st.Ksearch_spacep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:43, go-run-rio.act:180

		v, _ = st.Names["config"].(string)
		err, res = fnd3(act, "Config_" + v, v, "ref:Model.config:Config." + v,  "*", st.LineNo, "omni.unit:43, go-run-rio.act:184" );
		st.Kconfigp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApLayer {

//  omni.unit:52, go-run-rio.act:208

		v, _ = st.Names["parent_layer"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Layer_" + v,v, "ref_link:Layer.parent_layer:Model." + st.Parent + ".Layer." + v,  "*", st.LineNo, "omni.unit:52, go-run-rio.act:211" );
		st.Kparent_layerp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApBlock {

//  omni.unit:60, go-run-rio.act:180

		v, _ = st.Names["model"].(string)
		err, res = fnd3(act, "Model_" + v, v, "ref:Block.model:Model." + v,  "*", st.LineNo, "omni.unit:60, go-run-rio.act:184" );
		st.Kmodelp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApTensor {

//  omni.unit:68, go-run-rio.act:208

		v, _ = st.Names["producer"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Op_" + v,v, "ref_link:Tensor.producer:Model." + st.Parent + ".Op." + v,  "*", st.LineNo, "omni.unit:68, go-run-rio.act:211" );
		st.Kproducerp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:69, go-run-rio.act:180

		v, _ = st.Names["distribution"].(string)
		err, res = fnd3(act, "EnergyFunction_" + v, v, "ref:Tensor.distribution:EnergyFunction." + v,  "*", st.LineNo, "omni.unit:69, go-run-rio.act:184" );
		st.Kdistributionp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApOp {

//  omni.unit:79, go-run-rio.act:180

		v, _ = st.Names["hardware"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Op.hardware:Hardware." + v,  "*", st.LineNo, "omni.unit:79, go-run-rio.act:184" );
		st.Khardwarep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:80, go-run-rio.act:180

		v, _ = st.Names["energy_fn"].(string)
		err, res = fnd3(act, "EnergyFunction_" + v, v, "ref:Op.energy_fn:EnergyFunction." + v,  "*", st.LineNo, "omni.unit:80, go-run-rio.act:184" );
		st.Kenergy_fnp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:81, go-run-rio.act:180

		v, _ = st.Names["search_space"].(string)
		err, res = fnd3(act, "SearchSpace_" + v, v, "ref:Op.search_space:SearchSpace." + v,  "*", st.LineNo, "omni.unit:81, go-run-rio.act:184" );
		st.Ksearch_spacep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:82, go-run-rio.act:180

		v, _ = st.Names["strategy"].(string)
		err, res = fnd3(act, "Strategy_" + v, v, "ref:Op.strategy:Strategy." + v,  "*", st.LineNo, "omni.unit:82, go-run-rio.act:184" );
		st.Kstrategyp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:85, go-run-rio.act:208

		v, _ = st.Names["predicate"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Tensor_" + v,v, "ref_link:Op.predicate:Model." + st.Parent + ".Tensor." + v,  "*", st.LineNo, "omni.unit:85, go-run-rio.act:211" );
		st.Kpredicatep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:86, go-run-rio.act:208

		v, _ = st.Names["next_op"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Op_" + v,v, "ref_link:Op.next_op:Model." + st.Parent + ".Op." + v,  "*", st.LineNo, "omni.unit:86, go-run-rio.act:211" );
		st.Knext_opp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:87, go-run-rio.act:208

		v, _ = st.Names["layer"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kparentp) + "_Layer_" + v,v, "ref_link:Op.layer:Model." + st.Parent + ".Layer." + v,  "*", st.LineNo, "omni.unit:87, go-run-rio.act:211" );
		st.Klayerp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApConfig {

//  omni.unit:112, go-run-rio.act:180

		v, _ = st.Names["schedule"].(string)
		err, res = fnd3(act, "Strategy_" + v, v, "ref:Config.schedule:Strategy." + v,  "*", st.LineNo, "omni.unit:112, go-run-rio.act:184" );
		st.Kschedulep = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApKernel {

//  omni.unit:121, go-run-rio.act:180

		v, _ = st.Names["hardware"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Kernel.hardware:Hardware." + v,  "*", st.LineNo, "omni.unit:121, go-run-rio.act:184" );
		st.Khardwarep = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApSearchSpace {

//  omni.unit:137, go-run-rio.act:180

		v, _ = st.Names["target_model"].(string)
		err, res = fnd3(act, "Model_" + v, v, "ref:SearchSpace.target_model:Model." + v,  "*", st.LineNo, "omni.unit:137, go-run-rio.act:184" );
		st.Ktarget_modelp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApStrategy {

//  omni.unit:152, go-run-rio.act:180

		v, _ = st.Names["search_space"].(string)
		err, res = fnd3(act, "SearchSpace_" + v, v, "ref:Strategy.search_space:SearchSpace." + v,  "*", st.LineNo, "omni.unit:152, go-run-rio.act:184" );
		st.Ksearch_spacep = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:153, go-run-rio.act:180

		v, _ = st.Names["fitness"].(string)
		err, res = fnd3(act, "Metric_" + v, v, "ref:Strategy.fitness:Metric." + v,  "*", st.LineNo, "omni.unit:153, go-run-rio.act:184" );
		st.Kfitnessp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApConstraint {

//  omni.unit:162, go-run-rio.act:180

		v, _ = st.Names["target_hw"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Constraint.target_hw:Hardware." + v,  "*", st.LineNo, "omni.unit:162, go-run-rio.act:184" );
		st.Ktarget_hwp = res
		if (err == false) {
			errs += 1
		}
//  omni.unit:163, go-run-rio.act:180

		v, _ = st.Names["model"].(string)
		err, res = fnd3(act, "Model_" + v, v, "ref:Constraint.model:Model." + v,  "*", st.LineNo, "omni.unit:163, go-run-rio.act:184" );
		st.Kmodelp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMetric {

//  omni.unit:172, go-run-rio.act:180

		v, _ = st.Names["target_hw"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Metric.target_hw:Hardware." + v,  "*", st.LineNo, "omni.unit:172, go-run-rio.act:184" );
		st.Ktarget_hwp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApCheckpoint {

//  omni.unit:180, go-run-rio.act:180

		v, _ = st.Names["model"].(string)
		err, res = fnd3(act, "Model_" + v, v, "ref:Checkpoint.model:Model." + v,  "*", st.LineNo, "omni.unit:180, go-run-rio.act:184" );
		st.Kmodelp = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApFusion {

//  omni.unit:188, go-run-rio.act:180

		v, _ = st.Names["hardware"].(string)
		err, res = fnd3(act, "Hardware_" + v, v, "ref:Fusion.hardware:Hardware." + v,  "*", st.LineNo, "omni.unit:188, go-run-rio.act:184" );
		st.Khardwarep = res
		if (err == false) {
			errs += 1
		}
	}
	for _, st := range act.ApMemory {

//  cog.unit:48, go-run-rio.act:235

	if st.Kobjectivep < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Memory.invalidated_by unresolved from link:Memory.objective:Objective %s > cog.unit:48, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApObjective[st.Kobjectivep].MyName
		v, _ = st.Names["invalidated_by"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kobjectivep) + "_Thought_" + v, v, "ref_child:Memory.invalidated_by:Objective." + parent + "." + v + " from link:Memory.objective", "*", st.LineNo, "cog.unit:48, go-run-rio.act:246")
		st.Kinvalidated_byp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApMemorySource {

//  cog.unit:59, go-run-rio.act:258

	if st.Kimplp >= 0 {
		t := st.Kimplp
		st.Kcanonp = act.ApImplementation[t].Kcanonp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:MemorySource.canon unresolved from ref:MemorySource.impl:Implementation.x %s (-) > cog.unit:59, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
	}
	for _, st := range act.ApCitation {

//  cog.unit:90, go-run-rio.act:272
	p = st.Me
	p = act.ApCitation[p].Kparentp
	p = act.ApThought[p].Kparentp
	p = act.ApObjective[p].Kparentp
	if p >= 0 {
		st.Kagentp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Citation.agent unresolved from key:Citation.citation_id:..x %s (-) > cog.unit:90, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  cog.unit:91, go-run-rio.act:235

	if st.Kagentp < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Citation.memory unresolved from up_copy:Citation.agent:Agent %s > cog.unit:91, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagentp].MyName
		v, _ = st.Names["memory"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagentp) + "_Memory_" + v, v, "ref_child:Citation.memory:Agent." + parent + "." + v + " from up_copy:Citation.agent", "+", st.LineNo, "cog.unit:91, go-run-rio.act:246")
		st.Kmemoryp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApAction {

//  cog.unit:111, go-run-rio.act:272
	p = st.Me
	p = act.ApAction[p].Kparentp
	p = act.ApThought[p].Kparentp
	p = act.ApObjective[p].Kparentp
	if p >= 0 {
		st.Kagentp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Action.agent unresolved from text:Action.args:..x %s (-) > cog.unit:111, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  cog.unit:112, go-run-rio.act:235

	if st.Kagentp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Action.result_mem unresolved from up_copy:Action.agent:Agent %s > cog.unit:112, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagentp].MyName
		v, _ = st.Names["result_mem"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagentp) + "_Memory_" + v, v, "ref_child:Action.result_mem:Agent." + parent + "." + v + " from up_copy:Action.agent", "*", st.LineNo, "cog.unit:112, go-run-rio.act:246")
		st.Kresult_memp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApThoughtSource {

//  cog.unit:127, go-run-rio.act:258

	if st.Kimplp >= 0 {
		t := st.Kimplp
		st.Kcanonp = act.ApImplementation[t].Kcanonp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:ThoughtSource.canon unresolved from ref:ThoughtSource.impl:Implementation.x %s (-) > cog.unit:127, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
	}
	for _, st := range act.ApSession {

//  session.unit:15, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApSession[i].Names["comp"] {
			st.Kagent_id2p = act.ApSession[i].Kagent_id2p
			break;
		}
	}

	}
	for _, st := range act.ApSnapshot {

//  session.unit:34, go-run-rio.act:272
	p = st.Me
	p = act.ApSnapshot[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Snapshot.session unresolved from word:Snapshot.trigger:..x %s (-) > session.unit:34, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:35, go-run-rio.act:258

	if st.Ksessionp >= 0 {
		t := st.Ksessionp
		st.Kagent_idp = act.ApSession[t].Kagent_idp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Snapshot.agent_id unresolved from up_copy:Snapshot.session:Session.x %s (-) > session.unit:35, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  session.unit:36, go-run-rio.act:235

	if st.Kagent_idp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Snapshot.objectives unresolved from ref_copy:Snapshot.agent_id:Agent %s > session.unit:36, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_idp].MyName
		v, _ = st.Names["objectives"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_idp) + "_Objective_" + v, v, "ref_child:Snapshot.objectives:Agent." + parent + "." + v + " from ref_copy:Snapshot.agent_id", "*", st.LineNo, "session.unit:36, go-run-rio.act:246")
		st.Kobjectivesp = res
		if !err {
			errs += 1
		}
	}
//  session.unit:37, go-run-rio.act:235

	if st.Kobjectivesp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Snapshot.thoughts unresolved from ref_child:Snapshot.objectives:Objective %s > session.unit:37, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApObjective[st.Kobjectivesp].MyName
		v, _ = st.Names["thoughts"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kobjectivesp) + "_Thought_" + v, v, "ref_child:Snapshot.thoughts:Objective." + parent + "." + v + " from ref_child:Snapshot.objectives", "*", st.LineNo, "session.unit:37, go-run-rio.act:246")
		st.Kthoughtsp = res
		if !err {
			errs += 1
		}
	}
//  session.unit:38, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApSnapshot[i].Names["comp"] {
			st.Kagent_id2p = act.ApSnapshot[i].Kagent_id2p
			break;
		}
	}

//  session.unit:39, go-run-rio.act:235

	if st.Kagent_id2p < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Snapshot.memories unresolved from ref_share:Snapshot.agent_id2:Agent %s > session.unit:39, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_id2p].MyName
		v, _ = st.Names["memories"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_id2p) + "_Memory_" + v, v, "ref_child:Snapshot.memories:Agent." + parent + "." + v + " from ref_share:Snapshot.agent_id2", "*", st.LineNo, "session.unit:39, go-run-rio.act:246")
		st.Kmemoriesp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApRestoration {

//  session.unit:53, go-run-rio.act:258

	if st.Ksource_sessionp >= 0 {
		t := st.Ksource_sessionp
		st.Kagent_idp = act.ApSession[t].Kagent_idp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Restoration.agent_id unresolved from ref:Restoration.source_session:Session.x %s (-) > session.unit:53, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  session.unit:54, go-run-rio.act:235

	if st.Kagent_idp < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Restoration.source_memory unresolved from ref_copy:Restoration.agent_id:Agent %s > session.unit:54, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_idp].MyName
		v, _ = st.Names["source_memory"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_idp) + "_Memory_" + v, v, "ref_child:Restoration.source_memory:Agent." + parent + "." + v + " from ref_copy:Restoration.agent_id", "+", st.LineNo, "session.unit:54, go-run-rio.act:246")
		st.Ksource_memoryp = res
		if !err {
			errs += 1
		}
	}
//  session.unit:55, go-run-rio.act:272
	p = st.Me
	p = act.ApRestoration[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Restoration.session unresolved from ref_child:Restoration.source_memory:Memory.x %s (-) > session.unit:55, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:56, go-run-rio.act:258

	if st.Ksessionp >= 0 {
		t := st.Ksessionp
		st.Kagent_id2p = act.ApSession[t].Kagent_id2p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Restoration.agent_id2 unresolved from up_copy:Restoration.session:Session.x %s (-) > session.unit:56, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  session.unit:57, go-run-rio.act:235

	if st.Kagent_id2p < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Restoration.restored_as unresolved from ref_copy:Restoration.agent_id2:Agent %s > session.unit:57, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_id2p].MyName
		v, _ = st.Names["restored_as"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_id2p) + "_Memory_" + v, v, "ref_child:Restoration.restored_as:Agent." + parent + "." + v + " from ref_copy:Restoration.agent_id2", "+", st.LineNo, "session.unit:57, go-run-rio.act:246")
		st.Krestored_asp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApCycle {

//  session.unit:88, go-run-rio.act:272
	p = st.Me
	p = act.ApCycle[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Cycle.session unresolved from word:Cycle.trigger:..x %s (-) > session.unit:88, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:89, go-run-rio.act:258

	if st.Ksessionp >= 0 {
		t := st.Ksessionp
		st.Kagent_idp = act.ApSession[t].Kagent_idp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Cycle.agent_id unresolved from up_copy:Cycle.session:Session.x %s (-) > session.unit:89, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  session.unit:90, go-run-rio.act:235

	if st.Kagent_idp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Cycle.objective unresolved from ref_copy:Cycle.agent_id:Agent %s > session.unit:90, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_idp].MyName
		v, _ = st.Names["objective"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_idp) + "_Objective_" + v, v, "ref_child:Cycle.objective:Agent." + parent + "." + v + " from ref_copy:Cycle.agent_id", "*", st.LineNo, "session.unit:90, go-run-rio.act:246")
		st.Kobjectivep = res
		if !err {
			errs += 1
		}
	}
//  session.unit:91, go-run-rio.act:235

	if st.Kobjectivep < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Cycle.hypothesis unresolved from ref_child:Cycle.objective:Objective %s > session.unit:91, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApObjective[st.Kobjectivep].MyName
		v, _ = st.Names["hypothesis"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kobjectivep) + "_Thought_" + v, v, "ref_child:Cycle.hypothesis:Objective." + parent + "." + v + " from ref_child:Cycle.objective", "*", st.LineNo, "session.unit:91, go-run-rio.act:246")
		st.Khypothesisp = res
		if !err {
			errs += 1
		}
	}
//  session.unit:92, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApCycle[i].Names["comp"] {
			st.Kagent_id2p = act.ApCycle[i].Kagent_id2p
			break;
		}
	}

//  session.unit:93, go-run-rio.act:235

	if st.Kagent_id2p < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Cycle.baseline unresolved from ref_share:Cycle.agent_id2:Agent %s > session.unit:93, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_id2p].MyName
		v, _ = st.Names["baseline"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_id2p) + "_Memory_" + v, v, "ref_child:Cycle.baseline:Agent." + parent + "." + v + " from ref_share:Cycle.agent_id2", "*", st.LineNo, "session.unit:93, go-run-rio.act:246")
		st.Kbaselinep = res
		if !err {
			errs += 1
		}
	}
//  session.unit:94, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApCycle[i].Names["comp"] {
			st.Kobjective2p = act.ApCycle[i].Kobjective2p
			break;
		}
	}

//  session.unit:95, go-run-rio.act:235

	if st.Kobjective2p < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Cycle.intervention unresolved from ref_share:Cycle.objective2:Objective %s > session.unit:95, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApObjective[st.Kobjective2p].MyName
		v, _ = st.Names["intervention"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kobjective2p) + "_Thought_" + v, v, "ref_child:Cycle.intervention:Objective." + parent + "." + v + " from ref_share:Cycle.objective2", "+", st.LineNo, "session.unit:95, go-run-rio.act:246")
		st.Kinterventionp = res
		if !err {
			errs += 1
		}
	}
//  session.unit:96, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApCycle[i].Names["comp"] {
			st.Kagent_id3p = act.ApCycle[i].Kagent_id3p
			break;
		}
	}

//  session.unit:97, go-run-rio.act:235

	if st.Kagent_id3p < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Cycle.outcome unresolved from ref_share:Cycle.agent_id3:Agent %s > session.unit:97, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_id3p].MyName
		v, _ = st.Names["outcome"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_id3p) + "_Memory_" + v, v, "ref_child:Cycle.outcome:Agent." + parent + "." + v + " from ref_share:Cycle.agent_id3", "+", st.LineNo, "session.unit:97, go-run-rio.act:246")
		st.Koutcomep = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApRecommendation {

//  session.unit:113, go-run-rio.act:272
	p = st.Me
	p = act.ApRecommendation[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Recommendation.session unresolved from key:Recommendation.rec_id:..x %s (-) > session.unit:113, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:114, go-run-rio.act:258

	if st.Ksessionp >= 0 {
		t := st.Ksessionp
		st.Kagent_idp = act.ApSession[t].Kagent_idp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:Recommendation.agent_id unresolved from up_copy:Recommendation.session:Session.x %s (-) > session.unit:114, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  session.unit:115, go-run-rio.act:235

	if st.Kagent_idp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Recommendation.outcome unresolved from ref_copy:Recommendation.agent_id:Agent %s > session.unit:115, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_idp].MyName
		v, _ = st.Names["outcome"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_idp) + "_Memory_" + v, v, "ref_child:Recommendation.outcome:Agent." + parent + "." + v + " from ref_copy:Recommendation.agent_id", "*", st.LineNo, "session.unit:115, go-run-rio.act:246")
		st.Koutcomep = res
		if !err {
			errs += 1
		}
	}
//  session.unit:116, go-run-rio.act:222

	for i := st.Me - 1; i >= 0; i-- {
		if st.Names["comp"] == act.ApRecommendation[i].Names["comp"] {
			st.Kagent_id2p = act.ApRecommendation[i].Kagent_id2p
			break;
		}
	}

//  session.unit:117, go-run-rio.act:235

	if st.Kagent_id2p < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Recommendation.objective unresolved from ref_share:Recommendation.agent_id2:Agent %s > session.unit:117, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApAgent[st.Kagent_id2p].MyName
		v, _ = st.Names["objective"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kagent_id2p) + "_Objective_" + v, v, "ref_child:Recommendation.objective:Agent." + parent + "." + v + " from ref_share:Recommendation.agent_id2", "*", st.LineNo, "session.unit:117, go-run-rio.act:246")
		st.Kobjectivep = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApCanonCycle {

//  session.unit:140, go-run-rio.act:235

	if st.Kbaseline_canonp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:CanonCycle.new_version unresolved from ref:CanonCycle.baseline_canon:Canon %s > session.unit:140, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApCanon[st.Kbaseline_canonp].MyName
		v, _ = st.Names["new_version"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kbaseline_canonp) + "_CanonVersion_" + v, v, "ref_child:CanonCycle.new_version:Canon." + parent + "." + v + " from ref:CanonCycle.baseline_canon", "*", st.LineNo, "session.unit:140, go-run-rio.act:246")
		st.Knew_versionp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApPattern {

//  session.unit:155, go-run-rio.act:235

	if st.Kdiscovered_inp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Pattern.evidence unresolved from ref:Pattern.discovered_in:Session %s > session.unit:155, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApSession[st.Kdiscovered_inp].MyName
		v, _ = st.Names["evidence"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kdiscovered_inp) + "_Cycle_" + v, v, "ref_child:Pattern.evidence:Session." + parent + "." + v + " from ref:Pattern.discovered_in", "*", st.LineNo, "session.unit:155, go-run-rio.act:246")
		st.Kevidencep = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApPatternSource {

//  session.unit:168, go-run-rio.act:258

	if st.Kimplp >= 0 {
		t := st.Kimplp
		st.Kcanonp = act.ApImplementation[t].Kcanonp
	} else if "+" != "*" {
		fmt.Printf("ref_copy:PatternSource.canon unresolved from ref:PatternSource.impl:Implementation.x %s (+) > session.unit:168, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
	}
	for _, st := range act.ApGenExecution {

//  session.unit:233, go-run-rio.act:272
	p = st.Me
	p = act.ApGenExecution[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:GenExecution.session unresolved from ref:GenExecution.program:GenProgram.x %s (-) > session.unit:233, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:234, go-run-rio.act:258

	if st.Ksessionp >= 0 {
		t := st.Ksessionp
		st.Kagent_idp = act.ApSession[t].Kagent_idp
	} else if "-" != "*" {
		fmt.Printf("ref_copy:GenExecution.agent_id unresolved from up_copy:GenExecution.session:Session.x %s (-) > session.unit:234, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
	}
	for _, st := range act.ApGenLearning {

//  session.unit:290, go-run-rio.act:272
	p = st.Me
	p = act.ApGenLearning[p].Kparentp
	p = act.ApGenExecution[p].Kparentp
	if p >= 0 {
		st.Ksessionp = p
	} else if "E_O_L" != "*" {
		fmt.Printf("ref_copy:GenLearning.session unresolved from word:GenLearning.cycle:Cycle.x %s (E_O_L) > session.unit:290, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  session.unit:291, go-run-rio.act:235

	if st.Ksessionp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:GenLearning.baseline unresolved from up_copy:GenLearning.session:Session %s > session.unit:291, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApSession[st.Ksessionp].MyName
		v, _ = st.Names["baseline"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Ksessionp) + "_GenExecution_" + v, v, "ref_child:GenLearning.baseline:Session." + parent + "." + v + " from up_copy:GenLearning.session", "*", st.LineNo, "session.unit:291, go-run-rio.act:246")
		st.Kbaselinep = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApCanonVersion {

//  canon.unit:96, go-run-rio.act:272
	p = st.Me
	p = act.ApCanonVersion[p].Kparentp
	if p >= 0 {
		st.Kcanonp = p
	} else if "-" != "*" {
		fmt.Printf("ref_copy:CanonVersion.canon unresolved from key:CanonVersion.version:..x %s (-) > canon.unit:96, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  canon.unit:97, go-run-rio.act:235

	if st.Kcanonp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:CanonVersion.previous unresolved from up_copy:CanonVersion.canon:Canon %s > canon.unit:97, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApCanon[st.Kcanonp].MyName
		v, _ = st.Names["previous"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kcanonp) + "_CanonVersion_" + v, v, "ref_child:CanonVersion.previous:Canon." + parent + "." + v + " from up_copy:CanonVersion.canon", "*", st.LineNo, "canon.unit:97, go-run-rio.act:246")
		st.Kpreviousp = res
		if !err {
			errs += 1
		}
	}
//  canon.unit:102, go-run-rio.act:258

	if st.Kimplementationsp >= 0 {
		t := st.Kimplementationsp
		st.Ksessionp = act.ApImplementation[t].Ksessionp
	} else if "*" != "*" {
		fmt.Printf("ref_copy:CanonVersion.session unresolved from ref:CanonVersion.implementations:Implementation.x %s (*) > canon.unit:102, go-run-rio.act:264\n", st.LineNo)
		errs += 1
	}
//  canon.unit:103, go-run-rio.act:235

	if st.Ksessionp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:CanonVersion.evidence unresolved from ref_copy:CanonVersion.session:Session %s > canon.unit:103, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApSession[st.Ksessionp].MyName
		v, _ = st.Names["evidence"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Ksessionp) + "_Cycle_" + v, v, "ref_child:CanonVersion.evidence:Session." + parent + "." + v + " from ref_copy:CanonVersion.session", "*", st.LineNo, "canon.unit:103, go-run-rio.act:246")
		st.Kevidencep = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApArg {

//  omni.unit:104, go-run-rio.act:272
	p = st.Me
	p = act.ApArg[p].Kparentp
	p = act.ApOp[p].Kparentp
	if p >= 0 {
		st.Kmodelp = p
	} else if "*" != "*" {
		fmt.Printf("ref_copy:Arg.model unresolved from word:Arg.role:..x %s (*) > omni.unit:104, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  omni.unit:105, go-run-rio.act:235

	if st.Kmodelp < 0 {
		if "+" != "*" {
			fmt.Printf("ref_child:Arg.tensor unresolved from up_copy:Arg.model:Model %s > omni.unit:105, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApModel[st.Kmodelp].MyName
		v, _ = st.Names["tensor"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kmodelp) + "_Tensor_" + v, v, "ref_child:Arg.tensor:Model." + parent + "." + v + " from up_copy:Arg.model", "+", st.LineNo, "omni.unit:105, go-run-rio.act:246")
		st.Ktensorp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApConstraint {

//  omni.unit:164, go-run-rio.act:235

	if st.Kmodelp < 0 {
		if "*" != "*" {
			fmt.Printf("ref_child:Constraint.target_op unresolved from ref:Constraint.model:Model %s > omni.unit:164, go-run-rio.act:239", st.LineNo)
			errs += 1
		}
	} else {
		parent := act.ApModel[st.Kmodelp].MyName
		v, _ = st.Names["target_op"].(string)
		err, res = fnd3(act, strconv.Itoa(st.Kmodelp) + "_Op_" + v, v, "ref_child:Constraint.target_op:Model." + parent + "." + v + " from ref:Constraint.model", "*", st.LineNo, "omni.unit:164, go-run-rio.act:246")
		st.Ktarget_opp = res
		if !err {
			errs += 1
		}
	}
	}
	for _, st := range act.ApControlFlow {

//  omni.unit:196, go-run-rio.act:272
	p = st.Me
	p = act.ApControlFlow[p].Kparentp
	p = act.ApOp[p].Kparentp
	if p >= 0 {
		st.Kmodel1p = p
	} else if "*" != "*" {
		fmt.Printf("ref_copy:ControlFlow.model1 unresolved from key:ControlFlow.control:..x %s (*) > omni.unit:196, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  omni.unit:198, go-run-rio.act:272
	p = st.Me
	p = act.ApControlFlow[p].Kparentp
	p = act.ApOp[p].Kparentp
	if p >= 0 {
		st.Kmodel2p = p
	} else if "*" != "*" {
		fmt.Printf("ref_copy:ControlFlow.model2 unresolved from ref_cild:ControlFlow.predicate:Tensor.x %s (*) > omni.unit:198, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
//  omni.unit:200, go-run-rio.act:272
	p = st.Me
	p = act.ApControlFlow[p].Kparentp
	p = act.ApOp[p].Kparentp
	if p >= 0 {
		st.Kmodel3p = p
	} else if "*" != "*" {
		fmt.Printf("ref_copy:ControlFlow.model3 unresolved from ref_cild:ControlFlow.branch_true:Op.x %s (*) > omni.unit:200, go-run-rio.act:285\n", st.LineNo)
		errs += 1
	}
	}
	return(errs)
}

func DoAll(glob *GlobT, va []string, lno string) int {
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
	if va[0] == "Agent" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Agent_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApAgent[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApAgent[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApAgent {
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
	if va[0] == "MemorySource" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["MemorySource_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApMemorySource[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApMemorySource[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApMemorySource {
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
	if va[0] == "Citation" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Citation_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCitation[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCitation[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCitation {
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
	if va[0] == "Action" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Action_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApAction[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApAction[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApAction {
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
	if va[0] == "ThoughtSource" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["ThoughtSource_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApThoughtSource[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApThoughtSource[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApThoughtSource {
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
	if va[0] == "Tool" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Tool_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApTool[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApTool[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApTool {
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
	if va[0] == "Vector" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Vector_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApVector[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApVector[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApVector {
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
	if va[0] == "Session" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Session_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApSession[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApSession[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApSession {
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
	if va[0] == "Snapshot" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Snapshot_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApSnapshot[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApSnapshot[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApSnapshot {
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
	if va[0] == "Restoration" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Restoration_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApRestoration[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApRestoration[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApRestoration {
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
	if va[0] == "CrossRef" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["CrossRef_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCrossRef[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCrossRef[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCrossRef {
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
	if va[0] == "Cycle" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Cycle_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCycle[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCycle[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCycle {
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
	if va[0] == "Recommendation" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Recommendation_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApRecommendation[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApRecommendation[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApRecommendation {
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
	if va[0] == "CanonCycle" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["CanonCycle_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCanonCycle[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCanonCycle[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCanonCycle {
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
	if va[0] == "Pattern" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Pattern_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApPattern[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApPattern[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApPattern {
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
	if va[0] == "PatternSource" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["PatternSource_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApPatternSource[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApPatternSource[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApPatternSource {
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
	if va[0] == "Evolution" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Evolution_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApEvolution[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApEvolution[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApEvolution {
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
	if va[0] == "MetricS" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["MetricS_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApMetricS[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApMetricS[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApMetricS {
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
	if va[0] == "GenExecution" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenExecution_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenExecution[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenExecution[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenExecution {
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
	if va[0] == "GenValidation" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenValidation_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenValidation[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenValidation[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenValidation {
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
	if va[0] == "GenLearning" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenLearning_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenLearning[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenLearning[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenLearning {
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
	if va[0] == "GenDebug" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenDebug_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenDebug[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenDebug[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenDebug {
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
	if va[0] == "CanonMeta" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["CanonMeta_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCanonMeta[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCanonMeta[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCanonMeta {
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
	if va[0] == "Implementation" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Implementation_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApImplementation[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApImplementation[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApImplementation {
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
	if va[0] == "CanonVersion" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["CanonVersion_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCanonVersion[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCanonVersion[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCanonVersion {
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
	if va[0] == "Library" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Library_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApLibrary[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApLibrary[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApLibrary {
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
	if va[0] == "GenProgram" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenProgram_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenProgram[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenProgram[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenProgram {
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
	if va[0] == "GenBootstrap" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenBootstrap_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenBootstrap[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenBootstrap[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenBootstrap {
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
	if va[0] == "GenInput" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenInput_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenInput[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenInput[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenInput {
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
	if va[0] == "GenStrategy" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenStrategy_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenStrategy[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenStrategy[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenStrategy {
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
	if va[0] == "GenSearchSpace" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenSearchSpace_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenSearchSpace[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenSearchSpace[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenSearchSpace {
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
	if va[0] == "GenPatternRef" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenPatternRef_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenPatternRef[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenPatternRef[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenPatternRef {
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
	if va[0] == "GenOutputSpec" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenOutputSpec_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenOutputSpec[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenOutputSpec[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenOutputSpec {
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
	if va[0] == "GenMetric" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenMetric_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenMetric[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenMetric[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenMetric {
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
	if va[0] == "GenEvolution" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenEvolution_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenEvolution[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenEvolution[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenEvolution {
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
	if va[0] == "GenPipeline" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenPipeline_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenPipeline[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenPipeline[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenPipeline {
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
	if va[0] == "GenStage" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["GenStage_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApGenStage[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApGenStage[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApGenStage {
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
	if va[0] == "Project" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Project_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApProject[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApProject[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApProject {
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
	if va[0] == "Domain" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Domain_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApDomain[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApDomain[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApDomain {
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
	if va[0] == "Hardware" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Hardware_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApHardware[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApHardware[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApHardware {
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
	if va[0] == "Model" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Model_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApModel[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApModel[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApModel {
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
	if va[0] == "Layer" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Layer_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApLayer[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApLayer[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApLayer {
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
	if va[0] == "Block" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Block_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApBlock[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApBlock[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApBlock {
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
	if va[0] == "Tensor" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Tensor_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApTensor[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApTensor[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApTensor {
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
	if va[0] == "Op" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Op_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApOp[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApOp[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApOp {
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
	if va[0] == "Arg" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Arg_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApArg[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApArg[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApArg {
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
	if va[0] == "Config" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Config_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApConfig[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApConfig[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApConfig {
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
	if va[0] == "Kernel" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Kernel_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApKernel[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApKernel[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApKernel {
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
	if va[0] == "EnergyFunction" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["EnergyFunction_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApEnergyFunction[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApEnergyFunction[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApEnergyFunction {
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
	if va[0] == "SearchSpace" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["SearchSpace_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApSearchSpace[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApSearchSpace[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApSearchSpace {
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
	if va[0] == "Dimension" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Dimension_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApDimension[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApDimension[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApDimension {
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
	if va[0] == "Strategy" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Strategy_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApStrategy[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApStrategy[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApStrategy {
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
	if va[0] == "Constraint" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Constraint_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApConstraint[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApConstraint[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApConstraint {
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
	if va[0] == "Metric" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Metric_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApMetric[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApMetric[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApMetric {
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
	if va[0] == "Checkpoint" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Checkpoint_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApCheckpoint[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApCheckpoint[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApCheckpoint {
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
	if va[0] == "Fusion" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["Fusion_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApFusion[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApFusion[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApFusion {
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
	if va[0] == "ControlFlow" {
		if (len(va) > 1 && len(va[1]) > 0) {
			en, er := glob.Dats.index["ControlFlow_" + va[1] ];
			if !er {
				if len(va) > 2 {
					return( glob.Dats.ApControlFlow[en].DoIts(glob, va[2:], lno) )
				}
				return( GoAct(glob, glob.Dats.ApControlFlow[en]) )
			}
			return(0)
		}
		for _, st := range glob.Dats.ApControlFlow {
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
	fmt.Printf("?No all %s cmd ?%s? > Command line arguments, go-run-rio.act:43", va[0], lno);
	return 0;
}

func Load(act *ActT, toks string, ln string, pos int, lno string) int {
	errs := 0
	ss := strings.Split(toks,".")
	tok := ss[0]
	flag := ss[1:]
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
	return errs
}

func Loadh(act *ActT, toks string, ln string, pos int, lno string, nm map[string]any) int {
	errs := 0
	ss := strings.Split(toks,".")
	tok := ss[0]
	flag := ss[1:]
	if tok == "Agent" { errs += loadAgent(act,ln,pos,lno,flag,nm) }
	if tok == "Objective" { errs += loadObjective(act,ln,pos,lno,flag,nm) }
	if tok == "Memory" { errs += loadMemory(act,ln,pos,lno,flag,nm) }
	if tok == "MemorySource" { errs += loadMemorySource(act,ln,pos,lno,flag,nm) }
	if tok == "Thought" { errs += loadThought(act,ln,pos,lno,flag,nm) }
	if tok == "Citation" { errs += loadCitation(act,ln,pos,lno,flag,nm) }
	if tok == "Action" { errs += loadAction(act,ln,pos,lno,flag,nm) }
	if tok == "ThoughtSource" { errs += loadThoughtSource(act,ln,pos,lno,flag,nm) }
	if tok == "Tool" { errs += loadTool(act,ln,pos,lno,flag,nm) }
	if tok == "Vector" { errs += loadVector(act,ln,pos,lno,flag,nm) }
	if tok == "Session" { errs += loadSession(act,ln,pos,lno,flag,nm) }
	if tok == "Snapshot" { errs += loadSnapshot(act,ln,pos,lno,flag,nm) }
	if tok == "Restoration" { errs += loadRestoration(act,ln,pos,lno,flag,nm) }
	if tok == "CrossRef" { errs += loadCrossRef(act,ln,pos,lno,flag,nm) }
	if tok == "Cycle" { errs += loadCycle(act,ln,pos,lno,flag,nm) }
	if tok == "Recommendation" { errs += loadRecommendation(act,ln,pos,lno,flag,nm) }
	if tok == "CanonCycle" { errs += loadCanonCycle(act,ln,pos,lno,flag,nm) }
	if tok == "Pattern" { errs += loadPattern(act,ln,pos,lno,flag,nm) }
	if tok == "PatternSource" { errs += loadPatternSource(act,ln,pos,lno,flag,nm) }
	if tok == "Evolution" { errs += loadEvolution(act,ln,pos,lno,flag,nm) }
	if tok == "MetricS" { errs += loadMetricS(act,ln,pos,lno,flag,nm) }
	if tok == "Artifact" { errs += loadArtifact(act,ln,pos,lno,flag,nm) }
	if tok == "GenExecution" { errs += loadGenExecution(act,ln,pos,lno,flag,nm) }
	if tok == "GenValidation" { errs += loadGenValidation(act,ln,pos,lno,flag,nm) }
	if tok == "GenLearning" { errs += loadGenLearning(act,ln,pos,lno,flag,nm) }
	if tok == "GenDebug" { errs += loadGenDebug(act,ln,pos,lno,flag,nm) }
	if tok == "Canon" { errs += loadCanon(act,ln,pos,lno,flag,nm) }
	if tok == "Link" { errs += loadLink(act,ln,pos,lno,flag,nm) }
	if tok == "Section" { errs += loadSection(act,ln,pos,lno,flag,nm) }
	if tok == "CanonMeta" { errs += loadCanonMeta(act,ln,pos,lno,flag,nm) }
	if tok == "Implementation" { errs += loadImplementation(act,ln,pos,lno,flag,nm) }
	if tok == "CanonVersion" { errs += loadCanonVersion(act,ln,pos,lno,flag,nm) }
	if tok == "Library" { errs += loadLibrary(act,ln,pos,lno,flag,nm) }
	if tok == "GenProgram" { errs += loadGenProgram(act,ln,pos,lno,flag,nm) }
	if tok == "GenBootstrap" { errs += loadGenBootstrap(act,ln,pos,lno,flag,nm) }
	if tok == "GenInput" { errs += loadGenInput(act,ln,pos,lno,flag,nm) }
	if tok == "GenStrategy" { errs += loadGenStrategy(act,ln,pos,lno,flag,nm) }
	if tok == "GenSearchSpace" { errs += loadGenSearchSpace(act,ln,pos,lno,flag,nm) }
	if tok == "GenPatternRef" { errs += loadGenPatternRef(act,ln,pos,lno,flag,nm) }
	if tok == "GenOutputSpec" { errs += loadGenOutputSpec(act,ln,pos,lno,flag,nm) }
	if tok == "GenMetric" { errs += loadGenMetric(act,ln,pos,lno,flag,nm) }
	if tok == "GenEvolution" { errs += loadGenEvolution(act,ln,pos,lno,flag,nm) }
	if tok == "GenPipeline" { errs += loadGenPipeline(act,ln,pos,lno,flag,nm) }
	if tok == "GenStage" { errs += loadGenStage(act,ln,pos,lno,flag,nm) }
	if tok == "Project" { errs += loadProject(act,ln,pos,lno,flag,nm) }
	if tok == "Domain" { errs += loadDomain(act,ln,pos,lno,flag,nm) }
	if tok == "Hardware" { errs += loadHardware(act,ln,pos,lno,flag,nm) }
	if tok == "Model" { errs += loadModel(act,ln,pos,lno,flag,nm) }
	if tok == "Layer" { errs += loadLayer(act,ln,pos,lno,flag,nm) }
	if tok == "Block" { errs += loadBlock(act,ln,pos,lno,flag,nm) }
	if tok == "Tensor" { errs += loadTensor(act,ln,pos,lno,flag,nm) }
	if tok == "Op" { errs += loadOp(act,ln,pos,lno,flag,nm) }
	if tok == "Arg" { errs += loadArg(act,ln,pos,lno,flag,nm) }
	if tok == "Config" { errs += loadConfig(act,ln,pos,lno,flag,nm) }
	if tok == "Kernel" { errs += loadKernel(act,ln,pos,lno,flag,nm) }
	if tok == "EnergyFunction" { errs += loadEnergyFunction(act,ln,pos,lno,flag,nm) }
	if tok == "SearchSpace" { errs += loadSearchSpace(act,ln,pos,lno,flag,nm) }
	if tok == "Dimension" { errs += loadDimension(act,ln,pos,lno,flag,nm) }
	if tok == "Strategy" { errs += loadStrategy(act,ln,pos,lno,flag,nm) }
	if tok == "Constraint" { errs += loadConstraint(act,ln,pos,lno,flag,nm) }
	if tok == "Metric" { errs += loadMetric(act,ln,pos,lno,flag,nm) }
	if tok == "Checkpoint" { errs += loadCheckpoint(act,ln,pos,lno,flag,nm) }
	if tok == "Fusion" { errs += loadFusion(act,ln,pos,lno,flag,nm) }
	if tok == "ControlFlow" { errs += loadControlFlow(act,ln,pos,lno,flag,nm) }
	return errs
}

