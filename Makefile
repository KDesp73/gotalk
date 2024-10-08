all: server cli index

server:
	go build ./cmd/gotalk

cli:
	go build ./cmd/gotalk-cli

index:
	cc index.c -o webc-gen -lwebc
	./webc-gen -e

clean:
	rm ./gotalk
	rm ./gotalk-cli
	rm ./webc-gen


