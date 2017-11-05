.PHONY: mkbuilddir buildbot runtests clean

mkbuilddir:
	mkdir -p build

buildbot: mkbuilddir
	go build -o build/roger-bot ./examples/

runtests:
	go test ./...

gettestcoverage:
	go test -cover ./...

clean:
	rm -rf ./build
