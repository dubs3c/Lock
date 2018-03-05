

install:
	go build -o ./bin/lock

clean:
	rm -r ./bin

test:
	go test ./...