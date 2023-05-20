# go mod tidy
.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor

# go mod vendor
.PHONY: vendor
vendor:
	go mod vendor

# gRPC generate
.PHONY: bufgen
bufgen:
	buf generate

# make buflint
.PHONY: buflint
buflint:
	buf lint
