..\gencode go -schema=test.schema
..\gencode go -schema=fixed.schema
REM colf -f -b .. go
go test -bench=.