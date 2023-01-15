.PHONY: all
all: latc_x86_64 latc

.PHONY: latc_x86_64
latc_x86_64:
	cd src/bin/latc_x86_64; go build -o ../../../latc_x86_64 main.go

.PHONY: latc
latc:
	cd src/bin/latc; go build -o ../../../latc main.go


clean:
	rm latc_x86_64

test: 
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