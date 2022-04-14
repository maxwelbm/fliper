
run:
	go build -ldflags "-s -w" -o ./bin/fliper fliper.go;
	./bin/fliper
