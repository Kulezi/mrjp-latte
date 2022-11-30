.PHONY: all
all: latc

.PHONY: latc
latc:
	cd src; go build -o ../latc main.go

clean:
	rm latc