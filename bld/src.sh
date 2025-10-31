# The src/* goes to ../src

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go  g_run.act ml.unit,sample.unit,concept.unit,act.unit >src/run.go
if [ $? != 0 ]; then echo g_run.act sample.unit,app.unit,note.unit,act.unit,gen.unit has errors; fi

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go g_struct.act ml.unit,sample.unit,concept.unit,act.unit >src/structs.go
if [ $? != 0 ]; then echo g_struct.act sample.unit,app.unit,note.unit,act.unit,gen.unit has errors; fi


