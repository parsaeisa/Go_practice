# Bash tools 

Here you can see names and some tutorials of golang packages which run
in terminals . 
* These packages are so suitable for Makefiles in golang projects .

## Go commands

Creating vendor folder based on go.mod:
```bash
go mod vendor -v
```
After this command, GoLand starts indexing.

### gotestsum
`gotestsum` is like go test ./... but is utilized with a more preferable log . You can get this
by using below command : 
```bash
go install gotest.tools/gotestsum@latest
```

gotestsum features : 
- computing the coverage 
- running tests

`goi18n`

### golangci-lint
`golangci-lint` is used as a lint . Can be received by the command below : 

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin;
```
golang ci-lint takes a yml file as input. usage : 

```bash
golangci-lint run {path to yaml file} ./...
```
apparently `golint` is deprecated now and revive is used instead. 

This is a yml file that can be passed to the golang-ci lint : https://github.com/n25a/Boiler/blob/master/golang/golangci-lint.yml

### goimports
`goimports`

installing :
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

> before using any of these in your Makefile , check that wether they exist
or not by using 'which' command . for example : which gotestsum 

## Package managing
Installing project's packages :
```go
go mod vendor -v
```
the goproxy must be on . 

The bash command below matches the modules to the source code. 
```
go mod tidy
```

go mod tidy works by loading all of the packages in the main module and all of the packages they import, recursively. 

