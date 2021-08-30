#!/usr/bin/make -f

test:
	@go version
	go fmt ./...
	go mod tidy
	@echo
	go test        -cover -timeout=1s -race ./...
	@echo
	go test -short -cover -timeout=1s -race ./...

doc:
	printf '# ' > README.md && \
		head -n 1 go.mod | sed 's/^module //' >> README.md && \
		echo >> README.md && \
		echo >> README.md && \
		go doc -all suite  | sed 's/^/\t/' >> README.md && \
		echo '---' >> README.md && \
		echo >> README.md && \
		go doc -all should | sed 's/^/\t/' >> README.md && \
		echo >> README.md
