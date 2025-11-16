run="/home/theo/Downloads/go1.24.5.linux-amd64/go/bin/go run main.go gen.go run.go structs.go  collect.go hcon4.go"

$run  doc.act unit.unit,artifact.unit,act.unit > gen_doc.txt
if [ $? != 0 ]; then echo doc.act unit.unit,artifact.unit,act.unit has errors; fi

$run  doc.act tsu.unit,tsu-auto.unit,tsu-ext.unit > tsu-doc.txt
if [ $? != 0 ]; then echo doc.act tsu.unit,tsu-auto.unit,tsu-ext.unit has errors; fi

$run  doc.act net5.unit > ml-doc.txt
if [ $? != 0 ]; then echo doc.act net5.unit has errors; fi

$run  doc.act nexus.unit > nexus-doc.txt
if [ $? != 0 ]; then echo doc.act nexus.unit has errors; fi


