run="go run ../src/main.go ../src/gen.go ../src/run.go ../src/structs.go  ../src/collect.go ../src/rio.go"


$run  gen-json.act  pbit-maxcut.omni >max-cut.py
if [ $? != 0 ]; then echo gen-json.act  pbit-maxcut.omni has errors; fi

$run  pbit-maxcut.act  pbit-maxcut.omni >max-cut.json
if [ $? != 0 ]; then echo pbit-maxcut.act  pbit-maxcut.omni has errors; fi

