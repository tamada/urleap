PACKAGE_LIST := $(shell go list ./...)

urleap: test
	go build -o urleap $(PACKAGE_LIST)

test:
	go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)

clean:
	rm -f urleap
