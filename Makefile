SRC=$(wildcard *.go)
BIN=$(patsubst %.go, %.exe, $(SRC))

all: $(BIN)

%.exe: %.go
	go build -o $@ $<


.PHONY: clean

clean:
	rm -f $(BIN)
