GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

VRF_ABI_ARTIFACT := ./abis/VRFCoodinatorV2.sol/VRFCoordinatorV2.json


my-vrf-oracle:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/my-vrf-oracle

clean:
	rm cmd/my-vrf-oracle

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: binding-vrf


binding-vrf:
	$(eval temp := $(shell mktemp))

	cat $(VRF_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(VRF_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/VRFCoordinatorV2.go \
		--type VRFCoordinatorV2 \
		--bin $(temp)

		rm $(temp)


.PHONY: \
	my-vrf-oracle \
	bindings \
	binding-vrf \
	clean \
	test \
	lint