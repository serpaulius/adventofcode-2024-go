```
# enable profiling in main.go
go run main.go

# show performance metrics overview
go tool pprof -http=:8080 myprogram.prof

# NOTE: on Mac may need to install these before running pprof
# brew install svn
# brew install graphviz
```
