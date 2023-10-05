INCLUDEDIR = include/

CC = _CC

build: 
	@echo Compiling...
	$(CC) -o bin/main -I $(INCLUDEDIR) ./src/*.cpp

debug: 
	@echo Compiling with debug symbols...
	$(CC) --debug -o bin/main -I $(INCLUDEDIR) ./src/*.cpp

run: build
	./bin/main
