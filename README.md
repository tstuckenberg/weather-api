# Fibre-API
Fibre API with Golang

First we need to create our Go environment

```
mkdir bin # for go binaries
mkdir pkg # for go external libraries
mkdir src # for your own go code

export PATH=$PATH:[your_path_to]/bin

# Import goimports command
go get golang.org/x/tools/cmd/goimports


cd src ; git clone git@git.vocus.net:gis/fibre-api.git

cd fibre-api ; ls # should return as below:
LICENSE		README.md	apidoc.apib	fibre-api

```

Start working with Go

```
cd fibre-api
go run
```
