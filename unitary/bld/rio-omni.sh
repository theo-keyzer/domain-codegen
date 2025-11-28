run="go run ../../gen/main.go ../../gen/gen.go ../../gen/run.go ../../gen/structs.go ../../gen/collect.go"


$run go-run-rio.act act.unit,cog.unit,session.unit,canon.unit,gen-high.unit,omni.unit >rio-omni/run.go
if [ $? != 0 ]; then echo go-run-rio.act cog.unit,session.unit,act.unit has errors; fi

$run go-struct-rio.act act.unit,cog.unit,session.unit,canon.unit,gen-high.unit,omni.unit >rio-omni/structs.go
if [ $? != 0 ]; then echo go-struct-rio.act cog.unit,session.unit,act.unit has errors; fi

$run doc.act,check.act cog.unit,session.unit,canon.unit,gen-high.unit,omni.unit  >rio-omni/rio-omni.txt
if [ $? != 0 ]; then echo doc.act,check.act cog.unit,session.unit has errors; fi

$run gen-act.act cog.unit,session.unit,canon.unit,gen-high.unit,omni.unit  >rio-omni/gen-json.act
if [ $? != 0 ]; then echo gen-act.act gen-act.unit has errors; fi


