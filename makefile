.SILENT:
.PHONY: $(MAKECMDGOALS)

example:
	go run example/example.go

test:
	go test -v

vet:
	echo "\n\033[0;36mgo vet is vetting...\033[0m 🤨"
	go vet ./...
	echo "\n\033[0;33mlooks good!\033[0m 🚀\n"

lint:
	echo "\n\033[0;36mgolint is linting...\033[0m 🤨"
	$$GOPATH/bin/golint -set_exit_status .
	echo "\n\033[0;33mlooks good!\033[0m 🚀\n"

fmt:
	gofmt -w -s -d .

# workaround to run godoc with modules - https://github.com/golang/go/issues/26827#issuecomment-451476369
docs:
	mkdir -p /tmp/tmpgoroot/docs
	rm -rf /tmp/tmpgopath/src/github.com/taxjar/taxjar-go
	mkdir -p /tmp/tmpgopath/src/github.com/taxjar/taxjar-go
	tar -c --exclude='.git' --exclude='tmp' . | tar -x -C /tmp/tmpgopath/src/github.com/taxjar/taxjar-go
	echo "\n\033[0;36mopen \033[4;36mhttp://localhost:6060/pkg/github.com/taxjar/taxjar-go \033[0;36m for a preview\n\033[0m"
	GOROOT=/tmp/tmpgoroot/ GOPATH=/tmp/tmpgopath/ godoc -http=localhost:6060 -index -timestamps -play -analysis=type
