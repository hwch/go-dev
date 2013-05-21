SRC = $(wildcard *.go)
BIN = $(patsubst %.go, %.exe, $(SRC))
#GO = gccgo
GO = go
all: $(BIN)

%.exe: %.go
	$(GO) build -o $@ $<
#	$(GO) -I./src -o $@ $<

clean:
	rm -f $(BIN)
