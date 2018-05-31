LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse master)"

OS := $(shell uname)

build: clean
	go build $(LDFLAGS) -o testsvc

deps:
	dep ensure

test:
	go test -v .

clean:
	go clean
	rm -f test_svc
	rm -f bin/*

install: clean
ifeq ($(OS),Darwin)
	./build.sh darwin
	cp -f bin/test_svc-darwin /usr/local/bin/testsvc
endif 
ifeq ($(OS),Linux)
	./build.sh linux
	cp -f bin/test_svc-linux /usr/local/bin/testsvc
endif
ifeq ($(OS),FreeBSD)
	./build.sh freebsd
	cp -f bin/test_svc-freebsd /usr/local/bin/testsvc
endif
uninstall: 
	rm -f /usr/local/bin/testsvc*

release:
	./build.sh release