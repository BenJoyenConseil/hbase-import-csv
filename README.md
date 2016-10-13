# hbase-import-csv
Simple importer from CSV to HBase in Go (POC)

# setup

An easy way to test on a local machine is to use docker :

``docker run --name=hbase -p 2181:2181 -p 60000:60000 -p 60010:60010 -p 60020:60020 -p 60030:60030 -d nerdammer/hbase``

``go run import.go``
