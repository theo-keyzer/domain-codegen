# The src/* goes to ../src

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go  g_runh.act net5.unit,act.unit >srch/run.go
if [ $? != 0 ]; then echo g_runh.act net5.unit,act.unit has errors; fi

go run ../gen/main.go ../gen/gen.go ../gen/run.go ../gen/structs.go  ../gen/collect.go g_structh.act net5.unit,act.unit >srch/structs.go
if [ $? != 0 ]; then echo g_structh.act net5.unit,act.unit has errors; fi


