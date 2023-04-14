PACKAGE_LIST := $(shell go list ./...)

urleap:
	go build -o urleap $(PACKAGE_LIST)

test:
	go test $(PACKAGE_LIST)

clean:
	rm -f urleap
