# Bash tools 

Here you can see names and some tutorials of golang packages which run
in terminals . 
* These packages are so suitable for Makefiles in golang projects .

`gotestsum` is like go test ./... but a utilized with a more preferable log . You can get this
by using below command : 
```bash
go install gotest.tools/gotestsum@latest
```

`goi18n`

`golangci-lint` is used as a lint . Can be received by the command below : 

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin;
```

`goimports`

installing :
```bash
go install golang.org/x/tools/cmd/goimports@latest
```

* before using any of this in your Makefile , check that wether they exist
or not by using 'which' command . for example : which gotestsum 
