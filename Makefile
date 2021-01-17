SHELL := /usr/bin/env bash
PROJ := github.com/moshen/gotermimg
ALL := gotermimg
SUBP := $(shell ls -d -1 */ | grep -v vendor | paste -sd ',' -)
SAFERM := saferm () { for f in "$$@"; do ([[ -e "$$f" ]] && rm "$$f"); done; return 0; }; saferm

all: $(ALL) 

gotermimg:
	go build -mod vendor ./cmd/gotermimg

clean:
	go clean -r -i
	$(SAFERM) $(ALL)

fmt:
	go fmt $(PROJ)/ $(PROJ)/{$(SUBP)}...
