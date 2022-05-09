GOPATH	:= $(shell go env GOPATH)
GOPATH1	:= $(firstword $(subst :, ,$(GOPATH)))

# TYPE_MAPPINGS explained:
# `integer` - 1st override added.  https://swagger.io/specification/v2/ doesn't provide an unsigned int representation.
# `number` -  After specifying `integer` override, we discovered a need for a signed integer.  The kludge repurposes `number` because it's an unused data type.
TYPE_MAPPINGS = -type-mappings integer=uint64,number=int64

# `make all` or just `make` should be appropriate for dev work
all:	server/v2/generated/types.go server/v2/generated/routes.go server/v2/generated/private/types.go server/v2/generated/private/routes.go

# `make generate` should be able to replace old `generate.sh` script and be appropriate for build system use
generate:	oapi-codegen all

server/v2/generated/types.go:	algod.oas3.yml
	$(GOPATH1)/bin/oapi-codegen -package generated ${TYPE_MAPPINGS} -generate types -exclude-tags=private,common -o ./server/v2/generated/types.go algod.oas3.yml

server/v2/generated/routes.go:	algod.oas3.yml
	$(GOPATH1)/bin/oapi-codegen -package generated ${TYPE_MAPPINGS} -generate server,spec -exclude-tags=private,common -o ./server/v2/generated/routes.go algod.oas3.yml

server/v2/generated/private/types.go:	algod.oas3.yml
	$(GOPATH1)/bin/oapi-codegen -package private ${TYPE_MAPPINGS} -generate types -include-tags=private -o ./server/v2/generated/private/types.go algod.oas3.yml

server/v2/generated/private/routes.go:	algod.oas3.yml
	$(GOPATH1)/bin/oapi-codegen -package private ${TYPE_MAPPINGS} -generate server,spec -include-tags=private -o ./server/v2/generated/private/routes.go algod.oas3.yml

algod.oas3.yml:	algod.oas2.json
	curl -s -X POST "https://converter.swagger.io/api/convert" -H "accept: application/json" -H "Content-Type: application/json" -d @./algod.oas2.json -o .3tmp.json
	python3 jsoncanon.py < .3tmp.json > algod.oas3.yml
	rm -f .3tmp.json

oapi-codegen:	.PHONY
	../../../scripts/buildtools/install_buildtools.sh -o github.com/algorand/oapi-codegen -c github.com/algorand/oapi-codegen/cmd/oapi-codegen

clean:
	rm -rf server/v2/generated/types.go server/v2/generated/routes.go server/v2/generated/private/types.go server/v2/generated/private/routes.go algod.oas3.yml

.PHONY:
