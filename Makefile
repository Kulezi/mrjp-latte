.PHONY: all
all: latc_x86 latc

.PHONY: latc_x86
latc_x86:
	cd src/cmd/latc_x86; go build -o ../../../latc_x86 main.go

.PHONY: latc
latc:
	cd src/cmd/latc; go build -o ../../../latc main.go

clean:
	rm latc_x86 latc

test-typecheck:
	@cd src/compiler; go test . -run TestTypecheck -count=1

test: 
	@cd src/compiler; go test . -count=1

test-verbose: 
	@cd src/compiler; go test . -count=1 -v


.PHONY: cover
PACKAGES = $(shell cat testpackages.txt)
cover: 
	@cd src/compiler; $(foreach pkg,$(PACKAGES),\
		go test . -count=1 \
		-coverprofile=../../profile-$(shell basename $(pkg)).cov \
		-coverpkg=./$(pkg) \
		-covermode=count; \
		go tool cover -html=../../profile-$(shell basename $(pkg)).cov &)