# updates in 'test-coverage-badge' block
_COVERAGE_FLOOR=_ 

_CHECK_ERROR=if [ $$? != 0 ]; then exit 1; fi

.PHONY: default \
	install-deps \
	lint-run test-run \
	git-status git-push \
	test-coverage test-coverage-view test-coverage-treemap test-coverage-badge
default: lint-run test-run

### INSTALL

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.0
	go install github.com/nikolaydubina/go-cover-treemap@v1.4.2

### LINT

lint-run:
	go fmt ./...
	go vet ./...
	golangci-lint run -E "gosec,unconvert,gosimple,goconst,gocyclo,err113,ineffassign,unparam,unused,bodyclose,noctx,perfsprint,prealloc,gocritic,govet,revive,staticcheck,errcheck,errorlint,nestif,maintidx"

### TEST

test-run:
	go test -race -cover -count=1 .

### TEST COVERAGE

test-coverage:
	go test -coverpkg=./... -coverprofile=./test/coverage.out -count=1 .
	$(_CHECK_ERROR)

test-coverage-view:
	go tool cover -html=./test/coverage.out

test-coverage-treemap:
	go-cover-treemap -coverprofile=./test/coverage.out > ./test/coverage.svg

test-coverage-badge: 
	$(eval _COVERAGE_FLOOR=go tool cover -func=./test/coverage.out | grep total: | grep -oP '([0-9])+(?=\.[0-9]+)')
	if [ `${_COVERAGE_FLOOR}` -lt 60 ]; then \
		cat ./test/badge_coverage_template.svg | sed -e "s/{{.color}}/dc143c/g;s/{{.percent}}/`${_COVERAGE_FLOOR}`/g" > ./test/badge_coverage.svg; \
	elif [ `${_COVERAGE_FLOOR}` -gt 80 ]; then \
		cat ./test/badge_coverage_template.svg | sed -e "s/{{.color}}/97ca00/g;s/{{.percent}}/`${_COVERAGE_FLOOR}`/g" > ./test/badge_coverage.svg; \
	else \
		cat ./test/badge_coverage_template.svg | sed -e "s/{{.color}}/ff8c00/g;s/{{.percent}}/`${_COVERAGE_FLOOR}`/g" > ./test/badge_coverage.svg; \
	fi

### GIT

git-status: lint-run test-coverage test-coverage-treemap test-coverage-badge
	go fmt ./...
	git add .
	git status 

git-push:
	git commit -m "update"
	git push 
