run="go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go ../gen/collect.go"


$run go-run-rio.act one.unit,act.unit >one/run.go
if [ $? != 0 ]; then echo go-run-rio.act one.unit,act.unit has errors; fi

$run go-struct-rio.act one.unit,act.unit >one/structs.go
if [ $? != 0 ]; then echo go-struct-rio.act one.unit,act.unit has errors; fi

