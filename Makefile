SRC=$(wildcard *.go)
BIN=$(patsubst %.go, %.exe, $(SRC))
#GO = gccgo
GO = go

all: $(BIN)

%.exe: %.go
	$(GO) build -o $@ $<


.PHONY: clean

clean:
	rm -f $(BIN)
