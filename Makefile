
build:
	go build -ldflags "-s -w" -o ./bin/fliper fliper.go;
run:
	./bin/fliper -t="${t}"