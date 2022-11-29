.PHONY: all
all: latc_x86_64

.PHONY: latc_x86_64
latc_x86_64:
	cd src; go build -o ../latc_x86_64 main.go

clean:
	rm latc_x86_64
