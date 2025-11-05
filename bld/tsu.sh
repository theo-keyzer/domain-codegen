# The src/* goes to ../src

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go  g_runh.act tsu.unit,tsu-auto.unit,tsu-ext.unit,act.unit >tsu/run.go
if [ $? != 0 ]; then echo g_runh.act tsu.unit,tsu-auto.unit,tsu-ext.unit,act.unit has errors; fi

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go g_structh.act tsu.unit,tsu-auto.unit,tsu-ext.unit,act.unit >tsu/structs.go
if [ $? != 0 ]; then echo g_structh.act tsu.unit,tsu-auto.unit,tsu-ext.unit,act.unit has errors; fi


