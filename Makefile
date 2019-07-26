.PHONY: test
test:
	ginkgo -r -v

.PHONY: coverage
coverage:
	go test ./...  --cover

.PHONY: deps
deps:
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...
	go get ./...
