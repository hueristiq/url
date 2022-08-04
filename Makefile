TARGET   ?= url
GO       ?= go
GOFLAGS  ?= 

fmt:
	$(GO) $(GOFLAGS) fmt ./...; \
	echo "Done."

test:
	$(GO) $(GOFLAGS) test ./...