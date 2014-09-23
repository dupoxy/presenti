#!/bin/bash -e

# test install and generate doc.
cd ..
lister=`go list`
go test $lister
#go test -coverprofile=coverage.out $lister
go install $lister
misc/doc.sh presenti

# generate Article/Slide for presenti.
presenti -a main.go
presenti -s main.go
